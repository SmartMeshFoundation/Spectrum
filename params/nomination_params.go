package params

import (
	"crypto/ecdsa"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/crypto"
	"github.com/MeshBoxTech/mesh-chain/log"
)

var exclude = map[common.Address]bool{
	common.HexToAddress("0x3a5fBaC6CA913599C5fde8c1638dB58d01De8A48"): true,
	common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49"): true,
	common.HexToAddress("0xc22D53456ABd14Da347517a4B47ea24866B8E3Ae"): true,
}

// genesis node do not can be nomination
func CanNomination(key *ecdsa.PublicKey) bool {
	addr := crypto.PubkeyToAddress(*key)
	log.Debug("fetchVolunteer_callback.CanNomination ", "addr", addr.Hex(), "exclude", exclude[addr])
	if exclude[addr] {
		return false
	}
	return true
}
