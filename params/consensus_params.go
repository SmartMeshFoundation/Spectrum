package params

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
)

const (
	// at same account and block number to deploy this contract can be get the same address
	ChiefAddress = "0x9ec55c1dafd4a487e41da33e344aef86da41ab82" //chief contract address for consensus of tribe
)

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

// for chief
var (
	MboxChan = make(chan Mbox, 32)
	InitTribeStatus = make(chan struct{},1)
)


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
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}

