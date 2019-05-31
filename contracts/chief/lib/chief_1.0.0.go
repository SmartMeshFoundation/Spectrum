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

// ChiefBaseABI is the input ABI used to generate the binding from.
const ChiefBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getLeaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_leaderLimit\",\"type\":\"uint256\"}],\"name\":\"setLeaderLimit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"volunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_tribe\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tribeAddr\",\"type\":\"address\"}],\"name\":\"initTribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pocAddr\",\"type\":\"address\"}],\"name\":\"initPoc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLeaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vsn\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"pocChangeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"appendLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pocAddr\",\"type\":\"address\"},{\"name\":\"tribeAddr\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ChiefBaseBin is the compiled bytecode used for deploying new contracts.
const ChiefBaseBin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a090815261003e91600091906100b1565b5060058055600360065561181b600755600260085534801561005f57600080fd5b50604051604080610d7d8339810180604052604081101561007f57600080fd5b50805160209091015160028054600160a060020a0319163317905560079190915560068190556000190160085561014c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f257805160ff191683800117855561011f565b8280016001018555821561011f579182015b8281111561011f578251825591602001919060010190610104565b5061012b92915061012f565b5090565b61014991905b8082111561012b5760008155600101610135565b90565b610c228061015b6000396000f3fe608060405260043610610168576000357c0100000000000000000000000000000000000000000000000000000000900480639037f182116100d3578063bf8fb6141161008c578063bf8fb61414610486578063bfe3a272146104bf578063db512e85146104f2578063eb5c001114610539578063ed77db3f1461054e578063f09a40161461057857610168565b80639037f18214610357578063961c5c7a1461038a578063aa5647751461039f578063ac69cbea146103d2578063b2bdfa7b146103e7578063bbf2ef67146103fc57610168565b8063609bf73b11610125578063609bf73b1461029f5780636c151173146102d0578063757991a8146102e5578063893d20e8146102fa5780638b7204bb1461030f578063900cf0cf1461034257610168565b80630d6bb50f1461016d5780630e7279e4146101d257806313af4035146101f15780631dd47d3d1461021757806326b249c81461026357806357f607a11461028a575b600080fd5b34801561017957600080fd5b506101826105b3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101be5781810151838201526020016101a6565b505050509050019250505060405180910390f35b6101ef600480360360208110156101e857600080fd5b5035610616565b005b6101ef6004803603602081101561020757600080fd5b5035600160a060020a0316610632565b34801561022357600080fd5b5061024a6004803603602081101561023a57600080fd5b5035600160a060020a0316610678565b60408051600092830b90920b8252519081900360200190f35b34801561026f57600080fd5b50610278610717565b60408051918252519081900360200190f35b34801561029657600080fd5b5061027861071d565b3480156102ab57600080fd5b506102b4610723565b60408051600160a060020a039092168252519081900360200190f35b3480156102dc57600080fd5b50610278610732565b3480156102f157600080fd5b50610278610738565b34801561030657600080fd5b506102b461073e565b34801561031b57600080fd5b506101ef6004803603602081101561033257600080fd5b5035600160a060020a031661074d565b34801561034e57600080fd5b5061027861077c565b34801561036357600080fd5b506101ef6004803603602081101561037a57600080fd5b5035600160a060020a0316610782565b34801561039657600080fd5b50610278610895565b3480156103ab57600080fd5b506101ef600480360360208110156103c257600080fd5b5035600160a060020a031661089b565b3480156103de57600080fd5b506102786108ca565b3480156103f357600080fd5b506102b46108d0565b34801561040857600080fd5b506104116108df565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561044b578181015183820152602001610433565b50505050905090810190601f1680156104785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561049257600080fd5b506101ef600480360360408110156104a957600080fd5b50600160a060020a03813516906020013561096d565b3480156104cb57600080fd5b506101ef600480360360208110156104e257600080fd5b5035600160a060020a03166109f7565b3480156104fe57600080fd5b506105256004803603602081101561051557600080fd5b5035600160a060020a0316610ac5565b604080519115158252519081900360200190f35b34801561054557600080fd5b50610278610b2f565b34801561055a57600080fd5b506102b46004803603602081101561057157600080fd5b5035610b35565b34801561058457600080fd5b506101ef6004803603604081101561059b57600080fd5b50600160a060020a0381358116916020013516610b5d565b6060600480548060200260200160405190810160405280929190818152602001828054801561060b57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105ed575b505050505090505b90565b600254600160a060020a0316331461062d57600080fd5b600555565b600254600160a060020a0316331461064957600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154604080517f2988fcfd000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000939290921691632988fcfd9160248082019260209290919082900301818787803b1580156106e357600080fd5b505af11580156106f7573d6000803e3d6000fd5b505050506040513d602081101561070d57600080fd5b505190505b919050565b60085481565b60065481565b600354600160a060020a031681565b60055481565b60075490565b600254600160a060020a031690565b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60075481565b600254600160a060020a0316331461079957600080fd5b600454600114156107a957600080fd5b60005b6004548110156108915781600160a060020a03166004828154811015156107cf57fe5b600091825260209091200154600160a060020a0316141561088957805b6004546000190181101561086f57600480546001830190811061080b57fe5b60009182526020909120015460048054600160a060020a03909216918390811061083157fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016107ec565b50600480546000199283019201906108879082610baf565b505b6001016107ac565b5050565b60085490565b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055490565b600254600160a060020a031681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156109655780601f1061093a57610100808354040283529160200191610965565b820191906000526020600020905b81548152906001019060200180831161094857829003601f168201915b505050505081565b600154604080517f56df3db1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201859052915191909216916356df3db191604480830192600092919082900301818387803b1580156109db57600080fd5b505af11580156109ef573d6000803e3d6000fd5b505050505050565b600254600160a060020a03163314610a0e57600080fd5b6005546004541015610ac25760005b600454811015610a685781600160a060020a0316600482815481101515610a4057fe5b600091825260209091200154600160a060020a03161415610a6057600080fd5b600101610a1d565b50600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b6000805b600454811015610b2657600160a060020a03831615801590610b0f57506004805482908110610af457fe5b600091825260209091200154600160a060020a038481169116145b15610b1e576001915050610712565b600101610ac9565b50600092915050565b60065490565b6004805482908110610b4357fe5b600091825260209091200154600160a060020a0316905081565b600154600160a060020a0316158015610b7e5750600160a060020a03821615155b8015610b925750600160a060020a03811615155b1515610b9d57600080fd5b610ba68261089b565b6108918161074d565b815481835581811115610bd357600083815260209020610bd3918101908301610bd8565b505050565b61061391905b80821115610bf25760008155600101610bde565b509056fea165627a7a7230582045a355af8c85a8af743025e20fa1a87b4779d8d7ee3560cc3d53488e51f72fee0029`

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
	return address, tx, &ChiefBase{ChiefBaseCaller: ChiefBaseCaller{contract: contract}, ChiefBaseTransactor: ChiefBaseTransactor{contract: contract}}, nil
}

// ChiefBase is an auto generated Go binding around an Ethereum contract.
type ChiefBase struct {
	ChiefBaseCaller     // Read-only binding to the contract
	ChiefBaseTransactor // Write-only binding to the contract
}

// ChiefBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChiefBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiefBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiefBaseSession struct {
	Contract     *ChiefBase              // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChiefBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiefBaseCallerSession struct {
	Contract *ChiefBaseCaller        // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
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
	contract, err := bindChiefBase(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChiefBase{ChiefBaseCaller: ChiefBaseCaller{contract: contract}, ChiefBaseTransactor: ChiefBaseTransactor{contract: contract}}, nil
}

// NewChiefBaseCaller creates a new read-only instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBaseCaller(address common.Address, caller bind.ContractCaller) (*ChiefBaseCaller, error) {
	contract, err := bindChiefBase(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefBaseCaller{contract: contract}, nil
}

// NewChiefBaseTransactor creates a new write-only instance of ChiefBase, bound to a specific deployed contract.
func NewChiefBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ChiefBaseTransactor, error) {
	contract, err := bindChiefBase(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ChiefBaseTransactor{contract: contract}, nil
}

// bindChiefBase binds a generic wrapper to an already deployed contract.
func bindChiefBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiefBase *ChiefBaseRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase.Contract.ChiefBaseCaller.contract.CallWithNumber(opts, result, method, params...)
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
func (_ChiefBase *ChiefBaseCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase.Contract.contract.CallWithNumber(opts, result, method, params...)
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
func (_ChiefBase *ChiefBaseCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "_owner")
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
func (_ChiefBase *ChiefBaseCaller) Tribe(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "_tribe")
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
func (_ChiefBase *ChiefBaseCaller) Epoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "epoch")
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

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) GetEpoch() (*big.Int, error) {
	return _ChiefBase.Contract.GetEpoch(&_ChiefBase.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) GetEpoch() (*big.Int, error) {
	return _ChiefBase.Contract.GetEpoch(&_ChiefBase.CallOpts)
}

// GetLeaderLimit is a free data retrieval call binding the contract method 0xac69cbea.
//
// Solidity: function getLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) GetLeaderLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getLeaderLimit")
	return *ret0, err
}

// GetLeaderLimit is a free data retrieval call binding the contract method 0xac69cbea.
//
// Solidity: function getLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) GetLeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetLeaderLimit(&_ChiefBase.CallOpts)
}

// GetLeaderLimit is a free data retrieval call binding the contract method 0xac69cbea.
//
// Solidity: function getLeaderLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) GetLeaderLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetLeaderLimit(&_ChiefBase.CallOpts)
}

// GetLeaderList is a free data retrieval call binding the contract method 0x0d6bb50f.
//
// Solidity: function getLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseCaller) GetLeaderList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getLeaderList")
	return *ret0, err
}

// GetLeaderList is a free data retrieval call binding the contract method 0x0d6bb50f.
//
// Solidity: function getLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseSession) GetLeaderList() ([]common.Address, error) {
	return _ChiefBase.Contract.GetLeaderList(&_ChiefBase.CallOpts)
}

// GetLeaderList is a free data retrieval call binding the contract method 0x0d6bb50f.
//
// Solidity: function getLeaderList() constant returns(address[])
func (_ChiefBase *ChiefBaseCallerSession) GetLeaderList() ([]common.Address, error) {
	return _ChiefBase.Contract.GetLeaderList(&_ChiefBase.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ChiefBase *ChiefBaseCaller) GetOwner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ChiefBase *ChiefBaseSession) GetOwner() (common.Address, error) {
	return _ChiefBase.Contract.GetOwner(&_ChiefBase.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ChiefBase *ChiefBaseCallerSession) GetOwner() (common.Address, error) {
	return _ChiefBase.Contract.GetOwner(&_ChiefBase.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) GetSignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetSignerLimit(&_ChiefBase.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) GetSignerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetSignerLimit(&_ChiefBase.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseSession) GetVolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetVolunteerLimit(&_ChiefBase.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _ChiefBase.Contract.GetVolunteerLimit(&_ChiefBase.CallOpts)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase *ChiefBaseCaller) IsLeader(opts *bind.CallOptsWithNumber, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "isLeader", addr)
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
func (_ChiefBase *ChiefBaseCaller) LeaderLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "leaderLimit")
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
func (_ChiefBase *ChiefBaseCaller) LeaderList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "leaderList", arg0)
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

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) SignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "signerLimit")
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

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase *ChiefBaseCaller) VolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "volunteerLimit")
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
func (_ChiefBase *ChiefBaseCaller) Vsn(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChiefBase.contract.CallWithNumber(opts, out, "vsn")
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

// InitPoc is a paid mutator transaction binding the contract method 0xaa564775.
//
// Solidity: function initPoc(pocAddr address) returns()
func (_ChiefBase *ChiefBaseTransactor) InitPoc(opts *bind.TransactOpts, pocAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "initPoc", pocAddr)
}

// InitPoc is a paid mutator transaction binding the contract method 0xaa564775.
//
// Solidity: function initPoc(pocAddr address) returns()
func (_ChiefBase *ChiefBaseSession) InitPoc(pocAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.InitPoc(&_ChiefBase.TransactOpts, pocAddr)
}

// InitPoc is a paid mutator transaction binding the contract method 0xaa564775.
//
// Solidity: function initPoc(pocAddr address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) InitPoc(pocAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.InitPoc(&_ChiefBase.TransactOpts, pocAddr)
}

// InitTribe is a paid mutator transaction binding the contract method 0x8b7204bb.
//
// Solidity: function initTribe(tribeAddr address) returns()
func (_ChiefBase *ChiefBaseTransactor) InitTribe(opts *bind.TransactOpts, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.contract.Transact(opts, "initTribe", tribeAddr)
}

// InitTribe is a paid mutator transaction binding the contract method 0x8b7204bb.
//
// Solidity: function initTribe(tribeAddr address) returns()
func (_ChiefBase *ChiefBaseSession) InitTribe(tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.InitTribe(&_ChiefBase.TransactOpts, tribeAddr)
}

// InitTribe is a paid mutator transaction binding the contract method 0x8b7204bb.
//
// Solidity: function initTribe(tribeAddr address) returns()
func (_ChiefBase *ChiefBaseTransactorSession) InitTribe(tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase.Contract.InitTribe(&_ChiefBase.TransactOpts, tribeAddr)
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
const POCABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"depositTest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"minerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawWaitNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwnerEffectiveNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNormalList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"},{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStopList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// POCBin is the compiled bytecode used for deploying new contracts.
const POCBin = `0x608060405234801561001057600080fd5b506040516060806114bf8339810180604052606081101561003057600080fd5b5080516020820151604090920151600492909255600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a03909216600160a060020a0319909216919091179055600555611418806100a76000396000f3fe608060405260043610610105576000357c01000000000000000000000000000000000000000000000000000000009004806351cff8d9116100a7578063656cb0bc11610076578063656cb0bc146104475780639c52ade21461047a578063dd0b281e1461048f578063f2a558f9146104c257610105565b806351cff8d91461028e57806353ed5143146102c157806356df3db1146103f9578063581003701461043257610105565b80632988fcfd116100e35780632988fcfd146101b357806341b3d185146101ff5780634b58f36e1461021457806351393ce31461022957610105565b80630b631bbc1461010a5780630c58ed8e14610132578063147049701461018c575b600080fd5b6101306004803603602081101561012057600080fd5b5035600160a060020a03166104ee565b005b34801561013e57600080fd5b506101656004803603602081101561015557600080fd5b5035600160a060020a03166105db565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b34801561019857600080fd5b506101a1610608565b60408051918252519081900360200190f35b3480156101bf57600080fd5b506101e6600480360360208110156101d657600080fd5b5035600160a060020a031661060e565b60408051600092830b90920b8252519081900360200190f35b34801561020b57600080fd5b506101a16106a2565b34801561022057600080fd5b506101a16106a8565b34801561023557600080fd5b5061023e6106ae565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561027a578181015183820152602001610262565b505050509050019250505060405180910390f35b34801561029a57600080fd5b50610130600480360360208110156102b157600080fd5b5035600160a060020a0316610711565b3480156102cd57600080fd5b506102d66108c5565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b8381101561032257818101518382015260200161030a565b50505050905001858103845288818151815260200191508051906020019060200280838360005b83811015610361578181015183820152602001610349565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156103a0578181015183820152602001610388565b50505050905001858103825286818151815260200191508051906020019060200280838360005b838110156103df5781810151838201526020016103c7565b505050509050019850505050505050505060405180910390f35b34801561040557600080fd5b506101306004803603604081101561041c57600080fd5b50600160a060020a038135169060200135610c7f565b34801561043e57600080fd5b5061023e610db4565b34801561045357600080fd5b506101306004803603602081101561046a57600080fd5b5035600160a060020a0316610e14565b34801561048657600080fd5b5061023e610e57565b34801561049b57600080fd5b50610130600480360360208110156104b257600080fd5b5035600160a060020a0316610eb7565b610130600480360360608110156104d857600080fd5b508035906020810135906040013560ff16610ee6565b6004543410156104fd57600080fd5b600160a060020a038116151561051257600080fd5b600160a060020a0381166000908152602081905260409020600101541561053857600080fd5b604080516060810182523381523460208083019182526000838501818152600160a060020a039687168083529282905294812093518454961673ffffffffffffffffffffffffffffffffffffffff1996871617845591516001808501919091559351600290930192909255825480840184559290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69091018054909216179055565b600160a060020a039081166000908152602081905260409020600181015460028201549154909391921690565b60055481565b600380546000918290600019830183811061062557fe5b600091825260209091200154600160a060020a0316905060018211801561064d575060065443105b1561067c5760038054600119840190811061066457fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461069157600080fd5b61069a84611073565b949350505050565b60045481565b60065481565b6060600180548060200260200160405190810160405280929190818152602001828054801561070657602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106e8575b505050505090505b90565b600160a060020a0381811660009081526020819052604090205416331461073757600080fd5b600160a060020a0381166000908152602081905260408120600201541161075d57600080fd5b600554600160a060020a03821660009081526020819052604090206002015443031161078857600080fd5b600160a060020a0381166000908152602081905260408120600181018054825473ffffffffffffffffffffffffffffffffffffffff19168355908390556002918201839055905490915b6001820381101561087d5760028054829081106107eb57fe5b600091825260209091200154600160a060020a03858116911614156108755760028054600019840190811061081c57fe5b60009182526020909120015460028054600160a060020a03909216918390811061084257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061087d565b6001016107d2565b5060028054600019019061089190826113a5565b50604051339083156108fc029084906000818181858888f193505050501580156108bf573d6000803e3d6000fd5b50505050565b6002546001546040805182840180825260208082028301019092526060938493849384938490828015610902578160200160208202803883390190505b509050606082604051908082528060200260200182016040528015610931578160200160208202803883390190505b509050606083604051908082528060200260200182016040528015610960578160200160208202803883390190505b50905060608460405190808252806020026020018201604052801561098f578160200160208202803883390190505b50905060005b87811015610afb5760028054829081106109ab57fe5b6000918252602090912001548551600160a060020a03909116908690839081106109d157fe5b600160a060020a03909216602092830290910190910152600280546000918291849081106109fb57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600101548451859083908110610a3157fe5b6020908102909101015260028054600091829184908110610a4e57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600201548351849083908110610a8457fe5b6020908102909101015260028054600091829184908110610aa157fe5b6000918252602080832090910154600160a060020a0390811684529083019390935260409091019020548351911690839083908110610adc57fe5b600160a060020a03909216602092830290910190910152600101610995565b5060005b86811015610c6d576001805482908110610b1557fe5b6000918252602090912001548551600160a060020a039091169086908a8401908110610b3d57fe5b600160a060020a0390921660209283029091019091015260018054600091829184908110610b6757fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154845185908a8401908110610b9f57fe5b6020908102909101015260018054600091829184908110610bbc57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060020154835184908a8401908110610bf457fe5b6020908102909101015260018054600091829184908110610c1157fe5b6000918252602080832090910154600160a060020a039081168452908301939093526040909101902054835191169083908a8401908110610c4e57fe5b600160a060020a03909216602092830290910190910152600101610aff565b50929a91995097509095509350505050565b60038054906000906000198301838110610c9557fe5b600091825260209091200154600160a060020a03169050600182118015610cbd575060065443105b15610cec57600380546001198401908110610cd457fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610d0157600080fd5b600654431015610d5357600380548591906000198101908110610d2057fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610dac565b600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386161790555b505060065550565b6060600380548060200260200160405190810160405280929190818152602001828054801561070657602002820191906000526020600020908154600160a060020a031681526001909101906020018083116106e8575050505050905090565b600160a060020a03818116600090815260208190526040902054163314610e3a57600080fd5b610e4381611073565b60000b6001141515610e5457600080fd5b50565b6060600280548060200260200160405190810160405280929190818152602001828054801561070657602002820191906000526020600020908154600160a060020a031681526001909101906020018083116106e8575050505050905090565b600160a060020a03818116600090815260208190526040902054163314610edd57600080fd5b610e438161120d565b600454341015610ef557600080fd5b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff861660748501526094840188905260b484018790529351909260019260d4808301939192601f198301929081900390910190855afa158015610f84573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a0381161515610fa657600080fd5b600160a060020a03811660009081526020819052604090206001015415610fcc57600080fd5b604080516060810182523381523460208083019182526000838501818152600160a060020a039687168083529282905294812093518454961673ffffffffffffffffffffffffffffffffffffffff1996871617845591516001808501919091559351600290930192909255825480840184559290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6909101805490921617905550505050565b600160a060020a038116600090815260208190526040812060010154151561109e5750600019611208565b600160a060020a038216600090815260208190526040902060020154156110c757506000611208565b600160a060020a038216600090815260208190526040812043600290910155600154905b6001820381101561119657600180548290811061110457fe5b600091825260209091200154600160a060020a038581169116141561118e5760018054600019840190811061113557fe5b60009182526020909120015460018054600160a060020a03909216918390811061115b57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611196565b6001016110eb565b506001805460001901906111aa90826113a5565b505060028054600180820183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905590505b919050565b600160a060020a03811660009081526020819052604081206001015415156112385750600019611208565b600160a060020a038216600090815260208190526040902060020154151561126257506000611208565b600160a060020a0382166000908152602081905260408120600290810182905554905b6001820381101561133057600280548290811061129e57fe5b600091825260209091200154600160a060020a0385811691161415611328576002805460001984019081106112cf57fe5b60009182526020909120015460028054600160a060020a0390921691839081106112f557fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611330565b600101611285565b5060028054600019019061134490826113a5565b505060018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff199091161790559050919050565b8154818355818111156113c9576000838152602090206113c99181019083016113ce565b505050565b61070e91905b808211156113e857600081556001016113d4565b509056fea165627a7a72305820fb5d708faa2575ae9ecde0b2f6d0e8a2ebaedc3ea044741128353e81fb0ce3bb0029`

// DeployPOC deploys a new Ethereum contract, binding an instance of POC to it.
func DeployPOC(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address, _amount *big.Int, _number *big.Int) (common.Address, *types.Transaction, *POC, error) {
	parsed, err := abi.JSON(strings.NewReader(POCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(POCBin), backend, _addr, _amount, _number)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &POC{POCCaller: POCCaller{contract: contract}, POCTransactor: POCTransactor{contract: contract}}, nil
}

// POC is an auto generated Go binding around an Ethereum contract.
type POC struct {
	POCCaller     // Read-only binding to the contract
	POCTransactor // Write-only binding to the contract
}

// POCCaller is an auto generated read-only Go binding around an Ethereum contract.
type POCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type POCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type POCSession struct {
	Contract     *POC                    // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// POCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type POCCallerSession struct {
	Contract *POCCaller              // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
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
	contract, err := bindPOC(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &POC{POCCaller: POCCaller{contract: contract}, POCTransactor: POCTransactor{contract: contract}}, nil
}

// NewPOCCaller creates a new read-only instance of POC, bound to a specific deployed contract.
func NewPOCCaller(address common.Address, caller bind.ContractCaller) (*POCCaller, error) {
	contract, err := bindPOC(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &POCCaller{contract: contract}, nil
}

// NewPOCTransactor creates a new write-only instance of POC, bound to a specific deployed contract.
func NewPOCTransactor(address common.Address, transactor bind.ContractTransactor) (*POCTransactor, error) {
	contract, err := bindPOC(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &POCTransactor{contract: contract}, nil
}

// bindPOC binds a generic wrapper to an already deployed contract.
func bindPOC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(POCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POC *POCRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _POC.Contract.POCCaller.contract.CallWithNumber(opts, result, method, params...)
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
func (_POC *POCCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _POC.Contract.contract.CallWithNumber(opts, result, method, params...)
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

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[])
func (_POC *POCCaller) GetAll(opts *bind.CallOptsWithNumber) ([]common.Address, []*big.Int, []*big.Int, []common.Address, error) {
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
	err := _POC.contract.CallWithNumber(opts, out, "getAll")
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

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC *POCCaller) GetNormalList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "getNormalList")
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
func (_POC *POCCaller) GetOwnerList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "getOwnerList")
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
func (_POC *POCCaller) GetStopList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "getStopList")
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

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_POC *POCCaller) MinDeposit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "minDeposit")
	return *ret0, err
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_POC *POCSession) MinDeposit() (*big.Int, error) {
	return _POC.Contract.MinDeposit(&_POC.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_POC *POCCallerSession) MinDeposit() (*big.Int, error) {
	return _POC.Contract.MinDeposit(&_POC.CallOpts)
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC *POCCaller) MinerStatus(opts *bind.CallOptsWithNumber, _addr common.Address) (*big.Int, *big.Int, common.Address, error) {
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
	err := _POC.contract.CallWithNumber(opts, out, "minerStatus", _addr)
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
func (_POC *POCCaller) NewOwnerEffectiveNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "newOwnerEffectiveNumber")
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
func (_POC *POCCaller) WithdrawWaitNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC.contract.CallWithNumber(opts, out, "withdrawWaitNumber")
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

// DepositTest is a paid mutator transaction binding the contract method 0x0b631bbc.
//
// Solidity: function depositTest(minerAddress address) returns()
func (_POC *POCTransactor) DepositTest(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC.contract.Transact(opts, "depositTest", minerAddress)
}

// DepositTest is a paid mutator transaction binding the contract method 0x0b631bbc.
//
// Solidity: function depositTest(minerAddress address) returns()
func (_POC *POCSession) DepositTest(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.DepositTest(&_POC.TransactOpts, minerAddress)
}

// DepositTest is a paid mutator transaction binding the contract method 0x0b631bbc.
//
// Solidity: function depositTest(minerAddress address) returns()
func (_POC *POCTransactorSession) DepositTest(minerAddress common.Address) (*types.Transaction, error) {
	return _POC.Contract.DepositTest(&_POC.TransactOpts, minerAddress)
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

// TribeChief_1_0_0ABI is the input ABI used to generate the binding from.
const TribeChief_1_0_0ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"blMap\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_blackMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_v2s\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_clean_all_signer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLeader\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"signersMap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_blackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_volunteerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"volunteerMap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeBlackMember\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_set_leader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genSigners_clean_volunteer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_signerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"pushBlackMember\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"pushVolunteer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"baseAddress\",\"type\":\"address\"},{\"name\":\"pocAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_1_0_0Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_1_0_0Bin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a0908152620000409160009190620003af565b503480156200004e57600080fd5b50604051604080620022d7833981018060405260408110156200007057600080fd5b50805160209091015160018054600160a060020a031916600160a060020a038085169190911791829055604080517ff09a401600000000000000000000000000000000000000000000000000000000815282851660048201523060248201529051929091169163f09a40169160448082019260009290919082900301818387803b158015620000fe57600080fd5b505af115801562000113573d6000803e3d6000fd5b505050506060600160009054906101000a9004600160a060020a0316600160a060020a0316630d6bb50f6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160006040518083038186803b1580156200018457600080fd5b505afa15801562000199573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015620001c357600080fd5b810190808051640100000000811115620001dc57600080fd5b82016020810184811115620001f057600080fd5b81518560208202830111640100000000821117156200020e57600080fd5b5050805190945060001092506200022791505057600080fd5b6001600760008360008151811015156200023d57fe5b90602001906020020151600160a060020a0316600160a060020a031681526020019081526020016000208190555060048160008151811015156200027d57fe5b6020908102919091018101518254600181018455600093845291909220018054600160a060020a031916600160a060020a039092169190911790556004545b600160009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200032757600080fd5b505afa1580156200033c573d6000803e3d6000fd5b505050506040513d60208110156200035357600080fd5b5051811015620003a55760048054600181810183556000929092527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054600160a060020a031916905501620002bc565b5050505062000454565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003f257805160ff191683800117855562000422565b8280016001018555821562000422579182015b828111156200042257825182559160200191906001019062000405565b506200043092915062000434565b5090565b6200045191905b808211156200043057600081556001016200043b565b90565b611e7380620004646000396000f3fe60806040526004361061016a5760003560e060020a9004806365ddc059116100d5578063961c5c7a1161008e578063c8e0666a11610068578063c8e0666a146106ee578063d4d8c65714610714578063d7ca4a1c14610747578063eb5c0011146107fc5761016a565b8063961c5c7a1461069a578063aca680a2146106af578063b4deb2ca146106c45761016a565b806365ddc059146105d857806373d53ec714610602578063757991a8146106175780637cc81e5a1461062c5780638698514a1461065f5780638c607d25146106855761016a565b80634c051f14116101275780634c051f141461036f5780634e69d5601461038457806354fd4d50146104ca57806357e871e7146105545780635c7f94b71461057b578063609ed258146105ae5761016a565b8063113aa4711461016f5780631c1b8772146101ca57806320c1a518146101ff5780633aa07a36146102ff57806340267e93146103455780634b77da761461035a575b600080fd5b34801561017b57600080fd5b506101a26004803603602081101561019257600080fd5b5035600160a060020a0316610811565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b3480156101d657600080fd5b506101fd600480360360208110156101ed57600080fd5b5035600160a060020a031661083c565b005b34801561020b57600080fd5b506102af6004803603602081101561022257600080fd5b81019060208101813564010000000081111561023d57600080fd5b82018360208201111561024f57600080fd5b8035906020019184602083028401116401000000008311171561027157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610b9f945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102eb5781810151838201526020016102d3565b505050509050019250505060405180910390f35b34801561030b57600080fd5b506103296004803603602081101561032257600080fd5b5035610ba5565b60408051600160a060020a039092168252519081900360200190f35b34801561035157600080fd5b506101fd610bcd565b34801561036657600080fd5b506101fd610d6c565b34801561037b57600080fd5b506102af610e5b565b34801561039057600080fd5b50610399610f4a565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b838110156103f15781810151838201526020016103d9565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b83811015610430578181015183820152602001610418565b50505050905001858103835289818151815260200191508051906020019060200280838360005b8381101561046f578181015183820152602001610457565b50505050905001858103825288818151815260200191508051906020019060200280838360005b838110156104ae578181015183820152602001610496565b505050509050019a505050505050505050505060405180910390f35b3480156104d657600080fd5b506104df611104565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610519578181015183820152602001610501565b50505050905090810190601f1680156105465780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561056057600080fd5b5061056961119a565b60408051918252519081900360200190f35b34801561058757600080fd5b506105696004803603602081101561059e57600080fd5b5035600160a060020a03166111a0565b3480156105ba57600080fd5b50610329600480360360208110156105d157600080fd5b50356111b2565b3480156105e457600080fd5b50610329600480360360208110156105fb57600080fd5b50356111c0565b34801561060e57600080fd5b506101fd6111ce565b34801561062357600080fd5b506105696111f0565b34801561063857600080fd5b506105696004803603602081101561064f57600080fd5b5035600160a060020a031661127f565b6101fd6004803603602081101561067557600080fd5b5035600160a060020a0316611291565b34801561069157600080fd5b506101fd6114d7565b3480156106a657600080fd5b5061056961167b565b3480156106bb57600080fd5b506101fd6116d9565b3480156106d057600080fd5b50610329600480360360208110156106e757600080fd5b50356116fe565b6101fd6004803603602081101561070457600080fd5b5035600160a060020a031661170c565b34801561072057600080fd5b506101fd6004803603602081101561073757600080fd5b5035600160a060020a031661187d565b34801561075357600080fd5b5061075c611960565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b838110156107a657818101518382015260200161078e565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156107e55781810151838201526020016107cd565b505050509050019550505050505060405180910390f35b34801561080857600080fd5b506105696119f8565b600960205260009081526040902080546001820154600290920154600160a060020a03909116919083565b3380151561084957600080fd5b600160a060020a03811660009081526007602052604081205411806108fc5750600154604080517fdb512e85000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163db512e85916024808301926020929190829003018186803b1580156108cf57600080fd5b505afa1580156108e3573d6000803e3d6000fd5b505050506040513d60208110156108f957600080fd5b50515b151561090757600080fd5b4360065561091433611291565b600454600654600090829081151561092857fe5b069050600060048281548110151561093c57fe5b6000918252602082200154600160a060020a03169150821115610a8357600160a060020a03851615610971576109718561187d565b600160a060020a03811615801590610992575033600160a060020a03821614155b15610a8357600154604080517f1dd47d3d000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691631dd47d3d9160248083019260209291908290030181600087803b1580156109ff57600080fd5b505af1158015610a13573d6000803e3d6000fd5b505050506040513d6020811015610a2957600080fd5b5050600160a060020a03811660009081526007602052604081208190556004805484908110610a5457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b600160009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610ad457600080fd5b505afa158015610ae8573d6000803e3d6000fd5b505050506040513d6020811015610afe57600080fd5b505183148015610b8b575060018060009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610b5b57600080fd5b505afa158015610b6f573d6000803e3d6000fd5b505050506040513d6020811015610b8557600080fd5b50510382145b15610b9857610b986111ce565b5050505050565b50606090565b6003805482908110610bb357fe5b600091825260209091200154600160a060020a0316905081565b60005b600554811015610c99576000600582815481101515610beb57fe5b60009182526020918290200154600154604080517feb5c00110000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169550610c909486949093169263eb5c00119260048082019391829003018186803b158015610c5c57600080fd5b505afa158015610c70573d6000803e3d6000fd5b505050506040513d6020811015610c8657600080fd5b5051600101611a56565b50600101610bd0565b506004545b600160009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610cef57600080fd5b505afa158015610d03573d6000803e3d6000fd5b505050506040513d6020811015610d1957600080fd5b5051811015610d695760048054600181810183556000929092527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054600160a060020a031916905501610c9e565b50565b6004546040805182815260208084028201019091526060918015610d9a578160200160208202803883390190505b50905060005b8151811015610dfc576004805482908110610db757fe5b6000918252602090912001548251600160a060020a0390911690839083908110610ddd57fe5b600160a060020a03909216602092830290910190910152600101610da0565b5080515b6001811115610e57576000198101610e1781611c46565b60008382815181101515610e2757fe5b906020019060200201519050600082118015610e4b5750600160a060020a03811615155b50505060001901610e00565b5050565b600154604080517f0d6bb50f0000000000000000000000000000000000000000000000000000000081529051606092600160a060020a031691630d6bb50f916004808301926000929190829003018186803b158015610eb957600080fd5b505afa158015610ecd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ef657600080fd5b810190808051640100000000811115610f0e57600080fd5b82016020810184811115610f2157600080fd5b8151856020820283011164010000000082111715610f3e57600080fd5b50909450505050505b90565b606080606080600080600480549050604051908082528060200260200182016040528015610f82578160200160208202803883390190505b506004546040805182815260208084028201019091529195508015610fb1578160200160208202803883390190505b50925060005b6004548110156110345760008582815181101515610fd157fe5b602090810290910101526004805460079160009184908110610fef57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054845185908390811061102257fe5b60209081029091010152600101610fb7565b50600280548060200260200160405190810160405280929190818152602001828054801561108b57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161106d575b5050505050945060048054806020026020016040519081016040528092919081815260200182805480156110e857602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116110ca575b5050505050955060065490506005805490509150909192939495565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156111905780601f1061116557610100808354040283529160200191611190565b820191906000526020600020905b81548152906001019060200180831161117357829003601f168201915b5050505050905090565b60065481565b60076020526000908152604090205481565b6002805482908110610bb357fe5b6005805482908110610bb357fe5b6111d6610d6c565b6111de6114d7565b6111e6610bcd565b6111ee6116d9565b565b600154604080517f757991a80000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163757991a8916004808301926020929190829003018186803b15801561124e57600080fd5b505afa158015611262573d6000803e3d6000fd5b505050506040513d602081101561127857600080fd5b5051905090565b60086020526000908152604090205481565b600160a060020a0381166000908152600960205260408120600101541115610d6957600160a060020a0381166000908152600960205260409020600101546003116113bf5760005b6002548110156113bd5781600160a060020a03166002828154811015156112fc57fe5b600091825260209091200154600160a060020a031614156113b557600254600110156113a057805b6002546000190181101561139a57600280546001830190811061134357fe5b60009182526020909120015460028054600160a060020a03909216918390811061136957fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101611324565b50600019015b6002805460001901906113b39082611dda565b505b6001016112d9565b505b600160a060020a03811660009081526009602052604081208054600160a060020a0319168155600181018290556002018190555b600354811015610e575781600160a060020a031660038281548110151561141657fe5b600091825260209091200154600160a060020a031614156114cf57600354600110156114ba57805b600354600019018110156114b457600380546001830190811061145d57fe5b60009182526020909120015460038054600160a060020a03909216918390811061148357fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560010161143e565b50600019015b6003805460001901906114cd9082611dda565b505b6001016113f3565b6000600460008154811015156114e957fe5b6000918252602082200154600154604080517f0d6bb50f0000000000000000000000000000000000000000000000000000000081529051600160a060020a0393841695506060949390921692630d6bb50f92600480840193829003018186803b15801561155557600080fd5b505afa158015611569573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561159257600080fd5b8101908080516401000000008111156115aa57600080fd5b820160208101848111156115bd57600080fd5b81518560208202830111640100000000821117156115da57600080fd5b50909450600093505050505b815181101561167657600082828151811015156115ff57fe5b90602001906020020151905080600160a060020a031684600160a060020a0316141561166d5760018351038214156116595761165483600081518110151561164357fe5b906020019060200201516000611a56565b61166d565b61166d838360010181518110151561164357fe5b506001016115e6565b505050565b600154604080517f961c5c7a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163961c5c7a916004808301926020929190829003018186803b15801561124e57600080fd5b6005545b6000811115610d695760001981016116f481611d12565b50600019016116dd565b6004805482908110610bb357fe5b611714611dfe565b5060408051606081018252600160a060020a03831680825260016020808401829052438486015260009283526009905292902090910154156117fd57600160a060020a038216600090815260096020526040902060010154600311156117f857600160a060020a0382166000908152600960205260409020600181810180549091019081905543600290920191909155600314156117f857600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018054600160a060020a031916600160a060020a0384161790555b610e57565b600160a060020a039182166000818152600960209081526040808320855181549716600160a060020a03199788161781559185015160018084019190915594015160029091015560038054938401815590527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018054909216179055565b600160009054906101000a9004600160a060020a0316600160a060020a031663961c5c7a6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156118ce57600080fd5b505afa1580156118e2573d6000803e3d6000fd5b505050506040513d60208110156118f857600080fd5b50516005541015610d69576005805460018181019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018054600160a060020a038416600160a060020a0319909116811790915560009081526008602052604090205550565b6060806000600580549050604051908082528060200260200182016040528015611994578160200160208202803883390190505b5060058054604080516020808402820181019092528281529395508301828280156119e857602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116119ca575b5050505050925082519050909192565b600154604080517feb5c00110000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163eb5c0011916004808301926020929190829003018186803b15801561124e57600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611aa757600080fd5b505afa158015611abb573d6000803e3d6000fd5b505050506040513d6020811015611ad157600080fd5b50516004541015610e5757600160009054906101000a9004600160a060020a0316600160a060020a031663eb5c00116040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611b2d57600080fd5b505afa158015611b41573d6000803e3d6000fd5b505050506040513d6020811015611b5757600080fd5b5051811015611bdc5760076000600483815481101515611b7357fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556004805483919083908110611ba957fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611c28565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054600160a060020a031916600160a060020a0384161790555b50600160a060020a0316600090815260076020526040902060019055565b60045480821015610e575760076000600484815481101515611c6457fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055815b60018203811015611cfe576004805460018301908110611ca757fe5b60009182526020909120015460048054600160a060020a039092169183908110611ccd57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101611c8b565b506004805460001901906116769082611dda565b60055480821015610e575760086000600584815481101515611d3057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055815b60018203811015611dca576005805460018301908110611d7357fe5b60009182526020909120015460058054600160a060020a039092169183908110611d9957fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101611d57565b5060058054600019019061167690825b81548183558181111561167657600083815260209020611676918101908301611e29565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b610f4791905b80821115611e435760008155600101611e2f565b509056fea165627a7a72305820e4f653be4b9db1d948887a9db428d133c867c4931f3f5faa64f27c58279b931b0029`

