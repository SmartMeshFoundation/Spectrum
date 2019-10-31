package statute

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	chieflib "github.com/SmartMeshFoundation/Spectrum/contracts/chief/lib"

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
	accman        *accounts.Manager
	anmap_0_0_1   *anmaplib.Anmap
	meshbox_0_0_1 *meshboxlib.MeshBox
	meshbox_0_0_2 *meshboxlib.MeshBox_0_0_2
	poc_1         *chieflib.POC_1_0_0
	ipcpath       string
	server        *p2p.Server // peers and nodekey ...
	quit          chan int
	ethereum      *eth.Ethereum
}

var statuteService *StatuteService

func NewStatuteService(ctx *node.ServiceContext) (node.Service, error) {
	var ethereum *eth.Ethereum
	err := ctx.Service(&ethereum)
	if err != nil {
		log.Error("NewStatuteService", "err", err)
	}
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

func (self *StatuteService) startMeshbox(vsn string, backend *eth.ContractBackend) {
	defer func() {
		if r := recover(); r != nil {
			log.Warn("ignore_this_err", r)
		}
	}()
	var period = params.TribePeriod()
	for {
		var cn = self.ethereum.BlockChain().CurrentBlock().Number()
		if params.IsReadyMeshbox(cn) {
			mn, maddr := params.MeshboxInfo(cn, vsn)
			if maddr != common.HexToAddress("") {
				switch vsn {
				case "0.0.1":
					contract, err := meshboxlib.NewMeshBox(maddr, backend)
					if err != nil {
						panic(err)
					}
					statuteService.meshbox_0_0_1 = contract
				case "0.0.2":
					contract, err := meshboxlib.NewMeshBox_0_0_2(maddr, backend)
					if err != nil {
						panic(err)
					}
					statuteService.meshbox_0_0_2 = contract
				}
				close(params.InitMeshbox)
				log.Info("<<Meshbox.Start>> success ", "period", period, "vsn", vsn, "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else {
				//} else if cn.Cmp(mn) >= 0 {
				log.Info("<<Meshbox.Start>> cancel ", "period", period, "vsn", vsn, "cn", cn, "tn", mn)
				return
			}
		}
		<-time.After(time.Duration(period) * time.Second)
	}
}

func (self *StatuteService) startAnmap(vsn string, backend *eth.ContractBackend) {
	defer func() {
		if r := recover(); r != nil {
			log.Warn("ignore_this_err", r)
		}
	}()
	var period = params.TribePeriod()
	for {
		var cn = self.ethereum.BlockChain().CurrentBlock().Number()
		if params.IsReadyAnmap(cn) {
			mn, maddr := params.AnmapInfo(cn, vsn)
			if maddr != common.HexToAddress("") {
				switch vsn {
				case "0.0.1":
					contract, err := anmaplib.NewAnmap(maddr, backend)
					if err != nil {
						panic(err)
					}
					statuteService.anmap_0_0_1 = contract
				}
				close(params.InitAnmap)
				log.Info("<<Anmap.Start>> success ", "period", period, "vsn", vsn, "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else if cn.Cmp(mn) >= 0 {
				log.Info("<<Anmap.Start>> cancel ", "period", period, "vsn", vsn, "cn", cn.Int64(), "tn", mn.Int64())
				return
			}
		}
		<-time.After(time.Duration(period) * time.Second)
	}
}

func (self *StatuteService) Protocols() []p2p.Protocol { return nil }
func (self *StatuteService) APIs() []rpc.API           { return nil }
func (self *StatuteService) Start(server *p2p.Server) error {
	var be = eth.NewContractBackend(self.ethereum.ApiBackend)
	go self.startMeshbox("0.0.1", be)
	go self.startMeshbox("0.0.2", be)
	go self.startAnmap("0.0.1", be)
	if true {
		//poc contract service
		var err error
		self.poc_1, err = chieflib.NewPOC_1_0_0(params.POCInfo(), be)
		if err != nil {
			panic(err)
		}
	}

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
	bR, err := hex.DecodeString(sigHex[:64])
	bS, err := hex.DecodeString(sigHex[64:128])
	if err != nil {
		log.Error("sigSplit ", "sigHex", sigHex, "err", err)
	}
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
	vo, err := self.anmap_0_0_1.BindInfo(opts, addr)

	// anmap.sol bindInfo func has a problum , an/na return diff nodeids so query again
	if err == nil && len(vo.Nids) == 1 && vo.From != addr {
		vo, err = self.anmap_0_0_1.BindInfo(opts, vo.From)
	}

	from = vo.From
	nodeids = vo.Nids

	return
}

func (self *StatuteService) Bind(from, nodeAddr common.Address, sigHex string) (common.Hash, error) {
	a := accounts.Account{Address: from}
	w, err := self.accman.Find(a)
	if err != nil {
		return common.Hash{}, err
	}
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
	tx, err := self.anmap_0_0_1.Bind(opts, nodeAddr, v, r, s)
	log.Info("<<StatuteService.Bind>>", "err", err, "tx", tx)
	if err != nil {
		return common.HexToHash("0x"), err
	}
	return tx.Hash(), nil
}

func (self *StatuteService) Unbind(from, nodeAddr common.Address, sigHex string) (common.Hash, error) {
	a := accounts.Account{Address: from}
	w, err := self.accman.Find(a)
	if err != nil {
		log.Error("Unbindb find", "err", err)
	}
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
	tx, err := self.anmap_0_0_1.UnbindBySig(opts, nodeAddr, v, r, s)
	log.Info("<<StatuteService.Unbind>>", "err", err, "tx", tx)
	if err != nil {
		return common.HexToHash("0x"), err
	}
	return tx.Hash(), nil
}

func (self *StatuteService) getBalance(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	addr := mbox.Params["addr"].(common.Address)
	sdb, err := self.ethereum.BlockChain().State()
	if err != nil {
		log.Error("BlockChain().State()", "err", err)
	}
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
		_block := self.ethereum.BlockChain().GetBlockByHash(*blockHash)
		log.Debug("<<StatuteService_bindInfo>>", "hash", blockHash.Hex())
		if _block == nil {
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
from common.Address, sigHex string
*/
func (self *StatuteService) pocDeposit(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var (
		from   common.Address
		sigHex string
		err    error
		tx     *types.Transaction
	)
	defer func() {
		if err != nil {
			success.Success = false
			success.Entity = err
		} else {
			success.Entity = tx.Hash().String()
		}
		mbox.Rtn <- success
	}()
	from = mbox.Params["from"].(common.Address)
	sigHex = mbox.Params["sigHex"].(string)
	log.Info("mbox.params", "from", from.Hex(), "sigHex", sigHex, "pocaddr", params.POCInfo())

	a := accounts.Account{Address: from}
	w, err := self.accman.Find(a)
	if err != nil {
		return
	}
	//client   *ethclient.Client
	client, err := contracts.GetEthclientInstance()
	if err != nil {
		return
	}
	poc, err := chieflib.NewPOC_1_0_0(params.POCInfo(), client)
	if err != nil {
		return
	}
	//质押最小金额
	min, err := poc.MinDepositAmount(nil)
	if err != nil {
		log.Error(fmt.Sprintf("query min deposit err %s,poc addr=%s", err, params.POCInfo().String()))
		return
	}
	sdb, err := self.ethereum.BlockChain().State()
	if err != nil {
		return
	}
	balance := sdb.GetBalance(from)
	if balance.Cmp(min) <= 0 {
		err = fmt.Errorf("addr %s doesn't have enough balance, need=%s, have=%s", from.String(), min, balance)
	}
	opts := &bind.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return w.SignTx(a, tx, params.ChainID())
		},
		Value: min,
	}
	//if params.ChiefTxNonce > 0 {
	pnonce, perr := client.PendingNonceAt(context.Background(), from)
	if perr != nil {
		log.Debug("<<StatuteService_PocDeposit>> === nonce_err", "err", perr)
	} else {
		log.Debug("<<StatuteService_PocDeposit>> === nonce", "nonce", pnonce)
		opts.Nonce = new(big.Int).SetUint64(pnonce)
	}

	r, s, v := sigSplit(sigHex)
	tx, err = poc.Deposit(opts, r, s, v)
	log.Info("<<StatuteService.PocDeposit>>", "err", err, "tx", tx)
	return
}

func (self *StatuteService) pocStartAndStopAndWithdrawAndWithdrawSurplus(mbox params.Mbox, method string) {
	success := params.MBoxSuccess{Success: true}
	var (
		from   common.Address
		nodeID common.Address
		err    error
		tx     *types.Transaction
	)
	defer func() {
		if err != nil {
			success.Success = false
			success.Entity = err
		} else {
			success.Entity = tx.Hash().String()
		}
		mbox.Rtn <- success
	}()
	from = mbox.Params["from"].(common.Address)
	nodeID = mbox.Params["nodeid"].(common.Address)
	log.Info("mbox.params", "from", from.Hex(), "nodeid", nodeID, "pocaddr", params.POCInfo())

	a := accounts.Account{Address: from}
	w, err := self.accman.Find(a)
	if err != nil {
		return
	}
	opts := &bind.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return w.SignTx(a, tx, params.ChainID())
		},
	}
	//client   *ethclient.Client
	client, err := contracts.GetEthclientInstance()
	if err != nil {
		return
	}
	//if params.ChiefTxNonce > 0 {
	pnonce, perr := client.PendingNonceAt(context.Background(), from)
	if perr != nil {
		log.Debug("<<StatuteService_pocStartAndStopAndWithdrawAndWithdrawSurplus>> === nonce_err", "err", perr)
	} else {
		log.Debug("<<StatuteService_pocStartAndStopAndWithdrawAndWithdrawSurplus>> === nonce", "nonce", pnonce)
		opts.Nonce = new(big.Int).SetUint64(pnonce)
	}
	switch method {
	case params.POC_METHOD_START:
		tx, err = self.poc_1.Start(opts, nodeID)
	case params.POC_METHOD_STOP:
		tx, err = self.poc_1.Stop(opts, nodeID)
	case params.POC_METHOD_WITHDRAW:
		tx, err = self.poc_1.Withdraw(opts, nodeID)
	case params.POC_METHOD_WITHDRAW_SURPLUS:
		tx, err = self.poc_1.WithdrawSurplus(opts, nodeID)
	default:
		panic(method)
	}
	log.Info("<<StatuteService>>", "method", method, "err", err, "tx", tx)
	return
}

func (self *StatuteService) pocGetAll(mbox params.Mbox) {
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
	log.Debug("=>TribeService.pocGetAll", "blockNumber", blockNumber, "blockHash", blockHash.Hex())
	success := params.MBoxSuccess{Success: true}
	var (
		err error
		ps  *params.PocStatus
	)
	defer func() {
		if err != nil {
			success.Success = false
			success.Entity = err
		} else {
			success.Entity = ps
		}
		mbox.Rtn <- success
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	opts := &bind.CallOptsWithNumber{}
	opts.Context = ctx
	opts.Hash = blockHash
	minerList, amountList, blockList, ownerList, blackStatusList, err := self.poc_1.GetAll(opts)
	if err != nil {
		return
	}
	ps = &params.PocStatus{
		MinerList:       minerList,
		AmountList:      amountList,
		BlockList:       blockList,
		OwnerList:       ownerList,
		BlackStatusList: blackStatusList,
	}
	log.Debug("<<StatuteService.pocGetAll>>")
	return
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
		var (
			num  = self.ethereum.BlockChain().CurrentHeader().Number
			ctx  = context.Background()
			opts = new(bind.CallOptsWithNumber)
			i    *big.Int
		)
		vsn, err := params.MeshboxVsn(num)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		opts.Context = ctx
		switch vsn {
		case "0.0.1":
			i, err = self.meshbox_0_0_1.ExistAddress(opts, addr)
		case "0.0.2":
			i, err = self.meshbox_0_0_2.ExistAddress(opts, addr)
		}
		log.Debug("<<StatuteService.ExistAddress>>", "r", i, "err", err)
		return i, err
	default:
		return nil, errors.New("wait init")
	}
}

func (self *StatuteService) GetMeshboxList() ([]common.Address, error) {
	log.Debug("<<StatuteService.GetMeshboxList>>")
	select {
	case <-params.InitMeshbox:
		var (
			num  = self.ethereum.BlockChain().CurrentHeader().Number
			ctx  = context.Background()
			opts = new(bind.CallOptsWithNumber)
			mbs  []common.Address
		)
		vsn, err := params.MeshboxVsn(num)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		opts.Context = ctx
		switch vsn {
		case "0.0.2":
			mbs, err = self.meshbox_0_0_2.GetMeshboxList(opts)
		}
		log.Debug("<<StatuteService.GetMeshboxList>>", "err", err)
		return mbs, err
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
			case params.POC_METHOD_DEPOSIT:
				self.pocDeposit(mbox)
			case params.POC_METHOD_START:
				fallthrough
			case params.POC_METHOD_STOP:
				fallthrough
			case params.POC_METHOD_WITHDRAW:
				fallthrough
			case params.POC_METHOD_WITHDRAW_SURPLUS:
				self.pocStartAndStopAndWithdrawAndWithdrawSurplus(mbox, mbox.Method)
			case params.POC_METHOD_GET_STATUS:
				self.pocGetAll(mbox)
			}
		}
	}
}
