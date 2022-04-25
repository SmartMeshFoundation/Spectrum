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

// Package tracers is a collection of JavaScript transaction tracers.
package tracers

import (
	"encoding/json"
	"errors"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/core/vm"
)

var all = make(map[string]string)

// tracer retrieves a specific JavaScript tracer by name.
func tracer(name string) (string, bool) {
	if tracer, ok := all[name]; ok {
		return tracer, true
	}
	return "", false
}

// Context contains some contextual infos for a transaction execution that is not
// available from within the EVM object.
type Context struct {
	BlockHash common.Hash // Hash of the block the tx is contained within (zero if dangling tx or call)
	TxIndex   int         // Index of the transaction within a block (zero if dangling tx or call)
	TxHash    common.Hash // Hash of the transaction being traced (zero if dangling call)
}

type Tracer interface {
	vm.EVMLogger
	GetResult() (json.RawMessage, error)
	// Stop terminates execution of the tracer at the first opportune moment.
	Stop(err error)
}

type lookupFunc func(string, *Context) (Tracer, error)

var (
	lookups []lookupFunc
)

// RegisterLookup registers a method as a lookup for tracers, meaning that
// users can invoke a named tracer through that lookup. If 'wildcard' is true,
// then the lookup will be placed last. This is typically meant for interpreted
// engines (js) which can evaluate dynamic user-supplied code.
func RegisterLookup(wildcard bool, lookup lookupFunc) {
	if wildcard {
		lookups = append(lookups, lookup)
	} else {
		lookups = append([]lookupFunc{lookup}, lookups...)
	}
}

// New returns a new instance of a tracer, by iterating through the
// registered lookups.
func New(code string, ctx *Context) (Tracer, error) {
	for _, lookup := range lookups {
		if tracer, err := lookup(code, ctx); err == nil {
			return tracer, nil
		}
	}
	return nil, errors.New("tracer not found")
}
