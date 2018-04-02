package params

import (
	"testing"
	"sort"
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
)

var data = ChiefInfoList{
	newChiefInfo(40000, "0.0.4", "0x04"),
	newChiefInfo(3, "0.0.2", "0x02"),
	newChiefInfo(60000, "0.0.6", "0x06"),
	newChiefInfo(30000, "0.0.3", "0x03"),
	newChiefInfo(50000, "0.0.5", "0x05"),
}

func TestChiefInfoList(t *testing.T) {
	t.Log(data)
	sort.Sort(data)
	t.Log(data)
}

func TestGetChiefAddress(t *testing.T) {
	sort.Sort(data)
	t.Log(getChiefAddress(data,big.NewInt(5)))
	t.Log(getChiefAddress(data,big.NewInt(50010)))
	t.Log(getChiefAddress(data,big.NewInt(20010)))
	t.Log(getChiefAddress(data,big.NewInt(30000)))
	t.Log(getChiefAddress(data,big.NewInt(40010)))
	t.Log(getChiefAddress(data,big.NewInt(80000)))
}

func TestIsChiefAddress(t *testing.T) {
	t.Log(isChiefAddress(data,common.HexToAddress("0x04")))
	t.Log(isChiefAddress(data,common.HexToAddress("0x0f")))
}