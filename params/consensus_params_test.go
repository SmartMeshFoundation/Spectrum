package params

import (
	"testing"
	"sort"
	"math/big"
	"github.com/SmartMeshFoundation/Spectrum/common"
)

var data = ChiefInfoList{
	newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address,TribeChief_0_0_2ABI),
	newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address,TribeChief_0_0_3ABI),
	newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address,TribeChief_0_0_4ABI),
	newChiefInfo(TestnetChainConfig.Chief005Block, "0.0.5", TestnetChainConfig.Chief005Address,TribeChief_0_0_5ABI),
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

func TestIsChiefUpdate(t *testing.T) {
	data := []byte{28,27,135,114,0,0}
	t.Log(IsChiefUpdate(data))
	data = []byte{28,27,135,115,0,0}
	t.Log(IsChiefUpdate(data))
}
