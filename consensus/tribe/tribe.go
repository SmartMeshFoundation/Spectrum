// Copyright 2018 The SMChain Authors
// Package tribe expand the proof-of-authority consensus engine.
/*
TODO list :
	2018-02-14:
		1、在 console 中增加 tribe API ，用来查看和调试 chief 状态
		2、初始化签名人列表，找到那个更新 tribe.status 的节点 : 暗号 TODO 9999
*/

package tribe

import (
	"bytes"
	"errors"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/SmartMeshFoundation/SMChain/accounts"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/common/hexutil"
	"github.com/SmartMeshFoundation/SMChain/consensus"
	"github.com/SmartMeshFoundation/SMChain/consensus/misc"
	"github.com/SmartMeshFoundation/SMChain/core/state"
	"github.com/SmartMeshFoundation/SMChain/core/types"
	"github.com/SmartMeshFoundation/SMChain/crypto"
	"github.com/SmartMeshFoundation/SMChain/crypto/sha3"
	"github.com/SmartMeshFoundation/SMChain/ethdb"
	"github.com/SmartMeshFoundation/SMChain/log"
	"github.com/SmartMeshFoundation/SMChain/params"
	"github.com/SmartMeshFoundation/SMChain/rlp"
	"github.com/SmartMeshFoundation/SMChain/rpc"
	"github.com/hashicorp/golang-lru"
	"fmt"
)

const (
	historyLimit = 128 // Number of recent vote snapshots to keep in memory
	// TODO 3w 并不多，放到内存是否妥当？
	// TODO 30 for test , default 30000
	signerLimit = 111
	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers

)

var (
	// TODO 是否应该缩短选举周期呢?
	epochLength = uint64(signerLimit) // Default number of blocks after which to checkpoint and reset the pending votes
	blockPeriod = uint64(15)          // Default minimum difference between two consecutive block's timestamps

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	nonceSync  = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceAsync = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// only accept sync or async flag
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidNonce = errors.New("nonce not 0x00..0 or 0xff..f")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 or 2, or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// errUnauthorized is returned if a header is signed by a non-authorized entity.
	errUnauthorized = errors.New("unauthorized")

	// errWaitTransactions is returned if an empty block is attempted to be sealed
	// on an instant chain (0 second period). It's important to refuse these as the
	// block reward is zero, so an empty block just bloats the chain... fast.
	errWaitTransactions = errors.New("waiting for transactions")
)

