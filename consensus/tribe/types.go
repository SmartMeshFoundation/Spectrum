package tribe

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/common/hexutil"
	"github.com/SmartMeshFoundation/Spectrum/consensus"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/ethdb"
	"github.com/SmartMeshFoundation/Spectrum/params"
	"github.com/hashicorp/golang-lru"
)

const (
	// None -> Volunteer -> Signer -> Sinner
	LevelNone      = "None"
	LevelVolunteer = "Volunteer"
	LevelSigner    = "Signer"
	LevelSinner    = "Sinner"

	historyLimit = 2048
	wiggleTime   = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
	CHIEF_NUMBER = int64(1)
)

var (
	blockPeriod = uint64(15)                               // Default minimum difference between two consecutive block's timestamps
	extraVanity = 32                                       // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65                                       // Fixed number of extra-data suffix bytes reserved for signer seal
	nonceSync   = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceAsync  = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.
	uncleHash   = types.CalcUncleHash(nil)                 // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
	diffInTurn  = big.NewInt(2)                            // Block difficulty for in-turn signatures
	diffNoTurn  = big.NewInt(1)                            // Block difficulty for out-of-turn signatures
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// only accept sync or async flag
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidNonce = errors.New("nonce not 0x00..0 or 0xff..f")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 or 2, or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid__difficulty")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// errUnauthorized is returned if a header is signed by a non-authorized entity.
	errUnauthorized = errors.New("unauthorized")

	// errWaitTransactions is returned if an empty block is attempted to be sealed
	// on an instant chain (0 second period). It's important to refuse these as the
	// block reward is zero, so an empty block just bloats the chain... fast.
	errWaitTransactions = errors.New("waiting for transactions")
	// for tribe consensus block validator
	ErrTribeNotAllowEmptyTxList = errors.New("tribe not allow empty tx list")
	ErrTribeMustContainChiefTx  = errors.New("tribe must contain chief tx")
	ErrTribeChiefCannotRepeat   = errors.New("tribe chief tx cannot repeat")
)

type Tribe struct {
	config   *params.TribeConfig // Consensus engine configuration parameters
	db       ethdb.Database      // Database to store and retrieve snapshot checkpoints
	sigcache *lru.ARCCache       // mapping block.hash -> signer
	//signer      common.Address      // Ethereum address of the signing key
	signFn           SignerFn     // Signer function to authorize hashes with
	lock             sync.RWMutex // Protects the signer fields
	Status           *TribeStatus
	SealErrorCh      chan error // when error from seal fun, may be need commit a new work
	SealErrorCounter uint32     // less then 3 , retry commit new work
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

type History struct {
	Number     int64          `json:"number"`
	Hash       common.Hash    `json:"hash"`
	Signer     common.Address `json:"signer"`
	Difficulty int64          `json:"difficulty"`
}

type TribeStatus struct {
	Signers      []*Signer        `json:"signers"`
	Volunteers   []common.Address `json:"volunteers"`
	SignerLevel  string           `json:"signerLevel"` // None -> Volunteer -> Signer
	Number       int64            `json:"number"`      // last block.number
	BlackListLen int              `json:"totalSinner"` // length of blacklist
	blackList    []common.Address
	mining       int32 // 1 mining start , 0 mining stop
	nodeKey      *ecdsa.PrivateKey
	tribe        *Tribe
}

type TribeMiner struct {
	Address common.Address `json:"address"`
	Balance *big.Int       `json:"balance"`
	Level   string         `json:"level"` // None 、 Volunteer 、 Signer
}
