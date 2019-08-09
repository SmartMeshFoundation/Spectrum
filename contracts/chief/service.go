package chief

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/SmartMeshFoundation/Spectrum/contracts"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"math/big"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/contracts/statute"

	"fmt"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	chieflib "github.com/SmartMeshFoundation/Spectrum/contracts/chief/lib"
	"github.com/SmartMeshFoundation/Spectrum/eth"
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
	tribeChief_0_0_7 *chieflib.TribeChief_0_0_7
	tribeChief_1_0_0 *chieflib.TribeChief_1_0_0
	poc              *chieflib.POC_1_0_0
	base             *chieflib.ChiefBase_1_0_0
	quit             chan int
	server           *p2p.Server // peers and nodekey ...
	ethereum         *eth.Ethereum
	ctx              *node.ServiceContext
}

func NewTribeService(ctx *node.ServiceContext) (node.Service, error) {
	var (
		apiBackend ethapi.Backend
		ethereum   *eth.Ethereum
	)
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

	ts := &TribeService{
		quit:     make(chan int),
		ethereum: ethereum,
		ctx:      ctx,
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
	if v0_0_7 := params.GetChiefInfoByVsn("0.0.7"); v0_0_7 != nil {
		contract_0_0_7, err := chieflib.NewTribeChief_0_0_7(v0_0_7.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		ts.tribeChief_0_0_7 = contract_0_0_7
	}
	if v1_0_0 := params.GetChiefInfoByVsn("1.0.0"); v1_0_0 != nil {
		ab := eth.NewContractBackend(apiBackend)
		contract_1_0_0, err := chieflib.NewTribeChief_1_0_0(v1_0_0.Addr, ab)
		if err != nil {
			return nil, err
		}

		ts.tribeChief_1_0_0 = contract_1_0_0
		poc, err := chieflib.NewPOC_1_0_0(v1_0_0.PocAddr, ab)
		if err != nil {
			return nil, err
		}
		ts.poc = poc

		base, err := chieflib.NewChiefBase_1_0_0(v1_0_0.BaseAddr, ab)
		if err != nil {
			return nil, err
		}
		ts.base = base

		log.Info("<<TribeService>> chief-1.0.0 and poc init success.")
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
			case params.Chief100Update:
				self.chief100FetchNextRoundSigner(mbox)
			case "FilterVolunteer":
				self.filterVolunteer(mbox)
			case "GetVolunteers":
				self.getVolunteers(mbox)
			case "VerifyMiner": // for 1.0.0 vrf selected
				self.VerifyMiner(mbox)

			}
		}
	}
}

func (self *TribeService) Stop() error {
	close(self.quit)
	return nil
}

