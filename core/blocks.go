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

package core

import "github.com/SmartMeshFoundation/Spectrum/common"

// BadHashes represent a set of manually tracked bad hashes (usually hard forks)
var BadHashes = map[common.Hash]bool{
	// add by liangc : testnet : rollback 176222 in testnet for test rollback
	common.HexToHash("0xab10d12d320ec1d6f6e336b0d9f7e5aa250249618912222d0d902a1d3e3fcf89"): true,
	// add by liangc : testnet : 176222
	common.HexToHash("0x9423939f91a9edb51f68ec31cae3c46ab3dd853208e4266f403817bb514395dc"): true,
	// add by liangc : testnet : 176222 : fix gas error
	common.HexToHash("0xc01cdb386fa7c94b4830f97eeb51b1e3d9cdb6f890e13dbd1ebc902bdbdea243"): true,
	// add by liangc : testnet : rollback 180572 in testnet for fix the gaslimit verify
	common.HexToHash("0x4466da3caf8f797937d095f4da36635880ce1ebe088c0f0f78fb4dae66e73eec"): true,

	// add by liangc : testnet : rollback 525233 in testnet for test block time : 14 18 22
	common.HexToHash("0x49109562155ac528c52873dcede7bfa2666e5c7bc21f95d6f742baca7a4abbc7"): true,
	common.HexToHash("0x8857ce7c4e17c5ff1605a5da8bdd172b669e6eaa8a5b2ebf7d4b574f7e36ec97"): true,
	common.HexToHash("0xf127c1bb9061264744cc62654647562a058f510231d5203ff21f6f9299407bcf"): true,

	// add by liangc : testnet : rollback 595708 in testnet for test and update chief 0.0.6
	common.HexToHash("0x24f468cb2f0a77add1705a341277eadea9b64700981f6e0d94dc688c5dfbb9f2"): true,
}
