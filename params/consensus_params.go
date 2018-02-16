package params

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
)

const (
	ChiefAddress string = "0x9ec55c1dafd4a487e41da33e344aef86da41ab82" //chief contract address for consensus of tribe
)

type Mbox struct {
	Method string
	Rtn    chan MBoxSuccess
	Params map[string]interface{}
}

type MBoxSuccess struct {
	Success bool
	Entity  interface{}
}

// for chief
var MboxChan chan Mbox = make(chan Mbox, 32)

func SendToMsgBox(method string, rtn chan MBoxSuccess) {
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	MboxChan <- m
}

// chief 合约中 GetStatus 的返回值
type ChiefStatus struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}