// DeployTribeChief_1_0_0 deploys a new Ethereum contract, binding an instance of TribeChief_1_0_0 to it.
func DeployTribeChief_1_0_0(auth *bind.TransactOpts, backend bind.ContractBackend, baseAddress common.Address, pocAddress common.Address) (common.Address, *types.Transaction, *TribeChief_1_0_0, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_1_0_0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_1_0_0Bin), backend, baseAddress, pocAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_1_0_0{TribeChief_1_0_0Caller: TribeChief_1_0_0Caller{contract: contract}, TribeChief_1_0_0Transactor: TribeChief_1_0_0Transactor{contract: contract}}, nil
}

// TribeChief_1_0_0 is an auto generated Go binding around an Ethereum contract.
type TribeChief_1_0_0 struct {
	TribeChief_1_0_0Caller     // Read-only binding to the contract
	TribeChief_1_0_0Transactor // Write-only binding to the contract
}

// TribeChief_1_0_0Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_1_0_0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_1_0_0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_1_0_0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_1_0_0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_1_0_0Session struct {
	Contract     *TribeChief_1_0_0       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_1_0_0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_1_0_0CallerSession struct {
	Contract *TribeChief_1_0_0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_1_0_0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_1_0_0TransactorSession struct {
	Contract     *TribeChief_1_0_0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_1_0_0Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_1_0_0Raw struct {
	Contract *TribeChief_1_0_0 // Generic contract binding to access the raw methods on
}

// TribeChief_1_0_0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_1_0_0CallerRaw struct {
	Contract *TribeChief_1_0_0Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_1_0_0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_1_0_0TransactorRaw struct {
	Contract *TribeChief_1_0_0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_1_0_0 creates a new instance of TribeChief_1_0_0, bound to a specific deployed contract.
func NewTribeChief_1_0_0(address common.Address, backend bind.ContractBackend) (*TribeChief_1_0_0, error) {
	contract, err := bindTribeChief_1_0_0(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_1_0_0{TribeChief_1_0_0Caller: TribeChief_1_0_0Caller{contract: contract}, TribeChief_1_0_0Transactor: TribeChief_1_0_0Transactor{contract: contract}}, nil
}

// NewTribeChief_1_0_0Caller creates a new read-only instance of TribeChief_1_0_0, bound to a specific deployed contract.
func NewTribeChief_1_0_0Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_1_0_0Caller, error) {
	contract, err := bindTribeChief_1_0_0(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_1_0_0Caller{contract: contract}, nil
}

// NewTribeChief_1_0_0Transactor creates a new write-only instance of TribeChief_1_0_0, bound to a specific deployed contract.
func NewTribeChief_1_0_0Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_1_0_0Transactor, error) {
	contract, err := bindTribeChief_1_0_0(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_1_0_0Transactor{contract: contract}, nil
}

// bindTribeChief_1_0_0 binds a generic wrapper to an already deployed contract.
func bindTribeChief_1_0_0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_1_0_0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_1_0_0 *TribeChief_1_0_0Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_1_0_0.Contract.TribeChief_1_0_0Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_1_0_0 *TribeChief_1_0_0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.TribeChief_1_0_0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_1_0_0 *TribeChief_1_0_0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.TribeChief_1_0_0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_1_0_0.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.contract.Transact(opts, method, params...)
}

// BlackList is a free data retrieval call binding the contract method 0x609ed258.
//
// Solidity: function _blackList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) BlackList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "_blackList", arg0)
	return *ret0, err
}

// BlackList is a free data retrieval call binding the contract method 0x609ed258.
//
// Solidity: function _blackList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) BlackList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.BlackList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x609ed258.
//
// Solidity: function _blackList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) BlackList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.BlackList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlackMembers is a free data retrieval call binding the contract method 0x3aa07a36.
//
// Solidity: function _blackMembers( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) BlackMembers(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "_blackMembers", arg0)
	return *ret0, err
}

// BlackMembers is a free data retrieval call binding the contract method 0x3aa07a36.
//
// Solidity: function _blackMembers( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) BlackMembers(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.BlackMembers(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlackMembers is a free data retrieval call binding the contract method 0x3aa07a36.
//
// Solidity: function _blackMembers( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) BlackMembers(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.BlackMembers(&_TribeChief_1_0_0.CallOpts, arg0)
}

