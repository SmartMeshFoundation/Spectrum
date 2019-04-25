package statute

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/SmartMeshFoundation/Spectrum/accounts"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/contracts"
	"github.com/SmartMeshFoundation/Spectrum/contracts/statute/anmaplib"
	"github.com/SmartMeshFoundation/Spectrum/contracts/statute/meshboxlib"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/eth"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/node"
	"github.com/SmartMeshFoundation/Spectrum/p2p"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
	"math/big"
	"time"
)

type AnmapService interface {
	BindInfo(addr common.Address, blockNumber *big.Int, blockHash *common.Hash) (from common.Address, nodeids []common.Address, err error)
	Bind(from, nodeAddr common.Address, sigHex string) (common.Hash, error)
	Unbind(from, nodeAddr common.Address, sigHex string) (common.Hash, error)
}

type MeshboxService interface {
	ExistAddress(addr common.Address) (*big.Int, error)
}

type StatuteService struct {
	accman   *accounts.Manager
	anmap    *anmaplib.Anmap
	meshbox  *meshboxlib.MeshBox
	ipcpath  string
	server   *p2p.Server // peers and nodekey ...
	quit     chan int
	ethereum *eth.Ethereum
}

var statuteService *StatuteService

func NewStatuteService(ctx *node.ServiceContext) (node.Service, error) {
	var ethereum *eth.Ethereum
	ctx.Service(&ethereum)
	ipcpath := params.GetIPCPath()
	statuteService = &StatuteService{
		accman:   ctx.AccountManager,
		quit:     make(chan int),
		ipcpath:  ipcpath,
		ethereum: ethereum,
	}
	go statuteService.loop()
	return statuteService, nil
}

func (self *StatuteService) Protocols() []p2p.Protocol { return nil }
func (self *StatuteService) APIs() []rpc.API           { return nil }
func (self *StatuteService) Start(server *p2p.Server) error {
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
				statuteService.meshbox = contract
				close(params.InitMeshbox)
				log.Info("<<Meshbox.Start>> success ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else if cn.Cmp(mn) >= 0 {
				log.Info("<<Meshbox.Start>> cancel ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			}
			//log.Info("<<Meshbox.Start>> waiting... ", "cn", cn.Int64(), "tn", mn.Int64())
			<-time.After(14 * time.Second)
		}
	}()
	go func() {
		for {
			var (
				be = eth.NewContractBackend(self.ethereum.ApiBackend)
				cn = self.ethereum.BlockChain().CurrentBlock().Number()
			)
			mn, maddr := params.AnmapInfo(cn)
			if maddr != common.HexToAddress("") {
				contract, err := anmaplib.NewAnmap(maddr, be)
				if err != nil {
					panic(err)
				}
				statuteService.anmap = contract
				close(params.InitAnmap)
				log.Info("<<Anmap.Start>> success ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else if cn.Cmp(mn) >= 0 {
				log.Info("<<Anmap.Start>> cancel ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			}
			//log.Info("<<Anmap.Start>> waiting... ", "cn", cn.Int64(), "tn", mn.Int64())
			<-time.After(14 * time.Second)
		}
	}()
	self.server = server
	return nil
}

func (self *StatuteService) Stop() error {
	close(self.quit)
	return nil
}

// ===============================================================================
// biz functions
// ===============================================================================

func GetAnmapService() (AnmapService, error) {
	log.Debug("<<GetAnmapService>>")
	select {
	case <-params.InitMeshbox:
		return statuteService, nil
	default:
		return nil, errors.New("anmap wait init")
	}
}

func GetMeshboxService() (MeshboxService, error) {
	log.Debug("<<GetMeshboxService>>")
	select {
	case <-params.InitMeshbox:
		return statuteService, nil
	default:
		return nil, errors.New("meshbox wait init")
	}
}

func sigSplit(sigHex string) (R, S [32]byte, V uint8) {
	bR, _ := hex.DecodeString(sigHex[:64])
	bS, _ := hex.DecodeString(sigHex[64:128])
	copy(R[:], bR)
	copy(S[:], bS)
	V = 27
	switch sigHex[128:] {
	case "01":
		V = 28
	}
	return
}

func (self *StatuteService) BindInfo(addr common.Address, blockNumber *big.Int, blockHash *common.Hash) (from common.Address, nodeids []common.Address, err error) {
	chash := self.ethereum.BlockChain().CurrentBlock().Hash()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	if blockHash != nil {
		opts.Hash = blockHash
	}

	if blockNumber != nil {
		opts.Number = blockNumber
	}
	if blockNumber == nil && blockHash == nil {
		opts.Hash = &chash
	}
	vo, err := self.anmap.BindInfo(opts, addr)

	// anmap.sol bindInfo func has a problum , an/na return diff nodeids so query again
	if err == nil && len(vo.Nids) == 1 && vo.From != addr {
		vo, err = self.anmap.BindInfo(opts, vo.From)
	}

	from = vo.From
	nodeids = vo.Nids

	return
}

func (self *StatuteService) Bind(from, nodeAddr common.Address, sigHex string) (common.Hash, error) {
	a := accounts.Account{Address: from}
	w, _ := self.accman.Find(a)
	opts := &bind.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return w.SignTx(a, tx, params.ChainID())
		},
	}
	//client   *ethclient.Client
	client, err := contracts.GetEthclientInstance()
	if err != nil {
		return common.Hash{}, err
	}
	//if params.ChiefTxNonce > 0 {
	pnonce, perr := client.PendingNonceAt(context.Background(), from)
	if perr != nil {
		log.Debug("<<StatuteService_Bind>> === nonce_err", "err", perr)
	} else {
		log.Debug("<<StatuteService_Bind>> === nonce", "nonce", pnonce)
		opts.Nonce = new(big.Int).SetUint64(pnonce)
	}

	r, s, v := sigSplit(sigHex)
	tx, err := self.anmap.Bind(opts, nodeAddr, v, r, s)
	log.Info("<<StatuteService.Bind>>", "err", err, "tx", tx)
	if err != nil {
		return common.HexToHash("0x"), err
	}
	return tx.Hash(), nil
}

func (self *StatuteService) Unbind(from, nodeAddr common.Address, sigHex string) (common.Hash, error) {
	a := accounts.Account{Address: from}
	w, _ := self.accman.Find(a)
	opts := &bind.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return w.SignTx(a, tx, params.ChainID())
		},
	}
	r, s, v := sigSplit(sigHex)
	tx, err := self.anmap.UnbindBySig(opts, nodeAddr, v, r, s)
	log.Info("<<StatuteService.Unbind>>", "err", err, "tx", tx)
	if err != nil {
		return common.HexToHash("0x"), err
	}
	return tx.Hash(), nil
}

