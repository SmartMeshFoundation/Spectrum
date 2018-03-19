package tribe

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/rpc"
	"github.com/SmartMeshFoundation/SMChain/core/types"
	"errors"
	"github.com/SmartMeshFoundation/SMChain/params"
	"encoding/json"
	"sync/atomic"
	"github.com/SmartMeshFoundation/SMChain/crypto"
	"github.com/SmartMeshFoundation/SMChain/ethclient"
	"time"
	"context"
	"os"
	"github.com/SmartMeshFoundation/SMChain/log"
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

func (api *API) GetSigners(hash *common.Hash) ([]*Signer, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := common.HexToHash("0x")
		if hash!= nil {
			h = *hash
		}
		api.tribe.Status.LoadSignersFromChief(h)
	}
	return api.tribe.Status.Signers, nil
}

func (api *API) GetStatus(hash *common.Hash) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := common.HexToHash("0x")
		if hash!= nil {
			h = *hash
		}
		api.tribe.Status.LoadSignersFromChief(h)
	}
	return api.tribe.Status, nil
}

func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers:     make([]*Signer, 0),
		SignerLevel: LevelNone,
	}
	return ts
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

func (self *TribeStatus) GetSignersFromChiefByHash(hash common.Hash) ([]*Signer, error) {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash)
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
func (self *TribeStatus) LoadSignersFromChief(hash common.Hash) error {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash)
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
func (self *TribeStatus) InTurnForVerify(number int64,parentHash common.Hash, signer common.Address) *big.Int {
	parentNumber := number - 1
	var signers []*Signer
	//if number > 1 & self.Number != parentNumber {
	if number > 1 {
		var err error
		signers, err = self.GetSignersFromChiefByHash(parentHash)
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
func (self *TribeStatus) Update(currentNumber *big.Int,hash common.Hash) {
	if currentNumber.Int64() >= CHIEF_NUMBER && atomic.LoadInt32(&self.mining) == 1 {
		// mining start
		success := <-params.SendToMsgBox("Update")
		log.Debug("TribeStatus.Update :", "num", currentNumber.Int64(), "success", success)
	}
	self.LoadSignersFromChief(hash)
	return
}

func (self *TribeStatus) ValidateSigner(number int64,parentHash common.Hash, signer common.Address) bool {
	var signers []*Signer
	//if number > 1 && self.Number != parentNumber {
	if number > 1 {
		var err error
		signers, err = self.GetSignersFromChiefByHash(parentHash)
		if err != nil {
			log.Warn("TribeStatus.ValidateSigner : GetSignersFromChiefByNumber :", "err", err)
		}
	}
	if _, _, e := self.fetchOnSigners(signer, signers); e == nil {
		return true
	}
	return false
}

// every block
// sync download or mine
// check chief tx
func (self *TribeStatus) ValidatorBlock(block *types.Block) error {
	if block.Number().Int64() < 3 {
		return nil
	}
	// check first tx , must be chief.tx , and onely one chief.tx in tx list
	if block != nil && block.Transactions().Len() == 0 {
		return ErrTribeNotAllowEmptyTxList
	}
	var total = 0
	for _, tx := range block.Transactions() {
		if tx.To() != nil && common.HexToAddress(params.ChiefAddress) == *tx.To() {
			total ++
		}
	}
	if total == 0 {
		return ErrTribeMustContainChiefTx
	} else if total > 1 {
		return ErrTribeChiefCannotRepeat
	}
	return nil
}

func (self *TribeStatus) String() string {
	if b, e := json.Marshal(self); e != nil {
		return "error: " + e.Error()
	} else {
		return "status: " + string(b)
	}
}
