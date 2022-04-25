package params

import (
	"bytes"
	"encoding/hex"
	"github.com/MeshBoxTech/mesh-chain/accounts/abi"
	"github.com/MeshBoxTech/mesh-chain/common"
	"math/big"
	"sort"
	"strings"
	"sync"
	"testing"
)

var data = ChiefInfoList{
	newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address, TribeChief_0_0_2ABI),
	newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address, TribeChief_0_0_3ABI),
	newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address, TribeChief_0_0_4ABI),
	newChiefInfo(TestnetChainConfig.Chief005Block, "0.0.5", TestnetChainConfig.Chief005Address, TribeChief_0_0_5ABI),
}

func TestChiefInfoList(t *testing.T) {
	t.Log(data)
	sort.Sort(data)
	t.Log(data)
}

func TestGetChiefAddress(t *testing.T) {
	sort.Sort(data)
	t.Log(getChiefInfo(data, big.NewInt(5)))
	t.Log(getChiefInfo(data, big.NewInt(50010)))
	t.Log(getChiefInfo(data, big.NewInt(20010)))
	t.Log(getChiefInfo(data, big.NewInt(30000)))
	t.Log(getChiefInfo(data, big.NewInt(40010)))
	t.Log(getChiefInfo(data, big.NewInt(80000)))
}

func TestIsChiefAddress(t *testing.T) {
	t.Log(isChiefAddress(data, common.HexToAddress("0x04")))
	t.Log(isChiefAddress(data, common.HexToAddress("0x0f")))
}

type ChiefStatus1 struct {
	NumberList []*big.Int
	BlackList  []*big.Int
}

type ChiefStatus2 struct {
	NumberList []*big.Int
	BlackList  []*big.Int
}

func TestFooBar(t *testing.T) {
	a := ChiefStatus1{[]*big.Int{big.NewInt(1)}, []*big.Int{big.NewInt(2)}}
	t.Log(a)
	b := ChiefStatus2(a)
	t.Log(b)
	var x []common.Address
	x = nil
	t.Log(x == nil)
}

func TestAddr(t *testing.T) {
	add1 := common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88aB49")
	add2 := common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49")
	t.Log(add1 == add2)
}

func TestRegisterContract(t *testing.T) {
	t.Log(TribeChief_0_0_6ABI)
	hexdata := "1c1b87720000000000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(hexdata)
	t.Log("1 err=", err, data)
	_abi, err := abi.JSON(strings.NewReader(TribeChief_0_0_2ABI))
	t.Log("2 err=", err)
	method := _abi.Methods["update"]
	id := new(common.Address)
	r := []interface{}{id}
	t.Log(len(data[4:]), len(data), data[:4])
	err = method.Inputs.Unpack(id, data[4:])
	t.Log("4 err=", err, r, id.Hex())

	t.Log(bytes.Equal(data[:4], []byte{28, 27, 135, 114}), data[:4])

	rrr, _ := _abi.Pack("update", common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"))
	t.Log(rrr[4:])
	aaa := common.Bytes2Hex(rrr[4:])
	t.Log(common.HexToAddress(aaa) == common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"))
	bbb := common.Bytes2Hex([]byte{0, 0, 0, 0, 0, 0, 0})
	t.Log(common.HexToAddress(bbb) == common.HexToAddress(""))
}

func TestError(t *testing.T) {
	ch := make(chan int)
	sm := new(sync.Map)
	sm.Store("foo", "bar")
	sm.Store("hello", "world")

	sm.Range(func(k, v interface{}) bool {
		defer func() {
			if err := recover(); err != nil {
				t.Log(k, v, "err:", err)
			}
		}()
		defer close(ch)
		t.Log(k, v)
		return true
	})

}