func (self *StatuteService) getBalance(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	addr := mbox.Params["addr"].(common.Address)
	sdb, _ := self.ethereum.BlockChain().State()
	success.Entity = map[string]interface{}{"addr": addr, "balance": sdb.GetBalance(addr)}
	mbox.Rtn <- success
}

/*
args:
	addr
	hash : blockHash
*/
func (self *StatuteService) bindInfo(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var (
		addr        common.Address
		blockHash   *common.Hash
		blockNumber *big.Int
	)
	addr = mbox.Params["addr"].(common.Address)
	// hash and number can not nil
	if h, ok := mbox.Params["hash"]; ok {
		bh := h.(common.Hash)
		blockHash = &bh
		_header := self.ethereum.BlockChain().GetHeaderByHash(*blockHash)
		log.Debug("<<StatuteService_bindInfo>>", "hash", blockHash.Hex(), "header", _header)
		if _header == nil {
			blockHash = nil
		}
	}
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
	}
	f, n, err := self.BindInfo(addr, blockNumber, blockHash)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = map[string]interface{}{"from": f, "nodeids": n}
	}
	mbox.Rtn <- success
}

/*
args:
	from
	nodeid
	sigHex
*/
func (self *StatuteService) bind(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var (
		from, nodeid common.Address
		sigHex       string
	)
	from = mbox.Params["from"].(common.Address)
	nodeid = mbox.Params["nodeid"].(common.Address)
	sigHex = mbox.Params["sigHex"].(string)
	log.Info("mbox.params", "from", from.Hex(), "nodeid", nodeid, "sigHex", sigHex)
	txHash, err := self.Bind(from, nodeid, sigHex)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = txHash.Hex()
	}
	mbox.Rtn <- success
}

/*
args:
	from
	nodeid
	sigHex
*/
func (self *StatuteService) unbind(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var (
		from, nodeid common.Address
		sigHex       string
	)
	from = mbox.Params["from"].(common.Address)
	nodeid = mbox.Params["nodeid"].(common.Address)
	sigHex = mbox.Params["sigHex"].(string)
	log.Info("mbox.params", "from", from.Hex(), "nodeid", nodeid, "sigHex", sigHex)
	txHash, err := self.Unbind(from, nodeid, sigHex)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = txHash.Hex()
	}
	mbox.Rtn <- success
}

func (self *StatuteService) ExistAddress(addr common.Address) (*big.Int, error) {
	log.Debug("<<StatuteService.ExistAddress>>")
	select {
	case <-params.InitMeshbox:
		var ctx = context.Background()
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		i, err := self.meshbox.ExistAddress(opts, addr)
		log.Debug("<<StatuteService.ExistAddress>>", "r", i, "err", err)
		return i, err
	default:
		return nil, errors.New("wait init")
	}
}

func (self *StatuteService) existAddress(mbox params.Mbox) {
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

func (self *StatuteService) loop() {
	for {
		select {
		case <-self.quit:
			break
		case mbox := <-params.StatuteService:
			switch mbox.Method {
			case "getBalance":
				self.getBalance(mbox)
			case "bindInfo":
				self.bindInfo(mbox)
			case "bind":
				self.bind(mbox)
			case "unbind":
				self.unbind(mbox)
			case "existAddress":
				self.existAddress(mbox)
			}
		}
	}
}
