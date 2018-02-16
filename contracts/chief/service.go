package chief

import (
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/p2p"
	"github.com/SmartMeshFoundation/SMChain/rpc"
	"github.com/SmartMeshFoundation/SMChain/node"
	"github.com/SmartMeshFoundation/SMChain/eth"
	"github.com/SmartMeshFoundation/SMChain/les"
	"fmt"
	"github.com/SmartMeshFoundation/SMChain/internal/ethapi"
	"github.com/SmartMeshFoundation/SMChain/params"
	"time"
	"context"
	"github.com/SmartMeshFoundation/SMChain/accounts/abi/bind"
	"strings"
	"encoding/json"
)

const (
	//TODO : 换成 nodeinfo 中的 id 对应的账户
	key = `{"address":"4792ff97fbc79d659b46c56d009d74e5caee850e","crypto":{"cipher":"aes-128-ctr","ciphertext":"d6cb5b2b3e2e648a046f00cf0018541d0a43b1ede1d09d2c0247f8b6e26cafbf","cipherparams":{"iv":"b9ebe0e5364bd9bb465e4a0f971a5f43"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"17e870a3fbac58266843af921ba6c54bbc328a2ca9cabb0736f5c579bac73f42"},"mac":"e7622bb61948af35da78726bd6847827a4eb041a365bfb4daf5adf87d9de5558"},"id":"3c62b232-5c37-4e5c-a568-134ca6e3dc40","version":3}`
)

var ()
/*
type Service interface {
	Protocols() []p2p.Protocol
	APIs() []rpc.API
	Start(server *p2p.Server) error
	Stop() error
}
*/
// implements node.Service
type TribeService struct {
	tribeChief *TribeChief
	quit       chan int
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
	go self.loop()
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

func (self *TribeService) getstatus(mbox params.Mbox) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	opts := &bind.CallOpts{Context: ctx}
	success := params.MBoxSuccess{Success:true}
	chiefStatus , err := self.tribeChief.GetStatus(opts)
	success.Entity = params.ChiefStatus(chiefStatus)
	if err != nil {
		fmt.Println("-- tribeService.getStatus.err2 --> ",err)
		success.Success = false
		success.Entity = err
	} else {
		if b,e := json.Marshal(&chiefStatus) ; e==nil {
			fmt.Println("--XXXX-->",params.ChiefStatus(chiefStatus))
			fmt.Println("chief.getstatus --> ",string(b))
		}else{
			fmt.Println("-- tribeService.getStatus.err1 --> ",e)
		}
	}
	mbox.Rtn <- success
	fmt.Println("service.mbox.rtn: getstatus <- successed ")
}


func (self *TribeService) update(mbox params.Mbox) {
	//TODO change account
	key := `{"address":"4792ff97fbc79d659b46c56d009d74e5caee850e","crypto":{"cipher":"aes-128-ctr","ciphertext":"d6cb5b2b3e2e648a046f00cf0018541d0a43b1ede1d09d2c0247f8b6e26cafbf","cipherparams":{"iv":"b9ebe0e5364bd9bb465e4a0f971a5f43"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"17e870a3fbac58266843af921ba6c54bbc328a2ca9cabb0736f5c579bac73f42"},"mac":"e7622bb61948af35da78726bd6847827a4eb041a365bfb4daf5adf87d9de5558"},"id":"3c62b232-5c37-4e5c-a568-134ca6e3dc40","version":3}`
	auth, _ := bind.NewTransactor(strings.NewReader(key), "123456")
	auth.GasPrice = eth.DefaultConfig.GasPrice
	auth.GasLimit = params.GenesisGasLimit
	success := params.MBoxSuccess{Success:true}
	t,e := self.tribeChief.Update(auth,common.Address{})
	success.Entity = t.Hash().Hex()
	fmt.Println("-- update -->",e,t)
	if e != nil {
		success.Success = false
		success.Entity = e
	}
	mbox.Rtn <- success
	fmt.Println("service.mbox.rtn: update <- successed ")
}
