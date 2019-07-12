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
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := header.Hash()
		n := header.Number
		if hash != nil {
			h = *hash
			n = api.chain.GetHeaderByHash(h).Number
		}
		api.tribe.Status.LoadSignersFromChief(h, n)
	}
	return api.tribe.Status.Signers, nil
}

func (api *API) GetStatus(hash *common.Hash) (*TribeStatus, error) {
	header := api.chain.CurrentHeader()
	if header.Number.Int64() == 0 {
		api.tribe.Status.genesisSigner(header)
	} else {
		h := header.Hash()
		n := header.Number
		if hash != nil {
			h = *hash
			n = api.chain.GetHeaderByHash(h).Number
		}
		api.tribe.Status.LoadSignersFromChief(h, n)
		if chiefInfo := params.GetChiefInfo(n); chiefInfo != nil {
			api.tribe.Status.Vsn = chiefInfo.Version
		}
	}
	return api.tribe.Status, nil
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
