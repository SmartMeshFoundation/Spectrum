// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meshboxlib

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
)

// MeshBox_0_0_2ABI is the input ABI used to generate the binding from.
const MeshBox_0_0_2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMeshboxList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"delAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"existAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"addAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// MeshBox_0_0_2Bin is the compiled bytecode used for deploying new contracts.
const MeshBox_0_0_2Bin = `0x60018054600160a060020a031916905560c0604052600560808190527f302e302e3200000000000000000000000000000000000000000000000000000060a090815261004e9160029190610066565b5060008054600160a060020a03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610804806101106000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c01000000000000000000000000000000000000000000000000000000009004806380b7069d1161007857806380b7069d1461022c5780638da5cb5b14610264578063a6f9dae114610288578063ba4c18db146102ae576100a5565b80633ec7eed1146100aa5780635022edf81461010257806354fd4d50146101a757806379ba509714610224575b600080fd5b6100b2610353565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ee5781810151838201526020016100d6565b505050509050019250505060405180910390f35b6101a56004803603602081101561011857600080fd5b81019060208101813564010000000081111561013357600080fd5b82018360208201111561014557600080fd5b8035906020019184602083028401116401000000008311171561016757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506103b6945050505050565b005b6101af610407565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e95781810151838201526020016101d1565b50505050905090810190601f1680156102165780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a5610491565b6102526004803603602081101561024257600080fd5b5035600160a060020a0316610528565b60408051918252519081900360200190f35b61026c610543565b60408051600160a060020a039092168252519081900360200190f35b6101a56004803603602081101561029e57600080fd5b5035600160a060020a0316610552565b6101a5600480360360408110156102c457600080fd5b8101906020810181356401000000008111156102df57600080fd5b8201836020820111156102f157600080fd5b8035906020019184602083028401116401000000008311171561031357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506105b3915050565b606060048054806020026020016040519081016040528092919081815260200182805480156103ab57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161038d575b505050505090505b90565b600054600160a060020a031633146103cd57600080fd5b805160005b81811015610402576103fa83828151811015156103eb57fe5b90602001906020020151610696565b6001016103d2565b505050565b60028054604080516020601f60001961010060018716150201909416859004938401819004810282018101909252828152606093909290918301828280156103ab5780601f10610465576101008083540402835291602001916103ab565b820191906000526020600020905b81548152906001019060200180831161047357509395945050505050565b600154600160a060020a031633146104a857600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600160a060020a031660009081526003602052604090205490565b600054600160a060020a031681565b600054600160a060020a0316331461056957600080fd5b600054600160a060020a038281169116141561058457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146105ca57600080fd5b815160005b8181101561069057600084828151811015156105e757fe5b6020908102909101810151600160a060020a03811660009081526003909252604090912054909150151561066e57600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b600160a060020a031660009081526003602052604090208390556001016105cf565b50505050565b600160a060020a03811660009081526003602052604081208190555b6004548110156107975781600160a060020a03166004828154811015156106d557fe5b600091825260209091200154600160a060020a0316141561078f57805b6004546000190181101561077557600480546001830190811061071157fe5b60009182526020909120015460048054600160a060020a03909216918390811061073757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016106f2565b506004805460001992830192019061078d908261079b565b505b6001016106b2565b5050565b815481835581811115610402576000838152602090206104029181019083016103b391905b808211156107d457600081556001016107c0565b509056fea165627a7a723058209e3a10e2851ac0833bead3b88d1d6a611e324b8d6f40c0e7f2562596c37d3b230029`

