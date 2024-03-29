package params

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"sync"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

const (
	POC_METHOD_DEPOSIT          = "poc_deposit"
	POC_METHOD_START            = "poc_start"
	POC_METHOD_STOP             = "poc_stop"
	POC_METHOD_WITHDRAW         = "poc_withdraw"
	POC_METHOD_WITHDRAW_SURPLUS = "poc_withdrawsurplus"
	POC_METHOD_GET_STATUS       = "poc_getall"

	Chief100Update = "SIP100Update"
)

type ChiefInfo struct {
	StartNumber, PocStartNumber *big.Int // 0 is nil
	Version                     string
	Addr, PocAddr, BaseAddr     common.Address
}
type ChiefInfoList []*ChiefInfo

func (p ChiefInfoList) Len() int { return len(p) }
func (p ChiefInfoList) Less(i, j int) bool {
	if p[i].StartNumber == nil || p[j].StartNumber == nil {
		return true
	}
	return p[i].StartNumber.Int64() > p[j].StartNumber.Int64()
}
func (p ChiefInfoList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (self *ChiefInfo) String() string {
	return fmt.Sprintf("start: %d , vsn: %s , addr: %s", self.StartNumber.Int64(), self.Version, self.Addr.Hex())
}

func newChiefInfo(num *big.Int, vsn string, addr common.Address) *ChiefInfo {
	empty := common.HexToAddress("0x")
	return newChiefInfoWithPocBase(num, vsn, addr, empty, empty)
}

func newChiefInfoWithPocBase(num *big.Int, vsn string, addr, pocAddr, baseAddr common.Address) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: num,
		Version:     vsn,
		Addr:        addr,
		PocAddr:     pocAddr,
		BaseAddr:    baseAddr,
	}
}

var (
	ChiefBaseBalance = new(big.Int).Mul(big.NewInt(1), big.NewInt(Finney))
	MboxChan         = make(chan Mbox, 32)
	StatuteService   = make(chan Mbox, 384)
	//PocService 用于poc 服务
	PocService = make(chan Mbox, 32)
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

func ChainID() (id *big.Int) {
	if IsTestnet() {
		id = TestnetChainConfig.ChainId
	} else if IsDevnet() {
		id = DevnetChainConfig.ChainId
	} else {
		id = MainnetChainConfig.ChainId
	}
	return
}

func AnmapInfo(num *big.Int, vsn string) (n *big.Int, addr common.Address) {
	if IsTestnet() {
		n = TestnetChainConfig.Anmap001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = TestnetChainConfig.Anmap001Address
		}
	} else if IsDevnet() {
		n = DevnetChainConfig.Anmap001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = DevnetChainConfig.Anmap001Address
		}
	} else {
		n = MainnetChainConfig.Anmap001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			addr = MainnetChainConfig.Anmap001Address
		}
	}
	return
}

func MeshboxVsn(num *big.Int) (string, error) {
	var n *big.Int
	if IsTestnet() {
		n = TestnetChainConfig.Meshbox002Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.2", nil
		}
		n = TestnetChainConfig.Meshbox001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.1", nil
		}

	} else if IsDevnet() {
		n = DevnetChainConfig.Meshbox002Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.2", nil
		}
		n = DevnetChainConfig.Meshbox001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.1", nil
		}
	} else {
		n = MainnetChainConfig.Meshbox002Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.2", nil
		}
		n = MainnetChainConfig.Meshbox001Block
		if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
			return "0.0.1", nil
		}
	}
	return "", errors.New("meshbox_service_not_started")
}

//POCInfo: poc地址
func POCInfo() (addr common.Address) {
	if IsTestnet() {
		addr = TestnetChainConfig.PocAddress
	} else if IsDevnet() {
		addr = DevnetChainConfig.PocAddress
	} else {
		addr = MainnetChainConfig.PocAddress

	}
	return
}