// sigHash returns the hash which is used as input for the proof-of-authority
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func sigHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	hasher.Sum(hash[:0])
	return hash
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, t *Tribe) (common.Address, error) {
	sigcache := t.sigcache
	// XXXX : 掐头去尾 ，约定创世区块只能指定一个签名人，因为第一个块要部署合约
	if header.Number.Uint64() == 0 {
		signer := common.Address{}
		copy(signer[:], header.Extra[extraVanity:])
		t.status.LoadSigners([]*Signer{{signer, 3}})
		return signer, nil
	}

	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(sigHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

type Tribe struct {
	config   *params.TribeConfig // Consensus engine configuration parameters
	db       ethdb.Database      // Database to store and retrieve snapshot checkpoints
	sigcache *lru.ARCCache       // 缓存 block.hash -> signer
	signer   common.Address      // Ethereum address of the signing key
	signFn   SignerFn            // Signer function to authorize hashes with
	lock     sync.RWMutex        // Protects the signer fields

	status *TribeStatus // 签名者信息和当前的全部验证条件都装在这里
}

// signers set to the ones provided by the user.
func New(config *params.TribeConfig, db ethdb.Database) *Tribe {
	debug("New TribeConfig:", config)
	//TODO 启动时初始化 tribeStatus
	status := NewTribeStatus()
	sigcache, _ := lru.NewARC(historyLimit)
	conf := *config
	conf.Epoch = epochLength
	conf.Period = blockPeriod

	return &Tribe{
		config:   &conf,
		db:       db,
		status:   status,
		sigcache: sigcache,
	}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (t *Tribe) Author(header *types.Header) (a common.Address, e error) {
	a, e = ecrecover(header, t)
	return
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (t *Tribe) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return t.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (t *Tribe) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := t.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (t *Tribe) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return consensus.ErrFutureBlock
	}
	// Nonces must be 0x00..0 or 0xff..f, zeroes enforced on checkpoints
	if !bytes.Equal(header.Nonce[:], nonceSync) && !bytes.Equal(header.Nonce[:], nonceAsync) {
		return errInvalidNonce
	}
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return t.verifyCascadingFields(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (t *Tribe) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time.Uint64()+t.config.Period > header.Time.Uint64() {
		return ErrInvalidTimestamp
	}
	if header.Difficulty.Cmp(diffNoTurn) == 0 && parent.Time.Uint64()+t.config.Period == header.Time.Uint64() {
		return ErrInvalidTimestamp
	}
	// All basic checks passed, verify the seal and return
	return t.verifySeal(chain, header, parents)
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (t *Tribe) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (t *Tribe) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return t.verifySeal(chain, header, nil)
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (t *Tribe) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, t)
	if err != nil {
		return err
	}
	log.Info("verifySeal", "number", number, "signer", signer.Hex())
	// signer 必须在合法的签名人列表中
	if !t.status.ValidateSigner(signer) {
		return errUnauthorized
	}
	// 根据 signer 算出 Difficulty 并予以校验
	difficulty := t.status.InTurn(header.Number.Int64(), signer)
	if difficulty != header.Difficulty {
		return errInvalidDifficulty
	}
	// XXXX : 这些应该是在 verifyHeader 中校验
	// XXXXXXXX : time 不能小于规定时间间隔，而且 Difficulty == no-turn 时时间必须大于规定时间 500ms 以上
	// XXXXXXXX : 同一个 signer 不能连续出块
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (t *Tribe) Prepare(chain consensus.ChainReader, header *types.Header) error {
	// TODO : 初始的签名人只为打包 1 块，并部署一个 chief 合约，用来做后续的共识 >>>>
	// add by liangc : 满足出块条件时去调用 chief.update , 此任务先延后一些
	if chain.CurrentHeader().Number.Int64() > 1000 {
		rtn := make(chan params.MBoxSuccess)
		params.SendToMsgBox("Update", rtn)
		r := <-rtn
		fmt.Println("--- worker_update_chainHeadCh : chief_update --->", r)
	}
	// TODO : 初始的签名人只为打包 1 块，并部署一个 chief 合约，用来做后续的共识 <<<<

	number := header.Number.Uint64()
	//TODO tribe : 提名一个符合要求的邻居节点到候选人列表
	header.Coinbase = common.Address{}
	//TODO tribe : **** 是否需要同步，要看签名人列表有没有变化，这里是个难题，如何提前预测？
	// 按照当前得分看，如果这个块不是我的，那么应该出块的人如果等于1分，则预言此处为 SYNC
	header.Nonce = types.BlockNonce{}
	copy(header.Nonce[:], nonceAsync)

	//TODO tribe : Extra 存放签名信息, 规则需要重新定 : 暂时先用 POA 规则 >>>>
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)
	//TODO tribe : Extra 存放签名信息, 规则需要重新定 : 暂时先用 POA 规则 <<<<

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = new(big.Int).Add(parent.Time, new(big.Int).SetUint64(t.config.Period))
	if header.Time.Int64() < time.Now().Unix() {
		header.Time = big.NewInt(time.Now().Unix())
	}

	// tribe : in-turn signers[ block.number % len(signers) ] == currentSigner
	// Set the correct difficulty
	header.Difficulty = t.CalcDifficulty(chain, header.Time.Uint64(), parent)
	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (t *Tribe) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// No block rewards in Tribe, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
// TODO called by : func (s *Ethereum) StartMining(local bool) error
// 在上面那个方法中 通过 Etherbase() 得到 signer ,就是 etherbase
func (t *Tribe) Authorize(signer common.Address, signFn SignerFn) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.signer = signer
	t.signFn = signFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (t *Tribe) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	header := block.Header()
	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return nil, errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	// TODO ???
	if t.config.Period == 0 && len(block.Transactions()) == 0 {
		return nil, errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	t.lock.RLock()
	signer, signFn := t.signer, t.signFn
	t.lock.RUnlock()

	// 校验 signer 是否在最新的 signers 列表中
	if !t.status.ValidateSigner(t.signer) {
		return nil, errUnauthorized
	}

	delay := time.Unix(header.Time.Int64(), 0).Sub(time.Now())
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		wiggle := time.Duration(len(t.status.Signers)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))
	}
	select {
	case <-stop:
		return nil, nil
	case <-time.After(delay):
	}
	// Sign all the things!
	hash := sigHash(header).Bytes()
	sighash, err := signFn(accounts.Account{Address: signer}, hash)
	if err != nil {
		return nil, err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	return block.WithSeal(header), nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (t *Tribe) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	//TODO 通过 parent 块获取最新的 signers 并计算 difficulty
	// 并满足如下公式
	// in-turn : signers[ block.number % len(signers) ] == t.signer
	// no-turn : signers[ block.number % len(signers) ] != t.signer
	return t.status.InTurn(parent.Number.Int64()+1, t.signer)
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (t *Tribe) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "tribe",
		Version:   "1.0",
		Service:   &API{chain: chain, tribe: t},
		Public:    false,
	}}
}

func debug(msg ... interface{}) {
	msg = append([]interface{}{"<< tribe : liangc debug >> :"}, msg...)
	fmt.Println(msg ...)
}
