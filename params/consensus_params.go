package params

const (
	ChiefAddress string = "0x941a7fe21434baae92348c728276a39920c8c5ec" //chief contract address for consensus of tribe
)

type Mbox struct {
	Method string
	Rtn    chan interface{}
	Params map[string]interface{}
}
// for chief
var MboxChan chan Mbox = make(chan Mbox,32)
func SendToMsgBox(method string, rtn chan interface{}) {
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	MboxChan <- m
}
