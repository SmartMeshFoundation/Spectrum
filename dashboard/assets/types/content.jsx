// @flow

// Copyright 2017 The mesh-chain Authors
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

import type {ChartEntry} from './message';

export type Content = {
    home: Home,
    chain: Chain,
    txpool: TxPool,
    network: Network,
    system: System,
    logs: Logs,
};

export type Home = {
    memory: Array<ChartEntry>,
    traffic: Array<ChartEntry>,
};

export type Chain = {
    /* TODO (kurkomisi) */
};

export type TxPool = {
    /* TODO (kurkomisi) */
};

export type Network = {
    /* TODO (kurkomisi) */
};

export type System = {
    /* TODO (kurkomisi) */
};

export type Logs = {
    log: Array<string>,
};
