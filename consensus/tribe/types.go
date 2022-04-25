package tribe

import (
	"crypto/ecdsa"
	"errors"
	"github.com/MeshBoxTech/mesh-chain/accounts"
	"github.com/MeshBoxTech/mesh-chain/accounts/abi"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/common/hexutil"
	"github.com/MeshBoxTech/mesh-chain/consensus"
	"github.com/MeshBoxTech/mesh-chain/core/types"
	"github.com/MeshBoxTech/mesh-chain/ethdb"
	"github.com/MeshBoxTech/mesh-chain/params"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
	"sync"
)

const (

	historyLimit = 2048
	wiggleTime           = uint64(5)
	initialBackOffTime   = uint64(5) // second
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
)

var (
	blockPeriod  = uint64(15)                               // Default minimum difference between two consecutive block's timestamps
	extraVanity = 32                                       // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraVrf    = 161                                      // before SIP100 extra format is bytes[extraVanity+extraSeal], after is bytes[extraVrf+extraSeal]
	extraSeal    = 65                                       // Fixed number of extra-data suffix bytes reserved for signer seal
	nonceSync    = hexutil.MustDecode("0xffffffffffffffff") // TODO Reserved to control behavior
	nonceAsync   = hexutil.MustDecode("0x0000000000000000") // TODO Reserved to control behavior
	uncleHash    = types.CalcUncleHash(nil)                 // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
	validatorBytesLength = common.AddressLength
	diffInTurnMain = big.NewInt(3) // Block difficulty for in-turn Main
	diffInTurn     = big.NewInt(2) // Block difficulty for in-turn Sub
	diffNoTurn     = big.NewInt(1) // Block difficulty for out-of-turn Other
	// less than SIP100 <<<<<<<<<<<<<
	diff = int64(6) // SIP100 max diff is 6
	maxValidators = 21
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")
	// errUnauthorizedValidator is returned if a header is signed by a non-authorized entity.
	errUnauthorizedValidator = errors.New("unauthorized validator")


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

	// errInvalidExtraValidators is returned if validator data in extra-data field is invalid.
	errInvalidExtraValidators = errors.New("Invalid extra validators in extra data field")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 - 3 , or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid__difficulty")

	// errExtraValidators is returned if non-sprint-end block contain validator data in
	// their extra-data fields.
	errExtraValidators = errors.New("non-sprint-end block contains extra validator list")

	// errInvalidSpanValidators is returned if a block contains an
	// invalid list of validators (i.e. non divisible by 20 bytes).
	errInvalidSpanValidators = errors.New("invalid validator list on sprint end block")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp       = errors.New("invalid timestamp")
	ErrInvalidTimestampSIP002 = errors.New("invalid timestamp (SIP002)")

	// errUnauthorized is returned if a header is signed by a non-authorized entity.
	errUnauthorized = errors.New("unauthorized")

	// errInvalidCoinbase is returned if the coinbase isn't the validator of the block.
	errInvalidCoinbase = errors.New("Invalid coin base")

	// errInvalidValidatorLen is returned if validators length is zero or bigger than maxValidators.
	errInvalidValidatorsLength = errors.New("Invalid validators length")

	// errWaitTransactions is returned if an empty block is attempted to be sealed
	// on an instant chain (0 second period). It's important to refuse these as the
	// block reward is zero, so an empty block just bloats the chain... fast.
	errWaitTransactions = errors.New("waiting for transactions")
	// for tribe consensus block validator
	ErrTribeNotAllowEmptyTxList                 = errors.New("tribe not allow empty tx list")
	ErrTribeMustContainChiefTx                  = errors.New("tribe must contain chief tx")
	ErrTribeChiefVolunteerLowBalance            = errors.New("tribe chief volunteer low balance")
	ErrTribeChiefVolunteerFail                  = errors.New("tribe chief volunteer check fail")
	ErrTribeChiefTxMustAtPositionZero           = errors.New("tribe chief tx must at postion 0")
	ErrTribeChiefTxSignerAndBlockSignerNotMatch = errors.New("tribe chief update tx signer and block signer not match")
	ErrTribeValdateTxSenderCannotInSignerList   = errors.New("tx sender cannot in signerlist")

	BlockRewardReducedInterval                  = 2102400
	MeshRewardForValidator, _                   = new(big.Int).SetString("100000000000000000000", 10) //Block reward in wei for successfully mining a block
	MeshRewardForPom, _                         = new(big.Int).SetString("10000000000000000000000", 10) //Block reward in wei for successfully mining a block
)

type Tribe struct {
	accman   *accounts.Manager
	config   *params.TribeConfig // Consensus engine configuration parameters
	sigcache *lru.ARCCache       // mapping block.hash -> signer
	//Status   *TribeStatus
	//SealErrorCounter uint32     // less then 3 , retry commit new work
	isInit bool
	lock   sync.Mutex
	stateFn StateFn // Function to get state by state root
	abi map[string]abi.ABI // Interactive with system contracts
	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	db          ethdb.Database         // Database to store and retrieve snapshot checkpoints
	nodeKey   *ecdsa.PrivateKey  //miner
}

type API struct {
	accman *accounts.Manager
	chain  consensus.ChainReader
	tribe  *Tribe
}



type TribeMiner struct {
	Address common.Address `json:"address"`
	Balance *big.Int       `json:"balance"`
}
