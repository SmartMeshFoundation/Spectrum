// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

var MainnetBootnodes = []string{
	// cn : qingdao 1
	"enode://697d18773e08d2ac01c6d6b653ad4f8942bc805118cc09f41508c55e05e5043fb801df7783493e9f87be3b62a1d4da4d753f4b9bbf8921c09ebace0a165ac223@47.105.34.94:44944",
	// cn : shanghai
	"enode://a69b1cb395ae4d0520d8c82df448ce17daf8b4ce175a0fe892b8a7544abe55a3c3297e67ac7be5ed5d99c2d0b56e6cc3c1b4b33fc62ade55772d5d240ac0ab39@47.100.215.147:44944",
}

var TestnetBootnodes = []string{
	"enode://c65553f06d82bc26fb6e6126e619f682ae59529f5506ae98269cc60de3c803bf3d8abd99ef7a64ce777f554a4f5d3be70078d319afff59015b7a9d23a5e3c3a1@123.207.146.205:44944",
}

var RinkebyBootnodes = []string{
}

var RinkebyV5Bootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