func MeshboxInfo(num *big.Int, vsn string) (n *big.Int, addr common.Address) {
	switch vsn {
	case "0.0.1":
		if IsTestnet() {
			n = TestnetChainConfig.Meshbox001Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = TestnetChainConfig.Meshbox001Address
			}
		} else if IsDevnet() {
			n = DevnetChainConfig.Meshbox001Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = DevnetChainConfig.Meshbox001Address
			}
		} else {
			n = MainnetChainConfig.Meshbox001Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = MainnetChainConfig.Meshbox001Address
			}
		}
	case "0.0.2":
		if IsTestnet() {
			n = TestnetChainConfig.Meshbox002Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = TestnetChainConfig.Meshbox002Address
			}
		} else if IsDevnet() {
			n = DevnetChainConfig.Meshbox002Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = DevnetChainConfig.Meshbox002Address
			}
		} else {
			n = MainnetChainConfig.Meshbox002Block
			if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
				addr = MainnetChainConfig.Meshbox002Address
			}
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
	} else if IsDevnet() {
		if DevnetChainConfig.SIP001Block.Cmp(num) <= 0 {
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
	} else if IsDevnet() {
		if DevnetChainConfig.SIP002Block.Cmp(num) <= 0 {
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
	} else if IsDevnet() {
		if DevnetChainConfig.SIP003Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP003Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP003Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 22-01-29
func IsSIP004Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.Sip004Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.Sip004Block.Cmp(num) <= 0 {
			return true
		}
	} else if IsDevnet() {
		if DevnetChainConfig.Sip004Block.Cmp(big.NewInt(0)) > 0 && DevnetChainConfig.Sip004Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.Sip004Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.Sip004Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 22-01-29
func EqualSIP004Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.Sip004Block.Cmp(num) == 0 {
			return true
		}
	} else if IsDevnet() {
		if DevnetChainConfig.Sip004Block.Cmp(num) == 0 {
			return true
		}
	} else {
		if MainnetChainConfig.Sip004Block.Cmp(num) == 0 {
			return true
		}
	}
	return false
}

// add by liangc : 19-05-31 : for smc-1.0.0
// may be discard
func IsSIP100Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.Chief100Block != nil && TestnetChainConfig.Chief100Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.Chief100Block.Cmp(num) <= 0 {
			return true
		}
	} else if IsDevnet() {
		if DevnetChainConfig.Chief100Block != nil && DevnetChainConfig.Chief100Block.Cmp(big.NewInt(0)) > 0 && DevnetChainConfig.Chief100Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.Chief100Block != nil && MainnetChainConfig.Chief100Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.Chief100Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

func IsReadyMeshbox(num *big.Int) bool {
	n := MainnetChainConfig.Meshbox002Block
	if IsTestnet() {
		n = TestnetChainConfig.Meshbox002Block
	} else if IsDevnet() {
		n = DevnetChainConfig.Meshbox002Block
	}
	if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
		return true
	}

	n = MainnetChainConfig.Meshbox001Block
	if IsTestnet() {
		n = TestnetChainConfig.Meshbox001Block
	} else if IsDevnet() {
		n = DevnetChainConfig.Meshbox001Block
	}
	if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
		return true
	}
	return false
}

func IsReadyAnmap(num *big.Int) bool {
	n := MainnetChainConfig.Anmap001Block
	if IsTestnet() {
		n = TestnetChainConfig.Anmap001Block
	} else if IsDevnet() {
		n = DevnetChainConfig.Anmap001Block
	}
	if n != nil && n.Cmp(big.NewInt(0)) > 0 && n.Cmp(num) <= 0 {
		return true
	}
	return false
}

// startNumber and address must from chain's config
// 针对老的分叉版本在检测时跳过，从1.0.0版本开始不再跳过，所有以后版本本列表不再增加新的项
func beforechief100AddressList() (list ChiefInfoList) {
	if chiefInfoList != nil {
		return chiefInfoList
	}
	if IsTestnet() {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address),
			/*
				newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address),
				newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address),
				newChiefInfo(TestnetChainConfig.Chief005Block, "0.0.5", TestnetChainConfig.Chief005Address),
				newChiefInfo(TestnetChainConfig.Chief006Block, "0.0.6", TestnetChainConfig.Chief006Address),
			*/
		}
	} else if IsDevnet() {
		list = ChiefInfoList{
			newChiefInfo(DevnetChainConfig.Chief007Block, "0.0.7", DevnetChainConfig.Chief007Address),
		}
	} else {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(MainnetChainConfig.Chief005Block, "0.0.5", MainnetChainConfig.Chief005Address),
			newChiefInfo(MainnetChainConfig.Chief006Block, "0.0.6", MainnetChainConfig.Chief006Address),
		}
	}
	chiefInfoList = list
	return
}

