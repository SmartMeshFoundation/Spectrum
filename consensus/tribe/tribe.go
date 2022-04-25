// Copyright 2018 The mesh-chain Authors
package tribe

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/MeshBoxTech/mesh-chain/consensus/tribe/vmcaller"
	"math/big"
	"math/rand"
	"sort"
	"time"

	"github.com/MeshBoxTech/mesh-chain/accounts"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/common/math"
	"github.com/MeshBoxTech/mesh-chain/consensus"
	"github.com/MeshBoxTech/mesh-chain/consensus/misc"
	"github.com/MeshBoxTech/mesh-chain/core/state"
	"github.com/MeshBoxTech/mesh-chain/core/types"
	"github.com/MeshBoxTech/mesh-chain/crypto"
	"github.com/MeshBoxTech/mesh-chain/crypto/sha3"
	"github.com/MeshBoxTech/mesh-chain/ethdb"
	"github.com/MeshBoxTech/mesh-chain/log"
	"github.com/MeshBoxTech/mesh-chain/params"
	"github.com/MeshBoxTech/mesh-chain/rlp"
	"github.com/MeshBoxTech/mesh-chain/rpc"
	lru "github.com/hashicorp/golang-lru"

	sha33 "golang.org/x/crypto/sha3"

	"github.com/holiman/uint256"

)

// sigHash returns the hash which is used as input for the proof-of-authority
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.

// StateFn gets state by the state root hash.
type StateFn func(hash common.Hash) (*state.StateDB, error)

func sigHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewKeccak256()

	err := rlp.Encode(hasher, []interface{}{
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
	if err != nil {
		panic(err)
	}
	hasher.Sum(hash[:0])
	return hash
}

func ecrecoverPubkey(header *types.Header, signature []byte) ([]byte, error) {
	pubkey, err := crypto.Ecrecover(sigHash(header).Bytes(), signature)
	return pubkey, err
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, t *Tribe) (common.Address, error) {
	sigcache := t.sigcache

	// If the signature's already cached, return that
	hash := header.Hash()
	if sigcache != nil {
		if address, known := sigcache.Get(hash); known {
			return address.(common.Address), nil
		}
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	// Recover the public key and the Ethereum address
	pubkey, err := ecrecoverPubkey(header, header.Extra[len(header.Extra)-extraSeal:])
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
	if sigcache != nil {
		sigcache.Add(hash, signer)
	}

	return signer, nil
}

// signers set to the ones provided by the user.
func New(accman *accounts.Manager, config *params.TribeConfig, db ethdb.Database) *Tribe {
	sigcache, err := lru.NewARC(historyLimit)
	recents, _ := lru.NewARC(historyLimit)
	if err != nil {
		panic(err)
	}
	conf := *config
	if conf.Period <= 0 {
		conf.Period = blockPeriod
	}
	tribe := &Tribe{
		accman:   accman,
		config:   &conf,
		sigcache: sigcache,
		recents:  recents,
		db:db,
	}
	return tribe
}

func (t *Tribe) Init(fn StateFn,nodekey *ecdsa.PrivateKey) {
	    t.nodeKey= nodekey
		t.stateFn = fn
		t.abi = GetInteractiveABI()
		t.isInit = true
		log.Info("init tribe.status success.")
}
func (t *Tribe) GetConfig() *params.TribeConfig {
	return t.config
}
func (t *Tribe) SetConfig(config *params.TribeConfig) {
	t.config = config
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
	return t.verifyHeader(chain, header, nil)
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
	// check extra data
	isEpoch := number%t.config.Epoch == 0
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !isEpoch && signersBytes != 0 {
		return errExtraValidators
	}

	if isEpoch && (signersBytes-extraVrf)%validatorBytesLength != 0 {
		return errInvalidSpanValidators
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
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
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	//验证区块时间
	snap, err := t.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	err = t.blockTimeVerify(snap, header, parent)
	if err != nil {
		return err
	}

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

	minGasLimit := params.MinGasLimit



	if diff.Cmp(limit) >= 0 || header.GasLimit.Cmp(minGasLimit) < 0 {
		return fmt.Errorf("invalid gas limit: have %v, want %v += %v", header.GasLimit, parent.GasLimit, limit)
	}

	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}

	return t.verifySeal(chain,header,parents)
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
// don't support remote miner agent, these code never reached
func (t *Tribe) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	e := t.verifySeal(chain, header, nil)
	if e != nil {
		log.Error("Tribe.VerifySeal", "err", e)
	}
	return e
}

func (t *Tribe) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := t.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	//
	//// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, t)
	if err != nil {
		return err
	}
	if signer != header.Coinbase {
		return errInvalidCoinbase
	}

	if number%t.config.Epoch == 0 {
		//verify vrf
		msg := header.Number.Bytes()
		sig := header.Extra[len(header.Extra)-extraSeal:]
		pubbuf, err := ecrecoverPubkey(header, sig)
		if err != nil {
			return err
		}
		x, y := elliptic.Unmarshal(crypto.S256(), pubbuf)
		pubkey := ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}
		err = crypto.SimpleVRFVerify(&pubkey, msg, header.Extra[extraVanity:extraVanity+extraVrf])
		if err != nil {
			return err
		}
	}

	if _, ok := snap.Validators[signer]; !ok {
		return errUnauthorizedValidator
	}
	inturn := snap.inturn(signer)
	if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
		return errInvalidDifficulty
	}
	if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
		return errInvalidDifficulty
	}

	log.Debug("verifySeal", "number", number, "signer", signer.Hex())
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (t *Tribe) Prepare(chain consensus.ChainReader, header *types.Header) error {
	number := header.Number.Uint64()

	snap, err := t.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	header.Coinbase = t.GetMinerAddress()
	header.Nonce = types.BlockNonce{}
	copy(header.Nonce[:], nonceAsync)
	//log.Debug("fix extra", "extra-len", len(header.Extra), "extraVanity", extraVanity)
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	if number%t.config.Epoch == 0 {
		vrf, err := crypto.SimpleVRF2Bytes(t.nodeKey, header.Number.Bytes())
		if err != nil {
			return err
		}
		header.Extra = append(header.Extra, vrf...)
		newValidators, err := t.getNewValidators(chain, header)
		if err != nil {
		  	return err
		}
		for _, validator := range newValidators {
			header.Extra = append(header.Extra, validator.Bytes()...)
		 }
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Extra : append sig to last 65 bytes <<<<

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	// Set the correct difficulty
	header.Difficulty = t.CalcDifficulty(chain, header.Time.Uint64(), parent)
	header.Time = new(big.Int).Add(parent.Time, new(big.Int).SetUint64(t.config.Period))

	//替补出块延迟
	delay := backOffTime(snap, t.GetMinerAddress())
	header.Time = new(big.Int).Add(header.Time , new(big.Int).SetUint64(delay))

	if header.Time.Int64() < time.Now().Unix() {
		header.Time = big.NewInt(time.Now().Unix())
	}
	return nil
}
// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (t *Tribe) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Accumulate any block and uncle rewards and commit the final state root
	if header.Number.Cmp(common.Big1) == 0 {
		if err := t.initializeSystemContracts(chain, header, state); err != nil {
			log.Error("Initialize system contracts failed", "err", err)
			return nil,err
		}
	}

	if header.Difficulty.Cmp(diffInTurn) != 0 {
		if err := t.tryPunishValidator(chain, header, state); err != nil {
			return nil,err
		}
	}

	if header.Number.Uint64()%t.config.Epoch == 0 {
		newValidators, err := t.doSomethingAtEpoch(chain, header, state)
		if err!= nil {
			return nil,err
		}
		//verify validators
		validatorsBytes := make([]byte, len(newValidators)*common.AddressLength)
		for i, validator := range newValidators {
			copy(validatorsBytes[i*common.AddressLength:], validator.Bytes())
		}
		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity+extraVrf:extraSuffix], validatorsBytes) {
			return nil,errInvalidExtraValidators
		}
	}
	t.accumulateRewards(chain, state, header)

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	//there is no uncle in triple
	header.UncleHash = types.CalcUncleHash(nil)
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (t *Tribe) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	header := block.Header()
	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return nil,errUnknownBlock
	}


	snap, err := t.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return nil,err
	}
	// Bail out if we're unauthorized to sign a block
	if _, authorized := snap.Validators[t.GetMinerAddress()]; !authorized {
		return nil,errUnauthorizedValidator
	}

	now := time.Now()
	delay := time.Unix(header.Time.Int64(), 0).Sub(now)
	log.Info(fmt.Sprintf("Seal -> num=%d, diff=%d, miner=%s, delay=%d", number, header.Difficulty, header.Coinbase.Hex(), delay))
	select {
	case <-stop:
		log.Warn(fmt.Sprintf("🐦 cancel -> num=%d, diff=%d, miner=%s, delay=%d", number, header.Difficulty, header.Coinbase.Hex(), delay))
		return nil, nil
	case <-time.After(delay):
	}

	// Sign all the things!
	hash := sigHash(header).Bytes()
	sighash, err := crypto.Sign(hash, t.nodeKey)
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
	snap, err := t.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return calcDifficulty(snap, t.GetMinerAddress())
}
func calcDifficulty(snap *Snapshot, validator common.Address) *big.Int {
	if snap.inturn(validator) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}
func (self *Tribe) GetMinerAddress() common.Address {
	if self.nodeKey == nil {
		panic(errors.New("GetMinerAddress but nodekey not ready"))
	}
	pub := self.nodeKey.PublicKey
	add := crypto.PubkeyToAddress(pub)
	return add
}
func (self *Tribe) GetMinerAddressByChan(rtn chan common.Address) {
	go func() {
		for {
			if self.nodeKey != nil && self.isInit {
				break
			}
			<-time.After(time.Second)
		}
		pub := self.nodeKey.PublicKey
		rtn <- crypto.PubkeyToAddress(pub)
	}()
}
func (t *Tribe) getNodekey() *ecdsa.PrivateKey {
	if t.nodeKey == nil {
		panic(errors.New("GetNodekey but nodekey not ready"))
	}
	return t.nodeKey
}

