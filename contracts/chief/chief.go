// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chief

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/SMChain/accounts/abi"
	"github.com/SmartMeshFoundation/SMChain/accounts/abi/bind"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/core/types"
)

// TribeChiefABI is the input ABI used to generate the binding from.
const TribeChiefABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"int256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"score\",\"type\":\"int256\"}],\"name\":\"pushSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChiefBin is the compiled bytecode used for deploying new contracts.
const TribeChiefBin = `0x60606040526032600155600f600255341561001957600080fd5b60008054600160a060020a03191633600160a060020a0316178155600360208190527fb0d6da0981f9184e68c3a74cc6d81b2867cfa74ee652be6f621d78a7c1b3ede3805460ff1990811660019081179092557f5746fcaa35b8fe5d9c2023cf25b71497594d3c71f516ab9e4382a7895cde638880548216831790557f9c44b6fb5e3098c41faafd01076b84e13ff656f0882bfe8c08d14f839329c7e7805482168317905573adb3ea3ad356199206ca817b04fd668cc5196df2938490527fe8505a5ac9cb1417183861875ba9240bf4785e37ddba4197fc18d4309f87a24f8054909116909117905573ecce1549c2c803996e69930c97594525294de68f9173922316aefd06ae15ec9a797c88c695eaa4c3c32b9173ec61e6c0c17930fe1d8538de2c0b25644c542632919061015d9085906401000000006101c881026103d51704565b505061017d8360036101c8640100000000026103d5176401000000009004565b505061019d8260036101c8640100000000026103d5176401000000009004565b50506101bd8160036101c8640100000000026103d5176401000000009004565b5050505050506102da565b600580546000918291606590106101e55760009250829150610288565b80548190600181016101f78382610290565b5060009182526020909120018054600160a060020a031916600160a060020a03871617905560018181018054909181016102318382610290565b506000918252602090912001849055600281018054600181016102548382610290565b506000918252602080832043920191909155600160a060020a03871682526003830190526040902084905580549250600191505b509250929050565b8154818355818115116102b4576000838152602090206102b49181019083016102b9565b505050565b6102d791905b808211156102d357600081556001016102bf565b5090565b90565b61054c806102e96000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634e69d560811461005b578063a2e620451461014c578063da371bd214610161575b600080fd5b341561006657600080fd5b61006e6101aa565b60405180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b838110156100b657808201518382015260200161009e565b50505050905001848103835286818151815260200191508051906020019060200280838360005b838110156100f55780820151838201526020016100dd565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561013457808201518382015260200161011c565b50505050905001965050505050505060405180910390f35b341561015757600080fd5b61015f6102e2565b005b341561016c57600080fd5b61019073ffffffffffffffffffffffffffffffffffffffff600435166024356103d5565b604051918252151560208201526040908101905180910390f35b6101b26104c4565b6101ba6104c4565b6101c26104c4565b60058054600690600790839060208082020160405190810160405280929190818152602001828054801561022c57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610201575b505050505092508180548060200260200160405190810160405280929190818152602001828054801561027e57602002820191906000526020600020905b81548152602001906001019080831161026a575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156102d057602002820191906000526020600020905b8154815260200190600101908083116102bc575b50505050509050925092509250909192565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604081205490919082901361031857600080fd5b60068054600090811061032757fe5b60009182526020822001549250821315610382576006805460001990930192839190600090811061035457fe5b600091825260208220019190915560078054439290811061037157fe5b6000918252602090912001556103d1565b60058054600090811061039157fe5b60009182526020822001805473ffffffffffffffffffffffffffffffffffffffff191690556006805490919081106103c557fe5b60009182526020822001555b5050565b600580546000918291606590106103f257600092508291506104bc565b805481906001810161040483826104d6565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff8716179055600181810180549091810161045883826104d6565b5060009182526020909120018490556002810180546001810161047b83826104d6565b50600091825260208083204392019190915573ffffffffffffffffffffffffffffffffffffffff871682526003830190526040902084905580549250600191505b509250929050565b60206040519081016040526000815290565b8154818355818115116104fa576000838152602090206104fa9181019083016104ff565b505050565b61051d91905b808211156105195760008155600101610505565b5090565b905600a165627a7a723058202d13d364eaa7f6d08feddc015f27378ad0bbaac9908143eed93ccc37605733250029`

// DeployTribeChief deploys a new Ethereum contract, binding an instance of TribeChief to it.
func DeployTribeChief(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TribeChief, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChiefABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChiefBin), backend)
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
	Contract     *TribeChief       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TribeChiefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChiefCallerSession struct {
	Contract *TribeChiefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
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
func (_TribeChief *TribeChiefRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.TribeChiefCaller.contract.Call(opts, result, method, params...)
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
func (_TribeChief *TribeChiefCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.contract.Call(opts, result, method, params...)
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

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(address[], int256[], uint256[])
func (_TribeChief *TribeChiefCaller) GetStatus(opts *bind.CallOpts) ([]common.Address, []*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
		ret2 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TribeChief.contract.Call(opts, out, "getStatus")
	return *ret0, *ret1, *ret2, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(address[], int256[], uint256[])
func (_TribeChief *TribeChiefSession) GetStatus() ([]common.Address, []*big.Int, []*big.Int, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(address[], int256[], uint256[])
func (_TribeChief *TribeChiefCallerSession) GetStatus() ([]common.Address, []*big.Int, []*big.Int, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// PushSigner is a paid mutator transaction binding the contract method 0xda371bd2.
//
// Solidity: function pushSigner(signer address, score int256) returns(uint256, bool)
func (_TribeChief *TribeChiefTransactor) PushSigner(opts *bind.TransactOpts, signer common.Address, score *big.Int) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "pushSigner", signer, score)
}

// PushSigner is a paid mutator transaction binding the contract method 0xda371bd2.
//
// Solidity: function pushSigner(signer address, score int256) returns(uint256, bool)
func (_TribeChief *TribeChiefSession) PushSigner(signer common.Address, score *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.PushSigner(&_TribeChief.TransactOpts, signer, score)
}

// PushSigner is a paid mutator transaction binding the contract method 0xda371bd2.
//
// Solidity: function pushSigner(signer address, score int256) returns(uint256, bool)
func (_TribeChief *TribeChiefTransactorSession) PushSigner(signer common.Address, score *big.Int) (*types.Transaction, error) {
	return _TribeChief.Contract.PushSigner(&_TribeChief.TransactOpts, signer, score)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_TribeChief *TribeChiefTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_TribeChief *TribeChiefSession) Update() (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_TribeChief *TribeChiefTransactorSession) Update() (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts)
}
