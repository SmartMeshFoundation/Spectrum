// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chieflib

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/SMChain/accounts/abi"
	"github.com/SmartMeshFoundation/SMChain/accounts/abi/bind"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/core/types"
)

// TribeChief_0_0_3ABI is the input ABI used to generate the binding from.
const TribeChief_0_0_3ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"blackList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"name\":\"TribeChief\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TribeChief_0_0_3Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_3Bin = `0x606060405260408051908101604052600581527f302e302e330000000000000000000000000000000000000000000000000000006020820152600090805161004b92916020019061006a565b50612d006001556003805560048055341561006557600080fd5b610105565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ab57805160ff19168380011785556100d8565b828001600101855582156100d8579182015b828111156100d85782518255916020019190600101906100bd565b506100e49291506100e8565b5090565b61010291905b808211156100e457600081556001016100ee565b90565b610d4d806101146000396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b877281146100925780634e69d560146100b357806354fd4d501461023557806355cd14c9146102bf5780636b1f71ad146102d557806379fd787e14610324578063961c5c7a1461033a578063eb5c00111461035f575b600080fd5b341561009d57600080fd5b6100b1600160a060020a0360043516610372565b005b34156100be57600080fd5b6100c661058e565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b8381101561011c578082015183820152602001610104565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b8381101561015b578082015183820152602001610143565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b8381101561019a578082015183820152602001610182565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156101d95780820151838201526020016101c1565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610218578082015183820152602001610200565b505050509050019b50505050505050505050505060405180910390f35b341561024057600080fd5b610248610781565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561028457808201518382015260200161026c565b50505050905090810190601f1680156102b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102ca57600080fd5b6100b160043561082a565b34156102e057600080fd5b6100b1600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061085795505050505050565b341561032f57600080fd5b6100b1600435610999565b341561034557600080fd5b61034d6109c6565b60405190815260200160405180910390f35b341561036a57600080fd5b61034d6109cc565b33600160a060020a0381166000908152600a602052604081206001015490918291829182918290116103a357600080fd5b436005819055600154901180156103c55750600154438115156103c257fe5b06155b1561042c576008549450600093505b8484101561042057600b60006008868154811015156103ef57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600193909301926103d4565b61042c60086000610cb3565b60075460055481151561043b57fe5b0692506002600060078581548110151561045157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561051a57600a600060078581548110151561049257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206007805491935090849081106104c857fe5b60009182526020909120015433600160a060020a0390811691161461051557815460019011156105075781546000190182556005546001830155610510565b610510836109d2565b61051a565b600382555b600160a060020a038616156105325761053286610aba565b6003546007541080156105485750600854600090115b156105865761057c6008600081548110151561056057fe5b600091825260209091200154600160a060020a03166003610b68565b6105866000610bd8565b505050505050565b610596610cd1565b61059e610cd1565b6105a6610cd1565b6105ae610cd1565b6105b6610cd1565b60075460009081906040518059106105cb5750595b90808252806020026020018201604052506007549095506040518059106105ef5750595b90808252806020026020018201604052509350600090505b6007548110156106b957600a600060078381548110151561062457fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205485828151811061065557fe5b6020908102909101015260078054600a916000918490811061067357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600101548482815181106106a757fe5b60209081029091010152600101610607565b600880548060200260200160405190810160405280929190818152602001828054801561070f57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106f1575b50505050509650600780548060200260200160405190810160405280929190818152602001828054801561076c57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161074e575b50505050509550600554915050909192939495565b610789610cd1565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561081f5780601f106107f45761010080835404028352916020019161081f565b820191906000526020600020905b81548152906001019060200180831161080257829003601f168201915b505050505090505b90565b33600160a060020a0381166000908152600a60205260408120600101541161085157600080fd5b50600455565b6006805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a0316179055600080808080808651955060008611156108f657600094505b858510156108f1578685815181106108ac57fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905593506108e6846003610b68565b600190940193610898565b610990565b5050734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff19166001179055905073ca012e2facf405885293d8d3ba3f14fae1e58b7173adb3ea3ad356199206ca817b04fd668cc5196df261097a836003610b68565b610985826003610b68565b610990816003610b68565b50505050505050565b33600160a060020a0381166000908152600a6020526040812060010154116109c057600080fd5b50600355565b60045490565b60035490565b600754600081831015610ab557600a60006007858154811015156109f257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610aa0576007805460018301908110610a3c57fe5b60009182526020909120015460078054600160a060020a039092169183908110610a6257fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610a20565b600780546000190190610ab39082610ce3565b505b505050565b600454600854108015610ae35750600160a060020a0381166000908152600b6020526040902054155b8015610b085750600160a060020a0381166000908152600a6020526040902060010154155b15610b65576008805460018101610b1f8382610ce3565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600b9052604090204390555b50565b6003546007541015610bd4576007805460018101610b868382610ce3565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b600854600081831015610ab557600b6000600885815481101515610bf857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610ca0576008805460018301908110610c3c57fe5b60009182526020909120015460088054600160a060020a039092169183908110610c6257fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610c20565b600880546000190190610ab39082610ce3565b5080546000825590600052602060002090810190610b659190610d03565b60206040519081016040526000815290565b815481835581811511610ab557600083815260209020610ab59181019083015b61082791905b80821115610d1d5760008155600101610d09565b50905600a165627a7a72305820c8e4c7531967b7a19be38ca6c3c2b5d869488cc6f5998af48e898e5e0e0397bb0029`

