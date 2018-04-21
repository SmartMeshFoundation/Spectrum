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
	// add by liangc : rollback 176222 in testnet for test rollback
	common.HexToHash("0xab10d12d320ec1d6f6e336b0d9f7e5aa250249618912222d0d902a1d3e3fcf89"): true,
	// add by liangc : rollback 180575 in testnet for fix the gaslimit verify
	common.HexToHash("0x454b5e91b203232e1cf56a49297cfceb07ba945e046536ee1ace63d13d64239f"): true,
}
