package params

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/hashicorp/golang-lru"
	"math/big"
	"os"
	"sort"
	"strings"
)

type ChiefInfo struct {
	StartNumber *big.Int // 0 is nil
	Version     string
	Addr        common.Address
	Abi         string
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

func newChiefInfo(num *big.Int, vsn string, addr common.Address, abi string) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: num,
		Version:     vsn,
		Addr:        addr,
		Abi:         abi,
	}
}

// for chief
// TribeChiefABI is the input ABI used to generate the binding from.
const (
	// copy from chieflib
	TribeChief_0_0_2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	TribeChief_0_0_3ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	TribeChief_0_0_4ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	TribeChief_0_0_5ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	TribeChief_0_0_6ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fillVolunteerForTest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fillSignerForTest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
)

var (
	ChiefBaseBalance = new(big.Int).Mul(big.NewInt(1), big.NewInt(Finney))
	MboxChan         = make(chan Mbox, 32)
	//close at tribe.init
	TribeReadyForAcceptTxs = make(chan struct{})
	InitTribe              = make(chan struct{})
	//close at tribeService
	InitTribeStatus               = make(chan struct{})
	chiefInfoList   ChiefInfoList = nil
	// added by cai.zhihong
	// ChiefTxGas = big.NewInt(400000)
	abiCache *lru.Cache = nil
)

// if input num less then nr001block ,enable new role for chief.tx's gaspool
func IsNR001Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR001Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR001Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// new_rule_002 to change block period
// NR002Block must big then zero
func IsNR002Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR002Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.NR002Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR002Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.NR002Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 18-09-13 : incompatible HomesteadSigner begin at this number
func IsNR003Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR003Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.NR003Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR003Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.NR003Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// startNumber and address must from chain's config
func chiefAddressList() (list ChiefInfoList) {
	if chiefInfoList != nil {
		return chiefInfoList
	}
	if IsTestnet() {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address, TribeChief_0_0_2ABI),
			newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address, TribeChief_0_0_3ABI),
			newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address, TribeChief_0_0_4ABI),
			newChiefInfo(TestnetChainConfig.Chief005Block, "0.0.5", TestnetChainConfig.Chief005Address, TribeChief_0_0_5ABI),
			newChiefInfo(TestnetChainConfig.Chief006Block, "0.0.6", TestnetChainConfig.Chief006Address, TribeChief_0_0_6ABI),
		}
	} else {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(MainnetChainConfig.Chief005Block, "0.0.5", MainnetChainConfig.Chief005Address, TribeChief_0_0_5ABI),
			newChiefInfo(MainnetChainConfig.Chief006Block, "0.0.6", MainnetChainConfig.Chief006Address, TribeChief_0_0_6ABI),
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
	if blockNumber != nil && blockNumber.Cmp(big.NewInt(0)) > 0 {
		return getChiefInfo(chiefAddressList(), blockNumber)
	}
	return nil
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

func IsChiefUpdate(data []byte) bool {
	if abiCache == nil {
		abiCache, _ = lru.New(10)
	}
	if len(data) < 4 {
		return false
	}
	dk := append(data[:4], []byte{0, 0, 0, 0}...)
	if abiCache.Contains(string(dk)) {
		return true
	}
	if len(data) > 4 {
		for _, ci := range chiefAddressList() {
			reader := strings.NewReader(ci.Abi)
			dec := json.NewDecoder(reader)
			var abi abi.ABI
			if err := dec.Decode(&abi); err != nil {
				panic(fmt.Errorf("chief_abi_error : vsn=%s", ci.Version))
			}
			buf, _ := abi.Pack("update", common.Address{})
			bk := append(buf[:4], []byte{0, 0, 0, 0}...)
			abiCache.Add(string(bk), string(bk))
			if bytes.Equal(data[0:4], buf[0:4]) {
				log.Debug("is_chief_update_true", "input", data)
				return true
			}
		}
	}
	return false
}
func IsChiefAddress(addr common.Address) bool {
	return isChiefAddress(chiefAddressList(), addr)
}
func isChiefAddress(list ChiefInfoList, addr common.Address) bool {
	if addr == common.HexToAddress("0x") {
		log.Warn("--> isChiefAddress :: address_not_be_empty", "addr", addr)
		return false
	}
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

func SendToMsgBoxForFilterVolunteer(hash common.Hash, number *big.Int, addr common.Address) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "FilterVolunteer",
		Rtn:    rtn,
	}
	if number == nil || hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash, "number": number, "address": addr}
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
	VolunteerList  []common.Address
	SignerList     []common.Address
	BlackList      []common.Address // append by vsn 0.0.3
	ScoreList      []*big.Int
	NumberList     []*big.Int
	Number         *big.Int
	Epoch          *big.Int
	SignerLimit    *big.Int
	VolunteerLimit *big.Int
	TotalVolunteer *big.Int
}

// clone from chief.getVolunteers return struct
// for return to tribe by channel
// append by vsn 0.0.6
type ChiefVolunteers struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}

func GetIPCPath() string {
	return os.Getenv("IPCPATH")
}

func IsTestnet() bool {
	return os.Getenv("TESTNET") == "1"
}
