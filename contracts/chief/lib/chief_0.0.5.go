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

// TribeChief_0_0_5ABI is the input ABI used to generate the binding from.
const TribeChief_0_0_5ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setVolunteerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"setSingerLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_5Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_5Bin = `0x606060405260408051908101604052600581527f302e302e35000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000217565b506116806001556011600255604660035560015460045560025460065560035460075534156200007c57600080fd5b60405162001411380380620014118339810160405280805160098054600160a060020a03191633600160a060020a031617905591909101905060008080808451935060008411156200013f57600092505b838310156200013957848381518110620000e357fe5b90602001906020020151600160a060020a0381166000908152600560205260409020805460ff1916600117905591506200012d82600364010000000062000e7b620001b182021704565b600190920191620000cd565b620001a6565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260056020527fe972f0cd07b79b44c16c7f378a5866e05e66a2c3fdcdb77d23dbf43a635beeee805460ff19166001179055620001a681600364010000000062000e7b620001b182021704565b5050505050620002e8565b600654600a5410156200021357600a805460018101620001d283826200029c565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600d90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025a57805160ff19168380011785556200028a565b828001600101855582156200028a579182015b828111156200028a5782518255916020019190600101906200026d565b5062000298929150620002c8565b5090565b815481835581811511620002c357600083815260209020620002c3918101908301620002c8565b505050565b620002e591905b80821115620002985760008155600101620002cf565b90565b61111980620002f86000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ceb2cef811461009d5780631c1b8772146100b55780634e69d560146100d457806354fd4d501461025657806355cd14c9146102e0578063757991a8146102f657806379fd787e1461031b578063961c5c7a14610331578063eb5c001114610344575b600080fd5b34156100a857600080fd5b6100b3600435610357565b005b34156100c057600080fd5b6100b3600160a060020a0360043516610388565b34156100df57600080fd5b6100e7610580565b60405180806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b8381101561013d578082015183820152602001610125565b5050505090500186810385528b818151815260200191508051906020019060200280838360005b8381101561017c578082015183820152602001610164565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b838110156101bb5780820151838201526020016101a3565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156101fa5780820151838201526020016101e2565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610239578082015183820152602001610221565b505050509050019b50505050505050505050505060405180910390f35b341561026157600080fd5b6102696107d0565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a557808201518382015260200161028d565b50505050905090810190601f1680156102d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102eb57600080fd5b6100b3600435610879565b341561030157600080fd5b610309610924565b60405190815260200160405180910390f35b341561032657600080fd5b6100b360043561092a565b341561033c57600080fd5b6103096109de565b341561034f57600080fd5b6103096109e4565b6009543390600160a060020a0380831691161461037357600080fd5b60015482101561038257600080fd5b50600455565b33600160a060020a0381166000908152600d6020526040812060010154909182918290116103b557600080fd5b436008819055600454901180156103d75750600454438115156103d457fe5b06155b156103f4576103e46109ea565b6103ec610a48565b6103f4610aa2565b600a5460085481151561040357fe5b06925060056000600a8581548110151561041957fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16151561050e57600d6000600a8581548110151561045a57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600a8054919350908490811061049057fe5b60009182526020909120015433600160a060020a0390811691161461050957815460019011156104cf5781546000190182556008546001830155610504565b6104fb600a848154811015156104e157fe5b600091825260209091200154600160a060020a0316610c4b565b61050483610cc5565b61050e565b600382555b600160a060020a038416156105265761052684610dab565b600654600a5410801561053c5750600b54600090115b1561057a57610570600b600081548110151561055457fe5b600091825260209091200154600160a060020a03166003610e7b565b61057a6000610eeb565b50505050565b61058861107f565b61059061107f565b61059861107f565b6105a061107f565b6105a861107f565b600a5460009081906040518059106105bd5750595b9080825280602002602001820160405250600a549094506040518059106105e15750595b90808252806020026020018201604052509250600090505b600a548110156106ab57600d6000600a8381548110151561061657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484828151811061064757fe5b60209081029091010152600a8054600d916000918490811061066557fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061069957fe5b602090810290910101526001016105f9565b600b80548060200260200160405190810160405280929190818152602001828054801561070157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106e3575b50505050509650600a80548060200260200160405190810160405280929190818152602001828054801561075e57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610740575b50505050509550600c8054806020026020016040519081016040528092919081815260200182805480156107bb57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161079d575b50505050509450600854915050909192939495565b6107d861107f565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561086e5780601f106108435761010080835404028352916020019161086e565b820191906000526020600020905b81548152906001019060200180831161085157829003601f168201915b505050505090505b90565b60095460009081903390600160a060020a0380831691161461089a57600080fd5b6003548410156108a957600080fd5b600b5492508284101561091c578391505b8282101561090657600e6000600b848154811015156108d557fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600191909101906108ba565b600b805485850390039061091a9082611091565b505b505050600755565b60045490565b60095460009081903390600160a060020a0380831691161461094b57600080fd5b60025484101561095a57600080fd5b600a549250828410156109d6578391505b828210156109c057600d6000600a8481548110151561098657fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120818155600190810191909155919091019061096b565b600a80548585039003906109d49082611091565b505b505050600655565b60075490565b60065490565b600b5460005b81811015610a3857600e6000600b83815481101515610a0b57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016109f0565b610a44600b60006110b5565b5050565b600c5460005b81811015610a9657600f6000600c83815481101515610a6957fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600101610a4e565b610a44600c60006110b5565b6000610aac61107f565b600080600080600080600080985086604051805910610ac85750595b9080825280602002602001820160405250600a549098509650600095505b86861015610bbe57600a805487908110610afc57fe5b600091825260209091200154600160a060020a03169450610b21856000198901610fc6565b9350868410610b2f57610bb3565b600a805485908110610b3d57fe5b6000918252602080832090910154600160a060020a0316808352600590915260409091205490935060ff1615610b7257610bb3565b610b7c888561102d565b15610b8657610bb3565b83888a81518110610b9357fe5b60209081029091010152600387048910610bac57610bbe565b6001890198505b600190950194610ae6565b6000891115610c4057600091505b86821015610c4057878281518110610be057fe5b9060200190602002015190508015610c3557610c17600a898481518110610c0357fe5b90602001906020020151815481106104e157fe5b610c35888381518110610c2657fe5b90602001906020020151610cc5565b600190910190610bcc565b505050505050505050565b600160a060020a0381166000908152600f60205260409020541515610cc257600c805460018101610c7c8382611091565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091558252600f9052604090204390555b50565b600a54600081831015610da657600d6000600a85815481101515610ce557fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610d9357600a805460018301908110610d2f57fe5b600091825260209091200154600a8054600160a060020a039092169183908110610d5557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610d13565b600a8054600019019061057a9082611091565b505050565b600754600b54108015610dd45750600160a060020a0381166000908152600e6020526040902054155b8015610df65750600160a060020a0381166000908152600f6020526040902054155b8015610e1b5750600160a060020a0381166000908152600d6020526040902060010154155b15610cc257600b805460018101610e328382611091565b5060009182526020808320919091018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600e90526040902043905550565b600654600a541015610a4457600a805460018101610e998382611091565b5060009182526020808320919091018054600160a060020a03861673ffffffffffffffffffffffffffffffffffffffff1990911681179091558252600d90526040902081815543600191909101555050565b600b54600081831015610da657600e6000600b85815481101515610f0b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b60018203811015610fb357600b805460018301908110610f4f57fe5b600091825260209091200154600b8054600160a060020a039092169183908110610f7557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610f33565b600b8054600019019061057a9082611091565b600080808311610fd95760009150611026565b834442604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390209050828181151561102257fe5b0691505b5092915050565b600080600084511115611075575060005b8351811015611075578284828151811061105457fe5b90602001906020020151141561106d5760019150611026565b60010161103e565b5060009392505050565b60206040519081016040526000815290565b815481835581811511610da657600083815260209020610da69181019083016110cf565b5080546000825590600052602060002090810190610cc291905b61087691905b808211156110e957600081556001016110d5565b50905600a165627a7a72305820eecb79d8e9fc79e9f0b3270f551d2a6b05a13586a817989a1b94c467abbd11a80029`

