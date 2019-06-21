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
	StartNumber, PocStartNumber *big.Int // 0 is nil
	Version                     string
	Addr, PocAddr               common.Address
	Abi                         string
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
	return newChiefInfoWithPoc(num, nil, vsn, addr, common.HexToAddress("0x"), abi)
}

func newChiefInfoWithPoc(num, pocNum *big.Int, vsn string, addr, pocAddr common.Address, abi string) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: num,
		Version:     vsn,
		Addr:        addr,
		Abi:         abi,
		PocAddr:     pocAddr,
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
	TribeChief_0_0_7ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"},{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"},{\"name\":\"_volunteerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	TribeChief_1_0_0ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"blMap\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_blackMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_v2s\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_clean_all_signer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLeader\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"signersMap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_blackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_volunteerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"volunteerMap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeBlackMember\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_set_leader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_clean_volunteer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_signerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"pushBlackMember\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"pushVolunteer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"baseAddress\",\"type\":\"address\"},{\"name\":\"pocAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
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
	} else if IsDevnet() {
		mb = DevnetChainConfig.MinMinerBalance
	} else {
		mb = MainnetChainConfig.MinMinerBalance
	}
	return
}

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

// add by liangc : 19-03-27 : for smc-0.6.0
// https://github.com/SmartMeshFoundation/Spectrum/wiki/%5BChinese%5D-v0.6.0-Standard
func IsSIP004Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP004Block != nil && TestnetChainConfig.SIP004Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.SIP004Block.Cmp(num) <= 0 {
			return true
		}
	} else if IsDevnet() {
		if DevnetChainConfig.SIP004Block != nil && DevnetChainConfig.SIP004Block.Cmp(big.NewInt(0)) > 0 && DevnetChainConfig.SIP004Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP004Block != nil && MainnetChainConfig.SIP004Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP004Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 19-05-31 : for smc-1.0.0
// may be discard
func IsSIP005Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.SIP005Block != nil && TestnetChainConfig.SIP005Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.SIP005Block.Cmp(num) <= 0 {
			return true
		}
	} else if IsDevnet() {
		if DevnetChainConfig.SIP005Block != nil && DevnetChainConfig.SIP005Block.Cmp(big.NewInt(0)) > 0 && DevnetChainConfig.SIP005Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.SIP005Block != nil && MainnetChainConfig.SIP005Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.SIP005Block.Cmp(num) <= 0 {
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
	} else if IsDevnet() {
		list = ChiefInfoList{
			newChiefInfo(DevnetChainConfig.Chief007Block, "0.0.7", DevnetChainConfig.Chief007Address, TribeChief_0_0_7ABI),
			newChiefInfoWithPoc(DevnetChainConfig.Chief100Block, DevnetChainConfig.PocBlock, "1.0.0", DevnetChainConfig.Chief100Address, DevnetChainConfig.PocAddress, TribeChief_1_0_0ABI),
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
	if method == "Update" {
		log.Debug("TODO <<SendToMsgBoxWithNumber>> begin", "num", number)
	}
	MboxChan <- m
	if method == "Update" {
		log.Debug("TODO <<SendToMsgBoxWithNumber>> end", "num", number)
	}
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

func IsDevnet() bool {
	return os.Getenv("DEVNET") == "1"
}
