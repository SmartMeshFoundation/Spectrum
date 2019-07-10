// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chieflib

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
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
	return address, tx, &Chief{ChiefCaller: ChiefCaller{contract: contract}, ChiefTransactor: ChiefTransactor{contract: contract}, ChiefFilterer: ChiefFilterer{contract: contract}}, nil
}

// Chief is an auto generated Go binding around an Ethereum contract.
type Chief struct {
	ChiefCaller     // Read-only binding to the contract
	ChiefTransactor // Write-only binding to the contract
	ChiefFilterer   // Log filterer for contract events
}

// ChiefCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChiefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiefSession struct {
	Contract     *Chief            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChiefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiefCallerSession struct {
	Contract *ChiefCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
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
	contract, err := bindChief(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chief{ChiefCaller: ChiefCaller{contract: contract}, ChiefTransactor: ChiefTransactor{contract: contract}, ChiefFilterer: ChiefFilterer{contract: contract}}, nil
}

// NewChiefCaller creates a new read-only instance of Chief, bound to a specific deployed contract.
func NewChiefCaller(address common.Address, caller bind.ContractCaller) (*ChiefCaller, error) {
	contract, err := bindChief(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefCaller{contract: contract}, nil
}

// NewChiefTransactor creates a new write-only instance of Chief, bound to a specific deployed contract.
func NewChiefTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiefTransactor, error) {
	contract, err := bindChief(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefTransactor{contract: contract}, nil
}

// NewChiefFilterer creates a new log filterer instance of Chief, bound to a specific deployed contract.
func NewChiefFilterer(address common.Address, filterer bind.ContractFilterer) (*ChiefFilterer, error) {
	contract, err := bindChief(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChiefFilterer{contract: contract}, nil
}

// bindChief binds a generic wrapper to an already deployed contract.
func bindChief(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chief *ChiefRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Chief.Contract.ChiefCaller.contract.Call(opts, result, method, params...)
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
func (_Chief *ChiefCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Chief.Contract.contract.Call(opts, result, method, params...)
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
func (_Chief *ChiefCaller) FilterVolunteer(opts *bind.CallOpts, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Chief.contract.Call(opts, out, "filterVolunteer", volunteers)
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
func (_Chief *ChiefCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.Call(opts, out, "getEpoch")
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
func (_Chief *ChiefCaller) GetSignerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.Call(opts, out, "getSignerLimit")
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
func (_Chief *ChiefCaller) GetStatus(opts *bind.CallOpts) (struct {
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
	err := _Chief.contract.Call(opts, out, "getStatus")
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
func (_Chief *ChiefCaller) GetVolunteerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chief.contract.Call(opts, out, "getVolunteerLimit")
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
func (_Chief *ChiefCaller) GetVolunteers(opts *bind.CallOpts) (struct {
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
	err := _Chief.contract.Call(opts, out, "getVolunteers")
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
func (_Chief *ChiefCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Chief.contract.Call(opts, out, "version")
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

// ChiefBaseABI is the input ABI used to generate the binding from.
const ChiefBaseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_leaderLimit\",\"type\":\"uint256\"}],\"name\":\"setLeaderLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"volunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeLeaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocDelLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_tribe\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pocCleanBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddLockList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pocGetBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vsn\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"pocChangeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"appendLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pocAddr\",\"type\":\"address\"},{\"name\":\"tribeAddr\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeLeaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ChiefBaseBin is the compiled bytecode used for deploying new contracts.
const ChiefBaseBin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a090815261003e91600091906100b1565b5060058055600360065561181b600755600260085534801561005f57600080fd5b5060405160408061104d8339810180604052604081101561007f57600080fd5b50805160209091015160028054600160a060020a0319163317905560079190915560068190556000190160085561014c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f257805160ff191683800117855561011f565b8280016001018555821561011f579182015b8281111561011f578251825591602001919060010190610104565b5061012b92915061012f565b5090565b61014991905b8082111561012b5760008155600101610135565b90565b610ef28061015b6000396000f3fe608060405234801561001057600080fd5b50600436106101c6576000357c010000000000000000000000000000000000000000000000000000000090048063900cf0cf11610116578063bfe3a272116100b4578063db512e851161008e578063db512e851461049c578063ed77db3f146104d6578063f09a4016146104f3578063f0d5581714610521576101c6565b8063bfe3a27214610448578063c8ba1fd81461046e578063d0f2617114610494576101c6565b8063b175de17116100f0578063b175de171461033f578063b2bdfa7b14610397578063bbf2ef671461039f578063bf8fb6141461041c576101c6565b8063900cf0cf146102eb5780639037f182146102f3578063a9cfa49c14610319576101c6565b806340e2f99111610183578063609bf73b1161015d578063609bf73b146102af5780636c151173146102d35780636da50830146102db5780637216276b146102e3576101c6565b806340e2f991146102795780634bb19dac1461028157806357f607a1146102a7576101c6565b80630e7279e4146101cb57806313af4035146101ea5780631dd47d3d1461021057806326b249c81461024f578063282f78f2146102695780634067f0c814610271575b600080fd5b6101e8600480360360208110156101e157600080fd5b5035610529565b005b6101e86004803603602081101561020057600080fd5b5035600160a060020a0316610545565b6102366004803603602081101561022657600080fd5b5035600160a060020a031661058b565b60408051600092830b90920b8252519081900360200190f35b610257610641565b60408051918252519081900360200190f35b610257610647565b61025761064e565b610257610654565b6102366004803603602081101561029757600080fd5b5035600160a060020a031661065a565b6102576106dc565b6102b76106e2565b60408051600160a060020a039092168252519081900360200190f35b6102576106f1565b6102576106f7565b6101e86106fd565b61025761079a565b6101e86004803603602081101561030957600080fd5b5035600160a060020a03166107a0565b6101e86004803603602081101561032f57600080fd5b5035600160a060020a03166108b3565b61034761094c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561038357818101518382015260200161036b565b505050509050019250505060405180910390f35b6102b7610a3a565b6103a7610a49565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103e15781810151838201526020016103c9565b50505050905090810190601f16801561040e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e86004803603604081101561043257600080fd5b50600160a060020a038135169060200135610ad7565b6101e86004803603602081101561045e57600080fd5b5035600160a060020a0316610b78565b6101e86004803603602081101561048457600080fd5b5035600160a060020a0316610c46565b6102b7610cc4565b6104c2600480360360208110156104b257600080fd5b5035600160a060020a0316610cd3565b604080519115158252519081900360200190f35b6102b7600480360360208110156104ec57600080fd5b5035610d3d565b6101e86004803603604081101561050957600080fd5b50600160a060020a0381358116916020013516610d65565b610347610e1d565b600254600160a060020a0316331461054057600080fd5b600555565b600254600160a060020a0316331461055c57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600090600160a060020a031633146105a557600080fd5b600154604080517f2988fcfd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691632988fcfd9160248083019260209291908290030181600087803b15801561060d57600080fd5b505af1158015610621573d6000803e3d6000fd5b505050506040513d602081101561063757600080fd5b505190505b919050565b60085481565b6008545b90565b60075490565b60055490565b600354600090600160a060020a0316331461067457600080fd5b600154604080517f6fae4a5f000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691636fae4a5f9160248083019260209291908290030181600087803b15801561060d57600080fd5b60065481565b600354600160a060020a031681565b60055481565b60065490565b600354600160a060020a0316331461071457600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a0316633853d0a16040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561078057600080fd5b505af1158015610794573d6000803e3d6000fd5b50505050565b60075481565b600254600160a060020a031633146107b757600080fd5b600454600114156107c757600080fd5b60005b6004548110156108af5781600160a060020a03166004828154811015156107ed57fe5b600091825260209091200154600160a060020a031614156108a757805b6004546000190181101561088d57600480546001830190811061082957fe5b60009182526020909120015460048054600160a060020a03909216918390811061084f57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560010161080a565b50600480546000199283019201906108a59082610e7f565b505b6001016107ca565b5050565b600354600160a060020a031633146108ca57600080fd5b600154604080517f478e7d6e000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163478e7d6e91602480830192600092919082900301818387803b15801561093157600080fd5b505af1158015610945573d6000803e3d6000fd5b5050505050565b600154604080517f360b97b90000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163360b97b9916004808301926000929190829003018186803b1580156109aa57600080fd5b505afa1580156109be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156109e757600080fd5b8101908080516401000000008111156109ff57600080fd5b82016020810184811115610a1257600080fd5b8151856020820283011164010000000082111715610a2f57600080fd5b509094505050505090565b600254600160a060020a031681565b6000805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610acf5780601f10610aa457610100808354040283529160200191610acf565b820191906000526020600020905b815481529060010190602001808311610ab257829003601f168201915b505050505081565b600254600160a060020a03163314610aee57600080fd5b600154604080517f56df3db1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201859052915191909216916356df3db191604480830192600092919082900301818387803b158015610b5c57600080fd5b505af1158015610b70573d6000803e3d6000fd5b505050505050565b600254600160a060020a03163314610b8f57600080fd5b6005546004541015610c435760005b600454811015610be95781600160a060020a0316600482815481101515610bc157fe5b600091825260209091200154600160a060020a03161415610be157600080fd5b600101610b9e565b50600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600354600160a060020a03163314610c5d57600080fd5b600154604080517f31f23668000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916331f2366891602480830192600092919082900301818387803b15801561093157600080fd5b600254600160a060020a031690565b6000805b600454811015610d3457600160a060020a03831615801590610d1d57506004805482908110610d0257fe5b600091825260209091200154600160a060020a038481169116145b15610d2c57600191505061063c565b600101610cd7565b50600092915050565b6004805482908110610d4b57fe5b600091825260209091200154600160a060020a0316905081565b33600160a060020a03821614610d7a57600080fd5b600154600160a060020a0316158015610d9b5750600160a060020a03821615155b15610dc9576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b600160a060020a03811615801590610dea5750600354600160a060020a0316155b156108af5760038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790555050565b60606004805480602002602001604051908101604052809291908181526020018280548015610e7557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610e57575b5050505050905090565b815481835581811115610ea357600083815260209020610ea3918101908301610ea8565b505050565b61064b91905b80821115610ec25760008155600101610eae565b509056fea165627a7a723058201de76c5a398ef5d138769ec2652a73b2f884e34f39f0498d269b36cd21bbd83d0029`

// DeployChiefBase deploys a new Ethereum contract, binding an instance of ChiefBase to it.
func DeployChiefBase(auth *bind.TransactOpts, backend bind.ContractBackend, _epoch *big.Int, _signerLimit *big.Int) (common.Address, *types.Transaction, *ChiefBase, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChiefBaseBin), backend, _epoch, _signerLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChiefBase{ChiefBaseCaller: ChiefBaseCaller{contract: contract}, ChiefBaseTransactor: ChiefBaseTransactor{contract: contract}, ChiefBaseFilterer: ChiefBaseFilterer{contract: contract}}, nil
}

// ChiefBase is an auto generated Go binding around an Ethereum contract.
type ChiefBase struct {
	ChiefBaseCaller     // Read-only binding to the contract
	ChiefBaseTransactor // Write-only binding to the contract
	ChiefBaseFilterer   // Log filterer for contract events
}

// ChiefBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiefBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiefBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChiefBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiefBaseSession struct {
	Contract     *ChiefBase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChiefBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiefBaseCallerSession struct {
	Contract *ChiefBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChiefBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChiefBaseTransactorSession struct {
	Contract     *ChiefBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChiefBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChiefBaseRaw struct {
	Contract *ChiefBase // Generic contract binding to access the raw methods on
}

// ChiefBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChiefBaseCallerRaw struct {
	Contract *ChiefBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ChiefBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChiefBaseTransactorRaw struct {
	Contract *ChiefBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChiefBase creates a new instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBase(address common.Address, backend bind.ContractBackend) (*ChiefBase, error) {
	contract, err := bindChiefBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChiefBase{ChiefBaseCaller: ChiefBaseCaller{contract: contract}, ChiefBaseTransactor: ChiefBaseTransactor{contract: contract}, ChiefBaseFilterer: ChiefBaseFilterer{contract: contract}}, nil
}

// NewChiefBaseCaller creates a new read-only instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBaseCaller(address common.Address, caller bind.ContractCaller) (*ChiefBaseCaller, error) {
	contract, err := bindChiefBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefBaseCaller{contract: contract}, nil
}

// NewChiefBaseTransactor creates a new write-only instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiefBaseTransactor, error) {
	contract, err := bindChiefBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefBaseTransactor{contract: contract}, nil
}

// NewChiefBaseFilterer creates a new log filterer instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ChiefBaseFilterer, error) {
	contract, err := bindChiefBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChiefBaseFilterer{contract: contract}, nil
}

// bindChiefBase binds a generic wrapper to an already deployed contract.
func bindChiefBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiefBase *ChiefBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase.Contract.ChiefBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiefBase *ChiefBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase.Contract.ChiefBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChiefBase *ChiefBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChiefBase.Contract.ChiefBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiefBase *ChiefBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiefBase *ChiefBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChiefBase *ChiefBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChiefBase.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase *ChiefBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "_owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase *ChiefBaseSession) Owner() (common.Address, error) {
	return _ChiefBase.Contract.Owner(&_ChiefBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase *ChiefBaseCallerSession) Owner() (common.Address, error) {
	return _ChiefBase.Contract.Owner(&_ChiefBase.CallOpts)
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase *ChiefBaseCaller) Tribe(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "_tribe")
	return *ret0, err
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase *ChiefBaseSession) Tribe() (common.Address, error) {
	return _ChiefBase.Contract.Tribe(&_ChiefBase.CallOpts)
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase *ChiefBaseCallerSession) Tribe() (common.Address, error) {
	return _ChiefBase.Contract.Tribe(&_ChiefBase.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "epoch")
	return *ret0, err
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) Epoch() (*big.Int, error) {
	return _ChiefBase.Contract.Epoch(&_ChiefBase.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) Epoch() (*big.Int, error) {
	return _ChiefBase.Contract.Epoch(&_ChiefBase.CallOpts)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase *ChiefBaseCaller) IsLeader(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "isLeader", addr)
	return *ret0, err
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase *ChiefBaseSession) IsLeader(addr common.Address) (bool, error) {
	return _ChiefBase.Contract.IsLeader(&_ChiefBase.CallOpts, addr)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase *ChiefBaseCallerSession) IsLeader(addr common.Address) (bool, error) {
	return _ChiefBase.Contract.IsLeader(&_ChiefBase.CallOpts, addr)
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) LeaderLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "leaderLimit")
	return *ret0, err
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) LeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.LeaderLimit(&_ChiefBase.CallOpts)
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) LeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.LeaderLimit(&_ChiefBase.CallOpts)
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase *ChiefBaseCaller) LeaderList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "leaderList", arg0)
	return *ret0, err
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase *ChiefBaseSession) LeaderList(arg0 *big.Int) (common.Address, error) {
	return _ChiefBase.Contract.LeaderList(&_ChiefBase.CallOpts, arg0)
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase *ChiefBaseCallerSession) LeaderList(arg0 *big.Int) (common.Address, error) {
	return _ChiefBase.Contract.LeaderList(&_ChiefBase.CallOpts, arg0)
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase *ChiefBaseCaller) PocGetBlackList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "pocGetBlackList")
	return *ret0, err
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase *ChiefBaseSession) PocGetBlackList() ([]common.Address, error) {
	return _ChiefBase.Contract.PocGetBlackList(&_ChiefBase.CallOpts)
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase *ChiefBaseCallerSession) PocGetBlackList() ([]common.Address, error) {
	return _ChiefBase.Contract.PocGetBlackList(&_ChiefBase.CallOpts)
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) SignerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "signerLimit")
	return *ret0, err
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) SignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.SignerLimit(&_ChiefBase.CallOpts)
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) SignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.SignerLimit(&_ChiefBase.CallOpts)
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) TakeEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeEpoch")
	return *ret0, err
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) TakeEpoch() (*big.Int, error) {
	return _ChiefBase.Contract.TakeEpoch(&_ChiefBase.CallOpts)
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) TakeEpoch() (*big.Int, error) {
	return _ChiefBase.Contract.TakeEpoch(&_ChiefBase.CallOpts)
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) TakeLeaderLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeLeaderLimit")
	return *ret0, err
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) TakeLeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeLeaderLimit(&_ChiefBase.CallOpts)
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) TakeLeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeLeaderLimit(&_ChiefBase.CallOpts)
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseCaller) TakeLeaderList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeLeaderList")
	return *ret0, err
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseSession) TakeLeaderList() ([]common.Address, error) {
	return _ChiefBase.Contract.TakeLeaderList(&_ChiefBase.CallOpts)
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseCallerSession) TakeLeaderList() ([]common.Address, error) {
	return _ChiefBase.Contract.TakeLeaderList(&_ChiefBase.CallOpts)
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase *ChiefBaseCaller) TakeOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeOwner")
	return *ret0, err
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase *ChiefBaseSession) TakeOwner() (common.Address, error) {
	return _ChiefBase.Contract.TakeOwner(&_ChiefBase.CallOpts)
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase *ChiefBaseCallerSession) TakeOwner() (common.Address, error) {
	return _ChiefBase.Contract.TakeOwner(&_ChiefBase.CallOpts)
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) TakeSignerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeSignerLimit")
	return *ret0, err
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) TakeSignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeSignerLimit(&_ChiefBase.CallOpts)
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) TakeSignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeSignerLimit(&_ChiefBase.CallOpts)
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) TakeVolunteerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "takeVolunteerLimit")
	return *ret0, err
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) TakeVolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeVolunteerLimit(&_ChiefBase.CallOpts)
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) TakeVolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.TakeVolunteerLimit(&_ChiefBase.CallOpts)
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) VolunteerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "volunteerLimit")
	return *ret0, err
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) VolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.VolunteerLimit(&_ChiefBase.CallOpts)
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) VolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.VolunteerLimit(&_ChiefBase.CallOpts)
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase *ChiefBaseCaller) Vsn(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChiefBase.contract.Call(opts, out, "vsn")
	return *ret0, err
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase *ChiefBaseSession) Vsn() (string, error) {
	return _ChiefBase.Contract.Vsn(&_ChiefBase.CallOpts)
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase *ChiefBaseCallerSession) Vsn() (string, error) {
	return _ChiefBase.Contract.Vsn(&_ChiefBase.CallOpts)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase *ChiefBaseTransactor) AppendLeader(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "appendLeader", addr)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase *ChiefBaseSession) AppendLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.AppendLeader(&_ChiefBase.TransactOpts, addr)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) AppendLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.AppendLeader(&_ChiefBase.TransactOpts, addr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase *ChiefBaseTransactor) Init(opts *bind.TransactOpts, pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "init", pocAddr, tribeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase *ChiefBaseSession) Init(pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.Init(&_ChiefBase.TransactOpts, pocAddr, tribeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) Init(pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.Init(&_ChiefBase.TransactOpts, pocAddr, tribeAddr)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseTransactor) PocAddBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocAddBlackList", minerAddress)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseSession) PocAddBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddBlackList(&_ChiefBase.TransactOpts, minerAddress)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) PocAddBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddBlackList(&_ChiefBase.TransactOpts, minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseTransactor) PocAddLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocAddLockList", minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseSession) PocAddLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddLockList(&_ChiefBase.TransactOpts, minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) PocAddLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddLockList(&_ChiefBase.TransactOpts, minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseTransactor) PocAddStop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocAddStop", minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseSession) PocAddStop(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddStop(&_ChiefBase.TransactOpts, minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseTransactorSession) PocAddStop(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocAddStop(&_ChiefBase.TransactOpts, minerAddress)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase *ChiefBaseTransactor) PocChangeOwner(opts *bind.TransactOpts, newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocChangeOwner", newOwner, num)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase *ChiefBaseSession) PocChangeOwner(newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocChangeOwner(&_ChiefBase.TransactOpts, newOwner, num)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase *ChiefBaseTransactorSession) PocChangeOwner(newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocChangeOwner(&_ChiefBase.TransactOpts, newOwner, num)
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase *ChiefBaseTransactor) PocCleanBlackList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocCleanBlackList")
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase *ChiefBaseSession) PocCleanBlackList() (*types.Transaction, error) {
	return _ChiefBase.Contract.PocCleanBlackList(&_ChiefBase.TransactOpts)
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase *ChiefBaseTransactorSession) PocCleanBlackList() (*types.Transaction, error) {
	return _ChiefBase.Contract.PocCleanBlackList(&_ChiefBase.TransactOpts)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseTransactor) PocDelLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "pocDelLockList", minerAddress)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseSession) PocDelLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocDelLockList(&_ChiefBase.TransactOpts, minerAddress)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase *ChiefBaseTransactorSession) PocDelLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.PocDelLockList(&_ChiefBase.TransactOpts, minerAddress)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase *ChiefBaseTransactor) RemoveLeader(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "removeLeader", addr)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase *ChiefBaseSession) RemoveLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.RemoveLeader(&_ChiefBase.TransactOpts, addr)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) RemoveLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.RemoveLeader(&_ChiefBase.TransactOpts, addr)
}

// SetLeaderLimit is a paid mutator transaction binding the contract method 0x0e7279e4.
//
// Solidity: function setLeaderLimit(_leaderLimit uint256) returns()
func (_ChiefBase *ChiefBaseTransactor) SetLeaderLimit(opts *bind.TransactOpts, _leaderLimit *big.Int) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "setLeaderLimit", _leaderLimit)
}

// SetLeaderLimit is a paid mutator transaction binding the contract method 0x0e7279e4.
//
// Solidity: function setLeaderLimit(_leaderLimit uint256) returns()
func (_ChiefBase *ChiefBaseSession) SetLeaderLimit(_leaderLimit *big.Int) (*types.Transaction, error) {
	return _ChiefBase.Contract.SetLeaderLimit(&_ChiefBase.TransactOpts, _leaderLimit)
}

// SetLeaderLimit is a paid mutator transaction binding the contract method 0x0e7279e4.
//
// Solidity: function setLeaderLimit(_leaderLimit uint256) returns()
func (_ChiefBase *ChiefBaseTransactorSession) SetLeaderLimit(_leaderLimit *big.Int) (*types.Transaction, error) {
	return _ChiefBase.Contract.SetLeaderLimit(&_ChiefBase.TransactOpts, _leaderLimit)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase *ChiefBaseTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase *ChiefBaseSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.SetOwner(&_ChiefBase.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.SetOwner(&_ChiefBase.TransactOpts, newOwner)
}

// POCABI is the input ABI used to generate the binding from.
const POCABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"minerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPopBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawWaitNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPushBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerEmptyBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initMinDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPushLockList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwnerEffectiveNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNormalList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"},{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"withdrawSurplus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPopLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStopList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"blackAndLockStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositHalveLimitNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_waitNumber\",\"type\":\"uint256\"},{\"name\":\"_limitNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"Start\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// POCBin is the compiled bytecode used for deploying new contracts.
const POCBin = `0x608060405234801561001057600080fd5b5060405160808061207c8339810180604052608081101561003057600080fd5b50805160208201516040830151606090930151600a91909155600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a03909316600160a060020a031990931692909217909155600d91909155600c5543600b55611fc5806100b76000396000f3fe60806040526004361061019e576000357c01000000000000000000000000000000000000000000000000000000009004806353ed5143116100ee5780636fae4a5f116100a7578063e7b14a9311610081578063e7b14a9314610688578063ed1eeaf1146106bb578063f2a558f9146106d0578063f8864cfe146106fc5761019e565b80636fae4a5f1461060d5780639c52ade214610640578063dd0b281e146106555761019e565b806353ed51431461040c57806356df3db114610544578063581003701461057d578063645006ca14610592578063656cb0bc146105a757806365e73e32146105da5761019e565b8063360b97b91161015b578063478e7d6e11610135578063478e7d6e1461037c5780634b58f36e146103af57806351393ce3146103c457806351cff8d9146103d95761019e565b8063360b97b9146102ed5780633853d0a1146103525780633913468e146103675761019e565b80630c58ed8e146101a35780631103ec96146101fd578063147049701461024957806321615f6e146102705780632988fcfd1461028557806331f23668146102b8575b600080fd5b3480156101af57600080fd5b506101d6600480360360208110156101c657600080fd5b5035600160a060020a0316610711565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b34801561020957600080fd5b506102306004803603602081101561022057600080fd5b5035600160a060020a031661073e565b60408051600092830b90920b8252519081900360200190f35b34801561025557600080fd5b5061025e610906565b60408051918252519081900360200190f35b34801561027c57600080fd5b5061025e61090c565b34801561029157600080fd5b50610230600480360360208110156102a857600080fd5b5035600160a060020a0316610912565b3480156102c457600080fd5b506102eb600480360360208110156102db57600080fd5b5035600160a060020a03166109a6565b005b3480156102f957600080fd5b50610302610ae5565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561033e578181015183820152602001610326565b505050509050019250505060405180910390f35b34801561035e57600080fd5b506102eb610b48565b34801561037357600080fd5b5061025e610c25565b34801561038857600080fd5b506102eb6004803603602081101561039f57600080fd5b5035600160a060020a0316610c2b565b3480156103bb57600080fd5b5061025e610d68565b3480156103d057600080fd5b50610302610d6e565b3480156103e557600080fd5b506102eb600480360360208110156103fc57600080fd5b5035600160a060020a0316610dce565b34801561041857600080fd5b50610421610fe2565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b8381101561046d578181015183820152602001610455565b50505050905001858103845288818151815260200191508051906020019060200280838360005b838110156104ac578181015183820152602001610494565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156104eb5781810151838201526020016104d3565b50505050905001858103825286818151815260200191508051906020019060200280838360005b8381101561052a578181015183820152602001610512565b505050509050019850505050505050505060405180910390f35b34801561055057600080fd5b506102eb6004803603604081101561056757600080fd5b50600160a060020a03813516906020013561139c565b34801561058957600080fd5b506103026114c4565b34801561059e57600080fd5b5061025e611524565b3480156105b357600080fd5b506102eb600480360360208110156105ca57600080fd5b5035600160a060020a031661155f565b3480156105e657600080fd5b506102eb600480360360208110156105fd57600080fd5b5035600160a060020a03166115a2565b34801561061957600080fd5b506102306004803603602081101561063057600080fd5b5035600160a060020a0316611683565b34801561064c57600080fd5b50610302611818565b34801561066157600080fd5b506102eb6004803603602081101561067857600080fd5b5035600160a060020a0316611878565b34801561069457600080fd5b50610230600480360360208110156106ab57600080fd5b5035600160a060020a03166118a7565b3480156106c757600080fd5b5061030261195d565b6102eb600480360360608110156106e657600080fd5b508035906020810135906040013560ff166119bd565b34801561070857600080fd5b5061025e611b7f565b600160a060020a039081166000908152602081905260409020600181015460028201549154909391921690565b600980546000918290600019830183811061075557fe5b600091825260209091200154600160a060020a0316905060018211801561077d5750600e5443105b156107ac5760098054600119840190811061079457fe5b600091825260209091200154600160a060020a031690505b33600160a060020a038216146107c157600080fd5b600160a060020a03841660009081526006602052604090205460055460011180610817575084600160a060020a03166005828154811015156107ff57fe5b600091825260209091200154600160a060020a031614155b15610827576000199350506108ff565b60058054600019810190811061083957fe5b60009182526020909120015460058054600160a060020a03909216918390811061085f57fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560058054600019019061089b9082611f5c565b50600160a060020a0385166000908152600660205260408120556005548110156108f95780600660006005848154811015156108d357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b60019350505b5050919050565b600d5481565b600b5481565b600980546000918290600019830183811061092957fe5b600091825260209091200154600160a060020a031690506001821180156109515750600e5443105b156109805760098054600119840190811061096857fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461099557600080fd5b61099e84611b85565b949350505050565b600980549060009060001983018381106109bc57fe5b600091825260209091200154600160a060020a031690506001821180156109e45750600e5443105b15610a13576009805460011984019081106109fb57fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610a2857600080fd5b6005546000108015610a745750600160a060020a038316600081815260066020526040902054600580549091908110610a5d57fe5b600091825260209091200154600160a060020a0316145b15610a7e57610ae0565b600580546001810182557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018054600160a060020a031916600160a060020a038616908117909155905460009182526006602052604090912060001990910190555b505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610b3d57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b1f575b505050505090505b90565b60098054906000906000198301838110610b5e57fe5b600091825260209091200154600160a060020a03169050600182118015610b865750600e5443105b15610bb557600980546001198401908110610b9d57fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610bca57600080fd5b60005b600554811015610c175760066000600583815481101515610bea57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600101610bcd565b506000610ae0600582611f5c565b600a5481565b60098054906000906000198301838110610c4157fe5b600091825260209091200154600160a060020a03169050600182118015610c695750600e5443105b15610c9857600980546001198401908110610c8057fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610cad57600080fd5b6007546000108015610cf95750600160a060020a038316600081815260086020526040902054600780549091908110610ce257fe5b600091825260209091200154600160a060020a0316145b15610d0357610ae0565b5050600780546001810182557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018054600160a060020a031916600160a060020a03939093169283179055546000918252600860205260409091206000199091019055565b600e5481565b60606001805480602002602001604051908101604052809291908181526020018280548015610b3d57602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610b1f575050505050905090565b600160a060020a03818116600090815260208190526040902054163314610df457600080fd5b600160a060020a03811660009081526020819052604081206002015411610e1a57600080fd5b600d54600160a060020a038216600090815260208190526040902060020154430311610e4557600080fd5b610e4e816118a7565b60000b15610e5b57600080fd5b600160a060020a0381166000908152602081815260408083206001810180548254600160a060020a0319168355908590556002909101849055600490925290912054600380546000198101908110610eaf57fe5b60009182526020909120015460038054600160a060020a039092169183908110610ed557fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600380546000190190610f119082611f5c565b50600160a060020a038316600090815260046020526040812055600354811015610f6f578060046000600384815481101515610f4957fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b6040805183815290513391600160a060020a038616917f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9181900360200190a3604051339083156108fc029084906000818181858888f19350505050158015610fdc573d6000803e3d6000fd5b50505050565b600354600154604080518284018082526020808202830101909252606093849384938493849082801561101f578160200160208202803883390190505b50905060608260405190808252806020026020018201604052801561104e578160200160208202803883390190505b50905060608360405190808252806020026020018201604052801561107d578160200160208202803883390190505b5090506060846040519080825280602002602001820160405280156110ac578160200160208202803883390190505b50905060005b878110156112185760038054829081106110c857fe5b6000918252602090912001548551600160a060020a03909116908690839081106110ee57fe5b600160a060020a039092166020928302909101909101526003805460009182918490811061111857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154845185908390811061114e57fe5b602090810290910101526003805460009182918490811061116b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206002015483518490839081106111a157fe5b60209081029091010152600380546000918291849081106111be57fe5b6000918252602080832090910154600160a060020a03908116845290830193909352604090910190205483519116908390839081106111f957fe5b600160a060020a039092166020928302909101909101526001016110b2565b5060005b8681101561138a57600180548290811061123257fe5b6000918252602090912001548551600160a060020a039091169086908a840190811061125a57fe5b600160a060020a039092166020928302909101909101526001805460009182918490811061128457fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154845185908a84019081106112bc57fe5b60209081029091010152600180546000918291849081106112d957fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060020154835184908a840190811061131157fe5b602090810290910101526001805460009182918490811061132e57fe5b6000918252602080832090910154600160a060020a039081168452908301939093526040909101902054835191169083908a840190811061136b57fe5b600160a060020a0390921660209283029091019091015260010161121c565b50929a91995097509095509350505050565b600980549060009060001983018381106113b257fe5b600091825260209091200154600160a060020a031690506001821180156113da5750600e5443105b15611409576009805460011984019081106113f157fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461141e57600080fd5b600e544310156114705760098054859190600019810190811061143d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506114bc565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a031916600160a060020a0386161790555b5050600e5550565b60606009805480602002602001604051908101604052809291908181526020018280548015610b3d57602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610b1f575050505050905090565b600080600c54600b54430381151561153857fe5b0490506003811115611548575060035b8060020a600a5481151561155857fe5b0491505090565b600160a060020a0381811660009081526020819052604090205416331461158557600080fd5b61158e81611b85565b60000b600114151561159f57600080fd5b50565b600160a060020a038181166000908152602081905260409020541633146115c857600080fd5b60006115d2611524565b600160a060020a03831660009081526020819052604081206001015491909103915081116115ff57600080fd5b600160a060020a0382166000818152602081815260409182902060010180548590039055815184815291513393927f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb92908290030190a3604051339082156108fc029083906000818181858888f19350505050158015610ae0573d6000803e3d6000fd5b600980546000918290600019830183811061169a57fe5b600091825260209091200154600160a060020a031690506001821180156116c25750600e5443105b156116f1576009805460011984019081106116d957fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461170657600080fd5b600160a060020a0384166000908152600860205260409020546007546001118061175c575084600160a060020a031660078281548110151561174457fe5b600091825260209091200154600160a060020a031614155b1561176c576000199350506108ff565b60078054600019810190811061177e57fe5b60009182526020909120015460078054600160a060020a0390921691839081106117a457fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556007805460001901906117e09082611f5c565b50600160a060020a0385166000908152600860205260408120556007548110156108f95780600860006007848154811015156108d357fe5b60606003805480602002602001604051908101604052809291908181526020018280548015610b3d57602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610b1f575050505050905090565b600160a060020a0381811660009081526020819052604090205416331461189e57600080fd5b61158e81611d62565b6005546000908190811080156118f75750600160a060020a0383166000818152600660205260409020546005805490919081106118e057fe5b600091825260209091200154600160a060020a0316145b15611900576001015b600754600010801561194c5750600160a060020a03831660008181526008602052604090205460078054909190811061193557fe5b600091825260209091200154600160a060020a0316145b15611955576002015b90505b919050565b60606007805480602002602001604051908101604052809291908181526020018280548015610b3d57602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610b1f575050505050905090565b6119c5611524565b3410156119d157600080fd5b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff861660748501526094840188905260b484018790529351909260019260d4808301939192601f198301929081900390910190855afa158015611a60573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a0381161515611a8257600080fd5b600160a060020a03811660009081526020819052604090206001015415611aa857600080fd5b60408051606081018252338082523460208084018281526000858701818152600160a060020a03898116808452838652898420985189549216600160a060020a031992831617895593516001808a01919091559151600298890155815480830183557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180549091168417905554958352869020600019959095019094558451918252935191937f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6292918290030190a35050505050565b600c5481565b600160a060020a0381166000908152602081905260408120600101541515611bb05750600019611958565b600160a060020a03821660009081526020819052604090206002015415611bd957506000611958565b600160a060020a03821660009081526020818152604080832043600291820155909152902054600180546000198101908110611c1157fe5b60009182526020909120015460018054600160a060020a039092169183908110611c3757fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600180546000190190611c739082611f5c565b50600160a060020a038316600090815260026020526040812055600154811015611cd1578060026000600184815481101515611cab57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b600380546001810182557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a031916600160a060020a038616908117909155905460008281526004602052604080822060001990930190925590517f5f8b016b80ca3a499d7fc818528ee3c8658d421489a9efdf1d365c4241ddd8839190a250600192915050565b600160a060020a0381166000908152602081905260408120600101541515611d8d5750600019611958565b6000611d98836118a7565b60000b1315611da957506000611958565b600160a060020a0382166000908152602081905260409020600201541515611dd357506000611958565b600160a060020a0382166000908152602081815260408083206002018390556004909152902054600380546000198101908110611e0c57fe5b60009182526020909120015460038054600160a060020a039092169183908110611e3257fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600380546000190190611e6e9082611f5c565b50600160a060020a038316600090815260046020526040812055600354811015611ecc578060046000600384815481101515611ea657fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a031916600160a060020a038616908117909155905460008281526002602052604080822060001990930190925590517f84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a9659190a250600192915050565b815481835581811115610ae057600083815260209020610ae0918101908301610b4591905b80821115611f955760008155600101611f81565b509056fea165627a7a72305820c0c127e8251a1e682e30a9d11be2612f44eb952bf8ad54def69d14edb207d51e0029`

// DeployPOC deploys a new Ethereum contract, binding an instance of POC to it.
func DeployPOC(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address, _amount *big.Int, _waitNumber *big.Int, _limitNumber *big.Int) (common.Address, *types.Transaction, *POC, error) {
	parsed, err := abi.JSON(strings.NewReader(POCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(POCBin), backend, _addr, _amount, _waitNumber, _limitNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &POC{POCCaller: POCCaller{contract: contract}, POCTransactor: POCTransactor{contract: contract}, POCFilterer: POCFilterer{contract: contract}}, nil
}

// POC is an auto generated Go binding around an Ethereum contract.
type POC struct {
	POCCaller     // Read-only binding to the contract
	POCTransactor // Write-only binding to the contract
	POCFilterer   // Log filterer for contract events
}

// POCCaller is an auto generated read-only Go binding around an Ethereum contract.
type POCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type POCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type POCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type POCSession struct {
	Contract     *POC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// POCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type POCCallerSession struct {
	Contract *POCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// POCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type POCTransactorSession struct {
	Contract     *POCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// POCRaw is an auto generated low-level Go binding around an Ethereum contract.
type POCRaw struct {
	Contract *POC // Generic contract binding to access the raw methods on
}

// POCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type POCCallerRaw struct {
	Contract *POCCaller // Generic read-only contract binding to access the raw methods on
}

// POCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type POCTransactorRaw struct {
	Contract *POCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPOC creates a new instance of POC, bound to a specific deployed contract.
func NewPOC(address common.Address, backend bind.ContractBackend) (*POC, error) {
	contract, err := bindPOC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &POC{POCCaller: POCCaller{contract: contract}, POCTransactor: POCTransactor{contract: contract}, POCFilterer: POCFilterer{contract: contract}}, nil
}

// NewPOCCaller creates a new read-only instance of POC, bound to a specific deployed contract.
func NewPOCCaller(address common.Address, caller bind.ContractCaller) (*POCCaller, error) {
	contract, err := bindPOC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &POCCaller{contract: contract}, nil
}

// NewPOCTransactor creates a new write-only instance of POC, bound to a specific deployed contract.
func NewPOCTransactor(address common.Address, transactor bind.ContractTransactor) (*POCTransactor, error) {
	contract, err := bindPOC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &POCTransactor{contract: contract}, nil
}

// NewPOCFilterer creates a new log filterer instance of POC, bound to a specific deployed contract.
func NewPOCFilterer(address common.Address, filterer bind.ContractFilterer) (*POCFilterer, error) {
	contract, err := bindPOC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &POCFilterer{contract: contract}, nil
}

// bindPOC binds a generic wrapper to an already deployed contract.
func bindPOC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(POCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POC *POCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _POC.Contract.POCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POC *POCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC.Contract.POCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POC *POCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POC.Contract.POCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POC *POCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _POC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POC *POCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POC *POCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POC.Contract.contract.Transact(opts, method, params...)
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC *POCCaller) BlackAndLockStatus(opts *bind.CallOpts, minerAddress common.Address) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "blackAndLockStatus", minerAddress)
	return *ret0, err
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC *POCSession) BlackAndLockStatus(minerAddress common.Address) (int8, error) {
	return _POC.Contract.BlackAndLockStatus(&_POC.CallOpts, minerAddress)
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC *POCCallerSession) BlackAndLockStatus(minerAddress common.Address) (int8, error) {
	return _POC.Contract.BlackAndLockStatus(&_POC.CallOpts, minerAddress)
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC *POCCaller) DepositHalveLimitNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "depositHalveLimitNumber")
	return *ret0, err
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC *POCSession) DepositHalveLimitNumber() (*big.Int, error) {
	return _POC.Contract.DepositHalveLimitNumber(&_POC.CallOpts)
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC *POCCallerSession) DepositHalveLimitNumber() (*big.Int, error) {
	return _POC.Contract.DepositHalveLimitNumber(&_POC.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[])
func (_POC *POCCaller) GetAll(opts *bind.CallOpts) ([]common.Address, []*big.Int, []*big.Int, []common.Address, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
		ret2 = new([]*big.Int)
		ret3 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _POC.contract.Call(opts, out, "getAll")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[])
func (_POC *POCSession) GetAll() ([]common.Address, []*big.Int, []*big.Int, []common.Address, error) {
	return _POC.Contract.GetAll(&_POC.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[])
func (_POC *POCCallerSession) GetAll() ([]common.Address, []*big.Int, []*big.Int, []common.Address, error) {
	return _POC.Contract.GetAll(&_POC.CallOpts)
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC *POCCaller) GetBlackList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "getBlackList")
	return *ret0, err
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC *POCSession) GetBlackList() ([]common.Address, error) {
	return _POC.Contract.GetBlackList(&_POC.CallOpts)
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC *POCCallerSession) GetBlackList() ([]common.Address, error) {
	return _POC.Contract.GetBlackList(&_POC.CallOpts)
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC *POCCaller) GetLockList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "getLockList")
	return *ret0, err
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC *POCSession) GetLockList() ([]common.Address, error) {
	return _POC.Contract.GetLockList(&_POC.CallOpts)
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC *POCCallerSession) GetLockList() ([]common.Address, error) {
	return _POC.Contract.GetLockList(&_POC.CallOpts)
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC *POCCaller) GetNormalList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "getNormalList")
	return *ret0, err
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC *POCSession) GetNormalList() ([]common.Address, error) {
	return _POC.Contract.GetNormalList(&_POC.CallOpts)
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC *POCCallerSession) GetNormalList() ([]common.Address, error) {
	return _POC.Contract.GetNormalList(&_POC.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC *POCCaller) GetOwnerList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "getOwnerList")
	return *ret0, err
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC *POCSession) GetOwnerList() ([]common.Address, error) {
	return _POC.Contract.GetOwnerList(&_POC.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC *POCCallerSession) GetOwnerList() ([]common.Address, error) {
	return _POC.Contract.GetOwnerList(&_POC.CallOpts)
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC *POCCaller) GetStopList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "getStopList")
	return *ret0, err
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC *POCSession) GetStopList() ([]common.Address, error) {
	return _POC.Contract.GetStopList(&_POC.CallOpts)
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC *POCCallerSession) GetStopList() ([]common.Address, error) {
	return _POC.Contract.GetStopList(&_POC.CallOpts)
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC *POCCaller) InitBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "initBlockNumber")
	return *ret0, err
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC *POCSession) InitBlockNumber() (*big.Int, error) {
	return _POC.Contract.InitBlockNumber(&_POC.CallOpts)
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC *POCCallerSession) InitBlockNumber() (*big.Int, error) {
	return _POC.Contract.InitBlockNumber(&_POC.CallOpts)
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC *POCCaller) InitMinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "initMinDeposit")
	return *ret0, err
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC *POCSession) InitMinDeposit() (*big.Int, error) {
	return _POC.Contract.InitMinDeposit(&_POC.CallOpts)
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC *POCCallerSession) InitMinDeposit() (*big.Int, error) {
	return _POC.Contract.InitMinDeposit(&_POC.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC *POCCaller) MinDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "minDepositAmount")
	return *ret0, err
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC *POCSession) MinDepositAmount() (*big.Int, error) {
	return _POC.Contract.MinDepositAmount(&_POC.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC *POCCallerSession) MinDepositAmount() (*big.Int, error) {
	return _POC.Contract.MinDepositAmount(&_POC.CallOpts)
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC *POCCaller) MinerStatus(opts *bind.CallOpts, _addr common.Address) (*big.Int, *big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _POC.contract.Call(opts, out, "minerStatus", _addr)
	return *ret0, *ret1, *ret2, err
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC *POCSession) MinerStatus(_addr common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _POC.Contract.MinerStatus(&_POC.CallOpts, _addr)
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC *POCCallerSession) MinerStatus(_addr common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _POC.Contract.MinerStatus(&_POC.CallOpts, _addr)
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC *POCCaller) NewOwnerEffectiveNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "newOwnerEffectiveNumber")
	return *ret0, err
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC *POCSession) NewOwnerEffectiveNumber() (*big.Int, error) {
	return _POC.Contract.NewOwnerEffectiveNumber(&_POC.CallOpts)
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC *POCCallerSession) NewOwnerEffectiveNumber() (*big.Int, error) {
	return _POC.Contract.NewOwnerEffectiveNumber(&_POC.CallOpts)
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC *POCCaller) WithdrawWaitNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.Call(opts, out, "withdrawWaitNumber")
	return *ret0, err
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC *POCSession) WithdrawWaitNumber() (*big.Int, error) {
	return _POC.Contract.WithdrawWaitNumber(&_POC.CallOpts)
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC *POCCallerSession) WithdrawWaitNumber() (*big.Int, error) {
	return _POC.Contract.WithdrawWaitNumber(&_POC.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC *POCTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "changeOwner", _newOwner, _number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC *POCSession) ChangeOwner(_newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC.Contract.ChangeOwner(&_POC.TransactOpts, _newOwner, _number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC *POCTransactorSession) ChangeOwner(_newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC.Contract.ChangeOwner(&_POC.TransactOpts, _newOwner, _number)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC *POCTransactor) Deposit(opts *bind.TransactOpts, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "deposit", _r, _s, _v)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC *POCSession) Deposit(_r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC.Contract.Deposit(&_POC.TransactOpts, _r, _s, _v)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC *POCTransactorSession) Deposit(_r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC.Contract.Deposit(&_POC.TransactOpts, _r, _s, _v)
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC *POCTransactor) OwnerEmptyBlackList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerEmptyBlackList")
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC *POCSession) OwnerEmptyBlackList() (*types.Transaction, error) {
	return _POC.Contract.OwnerEmptyBlackList(&_POC.TransactOpts)
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC *POCTransactorSession) OwnerEmptyBlackList() (*types.Transaction, error) {
	return _POC.Contract.OwnerEmptyBlackList(&_POC.TransactOpts)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC *POCTransactor) OwnerPopBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerPopBlackList", minerAddress)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC *POCSession) OwnerPopBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPopBlackList(&_POC.TransactOpts, minerAddress)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC *POCTransactorSession) OwnerPopBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPopBlackList(&_POC.TransactOpts, minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC *POCTransactor) OwnerPopLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerPopLockList", minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC *POCSession) OwnerPopLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPopLockList(&_POC.TransactOpts, minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC *POCTransactorSession) OwnerPopLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPopLockList(&_POC.TransactOpts, minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC *POCTransactor) OwnerPushBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerPushBlackList", minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC *POCSession) OwnerPushBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPushBlackList(&_POC.TransactOpts, minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC *POCTransactorSession) OwnerPushBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPushBlackList(&_POC.TransactOpts, minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC *POCTransactor) OwnerPushLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerPushLockList", minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC *POCSession) OwnerPushLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPushLockList(&_POC.TransactOpts, minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC *POCTransactorSession) OwnerPushLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerPushLockList(&_POC.TransactOpts, minerAddress)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC *POCTransactor) OwnerStop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "ownerStop", minerAddress)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC *POCSession) OwnerStop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerStop(&_POC.TransactOpts, minerAddress)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC *POCTransactorSession) OwnerStop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.OwnerStop(&_POC.TransactOpts, minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC *POCTransactor) Start(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "start", minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC *POCSession) Start(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Start(&_POC.TransactOpts, minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC *POCTransactorSession) Start(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Start(&_POC.TransactOpts, minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC *POCTransactor) Stop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "stop", minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC *POCSession) Stop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Stop(&_POC.TransactOpts, minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC *POCTransactorSession) Stop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Stop(&_POC.TransactOpts, minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC *POCTransactor) Withdraw(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "withdraw", minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC *POCSession) Withdraw(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Withdraw(&_POC.TransactOpts, minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC *POCTransactorSession) Withdraw(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.Withdraw(&_POC.TransactOpts, minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC *POCTransactor) WithdrawSurplus(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "withdrawSurplus", minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC *POCSession) WithdrawSurplus(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.WithdrawSurplus(&_POC.TransactOpts, minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC *POCTransactorSession) WithdrawSurplus(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.WithdrawSurplus(&_POC.TransactOpts, minerAddress)
}

// POCDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the POC contract.
type POCDepositIterator struct {
	Event *POCDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POCDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POCDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POCDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POCDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POCDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POCDeposit represents a Deposit event raised by the POC contract.
type POCDeposit struct {
	Miner  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: e Deposit(miner indexed address, owner indexed address, amount uint256)
func (_POC *POCFilterer) FilterDeposit(opts *bind.FilterOpts, miner []common.Address, owner []common.Address) (*POCDepositIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _POC.contract.FilterLogs(opts, "Deposit", minerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &POCDepositIterator{contract: _POC.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: e Deposit(miner indexed address, owner indexed address, amount uint256)
func (_POC *POCFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *POCDeposit, miner []common.Address, owner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _POC.contract.WatchLogs(opts, "Deposit", minerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POCDeposit)
				if err := _POC.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// POCStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the POC contract.
type POCStartIterator struct {
	Event *POCStart // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POCStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POCStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POCStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POCStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POCStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POCStart represents a Start event raised by the POC contract.
type POCStart struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a965.
//
// Solidity: e Start(miner indexed address)
func (_POC *POCFilterer) FilterStart(opts *bind.FilterOpts, miner []common.Address) (*POCStartIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _POC.contract.FilterLogs(opts, "Start", minerRule)
	if err != nil {
		return nil, err
	}
	return &POCStartIterator{contract: _POC.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a965.
//
// Solidity: e Start(miner indexed address)
func (_POC *POCFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *POCStart, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _POC.contract.WatchLogs(opts, "Start", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POCStart)
				if err := _POC.contract.UnpackLog(event, "Start", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// POCStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the POC contract.
type POCStopIterator struct {
	Event *POCStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POCStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POCStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POCStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POCStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POCStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POCStop represents a Stop event raised by the POC contract.
type POCStop struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0x5f8b016b80ca3a499d7fc818528ee3c8658d421489a9efdf1d365c4241ddd883.
//
// Solidity: e Stop(miner indexed address)
func (_POC *POCFilterer) FilterStop(opts *bind.FilterOpts, miner []common.Address) (*POCStopIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _POC.contract.FilterLogs(opts, "Stop", minerRule)
	if err != nil {
		return nil, err
	}
	return &POCStopIterator{contract: _POC.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0x5f8b016b80ca3a499d7fc818528ee3c8658d421489a9efdf1d365c4241ddd883.
//
// Solidity: e Stop(miner indexed address)
func (_POC *POCFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *POCStop, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _POC.contract.WatchLogs(opts, "Stop", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POCStop)
				if err := _POC.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// POCWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the POC contract.
type POCWithdrawIterator struct {
	Event *POCWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POCWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POCWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POCWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POCWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POCWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POCWithdraw represents a Withdraw event raised by the POC contract.
type POCWithdraw struct {
	Miner  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: e Withdraw(miner indexed address, owner indexed address, amount uint256)
func (_POC *POCFilterer) FilterWithdraw(opts *bind.FilterOpts, miner []common.Address, owner []common.Address) (*POCWithdrawIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _POC.contract.FilterLogs(opts, "Withdraw", minerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &POCWithdrawIterator{contract: _POC.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: e Withdraw(miner indexed address, owner indexed address, amount uint256)
func (_POC *POCFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *POCWithdraw, miner []common.Address, owner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _POC.contract.WatchLogs(opts, "Withdraw", minerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POCWithdraw)
				if err := _POC.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TribeChief100ABI is the input ABI used to generate the binding from.
const TribeChief100ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextRoundSignerList\",\"outputs\":[{\"name\":\"nextRoundSignerList\",\"type\":\"address[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"baseAddress\",\"type\":\"address\"},{\"name\":\"pocAddress\",\"type\":\"address\"},{\"name\":\"startBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief100Bin is the compiled bytecode used for deploying new contracts.
const TribeChief100Bin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a0908152620000409160009190620003ad565b503480156200004e57600080fd5b50604051606080620018ae833981018060405260608110156200007057600080fd5b508051602082015160409283015160018054600160a060020a031916600160a060020a03808616919091179182905585517ff09a40160000000000000000000000000000000000000000000000000000000081528185166004820152306024820152955194959394929391169163f09a40169160448082019260009290919082900301818387803b1580156200010557600080fd5b505af11580156200011a573d6000803e3d6000fd5b5050506004828155600154604080517ff0d55817000000000000000000000000000000000000000000000000000000008152905160609450600160a060020a039092169263f0d55817928282019260009290829003018186803b1580156200018157600080fd5b505afa15801562000196573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015620001c057600080fd5b810190808051640100000000811115620001d957600080fd5b82016020810184811115620001ed57600080fd5b81518560208202830111640100000000821117156200020b57600080fd5b5050805190945060001092506200022491505057600080fd5b6001600560008360008151811015156200023a57fe5b90602001906020020151600160a060020a0316600160a060020a031681526020019081526020016000208190555060028160008151811015156200027a57fe5b6020908102919091018101518254600181018455600093845291909220018054600160a060020a031916600160a060020a039092169190911790556002545b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200032457600080fd5b505afa15801562000339573d6000803e3d6000fd5b505050506040513d60208110156200035057600080fd5b5051811015620003a25760028054600181810183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018054600160a060020a031916905501620002b9565b505050505062000452565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003f057805160ff191683800117855562000420565b8280016001018555821562000420579182015b828111156200042057825182559160200191906001019062000403565b506200042e92915062000432565b5090565b6200044f91905b808211156200042e576000815560010162000439565b90565b61144c80620004626000396000f3fe608060405234801561001057600080fd5b50600436106100a25760003560e060020a9004806373d53ec71161006a57806373d53ec7146103db578063757991a8146103e3578063961c5c7a146103fd578063d7ca4a1c14610405578063eb5c0011146104ad576100a2565b80631c1b8772146100a757806320c1a518146100cf5780634e69d560146101c257806353f197fe146102fb57806354fd4d501461035e575b600080fd5b6100cd600480360360208110156100bd57600080fd5b5035600160a060020a03166104b5565b005b610172600480360360208110156100e557600080fd5b81019060208101813564010000000081111561010057600080fd5b82018360208201111561011257600080fd5b8035906020019184602083028401116401000000008311171561013457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108bc945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101ae578181015183820152602001610196565b505050509050019250505060405180910390f35b6101ca6108c2565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b8381101561022257818101518382015260200161020a565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b83811015610261578181015183820152602001610249565b50505050905001858103835289818151815260200191508051906020019060200280838360005b838110156102a0578181015183820152602001610288565b50505050905001858103825288818151815260200191508051906020019060200280838360005b838110156102df5781810151838201526020016102c7565b505050509050019a505050505050505050505060405180910390f35b610303610aff565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b83811015610349578181015183820152602001610331565b50505050905001935050505060405180910390f35b610366610b68565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103a0578181015183820152602001610388565b50505050905090810190601f1680156103cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100cd610bff565b6103eb610c27565b60408051918252519081900360200190f35b6103eb610cb6565b61040d610d14565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b8381101561045757818101518382015260200161043f565b50505050905001838103825285818151815260200191508051906020019060200280838360005b8381101561049657818101518382015260200161047e565b505050509050019550505050505060405180910390f35b6103eb610d57565b338015156104c257600080fd5b600160a060020a03811660009081526005602052604081205411806105755750600154604080517fdb512e85000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163db512e85916024808301926020929190829003018186803b15801561054857600080fd5b505afa15801561055c573d6000803e3d6000fd5b505050506040513d602081101561057257600080fd5b50515b151561058057600080fd5b43600481905560025490600090829081151561059857fe5b06905060006002828154811015156105ac57fe5b6000918252602082200154600160a060020a0316915082111561077f57600160a060020a038516156105e1576105e185610db5565b600160a060020a03811615801590610602575033600160a060020a03821614155b1561077f57600154604080517f1dd47d3d000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000939290921691631dd47d3d9160248082019260209290919082900301818787803b15801561067257600080fd5b505af1158015610686573d6000803e3d6000fd5b505050506040513d602081101561069c57600080fd5b505160000b131561072757600154604080517fc8ba1fd8000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163c8ba1fd891602480830192600092919082900301818387803b15801561070e57600080fd5b505af1158015610722573d6000803e3d6000fd5b505050505b600160a060020a0381166000908152600560205260408120819055600280548490811061075057fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107d057600080fd5b505afa1580156107e4573d6000803e3d6000fd5b505050506040513d60208110156107fa57600080fd5b505183148015610887575060018060009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561085757600080fd5b505afa15801561086b573d6000803e3d6000fd5b505050506040513d602081101561088157600080fd5b50510382145b1561089457610894610bff565b61089c610c27565b838115156108a657fe5b0615156108b5576108b5610e98565b5050505050565b50606090565b6060806060806000806002805490506040519080825280602002602001820160405280156108fa578160200160208202803883390190505b506002546040805182815260208084028201019091529195508015610929578160200160208202803883390190505b50925060005b6002548110156109ac576000858281518110151561094957fe5b60209081029091010152600280546005916000918490811061096757fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054845185908390811061099a57fe5b6020908102909101015260010161092f565b50600160009054906101000a9004600160a060020a0316600160a060020a031663b175de176040518163ffffffff1660e060020a02815260040160006040518083038186803b1580156109fe57600080fd5b505afa158015610a12573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610a3b57600080fd5b810190808051640100000000811115610a5357600080fd5b82016020810184811115610a6657600080fd5b8151856020820283011164010000000082111715610a8357600080fd5b505092919050505094506002805480602002602001604051908101604052809291908181526020018280548015610ae357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610ac5575b5050505050955060045490506003805490509150909192939495565b606060006003805480602002602001604051908101604052809291908181526020018280548015610b5957602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b3b575b50505050509150815190509091565b60008054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610bf45780601f10610bc957610100808354040283529160200191610bf4565b820191906000526020600020905b815481529060010190602001808311610bd757829003601f168201915b505050505090505b90565b6000610c09610f05565b9050610c1481610fdc565b610c1c611164565b610c24611288565b50565b600154604080517f4067f0c80000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691634067f0c8916004808301926020929190829003018186803b158015610c8557600080fd5b505afa158015610c99573d6000803e3d6000fd5b505050506040513d6020811015610caf57600080fd5b5051905090565b600154604080517f282f78f20000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163282f78f2916004808301926020929190829003018186803b158015610c8557600080fd5b6060806000610d21610aff565b6040805182815260208084028201019091529194509150818015610d4f578160200160208202803883390190505b509150909192565b600154604080517f6da508300000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691636da50830916004808301926020929190829003018186803b158015610c8557600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a031663282f78f26040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610e0657600080fd5b505afa158015610e1a573d6000803e3d6000fd5b505050506040513d6020811015610e3057600080fd5b505160035410156100a257600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038316179055610c24565b600160009054906101000a9004600160a060020a0316600160a060020a0316637216276b6040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015610eeb57600080fd5b505af1158015610eff573d6000803e3d6000fd5b50505050565b600060026000815481101515610f1757fe5b600091825260209091200154600254600160a060020a039091169150600019015b60008110610fca57600060056000600284815481101515610f5557fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120919091556002805483908110610f8b57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560001901610f38565b506000610fd86002826113e7565b5090565b600154604080517ff0d558170000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163f0d55817916004808301926000929190829003018186803b15801561103a57600080fd5b505afa15801561104e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561107757600080fd5b81019080805164010000000081111561108f57600080fd5b820160208101848111156110a257600080fd5b81518560208202830111640100000000821117156110bf57600080fd5b50909450600093505050505b815181101561115f57600082828151811015156110e457fe5b90602001906020020151905080600160a060020a031684600160a060020a0316141561115657600183510382141561113c5761113783600081518110151561112857fe5b906020019060200201516112f6565b611150565b611150838360010181518110151561112857fe5b5061115f565b506001016110cb565b505050565b60005b6003548110156111ab57600060038281548110151561118257fe5b600091825260209091200154600160a060020a031690506111a2816112f6565b50600101611167565b506002545b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561120157600080fd5b505afa158015611215573d6000803e3d6000fd5b505050506040513d602081101561122b57600080fd5b5051811015610c245760028054600181810183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff19169055016111b0565b600354600019015b600081106112e85760006003828154811015156112a957fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560001901611290565b506000610c246003826113e7565b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561134757600080fd5b505afa15801561135b573d6000803e3d6000fd5b505050506040513d602081101561137157600080fd5b505160025410156100a2576002805460018082019092557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117909155600090815260056020526040902055610c24565b81548183558181111561115f5760008381526020902061115f918101908301610bfc91905b80821115610fd8576000815560010161140c56fea165627a7a72305820e7119be0ac8bf0342523acf16949e2bc06598846371692b9b71a7c210aaa6bfc0029`

// DeployTribeChief100 deploys a new Ethereum contract, binding an instance of TribeChief100 to it.
func DeployTribeChief100(auth *bind.TransactOpts, backend bind.ContractBackend, baseAddress common.Address, pocAddress common.Address, startBlockNumber *big.Int) (common.Address, *types.Transaction, *TribeChief100, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief100ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief100Bin), backend, baseAddress, pocAddress, startBlockNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief100{TribeChief100Caller: TribeChief100Caller{contract: contract}, TribeChief100Transactor: TribeChief100Transactor{contract: contract}, TribeChief100Filterer: TribeChief100Filterer{contract: contract}}, nil
}

// TribeChief100 is an auto generated Go binding around an Ethereum contract.
type TribeChief100 struct {
	TribeChief100Caller     // Read-only binding to the contract
	TribeChief100Transactor // Write-only binding to the contract
	TribeChief100Filterer   // Log filterer for contract events
}

// TribeChief100Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief100Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief100Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief100Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief100Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TribeChief100Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief100Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief100Session struct {
	Contract     *TribeChief100    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TribeChief100CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief100CallerSession struct {
	Contract *TribeChief100Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TribeChief100TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief100TransactorSession struct {
	Contract     *TribeChief100Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TribeChief100Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief100Raw struct {
	Contract *TribeChief100 // Generic contract binding to access the raw methods on
}

// TribeChief100CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief100CallerRaw struct {
	Contract *TribeChief100Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief100TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief100TransactorRaw struct {
	Contract *TribeChief100Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief100 creates a new instance of TribeChief100, bound to a specific deployed contract.
func NewTribeChief100(address common.Address, backend bind.ContractBackend) (*TribeChief100, error) {
	contract, err := bindTribeChief100(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief100{TribeChief100Caller: TribeChief100Caller{contract: contract}, TribeChief100Transactor: TribeChief100Transactor{contract: contract}, TribeChief100Filterer: TribeChief100Filterer{contract: contract}}, nil
}

// NewTribeChief100Caller creates a new read-only instance of TribeChief100, bound to a specific deployed contract.
func NewTribeChief100Caller(address common.Address, caller bind.ContractCaller) (*TribeChief100Caller, error) {
	contract, err := bindTribeChief100(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief100Caller{contract: contract}, nil
}

// NewTribeChief100Transactor creates a new write-only instance of TribeChief100, bound to a specific deployed contract.
func NewTribeChief100Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief100Transactor, error) {
	contract, err := bindTribeChief100(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief100Transactor{contract: contract}, nil
}

// NewTribeChief100Filterer creates a new log filterer instance of TribeChief100, bound to a specific deployed contract.
func NewTribeChief100Filterer(address common.Address, filterer bind.ContractFilterer) (*TribeChief100Filterer, error) {
	contract, err := bindTribeChief100(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TribeChief100Filterer{contract: contract}, nil
}

// bindTribeChief100 binds a generic wrapper to an already deployed contract.
func bindTribeChief100(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief100ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief100 *TribeChief100Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief100.Contract.TribeChief100Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief100 *TribeChief100Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief100.Contract.TribeChief100Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief100 *TribeChief100Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief100.Contract.TribeChief100Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief100 *TribeChief100CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief100.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief100 *TribeChief100TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief100.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief100 *TribeChief100TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief100.Contract.contract.Transact(opts, method, params...)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief100 *TribeChief100Caller) FilterVolunteer(opts *bind.CallOpts, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _TribeChief100.contract.Call(opts, out, "filterVolunteer", volunteers)
	return *ret0, err
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief100 *TribeChief100Session) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief100.Contract.FilterVolunteer(&_TribeChief100.CallOpts, volunteers)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief100 *TribeChief100CallerSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief100.Contract.FilterVolunteer(&_TribeChief100.CallOpts, volunteers)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief100 *TribeChief100Caller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief100.contract.Call(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief100 *TribeChief100Session) GetEpoch() (*big.Int, error) {
	return _TribeChief100.Contract.GetEpoch(&_TribeChief100.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief100.Contract.GetEpoch(&_TribeChief100.CallOpts)
}

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief100 *TribeChief100Caller) GetNextRoundSignerList(opts *bind.CallOpts) (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	ret := new(struct {
		NextRoundSignerList []common.Address
		Length              *big.Int
	})
	out := ret
	err := _TribeChief100.contract.Call(opts, out, "getNextRoundSignerList")
	return *ret, err
}

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief100 *TribeChief100Session) GetNextRoundSignerList() (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	return _TribeChief100.Contract.GetNextRoundSignerList(&_TribeChief100.CallOpts)
}

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetNextRoundSignerList() (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	return _TribeChief100.Contract.GetNextRoundSignerList(&_TribeChief100.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100Caller) GetSignerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief100.contract.Call(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief100.Contract.GetSignerLimit(&_TribeChief100.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief100.Contract.GetSignerLimit(&_TribeChief100.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief100 *TribeChief100Caller) GetStatus(opts *bind.CallOpts) (struct {
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
	err := _TribeChief100.contract.Call(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief100 *TribeChief100Session) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief100.Contract.GetStatus(&_TribeChief100.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief100.Contract.GetStatus(&_TribeChief100.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100Caller) GetVolunteerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief100.contract.Call(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief100.Contract.GetVolunteerLimit(&_TribeChief100.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief100.Contract.GetVolunteerLimit(&_TribeChief100.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief100 *TribeChief100Caller) GetVolunteers(opts *bind.CallOpts) (struct {
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
	err := _TribeChief100.contract.Call(opts, out, "getVolunteers")
	return *ret, err
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief100 *TribeChief100Session) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief100.Contract.GetVolunteers(&_TribeChief100.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief100 *TribeChief100CallerSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief100.Contract.GetVolunteers(&_TribeChief100.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief100 *TribeChief100Caller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief100.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief100 *TribeChief100Session) Version() (string, error) {
	return _TribeChief100.Contract.Version(&_TribeChief100.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief100 *TribeChief100CallerSession) Version() (string, error) {
	return _TribeChief100.Contract.Version(&_TribeChief100.CallOpts)
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief100 *TribeChief100Transactor) GenSigners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief100.contract.Transact(opts, "genSigners")
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief100 *TribeChief100Session) GenSigners() (*types.Transaction, error) {
	return _TribeChief100.Contract.GenSigners(&_TribeChief100.TransactOpts)
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief100 *TribeChief100TransactorSession) GenSigners() (*types.Transaction, error) {
	return _TribeChief100.Contract.GenSigners(&_TribeChief100.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief100 *TribeChief100Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief100.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief100 *TribeChief100Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief100.Contract.Update(&_TribeChief100.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief100 *TribeChief100TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief100.Contract.Update(&_TribeChief100.TransactOpts, volunteer)
}
