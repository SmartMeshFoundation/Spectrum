package tribe

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts"
	"github.com/SmartMeshFoundation/Spectrum/accounts/keystore"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/SmartMeshFoundation/Spectrum/rpc"
)

// fetchKeystore retrives the encrypted keystore from the account manager.
func fetchKeystore(am *accounts.Manager) *keystore.KeyStore {
	return am.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
}

func (api *API) BindSign(from *common.Address) (string, error) {
	if from == nil {
		return "", errors.New("args_can_not_empty")
	}
	nodekey := api.tribe.Status.getNodekey()
	msg := crypto.Keccak256(from.Bytes())
	sig, _ := crypto.Sign(msg, nodekey)
	sigHex := hex.EncodeToString(sig)
	return sigHex, nil
}

func (api *API) BindInfo(addr *common.Address, num *big.Int) (map[string]interface{}, error) {
	if addr == nil {
		nodekey := api.tribe.Status.getNodekey()
		_addr := crypto.PubkeyToAddress(nodekey.PublicKey)
		addr = &_addr
	}
	hash := api.chain.CurrentHeader().Hash()
	if num != nil {
		hash = api.chain.GetHeaderByNumber(num.Uint64()).Hash()
	}
	from, nodeids, err := params.AnmapBindInfo(*addr, hash)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["from"] = from
	m["nodeids"] = nodeids
	return m, nil
}

func (api *API) Bind(from *common.Address, passphrase string) (string, error) {
	if from == nil {
		return "", errors.New("args_can_not_empty")
	}
	a := accounts.Account{Address: *from}
	e := fetchKeystore(api.accman).TimedUnlock(a, passphrase, 60*time.Second)
	if e != nil {
		return "", e
	}
	nodekey := api.tribe.Status.getNodekey()
	nodeid := crypto.PubkeyToAddress(nodekey.PublicKey)
	msg := crypto.Keccak256(from.Bytes())
	sig, _ := crypto.Sign(msg, nodekey)
	sigHex := hex.EncodeToString(sig)
	tx, e := params.AnmapBind(*from, nodeid, sigHex)
	if e != nil {
		return "", e
	}
	log.Info("tribe.bind", "tx", tx, "from", from.Hex(), "nodeid", nodeid.Hex())
	return tx, nil
}

//PocDeposit 调用poc合约的deposit质押,满足金额以后,自动开启挖矿
func (api *API) PocDeposit(from *common.Address, passphrase string) (string, error) {
	if from == nil {
		return "", errors.New("args_can_not_empty")
	}
	a := accounts.Account{Address: *from}
	e := fetchKeystore(api.accman).TimedUnlock(a, passphrase, 60*time.Second)
	if e != nil {
		return "", e
	}
	nodekey := api.tribe.Status.getNodekey()
	nodeid := crypto.PubkeyToAddress(nodekey.PublicKey)
	msg := crypto.Keccak256(from.Bytes())
	sig, _ := crypto.Sign(msg, nodekey)
	sigHex := hex.EncodeToString(sig)
	tx, e := params.PocDeposit(*from, sigHex)
	if e != nil {
		return "", e
	}
	log.Info("tribe.pocDeposit", "tx", tx, "from", from.Hex(), "nodeid", nodeid.Hex())
	return tx, nil
}

//PocStart 调用poc合约的start,从stop状态恢复,开启挖矿
func (api *API) PocStart(from *common.Address, passphrase string) (string, error) {
	return api.pocHelper(from, passphrase, params.POC_METHOD_START)
}

//PocStop 调用poc合约的stop,停止挖矿,准备撤回抵押
func (api *API) PocStop(from *common.Address, passphrase string) (string, error) {
	return api.pocHelper(from, passphrase, params.POC_METHOD_STOP)
}

//PocWithdraw 调用poc合约的withdraw,在stop两周后,可以撤回押金
func (api *API) PocWithdraw(from *common.Address, passphrase string) (string, error) {
	return api.pocHelper(from, passphrase, params.POC_METHOD_WITHDRAW)
}

//PocWithdrawSurplus 调用poc合约的PocWithdrawSurplus,从合约中撤回多余的抵押押金
func (api *API) PocWithdrawSurplus(from *common.Address, passphrase string) (string, error) {
	return api.pocHelper(from, passphrase, params.POC_METHOD_WITHDRAW_SURPLUS)
}
func (api *API) pocHelper(from *common.Address, passphrase string, method string) (string, error) {
	if from == nil {
		return "", errors.New("args_can_not_empty")
	}
	var (
		tx string
		e  error
	)
	a := accounts.Account{Address: *from}
	e = fetchKeystore(api.accman).TimedUnlock(a, passphrase, 60*time.Second)
	if e != nil {
		return "", e
	}
	nodekey := api.tribe.Status.getNodekey()
	nodeID := crypto.PubkeyToAddress(nodekey.PublicKey)
	switch method {
	case params.POC_METHOD_START:
		tx, e = params.PocStart(*from, nodeID)
	case params.POC_METHOD_STOP:
		tx, e = params.PocStop(*from, nodeID)
	case params.POC_METHOD_WITHDRAW:
		tx, e = params.PocWithdraw(*from, nodeID)
	case params.POC_METHOD_WITHDRAW_SURPLUS:
		tx, e = params.PocWithdrawSurplus(*from, nodeID)
	}
	if e != nil {
		return "", e
	}
	log.Info("tribe.", "method", method, "tx", tx, "from", from.Hex(), "nodeid", nodeID.Hex())
	return tx, nil
}

