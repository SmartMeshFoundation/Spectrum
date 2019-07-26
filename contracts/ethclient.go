package contracts

import (
	"time"

	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
)

var client *ethclient.Client

func GetEthclientInstance() (*ethclient.Client, error) {
	if client == nil {
		var ec *ethclient.Client
		var err error
		for {
			ec, err = ethclient.Dial(params.GetIPCPath())
			if err != nil {
				log.Warn("GetEthclientInstance_fail", "err", err)
				time.Sleep(time.Millisecond * 100)
			} else {
				break
			}
		}

		client = ec
	}
	return client, nil
}
