// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chieflib

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
)

// ChiefABI is the input ABI used to generate the binding from.
const ChiefABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChiefBin is the compiled bytecode used for deploying new contracts.
const ChiefBin = `0x`

// DeployChief deploys a new Ethereum contract, binding an instance of Chief to it.
func DeployChief(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Chief, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChiefBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Chief{ChiefCaller: ChiefCaller{contract: contract}, ChiefTransactor: ChiefTransactor{contract: contract}}, nil
}

// Chief is an auto generated Go binding around an Ethereum contract.
type Chief struct {
	ChiefCaller     // Read-only binding to the contract
	ChiefTransactor // Write-only binding to the contract
}

// ChiefCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiefSession struct {
	Contract     *Chief                  // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChiefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiefCallerSession struct {
	Contract *ChiefCaller            // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// ChiefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChiefTransactorSession struct {
	Contract     *ChiefTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChiefRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChiefRaw struct {
	Contract *Chief // Generic contract binding to access the raw methods on
}

// ChiefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChiefCallerRaw struct {
	Contract *ChiefCaller // Generic read-only contract binding to access the raw methods on
}

// ChiefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChiefTransactorRaw struct {
	Contract *ChiefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChief creates a new instance of Chief, bound to a specific deployed contract.
func NewChief(address common.Address, backend bind.ContractBackend) (*Chief, error) {
	contract, err := bindChief(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chief{ChiefCaller: ChiefCaller{contract: contract}, ChiefTransactor: ChiefTransactor{contract: contract}}, nil
}

// NewChiefCaller creates a new read-only instance of Chief, bound to a specific deployed contract.
func NewChiefCaller(address common.Address, caller bind.ContractCaller) (*ChiefCaller, error) {
	contract, err := bindChief(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefCaller{contract: contract}, nil
}

// NewChiefTransactor creates a new write-only instance of Chief, bound to a specific deployed contract.
func NewChiefTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiefTransactor, error) {
	contract, err := bindChief(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ChiefTransactor{contract: contract}, nil
}

// bindChief binds a generic wrapper to an already deployed contract.
func bindChief(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chief *ChiefRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Chief.Contract.ChiefCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chief *ChiefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chief.Contract.ChiefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chief *ChiefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chief.Contract.ChiefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chief *ChiefCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Chief.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chief *ChiefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chief.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chief *ChiefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chief.Contract.contract.Transact(opts, method, params...)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_Chief *ChiefCaller) FilterVolunteer(opts *bind.CallOptsWithNumber, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Chief.contract.CallWithNumber(opts, out, "filterVolunteer", volunteers)
	return *ret0, err
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_Chief *ChiefSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _Chief.Contract.FilterVolunteer(&_Chief.CallOpts, volunteers)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_Chief *ChiefCallerSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _Chief.Contract.FilterVolunteer(&_Chief.CallOpts, volunteers)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_Chief *ChiefCaller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_Chief *ChiefSession) GetEpoch() (*big.Int, error) {
	return _Chief.Contract.GetEpoch(&_Chief.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_Chief *ChiefCallerSession) GetEpoch() (*big.Int, error) {
	return _Chief.Contract.GetEpoch(&_Chief.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_Chief *ChiefCaller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_Chief *ChiefSession) GetSignerLimit() (*big.Int, error) {
	return _Chief.Contract.GetSignerLimit(&_Chief.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_Chief *ChiefCallerSession) GetSignerLimit() (*big.Int, error) {
	return _Chief.Contract.GetSignerLimit(&_Chief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_Chief *ChiefCaller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	ret := new(struct {
		SignerList     []common.Address
		BlackList      []common.Address
		ScoreList      []*big.Int
		NumberList     []*big.Int
		TotalVolunteer *big.Int
		Number         *big.Int
	})
	out := ret
	err := _Chief.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_Chief *ChiefSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _Chief.Contract.GetStatus(&_Chief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_Chief *ChiefCallerSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _Chief.Contract.GetStatus(&_Chief.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_Chief *ChiefCaller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_Chief *ChiefSession) GetVolunteerLimit() (*big.Int, error) {
	return _Chief.Contract.GetVolunteerLimit(&_Chief.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_Chief *ChiefCallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _Chief.Contract.GetVolunteerLimit(&_Chief.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_Chief *ChiefCaller) GetVolunteers(opts *bind.CallOptsWithNumber) (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		WeightList    []*big.Int
		Length        *big.Int
	})
	out := ret
	err := _Chief.contract.CallWithNumber(opts, out, "getVolunteers")
	return *ret, err
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_Chief *ChiefSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _Chief.Contract.GetVolunteers(&_Chief.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_Chief *ChiefCallerSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _Chief.Contract.GetVolunteers(&_Chief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Chief *ChiefCaller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Chief.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Chief *ChiefSession) Version() (string, error) {
	return _Chief.Contract.Version(&_Chief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Chief *ChiefCallerSession) Version() (string, error) {
	return _Chief.Contract.Version(&_Chief.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_Chief *ChiefTransactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _Chief.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_Chief *ChiefSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _Chief.Contract.Update(&_Chief.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_Chief *ChiefTransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _Chief.Contract.Update(&_Chief.TransactOpts, volunteer)
}

// TribeChief_0_0_7ABI is the input ABI used to generate the binding from.
const TribeChief_0_0_7ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"},{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"},{\"name\":\"_volunteerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_7Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_7Bin = `0x60c0604052600560808190527f302e302e3700000000000000000000000000000000000000000000000000000060a0908152620000409160009190620002f1565b5061181b60015560116002556104d26003553480156200005f57600080fd5b5060405162001ac038038062001ac0833981018060405260808110156200008557600080fd5b8101908080516401000000008111156200009e57600080fd5b82016020810184811115620000b257600080fd5b8151856020820283011164010000000082111715620000d057600080fd5b50506020820151604083015160609093015191945092506000831115620000f75760018390555b6000821115620001075760028290555b6000811115620001175760038190555b60068054600160a060020a0319163317905583516000811115620001e55760005b81811015620001de57600086828151811015156200015257fe5b6020908102909101810151600160a060020a0381166000818152600490935260408320805460ff1916600190811790915560078054918201815590935260008051602062001aa08339815191529092018054600160a060020a0319169092179091559050811515620001d457620001d481600364010000000062000280810204565b5060010162000138565b5062000275565b60046020527f052387a23a063359ef51016da56c6ef818568f4a38e27c2728fda35f7b8ae85e805460ff1916600190811790915560078054918201815560005260008051602062001aa0833981519152018054600160a060020a031916734110bd1ff0b73fa12c259acf39c950277f2667879081179091556200027381600364010000000062000280810204565b505b505050505062000396565b6002546008541015620002ed576008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018054600160a060020a031916600160a060020a0385169081179091556000908152600b60205260409020828155439101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200033457805160ff191683800117855562000364565b8280016001018555821562000364579182015b828111156200036457825182559160200191906001019062000347565b506200037292915062000376565b5090565b6200039391905b808211156200037257600081556001016200037d565b90565b6116fa80620003a66000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063757991a811610078578063757991a81461037b578063961c5c7a14610395578063d7ca4a1c1461039d578063eb5c001114610445576100a5565b80631c1b8772146100aa57806320c1a518146100d25780634e69d560146101c557806354fd4d50146102fe575b600080fd5b6100d0600480360360208110156100c057600080fd5b5035600160a060020a031661044d565b005b610175600480360360208110156100e857600080fd5b81019060208101813564010000000081111561010357600080fd5b82018360208201111561011557600080fd5b8035906020019184602083028401116401000000008311171561013757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506104fa945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101b1578181015183820152602001610199565b505050509050019250505060405180910390f35b6101cd610615565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b8381101561022557818101518382015260200161020d565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b8381101561026457818101518382015260200161024c565b50505050905001858103835289818151815260200191508051906020019060200280838360005b838110156102a357818101518382015260200161028b565b50505050905001858103825288818151815260200191508051906020019060200280838360005b838110156102e25781810151838201526020016102ca565b505050509050019a505050505050505050505060405180910390f35b610306610800565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610340578181015183820152602001610328565b50505050905090810190601f16801561036d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610383610897565b60408051918252519081900360200190f35b61038361089d565b6103a56108a3565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b838110156103ef5781810151838201526020016103d7565b50505050905001838103825285818151815260200191508051906020019060200280838360005b8381101561042e578181015183820152602001610416565b505050509050019550505050505060405180910390f35b6103836109a3565b3380151561045a57600080fd5b600160a060020a0381166000908152600b60205260408120541161047d57600080fd5b4360058190556001541080156104a0575060015460055481151561049d57fe5b06155b156104ad576104ad6109a9565b600160a060020a038216156104c7576104c7826005610a07565b60025460085410806104dc5750600854600954105b156104ee576104e9610be2565b6104f6565b6104f6610d64565b5050565b60608151604051908082528060200260200182016040528015610527578160200160208202803883390190505b5060035460095491925011156106105760005b825181101561060e576000838281518110151561055357fe5b6020908102909101810151600160a060020a0381166000908152600c9092526040909120600101549091501580156105a15750600160a060020a0381166000908152600d6020526040902054155b80156105c65750600160a060020a0381166000908152600b6020526040902060010154155b156105ea57600183838151811015156105db57fe5b60209081029091010152610605565b600083838151811015156105fa57fe5b602090810290910101525b5060010161053a565b505b919050565b60608060608060008060088054905060405190808252806020026020018201604052801561064d578160200160208202803883390190505b50600854604080518281526020808402820101909152919550801561067c578160200160208202803883390190505b50925060005b60085481101561073857600b600060088381548110151561069f57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205485518690839081106106d257fe5b6020908102909101015260088054600b91600091849081106106f057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154845185908390811061072657fe5b60209081029091010152600101610682565b50600954600a80546040805160208084028201810190925282815293955083018282801561078f57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610771575b5050505050945060088054806020026020016040519081016040528092919081815260200182805480156107ec57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107ce575b505050505095506005549050909192939495565b60008054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561088c5780601f106108615761010080835404028352916020019161088c565b820191906000526020600020905b81548152906001019060200180831161086f57829003601f168201915b505050505090505b90565b60015490565b60035490565b60608060006009805490506040519080825280602002602001820160405280156108d7578160200160208202803883390190505b50600980546040805160208084028201810190925282815293955083018282801561092b57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161090d575b50939650600093505050505b60095481101561099957600c600060098381548110151561095457fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054835184908390811061098757fe5b60209081029091010152600101610937565b5082519050909192565b60025490565b600a5460005b818110156109f757600d6000600a838154811015156109ca57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016109af565b50610a04600a6000611672565b50565b801515610a745760095460001015610a6b5760005b600954811015610a69576009805482908110610a3457fe5b600091825260209091200154600160a060020a0384811691161415610a6157610a5c81610e3a565b610a69565b600101610a1c565b505b6104e982610f19565b600160a060020a03821660009081526004602052604090205460ff16158015610a9d5750806005145b8015610aac5750600354600954105b8015610ad15750600160a060020a0382166000908152600c6020526040902060010154155b8015610af35750600160a060020a0382166000908152600d6020526040902054155b8015610b185750600160a060020a0382166000908152600b6020526040902060010154155b15610b8e576009805460018082019092557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091556000908152600c60205260409020828155439101556104f6565b600581108015610bb75750600160a060020a0382166000908152600c6020526040812060010154115b156104f657600160a060020a0382166000908152600c60205260409020818155436001909101555050565b610bea610fb6565b60085460055460009190811515610bfd57fe5b06905060046000600883815481101515610c1357fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161515610d07576000600b6000600884815481101515610c5657fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600880549192509083908110610c8c57fe5b600091825260209091200154600160a060020a03163314610d0057805460011015610cc65780546000190181556005546001820155610cfb565b610cf2600883815481101515610cd857fe5b6000918252602082200154600160a060020a031690610a07565b610cfb826110b0565b610d05565b600381555b505b600254600854108015610d1c57506009546000105b15610a045760098054610d5491906000198101908110610d3857fe5b600091825260209091200154600160a060020a0316600361118f565b600954610a049060001901610e3a565b6008546005546000908290811515610d7857fe5b0690506000600882815481101515610d8c57fe5b6000918252602082200154600160a060020a03169150821115610e225733600160a060020a03821614610e2257610dc4816000610a07565b600160a060020a0381166000908152600b602052604081208181556001018190556008805484908110610df357fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b60018303821415610e3557610e3561120b565b505050565b600954808210156104f657600c6000600984815481101515610e5857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812081815560010155815b60018203811015610f05576009805460018301908110610ea157fe5b60009182526020909120015460098054600160a060020a039092169183908110610ec757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610e85565b50600980546000190190610e359082611690565b600160a060020a03811615801590610f475750600160a060020a0381166000908152600d6020526040902054155b15610a0457600a8054600181019091557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691821790556000908152600d60205260409020439055565b60005b60085481101561100657600880546000919083908110610fd557fe5b600091825260209091200154600160a060020a03161415610ffe57610ff9816110b0565b600019015b600101610fb9565b5060005b600854811015610a0457600060088281548110151561102557fe5b6000918252602080832090910154600160a060020a0316808352600c909152604082206001015490925011156110a75760005b6009548110156110a557600980548290811061107057fe5b600091825260209091200154600160a060020a038381169116141561109d5761109881610e3a565b6110a5565b600101611058565b505b5060010161100a565b600854808210156104f657600b60006008848154811015156110ce57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812081815560010155815b6001820381101561117b57600880546001830190811061111757fe5b60009182526020909120015460088054600160a060020a03909216918390811061113d57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016110fb565b50600880546000190190610e359082611690565b60025460085410156104f6576008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556000928352600b602052604090922090815543910155565b60006008600081548110151561121d57fe5b600091825260209182902001546008546040805182815282850281019094019052600160a060020a039091169250606091908015611265578160200160208202803883390190505b50905060005b81518110156112c757600880548290811061128257fe5b6000918252602090912001548251600160a060020a03909116908390839081106112a857fe5b600160a060020a0390921660209283029091019091015260010161126b565b5080515b60008111156113785760001981016112e2816110b0565b600083828151811015156112f257fe5b9060200190602002015190506000821180156113165750600160a060020a03811615155b1561136d57600160a060020a0381166000908152600c6020526040902054151561134557611345816005610a07565b600160a060020a0381166000908152600c602052604090205461136d90829060001901610a07565b5050600019016112cb565b50600160a060020a03821660009081526004602052604090205460ff1680156113a357506007546001105b156114375760005b6007548110156114315782600160a060020a03166007828154811015156113ce57fe5b600091825260209091200154600160a060020a03161415611429576007546000190181141561140f5761140a60076000815481101515610d3857fe5b611424565b611424600782600101815481101515610d3857fe5b611431565b6001016113ab565b5061144a565b61144a60076000815481101515610d3857fe5b6060600254604051908082528060200260200182016040528015611478578160200160208202803883390190505b5090506000805b60095481101561152157600254821061149757611521565b60006114cf6009838154811015156114ab57fe5b600091825260209091200154600954600160a060020a0390911690600019016115ac565b90506114db848261161f565b156114e65750611519565b6114f8600982815481101515610d3857fe5b80848481518110151561150757fe5b60209081029091010152506001909101905b60010161147f565b506002548110156115a65760005b6009548110156115a457600b600060098381548110151561154c57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154151561158c5761158c600982815481101515610d3857fe5b6002546008541061159c576115a4565b60010161152f565b505b50505050565b60008082116115bd57506000611619565b604080516c01000000000000000000000000601386810b900b0260208083019190915244603483015242605480840191909152835180840390910181526074909201909252805191012080838181151561161357fe5b06925050505b92915050565b600080835111156116695760005b83518110156116675782848281518110151561164557fe5b90602001906020020151141561165f576001915050611619565b60010161162d565b505b50600092915050565b5080546000825590600052602060002090810190610a0491906116b0565b815481835581811115610e3557600083815260209020610e359181019083015b61089491905b808211156116ca57600081556001016116b6565b509056fea165627a7a7230582024c8ccf95c83151d8e709451d6a6a137139753f0cc0546c8a1a6d048d339dc1f0029a66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688`

// DeployTribeChief_0_0_7 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_7 to it.
func DeployTribeChief_0_0_7(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address, _epoch *big.Int, _signerLimit *big.Int, _volunteerLimit *big.Int) (common.Address, *types.Transaction, *TribeChief_0_0_7, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_7ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_7Bin), backend, genesisSigners, _epoch, _signerLimit, _volunteerLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_0_0_7{TribeChief_0_0_7Caller: TribeChief_0_0_7Caller{contract: contract}, TribeChief_0_0_7Transactor: TribeChief_0_0_7Transactor{contract: contract}}, nil
}

// TribeChief_0_0_7 is an auto generated Go binding around an Ethereum contract.
type TribeChief_0_0_7 struct {
	TribeChief_0_0_7Caller     // Read-only binding to the contract
	TribeChief_0_0_7Transactor // Write-only binding to the contract
}

// TribeChief_0_0_7Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_0_0_7Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_7Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_0_0_7Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_7Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_0_0_7Session struct {
	Contract     *TribeChief_0_0_7       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_0_0_7CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_0_0_7CallerSession struct {
	Contract *TribeChief_0_0_7Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_0_0_7TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_0_0_7TransactorSession struct {
	Contract     *TribeChief_0_0_7Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_0_0_7Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_0_0_7Raw struct {
	Contract *TribeChief_0_0_7 // Generic contract binding to access the raw methods on
}

// TribeChief_0_0_7CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_0_0_7CallerRaw struct {
	Contract *TribeChief_0_0_7Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_0_0_7TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_0_0_7TransactorRaw struct {
	Contract *TribeChief_0_0_7Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_0_0_7 creates a new instance of TribeChief_0_0_7, bound to a specific deployed contract.
func NewTribeChief_0_0_7(address common.Address, backend bind.ContractBackend) (*TribeChief_0_0_7, error) {
	contract, err := bindTribeChief_0_0_7(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_7{TribeChief_0_0_7Caller: TribeChief_0_0_7Caller{contract: contract}, TribeChief_0_0_7Transactor: TribeChief_0_0_7Transactor{contract: contract}}, nil
}

// NewTribeChief_0_0_7Caller creates a new read-only instance of TribeChief_0_0_7, bound to a specific deployed contract.
func NewTribeChief_0_0_7Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_0_0_7Caller, error) {
	contract, err := bindTribeChief_0_0_7(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_7Caller{contract: contract}, nil
}

// NewTribeChief_0_0_7Transactor creates a new write-only instance of TribeChief_0_0_7, bound to a specific deployed contract.
func NewTribeChief_0_0_7Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_0_0_7Transactor, error) {
	contract, err := bindTribeChief_0_0_7(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_7Transactor{contract: contract}, nil
}

// bindTribeChief_0_0_7 binds a generic wrapper to an already deployed contract.
func bindTribeChief_0_0_7(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_7ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_7 *TribeChief_0_0_7Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_7.Contract.TribeChief_0_0_7Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_7 *TribeChief_0_0_7Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.TribeChief_0_0_7Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_7 *TribeChief_0_0_7Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.TribeChief_0_0_7Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_7.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_7 *TribeChief_0_0_7TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_7 *TribeChief_0_0_7TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.contract.Transact(opts, method, params...)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) FilterVolunteer(opts *bind.CallOptsWithNumber, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "filterVolunteer", volunteers)
	return *ret0, err
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_0_0_7.Contract.FilterVolunteer(&_TribeChief_0_0_7.CallOpts, volunteers)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_0_0_7.Contract.FilterVolunteer(&_TribeChief_0_0_7.CallOpts, volunteers)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetEpoch(&_TribeChief_0_0_7.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetEpoch(&_TribeChief_0_0_7.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetSignerLimit(&_TribeChief_0_0_7.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetSignerLimit(&_TribeChief_0_0_7.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	ret := new(struct {
		SignerList     []common.Address
		BlackList      []common.Address
		ScoreList      []*big.Int
		NumberList     []*big.Int
		TotalVolunteer *big.Int
		Number         *big.Int
	})
	out := ret
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_0_0_7.Contract.GetStatus(&_TribeChief_0_0_7.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_0_0_7.Contract.GetStatus(&_TribeChief_0_0_7.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetVolunteerLimit(&_TribeChief_0_0_7.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_7.Contract.GetVolunteerLimit(&_TribeChief_0_0_7.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) GetVolunteers(opts *bind.CallOptsWithNumber) (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		WeightList    []*big.Int
		Length        *big.Int
	})
	out := ret
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "getVolunteers")
	return *ret, err
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_0_0_7.Contract.GetVolunteers(&_TribeChief_0_0_7.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_0_0_7.Contract.GetVolunteers(&_TribeChief_0_0_7.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_0_0_7.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) Version() (string, error) {
	return _TribeChief_0_0_7.Contract.Version(&_TribeChief_0_0_7.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_7 *TribeChief_0_0_7CallerSession) Version() (string, error) {
	return _TribeChief_0_0_7.Contract.Version(&_TribeChief_0_0_7.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_7 *TribeChief_0_0_7Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_7.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_7 *TribeChief_0_0_7Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.Update(&_TribeChief_0_0_7.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_7 *TribeChief_0_0_7TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_7.Contract.Update(&_TribeChief_0_0_7.TransactOpts, volunteer)
}
