package tribe

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"sync/atomic"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
	"fmt"
)

func (api *API) GetMiner(number *rpc.BlockNumber) (*TribeMiner, error) {
	add := api.tribe.Status.GetMinerAddress()
	ipcpath := os.Getenv("IPCPATH")
	c, e := ethclient.Dial(ipcpath)
	if e != nil {
		return nil, e
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	b, e := c.BalanceAt(ctx, add, nil)
	if e != nil {
		return nil, e
	}
	return &TribeMiner{add, b, api.tribe.Status.SignerLevel}, nil
}

// chief-0.0.3 show blacklist
func (api *API) GetSinners(hash *common.Hash) ([]common.Address, error) {
	return api.tribe.Status.blackList, nil
}

func (api *API) GetSigners(hash *common.Hash) ([]*Signer, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := header.Hash()
		n := header.Number
		if hash != nil {
			h = *hash
			n = api.chain.GetHeaderByHash(h).Number
		}
		api.tribe.Status.LoadSignersFromChief(h, n)
	}
	return api.tribe.Status.Signers, nil
}

func (api *API) GetStatus(hash *common.Hash) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := header.Hash()
		n := header.Number
		if hash != nil {
			h = *hash
			n = api.chain.GetHeaderByHash(h).Number
		}
		api.tribe.Status.LoadSignersFromChief(h, n)
	}
	return api.tribe.Status, nil
}

func (api *API) GetHistory(last *big.Int) ([]map[string]string, error) {
	s := uint64(16)
	if last != nil {
		s = last.Uint64()
	}
	cn := api.chain.CurrentHeader().Number.Uint64()
	//history := make([]History, 0)
	_history := make([]map[string]string,0)
	for i := cn; i > cn-s ; i-- {
		_header := api.chain.GetHeaderByNumber(i)
		//_h := History{_header.Number.Int64(), _header.Hash(), _header.Coinbase, _header.Difficulty.Int64()}
		//history = append(history[:], _h)
		k := fmt.Sprintf("%d",_header.Number.Int64())
		v := fmt.Sprintf("%d -> %s",_header.Difficulty.Int64(),_header.Coinbase.Hex())
		_h := map[string]string{k:v}
		_history = append(_history,_h)
	}
	return _history, nil
}

func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers:     make([]*Signer, 0),
		SignerLevel: LevelNone,
	}
	return ts
}

func (self *TribeStatus) setTribe(tribe *Tribe) {
	self.tribe = tribe
}

func (self *TribeStatus) GetMinerAddress() common.Address {
	if self.nodeKey == nil {
		log.Warn("nodekey not ready")
		return common.Address{}
	}
	pub := self.nodeKey.PublicKey
	add := crypto.PubkeyToAddress(pub)
	return add
}

func (self *TribeStatus) GetSignersFromChiefByHash(hash common.Hash, number *big.Int) ([]*Signer, error) {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash, number)
	r := <-rtn
	if !r.Success {
		return nil, r.Entity.(error)
	}
	cs := r.Entity.(params.ChiefStatus)
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	return sl, nil
}

// 在 加载完所有 node.service 后，需要主动调用一次
func (self *TribeStatus) LoadSignersFromChief(hash common.Hash, number *big.Int) error {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash, number)
	r := <-rtn
	if !r.Success {
		return r.Entity.(error)
	}
	cs := r.Entity.(params.ChiefStatus)
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	self.Volunteers = cs.VolunteerList
	self.Number = cs.Number.Int64()
	self.BlackListLen = len(cs.BlackList) // chief-0.0.3
	self.blackList = cs.BlackList
	self.loadSigners(sl)
	go self.resetSignersLevel()
	return nil
}

func (self *TribeStatus) resetSignersLevel() {
	m := self.GetMinerAddress()
	for _, v := range self.Volunteers {
		if v == m {
			self.SignerLevel = LevelVolunteer
			return
		}
	}
	for _, s := range self.Signers {
		if s.Address == m {
			self.SignerLevel = LevelSigner
			return
		}
	}
	for _, s := range self.blackList {
		if s == m {
			self.SignerLevel = LevelSinner
			return
		}
	}
	// default none
	self.SignerLevel = LevelNone
}

func (self *TribeStatus) loadSigners(sl []*Signer) error {
	self.Signers = append(self.Signers[:0], sl...)
	return nil
}

func (self *TribeStatus) GetSigners() []*Signer {
	return self.Signers
}

