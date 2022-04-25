package tribe

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/MeshBoxTech/mesh-chain/ethclient"
	"math/big"
	"time"

	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/crypto"
	"github.com/MeshBoxTech/mesh-chain/params"
)

func (api *API) BindSign(from *common.Address) (string, error) {
	if from == nil {
		return "", errors.New("args_can_not_empty")
	}
	nodekey := api.tribe.getNodekey()
	msg := crypto.Keccak256(from.Bytes())
	sig, err := crypto.Sign(msg, nodekey)
	if err != nil {
		return "", err
	}
	sigHex := hex.EncodeToString(sig)
	return sigHex, nil
}

func (api *API) BindInfo(addr *common.Address, num *big.Int) (map[string]interface{}, error) {
	if addr == nil {
		nodekey := api.tribe.getNodekey()
		_addr := crypto.PubkeyToAddress(nodekey.PublicKey)
		addr = &_addr
	}
	header := api.chain.CurrentHeader()
	if num != nil && num.Cmp(api.chain.CurrentHeader().Number) <= 0{
		header  = api.chain.GetHeaderByNumber(num.Uint64())
	}
	bindinfo,err := api.tribe.getBindInfo(api.chain,header,*addr)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["from"] = bindinfo.From
	m["nodeids"] = bindinfo.Nids
	return m, nil
}
func (api *API) GetMiner() (*TribeMiner, error) {
	add := api.tribe.GetMinerAddress()
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
	return &TribeMiner{add, b}, nil
}

func (api *API) GetValidators(number *big.Int) ([]common.Address, error) {
	header := api.chain.CurrentHeader()
	if number != nil && number.Cmp(api.chain.CurrentHeader().Number) <= 0{
		header  = api.chain.GetHeaderByNumber(number.Uint64())
	}

	num := header.Number.Uint64()
	for num%api.tribe.config.Epoch != 0 {
		num -= 1
	}
	if num < header.Number.Uint64(){
		header = api.chain.GetHeaderByNumber(num)
	}
	return api.tribe.getNewValidators(api.chain,header)
}
