package statute

import (
	"context"
	"errors"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/contracts/statute/meshboxlib"
	"github.com/SmartMeshFoundation/Spectrum/eth"
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/node"
	"github.com/SmartMeshFoundation/Spectrum/p2p"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
	"math/big"
	"time"
)

type MeshboxService struct {
	meshbox  *meshboxlib.MeshBox
	ipcpath  string
	server   *p2p.Server // peers and nodekey ...
	quit     chan int
	client   *ethclient.Client
	ethereum *eth.Ethereum
}

var meshboxService *MeshboxService

func NewMeshboxService(ctx *node.ServiceContext) (node.Service, error) {
	var ethereum *eth.Ethereum
	ctx.Service(&ethereum)
	ipcpath := params.GetIPCPath()
	meshboxService = &MeshboxService{
		quit:     make(chan int),
		ipcpath:  ipcpath,
		ethereum: ethereum,
	}
	go meshboxService.loop()
	return meshboxService, nil
}

func (self *MeshboxService) Protocols() []p2p.Protocol { return nil }
func (self *MeshboxService) APIs() []rpc.API           { return nil }
func (self *MeshboxService) Start(server *p2p.Server) error {
	go func() {
		for {
			var (
				be = eth.NewContractBackend(self.ethereum.ApiBackend)
				cn = self.ethereum.BlockChain().CurrentBlock().Number()
			)
			mn, maddr := params.MeshboxInfo(cn)
			if maddr != common.HexToAddress("") {
				contract, err := meshboxlib.NewMeshBox(maddr, be)
				if err != nil {
					panic(err)
				}
				meshboxService.meshbox = contract
				close(params.InitMeshboxService)
				log.Info("<<MeshboxService.Start>> success ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else if cn.Cmp(mn) >= 0 {
				log.Info("<<MeshboxService.Start>> cancel ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			}
			log.Info("<<MeshboxService.Start>> waiting... ", "cn", cn.Int64(), "tn", mn.Int64())
			<-time.After(14 * time.Second)
		}
	}()
	self.server = server
	return nil
}

func (self *MeshboxService) Stop() error {
	close(self.quit)
	return nil
}

// ===============================================================================
// biz functions
// ===============================================================================

func GetMeshboxService() (*MeshboxService, error) {
	log.Debug("<<meshbox.service.GetMeshboxService>>")
	select {
	case <-params.InitMeshboxService:
		return meshboxService, nil
	default:
		return nil, errors.New("wait init")
	}
}

func (self *MeshboxService) ExistAddress(addr common.Address) (*big.Int, error) {
	log.Debug("<<meshbox.service.ExistAddress>>")
	select {
	case <-params.InitMeshboxService:
		var ctx = context.Background()
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		i, err := self.meshbox.ExistAddress(opts, addr)
		log.Debug("<<MeshboxService.ExistAddress>>", "r", i, "err", err)
		return i, err
	default:
		return nil, errors.New("wait init")
	}
}

func (self *MeshboxService) existAddress(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var addr common.Address
	addr = mbox.Params["addr"].(common.Address)
	log.Debug("mbox.params", "addr", addr.Hex())
	i, err := self.ExistAddress(addr)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = i.Int64()
	}
	mbox.Rtn <- success
}

func (self *MeshboxService) loop() {
	for {
		select {
		case <-self.quit:
			break
		case mbox := <-params.MeshboxService:
			switch mbox.Method {
			case "existAddress":
				self.existAddress(mbox)
			}
		}
	}
}