func (self *TribeService) _getVolunteers(blockNumber *big.Int, blockHash common.Hash) (params.ChiefVolunteers, error) {
	var (
		empty     = params.ChiefVolunteers{}
		chiefInfo = params.GetChiefInfo(blockNumber)
	)
	if chiefInfo == nil {
		log.Debug("=>TribeService.getVolunteers", "empty_chief", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
		return empty, errors.New("can_not_empty_chiefInfo")
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		opts.Hash = &blockHash
		switch chiefInfo.Version {
		case "0.0.6":
			v, err := self.tribeChief_0_0_6.GetVolunteers(opts)
			if err != nil {
				log.Error("=>TribeService.getVolunteers", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				return empty, err
			}
			return params.ChiefVolunteers{
				VolunteerList: v.VolunteerList,
				WeightList:    v.WeightList,
				Length:        v.Length,
			}, nil
		case "0.0.7":
			v, err := self.tribeChief_0_0_7.GetVolunteers(opts)
			if err != nil {
				log.Error("=>TribeService.getVolunteers", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				return empty, err
			}
			return params.ChiefVolunteers{
				VolunteerList: v.VolunteerList,
				WeightList:    v.WeightList,
				Length:        v.Length,
			}, nil
		case "1.0.0":
			// TODO
			v, err := self.tribeChief_1_0_0.GetVolunteers(opts)
			if err != nil {
				log.Error("=>TribeService.getVolunteers", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				return empty, err
			}
			return params.ChiefVolunteers{
				VolunteerList: v.VolunteerList,
				WeightList:    v.WeightList,
				Length:        v.Length,
			}, nil
		default:
			log.Error("=>TribeService.getVolunteers", "fail_vsn", chiefInfo.Version, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
			return empty, errors.New("fail_vsn_now")
		}
	}
}

func (self *TribeService) getVolunteers(mbox params.Mbox) {
	var (
		blockNumber *big.Int
		blockHash   common.Hash
		success     = params.MBoxSuccess{Success: true}
	)
	// hash and number can not nil
	if h, ok := mbox.Params["hash"]; ok {
		bh := h.(common.Hash)
		blockHash = bh
	}
	if n, ok := mbox.Params["number"]; ok {
		blockNumber = n.(*big.Int)
	}
	entity, err := self._getVolunteers(blockNumber, blockHash)
	if err == nil {
		success.Entity = entity
	} else {
		success.Success = false
		success.Entity = err
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
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		opts.Hash = blockHash
		switch chiefInfo.Version {
		case "0.0.6":
			rlist, err := self.tribeChief_0_0_6.FilterVolunteer(opts, vlist)
			if err != nil {
				log.Error("=>TribeService.filterVolunteer", "err", err, "blockNumber", blockNumber, "blockHash", blockHash.Hex())
				success.Success = false
				success.Entity = err
			}
			success.Entity = rlist[0]
		case "0.0.7":
			rlist, err := self.tribeChief_0_0_7.FilterVolunteer(opts, vlist)
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
func (self *TribeService) chief100FetchNextRoundSigner(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var (
		blockNumber *big.Int
		hash        common.Hash
		vrfn        *big.Int
		v           common.Address
	)
	hash = mbox.Params["hash"].(common.Hash)
	blockNumber = mbox.Params["number"].(*big.Int) //如果出错.立即崩溃亏即可
	vrfn = mbox.Params["vrfn"].(*big.Int)
	if params.IsSIP100Block(blockNumber) {
		nl := self.minerList(blockNumber, hash)
		v = self.takeMiner(nl, hash, vrfn.Bytes())
		log.Debug("chief.mbox.rtn chief100: update <-", "addr", v.String())
	} else {
		ch := self.ethereum.BlockChain().GetHeaderByHash(hash)
		TD := self.ethereum.BlockChain().GetTd(hash, blockNumber.Uint64())
		min := new(big.Int).Sub(TD, min_td)
		vsn := params.GetChiefInfo(blockNumber).Version
		vs := self.ethereum.FetchVolunteers(min, func(pk *ecdsa.PublicKey) bool {
			log.Debug("fetchVolunteer_callback", "vsn", vsn)
			if vsn == "0.0.6" || vsn == "0.0.7" {
				return params.CanNomination(pk)
			}
			return true
		})
		client, err := contracts.GetEthclientInstance()
		if err != nil {
			panic(fmt.Sprintf("GetEthclientInstance err:%s", err))
		}
		v = self.fetchVolunteer_0_0_6(client, blockNumber, vsn, vs, ch)
	}
	success.Success = true
	success.Entity = v
	mbox.Rtn <- success
	log.Debug("chief.mbox.rtn: update <-", "success", success)

}

// only for devnet and testnet
func (self *TribeService) fetchVolunteer_0_0_6(client *ethclient.Client, blockNumber *big.Int, vsn string, vs []*ecdsa.PublicKey, ch *types.Header) common.Address {
	vlist := make([]common.Address, 0, 0)
	for _, pub := range vs {
		add := crypto.PubkeyToAddress(*pub)
		vlist = append(vlist, add)
	}
	log.Debug("=> TribeService.fetchVolunteer :", "vsn", vsn, "vlist", len(vlist))
	if len(vlist) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		// only one instance between 0.0.6 and 0.0.7
		var (
			err   error
			rlist []*big.Int
		)

		rlist, err = self.tribeChief_0_0_6.FilterVolunteer(opts, vlist)
		if err == nil && len(rlist) > 0 {
			v := vlist[0]
			return v
		}

	}
	return common.Address{}
}

// --------------------------------------------------------------------------------------------------
// inner private
// --------------------------------------------------------------------------------------------------
func (self *TribeService) getChiefStatus(blockNumber *big.Int, blockHash *common.Hash) (params.ChiefStatus, error) {
	log.Debug(fmt.Sprintf("[getChiefStatus],blockNumber=%s,blockHash=%s", blockNumber, blockHash.String()))
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
		case "0.0.7":
			chiefStatus, err := self.tribeChief_0_0_7.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			epoch, err := self.tribeChief_0_0_7.GetEpoch(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			signerLimit, err := self.tribeChief_0_0_7.GetSignerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			volunteerLimit, err := self.tribeChief_0_0_7.GetVolunteerLimit(opts)
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
		case "1.0.0":
			chiefStatus, err := self.tribeChief_1_0_0.GetStatus(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			epoch, err := self.tribeChief_1_0_0.GetEpoch(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			signerLimit, err := self.tribeChief_1_0_0.GetSignerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			volunteerLimit, err := self.tribeChief_1_0_0.GetVolunteerLimit(opts)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			leaderList, leaderLimit, err := self.GetLeaders(blockNumber, blockHash)
			if err != nil {
				return params.ChiefStatus{}, err
			}
			return params.ChiefStatus{
				LeaderLimit:    leaderLimit,
				LeaderList:     leaderList,
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

func (self *TribeService) VerifyMiner(mbox params.Mbox) {
	var (
		parentblockHash common.Hash
		addr            common.Address
		vrfn            []byte
		result          bool
	)

	// hash and number can not nil
	if parenthash, ok := mbox.Params["parenthash"]; ok {
		bh := parenthash.(common.Hash)
		parentblockHash = bh
	}
	if a, ok := mbox.Params["addr"]; ok {
		addr = a.(common.Address)
	}
	if a, ok := mbox.Params["vrfn"]; ok {
		vrfn = a.([]byte)
	}
	success := params.MBoxSuccess{Success: true}
	defer func() {
		success.Entity = result
		mbox.Rtn <- success
	}()
	result = self.verifyMiner(addr, parentblockHash, vrfn)
	if !result {
		log.Error("VerifyMiner failed", "hash", parentblockHash.Hex(), "miners", success.Entity)
		return
	}
}

// poc normalList and meshboxList
func (self *TribeService) minerList(num *big.Int, hash common.Hash) []common.Address {
	var (
		ss   *statute.StatuteService
		nl   = make([]common.Address, 0)
		opts = new(bind.CallOptsWithNumber)
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	opts.Context = ctx
	opts.Hash = &hash

	vll, err := self.poc.GetNormalList(opts)
	if err != nil {
		log.Error("poc.GetNormalList__fail", "err", err)
	}

	err = self.ctx.Service(&ss)
	if err != nil {
		log.Error("get_StatuteService_fail", "err", err)
	}
	//暂时不允许meshbox参与出块,如果meshbox出快了也需要抵押
	//mbs, err := ss.GetMeshboxList()
	//if err != nil {
	//	log.Error("poc.GetMeshboxList__fail", "err", err)
	//}
	//if mbs != nil && len(mbs) > 0 {
	//	nl = append(nl, mbs...)
	//}

	if vll != nil && len(vll) > 0 {
		nl = append(nl, vll...)
	}
	return nl
}

//下一轮出块节点不能包含当前这一轮的出块人以及下一轮已经被选出来的出块人
func (self *TribeService) getNextRoundSignerExcludeList(blockNumber *big.Int, blockHash common.Hash) (addrs []common.Address) {
	//为了兼容考虑,这里的volunteers是已经选出的下一轮出块人列表,可能没有满员
	vl, err := self._getVolunteers(blockNumber, blockHash)
	if err != nil {
		log.Warn("tribeservice_getVolunteers_fail", "err", err)
	}
	for _, a := range vl.VolunteerList {
		addrs = append(addrs, a)
	}
	chiefStatus, err := self.getChiefStatus(blockNumber, &blockHash)
	if err != nil {
		log.Warn("getChiefStatus", "err", err)
	}
	for _, a := range chiefStatus.SignerList {
		addrs = append(addrs, a)
	}
	return addrs
}

//从nl也就是可能出块节列表中根据vrf选择一个,如果选中的人已经在下一轮出块列表中就尝试选择下一个,takerMiner只会在1.0.0版本后使用
func (self *TribeService) takeMiner(nl []common.Address, hash common.Hash, _vrfn []byte) common.Address {
	if nl != nil && len(nl) > 0 {
		var (
			block = self.ethereum.BlockChain().GetBlockByHash(hash)
			vrfn  = new(big.Int).SetBytes(_vrfn[:])
			fn    func(_vrfn *big.Int) common.Address
		)
		if block == nil {
			panic(errors.New("get block by hash fail"))
		}
		//排除当前signerList的原因是有可能被选中作为下一轮出块节点,但是同时又
		excludes := self.getNextRoundSignerExcludeList(block.Number(), block.Hash())
		fn = func(_vrfn *big.Int) common.Address {
			m := big.NewInt(int64(len(nl)))
			x := new(big.Int).Sub(_vrfn, vrfn)
			if x.Cmp(m) >= 0 {
				return common.Address{}
			}
			idx := new(big.Int).Mod(_vrfn, m)
			addrLog := make([]string, 0)
			for _, n := range nl {
				addrLog = append(addrLog[:], n.Hex())
			}
			log.Debug("fetchVolunteer-1.0.0-volunteers", "num", block.Number(), "addrList", addrLog)
			// skip if `n` in volunteer list
			v := nl[idx.Int64()]

			for _, vol := range excludes {
				if vol == v {
					return fn(new(big.Int).Add(_vrfn, big.NewInt(1)))
				}
			}
			log.Debug("fetchVolunteer-1.0.0-final", "num", block.Number(), "idx", idx.String(), "addr", v.Hex(), "vrfn", _vrfn)
			return v
		}
		return fn(vrfn)
	}
	return common.Address{}
}

func (self *TribeService) verifyMiner(vol common.Address, hash common.Hash, vrfn []byte) bool {
	block := self.ethereum.BlockChain().GetBlockByHash(hash)
	ci := params.GetChiefInfo(block.Number())
	switch ci.Version {
	case "1.0.0":
		m := self.takeMiner(self.minerList(block.Number(), block.Hash()), hash, vrfn)
		log.Debug("<<TribeService.verifyMiner>>", "result", vol == m, "c", vol.Hex(), "t", m.Hex())
		if vol == m {
			return true
		}
	default:
		return true
	}
	return false
}

func (self *TribeService) GetLeaders(num *big.Int, hash *common.Hash) ([]common.Address, *big.Int, error) {
	ci := params.GetChiefInfo(num)
	if ci != nil {
		switch ci.Version {
		case "1.0.0":
			var (
				leaders = make([]common.Address, 0)
				opts    = new(bind.CallOptsWithNumber)
			)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			opts.Context = ctx
			opts.Hash = hash
			leaders, err := self.base.TakeLeaderList(opts)
			if err != nil {
				return nil, nil, err
			}
			limit, err := self.base.TakeLeaderLimit(opts)
			if err != nil {
				return nil, nil, err
			}
			return leaders, limit, nil
		}
	}
	return nil, nil, errors.New(fmt.Sprintf("not_support_vsn : %s", ci.Version))
}
