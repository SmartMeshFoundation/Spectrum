// Copyright 2015 The mesh-chain Authors
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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/MeshBoxTech/mesh-chain/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("eth/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("eth/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("eth/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("eth/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("eth/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("eth/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("eth/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("eth/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("eth/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("eth/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("eth/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("eth/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("eth/downloader/states/in")
	stateDropMeter = metrics.NewMeter("eth/downloader/states/drop")
)