func (self *TribeStatus) InTurnForCalc(signer common.Address, parent *types.Header) *big.Int {
	number := parent.Number.Int64() + 1
	signers := self.GetSigners()
	log.Debug("-- tribe.InTurnForCalc -->", "signers", signers)
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		if number%int64(len(self.Signers)) == idx.Int64() {
			return diffInTurn
		}
	}
	return diffNoTurn
}
func (self *TribeStatus) InTurnForVerify(number int64, parentHash common.Hash, signer common.Address) *big.Int {
	parentNumber := number - 1
	var signers []*Signer
	//if number > 1 & self.Number != parentNumber {
	if number > 1 {
		var err error
		signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
		if err != nil {
			log.Warn("InTurn:GetSignersFromChiefByNumber:", "parentNumber", parentNumber, "err", err)
		}
	} else {
		return diffInTurn
	}
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		log.Debug("-- InTurnForVerify -->", "parent", parentNumber, "idx", idx.Int64(), "signers", signers)
		sl := len(signers)
		if sl > 0 && number%int64(sl) == idx.Int64() {
			return diffInTurn
		}
	}
	return diffNoTurn
}

func (self *TribeStatus) genesisSigner(header *types.Header) (common.Address, error) {
	signer := common.Address{}
	copy(signer[:], header.Extra[extraVanity:])
	self.loadSigners([]*Signer{{signer, 3}})
	return signer, nil
}

func (self *TribeStatus) fetchOnSigners(address common.Address, signers []*Signer) (*big.Int, *Signer, error) {
	if signers == nil {
		signers = self.Signers
	}
	if l := len(signers); l > 0 {
		for i := 0; i < l; i++ {
			if s := signers[i]; s.Address == address {
				return big.NewInt(int64(i)), s, nil
			}
		}
	}
	return nil, nil, errors.New("not_found")
}

// called by end of WriteBlockAndState
// if miner then execute chief.update and chief.getStatus
// else execute chief.getStatus only
func (self *TribeStatus) Update(currentNumber *big.Int, hash common.Hash) {
	if currentNumber.Int64() >= CHIEF_NUMBER && atomic.LoadInt32(&self.mining) == 1 {
		// mining start
		//success := <-params.SendToMsgBox("Update")
		success := <-params.SendToMsgBoxWithNumber("Update", currentNumber)
		log.Debug("TribeStatus.Update :", "num", currentNumber.Int64(), "success", success)
	}
	self.LoadSignersFromChief(hash, currentNumber)
	return
}

func (self *TribeStatus) ValidateSigner(number int64, parentHash common.Hash, signer common.Address) bool {
	var signers []*Signer
	//if number > 1 && self.Number != parentNumber {
	if number <= 3 {
		return true
	}
	var err error
	signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
	if err != nil {
		log.Warn("TribeStatus.ValidateSigner : GetSignersFromChiefByNumber :", "err", err)
	}
	if _, _, e := self.fetchOnSigners(signer, signers); e == nil {
		return true
	}
	return false
}

// every block
// sync download or mine
// check chief tx
func (self *TribeStatus) ValidateBlock(block *types.Block, validateSigner bool) error {
	if block.Number().Int64() < 3 {
		return nil
	}

	//TODO : ****** 这个地方是临时解决方案，后续需要做很大调整
	//TODO : ****** 这个地方是临时解决方案，后续需要做很大调整
	//如果 header.parent != currentBlockNumber , 则等一下，让 blockBody 追赶 header
	//cn := chain.CurrentHeader().Number.Uint64()
	//for cn < number-1 {
	//for cn < number-1 && chain.CurrentHeader().Hash() != header.ParentHash {
	/*
		fmt.Println(cn, "--------------------------->")
		fmt.Println(cn, "=1=> currentNum:", chain.CurrentHeader().Number.Int64(), "currentHash:", chain.CurrentHeader().Hash().Hex())
		fmt.Println(cn, "=2=> parentNum:", parent.Number.Int64(), "parentHash:", parent.Hash().Hex())
		fmt.Println(cn, "=3=> headerNum", header.Number.Int64(), "header.parentHex", header.ParentHash.Hex())
	*/
	header := block.Header()
	number := block.Number().Int64()
	// add by liangc : seal call this func must skip validate signer
	if (validateSigner) {
		signer, err := ecrecover(header, self.tribe)
		if err != nil {
			return err
		}
		if !self.ValidateSigner(number, header.ParentHash, signer) {
			return errUnauthorized
		}
	}
	// check first tx , must be chief.tx , and onely one chief.tx in tx list
	if block != nil && block.Transactions().Len() == 0 {
		return ErrTribeNotAllowEmptyTxList
	}
	var total = 0
	for _, tx := range block.Transactions() {
		if tx.To() != nil && params.IsChiefAddress(*tx.To()) {
			total++
		}
	}
	if total == 0 {
		return ErrTribeMustContainChiefTx
	} else if total > 1 {
		return ErrTribeChiefCannotRepeat
	}
	log.Debug("ValidateBlockp-->", "num", block.NumberU64(), "check_signer", validateSigner)
	return nil
}

func (self *TribeStatus) String() string {
	if b, e := json.Marshal(self); e != nil {
		return "error: " + e.Error()
	} else {
		return "status: " + string(b)
	}
}
