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
	ipcpath    string
	client     *ethclient.Client
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
	ipcpath := os.Getenv("IPCPATH")
	return &TribeService{
		tribeChief: contract,
		quit:       make(chan int),
		ipcpath:    ipcpath,
	}, nil
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
		blockNumber *big.Int = nil
		blockHash            = common.HexToHash("0x")
	)
	if h, ok := mbox.Params["hash"]; ok {
		blockHash = h.(common.Hash)
		log.Debug("=>TribeService.getstatus", "blockHash", blockHash.Hex())
	} else if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
		log.Debug("-> TribeService.getstatus", "blockNumber", blockNumber.Int64())
	}
	chiefStatus, err := self.getChiefStatus(blockNumber, blockHash)
	success := params.MBoxSuccess{Success: true}
	if err != nil {
		success.Success = false
		success.Entity = err
		log.Debug("chief.mbox.rtn: getstatus <-", "success", success.Success, "err", err)
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
		log.Warn(">>=== nonce_err=", "err",perr)
	} else {
		//pnonce0, perr0 := self.client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(prv.PublicKey))
		//nonce := params.ChiefTxNonce
		//fmt.Println(">>=== pnonce 0=", pnonce0, "perr=", perr0)
		//fmt.Println(">>=== pnonce 1=", pnonce, "perr=", perr)
		log.Debug(">>=== nonce=", "nonce",pnonce)
		auth.Nonce = new(big.Int).SetUint64(pnonce)
	}
	//}
	t, e := self.tribeChief.Update(auth, self.fetchVolunteer())
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
func (self *TribeService) getChiefStatus(blockNumber *big.Int, blockHash common.Hash) (params.ChiefStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	//opts := &bind.CallOpts{Context: ctx}
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	opts.Number = blockNumber
	opts.Hash = blockHash
	chiefStatus, err := self.tribeChief.GetStatus(opts)
	if err != nil {
		return params.ChiefStatus{}, err
	}
	return params.ChiefStatus(chiefStatus), nil
}

func (self *TribeService) isVolunteer(dict map[common.Address]interface{}, add common.Address) bool {
	//TODO ****** 关于选拔的各种规则
	// Rule.1 : Do not repeat the selection
	if _, ok := dict[add]; ok {
		return false
	}
	return true
}
func (self *TribeService) fetchVolunteer() common.Address {
	peers := self.server.Peers()
	if len(peers) > 0 {
		chiefStatus, err := self.getChiefStatus(nil, common.HexToHash("0x"))
		if err != nil {
			log.Error("getChiefStatus fail", "err", err)
		}
		vl := append(chiefStatus.VolunteerList[:], chiefStatus.SignerList...)
		vmap := make(map[common.Address]interface{})
		for _, v := range vl {
			vmap[v] = struct{}{}
		}
		//fmt.Println("#########################################################################")
		for _, peer := range peers {
			pub, _ := peer.ID().Pubkey()
			add := crypto.PubkeyToAddress(*pub)
			//fmt.Println(i, "id:", peer.ID().String())
			//fmt.Println(i, "hex:", add.Hex())
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			b, e := self.client.BalanceAt(ctx, add, nil)
			if e == nil && b.Cmp(params.ChiefBaseBalance) >= 0 && self.isVolunteer(vmap, add) {
				return add
			}
			//fmt.Println(i, "balance:", e, b.Int64())
		}
		//fmt.Println("#########################################################################")
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
