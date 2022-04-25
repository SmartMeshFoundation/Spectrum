// Copyright 2014 The mesh-chain Authors
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

// Package eth implements the Ethereum protocol.
package eth

import (
	"encoding/hex"
	"testing"
)

func TestMakeExtraData(t *testing.T) {
	extra := make([]byte, 0)
	extra = makeExtraData(extra)
	h := hex.EncodeToString(extra)
	t.Log(h)
	t.Log(extra)
}
