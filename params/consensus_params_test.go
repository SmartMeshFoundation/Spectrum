package params

import (
	"testing"
	"sort"
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
)

var data = ChiefInfoList{
	newChiefInfo(big.NewInt(40000), "0.0.4", common.HexToAddress("0x04")),
	newChiefInfo(big.NewInt(3), "0.0.2", common.HexToAddress("0x02")),
	newChiefInfo(big.NewInt(60000), "0.0.6", common.HexToAddress("0x06")),
	newChiefInfo(big.NewInt(30000), "0.0.3", common.HexToAddress("0x03")),
	newChiefInfo(big.NewInt(50000), "0.0.5", common.HexToAddress("0x05")),
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
