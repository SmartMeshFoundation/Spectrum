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

// TribeChief_0_0_4ABI is the input ABI used to generate the binding from.
const TribeChief_0_0_4ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_4Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_4Bin = `0x606060405260408051908101604052600581527f302e302e34000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000204565b5060146001556005600355600f60045534156200006957600080fd5b604051620012f5380380620012f58339810160405280805160068054600160a060020a03191633600160a060020a031617905591909101905060008080808451935060008411156200012c57600092505b838310156200012657848381518110620000d057fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff1916600117905591506200011a82600364010000000062000d726200019e82021704565b600190920191620000ba565b62000193565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff191660011790556200019381600364010000000062000d726200019e82021704565b5050505050620002d5565b600354600754101562000200576007805460018101620001bf838262000289565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600a90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024757805160ff191683800117855562000277565b8280016001018555821562000277579182015b82811115620002775782518255916020019190600101906200025a565b5062000285929150620002b5565b5090565b815481835581811511620002b057600083815260209020620002b0918101908301620002b5565b505050565b620002d291905b80821115620002855760008155600101620002bc565b90565b61101080620002e56000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ceb2cef811461009d5780631c1b8772146100b55780634e69d560146100d457806354fd4d501461025657806355cd14c9146102e0578063757991a8146102f657806379fd787e1461031b578063961c5c7a14610331578063eb5c001114610344575b600080fd5b34156100a857600080fd5b6100b3600435610357565b005b34156100c057600080fd5b6100b3600160a060020a0360043516610384565b34156100df57600080fd5b6100e761057c565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b8381101561013d578082015183820152602001610125565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b8381101561017c578082015183820152602001610164565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b838110156101bb5780820151838201526020016101a3565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156101fa5780820151838201526020016101e2565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610239578082015183820152602001610221565b505050509050019b50505050505050505050505060405180910390f35b341561026157600080fd5b6102696107cc565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a557808201518382015260200161028d565b50505050905090810190601f1680156102d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102eb57600080fd5b6100b3600435610875565b341561030157600080fd5b6103096108a2565b60405190815260200160405180910390f35b341561032657600080fd5b6100b36004356108a8565b341561033c57600080fd5b6103096108d5565b341561034f57600080fd5b6103096108db565b33600160a060020a0381166000908152600a60205260408120600101541161037e57600080fd5b50600155565b33600160a060020a0381166000908152600a6020526040812060010154909182918290116103b157600080fd5b436005819055600154901180156103d35750600154438115156103d057fe5b06155b156103f0576103e06108e1565b6103e861093f565b6103f0610999565b6007546005548115156103ff57fe5b0692506002600060078581548110151561041557fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561050a57600a600060078581548110151561045657fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060078054919350908490811061048c57fe5b60009182526020909120015433600160a060020a0390811691161461050557815460019011156104cb5781546000190182556005546001830155610500565b6104f76007848154811015156104dd57fe5b600091825260209091200154600160a060020a0316610b42565b61050083610bbc565b61050a565b600382555b600160a060020a038416156105225761052284610ca2565b6003546007541080156105385750600854600090115b156105765761056c6008600081548110151561055057fe5b600091825260209091200154600160a060020a03166003610d72565b6105766000610de2565b50505050565b610584610f76565b61058c610f76565b610594610f76565b61059c610f76565b6105a4610f76565b60075460009081906040518059106105b95750595b90808252806020026020018201604052506007549094506040518059106105dd5750595b90808252806020026020018201604052509250600090505b6007548110156106a757600a600060078381548110151561061257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484828151811061064357fe5b6020908102909101015260078054600a916000918490811061066157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061069557fe5b602090810290910101526001016105f5565b60088054806020026020016040519081016040528092919081815260200182805480156106fd57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106df575b50505050509650600780548060200260200160405190810160405280929190818152602001828054801561075a57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161073c575b5050505050955060098054806020026020016040519081016040528092919081815260200182805480156107b757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610799575b50505050509450600554915050909192939495565b6107d4610f76565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561086a5780601f1061083f5761010080835404028352916020019161086a565b820191906000526020600020905b81548152906001019060200180831161084d57829003601f168201915b505050505090505b90565b33600160a060020a0381166000908152600a60205260408120600101541161089c57600080fd5b50600455565b60015490565b33600160a060020a0381166000908152600a6020526040812060010154116108cf57600080fd5b50600355565b60045490565b60035490565b60085460005b8181101561092f57600b600060088381548110151561090257fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016108e7565b61093b60086000610f88565b5050565b60095460005b8181101561098d57600c600060098381548110151561096057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600101610945565b61093b60096000610f88565b60006109a3610f76565b6000806000806000806000809850866040518059106109bf5750595b90808252806020026020018201604052506007549098509650600095505b86861015610ab55760078054879081106109f357fe5b600091825260209091200154600160a060020a03169450610a18856000198901610ebd565b9350868410610a2657610aaa565b6007805485908110610a3457fe5b6000918252602080832090910154600160a060020a0316808352600290915260409091205490935060ff1615610a6957610aaa565b610a738885610f24565b15610a7d57610aaa565b83888a81518110610a8a57fe5b60209081029091010152600387048910610aa357610ab5565b6001890198505b6001909501946109dd565b6000891115610b3757600091505b86821015610b3757878281518110610ad757fe5b9060200190602002015190508015610b2c57610b0e6007898481518110610afa57fe5b90602001906020020151815481106104dd57fe5b610b2c888381518110610b1d57fe5b90602001906020020151610bbc565b600190910190610ac3565b505050505050505050565b600160a060020a0381166000908152600c60205260409020541515610bb9576009805460018101610b738382610fa6565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600c9052604090204390555b50565b600754600081831015610c9d57600a6000600785815481101515610bdc57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610c8a576007805460018301908110610c2657fe5b60009182526020909120015460078054600160a060020a039092169183908110610c4c57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610c0a565b6007805460001901906105769082610fa6565b505050565b600454600854108015610ccb5750600160a060020a0381166000908152600b6020526040902054155b8015610ced5750600160a060020a0381166000908152600c6020526040902054155b8015610d125750600160a060020a0381166000908152600a6020526040902060010154155b15610bb9576008805460018101610d298382610fa6565b5060009182526020808320919091018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600b90526040902043905550565b600354600754101561093b576007805460018101610d908382610fa6565b5060009182526020808320919091018054600160a060020a03861673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600a90526040902081815543600191909101555050565b600854600081831015610c9d57600b6000600885815481101515610e0257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610eaa576008805460018301908110610e4657fe5b60009182526020909120015460088054600160a060020a039092169183908110610e6c57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610e2a565b6008805460001901906105769082610fa6565b600080808311610ed05760009150610f1d565b834442604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902090508281811515610f1957fe5b0691505b5092915050565b600080600084511115610f6c575060005b8351811015610f6c5782848281518110610f4b57fe5b906020019060200201511415610f645760019150610f1d565b600101610f35565b5060009392505050565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610bb99190610fc6565b815481835581811511610c9d57600083815260209020610c9d9181019083015b61087291905b80821115610fe05760008155600101610fcc565b50905600a165627a7a72305820f4eb6bf3f470e46285f2169ab2a95c70d31f4e78ea14715eb79f29ebe64044d50029`

