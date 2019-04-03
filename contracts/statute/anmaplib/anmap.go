// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package anmaplib

import (
	"strings"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
)

// AnmapABI is the input ABI used to generate the binding from.
const AnmapABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"bindInfo\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"nodeid\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodePub\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"unbindBySig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodePub\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// AnmapBin is the compiled bytecode used for deploying new contracts.
const AnmapBin = `0x608060405260018054600160a060020a0319908116909155600080549091163317905561068e806100316000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c01000000000000000000000000000000000000000000000000000000009004806399ab9fca1161007857806399ab9fca14610119578063a6f9dae114610154578063b6b257421461017a578063f65f2eec146101825761009a565b806379ba50971461009f5780638d04cd30146100a95780638da5cb5b146100f5575b600080fd5b6100a76101bd565b005b6100cf600480360360208110156100bf57600080fd5b5035600160a060020a0316610254565b60408051600160a060020a03938416815291909216602082015281519081900390910190f35b6100fd6102fe565b60408051600160a060020a039092168252519081900360200190f35b6100a76004803603608081101561012f57600080fd5b50600160a060020a038135169060ff602082013516906040810135906060013561030d565b6100a76004803603602081101561016a57600080fd5b5035600160a060020a031661043c565b6100a761049d565b6100a76004803603608081101561019857600080fd5b50600160a060020a038135169060ff6020820135169060408101359060600135610510565b600154600160a060020a031633146101d457600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600160a060020a038082166000908152600360209081526040808320546002909252909120549082169116811580156102955750600160a060020a03811615155b156102b857600160a060020a038082166000908152600360205260409020541691505b600160a060020a038216158015906102d75750600160a060020a038116155b156102f95750600160a060020a03808216600090815260026020526040902054165b915091565b600054600160a060020a031681565b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff881660748501526094840187905260b484018690529351909260019260d4808301939192601f198301929081900390910190855afa15801561039c573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a03868116908216146103c157600080fd5b600160a060020a03808716600090815260036020526040902054168015156103e857600080fd5b600160a060020a039081166000908152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19908116909155999093168252600390522080549096169095555050505050565b600054600160a060020a0316331461045357600080fd5b600054600160a060020a038281169116141561046e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b33600090815260026020526040902054600160a060020a03168015156104c257600080fd5b336000908152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19908116909155600160a060020a039490941683526003909152902080549091169055565b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff881660748501526094840187905260b484018690529351909260019260d4808301939192601f198301929081900390910190855afa15801561059f573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a03868116908216146105c457600080fd5b33600090815260026020526040902054600160a060020a0316156105e757600080fd5b600160a060020a03868116600090815260036020526040902054161561060c57600080fd5b50503360008181526002602090815260408083208054600160a060020a0390991673ffffffffffffffffffffffffffffffffffffffff19998a16811790915583526003909152902080549095161790935550505056fea165627a7a72305820ca8a001d7c8bddce0a40d1a41a72918b4deda646d7c0f1015be4d56f64fdf9e00029`

