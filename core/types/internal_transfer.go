package types

import (
	"fmt"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/rlp"
)

//InternalTransfer:辅助合约内部smt转账
type InternalTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func (i *InternalTransfer) String() string {
	return fmt.Sprintf("{from:%s,to:%s,value:%s}", i.From.String(), i.To.String(), i.Value)
}

// Receipts is a wrapper around a Receipt array to implement DerivableList.
type InternalTransfers []*InternalTransfer

// Len returns the number of receipts in this list.
func (r InternalTransfers) Len() int { return len(r) }

// GetRlp returns the RLP encoding of one receipt from the list.
func (r InternalTransfers) GetRlp(i int) []byte {
	bytes, err := rlp.EncodeToBytes(r[i])
	if err != nil {
		panic(err)
	}
	return bytes
}
