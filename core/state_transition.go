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

package core

import (
	"errors"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/core/vm"
	"github.com/MeshBoxTech/mesh-chain/log"
	"github.com/MeshBoxTech/mesh-chain/params"
	"math/big"
)

var (
	Big0                         = big.NewInt(0)
	errInsufficientBalanceForGas = errors.New("insufficient balance to pay for gas")
)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay gas
3) Create a new state object if the recipient is \0*32
4) Value transfer
== If contract creation ==
  4a) Attempt to run transaction data
  4b) If valid, use result as code for the new state object
== end ==
5) Run Script section
6) Derive new state root
*/
type StateTransition struct {
	gp          *GasPool
	msg         Message
	gas         uint64
	gasPrice    *big.Int
	initialGas  *big.Int
	value       *big.Int
	data        []byte
	state       vm.StateDB
	evm         *vm.EVM
	blockNumber *big.Int
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	//FromFrontier() (common.Address, error)
	To() *common.Address

	GasPrice() *big.Int
	Gas() *big.Int
	Value() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}

// IntrinsicGas computes the 'intrinsic gas' for a message
// with the given data.
//
// TODO convert to uint64
func IntrinsicGas(data []byte, contractCreation, homestead bool) *big.Int {
	igas := new(big.Int)
	if contractCreation && homestead {
		igas.SetUint64(params.TxGasContractCreation)
	} else {
		igas.SetUint64(params.TxGas)
	}
	if len(data) > 0 {
		var nz int64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		m := big.NewInt(nz)
		m.Mul(m, new(big.Int).SetUint64(params.TxDataNonZeroGas))
		igas.Add(igas, m)
		m.SetInt64(int64(len(data)) - nz)
		m.Mul(m, new(big.Int).SetUint64(params.TxDataZeroGas))
		igas.Add(igas, m)
	}
	return igas
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, msg Message, gp *GasPool, blockNumber *big.Int) *StateTransition {
	return &StateTransition{
		gp:          gp,
		evm:         evm,
		msg:         msg,
		gasPrice:    msg.GasPrice(),
		initialGas:  new(big.Int),
		value:       msg.Value(),
		data:        msg.Data(),
		state:       evm.StateDB,
		blockNumber: blockNumber,
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg Message, gp *GasPool, blockNumber *big.Int) ([]byte, *big.Int, bool, error) {
	st := NewStateTransition(evm, msg, gp, blockNumber)

	ret, _, gasUsed, failed, err := st.TransitionDb()
	log.Debug("<<ApplyMessage>>", "err", err, "num", blockNumber, "from", st.from().Address().Hex())
	return ret, gasUsed, failed, err
}

func (st *StateTransition) from() vm.AccountRef {
	f := st.msg.From()
	if !st.state.Exist(f) {
		st.state.CreateAccount(f)
	}
	return vm.AccountRef(f)
}

func (st *StateTransition) to() vm.AccountRef {
	if st.msg == nil {
		return vm.AccountRef{}
	}
	to := st.msg.To()
	if to == nil {
		return vm.AccountRef{} // contract creation
	}

	reference := vm.AccountRef(*to)
	if !st.state.Exist(*to) {
		st.state.CreateAccount(*to)
	}
	return reference
}

func (st *StateTransition) useGas(amount uint64) error {
	if st.gas < amount {
		return vm.ErrOutOfGas
	}
	st.gas -= amount

	return nil
}

// TODO : if chief tx skip balance verify.
func (st *StateTransition) IsChiefSIP100() bool {
	if params.IsSIP100Block(st.blockNumber) &&
		st.to().Address() != common.HexToAddress("0x") &&
		params.IsChiefAddress(st.to().Address()) {
		return true
	}
	return false
}

func (st *StateTransition) buyGas() error {

	mgas := st.msg.Gas()

	if mgas.BitLen() > 64 {
		return vm.ErrOutOfGas
	}
	mgval := new(big.Int).Mul(mgas, st.gasPrice)

	var (
		state  = st.state
		sender = st.from()
	)

	if params.IsSIP001Block(st.blockNumber) {
		if err := st.gp.SubGas(mgas); err != nil {
			return err
		}
	}
	st.gas += mgas.Uint64()
	st.initialGas.Set(mgas)
	// if chief tx skip balance verify.
	if !st.IsChiefSIP100() {
		isok := state.GetBalance(sender.Address()).Cmp(mgval) < 0
		//if state.GetBalance(sender.Address()).Cmp(mgval) < 0 {
		if isok {
			return errInsufficientBalanceForGas
		}
	}

	log.Debug("<<StateTransition.buyGas>> 1", "num", st.evm.BlockNumber, "from", sender.Address().Hex(), "gas*price", mgval)
	state.SubBalance(sender.Address(), mgval)

	return nil
}

func (st *StateTransition) preCheck() error {
	msg := st.msg
	sender := st.from()

	// Make sure this transaction's nonce is correct
	if msg.CheckNonce() {
		nonce := st.state.GetNonce(sender.Address())

		if nonce < msg.Nonce() {
			return ErrNonceTooHigh
		} else if nonce > msg.Nonce() {
			return ErrNonceTooLow
		}
	}
	return st.buyGas()
}

// TransitionDb will transition the state by applying the current message and returning the result
// including the required gas for the operation as well as the used gas. It returns an error if it
// failed. An error indicates a consensus issue.
func (st *StateTransition) TransitionDb() (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {
	if err = st.preCheck(); err != nil {
		return
	}
	num := st.evm.BlockNumber
	msg := st.msg
	sender := st.from() // err checked in preCheck

	homestead := st.evm.ChainConfig().IsHomestead(num)
	contractCreation := msg.To() == nil

	// Pay intrinsic gas
	// TODO convert to uint64

	// add by liangc
	intrinsicGas := IntrinsicGas(st.data, contractCreation, homestead)

	if intrinsicGas.BitLen() > 64 {
		return nil, nil, nil, false, vm.ErrOutOfGas
	}
	if err = st.useGas(intrinsicGas.Uint64()); err != nil {
		return nil, nil, nil, false, err
	}

	var (
		evm = st.evm
		// vm errors do not effect consensus and are therefor
		// not assigned to err, except for insufficient balance
		// error.
		vmerr error
	)

	if contractCreation {
		ret, _, st.gas, vmerr = evm.Create(sender, st.data, st.gas, st.value)
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(sender.Address(), st.state.GetNonce(sender.Address())+1)
		//fmt.Println("-TransitionDb1>>>>>>",st.gas)
		ret, st.gas, vmerr = evm.Call(sender, st.to().Address(), st.data, st.gas, st.value)
		//fmt.Println("-TransitionDb2>>>>>>",st.gas)
	}
	if vmerr != nil {
		log.Debug("VM returned with error", "err", vmerr)
		// The only possible consensus-error would be if there wasn't
		// sufficient balance to make the transfer happen. The first
		// balance transfer may never fail.
		if vmerr == vm.ErrInsufficientBalance {
			return nil, nil, nil, false, vmerr
		}
	}

	requiredGas = new(big.Int).Set(st.gasUsed())

	st.refundGas()
	_m := st.evm.Coinbase
	_r := new(big.Int).Mul(st.gasUsed(), st.gasPrice)
	if params.IsChiefAddress(st.to().Address()) {
		_m = sender.Address()
	}

	log.Debug("<<StateTransition.TransitionDb>> 3", "num", st.evm.BlockNumber, "from", sender.Address().Hex(), "miner", st.evm.Coinbase.Hex(), "gas_reward", _r)
	st.state.AddBalance(_m, _r)
	return ret, requiredGas, st.gasUsed(), vmerr != nil, err
}

func (st *StateTransition) refundGas() {
	// Return eth for remaining gas to the sender account,
	// exchanged at the original rate.
	//sender := st.from() // err already checked
	//remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
	//log.Info("<<StateTransition.refundGas>> 2", "num", st.evm.BlockNumber, "from", sender.Address().Hex(), "gas*price", remaining)
	//st.state.AddBalance(sender.Address(), remaining)
	//
	//// Apply refund counter, capped to half of the used gas.
	////uhalf := remaining.Div(st.gasUsed(), common.Big2)
	////refund := math.BigMin(uhalf, st.state.GetRefund())
	////st.gas += refund.Uint64()
	////
	////st.state.AddBalance(sender.Address(), refund.Mul(refund, st.gasPrice))
	//
	//uhalf := remaining.Div(st.gasUsed(), common.Big2)
	//refund := uhalf.Uint64()
	//if uhalf.Uint64()> st.state.GetRefund(){
	//	refund = st.state.GetRefund()
	//}
	//st.gas += refund
	//
	//st.state.AddBalance(sender.Address(), new(big.Int).Mul(new(big.Int).SetUint64(refund), st.gasPrice))
	//
	//// Also return remaining gas to the block gas counter so it is
	//// available for the next transaction.
	//st.gp.AddGas(new(big.Int).SetUint64(st.gas))

	refund := st.gasUsed().Uint64() / 2
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.gas += refund
	// Return ETH for remaining gas, exchanged at the original rate.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(new(big.Int).SetUint64(st.gas))
}

func (st *StateTransition) gasUsed() *big.Int {
	return new(big.Int).Sub(st.initialGas, new(big.Int).SetUint64(st.gas))
}
