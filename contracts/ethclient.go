package contracts

import (
	"github.com/SmartMeshFoundation/Spectrum/ethclient"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
)

var client *ethclient.Client

func GetEthclientInstance() (*ethclient.Client, error) {
	if client == nil {
		ethclient, err := ethclient.Dial(params.GetIPCPath())
		if err != nil {
			log.Error("GetEthclientInstance_fail", "err", err)
			return nil, err
		}
		client = ethclient
	}
	return client, nil
}
