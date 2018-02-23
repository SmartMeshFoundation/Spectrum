package chief

import (
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/p2p"
	"github.com/SmartMeshFoundation/SMChain/rpc"
	"github.com/SmartMeshFoundation/SMChain/node"
	"github.com/SmartMeshFoundation/SMChain/eth"
	"github.com/SmartMeshFoundation/SMChain/les"
	"github.com/SmartMeshFoundation/SMChain/internal/ethapi"
	"github.com/SmartMeshFoundation/SMChain/params"
	"time"
	"context"
	"github.com/SmartMeshFoundation/SMChain/accounts/abi/bind"
	"github.com/SmartMeshFoundation/SMChain/log"
	"math/big"
	"fmt"
)

/*
type Service interface {
	Protocols() []p2p.Protocol
	APIs() []rpc.API
	Start(server *p2p.Server) error
	Stop() error
}
*/
//implements node.Service
type TribeService struct {
	tribeChief *TribeChief
	quit       chan int
	server     *p2p.Server // peers and nodekey ...
}

func NewTribeService(ctx *node.ServiceContext) (node.Service, error) {
	var apiBackend ethapi.Backend
	var ethereum *eth.Ethereum
	if err := ctx.Service(&ethereum); err == nil {
		apiBackend = ethereum.ApiBackend
	} else {
		var ethereum *les.LightEthereum
		if err := ctx.Service(&ethereum); err == nil {
			apiBackend = ethereum.ApiBackend
		} else {
			return nil, err
		}
	}
	contract, err := NewTribeChief(common.HexToAddress(params.ChiefAddress), eth.NewContractBackend(apiBackend))
	if err != nil {
		return nil, err
	}
	return &TribeService{
		tribeChief: contract,
		quit:       make(chan int),
	}, nil
}

func (self *TribeService) Protocols() []p2p.Protocol { return nil }
func (self *TribeService) APIs() []rpc.API           { return nil }

func (self *TribeService) Start(server *p2p.Server) error {
	self.server = server
	go self.loop()
	params.InitTribeStatus <- struct{}{}
	return nil
}
func (self *TribeService) loop() {
	for {
		select {
		case <-self.quit:
			break
		case mbox := <-params.MboxChan:
			switch mbox.Method {
			case "GetStatus":
				self.getstatus(mbox)
			case "GetNodeKey":
				self.getnodekey(mbox)
			case "Update":
				self.update(mbox)
			}
		}
	}
}

func (self *TribeService) Stop() error {
	self.quit <- 1
	return nil
}

func (self *TribeService) getnodekey(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	success.Entity = self.server.PrivateKey
	mbox.Rtn <- success
}

func (self *TribeService) getstatus(mbox params.Mbox) {
	var blockNumber *big.Int = nil
	if n,ok := mbox.Params["number"];ok {
		fmt.Println("--> TribeService.getstatus : blockNumber =",n)
		blockNumber = n.(*big.Int)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	//opts := &bind.CallOpts{Context: ctx}
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	opts.Number = blockNumber
	success := params.MBoxSuccess{Success: true}
	chiefStatus, err := self.tribeChief.GetStatus(opts)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = params.ChiefStatus(chiefStatus)
	}
	mbox.Rtn <- success
	log.Debug("chief.mbox.rtn: getstatus <-", "success", success)
}

func (self *TribeService) update(mbox params.Mbox) {
	prv := self.server.PrivateKey
	auth := bind.NewKeyedTransactor(prv)
	auth.GasPrice = eth.DefaultConfig.GasPrice
	auth.GasLimit = params.GenesisGasLimit
	success := params.MBoxSuccess{Success: true}

	if params.ChiefTxNonce > 0 {
		nonce := params.ChiefTxNonce+1
		auth.Nonce = new(big.Int).SetUint64(nonce)
		fmt.Println("YYYYYYYYYYYYYY>> ",auth.Nonce.Uint64())
	}
	t, e := self.tribeChief.Update(auth, common.Address{})
	if e != nil {
		success.Success = false
		success.Entity = e
		//log.Error("TribeService.update",e)
	} else {
		success.Entity = t.Hash().Hex()
	}
	mbox.Rtn <- success
	log.Debug("chief.mbox.rtn: update <-", "success", success)
}