// DeployTribeChief_0_0_3 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_3 to it.
func DeployTribeChief_0_0_3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TribeChief_0_0_3, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_0_0_3{TribeChief_0_0_3Caller: TribeChief_0_0_3Caller{contract: contract}, TribeChief_0_0_3Transactor: TribeChief_0_0_3Transactor{contract: contract}}, nil
}

// TribeChief_0_0_3 is an auto generated Go binding around an Ethereum contract.
type TribeChief_0_0_3 struct {
	TribeChief_0_0_3Caller     // Read-only binding to the contract
	TribeChief_0_0_3Transactor // Write-only binding to the contract
}

// TribeChief_0_0_3Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_0_0_3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_0_0_3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_0_0_3Session struct {
	Contract     *TribeChief_0_0_3       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_0_0_3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_0_0_3CallerSession struct {
	Contract *TribeChief_0_0_3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_0_0_3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_0_0_3TransactorSession struct {
	Contract     *TribeChief_0_0_3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_0_0_3Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_0_0_3Raw struct {
	Contract *TribeChief_0_0_3 // Generic contract binding to access the raw methods on
}

// TribeChief_0_0_3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_0_0_3CallerRaw struct {
	Contract *TribeChief_0_0_3Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_0_0_3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_0_0_3TransactorRaw struct {
	Contract *TribeChief_0_0_3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_0_0_3 creates a new instance of TribeChief_0_0_3, bound to a specific deployed contract.
func NewTribeChief_0_0_3(address common.Address, backend bind.ContractBackend) (*TribeChief_0_0_3, error) {
	contract, err := bindTribeChief_0_0_3(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_3{TribeChief_0_0_3Caller: TribeChief_0_0_3Caller{contract: contract}, TribeChief_0_0_3Transactor: TribeChief_0_0_3Transactor{contract: contract}}, nil
}

// NewTribeChief_0_0_3Caller creates a new read-only instance of TribeChief_0_0_3, bound to a specific deployed contract.
func NewTribeChief_0_0_3Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_0_0_3Caller, error) {
	contract, err := bindTribeChief_0_0_3(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_3Caller{contract: contract}, nil
}

// NewTribeChief_0_0_3Transactor creates a new write-only instance of TribeChief_0_0_3, bound to a specific deployed contract.
func NewTribeChief_0_0_3Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_0_0_3Transactor, error) {
	contract, err := bindTribeChief_0_0_3(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_3Transactor{contract: contract}, nil
}

// bindTribeChief_0_0_3 binds a generic wrapper to an already deployed contract.
func bindTribeChief_0_0_3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_3 *TribeChief_0_0_3Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_3.Contract.TribeChief_0_0_3Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_3 *TribeChief_0_0_3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.TribeChief_0_0_3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_3 *TribeChief_0_0_3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.TribeChief_0_0_3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_3.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.contract.Transact(opts, method, params...)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetSignerLimit(&_TribeChief_0_0_3.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetSignerLimit(&_TribeChief_0_0_3.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], blackList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	BlackList     []*big.Int
	Number        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		SignerList    []common.Address
		ScoreList     []*big.Int
		NumberList    []*big.Int
		BlackList     []*big.Int
		Number        *big.Int
	})
	out := ret
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], blackList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	BlackList     []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_3.Contract.GetStatus(&_TribeChief_0_0_3.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], blackList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	BlackList     []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_3.Contract.GetStatus(&_TribeChief_0_0_3.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetVolunteerLimit(&_TribeChief_0_0_3.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetVolunteerLimit(&_TribeChief_0_0_3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) Version() (string, error) {
	return _TribeChief_0_0_3.Contract.Version(&_TribeChief_0_0_3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) Version() (string, error) {
	return _TribeChief_0_0_3.Contract.Version(&_TribeChief_0_0_3.CallOpts)
}

// TribeChief is a paid mutator transaction binding the contract method 0x6b1f71ad.
//
// Solidity: function TribeChief(genesisSigners address[]) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) TribeChief(opts *bind.TransactOpts, genesisSigners []common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "TribeChief", genesisSigners)
}

// TribeChief is a paid mutator transaction binding the contract method 0x6b1f71ad.
//
// Solidity: function TribeChief(genesisSigners address[]) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) TribeChief(genesisSigners []common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.TribeChief(&_TribeChief_0_0_3.TransactOpts, genesisSigners)
}

// TribeChief is a paid mutator transaction binding the contract method 0x6b1f71ad.
//
// Solidity: function TribeChief(genesisSigners address[]) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) TribeChief(genesisSigners []common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.TribeChief(&_TribeChief_0_0_3.TransactOpts, genesisSigners)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) SetSingerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "setSingerLimit", n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetSingerLimit(&_TribeChief_0_0_3.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetSingerLimit(&_TribeChief_0_0_3.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) SetVolunteerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "setVolunteerLimit", n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetVolunteerLimit(&_TribeChief_0_0_3.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetVolunteerLimit(&_TribeChief_0_0_3.TransactOpts, n)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.Update(&_TribeChief_0_0_3.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.Update(&_TribeChief_0_0_3.TransactOpts, volunteer)
}
