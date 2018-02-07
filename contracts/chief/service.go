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
		fmt.Println("+++++++++++++++++ 1")
		select {
		case <-self.quit:
			break
		case mbox := <-params.MboxChan:
			fmt.Println("+++++++++++++++++ 2", mbox)
			switch mbox.Method {
			case "GetStatus":
				fmt.Println("+++++++++++++++++ 3", mbox)
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
	fmt.Println("++++++++ 111")
	Al, Sl, Bl, err := self.tribeChief.GetStatus(opts)
	fmt.Println("++++++++ 222",err)
	fmt.Println("-->", err)
	if err == nil {
		for i, a := range Al {
			fmt.Println("> ", i, a.Hex(), Sl[i], Bl[i])
		}
	}
	mbox.Rtn <- "getstatus_successed."
	fmt.Println("service.mbox.rtn <- successed ")
}

func (self *TribeService) update(mbox params.Mbox) {
	self.tribeChief.Update(nil)
}
