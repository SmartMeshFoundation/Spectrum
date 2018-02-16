package tribe

import (
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/accounts"
	"github.com/SmartMeshFoundation/SMChain/consensus"
)

type API struct {
	chain consensus.ChainReader
	tribe *Tribe
}

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account.
type SignerFn func(accounts.Account, []byte) ([]byte, error)

type Signer struct {
	Address common.Address `json:"address"` // 签名人
	Score   int64            `json:"score"`   // 分数
}

type TribeStatus struct {
	Signers []*Signer `json:"signers"` // 签名人
	Volunteers []common.Address `json:"volunteers"` //候选人
	Number  int64 `json:"number"` // last block.number
}


