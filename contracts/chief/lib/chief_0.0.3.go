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
const TribeChief_0_0_3ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_3Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_3Bin = `0x606060405260408051908101604052600581527f302e302e33000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000205565b50612d006001556005600355600f60045534156200006a57600080fd5b60405162001236380380620012368339810160405280805160068054600160a060020a03191633600160a060020a031617905591909101905060008080808451935060008411156200012d57600092505b838310156200012757848381518110620000d157fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905591506200011b82600364010000000062000d6b6200019f82021704565b600190920191620000bb565b62000194565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff191660011790556200019481600364010000000062000d6b6200019f82021704565b5050505050620002d6565b600354600754101562000201576007805460018101620001c083826200028a565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024857805160ff191683800117855562000278565b8280016001018555821562000278579182015b82811115620002785782518255916020019190600101906200025b565b5062000286929150620002b6565b5090565b815481835581811511620002b157600083815260209020620002b1918101908301620002b6565b505050565b620002d391905b80821115620002865760008155600101620002bd565b90565b610f5080620002e66000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b877281146100875780634e69d560146100a857806354fd4d501461022a57806355cd14c9146102b457806379fd787e146102ca578063961c5c7a146102e0578063eb5c001114610305575b600080fd5b341561009257600080fd5b6100a6600160a060020a0360043516610318565b005b34156100b357600080fd5b6100bb610731565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b838110156101115780820151838201526020016100f9565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b83811015610150578082015183820152602001610138565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b8381101561018f578082015183820152602001610177565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156101ce5780820151838201526020016101b6565b50505050905001868103825288818151815260200191508051906020019060200280838360005b8381101561020d5780820151838201526020016101f5565b505050509050019b50505050505050505050505060405180910390f35b341561023557600080fd5b61023d610981565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610279578082015183820152602001610261565b50505050905090810190601f1680156102a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102bf57600080fd5b6100a6600435610a2a565b34156102d557600080fd5b6100a6600435610a57565b34156102eb57600080fd5b6102f3610a84565b60405190815260200160405180910390f35b341561031057600080fd5b6102f3610a8a565b600080600080600080610329610eb6565b600080600080600080336000600a600083600160a060020a0316600160a060020a031681526020019081526020016000206001015411151561036a57600080fd5b4360058190556001549011801561038c57506001544381151561038957fe5b06155b156105b4576008549d5060009c505b8d8d10156103e757600b600060088f8154811015156103b657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560019c909c019b61039b565b6103f360086000610ec8565b6009549b5060009a505b8b8b101561044957600c600060098d81548110151561041857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205560019a909a01996103fd565b61045560096000610ec8565b6007549950600098506000198a0196505b6000871061053457600780548890811061047c57fe5b600091825260209091200154600160a060020a031695506104a1866000198c01610a90565b9450600260006007878154811015156104b657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16156104e857610528565b6104f28886610ae7565b156104fc57610528565b84888a8151811061050957fe5b6020908102919091010152600189019860038b04901061052857610534565b60001990960195610466565b6000885111156105b457600093505b87518410156105b45761058b600789868151811061055d57fe5b906020019060200201518154811061057157fe5b600091825260209091200154600160a060020a0316610b39565b6105a988858151811061059a57fe5b90602001906020020151610bb3565b600190930192610543565b6007546005548115156105c357fe5b069250600260006007858154811015156105d957fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615156106b457600a600060078581548110151561061a57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060078054919350908490811061065057fe5b60009182526020909120015433600160a060020a039081169116146106af578154600190111561068f57815460001901825560055460018301556106aa565b6106a160078481548110151561057157fe5b6106aa83610bb3565b6106b4565b600382555b600160a060020a038f16156106cc576106cc8f610c9b565b6003546007541080156106e25750600854600090115b1561072057610716600860008154811015156106fa57fe5b600091825260209091200154600160a060020a03166003610d6b565b6107206000610ddb565b505050505050505050505050505050565b610739610eb6565b610741610eb6565b610749610eb6565b610751610eb6565b610759610eb6565b600754600090819060405180591061076e5750595b90808252806020026020018201604052506007549094506040518059106107925750595b90808252806020026020018201604052509250600090505b60075481101561085c57600a60006007838154811015156107c757fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548482815181106107f857fe5b6020908102909101015260078054600a916000918490811061081657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061084a57fe5b602090810290910101526001016107aa565b60088054806020026020016040519081016040528092919081815260200182805480156108b257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610894575b50505050509650600780548060200260200160405190810160405280929190818152602001828054801561090f57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108f1575b50505050509550600980548060200260200160405190810160405280929190818152602001828054801561096c57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161094e575b50505050509450600554915050909192939495565b610989610eb6565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a1f5780601f106109f457610100808354040283529160200191610a1f565b820191906000526020600020905b815481529060010190602001808311610a0257829003601f168201915b505050505090505b90565b33600160a060020a0381166000908152600a602052604081206001015411610a5157600080fd5b50600455565b33600160a060020a0381166000908152600a602052604081206001015411610a7e57600080fd5b50600355565b60045490565b60035490565b600080834442604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902090508281811515610adc57fe5b0691505b5092915050565b600080600084511115610b2f575060005b8351811015610b2f5782848281518110610b0e57fe5b906020019060200201511415610b275760019150610ae0565b600101610af8565b5060009392505050565b600160a060020a0381166000908152600c60205260409020541515610bb0576009805460018101610b6a8382610ee6565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600c9052604090204390555b50565b600754600081831015610c9657600a6000600785815481101515610bd357fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610c81576007805460018301908110610c1d57fe5b60009182526020909120015460078054600160a060020a039092169183908110610c4357fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610c01565b600780546000190190610c949082610ee6565b505b505050565b600454600854108015610cc45750600160a060020a0381166000908152600b6020526040902054155b8015610ce65750600160a060020a0381166000908152600c6020526040902054155b8015610d0b5750600160a060020a0381166000908152600a6020526040902060010154155b15610bb0576008805460018101610d228382610ee6565b5060009182526020808320919091018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600b90526040902043905550565b6003546007541015610dd7576007805460018101610d898382610ee6565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b600854600081831015610c9657600b6000600885815481101515610dfb57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610ea3576008805460018301908110610e3f57fe5b60009182526020909120015460088054600160a060020a039092169183908110610e6557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610e23565b600880546000190190610c949082610ee6565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610bb09190610f06565b815481835581811511610c9657600083815260209020610c969181019083015b610a2791905b80821115610f205760008155600101610f0c565b50905600a165627a7a7230582007ba4eeced42ef961e2458c7ba80886eb4b23e42ee9c33934c09c7ecc04866a80029`

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
