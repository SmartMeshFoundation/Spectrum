// Copyright 2016 The mesh-chain Authors
// This file is part of the mesh-chain library.
//
// The mesh-chain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The mesh-chain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the mesh-chain library. If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/contracts/ens"
	"github.com/MeshBoxTech/mesh-chain/crypto"
	"github.com/MeshBoxTech/mesh-chain/log"
	"github.com/MeshBoxTech/mesh-chain/node"
	"github.com/MeshBoxTech/mesh-chain/swarm/network"
	"github.com/MeshBoxTech/mesh-chain/swarm/services/swap"
	"github.com/MeshBoxTech/mesh-chain/swarm/storage"
)

const (
	DefaultHTTPListenAddr = "127.0.0.1"
	DefaultHTTPPort       = "8500"
)

// separate bzz directories
// allow several bzz nodes running in parallel
type Config struct {
	// serialised/persisted fields
	*storage.StoreParams
	*storage.ChunkerParams
	*network.HiveParams
	Swap *swap.SwapParams
	*network.SyncParams
	Contract    common.Address
	EnsRoot     common.Address
	EnsApi      string
	Path        string
	ListenAddr  string
	Port        string
	PublicKey   string
	BzzKey      string
	NetworkId   uint64
	SwapEnabled bool
	SyncEnabled bool
	SwapApi     string
	Cors        string
	BzzAccount  string
	BootNodes   string
}

//create a default config with all parameters to set to defaults
func NewDefaultConfig() (self *Config) {

	self = &Config{
		StoreParams:   storage.NewDefaultStoreParams(),
		ChunkerParams: storage.NewChunkerParams(),
		HiveParams:    network.NewDefaultHiveParams(),
		SyncParams:    network.NewDefaultSyncParams(),
		Swap:          swap.NewDefaultSwapParams(),
		ListenAddr:    DefaultHTTPListenAddr,
		Port:          DefaultHTTPPort,
		Path:          node.DefaultDataDir(),
		EnsApi:        node.DefaultIPCEndpoint("smc"),
		EnsRoot:       ens.TestNetAddress,
		NetworkId:     network.NetworkId,
		SwapEnabled:   false,
		SyncEnabled:   true,
		SwapApi:       "",
		BootNodes:     "",
	}

	return
}

//some config params need to be initialized after the complete
//config building phase is completed (e.g. due to overriding flags)
func (self *Config) Init(prvKey *ecdsa.PrivateKey) {

	address := crypto.PubkeyToAddress(prvKey.PublicKey)
	self.Path = filepath.Join(self.Path, "bzz-"+common.Bytes2Hex(address.Bytes()))
	err := os.MkdirAll(self.Path, os.ModePerm)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating root swarm data directory: %v", err))
		return
	}

	pubkey := crypto.FromECDSAPub(&prvKey.PublicKey)
	pubkeyhex := common.ToHex(pubkey)
	keyhex := crypto.Keccak256Hash(pubkey).Hex()

	self.PublicKey = pubkeyhex
	self.BzzKey = keyhex

	self.Swap.Init(self.Contract, prvKey)
	self.SyncParams.Init(self.Path)
	self.HiveParams.Init(self.Path)
	self.StoreParams.Init(self.Path)
}
