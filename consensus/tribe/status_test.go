package tribe

import (
	"testing"
	"github.com/SmartMeshFoundation/SMChain/common"
)

var (
	address1 = common.HexToAddress("0x1")
	address2 = common.HexToAddress("0x2")
	address3 = common.HexToAddress("0x3")

	signerList1 = []Signer{{address1, 2}}
	signerList2 = []Signer{{address2, 2}, {address3, 2}}
	signerList3 = []Signer{{address1,3},{address2, 3}, {address3, 3}}
	signers     = NewSigners()
)

func TestSigners(t *testing.T) {
	signers.Reload(signerList1)
	t.Log(signers.Get(address2))
	t.Log(signers.Len())
	signers.Reload(signerList2)
	idx, r, e := signers.Get(address3)
	t.Log(idx, *r, e)
	t.Log(signers.Len())
}

func TestTribeStatus_InTurn(t *testing.T) {
	status := NewTribeStatus()
	status.signers.Reload(signerList3)
	for i:=1 ;i<10 ; i++ {
		t.Log(i,status.InTurn(int64(i),address2))
	}
}