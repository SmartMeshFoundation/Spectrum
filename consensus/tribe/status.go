package tribe

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/rpc"
	"github.com/SmartMeshFoundation/SMChain/core/types"
	"errors"
	"fmt"
	"github.com/SmartMeshFoundation/SMChain/params"
)

func (api *API) GetSigners(number *rpc.BlockNumber) ([]*Signer, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.status.genesisSigner(header)
	} else {
		//TODO 9999 : 在 chief 合约中获取

	}
	return api.tribe.status.Signers,nil
}

func (api *API) GetStatus(number *rpc.BlockNumber) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		fmt.Println("<><> genesis_signer_for_block1",header.Number.Int64())
		api.tribe.status.genesisSigner(header)
	} else {
		//在 chief 合约中获取
		api.tribe.status.LoadSignersFromChief()
	}
	return api.tribe.status,nil
}


func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers: make([]*Signer, 0, signerLimit),
	}
	/* 在这调用合约，会阻塞呀
	if err := ts.LoadSignersFromChief(); err!=nil {
		fmt.Println("--> tribeStatus.LoadSignersFromChief_error:",err)
	}
	*/
	return ts
}

func (self *TribeStatus) LoadSignersFromChief() error {
	rtn := make(chan params.MBoxSuccess)
	params.SendToMsgBox("GetStatus",rtn)
	r := <- rtn
	if !r.Success {
		return r.Entity.(error)
	}
	cs := r.Entity.(params.ChiefStatus)
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer,0,len(signers))
	for i,signer := range signers {
		score := scores[i]
		sl = append(sl,&Signer{signer,score.Int64()})
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

func (self *TribeStatus) genesisSigner(header *types.Header) (common.Address,error) {
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