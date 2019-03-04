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

// MeshBoxABI is the input ABI used to generate the binding from.
const MeshBoxABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"delAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"existAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"addAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// MeshBoxBin is the compiled bytecode used for deploying new contracts.
const MeshBoxBin = `0x608060405260018054600160a060020a03199081169091556000805490911633179055610475806100316000396000f3fe608060405234801561001057600080fd5b506004361061007e577c010000000000000000000000000000000000000000000000000000000060003504635022edf8811461008357806379ba50971461012857806380b7069d146101305780638da5cb5b14610168578063a6f9dae11461018c578063ba4c18db146101b2575b600080fd5b6101266004803603602081101561009957600080fd5b8101906020810181356401000000008111156100b457600080fd5b8201836020820111156100c657600080fd5b803590602001918460208302840111640100000000831117156100e857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610257945050505050565b005b6101266102bf565b6101566004803603602081101561014657600080fd5b5035600160a060020a0316610356565b60408051918252519081900360200190f35b610170610371565b60408051600160a060020a039092168252519081900360200190f35b610126600480360360208110156101a257600080fd5b5035600160a060020a0316610380565b610126600480360360408110156101c857600080fd5b8101906020810181356401000000008111156101e357600080fd5b8201836020820111156101f557600080fd5b8035906020019184602083028401116401000000008311171561021757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506103e1915050565b600054600160a060020a0316331461026e57600080fd5b805160005b818110156102ba57600060026000858481518110151561028f57fe5b6020908102909101810151600160a060020a0316825281019190915260400160002055600101610273565b505050565b600154600160a060020a031633146102d657600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600160a060020a031660009081526002602052604090205490565b600054600160a060020a031681565b600054600160a060020a0316331461039757600080fd5b600054600160a060020a03828116911614156103b257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146103f857600080fd5b815160005b81811015610443578260026000868481518110151561041857fe5b6020908102909101810151600160a060020a03168252810191909152604001600020556001016103fd565b5050505056fea165627a7a723058207fe03e92900c16c6d07c7edb9fdcb3f4f5ea3125ec5b8d82d0454220e0f0a6cd0029`

// DeployMeshBox deploys a new Ethereum contract, binding an instance of MeshBox to it.
func DeployMeshBox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MeshBox, error) {
	parsed, err := abi.JSON(strings.NewReader(MeshBoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MeshBoxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MeshBox{MeshBoxCaller: MeshBoxCaller{contract: contract}, MeshBoxTransactor: MeshBoxTransactor{contract: contract}}, nil
}

// MeshBox is an auto generated Go binding around an Ethereum contract.
type MeshBox struct {
	MeshBoxCaller     // Read-only binding to the contract
	MeshBoxTransactor // Write-only binding to the contract
}

// MeshBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MeshBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeshBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MeshBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeshBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MeshBoxSession struct {
	Contract     *MeshBox                // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MeshBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MeshBoxCallerSession struct {
	Contract *MeshBoxCaller          // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// MeshBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MeshBoxTransactorSession struct {
	Contract     *MeshBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MeshBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MeshBoxRaw struct {
	Contract *MeshBox // Generic contract binding to access the raw methods on
}

// MeshBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MeshBoxCallerRaw struct {
	Contract *MeshBoxCaller // Generic read-only contract binding to access the raw methods on
}

// MeshBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MeshBoxTransactorRaw struct {
	Contract *MeshBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeshBox creates a new instance of MeshBox, bound to a specific deployed contract.
func NewMeshBox(address common.Address, backend bind.ContractBackend) (*MeshBox, error) {
	contract, err := bindMeshBox(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MeshBox{MeshBoxCaller: MeshBoxCaller{contract: contract}, MeshBoxTransactor: MeshBoxTransactor{contract: contract}}, nil
}

// NewMeshBoxCaller creates a new read-only instance of MeshBox, bound to a specific deployed contract.
func NewMeshBoxCaller(address common.Address, caller bind.ContractCaller) (*MeshBoxCaller, error) {
	contract, err := bindMeshBox(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MeshBoxCaller{contract: contract}, nil
}

// NewMeshBoxTransactor creates a new write-only instance of MeshBox, bound to a specific deployed contract.
func NewMeshBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*MeshBoxTransactor, error) {
	contract, err := bindMeshBox(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MeshBoxTransactor{contract: contract}, nil
}

// bindMeshBox binds a generic wrapper to an already deployed contract.
func bindMeshBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MeshBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeshBox *MeshBoxRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _MeshBox.Contract.MeshBoxCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeshBox *MeshBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox.Contract.MeshBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeshBox *MeshBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeshBox.Contract.MeshBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeshBox *MeshBoxCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _MeshBox.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeshBox *MeshBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeshBox *MeshBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeshBox.Contract.contract.Transact(opts, method, params...)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox *MeshBoxCaller) ExistAddress(opts *bind.CallOptsWithNumber, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MeshBox.contract.CallWithNumber(opts, out, "existAddress", _owner)
	return *ret0, err
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox *MeshBoxSession) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _MeshBox.Contract.ExistAddress(&_MeshBox.CallOpts, _owner)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_MeshBox *MeshBoxCallerSession) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _MeshBox.Contract.ExistAddress(&_MeshBox.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox *MeshBoxCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MeshBox.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox *MeshBoxSession) Owner() (common.Address, error) {
	return _MeshBox.Contract.Owner(&_MeshBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MeshBox *MeshBoxCallerSession) Owner() (common.Address, error) {
	return _MeshBox.Contract.Owner(&_MeshBox.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox *MeshBoxTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeshBox.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox *MeshBoxSession) AcceptOwnership() (*types.Transaction, error) {
	return _MeshBox.Contract.AcceptOwnership(&_MeshBox.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MeshBox *MeshBoxTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MeshBox.Contract.AcceptOwnership(&_MeshBox.TransactOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_MeshBox *MeshBoxTransactor) AddAddress(opts *bind.TransactOpts, _owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _MeshBox.contract.Transact(opts, "addAddress", _owners, version)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_MeshBox *MeshBoxSession) AddAddress(_owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _MeshBox.Contract.AddAddress(&_MeshBox.TransactOpts, _owners, version)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_MeshBox *MeshBoxTransactorSession) AddAddress(_owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _MeshBox.Contract.AddAddress(&_MeshBox.TransactOpts, _owners, version)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox *MeshBoxTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox *MeshBoxSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox.Contract.ChangeOwner(&_MeshBox.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_MeshBox *MeshBoxTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _MeshBox.Contract.ChangeOwner(&_MeshBox.TransactOpts, _newOwner)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox *MeshBoxTransactor) DelAddress(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _MeshBox.contract.Transact(opts, "delAddress", _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox *MeshBoxSession) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _MeshBox.Contract.DelAddress(&_MeshBox.TransactOpts, _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_MeshBox *MeshBoxTransactorSession) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _MeshBox.Contract.DelAddress(&_MeshBox.TransactOpts, _owners)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405260018054600160a060020a031916905534801561002057600080fd5b5060008054600160a060020a031916331790556101e9806100426000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046379ba509781146100625780638da5cb5b1461006c578063a6f9dae114610090575b600080fd5b61006a6100b6565b005b61007461014d565b60408051600160a060020a039092168252519081900360200190f35b61006a600480360360208110156100a657600080fd5b5035600160a060020a031661015c565b600154600160a060020a031633146100cd57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600054600160a060020a0316331461017357600080fd5b600054600160a060020a038281169116141561018e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a7230582042576fa4333a57e21d2bfeeea4d8fa8673862f82ef805e512b1f34ce082646690029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned                  // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller            // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}