// DeployAnmap deploys a new Ethereum contract, binding an instance of Anmap to it.
func DeployAnmap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Anmap, error) {
	parsed, err := abi.JSON(strings.NewReader(AnmapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AnmapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Anmap{AnmapCaller: AnmapCaller{contract: contract}, AnmapTransactor: AnmapTransactor{contract: contract}}, nil
}

// Anmap is an auto generated Go binding around an Ethereum contract.
type Anmap struct {
	AnmapCaller     // Read-only binding to the contract
	AnmapTransactor // Write-only binding to the contract
}

// AnmapCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnmapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnmapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnmapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnmapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnmapSession struct {
	Contract     *Anmap                  // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AnmapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnmapCallerSession struct {
	Contract *AnmapCaller            // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// AnmapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnmapTransactorSession struct {
	Contract     *AnmapTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnmapRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnmapRaw struct {
	Contract *Anmap // Generic contract binding to access the raw methods on
}

// AnmapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnmapCallerRaw struct {
	Contract *AnmapCaller // Generic read-only contract binding to access the raw methods on
}

// AnmapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnmapTransactorRaw struct {
	Contract *AnmapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnmap creates a new instance of Anmap, bound to a specific deployed contract.
func NewAnmap(address common.Address, backend bind.ContractBackend) (*Anmap, error) {
	contract, err := bindAnmap(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Anmap{AnmapCaller: AnmapCaller{contract: contract}, AnmapTransactor: AnmapTransactor{contract: contract}}, nil
}

// NewAnmapCaller creates a new read-only instance of Anmap, bound to a specific deployed contract.
func NewAnmapCaller(address common.Address, caller bind.ContractCaller) (*AnmapCaller, error) {
	contract, err := bindAnmap(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AnmapCaller{contract: contract}, nil
}

// NewAnmapTransactor creates a new write-only instance of Anmap, bound to a specific deployed contract.
func NewAnmapTransactor(address common.Address, transactor bind.ContractTransactor) (*AnmapTransactor, error) {
	contract, err := bindAnmap(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AnmapTransactor{contract: contract}, nil
}

// bindAnmap binds a generic wrapper to an already deployed contract.
func bindAnmap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnmapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anmap *AnmapRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Anmap.Contract.AnmapCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anmap *AnmapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anmap.Contract.AnmapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anmap *AnmapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anmap.Contract.AnmapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anmap *AnmapCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Anmap.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anmap *AnmapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anmap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anmap *AnmapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anmap.Contract.contract.Transact(opts, method, params...)
}

// BindInfo is a free data retrieval call binding the contract method 0x8d04cd30.
//
// Solidity: function bindInfo(addr address) constant returns(from address, nodeid address)
func (_Anmap *AnmapCaller) BindInfo(opts *bind.CallOptsWithNumber, addr common.Address) (struct {
	From   common.Address
	Nodeid common.Address
}, error) {
	ret := new(struct {
		From   common.Address
		Nodeid common.Address
	})
	out := ret
	err := _Anmap.contract.CallWithNumber(opts, out, "bindInfo", addr)
	return *ret, err
}

// BindInfo is a free data retrieval call binding the contract method 0x8d04cd30.
//
// Solidity: function bindInfo(addr address) constant returns(from address, nodeid address)
func (_Anmap *AnmapSession) BindInfo(addr common.Address) (struct {
	From   common.Address
	Nodeid common.Address
}, error) {
	return _Anmap.Contract.BindInfo(&_Anmap.CallOpts, addr)
}

// BindInfo is a free data retrieval call binding the contract method 0x8d04cd30.
//
// Solidity: function bindInfo(addr address) constant returns(from address, nodeid address)
func (_Anmap *AnmapCallerSession) BindInfo(addr common.Address) (struct {
	From   common.Address
	Nodeid common.Address
}, error) {
	return _Anmap.Contract.BindInfo(&_Anmap.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Anmap *AnmapCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Anmap.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Anmap *AnmapSession) Owner() (common.Address, error) {
	return _Anmap.Contract.Owner(&_Anmap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Anmap *AnmapCallerSession) Owner() (common.Address, error) {
	return _Anmap.Contract.Owner(&_Anmap.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Anmap *AnmapTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anmap.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Anmap *AnmapSession) AcceptOwnership() (*types.Transaction, error) {
	return _Anmap.Contract.AcceptOwnership(&_Anmap.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Anmap *AnmapTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Anmap.Contract.AcceptOwnership(&_Anmap.TransactOpts)
}

// Bind is a paid mutator transaction binding the contract method 0xf65f2eec.
//
// Solidity: function bind(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapTransactor) Bind(opts *bind.TransactOpts, nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.contract.Transact(opts, "bind", nodePub, v, r, s)
}

// Bind is a paid mutator transaction binding the contract method 0xf65f2eec.
//
// Solidity: function bind(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapSession) Bind(nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.Contract.Bind(&_Anmap.TransactOpts, nodePub, v, r, s)
}

// Bind is a paid mutator transaction binding the contract method 0xf65f2eec.
//
// Solidity: function bind(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapTransactorSession) Bind(nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.Contract.Bind(&_Anmap.TransactOpts, nodePub, v, r, s)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Anmap *AnmapTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Anmap.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Anmap *AnmapSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Anmap.Contract.ChangeOwner(&_Anmap.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Anmap *AnmapTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Anmap.Contract.ChangeOwner(&_Anmap.TransactOpts, _newOwner)
}

// Unbind is a paid mutator transaction binding the contract method 0xb6b25742.
//
// Solidity: function unbind() returns()
func (_Anmap *AnmapTransactor) Unbind(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anmap.contract.Transact(opts, "unbind")
}

// Unbind is a paid mutator transaction binding the contract method 0xb6b25742.
//
// Solidity: function unbind() returns()
func (_Anmap *AnmapSession) Unbind() (*types.Transaction, error) {
	return _Anmap.Contract.Unbind(&_Anmap.TransactOpts)
}

// Unbind is a paid mutator transaction binding the contract method 0xb6b25742.
//
// Solidity: function unbind() returns()
func (_Anmap *AnmapTransactorSession) Unbind() (*types.Transaction, error) {
	return _Anmap.Contract.Unbind(&_Anmap.TransactOpts)
}

// UnbindBySig is a paid mutator transaction binding the contract method 0x99ab9fca.
//
// Solidity: function unbindBySig(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapTransactor) UnbindBySig(opts *bind.TransactOpts, nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.contract.Transact(opts, "unbindBySig", nodePub, v, r, s)
}

// UnbindBySig is a paid mutator transaction binding the contract method 0x99ab9fca.
//
// Solidity: function unbindBySig(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapSession) UnbindBySig(nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.Contract.UnbindBySig(&_Anmap.TransactOpts, nodePub, v, r, s)
}

// UnbindBySig is a paid mutator transaction binding the contract method 0x99ab9fca.
//
// Solidity: function unbindBySig(nodePub address, v uint8, r bytes32, s bytes32) returns()
func (_Anmap *AnmapTransactorSession) UnbindBySig(nodePub common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anmap.Contract.UnbindBySig(&_Anmap.TransactOpts, nodePub, v, r, s)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405260018054600160a060020a031916905534801561002057600080fd5b5060008054600160a060020a031916331790556101e9806100426000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046379ba509781146100625780638da5cb5b1461006c578063a6f9dae114610090575b600080fd5b61006a6100b6565b005b61007461014d565b60408051600160a060020a039092168252519081900360200190f35b61006a600480360360208110156100a657600080fd5b5035600160a060020a031661015c565b600154600160a060020a031633146100cd57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600054600160a060020a0316331461017357600080fd5b600054600160a060020a038281169116141561018e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a723058205b46aa774f6d4c950da426372226a4e1f10918038840db11f5fd98dd403ced7e0029`

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
