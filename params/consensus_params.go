package params

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
	"fmt"
	"sort"
)

const (
	// at same account and block number to deploy this contract can be get the same address
	ChiefAddress = "0x9ec55c1dafd4a487e41da33e344aef86da41ab82" //chief contract address for consensus of tribe
)

type ChiefInfo struct {
	StartNumber *big.Int
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

func newChiefInfo(num int64, vsn, addr string) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: big.NewInt(num),
		Version:     vsn,
		Addr:        common.HexToAddress(addr),
	}
}

// for chief
var (
	ChiefBaseBalance = new(big.Int).Mul(big.NewInt(1), big.NewInt(Finney))
	MboxChan         = make(chan Mbox, 32)
	InitTribeStatus  = make(chan struct{})
	chiefAddressList = ChiefInfoList{
		newChiefInfo(2, "0.0.2", "0x9ec55c1dafd4a487e41da33e344aef86da41ab82"),
		newChiefInfo(30000, "0.0.3", "0x01"),
	}
)

func GetChiefAddress(blockNumber *big.Int) *ChiefInfo {
	return getChiefAddress(chiefAddressList,blockNumber)
}
func getChiefAddress(list ChiefInfoList,blockNumber *big.Int) *ChiefInfo {
	// TODO sort once only
	sort.Sort(list)
	for _,c := range list {
		if blockNumber.Int64() >= c.StartNumber.Int64() {
			return c
		}
	}
	return nil
}

func IsChiefAddress(addr common.Address) bool {
	return isChiefAddress(chiefAddressList,addr)
}
func isChiefAddress(list ChiefInfoList,addr common.Address) bool {
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

func SendToMsgBoxWithHash(method string, hash common.Hash) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	if hash != common.HexToHash("0x") {
		m.Params = map[string]interface{}{"hash": hash}
	}
	MboxChan <- m
	return rtn
}

/*
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
*/

func SendToMsgBox(method string) chan MBoxSuccess {
	return SendToMsgBoxWithHash(method, common.HexToHash("0x"))
}

// clone from chief.getStatus return struct
// for return to tribe by channel
type ChiefStatus struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}
