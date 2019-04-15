package params

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"math/big"
	"os"
	"sort"
	"sync"
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
	StatuteService   = make(chan Mbox, 384)
	//close at tribe.init
	TribeReadyForAcceptTxs = make(chan struct{})
	InitTribe              = make(chan struct{})
	InitMeshbox            = make(chan struct{})
	InitAnmap              = make(chan struct{})
	//close at tribeService
	InitTribeStatus               = make(chan struct{})
	chiefInfoList   ChiefInfoList = nil
	// added by cai.zhihong
	// ChiefTxGas = big.NewInt(400000)
	//abiCache *lru.Cache = nil
	chiefContractCodeCache = new(sync.Map)
)

func GetMinMinerBalance() (mb *big.Int) {
	if IsTestnet() {
		mb = TestnetChainConfig.MinMinerBalance
	} else {
		mb = MainnetChainConfig.MinMinerBalance
	}
	return
}

func ChainID() (id *big.Int) {
	if IsTestnet() {
		id = TestnetChainConfig.ChainId
	} else {
		id = MainnetChainConfig.ChainId
	}
	return
}

func AnmapInfo(num *big.Int) (n *big.Int, addr common.Address) {
	if IsTestnet() {
		n = TestnetChainConfig.AnmapBlock
		if n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = TestnetChainConfig.AnmapAddress
		}
	} else {
		n = MainnetChainConfig.AnmapBlock
		if n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = MainnetChainConfig.AnmapAddress
		}
	}
	return
}

func MeshboxInfo(num *big.Int) (n *big.Int, addr common.Address) {
	if IsTestnet() {
		n = TestnetChainConfig.MeshboxBlock
		if n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = TestnetChainConfig.MeshboxAddress
		}
	} else {
		n = MainnetChainConfig.MeshboxBlock
		if n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = MainnetChainConfig.MeshboxAddress
		}
	}
	return
}

// if input num less then nr001block ,enable new role for chief.tx's gaspool
func IsSIP001Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP001Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP001Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// new_rule_002 to change block period
// SIP002Block must big then zero
func IsSIP002Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP002Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.SIP002Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP002Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP002Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 18-09-13 : incompatible HomesteadSigner begin at this number
func IsSIP003Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP003Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.SIP003Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP003Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP003Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 19-03-27 : for smc-0.6.0
// https://github.com/SmartMeshFoundation/Spectrum/wiki/%5BChinese%5D-v0.6.0-Standard
func IsSIP004Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP004Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.SIP004Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP004Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP004Block.Cmp(num) <= 0 {
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
	if len(data) < 4 {
		return false
	} else {
		if bytes.Equal(data[:4], []byte{28, 27, 135, 114}) {
			return true
		}
	}
	return false
}

func AnmapBindInfo(addr common.Address, blockHash common.Hash) (from common.Address, nodeids []common.Address, err error) {
	select {
	case <-InitAnmap:
		rtn := make(chan MBoxSuccess)
		m := Mbox{
			Method: "bindInfo",
			Rtn:    rtn,
		}
		emptyHash := common.Hash{} == blockHash
		if emptyHash {
			m.Params = map[string]interface{}{"addr": addr}
		} else {
			m.Params = map[string]interface{}{"addr": addr, "hash": blockHash}
		}
		StatuteService <- m
		success := <-rtn
		if success.Success {
			m := success.Entity.(map[string]interface{})
			from = m["from"].(common.Address)
			nodeids = m["nodeids"].([]common.Address)
		} else {
			err = success.Entity.(error)
		}
	default:
		err = errors.New("anmap_not_ready")
	}
	return
}

func AnmapBind(from, nodeid common.Address, sigHex string) (string, error) {
	select {
	case <-InitAnmap:
		rtn := make(chan MBoxSuccess)
		m := Mbox{
			Method: "bind",
			Rtn:    rtn,
		}
		m.Params = map[string]interface{}{"from": from, "nodeid": nodeid, "sigHex": sigHex}
		StatuteService <- m
		success := <-rtn
		if success.Success {
			return success.Entity.(string), nil
		} else {
			return "", success.Entity.(error)
		}
	default:
		return "", errors.New("anmap_not_ready")
	}
}

func AnmapUnbind(from, nodeid common.Address, sigHex string) (string, error) {
	select {
	case <-InitAnmap:
		rtn := make(chan MBoxSuccess)
		m := Mbox{
			Method: "unbind",
			Rtn:    rtn,
		}
		m.Params = map[string]interface{}{"from": from, "nodeid": nodeid, "sigHex": sigHex}
		StatuteService <- m
		success := <-rtn
		if success.Success {
			return success.Entity.(string), nil
		} else {
			return "", success.Entity.(error)
		}
	default:
		return "", errors.New("anmap_not_ready")
	}
}

func MeshboxExistAddress(addr common.Address) bool {
	select {
	case <-InitMeshbox:
		rtn := make(chan MBoxSuccess)
		m := Mbox{
			Method: "existAddress",
			Rtn:    rtn,
		}
		m.Params = map[string]interface{}{"addr": addr}
		StatuteService <- m
		success := <-rtn
		if success.Success && success.Entity.(int64) > 0 {
			return true
		}
	}
	return false
}

func SetChiefContractCode(addr common.Address, code []byte) {
	chiefContractCodeCache.Store(addr, code)
}

func GetChiefContractCode(addr common.Address) ([]byte, error) {
	val, ok := chiefContractCodeCache.Load(addr)
	if !ok {
		return nil, errors.New("not_found")
	}
	return val.([]byte), nil
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

func TribeGetStatus(num *big.Int, hash common.Hash) (ChiefStatus, error) {
	rtn := SendToMsgBoxWithHash("GetStatus", hash, num)
	r := <-rtn
	if !r.Success {
		return ChiefStatus{}, r.Entity.(error)
	}
	cs := r.Entity.(ChiefStatus)
	return cs, nil
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
