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
)

func TestSigners(t *testing.T) {
}

func TestTribeStatus_InTurn(t *testing.T) {
}