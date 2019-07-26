// Copyright 2018 The Spectrum Authors
package tribe

import (
	"bytes"
	"encoding/hex"
	"errors"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"crypto/ecdsa"
	"fmt"
	"sync/atomic"

	"github.com/SmartMeshFoundation/Spectrum/accounts"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/common/math"
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
	lru "github.com/hashicorp/golang-lru"
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

func ecrecoverPubkey(header *types.Header, signature []byte) ([]byte, error) {
	pubkey, err := crypto.Ecrecover(sigHash(header).Bytes(), signature)
	return pubkey, err
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, t *Tribe) (common.Address, error) {
	// XXXX : 掐头去尾 ，约定创世区块只能指定一个签名人，因为第一个块要部署合约
	extraVanity := extraVanityFn(header.Number)
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
	// Recover the public key and the Ethereum address
	pubkey, err := ecrecoverPubkey(header, header.Extra[len(header.Extra)-extraSeal:])

	//pubkey, err := crypto.Ecrecover(sigHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)

	return signer, nil
}

// signers set to the ones provided by the user.
func New(accman *accounts.Manager, config *params.TribeConfig, _ ethdb.Database) *Tribe {
	status := NewTribeStatus()
	sigcache, _ := lru.NewARC(historyLimit)
	conf := *config
	if conf.Period <= 0 {
		conf.Period = blockPeriod
	}
	tribe := &Tribe{
		accman:   accman,
		config:   &conf,
		Status:   status,
		sigcache: sigcache,
		//SealErrorCh: make(map[int64]error),
		SealErrorCh: new(sync.Map),
	}
	status.setTribe(tribe)
	return tribe
}

func (t *Tribe) Init(hash common.Hash, number *big.Int) {
	go func() {
		t.lock.Lock()
		defer t.lock.Unlock()
		if t.isInit {
			return
		}
		<-params.InitTribeStatus
		rtn := params.SendToMsgBox("GetNodeKey")
		success := <-rtn
		t.Status.nodeKey = success.Entity.(*ecdsa.PrivateKey)
		if params.InitTribe != nil {
			close(params.InitTribe)
			params.InitTribe = nil
		}
		log.Debug("init tribe.status when chiefservice start end.", "getnodekey", success.Success)
		if number.Int64() >= CHIEF_NUMBER {
			t.isInit = true
			t.Status.LoadSignersFromChief(hash, number)
		}
		log.Info("init tribe.status success.")
	}()
}

// called by miner.start
func (t *Tribe) WaitingNomination() {
	for {
		if t.Status.SignerLevel != LevelNone {
			return
		}
		<-time.After(time.Second * time.Duration(t.config.Period/2)) //每半块检查一次,这样保证等待时间不超过一块
	}
	return
}

// called by worker.start and worker.stop
// 1 mining start
// 0 mining stop
func (t *Tribe) SetMining(i int32, currentNumber *big.Int, currentBlockHash common.Hash) {
	log.Info("tribe.setMining", "mining", i)
	log.Debug("tribe.setMining_lock", "mining", i)
	atomic.StoreInt32(&t.Status.mining, i)
	if i == 1 {
		if currentNumber.Int64() >= CHIEF_NUMBER {
			log.Debug("><> tribe.SetMining -> Status.Update : may be pending")
			t.Status.Update(currentNumber, currentBlockHash)
			log.Debug("><> tribe.SetMining -> Status.Update : done")
		}
	}
	log.Debug("tribe.setMining_unlock", "mining", i)
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (t *Tribe) Author(header *types.Header) (common.Address, error) {
	var (
		coinbase common.Address
		err      error
	)
	if header.Coinbase == coinbase {
		coinbase, err = ecrecover(header, t)
	} else {
		coinbase = header.Coinbase
	}
	log.Debug("<<Tribe.Author>>", "num", header.Number, "coinbase", coinbase.Hex())
	return coinbase, err
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (t *Tribe) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	err := t.verifyHeader(chain, header, nil)
	if err == nil {
		//p := chain.GetHeaderByHash(header.ParentHash)
		//t.Status.LoadSignersFromChief(p.Hash(), p.Number)
	}
	return err
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (t *Tribe) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
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
func (t *Tribe) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) (err error) {
	defer func() {
		if err != nil {
			log.Info(fmt.Sprintf("verifyHeader err return number=%s,err=%s", header.Number, err))
		}
	}()
	extraVanity := extraVanityFn(header.Number)
	//if !t.isInit && header.Number.Int64() >= CHIEF_NUMBER {
	//	log.Info(fmt.Sprintf("verifyHeader init, header=%#v,parents=%#v", header, parents))
	//	t.Init(header.Hash(), header.Number)
	//}

	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return consensus.ErrFutureBlock
	}
	// Nonces must be 0x00..0 or 0xff..f, zeroes enforced on checkpoints
	/*
		if !bytes.Equal(header.Nonce[:], nonceSync) && !bytes.Equal(header.Nonce[:], nonceAsync) {
			return errInvalidNonce
		}*/

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
	if number > 0 && !params.IsBeforeChief100block(header.Number) {
		if ci := params.GetChiefInfo(header.Number); ci != nil {
			switch ci.Version {
			case "1.0.0":
				//TODO max value is a var ???
				if header.Difficulty == nil || header.Difficulty.Cmp(big.NewInt(6)) > 0 || header.Difficulty.Cmp(diffNoTurn) < 0 {
					log.Error("** verifyHeader ERROR CHIEF.v1.0.0 **", "diff", header.Difficulty.String(), "err", errInvalidDifficulty)
					return errInvalidDifficulty
				}
			default:
				if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurnMain) != 0 && header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
					log.Error("** verifyHeader ERROR **", "diff", header.Difficulty.String(), "err", errInvalidDifficulty)
					return errInvalidDifficulty
				}
			}
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	err = t.verifyCascadingFields(chain, header, parents)
	if err != nil {
		log.Error("verifyCascadingFields", "num", header.Number.Int64(), "err", err)
	}
	return err
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

	verifyTime := func() error {
		if params.IsSIP002Block(header.Number) {
			// first verification
			// second verification block time in validateSigner function
			// the min limit period is config.Period - 1
			if parent.Time.Uint64()+(t.config.Period-1) > header.Time.Uint64() {
				return ErrInvalidTimestampSIP002
			}
		} else {
			if parent.Time.Uint64()+t.config.Period > header.Time.Uint64() {
				return ErrInvalidTimestamp
			}
		}
		return nil
	}

	if len(parents) > 0 {
		parent = parents[len(parents)-1]
		if err := verifyTime(); err != nil {
			return err
		}
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
		if parent == nil || parent.Time == nil {
			e := errors.New(fmt.Sprintf("nil_parent_current_num=%d", header.Number.Int64()))
			log.Error("-->bad_block-->", "err", e)
			return e
		}
		if err := verifyTime(); err != nil {
			return err
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

	// Verify that the gas limit is <= 2^63-1
	if header.GasLimit.Cmp(math.MaxBig63) > 0 {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, math.MaxBig63)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed.Cmp(header.GasLimit) > 0 {
		return fmt.Errorf("invalid gasUsed: have %v, gasLimit %v", header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	diff := new(big.Int).Set(parent.GasLimit)
	diff = diff.Sub(diff, header.GasLimit)
	diff.Abs(diff)

	limit := new(big.Int).Set(parent.GasLimit)
	limit = limit.Div(limit, params.GasLimitBoundDivisor)

	if diff.Cmp(limit) >= 0 || header.GasLimit.Cmp(params.MinGasLimit) < 0 {
		return fmt.Errorf("invalid gas limit: have %v, want %v += %v", header.GasLimit, parent.GasLimit, limit)
	}

	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}

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

	if !t.Status.validateSigner(chain.GetHeaderByHash(header.ParentHash), header, signer) {
		return errUnauthorized
	}

	if number > 3 && !params.IsBeforeChief100block(header.Number) {
		difficulty := t.Status.InTurnForVerify(number, header.ParentHash, signer)
		if difficulty.Cmp(header.Difficulty) != 0 {
			log.Error("** verifySeal ERROR **", "diff", header.Difficulty.String(), "err", errInvalidDifficulty)
			return errInvalidDifficulty
		}
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (t *Tribe) Prepare(chain consensus.ChainReader, header *types.Header) error {
	number := header.Number.Uint64()
	if f, _, err := params.AnmapBindInfo(t.Status.GetMinerAddress(), chain.CurrentHeader().Hash()); err == nil && f != common.HexToAddress("0x") {
		header.Coinbase = f
	} else {
		header.Coinbase = t.Status.GetMinerAddress()
	}
	header.Nonce = types.BlockNonce{}
	copy(header.Nonce[:], nonceAsync)

	// Extra : append sig to last 65 bytes >>>>
	if params.IsSIP100Block(header.Number) {
		header.Extra = make([]byte, _extraVrf+extraSeal)
		// append vrf to header.Extra before sign
		parentHeader := chain.GetHeaderByHash(header.ParentHash)
		msg := append(parentHeader.Number.Bytes(), parentHeader.Extra[:32]...)
		vrfnp, err := crypto.SimpleVRF2Bytes(t.Status.nodeKey, msg)
		log.Debug("Tribe.Prepare --> params.GetVRFByHash", "err", err, "hash", header.ParentHash.Hex(), "vrfn", hex.EncodeToString(vrfnp[:32]))
		//vr, err := crypto.SimpleVRF2Bytes(t.Status.nodeKey, header.ParentHash.Bytes())
		if err != nil {
			panic(err)
		}

		copy(header.Extra[:len(header.Extra)-extraSeal], vrfnp)
		log.Debug("prepare_vrf", "num", header.Number, "parent_hash", header.ParentHash.Hex())
		log.Debug("prepare_vrf", "num", header.Number, "miner", crypto.PubkeyToAddress(t.Status.nodeKey.PublicKey).Hex())
		log.Debug("prepare_vrf", "num", header.Number, "err", err, "vrf", hex.EncodeToString(vrfnp))
		log.Debug("prepare_vrf", "num", header.Number, "extra", hex.EncodeToString(header.Extra))
	} else {
		extraVanity := extraVanityFn(header.Number)
		log.Debug("fix extra", "extra-len", len(header.Extra), "extraVanity", extraVanity)
		if len(header.Extra) < extraVanity {
			header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
		}
		header.Extra = header.Extra[:extraVanity]
		header.Extra = append(header.Extra, make([]byte, extraSeal)...)
	}
	// Extra : append sig to last 65 bytes <<<<

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	// tribe : in-turn signers[ block.number % len(signers) ] == currentSigner
	// Set the correct difficulty
	if number > 3 {
		header.Difficulty = t.CalcDifficulty(chain, header.Time.Uint64(), parent)
		/*
			按照sip100共识,替补只能是leader节点,如果该块不该我出,同时我也不是leader,那么就应该放弃出块
		*/
		if params.IsSIP100Block(header.Number) && header.Difficulty.Cmp(diffNoTurn) <= 0 {
			if !t.Status.IsLeader(t.Status.GetMinerAddress()) {
				return fmt.Errorf("%s is not signer and not leader at block %s,diff=%s", t.Status.GetMinerAddress().String(), header.Number, header.Difficulty)
			}
		}
	} else {
		header.Difficulty = diffInTurn
	}
	if params.IsSIP002Block(header.Number) {
		//modify by liangc : change period rule
		header.Time = new(big.Int).Add(parent.Time, new(big.Int).SetUint64(t.GetPeriod(header, nil)))
	} else {
		header.Time = new(big.Int).Add(parent.Time, new(big.Int).SetUint64(t.config.Period))
	}
	if header.Time.Int64() < time.Now().Unix() {
		header.Time = big.NewInt(time.Now().Unix())
	}
	return nil
}

var chiefGasLimit = big.NewInt(4712388)
var chiefGasPrice = big.NewInt(18000000000)

/*
获取要插入的chief Tx
*/
func (t *Tribe) GetChief100UpdateTx(chain consensus.ChainReader, header *types.Header, state *state.StateDB) *types.Transaction {
	parentHash := header.ParentHash
	parentNumber := new(big.Int).Set(header.Number)
	parentNumber.Sub(parentNumber, big.NewInt(1))
	if !params.IsSIP100Block(parentNumber) {
		return nil // sip100之前的update调用还是走老办法
	}
	vrf := new(big.Int).SetBytes(header.Extra[:32])
	nextRoundSigner := params.Chief100GetNextRoundSigner(parentHash, parentNumber, vrf)
	nonce := state.GetNonce(t.Status.GetMinerAddress())
	txData, _ := hex.DecodeString("1c1b8772000000000000000000000000")
	txData = append(txData, nextRoundSigner[:]...)
	rawTx := types.NewTransaction(nonce, params.GetChiefInfo(parentNumber).Addr, big.NewInt(0), chiefGasLimit, chiefGasPrice, txData)
	auth := bind.NewKeyedTransactor(t.Status.getNodekey())
	signedTx, err := auth.Signer(types.NewEIP155Signer(chain.Config().ChainId), auth.From, rawTx)
	if err != nil {
		panic(fmt.Sprintf("sign tx:%s", err))
	}
	return signedTx
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (t *Tribe) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Accumulate any block and uncle rewards and commit the final state root
	accumulateRewards(chain.Config(), state, header)
	//destroy SmartMeshFoundation 12% balance
	destroySmartMeshFoundation12Balance(chain.Config(), state, header)

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	//there is no uncle in triple
	header.UncleHash = types.CalcUncleHash(nil)
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (t *Tribe) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	if err := t.Status.ValidateBlock(nil, chain.GetBlock(block.ParentHash(), block.NumberU64()-1), block, false); err != nil {
		log.Error("Tribe_Seal", "number", block.Number().Int64(), "err", err)
		//log.Error("Tribe_Seal", "retry", atomic.LoadUint32(&t.SealErrorCounter), "number", block.Number().Int64(), "err", err)
		//t.SealErrorCh[chain.CurrentHeader().Number.Int64()] = err
		t.SealErrorCh.Store(chain.CurrentHeader().Number.Int64(), err) //因为没有收集到update调用的失败,所以再次尝试
		return nil, err
	}
	//atomic.StoreUint32(&t.SealErrorCounter, 0)
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
	// TODO liangc : How to stop the empty block ???
	if t.config.Period == 0 && len(block.Transactions()) == 0 {
		return nil, errWaitTransactions
	}
	if !t.Status.validateSigner(chain.GetHeaderByHash(block.ParentHash()), block.Header(), t.Status.GetMinerAddress()) {
		return nil, errUnauthorized
	}

	now := time.Now()
	delay := time.Unix(header.Time.Int64(), 0).Sub(now)
	fmt.Println(" ---->", "diff", header.Difficulty, "header.time=", header.Time.Int64(), "now=", now.Unix(), "delay=", delay)

	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		wiggle := time.Duration(len(t.Status.Signers)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))
	}
	select {
	case <-stop:
		log.Warn(fmt.Sprintf("🐦 cancel -> num=%d, diff=%d, miner=%s, delay=%d", number, header.Difficulty, header.Coinbase.Hex(), delay))
		return nil, nil
	case <-time.After(delay):
	}

	// Sign all the things!
	hash := sigHash(header).Bytes()
	sighash, err := crypto.Sign(hash, t.Status.nodeKey)
	if err != nil {
		return nil, err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	blk := block.WithSeal(header)
	return blk, nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (t *Tribe) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	log.Debug("CalcDifficulty", "ParentNumber", parent.Number.Int64(), "CurrentNumber:", chain.CurrentHeader().Number.Int64())
	currentNumber := new(big.Int).Add(chain.CurrentHeader().Number, big.NewInt(1))
	if ci := params.GetChiefInfo(currentNumber); ci != nil {
		switch ci.Version {
		case "1.0.0":
			return t.Status.InTurnForCalcChief100(t.Status.GetMinerAddress(), parent)
		}
	}
	return t.Status.InTurnForCalc(t.Status.GetMinerAddress(), parent)
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (t *Tribe) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "tribe",
		Version:   "0.0.1",
		Service:   &API{accman: t.accman, chain: chain, tribe: t},
		Public:    false,
	}}
}

/*
GetPeriodChief100: 计算出块节点对应的延时时间
依据:
普通节点替补出块规则调整为当前常委会出块节点替补出块，当前常委会出块节点替补出块规则调整为常委会节点列表顺序替补出块，出块难度依次递减。
当普通节点出块时，难度为6->5->4->3->2->1；当常委会节点出块时，难度为6->5->4->3->2。每次替补出块时间都顺延4秒，
下一个区块将在当前区块产生以后的14秒+ 以后产生。
总共只有5个常委会节点,其排序固定
signers:[0,...,16] 0号对应的是常委会节点,1-16对应的是普通出块节点
场景1:
假设`header`中的块该signers[3]出,
如果signers[3]出,则延时14秒,否则只能由其他常委会节点替代
假设当前signers[0]对应的是常委2,那么如果常委2出则延时18,常委3出则延时22,...,如果常委1出则延时34
场景2:
假设`header`中的块该signers[0]出,
如果signers[0]出,则延时14秒,否则只能由其他常委会节点替代
假设当前signers[0]对应的是常委2,那么常委3出则延时18,常委4出则延时22,...常委1出则延时30
*/
func (t *Tribe) GetPeriodChief100(header *types.Header, signers []*Signer) (p uint64) {
	var (
		// 14 , 18 , 22
		signature               = header.Extra[len(header.Extra)-extraSeal:]
		err                     error
		miner                   common.Address
		empty                   = make([]byte, len(signature))
		Main, Subs, Other, Else = t.config.Period - 1, t.config.Period + 3, t.config.Period + 7, uint64(86400 * 7)
		number, parentHash      = header.Number, header.ParentHash
	)

	p = Else
	if bytes.Equal(signature, empty) {
		miner = t.Status.GetMinerAddress()
	} else {
		miner, _ = ecrecover(header, t)
	}

	if signers == nil {
		signers, err = t.Status.GetSignersFromChiefByHash(parentHash, number)
	}

	if err != nil {
		log.Error("GetPeriod_getsigners_err", "err", err)
		return
	}

	sl := len(signers)
	if sl == 0 {
		log.Error("GetPeriod_signers_cannot_empty")
		return
	}

	// normal
	idx_m := number.Int64() % int64(sl)
	if miner == signers[idx_m].Address {
		p = Main
		return
	}

	// first leader
	if miner == signers[0].Address {
		p = Subs
		return
	}

	// other leader
	if leaders, err := leaderSort(signers[0].Address, t.Status.Leaders); err == nil {
		for i, leader := range leaders {
			if miner == leader && number.Int64()%int64(sl) == 0 {
				p = Subs + uint64(i)*(Subs-Main)
				return
			} else if miner == leader {
				p = Other + uint64(i)*(Subs-Main)
				return
			}
		}
	}

	return
}

func (t *Tribe) GetPeriod(header *types.Header, signers []*Signer) (p uint64) {
	if ci := params.GetChiefInfo(header.Number); ci != nil {
		switch ci.Version {
		case "1.0.0":
			p = t.GetPeriodChief100(header, signers)
			log.Debug("<<GetPeriodSIP005>>", "num", header.Number, "period", p)
			return
		}
	}
	// 14 , 18 , 22(random add 0~4.5s)
	signature := header.Extra[len(header.Extra)-extraSeal:]
	var (
		err   error
		miner common.Address
		empty = make([]byte, len(signature))
		/*
			替补出块按照顺序依次增加4秒的延迟
		*/
		Main, Subs, Other  = t.config.Period - 1, t.config.Period + 3, t.config.Period + 7
		number, parentHash = header.Number, header.ParentHash
	)

	p = Other
	if bytes.Equal(signature, empty) {
		miner = t.Status.GetMinerAddress()
	} else {
		miner, _ = ecrecover(header, t)
	}

	if number.Int64() <= 3 {
		p = Subs
		return
	}

	if signers == nil {
		signers, err = t.Status.GetSignersFromChiefByHash(parentHash, number)
	}

	if err != nil {
		log.Error("GetPeriod_getsigners_err", "err", err)
		p = Other
		return
	}

	sl := len(signers)
	if sl == 0 {
		log.Error("GetPeriod_signers_cannot_empty")
		p = Other
		return
	}

	idx_m, idx_s := number.Int64()%int64(sl), (number.Int64()+1)%int64(sl)

	if miner == signers[idx_m].Address {
		p = Main
		return
	}

	if miner == signers[idx_s].Address {
		p = Subs
		return
	}

	return
}

// AccumulateRewards credits the coinbase of the given block with the mining
// reward. The total reward consists of the static block reward and rewards for
// included uncles. The coinbase of each uncle block is also rewarded.
// add by liangc : no reward
func accumulateRewards(config *params.ChainConfig, state *state.StateDB, header *types.Header) {
	if config.Chief100Block == nil || header.Number.Cmp(config.Chief100Block) < 0 {
		return
	}
	// Select the correct block reward based on chain progression
	blockReward := new(big.Int).Set(Chief100BlockReward)

	// Accumulate the rewards for the miner and any included uncles
	number := new(big.Int).Set(header.Number)
	number = number.Sub(number, config.Chief100Block)
	number = number.Div(number, big.NewInt(int64(BlockRewardReducedInterval)))
	blockReward = blockReward.Rsh(blockReward, uint(number.Int64()))
	state.AddBalance(header.Coinbase, blockReward)
}

//
//销毁基金会账户12% token
//Destroy 12% token of Foundation Account

func destroySmartMeshFoundation12Balance(config *params.ChainConfig, state *state.StateDB, header *types.Header) {
	if !params.IsDevnet() && !params.IsTestnet() && config.Chief100Block != nil && header.Number.Cmp(config.Chief100Block) == 0 {
		ob := state.GetBalance(SmartMeshFoundationAccount)
		if ob.Cmp(SmartMeshFoundationAccountDestroyBalance) >= 0 {
			state.SubBalance(SmartMeshFoundationAccount, SmartMeshFoundationAccountDestroyBalance)
		} else {
			state.SubBalance(SmartMeshFoundationAccount, ob)
		}

	}
}