// SignerList is a free data retrieval call binding the contract method 0xb4deb2ca.
//
// Solidity: function _signerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) SignerList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "_signerList", arg0)
	return *ret0, err
}

// SignerList is a free data retrieval call binding the contract method 0xb4deb2ca.
//
// Solidity: function _signerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) SignerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.SignerList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// SignerList is a free data retrieval call binding the contract method 0xb4deb2ca.
//
// Solidity: function _signerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) SignerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.SignerList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// VolunteerList is a free data retrieval call binding the contract method 0x65ddc059.
//
// Solidity: function _volunteerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) VolunteerList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "_volunteerList", arg0)
	return *ret0, err
}

// VolunteerList is a free data retrieval call binding the contract method 0x65ddc059.
//
// Solidity: function _volunteerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) VolunteerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.VolunteerList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// VolunteerList is a free data retrieval call binding the contract method 0x65ddc059.
//
// Solidity: function _volunteerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) VolunteerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.VolunteerList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlMap is a free data retrieval call binding the contract method 0x113aa471.
//
// Solidity: function blMap( address) constant returns(addr address, score uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) BlMap(opts *bind.CallOptsWithNumber, arg0 common.Address) (struct {
	Addr   common.Address
	Score  *big.Int
	Number *big.Int
}, error) {
	ret := new(struct {
		Addr   common.Address
		Score  *big.Int
		Number *big.Int
	})
	out := ret
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "blMap", arg0)
	return *ret, err
}

