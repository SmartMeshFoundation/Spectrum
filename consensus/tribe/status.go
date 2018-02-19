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
	//TODO level from status
	level := "None"
	return &TribeMiner{add, b, level}, nil
}

func (api *API) GetSigners(number *rpc.BlockNumber) ([]*Signer, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		api.tribe.Status.LoadSignersFromChief()
	}
	return api.tribe.Status.Signers, nil
}

func (api *API) GetStatus(number *rpc.BlockNumber) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		api.tribe.Status.LoadSignersFromChief()
	}
	return api.tribe.Status, nil
}

func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers: make([]*Signer, 0, signerLimit),
	}
	return ts
}

func (self *TribeStatus) GetMinerAddress() common.Address {
	pub := self.nodeKey.PublicKey
	add := crypto.PubkeyToAddress(pub)
	return add
}

// 在 加载完所有 node.service 后，需要主动调用一次
func (self *TribeStatus) LoadSignersFromChief() error {
	rtn := params.SendToMsgBox("GetStatus")
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
	self.LoadSigners(sl)
	return nil
}
func (self *TribeStatus) LoadSigners(sl []*Signer) error {
	if len(sl) > signerLimit {
		return errors.New("size_not_allow")
	}
	self.Signers = append(self.Signers[:0], sl...)
	return nil
}

func (self *TribeStatus) GetSigners() []*Signer {
	return self.Signers
}

func (self *TribeStatus) InTurn(number int64, signer common.Address) *big.Int {
	if idx, _, err := self.fetchOnSigners(signer); err == nil {
		if number%int64(len(self.Signers)) == idx.Int64() {
			return diffInTurn
		}
	}
	return diffNoTurn
}

// 签名人是否在授权列表中
func (self *TribeStatus) ValidateSigner(signer common.Address) bool {
	if _, _, e := self.fetchOnSigners(signer); e == nil {
		return true
	}
	return false
}

func (self *TribeStatus) genesisSigner(header *types.Header) (common.Address, error) {
	signer := common.Address{}
	copy(signer[:], header.Extra[extraVanity:])
	self.LoadSigners([]*Signer{{signer, 3}})
	return signer, nil
}

func (self *TribeStatus) fetchOnSigners(address common.Address) (*big.Int, *Signer, error) {
	if l := len(self.Signers); l > 0 {
		for i := 0; i < l; i++ {
			if s := self.Signers[i]; s.Address == address {
				return big.NewInt(int64(i)), s, nil
			}
		}
	}
	return nil, nil, errors.New("not_found")
}

// called by end of WriteBlockAndState
// if miner then execute chief.update and chief.getStatus
// else execute chief.getStatus only
func (self *TribeStatus) Update(currentNumber *big.Int) {
	if currentNumber.Int64() > 1 && atomic.LoadInt32(&self.mining) == 1 {
		// mining start
		success := <-params.SendToMsgBox("Update")
		fmt.Println("<><><><> TribeStatus.Update :", success)
	}
	self.LoadSignersFromChief()
	return
}

func (self *TribeStatus) String() string {
	if b, e := json.Marshal(self); e != nil {
		return "error: " + e.Error()
	} else {
		return "status: " + string(b)
	}
}
