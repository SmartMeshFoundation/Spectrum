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
const TribeChief_0_0_3ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_3Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_3Bin = `0x606060405260408051908101604052600581527f302e302e33000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000204565b5060146001556005600355600f60045534156200006957600080fd5b604051620012bf380380620012bf8339810160405280805160068054600160a060020a03191633600160a060020a031617905591909101905060008080808451935060008411156200012c57600092505b838310156200012657848381518110620000d057fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905591506200011a82600364010000000062000df56200019e82021704565b600190920191620000ba565b62000193565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff191660011790556200019381600364010000000062000df56200019e82021704565b5050505050620002d5565b600354600754101562000200576007805460018101620001bf838262000289565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024757805160ff191683800117855562000277565b8280016001018555821562000277579182015b82811115620002775782518255916020019190600101906200025a565b5062000285929150620002b5565b5090565b815481835581811511620002b057600083815260209020620002b0918101908301620002b5565b505050565b620002d291905b80821115620002855760008155600101620002bc565b90565b610fda80620002e56000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ceb2cef811461009d5780631c1b8772146100b55780634e69d560146100d457806354fd4d501461025657806355cd14c9146102e0578063757991a8146102f657806379fd787e1461031b578063961c5c7a14610331578063eb5c001114610344575b600080fd5b34156100a857600080fd5b6100b3600435610357565b005b34156100c057600080fd5b6100b3600160a060020a0360043516610384565b34156100df57600080fd5b6100e76107a5565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b8381101561013d578082015183820152602001610125565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b8381101561017c578082015183820152602001610164565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b838110156101bb5780820151838201526020016101a3565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156101fa5780820151838201526020016101e2565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610239578082015183820152602001610221565b505050509050019b50505050505050505050505060405180910390f35b341561026157600080fd5b6102696109f5565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a557808201518382015260200161028d565b50505050905090810190601f1680156102d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102eb57600080fd5b6100b3600435610a9e565b341561030157600080fd5b610309610acb565b60405190815260200160405180910390f35b341561032657600080fd5b6100b3600435610ad1565b341561033c57600080fd5b610309610afe565b341561034f57600080fd5b610309610b04565b33600160a060020a0381166000908152600a60205260408120600101541161037e57600080fd5b50600155565b600080600080600080610395610f40565b600080600080600080336000600a600083600160a060020a0316600160a060020a03168152602001908152602001600020600101541115156103d657600080fd5b436005819055600154901180156103f85750600154438115156103f557fe5b06155b15610628576008549d5060009c505b8d8d101561045357600b600060088f81548110151561042257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560019c909c019b610407565b61045f60086000610f52565b6009549b5060009a505b8b8b10156104b557600c600060098d81548110151561048457fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560019a909a0199610469565b6104c160096000610f52565b6007549950600098506000198a0196505b600087106105a85760078054889081106104e857fe5b600091825260209091200154600160a060020a0316955061050d866000198c01610b0a565b9450898510158061055557506002600060078781548110151561052c57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff165b8061056557506105658886610b71565b1561056f5761059c565b84888a8151811061057c57fe5b6020908102909101015260038a048910610595576105a8565b6001890198505b600019909601956104d2565b60008851111561062857600093505b8751841015610628576105ff60078986815181106105d157fe5b90602001906020020151815481106105e557fe5b600091825260209091200154600160a060020a0316610bc3565b61061d88858151811061060e57fe5b90602001906020020151610c3d565b6001909301926105b7565b60075460055481151561063757fe5b0692506002600060078581548110151561064d57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561072857600a600060078581548110151561068e57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206007805491935090849081106106c457fe5b60009182526020909120015433600160a060020a039081169116146107235781546001901115610703578154600019018255600554600183015561071e565b6107156007848154811015156105e557fe5b61071e83610c3d565b610728565b600382555b600160a060020a038f1615610740576107408f610d25565b6003546007541080156107565750600854600090115b156107945761078a6008600081548110151561076e57fe5b600091825260209091200154600160a060020a03166003610df5565b6107946000610e65565b505050505050505050505050505050565b6107ad610f40565b6107b5610f40565b6107bd610f40565b6107c5610f40565b6107cd610f40565b60075460009081906040518059106107e25750595b90808252806020026020018201604052506007549094506040518059106108065750595b90808252806020026020018201604052509250600090505b6007548110156108d057600a600060078381548110151561083b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484828151811061086c57fe5b6020908102909101015260078054600a916000918490811061088a57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600101548382815181106108be57fe5b6020908102909101015260010161081e565b600880548060200260200160405190810160405280929190818152602001828054801561092657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610908575b50505050509650600780548060200260200160405190810160405280929190818152602001828054801561098357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610965575b5050505050955060098054806020026020016040519081016040528092919081815260200182805480156109e057602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116109c2575b50505050509450600554915050909192939495565b6109fd610f40565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a935780601f10610a6857610100808354040283529160200191610a93565b820191906000526020600020905b815481529060010190602001808311610a7657829003601f168201915b505050505090505b90565b33600160a060020a0381166000908152600a602052604081206001015411610ac557600080fd5b50600455565b60015490565b33600160a060020a0381166000908152600a602052604081206001015411610af857600080fd5b50600355565b60045490565b60035490565b600080808311610b1d5760009150610b6a565b834442604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902090508281811515610b6657fe5b0691505b5092915050565b600080600084511115610bb9575060005b8351811015610bb95782848281518110610b9857fe5b906020019060200201511415610bb15760019150610b6a565b600101610b82565b5060009392505050565b600160a060020a0381166000908152600c60205260409020541515610c3a576009805460018101610bf48382610f70565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600c9052604090204390555b50565b600754600081831015610d2057600a6000600785815481101515610c5d57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610d0b576007805460018301908110610ca757fe5b60009182526020909120015460078054600160a060020a039092169183908110610ccd57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610c8b565b600780546000190190610d1e9082610f70565b505b505050565b600454600854108015610d4e5750600160a060020a0381166000908152600b6020526040902054155b8015610d705750600160a060020a0381166000908152600c6020526040902054155b8015610d955750600160a060020a0381166000908152600a6020526040902060010154155b15610c3a576008805460018101610dac8382610f70565b5060009182526020808320919091018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600b90526040902043905550565b6003546007541015610e61576007805460018101610e138382610f70565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b600854600081831015610d2057600b6000600885815481101515610e8557fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610f2d576008805460018301908110610ec957fe5b60009182526020909120015460088054600160a060020a039092169183908110610eef57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610ead565b600880546000190190610d1e9082610f70565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610c3a9190610f90565b815481835581811511610d2057600083815260209020610d209181019083015b610a9b91905b80821115610faa5760008155600101610f96565b50905600a165627a7a7230582042b7168d63eb1475f5ade7365476f618dd4b75d3a7ee7d3ebe6e1e9389944ab50029`

// DeployTribeChief_0_0_3 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_3 to it.
func DeployTribeChief_0_0_3(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address) (common.Address, *types.Transaction, *TribeChief_0_0_3, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_3Bin), backend, genesisSigners)
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

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetEpoch(&_TribeChief_0_0_3.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_3.Contract.GetEpoch(&_TribeChief_0_0_3.CallOpts)
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
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		SignerList    []common.Address
		BlackList     []common.Address
		ScoreList     []*big.Int
		NumberList    []*big.Int
		Number        *big.Int
	})
	out := ret
	err := _TribeChief_0_0_3.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_3.Contract.GetStatus(&_TribeChief_0_0_3.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_3 *TribeChief_0_0_3CallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
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

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) SetEpoch(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "setEpoch", n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetEpoch(&_TribeChief_0_0_3.TransactOpts, n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.SetEpoch(&_TribeChief_0_0_3.TransactOpts, n)
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
