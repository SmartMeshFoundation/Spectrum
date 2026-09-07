// Copyright 2015 The Spectrum Authors
// This file is part of the Spectrum library.
//
// The Spectrum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Spectrum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Spectrum library. If not, see <http://www.gnu.org/licenses/>.

package params

var MainnetBootnodes = []string{

	// old bootnode
	"enode://697d18773e08d2ac01c6d6b653ad4f8942bc805118cc09f41508c55e05e5043fb801df7783493e9f87be3b62a1d4da4d753f4b9bbf8921c09ebace0a165ac223@47.105.34.94:44944",

	// --------------------
	// gm
	"enode://722cc5d7c735d472d59513c7170878f0402f7297e17e616dfbaf67efae1176773242683b0aca4aef7450f49c62d9f685a2ec08382ab189fcd69790cab102954e@8.211.125.195:44944",
	// hk
	"enode://a1b98a704a7bf23ecdebdc32488c5f20a41e20446d717e04c20f42bcf479d201e5db388d669df19310977aeba39615d116cb6fc7989679d29629b4fc97cf58df@8.217.96.55:44944",
	// cn : qingdao 1
	"enode://e6c934ad2af2fa35589b9c08cc7df48c497834a067c67de228e2cd3d46952b79dc2b00264f88083145a86812353997668f1d98595ac9eeb3d6bb4c0af58351f4@47.104.107.126:44944",
	// xjp
	"enode://61a5563bd1bf71e2ed7f0daed3c8e597836fcdbebe1e74f07d02c78c53700a193678dd2cff4459f2b3f41ef6e7da3980d332aa4116ba3fa6a9032d4742c4a567@47.84.106.225:44944",
	// us
	"enode://1c0f7ecf9bc730e9feaa723574516358d44ace33bed337340ff3886841886cbf2df0301146162711a75d3c0e59f09f09bcafb66e4ec250340e7072d84520bb54@47.77.213.106:44944",
}

var TestnetBootnodes = []string{
	"enode://c65553f06d82bc26fb6e6126e619f682ae59529f5506ae98269cc60de3c803bf3d8abd99ef7a64ce777f554a4f5d3be70078d319afff59015b7a9d23a5e3c3a1@123.207.146.205:44944",
}

var DevnetBootnodes = []string{}

var RinkebyBootnodes = []string{}

var RinkebyV5Bootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
