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
const TribeChief_0_0_3ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_moveSignersToBlacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_cleanVolunteerList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_cleanBlacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_3Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_3Bin = `0x606060405260408051908101604052600581527f302e302e33000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000204565b5060146001556005600355600f60045534156200006957600080fd5b60405162001349380380620013498339810160405280805160068054600160a060020a03191633600160a060020a031617905591909101905060008080808451935060008411156200012c57600092505b838310156200012657848381518110620000d057fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905591506200011a82600364010000000062000dc66200019e82021704565b600190920191620000ba565b62000193565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff191660011790556200019381600364010000000062000dc66200019e82021704565b5050505050620002d5565b600354600754101562000200576007805460018101620001bf838262000289565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024757805160ff191683800117855562000277565b8280016001018555821562000277579182015b82811115620002775782518255916020019190600101906200025a565b5062000285929150620002b5565b5090565b815481835581811511620002b057600083815260209020620002b0918101908301620002b5565b505050565b620002d291905b80821115620002855760008155600101620002bc565b90565b61106480620002e56000396000f3006060604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ceb2cef81146100be5780631c1b8772146100d6578063295cdd08146100f55780634e69d5601461010857806352eeea031461028a57806354fd4d501461029d57806355cd14c914610327578063757991a81461033d57806379fd787e14610362578063961c5c7a14610378578063b6e0d4981461038b578063eb5c00111461039e575b600080fd5b34156100c957600080fd5b6100d46004356103b1565b005b34156100e157600080fd5b6100d4600160a060020a03600435166103de565b341561010057600080fd5b6100d46105d6565b341561011357600080fd5b61011b610779565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b83811015610171578082015183820152602001610159565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b838110156101b0578082015183820152602001610198565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b838110156101ef5780820151838201526020016101d7565b50505050905001868103835289818151815260200191508051906020019060200280838360005b8381101561022e578082015183820152602001610216565b50505050905001868103825288818151815260200191508051906020019060200280838360005b8381101561026d578082015183820152602001610255565b505050509050019b50505050505050505050505060405180910390f35b341561029557600080fd5b6100d46109c9565b34156102a857600080fd5b6102b0610a27565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102ec5780820151838201526020016102d4565b50505050905090810190601f1680156103195780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561033257600080fd5b6100d4600435610ad0565b341561034857600080fd5b610350610afd565b60405190815260200160405180910390f35b341561036d57600080fd5b6100d4600435610b03565b341561038357600080fd5b610350610b30565b341561039657600080fd5b6100d4610b36565b34156103a957600080fd5b610350610b90565b33600160a060020a0381166000908152600a6020526040812060010154116103d857600080fd5b50600155565b33600160a060020a0381166000908152600a60205260408120600101549091829182901161040b57600080fd5b4360058190556001549011801561042d57506001544381151561042a57fe5b06155b1561044a5761043a6109c9565b610442610b36565b61044a6105d6565b60075460055481151561045957fe5b0692506002600060078581548110151561046f57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561056457600a60006007858154811015156104b057fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206007805491935090849081106104e657fe5b60009182526020909120015433600160a060020a0390811691161461055f5781546001901115610525578154600019018255600554600183015561055a565b61055160078481548110151561053757fe5b600091825260209091200154600160a060020a0316610b96565b61055a83610c10565b610564565b600382555b600160a060020a0384161561057c5761057c84610cf6565b6003546007541080156105925750600854600090115b156105d0576105c6600860008154811015156105aa57fe5b600091825260209091200154600160a060020a03166003610dc6565b6105d06000610e36565b50505050565b6000806105e1610fca565b6007549250600091508180808080876040518059106105fd5750595b908082528060200260200182016040525095506001880394505b600085106106ed57600780548690811061062d57fe5b600091825260209091200154600160a060020a03169350610652846000198a01610f11565b9250878310158061069a57506002600060078581548110151561067157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff165b806106aa57506106aa8684610f78565b156106b4576106e1565b828688815181106106c157fe5b602090810290910101526003880487106106da576106ed565b6001870196505b60001990940193610617565b600087111561076f57600091505b8782101561076f5785828151811061070f57fe5b906020019060200201519050801561076457610746600787848151811061073257fe5b906020019060200201518154811061053757fe5b61076486838151811061075557fe5b90602001906020020151610c10565b6001909101906106fb565b5050505050505050565b610781610fca565b610789610fca565b610791610fca565b610799610fca565b6107a1610fca565b60075460009081906040518059106107b65750595b90808252806020026020018201604052506007549094506040518059106107da5750595b90808252806020026020018201604052509250600090505b6007548110156108a457600a600060078381548110151561080f57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484828151811061084057fe5b6020908102909101015260078054600a916000918490811061085e57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061089257fe5b602090810290910101526001016107f2565b60088054806020026020016040519081016040528092919081815260200182805480156108fa57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108dc575b50505050509650600780548060200260200160405190810160405280929190818152602001828054801561095757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610939575b5050505050955060098054806020026020016040519081016040528092919081815260200182805480156109b457602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610996575b50505050509450600554915050909192939495565b60085460005b81811015610a1757600b60006008838154811015156109ea57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016109cf565b610a2360086000610fdc565b5050565b610a2f610fca565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ac55780601f10610a9a57610100808354040283529160200191610ac5565b820191906000526020600020905b815481529060010190602001808311610aa857829003601f168201915b505050505090505b90565b33600160a060020a0381166000908152600a602052604081206001015411610af757600080fd5b50600455565b60015490565b33600160a060020a0381166000908152600a602052604081206001015411610b2a57600080fd5b50600355565b60045490565b60095460005b81811015610b8457600c6000600983815481101515610b5757fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600101610b3c565b610a2360096000610fdc565b60035490565b600160a060020a0381166000908152600c60205260409020541515610c0d576009805460018101610bc78382610ffa565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600c9052604090204390555b50565b600754600081831015610cf157600a6000600785815481101515610c3057fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610cde576007805460018301908110610c7a57fe5b60009182526020909120015460078054600160a060020a039092169183908110610ca057fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610c5e565b6007805460001901906105d09082610ffa565b505050565b600454600854108015610d1f5750600160a060020a0381166000908152600b6020526040902054155b8015610d415750600160a060020a0381166000908152600c6020526040902054155b8015610d665750600160a060020a0381166000908152600a6020526040902060010154155b15610c0d576008805460018101610d7d8382610ffa565b5060009182526020808320919091018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600b90526040902043905550565b6003546007541015610a23576007805460018101610de48382610ffa565b5060009182526020808320919091018054600160a060020a03861673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600a90526040902081815543600191909101555050565b600854600081831015610cf157600b6000600885815481101515610e5657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610efe576008805460018301908110610e9a57fe5b60009182526020909120015460088054600160a060020a039092169183908110610ec057fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610e7e565b6008805460001901906105d09082610ffa565b600080808311610f245760009150610f71565b834442604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902090508281811515610f6d57fe5b0691505b5092915050565b600080600084511115610fc0575060005b8351811015610fc05782848281518110610f9f57fe5b906020019060200201511415610fb85760019150610f71565b600101610f89565b5060009392505050565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610c0d919061101a565b815481835581811511610cf157600083815260209020610cf19181019083015b610acd91905b808211156110345760008155600101611020565b50905600a165627a7a7230582003bb6871620b04ae83b46bbc853f346c0d82872b3e3487029010ca0b43f0edcf0029`

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

