package params

import (
	"math/big"
	"github.com/SmartMeshFoundation/SMChain/common"
	"fmt"
	"crypto/ecdsa"
	"github.com/SmartMeshFoundation/SMChain/crypto"
)

const (
	// at same account and block number to deploy this contract can be get the same address
	ChiefAddress = "0x9ec55c1dafd4a487e41da33e344aef86da41ab82" //chief contract address for consensus of tribe
)

// for chief
var (
	MboxChan        = make(chan Mbox, 32)
	InitTribeStatus = make(chan struct{}, 1)
	ChiefTxNonce    = uint64(0) //先放在这里吧，用来修正 chiefTx.nonce
	coinbaseAddress *common.Address
)

// called on worker.go : commitTransactions
func FixChiefTxNonce(from, to *common.Address, nonce uint64) {
	if coinbaseAddress == nil {
		rtn := SendToMsgBox("GetNodeKey")
		success := <-rtn
		nodeKey := success.Entity.(*ecdsa.PrivateKey)
		a := crypto.PubkeyToAddress(nodeKey.PublicKey)
		coinbaseAddress = &a
	}
	if from != nil && to != nil && *from == *coinbaseAddress && *to == common.HexToAddress(ChiefAddress) {
		ChiefTxNonce = nonce
		fmt.Println("><> ---- FixChiefTxNonce ----> ", ChiefTxNonce)
	}
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
	return SendToMsgBoxWithNumber(method, nil)
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