//PocGetStatus 查询poc合约状态
func (api *API) PocGetStatus() (*params.PocStatus, error) {
	return params.PocGetAll()
}
func (api *API) Unbind(from *common.Address, passphrase string) (string, error) {
	nodekey := api.tribe.Status.getNodekey()
	nodeid := crypto.PubkeyToAddress(nodekey.PublicKey)
	if from == nil {
		if m, err := api.BindInfo(&nodeid, nil); err != nil {
			return "", err
		} else {
			_from := m["from"].(common.Address)
			from = &_from
		}
	}
	a := accounts.Account{Address: *from}
	e := fetchKeystore(api.accman).TimedUnlock(a, passphrase, 60*time.Second)
	if e != nil {
		return "", e
	}
	msg := crypto.Keccak256(from.Bytes())
	sig, _ := crypto.Sign(msg, nodekey)
	sigHex := hex.EncodeToString(sig)
	tx, e := params.AnmapUnbind(*from, nodeid, sigHex)
	if e != nil {
		return "", e
	}
	log.Info("tribe.unbind", "tx", tx, "from", from.Hex(), "nodeid", nodeid.Hex())
	return tx, nil
}

func (api *API) GetMiner(number *rpc.BlockNumber) (*TribeMiner, error) {
	add := api.tribe.Status.GetMinerAddress()
	ipcpath := params.GetIPCPath()
	c, e := ethclient.Dial(ipcpath)
	if e != nil {
		return nil, e
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	b, e := c.BalanceAt(ctx, add, nil)
	if e != nil {
		return nil, e
	}
	return &TribeMiner{add, b, api.tribe.Status.SignerLevel}, nil
}

// chief-0.0.3 show blacklist
func (api *API) GetSinners(hash *common.Hash) ([]common.Address, error) {
	return api.tribe.Status.blackList, nil
}

func (api *API) GetSigners(hash *common.Hash) ([]*Signer, error) {
	header := api.chain.CurrentHeader()

	h := header.Hash()
	n := header.Number
	if hash != nil {
		h = *hash
		n = api.chain.GetHeaderByHash(h).Number
	}
	status, err := api.loadHistoryChiefStatus(h, n)
	if err != nil {
		return nil, err
	}
	return status.Signers, nil
}

func (api *API) GetStatus(hash *common.Hash) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	h := header.Hash()
	n := header.Number
	if hash != nil {
		h = *hash
		h2 := api.chain.GetHeaderByHash(h) //有可能返回0值,当hash不存在的时候
		if h2 == nil {
			return nil, errors.New("block not exist")
		}
		n = h2.Number
	}
	return api.loadHistoryChiefStatus(h, n)
}

func (api *API) GetVolunteers(hash *common.Hash) (*TribeVolunteers, error) {
	header := api.chain.CurrentHeader()
	if hash != nil {
		header = api.chain.GetHeaderByHash(*hash)
	} else {
		h := header.Hash()
		hash = &h
	}
	rtn := params.SendToMsgBoxWithHash("GetVolunteers", *hash, header.Number)
	r := <-rtn
	if !r.Success {
		return nil, r.Entity.(error)
	}
	cv := r.Entity.(params.ChiefVolunteers)
	vs := &TribeVolunteers{cv.Length, make([]*Volunteer, 0, 0)}
	if cv.Length != nil && cv.Length.Int64() > 0 {
		for i, volunteer := range cv.VolunteerList {
			weight := cv.WeightList[i]
			vs.Volunteers = append(vs.Volunteers, &Volunteer{volunteer, weight.Int64()})
		}
	}
	return vs, nil
}

func (api *API) GetHistory(last *big.Int, noRpc *bool) (interface{}, error) {
	s := uint64(16)
	if last != nil {
		s = last.Uint64()
	}
	cn := api.chain.CurrentHeader().Number.Uint64()
	if noRpc != nil && *noRpc {
		_history := make([]map[string]string, 0)
		for i := cn; i > cn-s; i-- {
			_header := api.chain.GetHeaderByNumber(i)
			k := "🔨"
			v := fmt.Sprintf("%d -> %s", _header.Number.Int64(), _header.Coinbase.Hex())
			if _header.Difficulty.Int64() == 1 {
				k = "👿"
			}
			_h := map[string]string{k: v}
			_history = append(_history, _h)
		}
		return _history, nil
	} else {
		_history := make([]History, 0)
		for i := cn; i > cn-s; i-- {
			_header := api.chain.GetHeaderByNumber(i)
			_h := History{_header.Number.Int64(), _header.Hash(), _header.Coinbase, _header.Difficulty, _header.Time}
			_history = append(_history[:], _h)
		}
		return _history, nil
	}
}

// 在 加载完所有 node.service 后，需要主动调用一次
func (api *API) loadHistoryChiefStatus(hash common.Hash, number *big.Int) (status *TribeStatus, err error) {
	//log.Info(fmt.Sprintf("LoadSignersFromChief hash=%s,number=%s", hash.String(), number))
	cs, err := params.TribeGetStatus(number, hash)
	if err != nil {
		return nil, err
	}
	status = &TribeStatus{}
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	status.LeaderLimit = cs.LeaderLimit
	status.Leaders = cs.LeaderList
	if len(status.Leaders) == 0 {
		panic(fmt.Sprintf("LoadSignersFromChief err ,hash=%s,number=%s,cs=%#v", hash.String(), number, cs))
	}
	status.Number = cs.Number.Int64()
	status.blackList = cs.BlackList
	status.Signers = sl
	status.Epoch, status.SignerLimit = cs.Epoch, cs.SignerLimit
	return
}
