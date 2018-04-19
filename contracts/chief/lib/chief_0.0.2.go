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

// TribeChiefABI is the input ABI used to generate the binding from.
const TribeChiefABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChiefBin is the compiled bytecode used for deploying new contracts.
const TribeChiefBin = `0x606060405260408051908101604052600581527f302e302e32000000000000000000000000000000000000000000000000000000602082015260009080516200004d9291602001906200026a565b50612d00600155600380556004805534156200006857600080fd5b60405162000eae38038062000eae8339810160405280805160068054600160a060020a03191633600160a060020a0316179055919091019050600080808080808651955060008611156200012d57600094505b858510156200012757868581518110620000d157fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905593506200011b8460036401000000006200097e6200020482021704565b600190940193620000bb565b620001f7565b5050734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff19166001179055905073ca012e2facf405885293d8d3ba3f14fae1e58b7173adb3ea3ad356199206ca817b04fd668cc5196df2620001c18360036401000000006200097e6200020482021704565b620001dc8260036401000000006200097e6200020482021704565b620001f78160036401000000006200097e6200020482021704565b505050505050506200033b565b600354600754101562000266576007805460018101620002258382620002ef565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600990526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ad57805160ff1916838001178555620002dd565b82800160010185558215620002dd579182015b82811115620002dd578251825591602001919060010190620002c0565b50620002eb9291506200031b565b5090565b8154818355818115116200031657600083815260209020620003169181019083016200031b565b505050565b6200033891905b80821115620002eb576000815560010162000322565b90565b610b63806200034b6000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b877281146100875780634e69d560146100a857806354fd4d50146101e557806355cd14c91461026f57806379fd787e14610285578063961c5c7a1461029b578063eb5c0011146102c0575b600080fd5b341561009257600080fd5b6100a6600160a060020a03600435166102d3565b005b34156100b357600080fd5b6100bb6104ef565b604051808060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019060200280838360005b8381101561010d5780820151838201526020016100f5565b50505050905001858103845289818151815260200191508051906020019060200280838360005b8381101561014c578082015183820152602001610134565b50505050905001858103835288818151815260200191508051906020019060200280838360005b8381101561018b578082015183820152602001610173565b50505050905001858103825287818151815260200191508051906020019060200280838360005b838110156101ca5780820151838201526020016101b2565b50505050905001995050505050505050505060405180910390f35b34156101f057600080fd5b6101f86106d9565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561023457808201518382015260200161021c565b50505050905090810190601f1680156102615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561027a57600080fd5b6100a6600435610782565b341561029057600080fd5b6100a66004356107af565b34156102a657600080fd5b6102ae6107dc565b60405190815260200160405180910390f35b34156102cb57600080fd5b6102ae6107e2565b33600160a060020a038116600090815260096020526040812060010154909182918291829182901161030457600080fd5b4360058190556001549011801561032657506001544381151561032357fe5b06155b1561038d576008549450600093505b8484101561038157600a600060088681548110151561035057fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560019390930192610335565b61038d60086000610ac9565b60075460055481151561039c57fe5b069250600260006007858154811015156103b257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561047b57600960006007858154811015156103f357fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060078054919350908490811061042957fe5b60009182526020909120015433600160a060020a0390811691161461047657815460019011156104685781546000190182556005546001830155610471565b610471836107e8565b61047b565b600382555b600160a060020a0386161561049357610493866108d0565b6003546007541080156104a95750600854600090115b156104e7576104dd600860008154811015156104c157fe5b600091825260209091200154600160a060020a0316600361097e565b6104e760006109ee565b505050505050565b6104f7610ae7565b6104ff610ae7565b610507610ae7565b61050f610ae7565b60075460009081906040518059106105245750595b90808252806020026020018201604052506007549094506040518059106105485750595b90808252806020026020018201604052509250600090505b600754811015610612576009600060078381548110151561057d57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548482815181106105ae57fe5b6020908102909101015260078054600991600091849081106105cc57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061060057fe5b60209081029091010152600101610560565b600880548060200260200160405190810160405280929190818152602001828054801561066857602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161064a575b5050505050955060078054806020026020016040519081016040528092919081815260200182805480156106c557602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106a7575b505050505094506005549150509091929394565b6106e1610ae7565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107775780601f1061074c57610100808354040283529160200191610777565b820191906000526020600020905b81548152906001019060200180831161075a57829003601f168201915b505050505090505b90565b33600160a060020a038116600090815260096020526040812060010154116107a957600080fd5b50600455565b33600160a060020a038116600090815260096020526040812060010154116107d657600080fd5b50600355565b60045490565b60035490565b6007546000818310156108cb576009600060078581548110151561080857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b600182038110156108b657600780546001830190811061085257fe5b60009182526020909120015460078054600160a060020a03909216918390811061087857fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610836565b6007805460001901906108c99082610af9565b505b505050565b6004546008541080156108f95750600160a060020a0381166000908152600a6020526040902054155b801561091e5750600160a060020a038116600090815260096020526040902060010154155b1561097b5760088054600181016109358382610af9565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600a9052604090204390555b50565b60035460075410156109ea57600780546001810161099c8382610af9565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558252600990526040902081815543600191909101555b5050565b6008546000818310156108cb57600a6000600885815481101515610a0e57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610ab6576008805460018301908110610a5257fe5b60009182526020909120015460088054600160a060020a039092169183908110610a7857fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610a36565b6008805460001901906108c99082610af9565b508054600082559060005260206000209081019061097b9190610b19565b60206040519081016040526000815290565b8154818355818115116108cb576000838152602090206108cb9181019083015b61077f91905b80821115610b335760008155600101610b1f565b50905600a165627a7a723058207b4818c6a089263c8aaab5297423704e330f405a4a5c3020628e103f78924fcc0029`

