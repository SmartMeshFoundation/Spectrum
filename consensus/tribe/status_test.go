package tribe

import (
	"github.com/SmartMeshFoundation/Spectrum/common"
	"testing"
)

var (
	address1 = common.HexToAddress("0x4110bd1ff0b73fa12c259acf39c950277f266787")
	address2 = common.HexToAddress("0xadb3ea3ad356199206ca817b04fd668cc5196df2")
	address3 = common.HexToAddress("0xb94b3aa41609e3f59cbaff3c2c298c6cc4c50b81")

	signerList1 = []*Signer{{address1, 2}}
	signerList2 = []*Signer{{address2, 2}, {address3, 2}}
	signerList3 = []*Signer{{address1, 3}, {address2, 3}, {address3, 3}}

	tribe = &TribeStatus{
		Signers: signerList3,
	}
)

type BUF []byte

func (a BUF) Str() string { return string(a[:]) }
func TestSigners(t *testing.T) {

}

func TestMinerSigner(t *testing.T) {
	miner := "0x4110bd1ff0b73fa12c259acf39c950277f266787"
	miner2 := "0x4110bd1ff0b73fa12c259acf39c950277f266787"
	add := common.HexToAddress(miner)
	add2 := common.HexToAddress(miner2)

	t.Log(add)
	t.Log(add == add2)
}

func TestAddr(t *testing.T) {
	var addr common.Address
	t.Log(common.HexToAddress("0x") == addr)
	t.Log(common.HexToAddress("0x") == common.HexToAddress(""))
}

func TestLeaderSort(t *testing.T) {
	f := common.HexToAddress("0x3")
	list := []common.Address{
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
	}
	r, err := leaderSort(f, list)
	t.Log(err)
	t.Log(list)
	t.Log(r)
}