// DeployTribeChief_0_0_4 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_4 to it.
func DeployTribeChief_0_0_4(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address) (common.Address, *types.Transaction, *TribeChief_0_0_4, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_4ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_4Bin), backend, genesisSigners)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_0_0_4{TribeChief_0_0_4Caller: TribeChief_0_0_4Caller{contract: contract}, TribeChief_0_0_4Transactor: TribeChief_0_0_4Transactor{contract: contract}}, nil
}

// TribeChief_0_0_4 is an auto generated Go binding around an Ethereum contract.
type TribeChief_0_0_4 struct {
	TribeChief_0_0_4Caller     // Read-only binding to the contract
	TribeChief_0_0_4Transactor // Write-only binding to the contract
}

// TribeChief_0_0_4Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_0_0_4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_0_0_4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_0_0_4Session struct {
	Contract     *TribeChief_0_0_4       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_0_0_4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_0_0_4CallerSession struct {
	Contract *TribeChief_0_0_4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_0_0_4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_0_0_4TransactorSession struct {
	Contract     *TribeChief_0_0_4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_0_0_4Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_0_0_4Raw struct {
	Contract *TribeChief_0_0_4 // Generic contract binding to access the raw methods on
}

// TribeChief_0_0_4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_0_0_4CallerRaw struct {
	Contract *TribeChief_0_0_4Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_0_0_4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_0_0_4TransactorRaw struct {
	Contract *TribeChief_0_0_4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_0_0_4 creates a new instance of TribeChief_0_0_4, bound to a specific deployed contract.
func NewTribeChief_0_0_4(address common.Address, backend bind.ContractBackend) (*TribeChief_0_0_4, error) {
	contract, err := bindTribeChief_0_0_4(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_4{TribeChief_0_0_4Caller: TribeChief_0_0_4Caller{contract: contract}, TribeChief_0_0_4Transactor: TribeChief_0_0_4Transactor{contract: contract}}, nil
}

// NewTribeChief_0_0_4Caller creates a new read-only instance of TribeChief_0_0_4, bound to a specific deployed contract.
func NewTribeChief_0_0_4Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_0_0_4Caller, error) {
	contract, err := bindTribeChief_0_0_4(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_4Caller{contract: contract}, nil
}

// NewTribeChief_0_0_4Transactor creates a new write-only instance of TribeChief_0_0_4, bound to a specific deployed contract.
func NewTribeChief_0_0_4Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_0_0_4Transactor, error) {
	contract, err := bindTribeChief_0_0_4(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_4Transactor{contract: contract}, nil
}

// bindTribeChief_0_0_4 binds a generic wrapper to an already deployed contract.
func bindTribeChief_0_0_4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_4 *TribeChief_0_0_4Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_4.Contract.TribeChief_0_0_4Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_4 *TribeChief_0_0_4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.TribeChief_0_0_4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_4 *TribeChief_0_0_4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.TribeChief_0_0_4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_4.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_4.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetEpoch(&_TribeChief_0_0_4.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetEpoch(&_TribeChief_0_0_4.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_4.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetSignerLimit(&_TribeChief_0_0_4.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetSignerLimit(&_TribeChief_0_0_4.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_0_0_4.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_4.Contract.GetStatus(&_TribeChief_0_0_4.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_4.Contract.GetStatus(&_TribeChief_0_0_4.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_4.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetVolunteerLimit(&_TribeChief_0_0_4.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_4.Contract.GetVolunteerLimit(&_TribeChief_0_0_4.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_0_0_4.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) Version() (string, error) {
	return _TribeChief_0_0_4.Contract.Version(&_TribeChief_0_0_4.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_4 *TribeChief_0_0_4CallerSession) Version() (string, error) {
	return _TribeChief_0_0_4.Contract.Version(&_TribeChief_0_0_4.CallOpts)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Transactor) SetEpoch(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.contract.Transact(opts, "setEpoch", n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetEpoch(&_TribeChief_0_0_4.TransactOpts, n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorSession) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetEpoch(&_TribeChief_0_0_4.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Transactor) SetSingerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.contract.Transact(opts, "setSingerLimit", n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetSingerLimit(&_TribeChief_0_0_4.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorSession) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetSingerLimit(&_TribeChief_0_0_4.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Transactor) SetVolunteerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.contract.Transact(opts, "setVolunteerLimit", n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetVolunteerLimit(&_TribeChief_0_0_4.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorSession) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.SetVolunteerLimit(&_TribeChief_0_0_4.TransactOpts, n)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_4.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.Update(&_TribeChief_0_0_4.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_4 *TribeChief_0_0_4TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_4.Contract.Update(&_TribeChief_0_0_4.TransactOpts, volunteer)
}