// DeployTribeChief deploys a new Ethereum contract, binding an instance of TribeChief to it.
func DeployTribeChief(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address) (common.Address, *types.Transaction, *TribeChief, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChiefABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChiefBin), backend, genesisSigners)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief{TribeChiefCaller: TribeChiefCaller{contract: contract}, TribeChiefTransactor: TribeChiefTransactor{contract: contract}}, nil
}

// TribeChief is an auto generated Go binding around an Ethereum contract.
type TribeChief struct {
	TribeChiefCaller     // Read-only binding to the contract
	TribeChiefTransactor // Write-only binding to the contract
}

// TribeChiefCaller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChiefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChiefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChiefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChiefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChiefSession struct {
	Contract     *TribeChief             // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChiefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChiefCallerSession struct {
	Contract *TribeChiefCaller       // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChiefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChiefTransactorSession struct {
	Contract     *TribeChiefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TribeChiefRaw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChiefRaw struct {
	Contract *TribeChief // Generic contract binding to access the raw methods on
}

// TribeChiefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChiefCallerRaw struct {
	Contract *TribeChiefCaller // Generic read-only contract binding to access the raw methods on
}

// TribeChiefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChiefTransactorRaw struct {
	Contract *TribeChiefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief creates a new instance of TribeChief, bound to a specific deployed contract.
func NewTribeChief(address common.Address, backend bind.ContractBackend) (*TribeChief, error) {
	contract, err := bindTribeChief(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief{TribeChiefCaller: TribeChiefCaller{contract: contract}, TribeChiefTransactor: TribeChiefTransactor{contract: contract}}, nil
}

// NewTribeChiefCaller creates a new read-only instance of TribeChief, bound to a specific deployed contract.
func NewTribeChiefCaller(address common.Address, caller bind.ContractCaller) (*TribeChiefCaller, error) {
	contract, err := bindTribeChief(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChiefCaller{contract: contract}, nil
}

// NewTribeChiefTransactor creates a new write-only instance of TribeChief, bound to a specific deployed contract.
func NewTribeChiefTransactor(address common.Address, transactor bind.ContractTransactor) (*TribeChiefTransactor, error) {
	contract, err := bindTribeChief(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChiefTransactor{contract: contract}, nil
}

// bindTribeChief binds a generic wrapper to an already deployed contract.
func bindTribeChief(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChiefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief *TribeChiefRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.TribeChiefCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief *TribeChiefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief.Contract.TribeChiefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief *TribeChiefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief.Contract.TribeChiefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief *TribeChiefCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief *TribeChiefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief *TribeChiefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief.Contract.contract.Transact(opts, method, params...)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefCaller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief.Contract.GetSignerLimit(&_TribeChief.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefCallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief.Contract.GetSignerLimit(&_TribeChief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCaller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		SignerList    []common.Address
		ScoreList     []*big.Int
		NumberList    []*big.Int
		Number        *big.Int
	})
	out := ret
	err := _TribeChief.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefCaller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief.Contract.GetVolunteerLimit(&_TribeChief.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief *TribeChiefCallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief.Contract.GetVolunteerLimit(&_TribeChief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefCaller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefSession) Version() (string, error) {
	return _TribeChief.Contract.Version(&_TribeChief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefCallerSession) Version() (string, error) {
	return _TribeChief.Contract.Version(&_TribeChief.CallOpts)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief *TribeChiefTransactor) SetSingerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "setSingerLimit", n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief *TribeChiefSession) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.SetSingerLimit(&_TribeChief.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief *TribeChiefTransactorSession) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.SetSingerLimit(&_TribeChief.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief *TribeChiefTransactor) SetVolunteerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "setVolunteerLimit", n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief *TribeChiefSession) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.SetVolunteerLimit(&_TribeChief.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief *TribeChiefTransactorSession) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.SetVolunteerLimit(&_TribeChief.TransactOpts, n)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefTransactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefTransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts, volunteer)
}
