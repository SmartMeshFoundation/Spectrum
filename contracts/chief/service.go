package chief

import (
	"context"
	"errors"
	"github.com/SmartMeshFoundation/Spectrum/contracts/statute"
	"math/big"
	"time"

	"crypto/ecdsa"
	"fmt"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/contracts/chief/lib"
	"github.com/SmartMeshFoundation/Spectrum/core"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/eth"
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/internal/ethapi"
	"github.com/SmartMeshFoundation/Spectrum/les"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/node"
	"github.com/SmartMeshFoundation/Spectrum/p2p"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
)

/*
type Service interface {
	Protocols() []p2p.Protocol
	APIs() []rpc.API
	Start(server *p2p.Server) error
	Stop() error
}
*/

// volunteer : peer.td - current.td < 200
var (
	min_td = big.NewInt(200)
)

//implements node.Service
type TribeService struct {
	tribeChief_0_0_2 *chieflib.TribeChief
	tribeChief_0_0_3 *chieflib.TribeChief_0_0_3
	tribeChief_0_0_4 *chieflib.TribeChief_0_0_4
	tribeChief_0_0_5 *chieflib.TribeChief_0_0_5
	tribeChief_0_0_6 *chieflib.TribeChief_0_0_6
	quit             chan int
	server           *p2p.Server // peers and nodekey ...
	ipcpath          string
	client           *ethclient.Client
	ethereum         *eth.Ethereum
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
	ipcpath := params.GetIPCPath()
	ts := &TribeService{
		quit:     make(chan int),
		ipcpath:  ipcpath,
		ethereum: ethereum,
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
	if v0_0_4 := params.GetChiefInfoByVsn("0.0.4"); v0_0_4 != nil {
		contract_0_0_4, err := chieflib.NewTribeChief_0_0_4(v0_0_4.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_4 = contract_0_0_4
	}
	if v0_0_5 := params.GetChiefInfoByVsn("0.0.5"); v0_0_5 != nil {
		contract_0_0_5, err := chieflib.NewTribeChief_0_0_5(v0_0_5.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_5 = contract_0_0_5
	}
	if v0_0_6 := params.GetChiefInfoByVsn("0.0.6"); v0_0_6 != nil {
		contract_0_0_6, err := chieflib.NewTribeChief_0_0_6(v0_0_6.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_6 = contract_0_0_6
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
			case "FilterVolunteer":
				self.filterVolunteer(mbox)
			case "GetVolunteers":
				self.getVolunteers(mbox)
			}
		}
	}
}

func (self *TribeService) Stop() error {
	close(self.quit)
	return nil
}

func (self *TribeService) getVolunteers(mbox params.Mbox) {
	var (
		blockNumber *big.Int
		blockHash   *common.Hash
		success     = params.MBoxSuccess{Success: true}
	)
	// hash and number can not nil
	if h, ok := mbox.Params["hash"]; ok {
		bh := h.(common.Hash)
		blockHash = &bh
	}
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
	}

	chiefInfo := params.GetChiefInfo(blockNumber)
	if chiefInfo == nil {
		log.Debug("=>TribeService.getVolunteers", "empty_chief", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
		success.Success = false
		success.Entity = errors.New("can_not_empty_chiefInfo")
	} else {
		switch chiefInfo.Version {
		case "0.0.6":
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			opts := new(bind.CallOptsWithNumber)
			opts.Context = ctx
			opts.Hash = blockHash
			v, err := self.tribeChief_0_0_6.GetVolunteers(opts)
			if err != nil {
				log.Error("=>TribeService.getVolunteers", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				success.Success = false
				success.Entity = err
			}
			success.Entity = params.ChiefVolunteers{v.VolunteerList, v.WeightList, v.Length}
		default:
			log.Error("=>TribeService.getVolunteers", "fail_vsn", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
			success.Success = false
			success.Entity = errors.New("fail_vsn_now")
		}
	}
	mbox.Rtn <- success
}

func (self *TribeService) filterVolunteer(mbox params.Mbox) {
	var (
		blockNumber *big.Int
		blockHash   *common.Hash
		addr        common.Address
		vlist       = make([]common.Address, 0, 1)
		success     = params.MBoxSuccess{Success: true}
	)
	// hash and number can not nil
	if h, ok := mbox.Params["hash"]; ok {
		bh := h.(common.Hash)
		blockHash = &bh
	}
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
	}
	if a, ok := mbox.Params["address"]; ok {
		addr = a.(common.Address)
		vlist = append(vlist[:], addr)
	}
	log.Debug("=>TribeService.filterVolunteer", "blockNumber", blockNumber, "blockHash", blockHash.Hex(), "addr", addr.Hex())

	chiefInfo := params.GetChiefInfo(blockNumber)
	if chiefInfo == nil {
		log.Error("=>TribeService.filterVolunteer", "empty_chief", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
		success.Success = false
		success.Entity = errors.New("cchiefInfo_can_not_empty")
	} else {
		switch chiefInfo.Version {
		case "0.0.6":
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			opts := new(bind.CallOptsWithNumber)
			opts.Context = ctx
			opts.Hash = blockHash
			rlist, err := self.tribeChief_0_0_6.FilterVolunteer(opts, vlist)
			if err != nil {
				log.Error("=>TribeService.filterVolunteer", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				success.Success = false
				success.Entity = err
			}
			success.Entity = rlist[0]
		default:
			log.Error("=>TribeService.filterVolunteer", "fail_vsn", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
			success.Success = false
			success.Entity = errors.New("fail_vsn_now")
		}
	}
	mbox.Rtn <- success
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
		log.Debug("chief.mbox.rtn: getstatus <-", "success", success.Success, "err", err)
	} else {
		entity := chiefStatus
		success.Entity = entity
		log.Debug("chief.mbox.rtn: getstatus <-", "success", success.Success)
		//log.Debug("chief.mbox.rtn: getstatus <-", "success", success.Success, "entity", entity)
	}
	mbox.Rtn <- success
}

func (self *TribeService) update(mbox params.Mbox) {
	prv := self.server.PrivateKey
	auth := bind.NewKeyedTransactor(prv)
	auth.GasPrice = eth.DefaultConfig.GasPrice
	//auth.GasLimit = params.GenesisGasLimit
	//auth.GasLimit = big.NewInt(params.ChiefTxGas.Int64())
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
		_b, _e := self.client.BlockByNumber(context.Background(), blockNumber)
		if _b == nil || _e != nil {
			log.Error("Tribe.update : getBlockError", "err", _e, "num", blockNumber.Int64())
			success.Entity = errors.New(fmt.Sprintf("TribeService.update : get_block_error : %d", blockNumber))
			mbox.Rtn <- success
			return
		}
		auth.GasLimit = core.CalcGasLimit(_b)
		log.Debug("-> TribeService.update", "blockNumber", blockNumber.Int64())
	} else {
		success.Entity = errors.New("TribeService.update : blockNumber not nil")
		mbox.Rtn <- success
		return
	}
	if chiefInfo := params.GetChiefInfo(blockNumber); chiefInfo != nil {
		switch chiefInfo.Version {
		case "0.0.2":
			t, e = self.tribeChief_0_0_2.Update(auth, self.fetchVolunteer(blockNumber, chiefInfo.Version))
		case "0.0.3":
			t, e = self.tribeChief_0_0_3.Update(auth, self.fetchVolunteer(blockNumber, chiefInfo.Version))
		case "0.0.4":
			t, e = self.tribeChief_0_0_4.Update(auth, self.fetchVolunteer(blockNumber, chiefInfo.Version))
		case "0.0.5":
			t, e = self.tribeChief_0_0_5.Update(auth, self.fetchVolunteer(blockNumber, chiefInfo.Version))
		case "0.0.6":
			t, e = self.tribeChief_0_0_6.Update(auth, self.fetchVolunteer(blockNumber, chiefInfo.Version))
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
			return params.ChiefStatus{
				VolunteerList: chiefStatus.VolunteerList,
				SignerList:    chiefStatus.SignerList,
				ScoreList:     chiefStatus.ScoreList,
				NumberList:    chiefStatus.NumberList,
				BlackList:     nil,
				Number:        chiefStatus.Number,
			}, nil
		case "0.0.4":
			chiefStatus, err := self.tribeChief_0_0_4.GetStatus(opts)
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
		case "0.0.5":
			chiefStatus, err := self.tribeChief_0_0_5.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			epoch, err := self.tribeChief_0_0_5.GetEpoch(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			signerLimit, err := self.tribeChief_0_0_5.GetSignerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			volunteerLimit, err := self.tribeChief_0_0_5.GetVolunteerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			return params.ChiefStatus{
				VolunteerList:  chiefStatus.VolunteerList,
				SignerList:     chiefStatus.SignerList,
				ScoreList:      chiefStatus.ScoreList,
				NumberList:     chiefStatus.NumberList,
				BlackList:      chiefStatus.BlackList,
				Number:         chiefStatus.Number,
				Epoch:          epoch,
				SignerLimit:    signerLimit,
				VolunteerLimit: volunteerLimit,
			}, nil
		case "0.0.6":
			chiefStatus, err := self.tribeChief_0_0_6.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			epoch, err := self.tribeChief_0_0_6.GetEpoch(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			signerLimit, err := self.tribeChief_0_0_6.GetSignerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			volunteerLimit, err := self.tribeChief_0_0_6.GetVolunteerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			return params.ChiefStatus{
				VolunteerList:  nil,
				SignerList:     chiefStatus.SignerList,
				ScoreList:      chiefStatus.ScoreList,
				NumberList:     chiefStatus.NumberList,
				BlackList:      chiefStatus.BlackList,
				Number:         chiefStatus.Number,
				Epoch:          epoch,
				SignerLimit:    signerLimit,
				VolunteerLimit: volunteerLimit,
				TotalVolunteer: chiefStatus.TotalVolunteer,
			}, nil
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

//0.0.6 : volunteerList is nil on vsn0.0.6
func (self *TribeService) fetchVolunteer(blockNumber *big.Int, vsn string) common.Address {
	ch := self.ethereum.BlockChain().CurrentHeader()
	TD := self.ethereum.BlockChain().GetTd(ch.Hash(), ch.Number.Uint64())
	min := new(big.Int).Sub(TD, min_td)
	vs := self.ethereum.FetchVolunteers(min, func(pk *ecdsa.PublicKey) bool {
		log.Debug("fetchVolunteer_callback", "vsn", vsn)
		if vsn == "0.0.6" {
			return params.CanNomination(pk)
		}
		return true
	})

	if len(vs) > 0 {
		chiefStatus, err := self.getChiefStatus(blockNumber, nil)
		if err != nil {
			log.Error("getChiefStatus fail", "err", err)
		}
		// exclude signers
		vl := chiefStatus.SignerList
		if chiefStatus.VolunteerList != nil {
			// exclude volunteers
			vl = append(vl[:], chiefStatus.VolunteerList...)
		}
		if chiefStatus.BlackList != nil {
			// exclude blacklist
			vl = append(vl[:], chiefStatus.BlackList...)
		}
		switch vsn {
		case "0.0.2", "0.0.3", "0.0.4", "0.0.5":
			vmap := make(map[common.Address]interface{})
			for _, v := range vl {
				vmap[v] = struct{}{}
			}
			for _, pub := range vs {
				add := crypto.PubkeyToAddress(*pub)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				b, e := self.client.BalanceAt(ctx, add, nil)
				if e == nil && b.Cmp(params.ChiefBaseBalance) >= 0 && self.isVolunteer(vmap, add) {
					return add
				}
			}
		case "0.0.6", "0.0.7":
			vlist := make([]common.Address, 0, 0)
			for _, pub := range vs {
				add := crypto.PubkeyToAddress(*pub)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				b, e := self.client.BalanceAt(ctx, add, nil)
				if e == nil && b.Cmp(params.ChiefBaseBalance) >= 0 {
					vlist = append(vlist, add)
				}
			}
			log.Debug("=> TribeService.fetchVolunteer :", "vsn", vsn, "vlist", len(vlist))
			if len(vlist) > 0 {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				opts := new(bind.CallOptsWithNumber)
				opts.Context = ctx
				if rlist, err := self.tribeChief_0_0_6.FilterVolunteer(opts, vlist); err == nil {
					log.Debug("=> TribeService.fetchVolunteer :", "vsn", vsn, "len", len(rlist), "rlist", rlist)
					sdb, _ := self.ethereum.BlockChain().State()
					for i, r := range rlist {
						if r.Int64() > 0 {
							v := vlist[i]
							if ms, err := statute.GetMeshboxService(); err == nil {
								// first : meshbox.sol
								log.Debug("<< FilterVolunteer.meshbox-rule >>")
								if w, err := ms.ExistAddress(v); err == nil && w != nil && w.Int64() > 0 {
									return v
								}
								log.Debug("<< FilterVolunteer.skip >> not_a_meshbox", "addr", v.Hex())
							}
							// second : if vlist not in meshbox contract the balance great than 10w SMT is requirement
							// check nodeid&account mapping contract and review balance
							var _m = v
							if as, err := statute.GetAnmapService(); err == nil {
								cbh := ch.Hash()
								if f, _, err := as.BindInfo(v, nil, &cbh); err == nil {
									log.Info("1.<<statute.GetAnmapService>>", "f", f.Hex(), "n", v.Hex())
									_m = f
								}
							}
							b := sdb.GetBalance(_m)
							if b.Cmp(params.GetMinMinerBalance()) >= 0 {
								log.Error("2.<<statute.GetAnmapService>>", "v", v.Hex(), "_m", _m.Hex(), "balance", b.String())
								return v
							}
						}
					}

				}

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
