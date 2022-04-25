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

package node

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/MeshBoxTech/mesh-chain/p2p"
	"github.com/MeshBoxTech/mesh-chain/p2p/nat"
)

const (
	DefaultHTTPHost = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort = 18545       // Default TCP port for the HTTP RPC server
	DefaultWSHost   = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort   = 18546       // Default TCP port for the websocket RPC server
)

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{
	DataDir:     DefaultDataDir(),
	HTTPPort:    DefaultHTTPPort,
	HTTPModules: []string{"net", "web3"},
	WSPort:      DefaultWSPort,
	WSModules:   []string{"net", "web3"},
	P2P: p2p.Config{
		ListenAddr:      ":60303",
		DiscoveryV5Addr: ":44945",
		MaxPeers:        25,
		NAT:             nat.Any(),
	},
}

func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "mesh-chain")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "mesh-chain")
		} else {
			return filepath.Join(home, ".mesh-chain")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

//add by liangc : for testnet build ipc path
func TestDataDir() string {
	home := homeDir()
	if home != "" {
		testnet := "testnet"
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "mesh-chain", testnet)
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "mesh-chain", testnet)
		} else {
			return filepath.Join(home, ".mesh-chain", testnet)
		}
	}
	return ""
}

//add by liangc : for testnet build ipc path
func DevDataDir() string {
	home := homeDir()
	if home != "" {
		devnet := "devnet"
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "mesh-chain", devnet)
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "mesh-chain", devnet)
		} else {
			return filepath.Join(home, ".mesh-chain", devnet)
		}
	}
	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
