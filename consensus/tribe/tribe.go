// Copyright 2018 The Spectrum Authors
package tribe

import (
	"bytes"
	"errors"
	"math/big"
	"math/rand"
	"time"

	"crypto/ecdsa"
	"fmt"
	"sync/atomic"

	"github.com/SmartMeshFoundation/Spectrum/accounts"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/consensus"
	"github.com/SmartMeshFoundation/Spectrum/consensus/misc"
	"github.com/SmartMeshFoundation/Spectrum/core/state"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/crypto/sha3"
	"github.com/SmartMeshFoundation/Spectrum/ethdb"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rlp"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
	"github.com/hashicorp/golang-lru"
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
	// XXXX : 掐头去尾 ，约定创世区块只能指定一个签名人，因为第一个块要部署合约
	if header.Number.Uint64() == 0 {
		signer := common.Address{}
		copy(signer[:], header.Extra[extraVanity:])
		t.Status.loadSigners([]*Signer{{signer, 3}})
		return signer, nil
	}
	sigcache := t.sigcache
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

// signers set to the ones provided by the user.
func New(config *params.TribeConfig, db ethdb.Database) *Tribe {
	status := NewTribeStatus()
	sigcache, _ := lru.NewARC(historyLimit)
	conf := *config
	if conf.Period <= 0 {
		conf.Period = blockPeriod
	}
	tribe := &Tribe{
		config:      &conf,
		db:          db,
		Status:      status,
		sigcache:    sigcache,
		SealErrorCh: make(chan error, 1),
	}
	status.setTribe(tribe)
	return tribe
}

func (t *Tribe) Init(hash common.Hash, number *big.Int) {
	go func() {
		log.Info("init tribe.status when chiefservice start end.")
		<-params.InitTribeStatus
		t.Status.LoadSignersFromChief(hash, number)
		rtn := params.SendToMsgBox("GetNodeKey")
		success := <-rtn
		t.Status.nodeKey = success.Entity.(*ecdsa.PrivateKey)
		log.Info("init tribe.status success.")
	}()
}

// called by miner.start
func (t *Tribe) WaitingNomination() chan struct{} {
	ch := make(chan struct{})
	go func() {
		for {
			if t.Status.SignerLevel != LevelNone {
				ch <- struct{}{}
				return
			}
			//TODO
			fmt.Println(":: WaitingNomination ::> ", t.Status.GetMinerAddress().Hex(), t.Status.SignerLevel)
			<-time.After(time.Second * 3)
		}
	}()
	return ch
}

// called by worker.start and worker.stop
// 1 mining start
// 0 mining stop
func (t *Tribe) SetMining(i int32, currentNumber *big.Int, currentBlockHash common.Hash) {
	log.Info("tribe.setMining", "mining", i)
	t.lock.Lock()
	log.Debug("tribe.setMining_lock", "mining", i)
	defer t.lock.Unlock()
	atomic.StoreInt32(&t.Status.mining, i)
	if i == 1 {
		if currentNumber.Int64() >= CHIEF_NUMBER {
			fmt.Println(currentNumber.Int64(), "><> tribe.SetMining -> Status.Update : may be pending")
			t.Status.Update(currentNumber, currentBlockHash)
			fmt.Println(currentNumber.Int64(), "><> tribe.SetMining -> Status.Update : done")
		}
	}
	log.Debug("tribe.setMining_unlock", "mining", i)
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
	// TODO ********** 一个临时方案，此处需要优化，最终目标是 async / sync 通过 header.nonce 来同步
	results := make(chan error)
	log.Debug("==> VerifyHeaders ", "currentNum", chain.CurrentHeader().Number.Int64(), "headers.len", len(headers))
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
	if number > 0 && !params.IsChiefBlock(header.Number) {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			log.Error("** verifyHeader ERROR **", "diff", header.Difficulty.String(), "err", errInvalidDifficulty)
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
func (t *Tribe) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) (err error) {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
		if parent.Time.Uint64()+t.config.Period > header.Time.Uint64() {
			return ErrInvalidTimestamp
		}
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
		if parent.Time.Uint64()+t.config.Period > header.Time.Uint64() {
			return ErrInvalidTimestamp
		}
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	/* TODO timestamp unit second , how to do this logic ?
	if header.Difficulty.Cmp(diffNoTurn) == 0 && parent.Time.Uint64()+t.config.Period == header.Time.Uint64() {
		return ErrInvalidTimestamp
	}
	*/
	return
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
	e := t.verifySeal(chain, header, nil)
	if e != nil {
		log.Error("Tribe.VerifySeal", "err", e)
	}
	return e
}

func (t *Tribe) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Int64()
	if number == 0 {
		return errUnknownBlock
	}
	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, t)
	if err != nil {
		return err
	}
	log.Debug("verifySeal", "number", number, "signer", signer.Hex())
	// signer 必须在合法的签名人列表中
	if !t.Status.ValidateSigner(number, header.ParentHash, signer) {
		return errUnauthorized
	}
	// 根据 signer 算出 Difficulty 并予以校验
	if number > 3 && !params.IsChiefBlock(header.Number) {
		difficulty := t.Status.InTurnForVerify(number, header.ParentHash, signer)
		if difficulty.Cmp(header.Difficulty) != 0 {
			log.Error("** verifySeal ERROR **", "diff", header.Difficulty.String(), "err", errInvalidDifficulty)
			log.Error("??is_chiefBlock??", "chiefBlock", params.IsChiefBlock(header.Number), "num", header.Number.Int64())
			return errInvalidDifficulty
		}
	}
	// XXXX : 这些应该是在 verifyHeader 中校验
	// XXXXXXXX : time 不能小于规定时间间隔，而且 Difficulty == no-turn 时时间必须大于规定时间 500ms 以上
	// XXXXXXXX : 某些条件下同一个 signer 不能连续出块
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (t *Tribe) Prepare(chain consensus.ChainReader, header *types.Header) error {
	number := header.Number.Uint64()
	//header.Coinbase = common.Address{}
	// 没有实际左右，只是为了观察日志，后续会去掉
	header.Coinbase = t.Status.GetMinerAddress()
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
	if number > 3 {
		header.Difficulty = t.CalcDifficulty(chain, header.Time.Uint64(), parent)
	} else {
		header.Difficulty = diffInTurn
	}
	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (t *Tribe) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	//fmt.Println("-- Tribe.Finalize --> txs :",txs)
	// No block rewards in Tribe, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (t *Tribe) Authorize(signer common.Address, signFn SignerFn) {
	t.lock.Lock()
	defer t.lock.Unlock()
	prv := t.Status.nodeKey
	//t.signer = t.Status.GetMinerAddress()
	t.signFn = func(a accounts.Account, hex []byte) ([]byte, error) {
		return crypto.Sign(hex, prv)
	}
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (t *Tribe) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	if err := t.Status.ValidateBlock(block,false); err != nil {
		log.Error("Tribe_Seal", "retry", atomic.LoadUint32(&t.SealErrorCounter), "number", block.Number().Int64(), "err", err)
		t.SealErrorCh <- err
		return nil, err
	}
	atomic.StoreUint32(&t.SealErrorCounter, 0)
	header := block.Header()
	// Sealing the genesis block is not supported
	number := header.Number.Int64()
	if number == 0 {
		if genesisSigner, e := t.Status.genesisSigner(header); e == nil && genesisSigner == t.Status.GetMinerAddress() {
			t.Status.SignerLevel = LevelSigner
		}
		return nil, errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	// TODO ???
	if t.config.Period == 0 && len(block.Transactions()) == 0 {
		return nil, errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	t.lock.RLock()
	signer, signFn := t.Status.GetMinerAddress(), t.signFn
	//signer, signFn := t.signer, t.signFn
	t.lock.RUnlock()

	// 校验 signer 是否在最新的 signers 列表中
	if !t.Status.ValidateSigner(number, block.ParentHash(), t.Status.GetMinerAddress()) {
		return nil, errUnauthorized
	}

	delay := time.Unix(header.Time.Int64(), 0).Sub(time.Now())
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		wiggle := time.Duration(len(t.Status.Signers)/2+1) * wiggleTime
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
	blk := block.WithSeal(header)
	return blk,nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (t *Tribe) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	log.Debug("CalcDifficulty", "ParentNumber", parent.Number.Int64(), "CurrentNumber:", chain.CurrentHeader().Number.Int64())
	//TODO 通过 parent 块获取最新的 signers 并计算 difficulty
	// 并满足如下公式
	// in-turn : signers[ block.number % len(signers) ] == t.signer
	// no-turn : signers[ block.number % len(signers) ] != t.signer
	return t.Status.InTurnForCalc(t.Status.GetMinerAddress(), parent)
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (t *Tribe) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "tribe",
		Version:   "0.0.1",
		Service:   &API{chain: chain, tribe: t},
		Public:    false,
	}}
}
