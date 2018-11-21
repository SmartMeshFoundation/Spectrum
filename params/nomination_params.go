package params

import (
	"crypto/ecdsa"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/common"
)

var skip = []common.Address{
	common.HexToAddress("0x3a5fBaC6CA913599C5fde8c1638dB58d01De8A48"),
	common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"),
	common.HexToAddress("0xc22D53456ABd14Da347517a4B47ea24866B8E3Ae"),
}

// genesis node do not can be nomination
func CanNomination(key *ecdsa.PublicKey) bool {
	addr := crypto.PubkeyToAddress(*key)
	for _, a := range skip {
		if addr == a {
			return false
		}
	}
	return true
}