// DeployTribeChief_0_0_5 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_5 to it.
func DeployTribeChief_0_0_5(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address) (common.Address, *types.Transaction, *TribeChief_0_0_5, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_5ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_5Bin), backend, genesisSigners)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_0_0_5{TribeChief_0_0_5Caller: TribeChief_0_0_5Caller{contract: contract}, TribeChief_0_0_5Transactor: TribeChief_0_0_5Transactor{contract: contract}}, nil
}

// TribeChief_0_0_5 is an auto generated Go binding around an Ethereum contract.
type TribeChief_0_0_5 struct {
	TribeChief_0_0_5Caller     // Read-only binding to the contract
	TribeChief_0_0_5Transactor // Write-only binding to the contract
}

// TribeChief_0_0_5Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_0_0_5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_0_0_5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_0_0_5Session struct {
	Contract     *TribeChief_0_0_5       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_0_0_5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_0_0_5CallerSession struct {
	Contract *TribeChief_0_0_5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_0_0_5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_0_0_5TransactorSession struct {
	Contract     *TribeChief_0_0_5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_0_0_5Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_0_0_5Raw struct {
	Contract *TribeChief_0_0_5 // Generic contract binding to access the raw methods on
}

// TribeChief_0_0_5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_0_0_5CallerRaw struct {
	Contract *TribeChief_0_0_5Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_0_0_5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_0_0_5TransactorRaw struct {
	Contract *TribeChief_0_0_5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_0_0_5 creates a new instance of TribeChief_0_0_5, bound to a specific deployed contract.
func NewTribeChief_0_0_5(address common.Address, backend bind.ContractBackend) (*TribeChief_0_0_5, error) {
	contract, err := bindTribeChief_0_0_5(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_5{TribeChief_0_0_5Caller: TribeChief_0_0_5Caller{contract: contract}, TribeChief_0_0_5Transactor: TribeChief_0_0_5Transactor{contract: contract}}, nil
}

// NewTribeChief_0_0_5Caller creates a new read-only instance of TribeChief_0_0_5, bound to a specific deployed contract.
func NewTribeChief_0_0_5Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_0_0_5Caller, error) {
	contract, err := bindTribeChief_0_0_5(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_5Caller{contract: contract}, nil
}

// NewTribeChief_0_0_5Transactor creates a new write-only instance of TribeChief_0_0_5, bound to a specific deployed contract.
func NewTribeChief_0_0_5Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_0_0_5Transactor, error) {
	contract, err := bindTribeChief_0_0_5(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_5Transactor{contract: contract}, nil
}

// bindTribeChief_0_0_5 binds a generic wrapper to an already deployed contract.
func bindTribeChief_0_0_5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_5 *TribeChief_0_0_5Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_5.Contract.TribeChief_0_0_5Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_5 *TribeChief_0_0_5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.TribeChief_0_0_5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_5 *TribeChief_0_0_5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.TribeChief_0_0_5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_5.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_5.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetEpoch(&_TribeChief_0_0_5.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetEpoch(&_TribeChief_0_0_5.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_5.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetSignerLimit(&_TribeChief_0_0_5.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetSignerLimit(&_TribeChief_0_0_5.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_0_0_5.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_5.Contract.GetStatus(&_TribeChief_0_0_5.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], blackList address[], scoreList uint256[], numberList uint256[], number uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	BlackList     []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief_0_0_5.Contract.GetStatus(&_TribeChief_0_0_5.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_5.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetVolunteerLimit(&_TribeChief_0_0_5.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_5.Contract.GetVolunteerLimit(&_TribeChief_0_0_5.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_0_0_5.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) Version() (string, error) {
	return _TribeChief_0_0_5.Contract.Version(&_TribeChief_0_0_5.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_5 *TribeChief_0_0_5CallerSession) Version() (string, error) {
	return _TribeChief_0_0_5.Contract.Version(&_TribeChief_0_0_5.CallOpts)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Transactor) SetEpoch(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.contract.Transact(opts, "setEpoch", n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetEpoch(&_TribeChief_0_0_5.TransactOpts, n)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorSession) SetEpoch(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetEpoch(&_TribeChief_0_0_5.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Transactor) SetSingerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.contract.Transact(opts, "setSingerLimit", n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetSingerLimit(&_TribeChief_0_0_5.TransactOpts, n)
}

// SetSingerLimit is a paid mutator transaction binding the contract method 0x79fd787e.
//
// Solidity: function setSingerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorSession) SetSingerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetSingerLimit(&_TribeChief_0_0_5.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Transactor) SetVolunteerLimit(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.contract.Transact(opts, "setVolunteerLimit", n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetVolunteerLimit(&_TribeChief_0_0_5.TransactOpts, n)
}

// SetVolunteerLimit is a paid mutator transaction binding the contract method 0x55cd14c9.
//
// Solidity: function setVolunteerLimit(n uint256) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorSession) SetVolunteerLimit(n *big.Int) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.SetVolunteerLimit(&_TribeChief_0_0_5.TransactOpts, n)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_5.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.Update(&_TribeChief_0_0_5.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_5 *TribeChief_0_0_5TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_5.Contract.Update(&_TribeChief_0_0_5.TransactOpts, volunteer)
}