// initializeSystemContracts initializes all genesis system contracts.
func (t *Tribe) initializeSystemContracts(chain consensus.ChainReader, header *types.Header, state *state.StateDB) error {
	snap, err := t.snapshot(chain, 0, header.ParentHash, nil)
	if err != nil {
		return err
	}

	genesisValidators := snap.validators()
	if len(genesisValidators) == 0 || len(genesisValidators) > maxValidators {
		return errInvalidValidatorsLength
	}

	method := "initialize"
	contracts := []struct {
		addr    common.Address
		packFun func() ([]byte, error)
	}{
		{params.ValidatorsContractAddr, func() ([]byte, error) {
			return t.abi[ValidatorsContractName].Pack(method, genesisValidators,params.OwnerAddress)
		}},
	}

	for _, contract := range contracts {
		data, err := contract.packFun()
		if err != nil {
			return err
		}
		chainConfig := params.MainnetChainConfig
		if params.IsTestnet(){
			chainConfig = params.TestChainConfig
		}
		nonce := state.GetNonce(header.Coinbase)
		msg := vmcaller.NewLegacyMessage(header.Coinbase, &contract.addr, nonce, new(big.Int), new(big.Int).SetUint64(math.MaxUint64), new(big.Int), data, true)
		if _, err := vmcaller.ExecuteMsg(msg, state, header, newChainContext(chain, t), chainConfig); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tribe) getNewValidators(chain consensus.ChainReader, header *types.Header) ([]common.Address, error) {
	if header.Number.Uint64() ==0 {
		validators := make([]common.Address, (len(header.Extra)-extraVanity-extraVrf-extraSeal)/common.AddressLength)
		for i := 0; i < len(validators); i++ {
		   copy(validators[i][:], header.Extra[extraVanity+extraVrf+i*common.AddressLength:])
	    }
	    return validators,nil
    }
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return []common.Address{}, consensus.ErrUnknownAncestor
	}
	number := header.Number.Uint64()
	if number%t.config.Epoch != 0 {
		return []common.Address{}, consensus.ErrInvalidNumber
	}

	statedb, err := t.stateFn(parent.Root)
	if err != nil {
		return []common.Address{}, err
	}

	// method
	method := "getNewValidators"
	vrf :=header.Extra[extraVanity:extraVanity+extraVrf]
	rand := new(big.Int).SetBytes(vrf)
	v,_ := uint256.FromBig(rand)
	data, err := t.abi[ValidatorsContractName].Pack(method,v.ToBig())
	if err != nil {
		return nil,err
	}

	// call contract
	nonce := statedb.GetNonce(header.Coinbase)
	msg := vmcaller.NewLegacyMessage(header.Coinbase, &params.ValidatorsContractAddr, nonce, new(big.Int), new(big.Int).SetUint64(math.MaxUint64), new(big.Int), data, false)
	//
	chainConfig := params.MainnetChainConfig
	if params.IsTestnet(){
		chainConfig = params.TestChainConfig
	}
	result, err := vmcaller.ExecuteMsg(msg, statedb, parent, newChainContext(chain, t), chainConfig);
	if  err != nil {
		log.Error("Can't decrease missed blocks counter for validator", "err", err)
		return nil,err
	}
	var out []common.Address
	if err := t.abi[ValidatorsContractName].Unpack(&out, method, result); err != nil {
		return []common.Address{}, err
	}
	sort.Sort(validatorsAscending(out))
	return out, nil
}
func (t *Tribe) punishValidator(val common.Address, chain consensus.ChainReader, header *types.Header, state *state.StateDB) error {
	// method
	method := "punishValidator"
	data, err := t.abi[ValidatorsContractName].Pack(method, val)
	if err != nil {
		log.Error("Can't pack data for punish", "error", err)
		return err
	}
	log.Warn("tryPunishValidator", "addr=", val,"number=",header.Number.Uint64())
	// call contract
	nonce := state.GetNonce(header.Coinbase)
	chainConfig := params.MainnetChainConfig
	if params.IsTestnet(){
		chainConfig = params.TestChainConfig
	}
	msg := vmcaller.NewLegacyMessage(header.Coinbase, &params.ValidatorsContractAddr, nonce, new(big.Int), new(big.Int).SetUint64(math.MaxUint64), new(big.Int), data, true)
	if _, err := vmcaller.ExecuteMsg(msg, state, header, newChainContext(chain, t), chainConfig); err != nil {
		log.Error("Can't punish validator", "err", err)
		return err
	}
	//log.Warn("punish validator success","addr=",val.String())
	return nil
}

func (t *Tribe) tryPunishValidator(chain consensus.ChainReader, header *types.Header, state *state.StateDB) error {
	number := header.Number.Uint64()
	snap, err := t.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	validators := snap.validators()
	outTurnValidator := validators[number%uint64(len(validators))]
	if err := t.punishValidator(outTurnValidator, chain, header, state); err != nil {
			return err
	}
	return nil
}