// startNumber and address must from chain's config
func chiefAddressList() (list ChiefInfoList) {
	if chiefInfoList != nil {
		return chiefInfoList
	}
	if IsTestnet() {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(TestnetChainConfig.Chief002Block, "0.0.2", TestnetChainConfig.Chief002Address),
			/*
				newChiefInfo(TestnetChainConfig.Chief003Block, "0.0.3", TestnetChainConfig.Chief003Address),
				newChiefInfo(TestnetChainConfig.Chief004Block, "0.0.4", TestnetChainConfig.Chief004Address),
				newChiefInfo(TestnetChainConfig.Chief005Block, "0.0.5", TestnetChainConfig.Chief005Address),
				newChiefInfo(TestnetChainConfig.Chief006Block, "0.0.6", TestnetChainConfig.Chief006Address),
			*/
			newChiefInfoWithPocBase(TestnetChainConfig.Chief100Block, "1.0.0", TestnetChainConfig.Chief100Address, TestnetChainConfig.PocAddress, TestnetChainConfig.ChiefBaseAddress),
		}
	} else if IsDevnet() {
		list = ChiefInfoList{
			newChiefInfo(DevnetChainConfig.Chief007Block, "0.0.7", DevnetChainConfig.Chief007Address),
			newChiefInfoWithPocBase(DevnetChainConfig.Chief100Block, "1.0.0", DevnetChainConfig.Chief100Address, DevnetChainConfig.PocAddress, DevnetChainConfig.ChiefBaseAddress),
		}
	} else {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(MainnetChainConfig.Chief005Block, "0.0.5", MainnetChainConfig.Chief005Address),
			newChiefInfo(MainnetChainConfig.Chief006Block, "0.0.6", MainnetChainConfig.Chief006Address),
			newChiefInfoWithPocBase(MainnetChainConfig.Chief100Block, "1.0.0", MainnetChainConfig.Chief100Address, MainnetChainConfig.PocAddress, MainnetChainConfig.ChiefBaseAddress),
		}
	}
	chiefInfoList = list
	return
}

