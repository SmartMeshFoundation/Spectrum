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
const AnmapABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"bindInfo\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"nids\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodePub\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"unbindBySig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodePub\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AnmapBin is the compiled bytecode used for deploying new contracts.
const AnmapBin = `0x608060405234801561001057600080fd5b506106d9806100206000396000f3fe608060405234801561001057600080fd5b506004361061005d577c010000000000000000000000000000000000000000000000000000000060003504638d04cd30811461006257806399ab9fca146100f5578063f65f2eec14610132575b600080fd5b6100886004803603602081101561007857600080fd5b5035600160a060020a031661016d565b6040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156100e05781810151838201526020016100c8565b50505050905001935050505060405180910390f35b6101306004803603608081101561010b57600080fd5b50600160a060020a038135169060ff6020820135169060408101359060600135610253565b005b6101306004803603608081101561014857600080fd5b50600160a060020a038135169060ff6020820135169060408101359060600135610508565b600160a060020a0380821660009081526001602052604090205416606081151561020657600160a060020a03831660009081526020818152604091829020805483518184028101840190945280845286955090918301828280156101fa57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116101dc575b5050505050905061024e565b60408051600180825281830190925290602080830190803883390190505090508281600081518110151561023657fe5b600160a060020a039092166020928302909101909101525b915091565b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff881660748501526094840187905260b484018690529351909260019260d4808301939192601f198301929081900390910190855afa1580156102e2573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a038681169082161461030757600080fd5b600160a060020a038082166000908152600160205260409020541680151561032e57600080fd5b600160a060020a038083166000908152600160209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19169055928416825281905220546002111561039d57600160a060020a038116600090815260208190526040812061039891610642565b6104ff565b60005b600160a060020a0382166000908152602081905260409020548110156104fd57600160a060020a038281166000908152602081905260409020805491851691839081106103e957fe5b600091825260209091200154600160a060020a031614156104f557805b600160a060020a038316600090815260208190526040902054600019018110156104c557600160a060020a038316600090815260208190526040902080546001830190811061045157fe5b6000918252602080832090910154600160a060020a03868116845291839052604090922080549190921691908390811061048757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610406565b50600160a060020a0382166000908152602081905260409020805460001901906104ef9082610663565b506104fd565b6001016103a0565b505b50505050505050565b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff881660748501526094840187905260b484018690529351909260019260d4808301939192601f198301929081900390910190855afa158015610597573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a03868116908216146105bc57600080fd5b600160a060020a0386811660009081526001602052604090205416156105e157600080fd5b600160a060020a0316600081815260016020818152604080842080543373ffffffffffffffffffffffffffffffffffffffff199182168117909255908552848352908420805493840181558452922001805490911690911790555050505050565b5080546000825590600052602060002090810190610660919061068c565b50565b8154818355818111156106875760008381526020902061068791810190830161068c565b505050565b6106aa91905b808211156106a65760008155600101610692565b5090565b9056fea165627a7a72305820e39dcea5237991bf9127aa6b3511c4ddedb63f4192e19becc42c76977d9656380029`

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
// Solidity: function bindInfo(addr address) constant returns(from address, nids address[])
func (_Anmap *AnmapCaller) BindInfo(opts *bind.CallOptsWithNumber, addr common.Address) (struct {
	From common.Address
	Nids []common.Address
}, error) {
	ret := new(struct {
		From common.Address
		Nids []common.Address
	})
	out := ret
	err := _Anmap.contract.CallWithNumber(opts, out, "bindInfo", addr)
	return *ret, err
}

// BindInfo is a free data retrieval call binding the contract method 0x8d04cd30.
//
// Solidity: function bindInfo(addr address) constant returns(from address, nids address[])
func (_Anmap *AnmapSession) BindInfo(addr common.Address) (struct {
	From common.Address
	Nids []common.Address
}, error) {
	return _Anmap.Contract.BindInfo(&_Anmap.CallOpts, addr)
}

// BindInfo is a free data retrieval call binding the contract method 0x8d04cd30.
//
// Solidity: function bindInfo(addr address) constant returns(from address, nids address[])
func (_Anmap *AnmapCallerSession) BindInfo(addr common.Address) (struct {
	From common.Address
	Nids []common.Address
}, error) {
	return _Anmap.Contract.BindInfo(&_Anmap.CallOpts, addr)
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
