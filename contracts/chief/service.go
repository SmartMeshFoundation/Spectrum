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
	"github.com/SmartMeshFoundation/SMChain/crypto"
	"github.com/SmartMeshFoundation/SMChain/ethclient"
	"os"
	"github.com/SmartMeshFoundation/SMChain/contracts/chief/lib"
	"github.com/SmartMeshFoundation/SMChain/core/types"
	"errors"
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
	tribeChief_0_0_2 *chieflib.TribeChief
	tribeChief_0_0_3 *chieflib.TribeChief_0_0_3
	quit             chan int
	server           *p2p.Server // peers and nodekey ...
	ipcpath          string
	client           *ethclient.Client
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
	ipcpath := os.Getenv("IPCPATH")
	ts := &TribeService{
		quit:    make(chan int),
		ipcpath: ipcpath,
	}
	if v0_0_2 := params.GetChiefInfoByVsn("0.0.2"); v0_0_2 != nil {
		contract_0_0_2, err := chieflib.NewTribeChief(v0_0_2.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_2 = contract_0_0_2
	}
	if v0_0_3 := params.GetChiefInfoByVsn("0.0.3"); v0_0_3 != nil {
		contract_0_0_3, err := chieflib.NewTribeChief_0_0_3(v0_0_3.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_3 = contract_0_0_3
	}
	return ts, nil
}

func (self *TribeService) Protocols() []p2p.Protocol { return nil }
func (self *TribeService) APIs() []rpc.API           { return nil }

func (self *TribeService) Start(server *p2p.Server) error {
	self.server = server
	go self.loop()
	close(params.InitTribeStatus)
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
	var (
		blockNumber *big.Int     = nil
		blockHash   *common.Hash = nil
	)
	// hash and number can not nil
	if h, ok := mbox.Params["hash"]; ok {
		bh := h.(common.Hash)
		blockHash = &bh
	}
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
	}
	log.Debug("=>TribeService.getstatus", "blockNumber", blockNumber, "blockHash", blockHash.Hex())

	success := params.MBoxSuccess{Success: true}
	chiefStatus, err := self.getChiefStatus(blockNumber, blockHash)
	if err != nil {
		success.Success = false
		success.Entity = err
		log.Error("chief.mbox.rtn: getstatus <-", "success", success.Success, "err", err)
	} else {
		entity := chiefStatus
		success.Entity = entity
		log.Debug("chief.mbox.rtn: getstatus <-", "success", success.Success, "entity", entity)
	}
	mbox.Rtn <- success
}

func (self *TribeService) update(mbox params.Mbox) {
	prv := self.server.PrivateKey
	auth := bind.NewKeyedTransactor(prv)
	auth.GasPrice = eth.DefaultConfig.GasPrice
	auth.GasLimit = params.GenesisGasLimit
	success := params.MBoxSuccess{Success: false}

	if err := self.initEthclient(); err != nil {
		success.Entity = err
		mbox.Rtn <- success
		return
	}
	//if params.ChiefTxNonce > 0 {
	pnonce, perr := self.client.NonceAt(context.Background(), crypto.PubkeyToAddress(prv.PublicKey), nil)
	if perr != nil {
		log.Warn(">>=== nonce_err=", "err", perr)
	} else {
		log.Debug(">>=== nonce=", "nonce", pnonce)
		auth.Nonce = new(big.Int).SetUint64(pnonce)
	}
	//}
	var (
		t           *types.Transaction
		e           error
		blockNumber *big.Int
	)
	// not nil
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
		log.Debug("-> TribeService.update", "blockNumber", blockNumber.Int64())
	} else {
		success.Entity = errors.New("TribeService.update : blockNumber not nil")
		mbox.Rtn <- success
		return
	}
	if chiefInfo := params.GetChiefInfo(blockNumber); chiefInfo != nil {
		switch chiefInfo.Version {
		case "0.0.2":
			t, e = self.tribeChief_0_0_2.Update(auth, self.fetchVolunteer(blockNumber))
		case "0.0.3":
			t, e = self.tribeChief_0_0_3.Update(auth, self.fetchVolunteer(blockNumber))
		}
	}

	if e != nil {
		success.Entity = e
	} else {
		success.Success = true
		success.Entity = t.Hash().Hex()
	}
	mbox.Rtn <- success
	log.Debug("chief.mbox.rtn: update <-", "success", success)
}

// --------------------------------------------------------------------------------------------------
// inner private
// --------------------------------------------------------------------------------------------------
func (self *TribeService) getChiefStatus(blockNumber *big.Int, blockHash *common.Hash) (params.ChiefStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	//opts := &bind.CallOpts{Context: ctx}
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	opts.Hash = blockHash
	if chiefInfo := params.GetChiefInfo(blockNumber); chiefInfo != nil {
		switch chiefInfo.Version {
		case "0.0.2":
			chiefStatus, err := self.tribeChief_0_0_2.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			return params.ChiefStatus{
				VolunteerList: chiefStatus.VolunteerList,
				SignerList:    chiefStatus.SignerList,
				ScoreList:     chiefStatus.ScoreList,
				NumberList:    chiefStatus.NumberList,
				BlackList:     nil,
				Number:        chiefStatus.Number,
			}, nil
		case "0.0.3":
			chiefStatus, err := self.tribeChief_0_0_3.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			return params.ChiefStatus(chiefStatus), err
		}
	}
	return params.ChiefStatus{}, errors.New("status_not_found")
}

func (self *TribeService) isVolunteer(dict map[common.Address]interface{}, add common.Address) bool {
	//TODO ****** 关于选拔的各种规则
	// Rule.1 : Do not repeat the selection
	if _, ok := dict[add]; ok {
		return false
	}
	return true
}
func (self *TribeService) fetchVolunteer(blockNumber *big.Int) common.Address {
	peers := self.server.Peers()
	if len(peers) > 0 {
		chiefStatus, err := self.getChiefStatus(blockNumber, nil)
		if err != nil {
			log.Error("getChiefStatus fail", "err", err)
		}
		// exclude signers and volunteers
		vl := append(chiefStatus.VolunteerList[:], chiefStatus.SignerList...)
		if chiefStatus.BlackList != nil {
			// exclude blacklist address
			vl = append(vl[:], chiefStatus.BlackList...)
		}
		vmap := make(map[common.Address]interface{})
		for _, v := range vl {
			vmap[v] = struct{}{}
		}
		for _, peer := range peers {
			pub, _ := peer.ID().Pubkey()
			add := crypto.PubkeyToAddress(*pub)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			b, e := self.client.BalanceAt(ctx, add, nil)
			if e == nil && b.Cmp(params.ChiefBaseBalance) >= 0 && self.isVolunteer(vmap, add) {
				return add
			}
		}
	}
	return common.Address{}
}

func (self *TribeService) initEthclient() error {
	if self.client == nil {
		ethclient, err := ethclient.Dial(self.ipcpath)
		if err != nil {
			log.Error("ipc error at tribeservice.update", "err", err)
			return err
		}
		self.client = ethclient
	}
	return nil
}