func GetChiefInfoByVsn(vsn string) *ChiefInfo {
	for _, ci := range chiefAddressList() {
		if ci.StartNumber != nil && ci.StartNumber.Int64() > 0 && ci.Version == vsn {
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
		if c.StartNumber != nil && blockNumber.Int64() >= c.StartNumber.Int64() {
			return c
		}
	}
	return nil
}

// skip verify difficulty on this old hardfork block number
func IsBeforeChief100block(blockNumber *big.Int) bool {
	//return isChiefBlock(oldchiefAddressList(), blockNumber)
	for _, ci := range beforechief100AddressList() {
		//log.Info("isChief", "a", ci.StartNumber, "b", blockNumber)
		if ci.StartNumber != nil && ci.StartNumber.Cmp(blockNumber) == 0 {
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

//PocDeposit poc质押操作
func PocDeposit(from common.Address, sigHex string) (string, error) {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "poc_deposit",
		Rtn:    rtn,
	}
	m.Params = map[string]interface{}{"from": from, "sigHex": sigHex}
	StatuteService <- m
	success := <-rtn
	if success.Success {
		return success.Entity.(string), nil
	} else {
		return "", success.Entity.(error)
	}
}
func pocHelper(method string, from, nodeID common.Address) (string, error) {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	m.Params = map[string]interface{}{"from": from, "nodeid": nodeID}
	StatuteService <- m
	success := <-rtn
	if success.Success {
		return success.Entity.(string), nil
	} else {
		return "", success.Entity.(error)
	}
}

//PocStart 因为错过出块被拉黑以后,在一个epoch之后需要手工恢复出块资格
func PocStart(from, nodeID common.Address) (string, error) {
	return pocHelper(POC_METHOD_START, from, nodeID)
}

//PocStop 因为不想参与挖矿,准备撤回抵押
func PocStop(from, nodeID common.Address) (string, error) {
	return pocHelper(POC_METHOD_STOP, from, nodeID)
}

//PocWithdraw 在停止PocStop两周后,可以从poc合约中撤回押金到自己账户
func PocWithdraw(from, nodeID common.Address) (string, error) {
	return pocHelper(POC_METHOD_WITHDRAW, from, nodeID)
}

//PocWithdrawSurplus  因为手工调用Poc Withdraw合约,一次性抵押过多,后续可以选择撤回多余的抵押
func PocWithdrawSurplus(from, nodeID common.Address) (string, error) {
	return pocHelper(POC_METHOD_WITHDRAW_SURPLUS, from, nodeID)
}

//PocStatus Poc状态
type PocStatus struct {
	MinerList       []common.Address
	AmountList      []*big.Int
	BlockList       []*big.Int
	OwnerList       []common.Address
	BlackStatusList []*big.Int
}

//PocGetAll 获取poc status
func PocGetAll(hash common.Hash, number *big.Int) (ps *PocStatus, err error) {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: POC_METHOD_GET_STATUS,
		Rtn:    rtn,
	}
	if number == nil || hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash, "number": number}
	StatuteService <- m
	success := <-rtn
	if success.Success {
		return success.Entity.(*PocStatus), nil
	} else {
		return nil, success.Entity.(error)
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
	default:
		return false
	}
	return false
}

/*
func GetVRFByHash(hash common.Hash) ([]byte, error) {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "GetVRF",
		Rtn:    rtn,
	}
	if hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash}
	MboxChan <- m

	success := <-rtn
	if success.Success {
		return success.Entity.([]byte), nil
	}
	return nil, success.Entity.(error)
}*/

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

var chiefCalledMap = map[string]string{
	TestnetChainConfig.Chief100Address.Hex():  TestnetChainConfig.ChiefBaseAddress.Hex(),
	TestnetChainConfig.ChiefBaseAddress.Hex(): TestnetChainConfig.PocAddress.Hex(),

	MainnetChainConfig.Chief100Address.Hex():  MainnetChainConfig.ChiefBaseAddress.Hex(),
	MainnetChainConfig.ChiefBaseAddress.Hex(): MainnetChainConfig.PocAddress.Hex(),

	DevnetChainConfig.Chief100Address.Hex():  DevnetChainConfig.ChiefBaseAddress.Hex(),
	DevnetChainConfig.ChiefBaseAddress.Hex(): DevnetChainConfig.PocAddress.Hex(),
}

func IsChiefCalled(from, to common.Address) (yes bool) {
	if t, ok := chiefCalledMap[from.Hex()]; ok && t == to.Hex() {
		yes = true
	}
	//fmt.Println("params.IsChiefCalled", from.Hex(), "->", to.Hex(), yes)
	return
}

func IsChiefAddress(addr common.Address) (yes bool) {
	yes = isChiefAddress(chiefAddressList(), addr)
	//t := stack.Trace()
	return
}

func isChiefAddress(list ChiefInfoList, addr common.Address) bool {
	if addr == common.HexToAddress("0x") {
		return false
	}
	for _, ci := range list {
		if ci.Addr == addr {
			return true
		}
	}
	return false
}

func IsChiefAddressOnBlock(number *big.Int, addr common.Address) (yes bool) {
	yes = isChiefAddressOnBlock(number, chiefAddressList(), addr)
	//t := stack.Trace()
	return
}
func isChiefAddressOnBlock(number *big.Int, list ChiefInfoList, addr common.Address) bool {
	if addr == common.HexToAddress("0x") {
		return false
	}
	if IsSIP100Block(number) { //sip100以后需要验证to地址必须和新的共识合约相同
		for _, ci := range list {
			if ci.Addr == addr && ci.Version == "1.0.0" {
				return true
			}
		}
	} else {
		for _, ci := range list {
			if ci.Addr == addr {
				return true
			}
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

/*
依据vrf,依据上一块的状态来选取下一轮的出块地址
*/
func Chief100GetNextRoundSigner(hash common.Hash, number *big.Int, vrf *big.Int) common.Address {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: Chief100Update,
		Rtn:    rtn,
	}
	m.Params = map[string]interface{}{
		"hash":   hash,
		"number": number,
		"vrfn":   vrf,
	}
	MboxChan <- m
	success := <-rtn
	return success.Entity.(common.Address) //必须成功,不成功就返回空地址

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

func VerifyMiner(parenthash common.Hash, addr common.Address, vrfn []byte) bool {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "VerifyMiner",
		Rtn:    rtn,
	}
	if parenthash == common.HexToHash("0x") {
		panic(errors.New("hash can not nil"))
	}
	if vrfn == nil {
		panic(errors.New("vrfn can not nil"))
	}
	m.Params = map[string]interface{}{"parenthash": parenthash, "addr": addr, "vrfn": vrfn}
	MboxChan <- m
	success := <-rtn
	if success.Success {
		return success.Entity.(bool)
	}
	return false
}

// clone from chief.getStatus return struct
// for return to tribe by channel
type ChiefStatus struct {
	LeaderList     []common.Address
	VolunteerList  []common.Address
	SignerList     []common.Address
	BlackList      []common.Address // append by vsn 0.0.3
	ScoreList      []*big.Int
	NumberList     []*big.Int
	Number         *big.Int
	Epoch          *big.Int
	LeaderLimit    *big.Int
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

func IsDevnet() bool {
	return os.Getenv("DEVNET") == "1"
}

func TribePeriod() uint64 {
	if IsTestnet() && TestnetChainConfig.Tribe.Period > 0 {
		return TestnetChainConfig.Tribe.Period
	} else if IsDevnet() && DevnetChainConfig.Tribe.Period > 0 {
		return DevnetChainConfig.Tribe.Period
	} else if MainnetChainConfig.Tribe.Period > 0 {
		return MainnetChainConfig.Tribe.Period
	}
	return 14
}
