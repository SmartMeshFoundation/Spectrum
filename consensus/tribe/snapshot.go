// Copyright 2017 The go-ethereum Authors
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

package tribe

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/consensus"
	"github.com/MeshBoxTech/mesh-chain/core/types"
	"github.com/MeshBoxTech/mesh-chain/ethdb"
	"github.com/MeshBoxTech/mesh-chain/params"
)
// Snapshot is the state of the authorization voting at a given point in time.
type Snapshot struct {
	config   *params.TribeConfig // Consensus engine parameters to fine tune behavior
	Number     uint64                      `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                 `json:"hash"`       // Block hash where the snapshot was created
	Validators map[common.Address]struct{} `json:"validators"` // Set of authorized validators at this moment
}

// validatorsAscending implements the sort interface to allow sorting a list of addresses
type validatorsAscending []common.Address

func (s validatorsAscending) Len() int           { return len(s) }
func (s validatorsAscending) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }
func (s validatorsAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent validators, so only ever use if for
// the genesis block.
func newSnapshot(config *params.TribeConfig, number uint64, hash common.Hash, validators []common.Address) *Snapshot {
	snap := &Snapshot{
		config:     config,
		Number:     number,
		Hash:       hash,
		Validators: make(map[common.Address]struct{}),
	}
	for _, validator := range validators {
		snap.Validators[validator] = struct{}{}
	}
	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(config *params.TribeConfig, db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	blob, err := db.Get(append([]byte("tribe-"), hash[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	snap.config = config

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("tribe-"), s.Hash[:]...), blob)
}

// copy creates a deep copy of the snapshot, though not the individual votes.
func (s *Snapshot) copy() *Snapshot {
	cpy := &Snapshot{
		config:     s.config,
		Number:     s.Number,
		Hash:       s.Hash,
		Validators: make(map[common.Address]struct{}),
	}
	for validator := range s.Validators {
		cpy.Validators[validator] = struct{}{}
	}

	return cpy
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot) apply(headers []*types.Header, chain consensus.ChainReader, parents []*types.Header,t *Tribe) (*Snapshot, error) {
	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}
	// Sanity check that the headers can be applied
	for i := 0; i < len(headers)-1; i++ {
		if headers[i+1].Number.Uint64() != headers[i].Number.Uint64()+1 {
			return nil, errInvalidVotingChain
		}
	}
	if headers[0].Number.Uint64() != s.Number+1 {
		return nil, errInvalidVotingChain
	}
	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	for _, header := range headers {
		// Remove any votes on checkpoint blocks
		number := header.Number.Uint64()
		// Resolve the authorization key and check against validators
		validator, err := ecrecover(header, t)
		if err != nil {
			return nil, err
		}
		if _, ok := snap.Validators[validator]; !ok {
			return nil, errUnauthorizedValidator
		}

		// update validators at the first block at epoch
		if number > 0 && number%s.config.Epoch == 0 {
			checkpointHeader := header
			// get validators from headers and use that for new validator set
			validators := make([]common.Address, (len(checkpointHeader.Extra)-extraVanity-extraSeal-extraVrf)/common.AddressLength)
			for i := 0; i < len(validators); i++ {
				copy(validators[i][:], checkpointHeader.Extra[extraVanity+extraVrf+i*common.AddressLength:])
			}

			newValidators := make(map[common.Address]struct{})
			for _, validator := range validators {
				newValidators[validator] = struct{}{}
			}

			snap.Validators = newValidators
		}
	}

	snap.Number += uint64(len(headers))
	snap.Hash = headers[len(headers)-1].Hash()

	return snap, nil
}

// validators retrieves the list of authorized validators in ascending order.
func (s *Snapshot) validators() []common.Address {
	sigs := make([]common.Address, 0, len(s.Validators))
	for sig := range s.Validators {
		sigs = append(sigs, sig)
	}
	sort.Sort(validatorsAscending(sigs))
	return sigs
}

// inturn returns if a validator at a given block height is in-turn or not.
func (s *Snapshot) inturn(validator common.Address) bool {
	validators := s.validators()
	offset := (s.Number + 1) % uint64(len(validators))
	return validators[offset] == validator
}

func (s *Snapshot) indexOfVal(validator common.Address) int {
	validators := s.validators()
	for idx, val := range validators {
		if val == validator {
			return idx
		}
	}
	return -1
}