// DeployMeshBox_0_0_2 deploys a new Ethereum contract, binding an instance of MeshBox_0_0_2 to it.
func DeployMeshBox_0_0_2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MeshBox_0_0_2, error) {
	parsed, err := abi.JSON(strings.NewReader(MeshBox_0_0_2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MeshBox_0_0_2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MeshBox_0_0_2{MeshBox_0_0_2Caller: MeshBox_0_0_2Caller{contract: contract}, MeshBox_0_0_2Transactor: MeshBox_0_0_2Transactor{contract: contract}}, nil
}

// MeshBox_0_0_2 is an auto generated Go binding around an Ethereum contract.
type MeshBox_0_0_2 struct {
	MeshBox_0_0_2Caller     // Read-only binding to the contract
	MeshBox_0_0_2Transactor // Write-only binding to the contract
}

// MeshBox_0_0_2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MeshBox_0_0_2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeshBox_0_0_2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MeshBox_0_0_2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeshBox_0_0_2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MeshBox_0_0_2Session struct {
	Contract     *MeshBox_0_0_2          // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MeshBox_0_0_2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MeshBox_0_0_2CallerSession struct {
	Contract *MeshBox_0_0_2Caller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// MeshBox_0_0_2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MeshBox_0_0_2TransactorSession struct {
	Contract     *MeshBox_0_0_2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MeshBox_0_0_2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MeshBox_0_0_2Raw struct {
	Contract *MeshBox_0_0_2 // Generic contract binding to access the raw methods on
}

// MeshBox_0_0_2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MeshBox_0_0_2CallerRaw struct {
	Contract *MeshBox_0_0_2Caller // Generic read-only contract binding to access the raw methods on
}

// MeshBox_0_0_2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MeshBox_0_0_2TransactorRaw struct {
	Contract *MeshBox_0_0_2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMeshBox_0_0_2 creates a new instance of MeshBox_0_0_2, bound to a specific deployed contract.
func NewMeshBox_0_0_2(address common.Address, backend bind.ContractBackend) (*MeshBox_0_0_2, error) {
	contract, err := bindMeshBox_0_0_2(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MeshBox_0_0_2{MeshBox_0_0_2Caller: MeshBox_0_0_2Caller{contract: contract}, MeshBox_0_0_2Transactor: MeshBox_0_0_2Transactor{contract: contract}}, nil
}

// NewMeshBox_0_0_2Caller creates a new read-only instance of MeshBox_0_0_2, bound to a specific deployed contract.
func NewMeshBox_0_0_2Caller(address common.Address, caller bind.ContractCaller) (*MeshBox_0_0_2Caller, error) {
	contract, err := bindMeshBox_0_0_2(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MeshBox_0_0_2Caller{contract: contract}, nil
}

// NewMeshBox_0_0_2Transactor creates a new write-only instance of MeshBox_0_0_2, bound to a specific deployed contract.
func NewMeshBox_0_0_2Transactor(address common.Address, transactor bind.ContractTransactor) (*MeshBox_0_0_2Transactor, error) {
	contract, err := bindMeshBox_0_0_2(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MeshBox_0_0_2Transactor{contract: contract}, nil
}

// bindMeshBox_0_0_2 binds a generic wrapper to an already deployed contract.
func bindMeshBox_0_0_2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MeshBox_0_0_2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeshBox_0_0_2 *MeshBox_0_0_2Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _MeshBox_0_0_2.Contract.MeshBox_0_0_2Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeshBox_0_0_2 *MeshBox_0_0_2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.MeshBox_0_0_2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeshBox_0_0_2 *MeshBox_0_0_2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.MeshBox_0_0_2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeshBox_0_0_2 *MeshBox_0_0_2CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _MeshBox_0_0_2.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.contract.Transact(opts, method, params...)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Caller) ExistAddress(opts *bind.CallOptsWithNumber, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MeshBox_0_0_2.contract.CallWithNumber(opts, out, "existAddress", _owner)
	return *ret0, err
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _MeshBox_0_0_2.Contract.ExistAddress(&_MeshBox_0_0_2.CallOpts, _owner)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox_0_0_2 *MeshBox_0_0_2CallerSession) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _MeshBox_0_0_2.Contract.ExistAddress(&_MeshBox_0_0_2.CallOpts, _owner)
}

// GetMeshboxList is a free data retrieval call binding the contract method 0x3ec7eed1.
//
// Solidity: function getMeshboxList() constant returns(address[])
func (_MeshBox_0_0_2 *MeshBox_0_0_2Caller) GetMeshboxList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MeshBox_0_0_2.contract.CallWithNumber(opts, out, "getMeshboxList")
	return *ret0, err
}

// GetMeshboxList is a free data retrieval call binding the contract method 0x3ec7eed1.
//
// Solidity: function getMeshboxList() constant returns(address[])
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) GetMeshboxList() ([]common.Address, error) {
	return _MeshBox_0_0_2.Contract.GetMeshboxList(&_MeshBox_0_0_2.CallOpts)
}

// GetMeshboxList is a free data retrieval call binding the contract method 0x3ec7eed1.
//
// Solidity: function getMeshboxList() constant returns(address[])
func (_MeshBox_0_0_2 *MeshBox_0_0_2CallerSession) GetMeshboxList() ([]common.Address, error) {
	return _MeshBox_0_0_2.Contract.GetMeshboxList(&_MeshBox_0_0_2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Caller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MeshBox_0_0_2.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) Owner() (common.Address, error) {
	return _MeshBox_0_0_2.Contract.Owner(&_MeshBox_0_0_2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox_0_0_2 *MeshBox_0_0_2CallerSession) Owner() (common.Address, error) {
	return _MeshBox_0_0_2.Contract.Owner(&_MeshBox_0_0_2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MeshBox_0_0_2.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) Version() (string, error) {
	return _MeshBox_0_0_2.Contract.Version(&_MeshBox_0_0_2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_MeshBox_0_0_2 *MeshBox_0_0_2CallerSession) Version() (string, error) {
	return _MeshBox_0_0_2.Contract.Version(&_MeshBox_0_0_2.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox_0_0_2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) AcceptOwnership() (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.AcceptOwnership(&_MeshBox_0_0_2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.AcceptOwnership(&_MeshBox_0_0_2.TransactOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], weight uint256) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Transactor) AddAddress(opts *bind.TransactOpts, _owners []common.Address, weight *big.Int) (*types.Transaction, error) {
	return _MeshBox_0_0_2.contract.Transact(opts, "addAddress", _owners, weight)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], weight uint256) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) AddAddress(_owners []common.Address, weight *big.Int) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.AddAddress(&_MeshBox_0_0_2.TransactOpts, _owners, weight)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], weight uint256) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorSession) AddAddress(_owners []common.Address, weight *big.Int) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.AddAddress(&_MeshBox_0_0_2.TransactOpts, _owners, weight)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Transactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.ChangeOwner(&_MeshBox_0_0_2.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.ChangeOwner(&_MeshBox_0_0_2.TransactOpts, _newOwner)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Transactor) DelAddress(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.contract.Transact(opts, "delAddress", _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2Session) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.DelAddress(&_MeshBox_0_0_2.TransactOpts, _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox_0_0_2 *MeshBox_0_0_2TransactorSession) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _MeshBox_0_0_2.Contract.DelAddress(&_MeshBox_0_0_2.TransactOpts, _owners)
}
