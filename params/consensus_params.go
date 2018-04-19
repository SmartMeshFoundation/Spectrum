package params

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"

	"github.com/SmartMeshFoundation/SMChain/common"
)

type ChiefInfo struct {
	StartNumber *big.Int // 0 is nil
	Version     string
	Addr        common.Address
}
type ChiefInfoList []*ChiefInfo

func (p ChiefInfoList) Len() int { return len(p) }
func (p ChiefInfoList) Less(i, j int) bool {
	return p[i].StartNumber.Int64() > p[j].StartNumber.Int64()
}
func (p ChiefInfoList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (self *ChiefInfo) String() string {
	return fmt.Sprintf("start: %d , vsn: %s , addr: %s", self.StartNumber.Int64(), self.Version, self.Addr.Hex())
}

func newChiefInfo(num *big.Int, vsn string, addr common.Address) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: num,
		Version:     vsn,
		Addr:        addr,
	}
}

// for chief
var (
	ChiefBaseBalance               = new(big.Int).Mul(big.NewInt(1), big.NewInt(Finney))
	MboxChan                       = make(chan Mbox, 32)
	InitTribeStatus                = make(chan struct{})
	chiefInfoList    ChiefInfoList = nil
	// added by cai.zhihong
	// ChiefTxGas = big.NewInt(400000)
)

// startNumber and address must from chain's config
func chiefAddressList() (list ChiefInfoList) {
	if chiefInfoList != nil {
		return chiefInfoList
	}
	if os.Getenv("TESTNET") == "1" {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address),
			newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address),
			newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address),
		}
	} else {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(MainnetChainConfig.Chief002Block, "0.0.2", MainnetChainConfig.Chief002Address),
			newChiefInfo(MainnetChainConfig.Chief003Block, "0.0.3", MainnetChainConfig.Chief003Address),
			newChiefInfo(MainnetChainConfig.Chief004Block, "0.0.4", MainnetChainConfig.Chief004Address),
		}
	}
	chiefInfoList = list
	return
}

func GetChiefInfoByVsn(vsn string) *ChiefInfo {
	for _, ci := range chiefAddressList() {
		if ci.StartNumber.Int64() > 0 && ci.Version == vsn {
			return ci
		}
	}
	return nil
}
func GetChiefInfo(blockNumber *big.Int) *ChiefInfo {
	return getChiefInfo(chiefAddressList(), blockNumber)
}
func getChiefInfo(list ChiefInfoList, blockNumber *big.Int) *ChiefInfo {
	// TODO sort once only
	sort.Sort(list)
	for _, c := range list {
		if blockNumber.Int64() >= c.StartNumber.Int64() {
			return c
		}
	}
	return nil
}

// skip verify difficulty on this block number
func IsChiefBlock(blockNumber *big.Int) bool {
	return isChiefBlock(chiefAddressList(), blockNumber)
}

func isChiefBlock(list ChiefInfoList, blockNumber *big.Int) bool {
	for _, ci := range list {
		//log.Info("isChief", "a", ci.StartNumber, "b", blockNumber)
		if ci.StartNumber.Cmp(blockNumber) == 0 {
			return true
		}
	}
	return false
}

func IsChiefAddress(addr common.Address) bool {
	return isChiefAddress(chiefAddressList(), addr)
}
func isChiefAddress(list ChiefInfoList, addr common.Address) bool {
	for _, ci := range list {
		if ci.Addr == addr {
			return true
		}
	}
	return false
}

// chief service message box obj
type Mbox struct {
	Method string
	Rtn    chan MBoxSuccess
	Params map[string]interface{}
}

type MBoxSuccess struct {
	Success bool
	Entity  interface{}
}

// called by chief.GetStatus
func SendToMsgBoxWithHash(method string, hash common.Hash, number *big.Int) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	if number == nil || hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash, "number": number}
	MboxChan <- m
	return rtn
}

// called by chief.Update
func SendToMsgBoxWithNumber(method string, number *big.Int) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	if number != nil {
		m.Params = map[string]interface{}{"number": number}
	}
	MboxChan <- m
	return rtn
}

func SendToMsgBox(method string) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	MboxChan <- m
	return rtn
}

// clone from chief.getStatus return struct
// for return to tribe by channel
type ChiefStatus struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address // append by vsn 0.0.3
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}