// CleanBlacklist is a paid mutator transaction binding the contract method 0xb6e0d498.
//
// Solidity: function _cleanBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) CleanBlacklist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "_cleanBlacklist")
}

// CleanBlacklist is a paid mutator transaction binding the contract method 0xb6e0d498.
//
// Solidity: function _cleanBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) CleanBlacklist() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.CleanBlacklist(&_TribeChief_0_0_3.TransactOpts)
}

// CleanBlacklist is a paid mutator transaction binding the contract method 0xb6e0d498.
//
// Solidity: function _cleanBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) CleanBlacklist() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.CleanBlacklist(&_TribeChief_0_0_3.TransactOpts)
}

// CleanVolunteerList is a paid mutator transaction binding the contract method 0x52eeea03.
//
// Solidity: function _cleanVolunteerList() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) CleanVolunteerList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "_cleanVolunteerList")
}

// CleanVolunteerList is a paid mutator transaction binding the contract method 0x52eeea03.
//
// Solidity: function _cleanVolunteerList() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) CleanVolunteerList() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.CleanVolunteerList(&_TribeChief_0_0_3.TransactOpts)
}

// CleanVolunteerList is a paid mutator transaction binding the contract method 0x52eeea03.
//
// Solidity: function _cleanVolunteerList() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) CleanVolunteerList() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.CleanVolunteerList(&_TribeChief_0_0_3.TransactOpts)
}

// MoveSignersToBlacklist is a paid mutator transaction binding the contract method 0x295cdd08.
//
// Solidity: function _moveSignersToBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Transactor) MoveSignersToBlacklist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_3.contract.Transact(opts, "_moveSignersToBlacklist")
}

// MoveSignersToBlacklist is a paid mutator transaction binding the contract method 0x295cdd08.
//
// Solidity: function _moveSignersToBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3Session) MoveSignersToBlacklist() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.MoveSignersToBlacklist(&_TribeChief_0_0_3.TransactOpts)
}

// MoveSignersToBlacklist is a paid mutator transaction binding the contract method 0x295cdd08.
//
// Solidity: function _moveSignersToBlacklist() returns()
func (_TribeChief_0_0_3 *TribeChief_0_0_3TransactorSession) MoveSignersToBlacklist() (*types.Transaction, error) {
	return _TribeChief_0_0_3.Contract.MoveSignersToBlacklist(&_TribeChief_0_0_3.TransactOpts)
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