func (t *Tribe) doSomethingAtEpoch(chain consensus.ChainReader, header *types.Header, state *state.StateDB) ([]common.Address, error) {
	newValidators, err := t.getNewValidators(chain, header)
	if err != nil {
		return []common.Address{}, err
	}
	//TO DO
	//调用pom合约分配发送mesh奖励
	////模拟奖励poc相关节点
	//addrs := [1000]string{"0x9c4a0733fb0bc42827891ea8029901f28a60b63d","0xa8e5caf50e749f040c5be8e3728fd49d3d891167","0x113de97d6d1ad27e63694b07ba0bcc4a9c959d5d","0xde6701c5563ea30b26df04c11f440e8d0ebb00e8","0x2ea87d4c36e4dd78f738917a553b3472e44891c8","0x3b8c788a375eaeaf3f9b892f0d7c68032dbb350d","0xb9b2e02907f01b349e5693f741ea6c7257ce4fda","0x5ce59142cff5c05f03b6041d3ee6b122cb12bb16","0x129bc5eb98eb0314522c7176d52b2a7e1f199b4f","0x5c7ac299b669d2466fdea5423301b226aebd5212","0x4d2f2e58fe82c6766ed8583becc959c3c8d68c34","0x850c19b06f681b2b84b084cf16bd9443dad1fd98","0xe11d99dea38311f590f1a8630d708fb0b4fc480e","0x02dc1f435717bd54fa09bb139165b2f8d756e922","0x3b13f41bf791930bc9cb7fbeaf878d3d0c7a1bbf","0x09be75b20c2f5205cd7fb8ac42fc2b16561d827b","0x4d27b5dd7e71b210fa4265c8a4302d3687129d9a","0x740b5f239e5f180c063068bd9a6b9349f0dd74a2","0x198277b7e24f5349f55c20ee57106cf7999b432d","0x7716d4e3814b7889ddf9be71133d00b40a8e5b25","0xcee7876caf7a569c4a638cd70a64242a8d01415f","0x17b9b9f2ffe1db30d533d1e09bdedfe8bd0d7ef7","0x7cb7b2ca6eac00f2322c97de75f1f028a10953a0","0x2f8e03e71f29246e88f95b41a0f5da31ccdf6aa0","0x5578af06c35970eaa1bacc6d32aa0dbd239200c5","0xb477dccc65722cb5a4114ca3ff017d7ba6eabac8","0xf1db441549aceb76a72fb9a408105efdc2f6933f","0xfc6496bf4eff99ec517e3cd3bec04a2b67319b0d","0x3d664ff0745607a649e675c8c7cc63db9385f5f9","0x928e95d633dbd501b09ca2befd782f0834e73055","0x2e7dc3578cdc6dcc3867619e8742ddb30880ec82","0xaf4afaf2483875b9ac274f012073ed0922240879","0xa19154abe2c9a39baa44ea86e41dedce3d941a62","0x2af83b041769453d16553b52c209708a2f5e67c0","0xb0aea9733c4a2644e56abcf1746d5bbabfd1eecc","0x9bc5603483b254200fcf059e0177d5ba1517070e","0xe0bbf1e89fff89c46ab4b5e32ecae015e5cbb08a","0x3cb665ecca3547c2f8d8f96efda2ff7c7fad9a87","0x5454c86fe2a5ea4775f41406dc93f8557e0d784d","0x9086eb86c714f89fd61e59154ec67e74c678606e","0xec7e322ef315275778ac3ba78e3c1ca29c42a850","0x8628095455180a22584d1f70c737556071dfe17d","0x654bf145cc6e4902b30ded52ce37a6ae646ad7e0","0xaa179cfd674671886b5086e985a0abee6117cf0c","0x25c14e28fae71a4292439ac2a9484f11c9ef4439","0x267009d6fec040a37ccd75445ea73c2decedc186","0x1cc6b31053ad164c698ff98e6fee9622e12de8d0","0x27a4366ef50adeea9e3be1c2a451be3d53e6b30f","0x941c0b16fb7305029536bba546d3a6cdfc1bb847","0xd25fd72ab4dc2603e3ce0fb02b41669fa8b519ff","0x39ed8b16c51188c1df351db3e4ed78e6a8db2700","0x38d294d212b90f5cc794ff51da1f70ab65fecda4","0x2cad9a7fdc0b3f05e3de5890835921baa5e5c6c5","0xbf665d3cf9206aed1b5e4b8e218564a466327b68","0xa015a6e94fa3e5839ee6d0d4d31b007410d93436","0x897a4de5f4ae67ec04ce261ade210e4bd0a4817d","0x27ebc107ebcd7c3d2d22b743f4ddf1c0cbb596c4","0x87527d6601b2207b9b4db1b3da6236f5f5e0a98d","0xaaeeab2273d04c3d78a79148bafac1622e52dfc8","0x78d7ccc8a47174b8b3f60b0eb2400808250a2e37","0x4358c32cde25e1bae80ad4f68d202a3e058b2d06","0xe207fedcde2036efbdcf58957175296c42febdea","0x60f0818b8c2e7acd77a1e0d8cc84f6c7c925989a","0x6a047bf4cd1578b0176c9595d58613284f58c82f","0xe61966403e1e72d4b637baaef8683433ef2b9cd2","0x5154927aeb019db0daf72edd0f50e86085169bcc","0x80c87f8a54543c01a8553e5ed5477cce8e620a00","0x6b3cca450a49c6a9d18df1ae9c91387aaa6cd70a","0x06d539a7c0f8f10d06c709c3dc389e66d8e84d91","0xa92e5288d49fb881a0c448df723015c4d5a08808","0xcc473fe389bf53bb8ce2fbe45c53d557fef42e64","0x5999b4399cb94c7601a4694415d7d699546343c1","0xb19a4f05ec0ebaec5cb3493aaec635422d3e4cf5","0x606d190f7f9ed187b8687f746d8ae8e3fc994e9c","0x98c086f9c4ddad9dfadb87d0fa8b9a24eda1eadb","0x2d5531622febc738ef7317649e9e6db3087940a1","0x4773e84639a7c3f73967b90de4f66f843f7d9cde","0xcaf6b39628887d78ef6fa16dc64a307c439febc7","0xa5cc6a3fbcc5fa06e583a371a36c85dda6a100de","0xb2f47acd1b42a6f5b2f8f9bb7e5d84599d9e7af3","0x93db03d6e6517a1fcdfd2287f6c43e0689943e1f","0x74cc9662565401f61fa95bbd573656a6466dc11f","0x74dbc6be47839fff6f2318accea57af49c97dd7c","0x67734eb43ac5eaeffee0f2287338dfb7d14013c8","0xa625910a2c7c1322a4e17aec733d9562c4f0027e","0x6ff579f6b895b09415c2752d667a1277812db6f6","0x3ce1e4cecbca6bed37a0c0bbc1b6531e11baf4f0","0x72d97b084c34d01ba3e37791e3a5f3bbc7b71fe7","0x3163132dc37a07c5e9d5d1085ca998a57931d711","0xe64542fc29ef14899252728a9b10609ba019490b","0xf54ba48bb0e47394047ba1d1fde8f48eccce8d23","0xed19d7449f3c004edb30d378e703208d93c40369","0x5da90151147628184b7296ea4c151bd864aff6e7","0x32afb63bac06c268c03e0a425c442847c52bbb95","0x4b1a2ee63772f3ed4b00d7a8bdf25a3e614f59e6","0xd8b03a05da9c1bdc874928f399524b962f98dab4","0x13a640466a15a976293d6729076b2d39c09a3e03","0x765416beb7d39fd69c33a3968ffd9a99054dadff","0x94e0dc203bd993f693f24f1d04c49ef15d6e8701","0x3881dffe4b286db1f1c9ee2a08c720aab8295d94","0xafce42545c878bfa39a0e46bce1168fcf6bf6a5d","0x071370e5ea9c4859f43b87cbfe91f0de57429e3c","0xb3b67cb8651e386c35bc27f158a3a8b9880ebd26","0xb834ec80b5d25f803454a72fe2fda13b7f9b579e","0x44146c8d7cf663a186f2bc590f69b822f7f8a21a","0xdd9459845b56b0a5018b10bc52cee57a0289ae0b","0x2df5e5cdf901a91aa402f01a3661ec432bf44e25","0x322490434aea751ba3ea351812d26f0407443fd5","0x27101e8de9ac2ebe13f70426b6757dcd3997eaa9","0x6ca7bfd01b89b265fa23d24b37ded5c9c3dfb638","0xeab362755603ef23ec788f270cf9692121599c8f","0xefba0556ff55c5818e42268b15d95a3faa50ce29","0x770e4e185033352d884ce622bbed22a4a7f1e449","0x978538d25238fc7a8be3f398376a8b95015b3872","0x0fcea2b947ac8bdb06904d35d3f884cb6bb609a9","0x29d3db28fcec10724d27cb483fc5177c469ba14d","0xeebc0cf04528aa9349908bc727848118aede1964","0x78e199c43fc5c9a7becaa92d860d4bde5628711a","0x2bd1c6c35708ec52b812a024de77687764ad172c","0x1676fceb6558a03b1d98fa4222d073e0c5519f93","0x17c3dd41f542edb7bd31b22e02df96d138832910","0x0327020b61a9b20a8e8f03d0d17f9019735eec71","0x3cc746b395758c690aed43e6a60ee04e3e21f016","0x49f5a4754f8c549fc14a390eb39d7feb38bc624c","0xf9453a19223243a0cf3d09c1e7384b73451393e1","0x5238fa54f03169f5004e38d8aae3bc75008a81b5","0x70cf7f9f812ece67cf8a9958e9da7af4f9eaad76","0x51e6c444f85387d5a0f74fd8fd2e7ac82a7ce151","0x3dab35b24386c2d947f4d1fd5fcdd7ea56bbaa70","0x8cf689ae55261392adfc23b91eeca19dd4f275ae","0x93b7eba99de30f38b3af658817500465213f24c4","0xd1867353c5c21a22246002cc467ce91998495031","0xf2d343ed83e485c4d4c528fc7456092b6d3f2b80","0xea15691b925e8c21cbc91b245c20e71dcc6088cf","0xba4d0903b365cabb0cdd7d2e0e313fd983e7bb3e","0x3c057d8516f67071f74709c8223759359ddc098b","0xfab4ea04adba1437e708e8bb0438a64eff542f99","0x5b96ad7220e944625f07cf57fb7af0a7876721b3","0x246c5e394aa40df1b20069d87c71b055a19911ed","0x1eea6582c6b0c5d00fc440b9ee3a425b64fbb55f","0xdf983ccdb515e78ff54de4bf148cbc19a8fe4bbb","0xc8f4003b22b8c734774c378d491b8d3cf777a244","0x13d82dab72fe7a3fa920017be08c0ed036653982","0xb011aa5cb3aaa316174e1cab1569ec606f6be5d3","0x12b369d568fa475df9220ce9c25c6effe3c6c9fb","0xa1ffa1ffa80976d72067f439e26464a2189a86de","0x25064b66dc540af56fd6c763ba708fea0d684762","0x54131a5c17da175f7781ceea1591bd6b628016d3","0xb2cfddaffb2227057aa0b0751938160a9d091718","0x348e21e2809a109e39a4fe8d896458295ad439ad","0xcf403547b1b9c85cc4afbd43420e5500945bef65","0x3b6aee2c0b778d139edb2acfff0b0b3a10b7347b","0xc73170cacc9f961f4d1754b4c48c33358212438c","0xd5dad60a1fd87e193bcf2975e5dd547967ee69d2","0x66166184ee4dea48d032d477e48dac649543ae3c","0x6acd973ab2a17df22ee242fadbe14c1beb2271b6","0x4f25abead90e1508f7ef3ce36f0356d4d21346d3","0x44dd607ef6cd88e86e9dca6f3c16dbed8dce4a7e","0x27df9cd360cf3e39f51b6aa0eac4e06af1c5a251","0xe7263020eeb9c2c4698780985f420a28ab055dc6","0xebdb10211e4b3ba674b65cfb5871db79589342d2","0x859b64035ccc37d5ebf6cb9c202261bfcd155454","0x2fda22ed61b4db3c3641a19638b56b28db03b741","0xa809efbbafdce03c9a1308f24b0ecbf2d9013430","0xb1bd931e9ea9193e061ea1beb08254cd97fe54fb","0x6611998e9591b2f3d08ef9628220bae5f073105d","0x27bdb45dca9132aad2e4bef7cc16b4a109153127","0x0cf11bd6a804ed5adec83b6be9f4d4fa5fa9138d","0xc611c08ceb0acd519b964cd26161894d78324a3c","0x9b7593a03e9433266ca44f2d7ec7c0a3a434edcc","0x9b9ab0c7e69368c735fbe5b37b1e88d84bdaf71b","0x8a291e2f79989648056df42d863eee51dfeaef7b","0x0411a238c2d61b379f07cfb206ce16ae14c1dc1f","0xb3f5bcc8bb1325107e70cf54156cf286a45275fe","0x5f2131060fb280577a779c46fbc4a557d6691b0d","0x50d2b007f539d0190a6e2dc825916b04c280b255","0x09a1eb1035e5bfacb8f2f6666a69520cce9087ef","0x40fb80b2a91bf0f3ae8a29d5ba842d7cb7838602","0x7075c2f7960abe5b6619b56cef2e0cff039b0451","0x238139f2a8d02e8c5273c52fada13911fd506793","0x77eec89aaf4029da2e634d9062afc35d49ee58d7","0x0a707f27d4d3a2fefd9734c41430fcfe0fce60af","0x6db4f5af1f25b203eb969e37927e4f875df7ab73","0x11f9814d630d8d90819544a32ccc528a9235d0cf","0x841645f58aeb8e42eda15f577ff84833627e2b41","0x6100d045d04a0e01c82741b68d5087da1848831c","0x1ebc9bb44346335ff9a9885f663c72ca7c64f8c7","0x73bb51f2eb7951e84d105db45d2bdf7ff7e983a5","0x267d9ce90b67c2031c729dd82948daecbca69eb0","0x8ba163353fd32c8df35937092bef2087f4b61c29","0x0720bddf73c0eadb951c457d1b1b3c1e1113ae55","0x82bbd1b209374a0fc052f17d803d7a366d1b9a95","0x6d3646aa35d2ab9cd491b485a95dd23523fe202e","0x0b4d4ec4e396a2d27f1edb586fca5063c5746403","0x1f0fbf02e6870ebaa9e31121781ee8083ace13f8","0x653596c50b912650f84b84929de2c8721e449b66","0x08ca3d9284c7159c32fd1fd2238ac25390d9adbb","0xcc0c6d05470cce763311452a785f230231a6e255","0xcea09092f16d55c335c986b91e1f13dcd2986dde","0x8e374494561dbfb02e484d70d8cdb4923d4132ad","0x7e19e07b76c2b4abd4bf3f2db6cf8e3fdde95be1","0xb9f3bd0d06c7215c77a9097abbb91b5c1cf06e55","0x2dc266f198c4d72cf77440bcb32da0564da2022e","0xc0066b8826078ba46b6c41b57165d52038e56036","0x6389d0b85b84139bcdb3378881fb2b02054cb9a7","0xc8972f5507c87aca87284ed0cf5ed54f2eb47fca","0x12fda30b3c97d6768c0ce28c247b1a06ede2b823","0x146345fc209d42a26a3f1812813465c895336d87","0x32f9f121d55d0f7a68cda73d9c4e20695b4c2fee","0x3671517cfb4bb9edcd77b4259c72d49bdc1c311f","0x43b555590b4b0508662ca6b4ff3157d04e7ab356","0x9ab7cc2cc46dfbcd796f5e295c6317d362cba024","0x4bdaa00390e811862bb70a9ef635d51a250aebbb","0x85fd1f09288dfcb82fbd27140238b5f38a733a1c","0xa441f87a359dc55e08d0719c86bcd1f0c643033f","0xcff1450c45b739a84dce35249505edb4a61f2a1f","0x0a79b85ceabbf264ecb2ef35c723c2fd159e3d49","0x3fc669a1e202c875d99fa9d7a731fed3a4c50887","0xd268e38b0d653c7119fea1a07cecae9679aec195","0xa71d869b378c4855a88d1cdf53ff8660577faaa3","0x6059e900e5c57fb3e114d109442c68ebd14f92fc","0xb4bfbb77434aa963317f394fcc8ddddd8e962934","0xfe6c8f655709dddabf92079e11e59c0bdaca17d6","0x720346aba08becbbfeadcc4b41b8b458cfd4fe27","0xdd47a82857016e53fc55fff7b1378418d201ce1c","0xfa7bcd326a21a483c307ffe43358ad4aa9183a90","0x91955d6870bd4c03aac36e101bfb4105c033d2d3","0xa4468afe36432bd4aa6926ca04e28031d8538153","0xfad56d34df56570e6add79642ca645623fcff762","0xb3b603611a10d157738c5c0c55c9269f9d393d3e","0xa24d297f1492a9210a6fadc7afe83b87bb9e1e95","0xd10939f1ba6b9a28c36f1f67de63a972403157e1","0xf208237342c445ae3621b01fe0c4f4807007bb9a","0x4876ae461a7c34a0ff791581b337a1e6f27c9205","0xbe6413383a1f497157a2f526389b37c1f18c6996","0x6ffa7396a6352c662b63e3e9dcab094246be7b52","0x515e5fcf05a41b493f0dc4859f2b3ba773811663","0xb4e983921059db33aa656a95a62516b17ea43d46","0x1f08d8e24427ac130b75fa2744ee6c3454d6841d","0xc4760921aafeb25eececddcfd96f9f77bccd36a0","0xf3130c8245e44c58d42cd99c5785e19824865b54","0x69a0e4037c250df13b0965c6120fbc60223118e4","0x24a5e51c08db9853b04b7e72b39605e0b166a835","0x66487cbbd4a3089512cb30087f39c038d6e76989","0x1c8560b3ddf7095aa42ddc893fdbd8ce40b6a34e","0x1ae7cb49913e38014b643daa801680c2afbf59fc","0x4a39bceae5ea811409a85cfa243200575a0dce7e","0xba5ad2a378050d50a644070c34332f393f1cbbb1","0x0dcb167147b630ade6503d0e113cf9c6bd2e33f7","0xec9eabcca62bead1db038b9ef9486439ecca33d2","0xdbae08c849aad316f88b703a6b874679d28809c8","0x160219e25c0793936f673a903b04a7103567cb3c","0xcced69b5fc73094cdb94a9f19cf0f7276650fa71","0x7125179e180471a929ab301f601cd7ebba8c381b","0xf693e8bfae2e49e86c7976189b17996fbe82fddd","0x6eaf3d45bc62db5a91af9c43a5a984554dcd27d5","0x4dd92d714d092cabfc413c79da4b3829a23eb777","0x145aff5dd52d271bb0fe04cb8e5ece383835a01f","0x19148937a53a94964e3a364d8a3337083f5c5751","0xda09c9f963a6c5b40555140642000f025582d966","0xf3576fe63a09231d9f73c3cd0d5e0f7636a4a81e","0xccd0bec3ad7e0ff14cc3c784415d1d3bf44a68e4","0xfdd53c571c4d71a5a2a990e66c10dcc887e6f6bf","0x93fd4a51198a72ef6643e3b1d92230ca37cf4396","0x4f9c749356b5ac2d8a736b783e384dc90cf198ae","0x4d8e3def68d607929a3c0adc40c5fc045bc4b956","0x0e7e9ee4ea32e0bac18c21aa9a475b34f9d2a3c9","0x1cc1a654762fb1ab03e87b7f66a1d3f574f46a18","0x0a55e4b77ce69c9c6e03422e00b60ba78cf72719","0x16ff91ad13b3171a196a143c6583991a8a77fd6b","0x9a0a1c1006e0d4cb5d95ea81f9d7af2c4ddf6f64","0x9bc6bf61ac3e22fe89a9a676094f78ee4c2a05da","0x4617b57fe0ca6fd3703c895f04e7506342385eee","0x1898187787c786b8e957c0c83c9b91f826e4f072","0xd1896d184b734bb336d55abbf771eca38aa1b1d9","0x855e7660768d309a85293e45b3128c1a2c2ca1f0","0xdd114741239f7aab9fb05a7023e6ccf34918eb7c","0x36f8f1db3884815f6a99cccfb5561cdb9481658a","0xf5624f4023a39cfd610d122990fbe01457437b63","0x91f78e24427759b6d2f6c781eedef4d88ca524ad","0x20d3f0e7ce333d2b9c14ce77061db8b49c0c5012","0xc331c0c12cd0ed34a2650b16580a6d7a0e5f6869","0xcd083e9dd75669be6753290f935e0fce319e6f02","0x4c65a1ec139172ace68f6b1b292d7c796433f577","0x8a35014f03c8fb9365d2b785b5622f2aded36627","0x95556ac7d32c518cab76159a5343ff4954d4da78","0x6988476c7e260d0f88b6a335e670644578c7fb90","0xff31a60907d4870749dabc49f488f1f6cb00bee0","0x63b9c9e80d3d9a02bd5e2cdb4fd612598e9e9c41","0x4d7070dec130e1ebf9e0b4c89cbef02766ca40f7","0x0b0777d1c91c01a60610272094fbf21bdcbea117","0xd9921b4aafe460ad2cbdd5d5e17e09215e75a2df","0xd052de9e4dd501357598b4eef3176cec28e08eb5","0x9e5e9c89ac0c2a0a13978d80daea3de5bd36e5b9","0x4272a809a0f6f03434d0b043386bfac9dae4fe1e","0x9e2b58741a381e9db28d398be1a746f97489ae31","0xa10b4f8b396f386cb940162076c4ca5c012f544e","0x2a4166a915df05a1adf556657b47f74e484a7207","0x96612b78ab2e829d75e3373b150aabf921d5a632","0x8ce0ecbac342841581cee5a5f77348a12d36b11e","0x910863e79d8a65ea2ccfa48837f4e8fe430981d4","0x541221bdcf48f58e5017a801373f17e42d4b39a9","0x44ffef4bfe29b8e092a8da1a0f46a6ef4a9cad99","0x368c24549e76f48dec6b20f5b9f1d04e2f069b7c","0x9a32e39f454f74d7f6c3ba376dec64d41276d5a9","0x0519942ba492276502991111ab65042b2fa47901","0x291695653b148cb331ae1cf979e00a4d5e46e8b8","0x60eae0b9c1c83925c4ef7d6da88ae449982292da","0x642f7e0e69715e11398c1f220442e487626260f5","0xaa01396eb71948ac2ea08af2762f9e3952ee2c3c","0xa20b8cfc303093661719db542d1a0f7c386bb434","0x97dc3fd874c5eb99235a92a45091506d9ecaae6e","0xe6b6e6b0edd8b83713662d50d7493133594988c1","0x46dc3510764e4c51681f6fb137ccdc4a3c2606fa","0x610e034e85a6e9e782fc26b3d6d6fff0abfa32de","0x2e36e17f40a6228b3e39072d8eca52818085a8f4","0x0843eb850938c32af7a66afa052e60afc8d6ff7e","0x5505e5a3ffa5e21cc01a121a4d4988a6e0dc5779","0xc64d3de2193c64df21ef6367ebc497d59399470c","0x784126bdb17a7406ec8c9a4c420f73330bc0aea4","0x9b2f38d060eef68959551d0035704fe92eea7837","0x112aa9bade11695416bbd8e7ce5f5352574f5b0b","0x34355afe7ef13c7e217ba94a5562e9e916d58385","0x2337532e4d6e70ed998bad5f1a62027594a854b3","0x241af9742e92f9e90a7b5a3ecaaed6de1c198bc5","0xba18b158b49f1774bfc43017b9b114975408fcc5","0x72e3f3f694c3045c1748cb08b4fcbb6f0af937e0","0xbd4f16711cc3faf3a9daef0349e3e1b2ae9f8341","0xb92cf71133eff74ffe93bbeedbf60e58ae628913","0x866d272756cf3d3a39ef74323d17dc100d432ca2","0x457114b447dfaebe54740f530e27f240f23ec15b","0x518d5c7930126005404292926b58ca2ea63f4ebf","0x58c62d921f5d3ba3602b205e4b1e9123e9d4a2e3","0xfa0c934e2c9eb31b097b0940a618b78de3d3fbd5","0xb80452abb9eb424d8745ae12be1b9c46c871b72a","0x67088c198c02a0f049839574bab992f0ac91a9b8","0x1ebf4e38ae570c11614d57fbef2e5fbd2dacb2f4","0x3bed1745f40466c9badf3843f3d42231c8678592","0xcaf40aced05270b6f1356a5a62e58b36238060ae","0xc03d4611d4ed704081fd1b0d208a38c75d617ed0","0x093a1c5b3130650545edebd8355289233815d633","0x08ad4f55b7c65ac42a2d4feaa1aa77993c936122","0xd22b6346b28ba50d8895f234975c6ebf01b4faa7","0xa63c334fc66daf4107b4b565850c0377936932e8","0x8f27a5321f0311d3f85cd2548d2eab542fd7b793","0xd5e3ffa58e4bbfd9d2ffdc3806fd65dd25460600","0x9d026e57fcb7af927f26ac49a598d06014cb67a5","0xa8186e86dc68f820175ebb431908b5b77bf6e6fd","0x6f2395fb6b09e9380e3b3491ffe909c06f8fe019","0xfac20752fe58abbddc7eefa8dd17ed92950b7801","0x69552a6e7c6b0367999082ab78b6e94681135c13","0x8c746626cbd2d281a90e6aff9187d062bcbbc521","0x9752a4f5f6534dca29534779e7dadb70f977a83b","0x0384dce271496949a9ac26ec76339dddf1586292","0xdb97b26d671119c904a0213cf6168196146bb5c6","0x7e283da34f2782cd8763a3bce803f8f59994732a","0x50f8f70b9bab5b6e54719c2852f2559185821ba2","0xd0a8dd5787cc61d9102fa3bdba7d6ed3406bd606","0x28a4f91d58026f623c46faacc766156638ae2c23","0x0bc80eb3695ed5b9a6729141c10ba2b60eae0a3d","0x88a39c1223d04e31b0a4c529d8c930a50c93d495","0xcd516ecda8179833ec6895cff63de75d2be6515b","0xd2392c95b6f6cc4fb77692b621b132fe1fe53683","0xb46f95d2712e7b0c0b16f3b50c2ba88b873ae8be","0xbd84d0475452c421a23138e58e97358e3cb3b24c","0x5dfdfbe7baf6e7298fd8804a6ea98cb0d08a08c8","0xbe98ccce56f0bb762238ea942eba12a91cbd4742","0x88f0fe0e9698c1ca558847129f39570fa2275b32","0x1b02b04c7aac623cae94da251fcc998b4d07987b","0xf4a329f401a56f401f36242643f89ac61c453400","0x0140f93d2c87c2691f4602f70630f8b052bf102b","0x1be836d2ff73a2ea2a05d0dad74d5eae043787dc","0x5829f4444dcd1e962fb56b977582a92dc92c58b2","0xc236d9690bf22b55b05438741e2e48f037a158d3","0x3d8e802a0a6bfbe8d32a5b582c59ae54cd80a5bb","0x23d5a8846f892286e04e33680623ddfc1fe17f7d","0x0a877c987e629ba5c0cba7557cb9336ac8b45056","0x2a110a821c9ffad47d8054f446b0fd5ff7ba74e5","0x59a26c18e23a9d29ccc69b6918e499406d25bbc7","0xa85cb9714fc9159f1b2ab5efb0f9ec3ac669ce64","0x80a565e16503649c097139e6def26573b860cd34","0xf5182ddb81d5a07343f4eb37d17b27a916ebbf9c","0x1d7a0ae6855e6436ea2fd7af2c04e8b5f41486c6","0x02b7c5722f99ff09f9df1c13a627b73e1b7caca0","0xbc76e020fde477c5300371cccd8d46c27133e110","0x74a341b8721c07328b9fa7ec4b10d4e9872af8ab","0xb7da8ca706d9de3f221978b5d88b84e284a8abe9","0xdf320939f6841019426f69ad3e50fae52022bad0","0x7da370f08cdf222b1df168ecb296edb853013616","0x66f07a9bafeb8a86e51a20551a997c9b7227b9ba","0x26d0a0b66ba1128c8339efa4ad2e8b6218753d73","0x6ebd72bd0b07119bd3e78f6b1c33a95b6ad208f2","0x77b8ef32f128e832a849f7e3012673ba396816c8","0x21988be88231144a6627d71e74f1aaf226381987","0x45815d79d364d45bc89440d02dbe45265e4c0f48","0x2016c474c8f835cb3be7330a1b83116e85b40f59","0x29fb70f3912862191bcf834208f89f696a7d98b8","0x2196e07722e88476c1098ced4438e4e3af9f4b14","0xc0c4a2068fcebff4ad955fb542a6d8ff4fc8204a","0xddb33d510ad002e5f9a633bfc0efd3511fc3135b","0xdef3d853a32ee78a5f358a4ed3dbe64546fd027b","0x0be4a25e75a450af0688608e31c3d96ecaf3407a","0x87fca0a1ef13c454d3aa5286ff49fce5efd0fdd9","0xf71947e105cb433f91eb99c2d80dd3e7249091ee","0xda2008b0c29b7970829f003b2cc4937e75db37b5","0xbaea4c8b69ef8fa180447eb529735cdcfd880845","0x88feb819d9192d7a39d8a0cc808aa41b9e9547be","0x3296e32cc223c5bdda601e5dc98787597d105170","0xfe9bf3b670803dcad9bb8c1584c130190e30b26b","0x59803a624cf7e22f3eaad907aea0f882be7abb8e","0x18cdbed3ecc458d24a815bc50cb3c0b182d1b47a","0xca02d520e411c7a754abd1b1b4935852b16646e6","0x3438deecd4803c8a89bd49061c98644f877eaa1e","0xda972926f3f4d6fdbb4d5ce51110b2b5e3b4d80f","0x8adc1900f709ae15606cb3a4a4d6d4e4493393f1","0xf8e2331f65aa76174059c67a8d2e04119a3d19d7","0x5c2941538784842abe1da5a2c54aca5e2d331198","0xb8ef574efc84fa2caa4146cd05b9394b914e5fdf","0xbe65ab7a0caba885a789dc153f4d56ec91c5fc87","0x1b9c9092925b9c9711194e13926d3b254209d89a","0x095dafcdfced58f6758418256598b9d67b4f80dc","0x92bc3b71d32297fdf8468a51c6e2798fdd04f44b","0x1425615ff7b11725422dd3919c7bb0122b2b3fdb","0x9668a641b17d18d09310dfe050471bdf7d0222cd","0xff0cc279c42fef9051abbc2291d1cc8f3699f209","0x43dc9215f8dc322195fa85fd163017dd48764a37","0x47c5075a9f3e1cfd694fdcf9cccbebd3f3d85f25","0xb991a8f541f3983a2afdad66f02c5dbe99de5869","0x54bd7a0668bc5fc02c0d2af030227d850e6626b8","0x151d6689953784b9a38561c82459f53b1801c236","0x45ba04c42aec064fafdc4779f30d2adca8c6a971","0xfbf37bfb3b059645b9c26189d0388a983095ee8c","0x450063252dfa357c5aadfa2f075f65cc723810f0","0x6a3d516367ed3664c33d9a8ffe0c85beaeea151a","0x3233cdd6b41363348859c65556d161aa74e7dd06","0xcdd9858f57584f659e18d9f70b532b875e0ed663","0x71d59d13c1b5ca1761676fed8ac2268b56f84c8f","0xade32fa40d88176853d489df8e81e2cf51b1e46e","0xd19afef18cddca078bb45553ee853236db99945c","0xc212489ff3a748e69dbc4a14f851264405696419","0xb7aeacc73bbf0ea2c176e2ac4af113b5a3ff9e32","0x9313d8aabf5ff1568ddbe7c7ab181aeddcfc8c87","0xfb8a2e5e4c47a78b0db99b133d959a2fc51e5c47","0x2c7dc2dc02a10feb8193306028fe2deeaf2b6fca","0x0400b92a539e5a6ac1052c0615753827fafe102a","0xae4f02078202a9ab64235de5b76286987ef00206","0xbedc785038070813e8218bef3f6c35c581637858","0x179482c93409a97449da4bb04c97df730bb5ca9f","0x1a5e158e1e0399b47b426fcae560fdd2d4723fe5","0x3c58ec59a1d1de9c1c172066ca68d307e7160b34","0xe91e256ed380b0d1414b3b27b2b06ac4ff70a7b9","0x1ecf31594cb9e49c82df636d957f83bb94b3c777","0x5967ac5dfd84d431c08deb921ae071a02cebca3d","0xe6bca82fa5694e847f1ae7085e693b72b00ce4b2","0xaa2340b6ccdda684d0a6f20d9a936f970d603e6d","0x3d58bfc87bf677c2fe494d9d1f3c74f29f15b5ba","0x9144425177842668e2dde1268a4ec9575f2364d9","0x140ec304a1f345c8dab816b9617baced220927a0","0x9d1a7b346d29b1f917412bace8ecb066026bee5e","0x3e54f542a7aeacf514700d9bf0376b616aeba097","0x75c39a74c78526b96c75f6f89c5f8742f5bcb16d","0x9c5652aa4143c69ecd486eb87fcbf96894ecc6fd","0x77dd634b89b4cf5ea65d4fd9e00a385c3d3ec90c","0x9c250341d1cb2e2a91b3c2d1fb0849fcf100fc82","0x51d35b3663cf150bb7fa0ad412ec192bc82fc1f7","0x93ff1a8bed8a50e942823c8487d85a9b59d2ea32","0x25e143aa4608b15fd812fac483fac784c9785ec3","0xe3c9c5c4980b2360bcdceaa83914d0dc2eaca884","0xb1961f5e6df2b2c8b34333a91482724ad4002d25","0x925dc88214a8d7495717a0b3bc22ff3f381163a4","0x38f369fcb766143c085291b27c614af63da80a08","0x0b27b8bed1f0492ddee50c2052158911b8c94363","0x4ab8827ccc913d101b92bb8afb2bf4b77b9d4082","0xb730a799ccb0ee151cb84146a0b8f1b857dd51ab","0x6eeafafd29791e3b7a2db2969cc3f937f224c823","0x9baca419286eb638eb97b3d27ce8f79a416fd9d4","0x8b23b0e851bc54351b50cd5c7d7164bac883aa11","0x249c2d86d79ebecff08a2c15a4138d1307f51f87","0x98aef647c424328d32a18805af58d18462c55945","0xd9c0534cff8a243e44eb2a6c5b8e6cf9c304c7ab","0xf342b0ee102a14f3010fe4b73b098dd6ff168865","0x38006f565fc73d57f9899e2d7a5faff793cfc6f5","0x98672f06f29954c2a9919515b51a00207b30287e","0x1173de299dac7056d947d4d9d56960bdb01cc5ba","0xdaee0b6f121cd4f0f29468c99cdafe89998a68eb","0x6f52f07bfdf49e499f0de918795b7aa29506ae4f","0xcd2d90acfd6bb81323e45e06c500231a6d0bdf1b","0x0a6f351206869bb3a10e57cd75fb48fb240ac30f","0xbbed799ada496a1b51caaa58f8253f0996c42546","0x8b9bdc6d573b73713f1eb208447d2c86068067cf","0x44a3d0350c59269db2c6000eb21bac9f6ff727a7","0x7b106cc43f0e01b82135373084150646de85842c","0x5a041d7294ab412e8749c034347320eef25dbdac","0x6f83538034ab1cf12d1a1e7d55d221dd996394e7","0xd73810f4397d30622b36dc3a1c31ccff3d94b5c9","0xa64d00852fae1d13dff8ca0a691564895f25b49c","0x143c94f65b3879e94f2c636fce98f7a64bebfb65","0xf1d1f58baab02ea09167be5147e20e7094520137","0x6c000f66d2239c51ac03d2b23267d5b748e0e1fc","0xd02db07eb758eaa2b96ffa09094030eab013702e","0x789ea342f7a723bb506804b6c4220f4f7cdee0af","0x58538a11db127fac7e2b9e46b2f3fd3aed59abca","0x7026dc43e422ae513d20b8f9871a9d3606ad050f","0x412ff22fdcbf4e7230615c23700608cad97d957e","0x53f78120b67fd1f5a27d62c9eb5e15ef1981e28b","0x843cb541249f352b16c7d357a50d5bdd90e81481","0x6b2e2c5af1b9cbc47549b6656804d0e197331238","0xb0a1a144059051a3573cb8eac3e82e9d1807d199","0xa02311bbae01502ae36af713483fd4324aa4aa50","0xc5cc708d0442359a0ff5c93c4975ecce625aa083","0x9169e30c7fa8ed370142407c65ffc9516b86befd","0xf0f3e69f516b249e107b47355e327c19c34a98df","0x6c77f95dbd22beae7fbc5ab9cdf92bfa1c34042d","0xec58a99bcad82cfd878907982aa06f661bbf1f56","0xf42b25e688c77c1ca5880d34f0fb31883c7b56c0","0x55560cb4def47f19c3272195556797e750f9a750","0xd6054bc3556a81e079e36defd2cfdffaf58e1514","0x0469046e40cc7910224c14ac6e51445958b1f898","0x430f24381b6534c811f616c8ef587443434f70e4","0xc6a14e6943c50a5a29b7d26733e83addb0d200fd","0xef6e476c6c9ad16c2520575872672356679155e0","0x9f40dddba847ac03784e1e85cdb640ff2638ca02","0xa099bc084c03c7badf16f244e926094e6062e242","0xf4e87bba5019977b6c1bf61afbd534d33425e389","0x1e877971ab344f97805ddac76fd2430e92f77f24","0xce4932a4d6d6403d7157fb7e34cad775e733a2da","0x772296f17ded73d47aa15aad335c55afe08276e9","0x444a9787bb6d4ae173f72cf7d71d8e9bc90f0a3f","0xcd42c8e7c78c8a1daafa7cf8fff57d186223eac2","0x9ab5cff9b19b6105366412a3f48e299ae0561a7a","0xd33b34d15c2df6a19df9f0d94750ae175f069a35","0x1a2c1db231b5f1cab5db5ca170421a9c5af3d14f","0x0abddce18c3828b53f3e1a31d67e7623919cf53f","0xedc07526b6b9e915d4e0c82a99912e359efdba84","0xeed6452777f6209ecdddb7800a4aafd35ccbf371","0x9cdd6265f1d3864cc4094a4db6828ccc6b0bfa3c","0x45c54824c8ba03d5f007a7ba2371bdf9aea33303","0xac42e2c70097494323c24d210273c38fdca2d339","0xc61cf6eb54344277687f77788f4f6bf3e5e8bac9","0x385db90719812452d2b1ebb54e08b58e84d41106","0x2dfde8bd3bae11f07b7d6620c7e1a068576f9ca2","0xac19e31a0c6111162b16cc1ad91560b9e45a16f9","0x557f7bbbb61f36f2d6b0b5916d2803815a5f8bd0","0x14d0a66d8d5e483363e31f3b26c90786384cc726","0x37d91da1aad9f7073ee1ab1eb0a3e26d3761ac0a","0x8d2c65fe66b6c7031c7c1a610c1b7ffe9a95d576","0x2e85eacee199c3fcf0df511765cf466f5040e1c7","0x844f567a8858d1cf0e118cc3560d735ca997f40e","0xaa10342cfdf4f624e3804a398bdffa4c2f31484d","0x3a5f306e0599fdefb550bd5fb0811b1a0be2848a","0x1b8fb8154da812504034287aa747544793948002","0xc242477deaf774c6e460eaff362e584a29485703","0xb5b962f539139583e3714e98406af2b2078cc0f4","0x7de7c27c165dd945e884dc92546be181a793aa76","0x92ca1e200d9caf9976378d97262a4c8552b962c7","0x25717987bfe3a375370be97810d27a3fa3b97a40","0xf8da40a4ffd0c239808a578f6dcfebe20f731a8c","0x5d89a8e05d46ccfb8851ac41f793073ca8a71645","0xecf6761deec56ecb9bf7bfc9bfd278a331347e4a","0x8b19ba88665b0cd433ca8a2316655bd9fd7bc5b0","0xc4ee5a604eddfe86d18ce289191b9bf8fbfd7ff4","0x49b092de10f1e876ed783af20f700ef77ee1310c","0x4da4bde78e1474e9b1e0e908435ab74931e21c4a","0xaf11f6ed04b633d7307ede5715fcef52c0c0847e","0xa21ad53846eec77ca609dbb49a247a0d4496a054","0x63fcfa4377a561be6062eef5c907de440fd3ab8a","0xc5959198248c05af0f8681f76b5ede542381dbcb","0x611107c513bc21d94511ba7c02e6cecc409e562b","0x54d0d22e06dfd1c042bee3db92827a2c57a3d45e","0x7152bd82809d2c0f08a0410e570da01ea0b18f71","0xb5db4be76704075982731be5a2f1714e3825a0e2","0x4f8a0c442f93309e0afc9a0fcef65cc6939e81a3","0xc35de6f2e43e55ba155c3efdf5b950a987678c8c","0x2fed2b450a210460a7d8865da2cb69295da891df","0x49de46123eb6e0f851862cb74c26e9be47acd266","0xa17670f5bfc90e912b01ef021621ed42dea3fa86","0xb0cb69d59558ba9458e2a7aedcd9813f81cfec4d","0x23686bf533ff41e31092c8af6192b7429e34b5dd","0x339bff51385634a5ff84bd24a63144e3639daf64","0x214330c9909a5b4f2cf811714d34c2b915ea3ad5","0x6b2eaed88091f734d2409276a3c2e8f1a39139d2","0x6e88de022c16b6e95f816fac0f45a87be16c0ca6","0x6e8de3d419dfa8f6a70faf430ff1a1fa54d0238b","0x254e5be03fad053f2ebf36e5ad2a1c39d41813cc","0xec3e5fc5f80a98fc4c983f5fa84cff920e0bb0c3","0x6a8c9484d901b8259288d185709f931436b2c7b2","0xdc16d2afaf7b9a5bb01de95ead584e90faae3b47","0x14b1fd5dab68ab1f7594a2e42c87e7dc6b75e292","0x91e167c242678d282edb56cfce241ed5317fdf90","0xcea4be80d27e46eafa94b9d5f3645e40eb73085d","0xdbd5610fc2568d4e5c88b200c8cfe9aeaaa675ed","0xb5c45b36603768b07dfc6629c11a87855453176f","0xb6ea1b1131c4b4988d67c8395587cbfa1730c60a","0xee6b25004f4657cef17d362c8aff5fbf1809ccb0","0x576bc2dcd4e09c218ae73fe22f9511fed02ed668","0xf98d7e6341b0842ba8c09ff5550cd86c51df5af7","0xe9a3f84a342612dbbbb4b3cad3211638e585ebae","0x085ab8db4577f5734a05d3c081a1bc79e0addf5b","0x2e1604e0dff44fbc69dc937fb90c63f7f0b10a1a","0x34600b3b58944662e14c193ee0e3ab9fcdd18884","0xcaeb2b0e4d72b1b332058d11b4dbbf276725c49c","0xf8064c95032379716181e36c8d7f07cd4ad9c984","0x9f83c0ecaef7a2819df1ad15b2f8772e4fa978d1","0x029336790913ed03a86fdda276f07b825a31a475","0x74abdc126528c388da66d0965336f2d46eb8df09","0x2b07cb0416d44408e8f1d6eb25484f06141325b4","0x32df400f37243318cee0a20229939ab8d687c48e","0xb93216153c65d2c0b18c06ef15d81863e712976f","0xc221358c3c720fef63a57d39434355c2c59742df","0xa89c4908ef951cd2c2c063e5f4c4bdfba2fbf4c3","0x48a23f2e49090a7b275a7a3d4301f13230773f47","0xe8e333853b1a4e72ad4d9497b707b5bf754ae0ed","0x66b533c4cd000cb94917bf59aca5f0922598683f","0xf98556ce060893440a3c24cf381bdb1a56e4dc19","0x97049ec126c24f413664a2792baefca136e85d9c","0xa3ac59edf04442c02ded43ad409b2c7e41cfac5e","0xb7c9e820045af0ae2b406007b44820da4bc4e8cd","0xcb8a719a8a44d533df2ecdc27078bc788bc5198b","0x61aab96ab83870a0bdcd3654db1c3e84523e4506","0xf5eb709b33525546649fef253c1a29eb6c1b72c1","0x40ba85c0d33a9d95655a6c9f35ef4ebaa4ab86eb","0x3fa94172d0be9fcd309f693dd1a2730c2736cfe5","0xee409ab65f4663931854a15ef3e4cd16f6bcb227","0x86a150171f0b825d11e0b87d3be99f3c6ff82dc1","0xa45f3c031e2e590f3cf84edb6403ab0ce48b1655","0x02658c53a9b63b6f902030c3f0f2357a535b3301","0xb6b580b177f34fdca4878e74826d3277a4bd7701","0x51baa204128215fc695b01e4127ad6a1c6d89ce1","0x6cb53291251373752f25888bb399a86108600cc3","0xba67eea0d94ef50bacbca47d72d871e19b8da322","0xeb29f7c82a3236add2bbeb737ca39d456e024641","0xbab2a9a167abd05187e2885b269b8782fd57a116","0x7a0103a1d03996830c11385f79c9c7045e67c9a1","0x5cd0e49a13f4b298586e462948ad99ae42bb7c38","0x43cfbe0fe238df99fd2d1b7ec134b52fecd01acc","0xb8034ca5eeaedd118cfb75ae849fcba077824706","0x66482d1e81c34a3e6d00fafbe320967aead680f8","0x3d2275ecad35a56c7067ff3ae04ebfd21a0e8803","0x2f48af5b41ed95d850dd939f92008b0e61d116d4","0x0ad8753e9f909844648268a19ae42a43273e6438","0x8304bf45209c53c7a40587dc00745596e984a90c","0xee8ebc423489623a1d4cff3577de24f5fa446fdb","0x277ef1f8c89f7486aea1b635d727161cb057a7d8","0x5ec02b4b9b2b8fd4d71b6e611127999ad1c8f8b4","0xba918d099d0efb6b71e758b09be3a0ebc533b461","0xe9a82e01bc6d1a4373de6efec489c342b841a064","0x841c47a05f0e73ff16424d4829cf6785fc6f0981","0x6c6a56395fd321b0fdf40371fcf62009f45dabad","0x2d2e0c139e6d083c3ae10b918f9a9ad62f47472f","0x447f14f24c5f6a4861bce043ef9115f56862b941","0x0533d9f716825610ae4d413ac836f8f07e0192b1","0xdef44517c2c20a89bb55c1e228c8b4374245d2b1","0x2573e94930177553ac3d483f91f9d42f85da5572","0x75975258d9ff07a2a9331756a19401699b46f860","0xe256e317689226a51fff9a8b52814f053e436635","0x9137b0958449c2600b3901df8f67a11f67fcee8e","0xbe384a20f0c1132656742d3a9a3f121cd38c6e73","0x974f07b528c92b63f0c9cd1a9f999267daf7f907","0x909852b0dd1f54ed5e523bb3d217ef3aaad1649f","0x4a6116f3085ecc49060b99c43269722d43a878a7","0xc33016ae462e759994f4d921126a3c2380c832e2","0x1674681f9277a658fb28eefaa388c1067fb442a9","0xffa72dc3f5a4f6a1b986918e9c84ecf2d056a5a0","0xb200b996c17d0765ab7ada7f454605d6764ed846","0x0df252ec883c900c1b9787d0ac1fede0998dbca1","0x1a9ceb2a73abf74a5280a9f5dce428015e863028","0x09b4620bdd048eebbc840506441aba343697e5a7","0x01539725c92b942a28897e95ff0e41b64821a699","0xd5edbdd1545fcdd4b39e1a2e3cebef71545addf8","0x901a140c4af49fd98a8c2234092839a59039e412","0xe52f1125642359e87c3dba9790eb6f6a1db504b2","0xa08dc2ac6c922dd1e9935175e7f842e3fd364250","0x158f4d50aa73bcf5c9e67fda682b020cbb4d2c34","0x2450e41fb788e7ce323be0ce1ba9692799037776","0x558f2db0378a8a2a0571f292656cbdd62cd81670","0x7151170d372a1ca306e89d7854bb2eca329cc078","0xb8921fe4b561a81d633a7fd74e9d33d0879e22b4","0x154384b276e2776ef318b68070f2ef4559090056","0xc0c8b614c8fe147a76e0d6bc18776179018f53d5","0xf5270b4654f359bb1537589a4a775a047a204485","0x54cf0233f790095ac447e272b659457af4c49c0e","0x9e2da98fc24ccdb8a061ef0cd76cd3928e377f3b","0x3acd61d278d829483bc0f3996b3a670cceb4f3f7","0xce4f5a1983c441539c47deee4ec96c86ced3f908","0x1b12b7d1d4737b0b2dc927e35cdc82f775765b6b","0xbdaa9fda47285fc49df4c72a8b0af64ded8e6b54","0xc9c693cc1031e289c832280f7497e8554600fdbf","0xa910dbc53e757b232359e619a64889ec874d51eb","0xb565c9762ffacd18ce0960a163c507fb0d47b476","0x1f7b9e2e4ac2b57a95f03bdbde7be8e8d4a5e436","0xd80f9e4e2a36e1383b54a19352e9fed54f62e28e","0xd001ce773109d78dceb454cf61986249ccb00f4f","0xb6937fd9f49b68d5f901d87fe27b646ed1db3512","0xf9061a6e57bd2f49912f57bb48c7705384fef4be","0xf15420bc2fe0953512266f6c73fa1d22f026448e","0x8b55c592ebb8f10895021a93287e01318de4c665","0x7f73244c55b992e9eada1ba95372582015458bc6","0x8c405505a97dc42478d51a08fc35553b775fdb5c","0xcdb8657eebba335182f4a38e6f0db3720def28fd","0x82cebe5eb9777664711a8d1dc66af921665d065a","0xbc286cf29df96df18001382a68e6967f59a95bd1","0x4bded6a82dcb847c1dc36d45752db11cb89a43dd","0x380c7acca0b91e796f8c426e889a25276d998efc","0x5fad34000c0cc68f50cd0546c55f85660468eee8","0xef581055208a061175513d00a8ea2b24813562b9","0x99d890fdcf1f655d1b0f9f40bc6c26e5df6a65eb","0x7f78afefb2d1163aa2e06e2f41bd088dbd8b0272","0x2ff7afa90fb3b9ccfe4492130fdd627488f0235b","0x09488c71e799bc683741d990572e1528821ddcc5","0x085baa1aa85ec000b74b4dd246d831b54dd804d8","0x7cdf78c84f98c3bc9bb08f33d9190185ebec7d48","0x297ece93c10ec7c11b25d11df22dd83cd99cae1f","0x121a712ace58d917be5bd5c86fb650709f9861e6","0x2470ade54d25d0ed5be0f030521e5db5b5ca8149","0x65dad0296086e2bc7c1000cf9abb5ba6f375b046","0xd919002da5689683b7c9f1302aebc52bde9a33cb","0xfe1c83a0637112310df9f1e999b154f5f2fb66b1","0x864ff038f2272aa1d435e58da4416f21dea7d746","0x10b12757c93923efcf40ed0ef2f4c612c18d1616","0x0e5466a031b72560bcaff53a3b6cfaeb62fd71aa","0x3ef761cb79d4e0e84f69fc088abca882e8b59ef8","0xa6aeeb03200843f3521cfee9caf8cb928567f1bf","0x47717ef58e6a4c85388fd4554a6ae992b4d8bd4f","0xfbc025fa56ef25a595c705c6658d4c4e6fe8e8dd","0x033661b5d736cf3a6bd00f73c826675526a80ff1","0x73b34665d1b94c3e6504df568dfd9b99dbe1f9b6","0x82f88b31762577baeb74252ce104c347cd822a5a","0x5ae9212ef0eb09da9c86b94540034e4fd0ac17f3","0x166f9f12d12607389d0b031882d7828cafe4a6dc","0xbb8a901fca75455478644aa6124415b62e36882b","0x0f13822341b319ce3686ef31e31779d7213aa5b5","0x5c29167988b40b29247fec339d67720c34a3e910","0x564773055bd40aba1bb6dc1c562899546318188a","0xc5ebdd9a1a40823e29d9f2f7baa52d635965ce64","0xb10bb3503f5d8d708a813d803c1a2bc57dafe9c0","0x0e00e0d8783c0e8eddee89a54991200bd7580519","0x0ce77bf187eab5df41bcb1e9d72726b1405c9f05","0x0f6895b0735569136a584435e501ed15909828f1","0xf1d898a04e954badc846e0df993d20a4de8e97ba","0x31ce4f459ad0d9694a9284c59c01e49b0d3f4c7f","0x2096e5924efbd8719927b884c0927bbdf89af902","0x79e4b7e07c69655b73bc17ddae129c2858e63db6","0x23d51b8922e93e49b00d65b83ce0ca57639d8c4d","0x8d2aa345e91e78232c9236fe9a8067bdd7129e7c","0xcfafd6ee83a9e57f6f9f30af804cdcad7b22ac6a","0x916674c38bb14f21746e1e2284472550aa9984fb","0xaabeb552134945c900c09d84e49acf6650c98181","0xbbad9e0648bd1603bb044b8e6cf71e6c2615af5f","0x890e293ced1488c8d863fcc9ea562aac028de869","0x04e50fc2264fe4d57734616bc6e6ec6206690e26","0x9148a3f82d798a6c7e301b4b92f379889979cc68","0x76ae873aff0cead287615d6ac7f980d83410c71a","0x01fc32c69399828df5a3d0d23f107056435bf48d","0xd53164b920f88f91fd339d8e278851ab1a54f2f9","0x541bc8a55909f479c2d475b2d472d29c9ec17cc3","0x0c7779428a567535289bd4a4f8884082346d0a7d","0x8390e57de9879b14bdae1b7c926d8a54f3c9a938","0xc744b7db5ec3fc92bd19e5ae21423f598f77bad3","0x22882ec8f85dfb3795095e0709e0ece2fd18b0bf","0xdcd73d62c39545ff94fb17a5e085e7996c497718","0x7b3c4887048b60e8818e4a94a7c892adbdead755","0x5f4f79427b4fd74a399e780cc0dfa49b1dae6378","0x57865422eef064aed518d81099911af3c946139b","0xcd6e97167e189969ba6f3a73d9cb77d129dbfe28","0x97b7955858890b1f9409a5cfe3caba90d114dffe","0xb4f05813582fc6d59ceb3a0836fb94a50f5abc7c","0xecd9dfd752be871031bb459965437a1c5a2d0f7f","0xd38d3f5c4a54f88c1fb5d8d66b0cdf5082853afc","0xa21e85af1b9eed2df777f03141c6927f9a578e53","0xb21cfc994e44de70d97fdc87f1b533c44b99d87e","0xfce65a3aef5a7eec0d7c122d9057db48b39c3b02","0x3c179465fc3c5cf182ddf44baf6a0e4b6298e1ab","0x72335a0ae533acd167e13a30240f840f91c30bec","0xd1dd63d346c0ff8a1cc459225789508ce23043bd","0x14e59039c455afd8a0618be77245439fa4fc376a","0xbfbae5cc75d503c5ca0d69361bf39f8d79e23b71","0xcd9b665bb0f690bd77571973c52efccb5edc4f16","0x2a9ebe79aa09c8befb4b17c29c36bc0e8296a436","0x10d6a7f1cf050ed76d04500013a0f5c926095ad6","0x01de8c5bd8050029518fbff6038817449f1d0b3e","0x07c00a5cf58c3e7e56b1d0be0dffcfe77078ae36","0xb207e1f23356afa8395ede3c8367ea0e8f7bdc0e","0x494f3c38a0d4215737bb5b8007e773151d5fcd53","0x347688fbb0bd5450f080d4eff0b802e22d32d5c3","0x26a39d761fe10d7c9aecddaf3408f2cc73a0f3d8","0xb5d18ac50e0520d0a7067c44c7ae8b78e5af0114","0xd350387fa9bd868a8652fb14eb8753eb91439014","0x912c1407ae9a3ac71b84da33bcd5290587abe488","0xe198857cadb5f91cdad6d9c86f0fe89166b786e8","0x67cab74be88bac1ff4cca59e7dfaf47f43131d75","0xfa23ef7361468f93ac5c30568ab549f4b6feff89","0x35f8b36f5c329497242bf3782c2bc589455083b7","0xa17145a20294adaee284e69a80bf8ae7747f63b2","0x931c3c240cadd20e1a5aff4f17071f53eea05ff7","0x62759a751072cb63a7b67a1d6214b8508829cdee","0xb40e7e81625b6c73e5eb94dba4f8ebf91b2b6fae","0xccb7c836b1bdf70431e87d9ccee579824aed0c0e","0x08e4c6ecabb7a513aec9fea0cd01b6e61d577442","0x82e959dbf3204f5dc1feb3a838e867e63ef36dd0","0xb55a1ff86f1ccd15b8699b58bb2e69eb21928d50","0x4f39a4bca4aa36c316baaa2b3d771c61d99bb9f1","0x229e7ba80b4b3db12b4f69cc4a6b669c716b320c","0x00b962c4115c0952a3a1521f3430d8fe43f33fa4","0x8ea8beb3b2e6d7846add8b4c8e118bb4aade98b9","0x4d7e37e4c8fbdd5d86581cbb018fbeec677de031","0xfed445ccc0e32ecce5f63230d459984bb17e7bb9","0x387070060a5081b7d6830d119b7283f9f5a86243","0xcb5901608cab4ef6351a107320f93627d37f86ec","0x6c6b493b149f940ea8b08b79ad5d642f265443ce","0xad74ffa0baecdbe2580def84f07f4242725eed5a","0x67a34872faa6de3dfaa449486104abea70c61d47","0xd54ba0559a7796f2d1ac257ca6804c6f5c4cb6f7","0xf09757e187c8789a12b9b6c407a47a7e1ed183aa","0x08d1ebc9464307d2f0e9203b5fdb181b67d8af01","0xe8e969558e659228cd12d382139a707f95127ee1","0x5a8394242b04b5961f2e183d816299124698c55b","0xfd037f238032327b416b5e39c222137d31ccf1b7","0xdfd3b4ad7ef6f7ce5a3271c53094a6ad21d0ec81","0x6f47c8fc583955fa94a496ced71ad04081ebc19e","0x823d0e75248cbfa6f073f16f3d5e02cd66fc1664","0x52752ade3e0cce3afb2f8f925b9c1bd43c14800e","0xd51d6f79fa21aab94390e9122bdefd4e10efb6ce","0x1a4376b650982d890982f0eb05bcc899dcd5fafb","0xf13dd4007089a9fededededa445ecd0a3f954d44","0x23aded569591c8b4b11cee8f76cef49c1c746db2","0x7fbcd6589ce1c711919429447b326b96a9f9a1a2","0x4ce242d51e0e2742aca6ae562d9f5aa34f181c30","0xaaab896a342689354cb976ae46f3d93ce22f2e20","0x77c7ebcea53b301f6c945f6a0f1e6fd7d1a05b79","0x193287d0e924106a5cd90810417893e5b86e0cd9","0x7cd3f6980e6ef5bdcaced88de3a75818a885fe98","0x1e009597aae5e32c0b15c42abb169196bcd5ac5e","0x57bb56c07ae9ea66dbf15ea75cfe12a54bc541b3","0x1ec94e73127a7dd086c7837579730cfb6d758d7d","0xd6eb8eb62ced9ad2b79727d6a7dad1dd7a800d2d","0xdcc2969c4fbaf9761a8882589414c74c388ff02a","0xdaf9aa7481c33ff953765febd9a3e8d3cc98eca1","0x19a4baf9561e3baa580a9955c01ba5c15a81dedb","0x0d6468c95b35262f2ce4fc7a2c4402ca1afa4d27","0x57f03bebc977a49de78cddff96704f4ea922ac65","0xca5116dd5ea6db38be2e310841e8b92c07a99f57","0x44156f46339ab176241d484f0a6a2f0034fc7a93","0xfbbbc308ca8f305bde1907d3f1e02b8029ae7f8a","0xbc46af21df5b53ee2d4c1af141de2edc8174f8f8","0xb5abdcf0a7d64523da5e9915262f6c7d1a0ccf04","0x9f7493ad51f4b08fdf51b7134dc1ab3eb177b7db","0x034a66b538a494bdc47577da0ea49305726c9231","0x8df02afbedfeb4f48548f031cc089ab9f85141bc","0x6855d2ce8112b22e0d94012e60f9fd74547314f1","0xe45ec3902ee96475b9f23d4b53287f4cd966932f","0x4153a764501c1a629a23cb1e9c1b15f2b88ea8fe","0x9457f06d54e9d84c002410a71ba29417da550434","0x643a6a1640957ade26df7ac4bca02b5ef6918fc5","0x1ac8512c573114afa0f9b3568c70ce16b1052dfd","0x24d322cc4028f909092d9eb1208211755f9e6912","0x51f5f2f195a2fcc847fe3a0f498964501820a0f6","0x17506cb42f4a836e11663686b75ea6156b1475e8","0x6ae38bc6000379966e6896da15678178cc5a4ac5","0x0769eb931e07a16d461edd619621219e5a9acca7","0x677a743888266af329b775c85eca67190eda2163","0xf0f678186e33cc30834e2553e0963d5f379ce386","0x908b497b342219e5bbdb2cb7c0cdba9bc107bb22","0x289481eceaa97d5e37086fb33cff267a7ebcef65","0xea13a1dcd72ee79e74701fbdf04b3708a4216468","0xb771ec0f93ef6b037b9b97d7a7ef1d227f86cc93","0x04f61fb4ba3835585d7998cb2e904d028c36b70c","0x007319cceeaa4ca8ade8f02a9ae3cd95a42b93f3","0x2ea1bdc9ee45aa191228b7c1ef05a074881ae4df","0x545b85be01a74a2e5def0aef8f438655698b4109","0x741f88cb56fb7f452df06a53fc330e102e144981","0x6da2a36ad2aa3f6d3c95f0edd4b032e3a9ae8e7d","0xd2fc0c9c7d3f8bcb21ee1d13f329b69b8168e3ae","0x52c7b210f4aabef126e224abfd47ad9ca7cf999a","0xaa904432fa4341f129d474955dd8d5d90b340d4f","0xb28e68a936833b50dd6ba27959727bd60642a2c9","0xd2085ee3c1c39622644dea1c72ccac67ac443fef","0xd75923a7d130454fb5eba37e4513b900fdb9a2b8","0x5875082fd285f55f0fb617161caa4eae9549f43a","0xcbd32252809ec8a7a807ac3bb74fb93d379ccb6f","0x4be852589e77e6ea846a08ea678c6523243477c5","0x1abb01c0b3c139f4bd48872377074a548a85a086","0x6706b65f2c60dc19602e173b36d5c6006b511312","0x3af723942827225ae35ff385c7a4749aa26354a5","0x90dbc20a97a184af4aba57fbd99d93f54638949f","0x5c91bbfec64e22b64e9d2b06fba02852af132cc7","0x85c7f0dd9389ea5339bc08fdc8f4e85b081da527","0xe2d1866c122f96d3124d4e6d10be99b23a05f022","0x229cfff914e04c8a1931bb85b668f7ff09babcc9","0x1d05a0db9327f759b37312bc6f25d86f6d97d345","0x92e7799baa87a33b715f928a19ecf3944e51bb60","0x54f19b6656c75d4ec36af0c0e1d35a7bf7a6e09b","0x6ff9001ec54b848741ec85c0160e6fefe64610f4","0xae7af022fa8cdd922158d4784914f83f078a9303","0xf2d210c721e1c83a62949d9ac43e4c35f31672f7","0x7248e2b77d1236182d91cf770657217bbde7e396","0xcacff597276467f46719220e10e0dfaa1d22d7dd","0xf7681247f507fd5807899cc67b335b73fa113302","0x86d0f0220c441685f7468405047667aae40cbb4a","0x9a950ab338803e5c807182966ede87b452d39711","0xf5dee2328d460fe1a58146b20ce74ddbd8ff8f31","0x07a2405c48ae415e16835abd3b5a232941095baa","0xf9b3dcbcbd133f105f263ca8f3328f4af1249fb1","0x66e63183ab848dab4b510899dc5de903e07a4684","0xa450f6ab105dc5bfe61d9afdf57b0ea61d8df7a4","0x4f4300f036c0770853770018ab952c8c54ae7bed","0xd6b79acd693b42b43db3935b17ea2df8a0bb266d","0x9b23b14ea302af7f93ae113d356fed5637c4021a","0x176972db2a971406f13d30364bdaa5c13e3fde2a","0xc63367bc04058fd090beace0fe3282bb0475b6d2","0xcc76708a0355a1ee0684930a04c82a7aee54b77f","0x2c7c8ea3aeaf1ed00ffd0a5d4826d0b6d314ce5a","0xab460892e2df0bdb9c10f2bb7079e6f5ce9a6276","0x342ed52857587f70c0d4ce1e28e4f8588f7c497e","0x24c9144323a23b84b7de2b7aa9f52e7bc0066faf","0x06cf215de35818d05df86d8eecdcb136c27aec4e","0x5d9e4a931541f9766f026918b6dca47afff76ebe","0x6396cd44f48d6a9831186608e7a53016eb35009f","0x6fd5d6096631f55987e2f88099f7831abe6fd54a","0x30f5caa395bc3797dd941ace47271ebe2c92464c","0x4ed2966f8393ebf696fe744891861a828497714f","0xfd79e320c6c3bff4259571c7704bd90eda882892","0x8a923a618049fb02e4f64eea7da705e81bf06ddf","0x22b8707cd21d403ff4001b5a0f6b04bf8a1344a1","0xb7f660e134fcc08cfa94b2f99142fcfc38ed0098","0x782812ecf2a390cfd8fc1b9e4bb14497d8056f76","0x3a28893e62f9b0b68a33c86fc00727ddf022ad86","0x91555a7d3c0d3591ab211a5de5616dd802b168a2","0xd75da654e6a8a3a4417607ccd549152fded23e49","0x4be50d04ace1b3f345237c813151b9816b6a9947","0xf9e63445b295e1d27a2f57831173b966585d4d60","0x7326e64d4045efb20c27bc8af6abeadf20945c79","0x22392aad12963b5a985f263b6ada280b2d86e98b","0x5c59250bd23e6ccf112d1c62a14a87078b84fe70","0xefe213da830527654ddc5686766f8148f86a28f6","0xd8afcd801344ecbc9fe59bba544a393c70acb1d6","0xd281374b795ccdc875d04161c26fc82c33d332df","0x605913d9c4fe334e596f6f66effa5e3b572182c8","0x079a9445ae06980b1c68e0ccc783b35e284c1c67","0x1cc4d13124b479e35e9f930852f6cffc0132d0ad","0x53f55d1aee736f2bdb45aa578d5c5fd46b2a3697","0x4e5149500a90dae83343edf89843e67c16cd9666","0xff235de6eb8f9bca8c1ce2a7db9534709cf74bd4","0x16f9b5ab0903092768f6b1aa33c22992e8a30c45","0x9e2ec109f374c8e9451e5e632743ccd49bcbdbd1","0xd8811bd4d1ff378349a60d7de5c33fb4b7c8e145","0xbc139c8f9a56e21b4b8a68aebec41067e35b00a4","0xe59bdd0b1763c271cb6e97b89dbf2eb88331890c","0x6b01ca54051e49b1adfb5ba8df038167bb440aa3","0x6bd7d388e9a4dfe1e9a78ab149f0167e813cce09","0x50f87036e95e7b6f6e22203efb17ff098d7233a1","0x8098894bb8740cf99b70095b9765c97526958387","0x1cf56507de8ed962980d46e8867ea866f3bffbe6","0x62758d0a589cbd644c26d6e50db6db34154ffa9b","0x6aa1f548470c61de00005359b995a00e7bc6165f","0xedf468a310213a0c782595e2c765d51e3426b993","0x796211209ad9e8fc9e8617c463ea1469fae60c00","0x761900a9fd88ea68b277fc17aae7324c50f1c7ca","0x353963dc4dd20e4f6471aca125d6ab938cde01d4","0xd63877525cc15397e104db9f48a001ccfbc58fe6","0x2fd05847ba97966b4f15a7afa1253bcd0d732c06","0x3835c3024ab74c57f00f2e2a3d7c63face743e7e","0xea03ef8026faf406dd8249ea40366b5905edc6c9","0xaafad3d02bd382e5d1cd1474b3fa5f364d03975d","0xe168f96366c27e5e7de3298fe1b01b1459d8cbfa","0xa2d6b00a7864458e7b35d14dd994143b85aa0194","0x81e8c02d027e7db70fa5d52cb90840b54245886e","0xc91adfdb6ad355ace25194f032c5af452e3b1d69","0x7405b5fcb1895b4df8b2986925f3c14d4900c894","0x8cccce4e4b43e968d66377efc05a3bca0c1d80e6","0x65957509e05543c2953ef145412e1748148a182f","0x87154250300adaa67154839567d937429835259d","0xb5b5c4dc9a27b7187febb3d25024417475fd4677","0x334ba882ccbc5efbfa6c78fdfbac740d0ca00d9b","0x4c6102fc2f123ae6fe2e81d96b6335169c839dd5","0x2cadab00407014e532cdd852fdb606bb2661ee23","0xfcd0a71b2e68f71b67eff8913e21da16f671058f","0x648143617e17abb82a0e2c7d27351564585f30be","0x8811c903c2a9b772dbb2ba4863f00d92e495f253","0xb93c7e617bbf921bee23a9309e6469774d2b0e16","0xe6999accad77359222df31f792c3350f3cf0c137","0xb5ddce67faf4b7886e892996a14ea8bcd218395c","0x4af26d16ffdde298c41b9f682eea5f6bb1e57d59","0xc4ebbf2616c2a7e23f0e3b83621dc8e4740c2914","0xfd289ad412c715ee8f577aadffa60c8cc3591914","0xe770e31095c6cbca4d12b48019fd0b669e1531b2","0xacc65b7f4dc842449a10039a591e1e6ff4f2c339","0x9f8fefc9e441c8d7b5f3e0ed94e0e73e0d8972b7","0xe729257f139c4bfce982c809b57dc1adfa1c5bb0","0xdb23b94699b49baffbfd56ab42eb18f37b8dd395","0xdf3b14769cd163fc2dd9935f43280d3dac194ee0","0x6b588cc1b5d0325630db2993c4334137ff45acde","0x1855c911f32005fb64bb331df611cc44c406a616","0xdaefc682eec5de3c1b7fdb335bfde99beab999a5","0x46ef624d4c717e6ab3de9cf6cfdcafb9fc0f31a6","0x66e19cb7b6e7197078a1ec6baf249804dcba803d","0xc234c2552f7e6a6ac290ba8ffe172dca41401b86","0xcd6c3e1402695aacc0c1b5f4a3f6358f87bc4604","0xf1d9a19978eea988cc02af077014169436fe8aee","0x924cb13986b138e2c1e3434d970eb9d80f24ddc9","0x604cff63c89d0f35351ae87e5768baaf2739fc0d"}
	//for i := 0;i< len(addrs);i++ {
	//	key := GetMESHBalanceKey(common.HexToAddress(addrs[i]))
	//	val := state.GetState(MeshAddress,key)
	//	newVal := val.Big().Add(val.Big(),blockReward)
	//	newvalKey := common.BytesToHash(newVal.Bytes())
	//	state.SetState(MeshAddress,key,newvalKey)
	//}

	return newValidators, nil
}
// snapshot retrieves the authorization snapshot at a given point in time.
func (t *Tribe) snapshot(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := t.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}
		// If an on-disk checkpoint snapshot can be found, use that
		if number%checkpointInterval == 0 {
			if s, err := loadSnapshot(t.config, t.db, hash); err == nil {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}
		// If we're at the genesis, snapshot the initial state. Alternatively if we're
		// at a checkpoint block without a parent (light client CHT), or we have piled
		// up more headers than allowed to be reorged (chain reinit from a freezer),
		// consider the checkpoint trusted and snapshot it.
		if number == 0 {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				validators := make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraVrf-extraSeal)/common.AddressLength)
				for i := 0; i < len(validators); i++ {
					copy(validators[i][:], checkpoint.Extra[extraVanity+extraVrf+i*common.AddressLength:])
				}
				snap = newSnapshot(t.config, number, hash, validators)
				if err := snap.store(t.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}
		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}
	snap, err := snap.apply(headers, chain, parents,t)
	if err != nil {
		return nil, err
	}
	t.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
		if err = snap.store(t.db); err != nil {
			return nil, err
		}
		log.Trace("Stored snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}
func (t *Tribe) blockTimeVerify(snap *Snapshot, header, parent *types.Header) error {
	if header.Time.Uint64() < parent.Time.Uint64()+t.config.Period+backOffTime(snap, header.Coinbase) {
		return consensus.ErrFutureBlock
	}
	return nil
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

// AccumulateRewards credits the coinbase of the given block with the validator
func accumulateTotalBalance(state *state.StateDB,blockReward *big.Int) {
    val := state.GetState(params.MeshContractAddress,params.TotalMeshHash)
    newValue := val.Big().Add(val.Big(),blockReward)
    vals := common.BytesToHash(newValue.Bytes())
    state.SetState(params.MeshContractAddress,params.TotalMeshHash,vals)
}

func GetMESHBalanceKey(addr common.Address) common.Hash {
	position := common.Big0
	hasher := sha33.NewLegacyKeccak256()
	hasher.Write(common.LeftPadBytes(addr.Bytes(), 32))
	hasher.Write(common.LeftPadBytes(position.Bytes(), 32))
	digest := hasher.Sum(nil)
	return common.BytesToHash(digest)
}

type bindInfo struct{
	From common.Address
    Nids []common.Address
}

func (t *Tribe) getBindInfo(chain consensus.ChainReader,header *types.Header,addr common.Address) (bindInfo,error){
	if header.Number.Uint64() ==0 {
		return bindInfo{From:addr,Nids:make([]common.Address,0)},nil
	}
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return bindInfo{},consensus.ErrUnknownAncestor
	}
	statedb, err := t.stateFn(parent.Root)
	if err != nil {
		return bindInfo{},err
	}
	// method
	method := "bindInfo"
	data, err := t.abi[ValidatorsContractName].Pack(method,addr)
	if err != nil {
		return bindInfo{},err
	}

	// call contract
	nonce := statedb.GetNonce(header.Coinbase)
	msg := vmcaller.NewLegacyMessage(header.Coinbase, &params.ValidatorsContractAddr, nonce, new(big.Int), new(big.Int).SetUint64(math.MaxUint64), new(big.Int), data, false)

	chainConfig := params.MainnetChainConfig
	if params.IsTestnet(){
		chainConfig = params.TestChainConfig
	}
	result, err := vmcaller.ExecuteMsg(msg, statedb, parent, newChainContext(chain, t), chainConfig);
	if  err != nil {
		log.Error("Can't decrease missed blocks counter for validator", "err", err)
		return bindInfo{},err
	}

	var out = bindInfo{}
	if err := t.abi[ValidatorsContractName].Unpack(&out, method, result); err != nil {
		return bindInfo{},err
	}
	return out,nil
}
func (t *Tribe) accumulateAccountsBalance(chain consensus.ChainReader,header *types.Header,state *state.StateDB,blockReward *big.Int, addr common.Address) {
	//get miner bind wallet for receive rewards
	bindInfo,err := t.getBindInfo(chain,header,addr)
	if err == nil{
		addr = bindInfo.From
	}
	key := GetMESHBalanceKey(addr)
	val := state.GetState(params.MeshContractAddress,key)
	newVal := val.Big().Add(val.Big(),blockReward)
	state.SetState(params.MeshContractAddress,key,common.BytesToHash(newVal.Bytes()))
}

//TO DO
func (t *Tribe) accumulatePOMRewards(chain consensus.ChainReader, state *state.StateDB, header *types.Header) {
	// Select the correct block reward based on chain progression
	blockReward := new(big.Int).Set(MeshRewardForPom)

	number := new(big.Int).Set(header.Number)
	number = number.Div(number, big.NewInt(int64(BlockRewardReducedInterval)))
	blockReward = blockReward.Rsh(blockReward, uint(number.Int64()))

	// Miner will send tx to deposit block rewards to contract, add to his balance first.
	key := GetMESHBalanceKey(header.Coinbase)
	val := state.GetState(params.MeshContractAddress,key)
	newVal := val.Big().Add(val.Big(),blockReward)
	state.SetState(params.MeshContractAddress,key,common.BytesToHash(newVal.Bytes()))

	//then miner send block reward to POM contract
	method := "distributePOMReward"
	data, err := t.abi[ValidatorsContractName].Pack(method,blockReward)
	if err != nil {
		log.Error("Can't pack data for distributeBlockReward", "err", err)
		return
	}

	nonce := state.GetNonce(header.Coinbase)
	msg := vmcaller.NewLegacyMessage(header.Coinbase, &params.ValidatorsContractAddr, nonce, new(big.Int), new(big.Int).SetUint64(math.MaxUint64), new(big.Int), data, true)
	chainConfig := params.MainnetChainConfig
	if params.IsTestnet(){
		chainConfig = params.TestChainConfig
	}
	if _, err := vmcaller.ExecuteMsg(msg, state, header, newChainContext(chain, t), chainConfig); err != nil {
		return
	}

	return

}
func (t *Tribe) accumulateRewards(chain consensus.ChainReader, state *state.StateDB, header *types.Header) {
	// Select the correct block reward based on chain progression
	blockReward := new(big.Int).Set(MeshRewardForValidator)
	number := new(big.Int).Set(header.Number)
	number = number.Div(number, big.NewInt(int64(BlockRewardReducedInterval)))
	blockReward = blockReward.Rsh(blockReward, uint(number.Int64()))

	accumulateTotalBalance(state,blockReward)
	t.accumulateAccountsBalance(chain,header,state,blockReward,header.Coinbase)
	//t.accumulatePOMRewards(chain,state, header)

}

func backOffTime(snap *Snapshot, val common.Address) uint64 {
	if snap.inturn(val) {
		return 0
	} else {
		idx := snap.indexOfVal(val)
		if idx < 0 {
			// The backOffTime does not matter when a validator is not authorized.
			return 0
		}
		s := rand.NewSource(int64(snap.Number))
		r := rand.New(s)
		n := len(snap.Validators)
		backOffSteps := make([]uint64, 0, n)
		for idx := uint64(0); idx < uint64(n); idx++ {
			backOffSteps = append(backOffSteps, idx)
		}
		r.Shuffle(n, func(i, j int) {
			backOffSteps[i], backOffSteps[j] = backOffSteps[j], backOffSteps[i]
		})
		delay := initialBackOffTime + backOffSteps[idx]*wiggleTime
		return delay
	}
}


