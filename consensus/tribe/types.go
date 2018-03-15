package tribe

import (
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/accounts"
	"github.com/SmartMeshFoundation/SMChain/consensus"
	"crypto/ecdsa"
	"math/big"
	"fmt"
	"sync"
	"github.com/SmartMeshFoundation/SMChain/params"
	"github.com/SmartMeshFoundation/SMChain/ethdb"
	"github.com/hashicorp/golang-lru"
)

const (
	// None -> Volunteer -> Signer
	LevelNone      = "None"
	LevelVolunteer = "Volunteer"
	LevelSigner    = "Signer"
)

type Tribe struct {
	config   *params.TribeConfig // Consensus engine configuration parameters
	db       ethdb.Database      // Database to store and retrieve snapshot checkpoints
	sigcache *lru.ARCCache       // mapping block.hash -> signer
	//signer      common.Address      // Ethereum address of the signing key
	signFn SignerFn     // Signer function to authorize hashes with
	lock   sync.RWMutex // Protects the signer fields
	Status *TribeStatus
}

type API struct {
	chain consensus.ChainReader
	tribe *Tribe
}

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account.
type SignerFn func(accounts.Account, []byte) ([]byte, error)

type Signer struct {
	Address common.Address `json:"address"` // 签名人
	Score   int64          `json:"score"`   // 分数
}

func (self *Signer) String() string {
	return fmt.Sprintf("%s:%d", self.Address.Hex(), self.Score)
}

type TribeStatus struct {
	Signers     []*Signer        `json:"signers"`    // 签名人
	Volunteers  []common.Address `json:"volunteers"` //候选人
	SignerLevel string                               // None -> Volunteer -> Signer
	Number      int64            `json:"number"`     // last block.number
	mining      int32                                // 1 mining start , 0 mining stop
	nodeKey     *ecdsa.PrivateKey
}

type TribeMiner struct {
	Address common.Address `json:"address"`
	Balance *big.Int       `json:"balance"`
	Level   string         `json:"level"` // None 、 Volunteer 、 Signer
}