// BlMap is a free data retrieval call binding the contract method 0x113aa471.
//
// Solidity: function blMap( address) constant returns(addr address, score uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) BlMap(arg0 common.Address) (struct {
	Addr   common.Address
	Score  *big.Int
	Number *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.BlMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlMap is a free data retrieval call binding the contract method 0x113aa471.
//
// Solidity: function blMap( address) constant returns(addr address, score uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) BlMap(arg0 common.Address) (struct {
	Addr   common.Address
	Score  *big.Int
	Number *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.BlMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) BlockNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "blockNumber")
	return *ret0, err
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) BlockNumber() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.BlockNumber(&_TribeChief_1_0_0.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) BlockNumber() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.BlockNumber(&_TribeChief_1_0_0.CallOpts)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) FilterVolunteer(opts *bind.CallOptsWithNumber, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "filterVolunteer", volunteers)
	return *ret0, err
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_1_0_0.Contract.FilterVolunteer(&_TribeChief_1_0_0.CallOpts, volunteers)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_1_0_0.Contract.FilterVolunteer(&_TribeChief_1_0_0.CallOpts, volunteers)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetEpoch(&_TribeChief_1_0_0.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetEpoch(&_TribeChief_1_0_0.CallOpts)
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() constant returns(address[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetLeader(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getLeader")
	return *ret0, err
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() constant returns(address[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetLeader() ([]common.Address, error) {
	return _TribeChief_1_0_0.Contract.GetLeader(&_TribeChief_1_0_0.CallOpts)
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() constant returns(address[])
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetLeader() ([]common.Address, error) {
	return _TribeChief_1_0_0.Contract.GetLeader(&_TribeChief_1_0_0.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetSignerLimit(&_TribeChief_1_0_0.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetSignerLimit(&_TribeChief_1_0_0.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetStatus(&_TribeChief_1_0_0.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetStatus(&_TribeChief_1_0_0.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetVolunteerLimit(&_TribeChief_1_0_0.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.GetVolunteerLimit(&_TribeChief_1_0_0.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetVolunteers(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getVolunteers")
	return *ret, err
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetVolunteers(&_TribeChief_1_0_0.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetVolunteers(&_TribeChief_1_0_0.CallOpts)
}

// SignersMap is a free data retrieval call binding the contract method 0x5c7f94b7.
//
// Solidity: function signersMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) SignersMap(opts *bind.CallOptsWithNumber, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "signersMap", arg0)
	return *ret0, err
}

// SignersMap is a free data retrieval call binding the contract method 0x5c7f94b7.
//
// Solidity: function signersMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) SignersMap(arg0 common.Address) (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.SignersMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// SignersMap is a free data retrieval call binding the contract method 0x5c7f94b7.
//
// Solidity: function signersMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) SignersMap(arg0 common.Address) (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.SignersMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) Version() (string, error) {
	return _TribeChief_1_0_0.Contract.Version(&_TribeChief_1_0_0.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) Version() (string, error) {
	return _TribeChief_1_0_0.Contract.Version(&_TribeChief_1_0_0.CallOpts)
}

// VolunteerMap is a free data retrieval call binding the contract method 0x7cc81e5a.
//
// Solidity: function volunteerMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) VolunteerMap(opts *bind.CallOptsWithNumber, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "volunteerMap", arg0)
	return *ret0, err
}

// VolunteerMap is a free data retrieval call binding the contract method 0x7cc81e5a.
//
// Solidity: function volunteerMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) VolunteerMap(arg0 common.Address) (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.VolunteerMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// VolunteerMap is a free data retrieval call binding the contract method 0x7cc81e5a.
//
// Solidity: function volunteerMap( address) constant returns(uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) VolunteerMap(arg0 common.Address) (*big.Int, error) {
	return _TribeChief_1_0_0.Contract.VolunteerMap(&_TribeChief_1_0_0.CallOpts, arg0)
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) GenSigners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "genSigners")
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GenSigners() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners is a paid mutator transaction binding the contract method 0x73d53ec7.
//
// Solidity: function genSigners() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) GenSigners() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_clean_all_signer is a paid mutator transaction binding the contract method 0x4b77da76.
//
// Solidity: function genSigners_clean_all_signer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) GenSigners_clean_all_signer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "genSigners_clean_all_signer")
}

// GenSigners_clean_all_signer is a paid mutator transaction binding the contract method 0x4b77da76.
//
// Solidity: function genSigners_clean_all_signer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GenSigners_clean_all_signer() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_clean_all_signer(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_clean_all_signer is a paid mutator transaction binding the contract method 0x4b77da76.
//
// Solidity: function genSigners_clean_all_signer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) GenSigners_clean_all_signer() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_clean_all_signer(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_clean_volunteer is a paid mutator transaction binding the contract method 0xaca680a2.
//
// Solidity: function genSigners_clean_volunteer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) GenSigners_clean_volunteer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "genSigners_clean_volunteer")
}

// GenSigners_clean_volunteer is a paid mutator transaction binding the contract method 0xaca680a2.
//
// Solidity: function genSigners_clean_volunteer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GenSigners_clean_volunteer() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_clean_volunteer(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_clean_volunteer is a paid mutator transaction binding the contract method 0xaca680a2.
//
// Solidity: function genSigners_clean_volunteer() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) GenSigners_clean_volunteer() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_clean_volunteer(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_set_leader is a paid mutator transaction binding the contract method 0x8c607d25.
//
// Solidity: function genSigners_set_leader() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) GenSigners_set_leader(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "genSigners_set_leader")
}

// GenSigners_set_leader is a paid mutator transaction binding the contract method 0x8c607d25.
//
// Solidity: function genSigners_set_leader() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GenSigners_set_leader() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_set_leader(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_set_leader is a paid mutator transaction binding the contract method 0x8c607d25.
//
// Solidity: function genSigners_set_leader() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) GenSigners_set_leader() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_set_leader(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_v2s is a paid mutator transaction binding the contract method 0x40267e93.
//
// Solidity: function genSigners_v2s() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) GenSigners_v2s(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "genSigners_v2s")
}

// GenSigners_v2s is a paid mutator transaction binding the contract method 0x40267e93.
//
// Solidity: function genSigners_v2s() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GenSigners_v2s() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_v2s(&_TribeChief_1_0_0.TransactOpts)
}

// GenSigners_v2s is a paid mutator transaction binding the contract method 0x40267e93.
//
// Solidity: function genSigners_v2s() returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) GenSigners_v2s() (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.GenSigners_v2s(&_TribeChief_1_0_0.TransactOpts)
}

// PushBlackMember is a paid mutator transaction binding the contract method 0xc8e0666a.
//
// Solidity: function pushBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) PushBlackMember(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "pushBlackMember", addr)
}

// PushBlackMember is a paid mutator transaction binding the contract method 0xc8e0666a.
//
// Solidity: function pushBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) PushBlackMember(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.PushBlackMember(&_TribeChief_1_0_0.TransactOpts, addr)
}

// PushBlackMember is a paid mutator transaction binding the contract method 0xc8e0666a.
//
// Solidity: function pushBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) PushBlackMember(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.PushBlackMember(&_TribeChief_1_0_0.TransactOpts, addr)
}

// PushVolunteer is a paid mutator transaction binding the contract method 0xd4d8c657.
//
// Solidity: function pushVolunteer(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) PushVolunteer(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "pushVolunteer", addr)
}

// PushVolunteer is a paid mutator transaction binding the contract method 0xd4d8c657.
//
// Solidity: function pushVolunteer(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) PushVolunteer(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.PushVolunteer(&_TribeChief_1_0_0.TransactOpts, addr)
}

// PushVolunteer is a paid mutator transaction binding the contract method 0xd4d8c657.
//
// Solidity: function pushVolunteer(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) PushVolunteer(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.PushVolunteer(&_TribeChief_1_0_0.TransactOpts, addr)
}

// RemoveBlackMember is a paid mutator transaction binding the contract method 0x8698514a.
//
// Solidity: function removeBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) RemoveBlackMember(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "removeBlackMember", addr)
}

// RemoveBlackMember is a paid mutator transaction binding the contract method 0x8698514a.
//
// Solidity: function removeBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) RemoveBlackMember(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.RemoveBlackMember(&_TribeChief_1_0_0.TransactOpts, addr)
}

// RemoveBlackMember is a paid mutator transaction binding the contract method 0x8698514a.
//
// Solidity: function removeBlackMember(addr address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) RemoveBlackMember(addr common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.RemoveBlackMember(&_TribeChief_1_0_0.TransactOpts, addr)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.Update(&_TribeChief_1_0_0.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_1_0_0 *TribeChief_1_0_0TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_1_0_0.Contract.Update(&_TribeChief_1_0_0.TransactOpts, volunteer)
}
