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

// ChiefBase_1_0_0ABI is the input ABI used to generate the binding from.
const ChiefBase_1_0_0ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"volunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeLeaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocDelLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_tribe\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leaderLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pocCleanBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddLockList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pocGetBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vsn\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"pocChangeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"appendLeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"pocAddBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pocAddr\",\"type\":\"address\"},{\"name\":\"tribeAddr\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takeLeaderList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ChiefBase_1_0_0Bin is the compiled bytecode used for deploying new contracts.
const ChiefBase_1_0_0Bin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a090815261003e91600091906100b1565b5060058055600360065561181b600755600260085534801561005f57600080fd5b50604051604080610ff98339810180604052604081101561007f57600080fd5b50805160209091015160028054600160a060020a0319163317905560079190915560068190556000190160085561014c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f257805160ff191683800117855561011f565b8280016001018555821561011f579182015b8281111561011f578251825591602001919060010190610104565b5061012b92915061012f565b5090565b61014991905b8082111561012b5760008155600101610135565b90565b610e9e8061015b6000396000f3fe608060405234801561001057600080fd5b50600436106101ab576000357c0100000000000000000000000000000000000000000000000000000000900480639037f182116100fb578063bfe3a272116100b4578063db512e851161008e578063db512e8514610464578063ed77db3f1461049e578063f09a4016146104bb578063f0d55817146104e9576101ab565b8063bfe3a27214610410578063c8ba1fd814610436578063d0f261711461045c576101ab565b80639037f182146102bb578063a9cfa49c146102e1578063b175de1714610307578063b2bdfa7b1461035f578063bbf2ef6714610367578063bf8fb614146103e4576101ab565b80634bb19dac116101685780636c151173116101425780636c1511731461029b5780636da50830146102a35780637216276b146102ab578063900cf0cf146102b3576101ab565b80634bb19dac1461024957806357f607a11461026f578063609bf73b14610277576101ab565b806313af4035146101b05780631dd47d3d146101d857806326b249c814610217578063282f78f2146102315780634067f0c81461023957806340e2f99114610241575b600080fd5b6101d6600480360360208110156101c657600080fd5b5035600160a060020a03166104f1565b005b6101fe600480360360208110156101ee57600080fd5b5035600160a060020a0316610537565b60408051600092830b90920b8252519081900360200190f35b61021f6105ed565b60408051918252519081900360200190f35b61021f6105f3565b61021f6105fa565b61021f610600565b6101fe6004803603602081101561025f57600080fd5b5035600160a060020a0316610606565b61021f610688565b61027f61068e565b60408051600160a060020a039092168252519081900360200190f35b61021f61069d565b61021f6106a3565b6101d66106a9565b61021f610746565b6101d6600480360360208110156102d157600080fd5b5035600160a060020a031661074c565b6101d6600480360360208110156102f757600080fd5b5035600160a060020a031661085f565b61030f6108f8565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561034b578181015183820152602001610333565b505050509050019250505060405180910390f35b61027f6109e6565b61036f6109f5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103a9578181015183820152602001610391565b50505050905090810190601f1680156103d65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101d6600480360360408110156103fa57600080fd5b50600160a060020a038135169060200135610a83565b6101d66004803603602081101561042657600080fd5b5035600160a060020a0316610b24565b6101d66004803603602081101561044c57600080fd5b5035600160a060020a0316610bf2565b61027f610c70565b61048a6004803603602081101561047a57600080fd5b5035600160a060020a0316610c7f565b604080519115158252519081900360200190f35b61027f600480360360208110156104b457600080fd5b5035610ce9565b6101d6600480360360408110156104d157600080fd5b50600160a060020a0381358116916020013516610d11565b61030f610dc9565b600254600160a060020a0316331461050857600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600090600160a060020a0316331461055157600080fd5b600154604080517f2988fcfd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691632988fcfd9160248083019260209291908290030181600087803b1580156105b957600080fd5b505af11580156105cd573d6000803e3d6000fd5b505050506040513d60208110156105e357600080fd5b505190505b919050565b60085481565b6008545b90565b60075490565b60055490565b600354600090600160a060020a0316331461062057600080fd5b600154604080517f6fae4a5f000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691636fae4a5f9160248083019260209291908290030181600087803b1580156105b957600080fd5b60065481565b600354600160a060020a031681565b60055481565b60065490565b600354600160a060020a031633146106c057600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a0316633853d0a16040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b50505050565b60075481565b600254600160a060020a0316331461076357600080fd5b6004546001141561077357600080fd5b60005b60045481101561085b5781600160a060020a031660048281548110151561079957fe5b600091825260209091200154600160a060020a0316141561085357805b600454600019018110156108395760048054600183019081106107d557fe5b60009182526020909120015460048054600160a060020a0390921691839081106107fb57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556001016107b6565b50600480546000199283019201906108519082610e2b565b505b600101610776565b5050565b600354600160a060020a0316331461087657600080fd5b600154604080517f478e7d6e000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163478e7d6e91602480830192600092919082900301818387803b1580156108dd57600080fd5b505af11580156108f1573d6000803e3d6000fd5b5050505050565b600154604080517f360b97b90000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163360b97b9916004808301926000929190829003018186803b15801561095657600080fd5b505afa15801561096a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561099357600080fd5b8101908080516401000000008111156109ab57600080fd5b820160208101848111156109be57600080fd5b81518560208202830111640100000000821117156109db57600080fd5b509094505050505090565b600254600160a060020a031681565b6000805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a7b5780601f10610a5057610100808354040283529160200191610a7b565b820191906000526020600020905b815481529060010190602001808311610a5e57829003601f168201915b505050505081565b600254600160a060020a03163314610a9a57600080fd5b600154604080517f56df3db1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201859052915191909216916356df3db191604480830192600092919082900301818387803b158015610b0857600080fd5b505af1158015610b1c573d6000803e3d6000fd5b505050505050565b600254600160a060020a03163314610b3b57600080fd5b6005546004541015610bef5760005b600454811015610b955781600160a060020a0316600482815481101515610b6d57fe5b600091825260209091200154600160a060020a03161415610b8d57600080fd5b600101610b4a565b50600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600354600160a060020a03163314610c0957600080fd5b600154604080517f31f23668000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916331f2366891602480830192600092919082900301818387803b1580156108dd57600080fd5b600254600160a060020a031690565b6000805b600454811015610ce057600160a060020a03831615801590610cc957506004805482908110610cae57fe5b600091825260209091200154600160a060020a038481169116145b15610cd85760019150506105e8565b600101610c83565b50600092915050565b6004805482908110610cf757fe5b600091825260209091200154600160a060020a0316905081565b33600160a060020a03821614610d2657600080fd5b600154600160a060020a0316158015610d475750600160a060020a03821615155b15610d75576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b600160a060020a03811615801590610d965750600354600160a060020a0316155b1561085b5760038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790555050565b60606004805480602002602001604051908101604052809291908181526020018280548015610e2157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610e03575b5050505050905090565b815481835581811115610e4f57600083815260209020610e4f918101908301610e54565b505050565b6105f791905b80821115610e6e5760008155600101610e5a565b509056fea165627a7a72305820069b96ecffe27b4a58d1e55835a06f38e03f3f0355a8cd0381ce76517d01ca120029`

// DeployChiefBase_1_0_0 deploys a new Ethereum contract, binding an instance of ChiefBase_1_0_0 to it.
func DeployChiefBase_1_0_0(auth *bind.TransactOpts, backend bind.ContractBackend, _epoch *big.Int, _signerLimit *big.Int) (common.Address, *types.Transaction, *ChiefBase_1_0_0, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefBase_1_0_0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChiefBase_1_0_0Bin), backend, _epoch, _signerLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChiefBase_1_0_0{ChiefBase_1_0_0Caller: ChiefBase_1_0_0Caller{contract: contract}, ChiefBase_1_0_0Transactor: ChiefBase_1_0_0Transactor{contract: contract}}, nil
}

// ChiefBase_1_0_0 is an auto generated Go binding around an Ethereum contract.
type ChiefBase_1_0_0 struct {
	ChiefBase_1_0_0Caller     // Read-only binding to the contract
	ChiefBase_1_0_0Transactor // Write-only binding to the contract
}

// ChiefBase_1_0_0Caller is an auto generated read-only Go binding around an Ethereum contract.
type ChiefBase_1_0_0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBase_1_0_0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ChiefBase_1_0_0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChiefBase_1_0_0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChiefBase_1_0_0Session struct {
	Contract     *ChiefBase_1_0_0        // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChiefBase_1_0_0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChiefBase_1_0_0CallerSession struct {
	Contract *ChiefBase_1_0_0Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// ChiefBase_1_0_0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChiefBase_1_0_0TransactorSession struct {
	Contract     *ChiefBase_1_0_0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChiefBase_1_0_0Raw is an auto generated low-level Go binding around an Ethereum contract.
type ChiefBase_1_0_0Raw struct {
	Contract *ChiefBase_1_0_0 // Generic contract binding to access the raw methods on
}

// ChiefBase_1_0_0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChiefBase_1_0_0CallerRaw struct {
	Contract *ChiefBase_1_0_0Caller // Generic read-only contract binding to access the raw methods on
}

// ChiefBase_1_0_0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChiefBase_1_0_0TransactorRaw struct {
	Contract *ChiefBase_1_0_0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewChiefBase_1_0_0 creates a new instance of ChiefBase_1_0_0, bound to a specific deployed contract.
func NewChiefBase_1_0_0(address common.Address, backend bind.ContractBackend) (*ChiefBase_1_0_0, error) {
	contract, err := bindChiefBase_1_0_0(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChiefBase_1_0_0{ChiefBase_1_0_0Caller: ChiefBase_1_0_0Caller{contract: contract}, ChiefBase_1_0_0Transactor: ChiefBase_1_0_0Transactor{contract: contract}}, nil
}

// NewChiefBase_1_0_0Caller creates a new read-only instance of ChiefBase_1_0_0, bound to a specific deployed contract.
func NewChiefBase_1_0_0Caller(address common.Address, caller bind.ContractCaller) (*ChiefBase_1_0_0Caller, error) {
	contract, err := bindChiefBase_1_0_0(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ChiefBase_1_0_0Caller{contract: contract}, nil
}

// NewChiefBase_1_0_0Transactor creates a new write-only instance of ChiefBase_1_0_0, bound to a specific deployed contract.
func NewChiefBase_1_0_0Transactor(address common.Address, transactor bind.ContractTransactor) (*ChiefBase_1_0_0Transactor, error) {
	contract, err := bindChiefBase_1_0_0(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ChiefBase_1_0_0Transactor{contract: contract}, nil
}

// bindChiefBase_1_0_0 binds a generic wrapper to an already deployed contract.
func bindChiefBase_1_0_0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChiefBase_1_0_0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase_1_0_0.Contract.ChiefBase_1_0_0Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.ChiefBase_1_0_0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.ChiefBase_1_0_0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ChiefBase_1_0_0.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "_owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) Owner() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.Owner(&_ChiefBase_1_0_0.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) Owner() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.Owner(&_ChiefBase_1_0_0.CallOpts)
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) Tribe(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "_tribe")
	return *ret0, err
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) Tribe() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.Tribe(&_ChiefBase_1_0_0.CallOpts)
}

// Tribe is a free data retrieval call binding the contract method 0x609bf73b.
//
// Solidity: function _tribe() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) Tribe() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.Tribe(&_ChiefBase_1_0_0.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) Epoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "epoch")
	return *ret0, err
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) Epoch() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.Epoch(&_ChiefBase_1_0_0.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) Epoch() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.Epoch(&_ChiefBase_1_0_0.CallOpts)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) IsLeader(opts *bind.CallOptsWithNumber, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "isLeader", addr)
	return *ret0, err
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) IsLeader(addr common.Address) (bool, error) {
	return _ChiefBase_1_0_0.Contract.IsLeader(&_ChiefBase_1_0_0.CallOpts, addr)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(addr address) constant returns(bool)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) IsLeader(addr common.Address) (bool, error) {
	return _ChiefBase_1_0_0.Contract.IsLeader(&_ChiefBase_1_0_0.CallOpts, addr)
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) LeaderLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "leaderLimit")
	return *ret0, err
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) LeaderLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.LeaderLimit(&_ChiefBase_1_0_0.CallOpts)
}

// LeaderLimit is a free data retrieval call binding the contract method 0x6c151173.
//
// Solidity: function leaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) LeaderLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.LeaderLimit(&_ChiefBase_1_0_0.CallOpts)
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) LeaderList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "leaderList", arg0)
	return *ret0, err
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) LeaderList(arg0 *big.Int) (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.LeaderList(&_ChiefBase_1_0_0.CallOpts, arg0)
}

// LeaderList is a free data retrieval call binding the contract method 0xed77db3f.
//
// Solidity: function leaderList( uint256) constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) LeaderList(arg0 *big.Int) (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.LeaderList(&_ChiefBase_1_0_0.CallOpts, arg0)
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) PocGetBlackList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "pocGetBlackList")
	return *ret0, err
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocGetBlackList() ([]common.Address, error) {
	return _ChiefBase_1_0_0.Contract.PocGetBlackList(&_ChiefBase_1_0_0.CallOpts)
}

// PocGetBlackList is a free data retrieval call binding the contract method 0xb175de17.
//
// Solidity: function pocGetBlackList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) PocGetBlackList() ([]common.Address, error) {
	return _ChiefBase_1_0_0.Contract.PocGetBlackList(&_ChiefBase_1_0_0.CallOpts)
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) SignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "signerLimit")
	return *ret0, err
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) SignerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.SignerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// SignerLimit is a free data retrieval call binding the contract method 0x57f607a1.
//
// Solidity: function signerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) SignerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.SignerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeEpoch")
	return *ret0, err
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeEpoch() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeEpoch(&_ChiefBase_1_0_0.CallOpts)
}

// TakeEpoch is a free data retrieval call binding the contract method 0x4067f0c8.
//
// Solidity: function takeEpoch() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeEpoch() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeEpoch(&_ChiefBase_1_0_0.CallOpts)
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeLeaderLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeLeaderLimit")
	return *ret0, err
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeLeaderLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeLeaderLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeLeaderLimit is a free data retrieval call binding the contract method 0x40e2f991.
//
// Solidity: function takeLeaderLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeLeaderLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeLeaderLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeLeaderList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeLeaderList")
	return *ret0, err
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeLeaderList() ([]common.Address, error) {
	return _ChiefBase_1_0_0.Contract.TakeLeaderList(&_ChiefBase_1_0_0.CallOpts)
}

// TakeLeaderList is a free data retrieval call binding the contract method 0xf0d55817.
//
// Solidity: function takeLeaderList() constant returns(address[])
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeLeaderList() ([]common.Address, error) {
	return _ChiefBase_1_0_0.Contract.TakeLeaderList(&_ChiefBase_1_0_0.CallOpts)
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeOwner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeOwner")
	return *ret0, err
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeOwner() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.TakeOwner(&_ChiefBase_1_0_0.CallOpts)
}

// TakeOwner is a free data retrieval call binding the contract method 0xd0f26171.
//
// Solidity: function takeOwner() constant returns(address)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeOwner() (common.Address, error) {
	return _ChiefBase_1_0_0.Contract.TakeOwner(&_ChiefBase_1_0_0.CallOpts)
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeSignerLimit")
	return *ret0, err
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeSignerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeSignerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeSignerLimit is a free data retrieval call binding the contract method 0x6da50830.
//
// Solidity: function takeSignerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeSignerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeSignerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) TakeVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "takeVolunteerLimit")
	return *ret0, err
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) TakeVolunteerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeVolunteerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// TakeVolunteerLimit is a free data retrieval call binding the contract method 0x282f78f2.
//
// Solidity: function takeVolunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) TakeVolunteerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.TakeVolunteerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) VolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "volunteerLimit")
	return *ret0, err
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) VolunteerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.VolunteerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// VolunteerLimit is a free data retrieval call binding the contract method 0x26b249c8.
//
// Solidity: function volunteerLimit() constant returns(uint256)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) VolunteerLimit() (*big.Int, error) {
	return _ChiefBase_1_0_0.Contract.VolunteerLimit(&_ChiefBase_1_0_0.CallOpts)
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Caller) Vsn(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChiefBase_1_0_0.contract.CallWithNumber(opts, out, "vsn")
	return *ret0, err
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) Vsn() (string, error) {
	return _ChiefBase_1_0_0.Contract.Vsn(&_ChiefBase_1_0_0.CallOpts)
}

// Vsn is a free data retrieval call binding the contract method 0xbbf2ef67.
//
// Solidity: function vsn() constant returns(string)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0CallerSession) Vsn() (string, error) {
	return _ChiefBase_1_0_0.Contract.Vsn(&_ChiefBase_1_0_0.CallOpts)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) AppendLeader(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "appendLeader", addr)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) AppendLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.AppendLeader(&_ChiefBase_1_0_0.TransactOpts, addr)
}

// AppendLeader is a paid mutator transaction binding the contract method 0xbfe3a272.
//
// Solidity: function appendLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) AppendLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.AppendLeader(&_ChiefBase_1_0_0.TransactOpts, addr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) Init(opts *bind.TransactOpts, pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "init", pocAddr, tribeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) Init(pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.Init(&_ChiefBase_1_0_0.TransactOpts, pocAddr, tribeAddr)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(pocAddr address, tribeAddr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) Init(pocAddr common.Address, tribeAddr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.Init(&_ChiefBase_1_0_0.TransactOpts, pocAddr, tribeAddr)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocAddBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocAddBlackList", minerAddress)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocAddBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddBlackList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocAddBlackList is a paid mutator transaction binding the contract method 0xc8ba1fd8.
//
// Solidity: function pocAddBlackList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocAddBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddBlackList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocAddLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocAddLockList", minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocAddLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddLockList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocAddLockList is a paid mutator transaction binding the contract method 0xa9cfa49c.
//
// Solidity: function pocAddLockList(minerAddress address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocAddLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddLockList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocAddStop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocAddStop", minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocAddStop(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddStop(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocAddStop is a paid mutator transaction binding the contract method 0x1dd47d3d.
//
// Solidity: function pocAddStop(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocAddStop(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocAddStop(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocChangeOwner(opts *bind.TransactOpts, newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocChangeOwner", newOwner, num)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocChangeOwner(newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocChangeOwner(&_ChiefBase_1_0_0.TransactOpts, newOwner, num)
}

// PocChangeOwner is a paid mutator transaction binding the contract method 0xbf8fb614.
//
// Solidity: function pocChangeOwner(newOwner address, num uint256) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocChangeOwner(newOwner common.Address, num *big.Int) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocChangeOwner(&_ChiefBase_1_0_0.TransactOpts, newOwner, num)
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocCleanBlackList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocCleanBlackList")
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocCleanBlackList() (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocCleanBlackList(&_ChiefBase_1_0_0.TransactOpts)
}

// PocCleanBlackList is a paid mutator transaction binding the contract method 0x7216276b.
//
// Solidity: function pocCleanBlackList() returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocCleanBlackList() (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocCleanBlackList(&_ChiefBase_1_0_0.TransactOpts)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) PocDelLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "pocDelLockList", minerAddress)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) PocDelLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocDelLockList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// PocDelLockList is a paid mutator transaction binding the contract method 0x4bb19dac.
//
// Solidity: function pocDelLockList(minerAddress address) returns(int8)
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) PocDelLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.PocDelLockList(&_ChiefBase_1_0_0.TransactOpts, minerAddress)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) RemoveLeader(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "removeLeader", addr)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) RemoveLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.RemoveLeader(&_ChiefBase_1_0_0.TransactOpts, addr)
}

// RemoveLeader is a paid mutator transaction binding the contract method 0x9037f182.
//
// Solidity: function removeLeader(addr address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) RemoveLeader(addr common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.RemoveLeader(&_ChiefBase_1_0_0.TransactOpts, addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Transactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0Session) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.SetOwner(&_ChiefBase_1_0_0.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_ChiefBase_1_0_0 *ChiefBase_1_0_0TransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChiefBase_1_0_0.Contract.SetOwner(&_ChiefBase_1_0_0.TransactOpts, newOwner)
}

// POC_1_0_0ABI is the input ABI used to generate the binding from.
const POC_1_0_0ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"minerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPopBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawWaitNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerStop\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPushBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerEmptyBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initMinDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPushLockList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwnerEffectiveNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNormalList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"},{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"withdrawSurplus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"ownerPopLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStopList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"blackAndLockStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLockList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"ownerSetInitBlockNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositHalveLimitNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_waitNumber\",\"type\":\"uint256\"},{\"name\":\"_limitNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"Start\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// POC_1_0_0Bin is the compiled bytecode used for deploying new contracts.
const POC_1_0_0Bin = `0x608060405234801561001057600080fd5b506040516080806123628339810180604052608081101561003057600080fd5b50805160208201516040830151606090930151600a91909155600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a03909316600160a060020a031990931692909217909155600d91909155600c5543600b556122ab806100b76000396000f3fe6080604052600436106101b9576000357c01000000000000000000000000000000000000000000000000000000009004806353ed5143116101095780639c52ade2116100a7578063ed1eeaf111610081578063ed1eeaf11461071b578063f2a558f914610730578063f30011b41461075c578063f8864cfe14610786576101b9565b80639c52ade2146106a0578063dd0b281e146106b5578063e7b14a93146106e8576101b9565b8063645006ca116100e3578063645006ca146105f2578063656cb0bc1461060757806365e73e321461063a5780636fae4a5f1461066d576101b9565b806353ed51431461042757806356df3db1146105a457806358100370146105dd576101b9565b8063360b97b911610176578063478e7d6e11610150578063478e7d6e146103975780634b58f36e146103ca57806351393ce3146103df57806351cff8d9146103f4576101b9565b8063360b97b9146103085780633853d0a11461036d5780633913468e14610382576101b9565b80630c58ed8e146101be5780631103ec9614610218578063147049701461026457806321615f6e1461028b5780632988fcfd146102a057806331f23668146102d3575b600080fd5b3480156101ca57600080fd5b506101f1600480360360208110156101e157600080fd5b5035600160a060020a031661079b565b604080519384526020840192909252600160a060020a031682820152519081900360600190f35b34801561022457600080fd5b5061024b6004803603602081101561023b57600080fd5b5035600160a060020a03166107c8565b60408051600092830b90920b8252519081900360200190f35b34801561027057600080fd5b50610279610990565b60408051918252519081900360200190f35b34801561029757600080fd5b50610279610996565b3480156102ac57600080fd5b5061024b600480360360208110156102c357600080fd5b5035600160a060020a031661099c565b3480156102df57600080fd5b50610306600480360360208110156102f657600080fd5b5035600160a060020a0316610a30565b005b34801561031457600080fd5b5061031d610b6f565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610359578181015183820152602001610341565b505050509050019250505060405180910390f35b34801561037957600080fd5b50610306610bd2565b34801561038e57600080fd5b50610279610caf565b3480156103a357600080fd5b50610306600480360360208110156103ba57600080fd5b5035600160a060020a0316610cb5565b3480156103d657600080fd5b50610279610df2565b3480156103eb57600080fd5b5061031d610df8565b34801561040057600080fd5b506103066004803603602081101561041757600080fd5b5035600160a060020a0316610e58565b34801561043357600080fd5b5061043c61106c565b60405180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561048c578181015183820152602001610474565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b838110156104cb5781810151838201526020016104b3565b50505050905001868103845289818151815260200191508051906020019060200280838360005b8381101561050a5781810151838201526020016104f2565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015610549578181015183820152602001610531565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015610588578181015183820152602001610570565b505050509050019a505050505050505050505060405180910390f35b3480156105b057600080fd5b50610306600480360360408110156105c757600080fd5b50600160a060020a0381351690602001356115f9565b3480156105e957600080fd5b5061031d611721565b3480156105fe57600080fd5b50610279611781565b34801561061357600080fd5b506103066004803603602081101561062a57600080fd5b5035600160a060020a03166117bc565b34801561064657600080fd5b506103066004803603602081101561065d57600080fd5b5035600160a060020a03166117ff565b34801561067957600080fd5b5061024b6004803603602081101561069057600080fd5b5035600160a060020a03166118e0565b3480156106ac57600080fd5b5061031d611a75565b3480156106c157600080fd5b50610306600480360360208110156106d857600080fd5b5035600160a060020a0316611ad5565b3480156106f457600080fd5b5061024b6004803603602081101561070b57600080fd5b5035600160a060020a0316611b04565b34801561072757600080fd5b5061031d611bba565b6103066004803603606081101561074657600080fd5b508035906020810135906040013560ff16611c1a565b34801561076857600080fd5b506103066004803603602081101561077f57600080fd5b5035611ddc565b34801561079257600080fd5b50610279611e65565b600160a060020a039081166000908152602081905260409020600181015460028201549154909391921690565b60098054600091829060001983018381106107df57fe5b600091825260209091200154600160a060020a031690506001821180156108075750600e5443105b156108365760098054600119840190811061081e57fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461084b57600080fd5b600160a060020a038416600090815260066020526040902054600554600111806108a1575084600160a060020a031660058281548110151561088957fe5b600091825260209091200154600160a060020a031614155b156108b157600019935050610989565b6005805460001981019081106108c357fe5b60009182526020909120015460058054600160a060020a0390921691839081106108e957fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556005805460001901906109259082612242565b50600160a060020a03851660009081526006602052604081205560055481101561098357806006600060058481548110151561095d57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b60019350505b5050919050565b600d5481565b600b5481565b60098054600091829060001983018381106109b357fe5b600091825260209091200154600160a060020a031690506001821180156109db5750600e5443105b15610a0a576009805460011984019081106109f257fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610a1f57600080fd5b610a2884611e6b565b949350505050565b60098054906000906000198301838110610a4657fe5b600091825260209091200154600160a060020a03169050600182118015610a6e5750600e5443105b15610a9d57600980546001198401908110610a8557fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610ab257600080fd5b6005546000108015610afe5750600160a060020a038316600081815260066020526040902054600580549091908110610ae757fe5b600091825260209091200154600160a060020a0316145b15610b0857610b6a565b600580546001810182557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018054600160a060020a031916600160a060020a038616908117909155905460009182526006602052604090912060001990910190555b505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610bc757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610ba9575b505050505090505b90565b60098054906000906000198301838110610be857fe5b600091825260209091200154600160a060020a03169050600182118015610c105750600e5443105b15610c3f57600980546001198401908110610c2757fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610c5457600080fd5b60005b600554811015610ca15760066000600583815481101515610c7457fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600101610c57565b506000610b6a600582612242565b600a5481565b60098054906000906000198301838110610ccb57fe5b600091825260209091200154600160a060020a03169050600182118015610cf35750600e5443105b15610d2257600980546001198401908110610d0a57fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614610d3757600080fd5b6007546000108015610d835750600160a060020a038316600081815260086020526040902054600780549091908110610d6c57fe5b600091825260209091200154600160a060020a0316145b15610d8d57610b6a565b5050600780546001810182557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018054600160a060020a031916600160a060020a03939093169283179055546000918252600860205260409091206000199091019055565b600e5481565b60606001805480602002602001604051908101604052809291908181526020018280548015610bc757602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610ba9575050505050905090565b600160a060020a03818116600090815260208190526040902054163314610e7e57600080fd5b600160a060020a03811660009081526020819052604081206002015411610ea457600080fd5b600d54600160a060020a038216600090815260208190526040902060020154430311610ecf57600080fd5b610ed881611b04565b60000b15610ee557600080fd5b600160a060020a0381166000908152602081815260408083206001810180548254600160a060020a0319168355908590556002909101849055600490925290912054600380546000198101908110610f3957fe5b60009182526020909120015460038054600160a060020a039092169183908110610f5f57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600380546000190190610f9b9082612242565b50600160a060020a038316600090815260046020526040812055600354811015610ff9578060046000600384815481101515610fd357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b6040805183815290513391600160a060020a038616917f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9181900360200190a3604051339083156108fc029084906000818181858888f19350505050158015611066573d6000803e3d6000fd5b50505050565b600354600154604080518284018082526020808202830101909252606093849384938493849392919084908280156110ae578160200160208202803883390190505b5090506060826040519080825280602002602001820160405280156110dd578160200160208202803883390190505b50905060608360405190808252806020026020018201604052801561110c578160200160208202803883390190505b50905060608460405190808252806020026020018201604052801561113b578160200160208202803883390190505b50905060608560405190808252806020026020018201604052801561116a578160200160208202803883390190505b50905060005b888110156113a257600380548290811061118657fe5b6000918252602090912001548651600160a060020a03909116908790839081106111ac57fe5b600160a060020a03909216602092830290910190910152600380546000918291849081106111d657fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154855186908390811061120c57fe5b602090810290910101526003805460009182918490811061122957fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060020154845185908390811061125f57fe5b602090810290910101526003805460009182918490811061127c57fe5b6000918252602080832090910154600160a060020a03908116845290830193909352604090910190205484519116908490839081106112b757fe5b600160a060020a03909216602092830290910190910152600554600010801561135b575060038054829081106112e957fe5b600091825260208220015460038054600160a060020a0390921692600592600692908690811061131557fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548154811061134457fe5b600091825260209091200154600160a060020a0316145b1561137f576001828281518110151561137057fe5b6020908102909101015261139a565b6000828281518110151561138f57fe5b602090810290910101525b600101611170565b5060005b878110156115e45760018054829081106113bc57fe5b6000918252602090912001548651600160a060020a039091169087908b84019081106113e457fe5b600160a060020a039092166020928302909101909101526001805460009182918490811061140e57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154855186908b840190811061144657fe5b602090810290910101526001805460009182918490811061146357fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060020154845185908b840190811061149b57fe5b60209081029091010152600180546000918291849081106114b857fe5b6000918252602080832090910154600160a060020a039081168452908301939093526040909101902054845191169084908b84019081106114f557fe5b600160a060020a0390921660209283029091019091015260055460001080156115995750600180548290811061152757fe5b600091825260208220015460018054600160a060020a0390921692600592600692908690811061155357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548154811061158257fe5b600091825260209091200154600160a060020a0316145b156115bf57600182828b018151811015156115b057fe5b602090810290910101526115dc565b600082828b018151811015156115d157fe5b602090810290910101525b6001016113a6565b50939c929b5090995097509095509350505050565b6009805490600090600019830183811061160f57fe5b600091825260209091200154600160a060020a031690506001821180156116375750600e5443105b156116665760098054600119840190811061164e57fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461167b57600080fd5b600e544310156116cd5760098054859190600019810190811061169a57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611719565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a031916600160a060020a0386161790555b5050600e5550565b60606009805480602002602001604051908101604052809291908181526020018280548015610bc757602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610ba9575050505050905090565b600080600c54600b54430381151561179557fe5b04905060038111156117a5575060035b8060020a600a548115156117b557fe5b0491505090565b600160a060020a038181166000908152602081905260409020541633146117e257600080fd5b6117eb81611e6b565b60000b60011415156117fc57600080fd5b50565b600160a060020a0381811660009081526020819052604090205416331461182557600080fd5b600061182f611781565b600160a060020a038316600090815260208190526040812060010154919091039150811161185c57600080fd5b600160a060020a0382166000818152602081815260409182902060010180548590039055815184815291513393927f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb92908290030190a3604051339082156108fc029083906000818181858888f19350505050158015610b6a573d6000803e3d6000fd5b60098054600091829060001983018381106118f757fe5b600091825260209091200154600160a060020a0316905060018211801561191f5750600e5443105b1561194e5760098054600119840190811061193657fe5b600091825260209091200154600160a060020a031690505b33600160a060020a0382161461196357600080fd5b600160a060020a038416600090815260086020526040902054600754600111806119b9575084600160a060020a03166007828154811015156119a157fe5b600091825260209091200154600160a060020a031614155b156119c957600019935050610989565b6007805460001981019081106119db57fe5b60009182526020909120015460078054600160a060020a039092169183908110611a0157fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600780546000190190611a3d9082612242565b50600160a060020a03851660009081526008602052604081205560075481101561098357806008600060078481548110151561095d57fe5b60606003805480602002602001604051908101604052809291908181526020018280548015610bc757602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610ba9575050505050905090565b600160a060020a03818116600090815260208190526040902054163314611afb57600080fd5b6117eb81612048565b600554600090819081108015611b545750600160a060020a038316600081815260066020526040902054600580549091908110611b3d57fe5b600091825260209091200154600160a060020a0316145b15611b5d576001015b6007546000108015611ba95750600160a060020a038316600081815260086020526040902054600780549091908110611b9257fe5b600091825260209091200154600160a060020a0316145b15611bb2576002015b90505b919050565b60606007805480602002602001604051908101604052809291908181526020018280548015610bc757602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610ba9575050505050905090565b611c22611781565b341015611c2e57600080fd5b604080516c01000000000000000000000000330260208083019190915282518083036014018152603483018085528151918301919091206000918290526054840180865281905260ff861660748501526094840188905260b484018790529351909260019260d4808301939192601f198301929081900390910190855afa158015611cbd573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a0381161515611cdf57600080fd5b600160a060020a03811660009081526020819052604090206001015415611d0557600080fd5b60408051606081018252338082523460208084018281526000858701818152600160a060020a03898116808452838652898420985189549216600160a060020a031992831617895593516001808a01919091559151600298890155815480830183557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180549091168417905554958352869020600019959095019094558451918252935191937f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6292918290030190a35050505050565b60098054906000906000198301838110611df257fe5b600091825260209091200154600160a060020a03169050600182118015611e1a5750600e5443105b15611e4957600980546001198401908110611e3157fe5b600091825260209091200154600160a060020a031690505b33600160a060020a03821614611e5e57600080fd5b5050600b55565b600c5481565b600160a060020a0381166000908152602081905260408120600101541515611e965750600019611bb5565b600160a060020a03821660009081526020819052604090206002015415611ebf57506000611bb5565b600160a060020a03821660009081526020818152604080832043600291820155909152902054600180546000198101908110611ef757fe5b60009182526020909120015460018054600160a060020a039092169183908110611f1d57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600180546000190190611f599082612242565b50600160a060020a038316600090815260026020526040812055600154811015611fb7578060026000600184815481101515611f9157fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b600380546001810182557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a031916600160a060020a038616908117909155905460008281526004602052604080822060001990930190925590517f5f8b016b80ca3a499d7fc818528ee3c8658d421489a9efdf1d365c4241ddd8839190a250600192915050565b600160a060020a03811660009081526020819052604081206001015415156120735750600019611bb5565b600061207e83611b04565b60000b131561208f57506000611bb5565b600160a060020a03821660009081526020819052604090206002015415156120b957506000611bb5565b600160a060020a03821660009081526020818152604080832060020183905560049091529020546003805460001981019081106120f257fe5b60009182526020909120015460038054600160a060020a03909216918390811061211857fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556003805460001901906121549082612242565b50600160a060020a0383166000908152600460205260408120556003548110156121b257806004600060038481548110151561218c57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020555b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a031916600160a060020a038616908117909155905460008281526002602052604080822060001990930190925590517f84d0447ca38875fa61115673259a210915bc1dd53a3c112d6f0790f15956a9659190a250600192915050565b815481835581811115610b6a57600083815260209020610b6a918101908301610bcf91905b8082111561227b5760008155600101612267565b509056fea165627a7a72305820221ad753e122d573f86c27dc70c55f5f486a0af73db10f59d506448dee597ef30029`

// DeployPOC_1_0_0 deploys a new Ethereum contract, binding an instance of POC_1_0_0 to it.
func DeployPOC_1_0_0(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address, _amount *big.Int, _waitNumber *big.Int, _limitNumber *big.Int) (common.Address, *types.Transaction, *POC_1_0_0, error) {
	parsed, err := abi.JSON(strings.NewReader(POC_1_0_0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(POC_1_0_0Bin), backend, _addr, _amount, _waitNumber, _limitNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &POC_1_0_0{POC_1_0_0Caller: POC_1_0_0Caller{contract: contract}, POC_1_0_0Transactor: POC_1_0_0Transactor{contract: contract}}, nil
}

// POC_1_0_0 is an auto generated Go binding around an Ethereum contract.
type POC_1_0_0 struct {
	POC_1_0_0Caller     // Read-only binding to the contract
	POC_1_0_0Transactor // Write-only binding to the contract
}

// POC_1_0_0Caller is an auto generated read-only Go binding around an Ethereum contract.
type POC_1_0_0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POC_1_0_0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type POC_1_0_0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POC_1_0_0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type POC_1_0_0Session struct {
	Contract     *POC_1_0_0              // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// POC_1_0_0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type POC_1_0_0CallerSession struct {
	Contract *POC_1_0_0Caller        // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// POC_1_0_0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type POC_1_0_0TransactorSession struct {
	Contract     *POC_1_0_0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// POC_1_0_0Raw is an auto generated low-level Go binding around an Ethereum contract.
type POC_1_0_0Raw struct {
	Contract *POC_1_0_0 // Generic contract binding to access the raw methods on
}

// POC_1_0_0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type POC_1_0_0CallerRaw struct {
	Contract *POC_1_0_0Caller // Generic read-only contract binding to access the raw methods on
}

// POC_1_0_0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type POC_1_0_0TransactorRaw struct {
	Contract *POC_1_0_0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPOC_1_0_0 creates a new instance of POC_1_0_0, bound to a specific deployed contract.
func NewPOC_1_0_0(address common.Address, backend bind.ContractBackend) (*POC_1_0_0, error) {
	contract, err := bindPOC_1_0_0(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &POC_1_0_0{POC_1_0_0Caller: POC_1_0_0Caller{contract: contract}, POC_1_0_0Transactor: POC_1_0_0Transactor{contract: contract}}, nil
}

// NewPOC_1_0_0Caller creates a new read-only instance of POC_1_0_0, bound to a specific deployed contract.
func NewPOC_1_0_0Caller(address common.Address, caller bind.ContractCaller) (*POC_1_0_0Caller, error) {
	contract, err := bindPOC_1_0_0(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &POC_1_0_0Caller{contract: contract}, nil
}

// NewPOC_1_0_0Transactor creates a new write-only instance of POC_1_0_0, bound to a specific deployed contract.
func NewPOC_1_0_0Transactor(address common.Address, transactor bind.ContractTransactor) (*POC_1_0_0Transactor, error) {
	contract, err := bindPOC_1_0_0(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &POC_1_0_0Transactor{contract: contract}, nil
}

// bindPOC_1_0_0 binds a generic wrapper to an already deployed contract.
func bindPOC_1_0_0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(POC_1_0_0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POC_1_0_0 *POC_1_0_0Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _POC_1_0_0.Contract.POC_1_0_0Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POC_1_0_0 *POC_1_0_0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.POC_1_0_0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POC_1_0_0 *POC_1_0_0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.POC_1_0_0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POC_1_0_0 *POC_1_0_0CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _POC_1_0_0.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POC_1_0_0 *POC_1_0_0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POC_1_0_0 *POC_1_0_0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.contract.Transact(opts, method, params...)
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC_1_0_0 *POC_1_0_0Caller) BlackAndLockStatus(opts *bind.CallOptsWithNumber, minerAddress common.Address) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "blackAndLockStatus", minerAddress)
	return *ret0, err
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC_1_0_0 *POC_1_0_0Session) BlackAndLockStatus(minerAddress common.Address) (int8, error) {
	return _POC_1_0_0.Contract.BlackAndLockStatus(&_POC_1_0_0.CallOpts, minerAddress)
}

// BlackAndLockStatus is a free data retrieval call binding the contract method 0xe7b14a93.
//
// Solidity: function blackAndLockStatus(minerAddress address) constant returns(int8)
func (_POC_1_0_0 *POC_1_0_0CallerSession) BlackAndLockStatus(minerAddress common.Address) (int8, error) {
	return _POC_1_0_0.Contract.BlackAndLockStatus(&_POC_1_0_0.CallOpts, minerAddress)
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) DepositHalveLimitNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "depositHalveLimitNumber")
	return *ret0, err
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) DepositHalveLimitNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.DepositHalveLimitNumber(&_POC_1_0_0.CallOpts)
}

// DepositHalveLimitNumber is a free data retrieval call binding the contract method 0xf8864cfe.
//
// Solidity: function depositHalveLimitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) DepositHalveLimitNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.DepositHalveLimitNumber(&_POC_1_0_0.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[], uint256[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetAll(opts *bind.CallOptsWithNumber) ([]common.Address, []*big.Int, []*big.Int, []common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
		ret2 = new([]*big.Int)
		ret3 = new([]common.Address)
		ret4 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getAll")
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[], uint256[])
func (_POC_1_0_0 *POC_1_0_0Session) GetAll() ([]common.Address, []*big.Int, []*big.Int, []common.Address, []*big.Int, error) {
	return _POC_1_0_0.Contract.GetAll(&_POC_1_0_0.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() constant returns(address[], uint256[], uint256[], address[], uint256[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetAll() ([]common.Address, []*big.Int, []*big.Int, []common.Address, []*big.Int, error) {
	return _POC_1_0_0.Contract.GetAll(&_POC_1_0_0.CallOpts)
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetBlackList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getBlackList")
	return *ret0, err
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Session) GetBlackList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetBlackList(&_POC_1_0_0.CallOpts)
}

// GetBlackList is a free data retrieval call binding the contract method 0x360b97b9.
//
// Solidity: function getBlackList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetBlackList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetBlackList(&_POC_1_0_0.CallOpts)
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetLockList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getLockList")
	return *ret0, err
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Session) GetLockList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetLockList(&_POC_1_0_0.CallOpts)
}

// GetLockList is a free data retrieval call binding the contract method 0xed1eeaf1.
//
// Solidity: function getLockList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetLockList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetLockList(&_POC_1_0_0.CallOpts)
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetNormalList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getNormalList")
	return *ret0, err
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Session) GetNormalList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetNormalList(&_POC_1_0_0.CallOpts)
}

// GetNormalList is a free data retrieval call binding the contract method 0x51393ce3.
//
// Solidity: function getNormalList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetNormalList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetNormalList(&_POC_1_0_0.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetOwnerList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getOwnerList")
	return *ret0, err
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Session) GetOwnerList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetOwnerList(&_POC_1_0_0.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x58100370.
//
// Solidity: function getOwnerList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetOwnerList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetOwnerList(&_POC_1_0_0.CallOpts)
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Caller) GetStopList(opts *bind.CallOptsWithNumber) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "getStopList")
	return *ret0, err
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0Session) GetStopList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetStopList(&_POC_1_0_0.CallOpts)
}

// GetStopList is a free data retrieval call binding the contract method 0x9c52ade2.
//
// Solidity: function getStopList() constant returns(address[])
func (_POC_1_0_0 *POC_1_0_0CallerSession) GetStopList() ([]common.Address, error) {
	return _POC_1_0_0.Contract.GetStopList(&_POC_1_0_0.CallOpts)
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) InitBlockNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "initBlockNumber")
	return *ret0, err
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) InitBlockNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.InitBlockNumber(&_POC_1_0_0.CallOpts)
}

// InitBlockNumber is a free data retrieval call binding the contract method 0x21615f6e.
//
// Solidity: function initBlockNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) InitBlockNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.InitBlockNumber(&_POC_1_0_0.CallOpts)
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) InitMinDeposit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "initMinDeposit")
	return *ret0, err
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) InitMinDeposit() (*big.Int, error) {
	return _POC_1_0_0.Contract.InitMinDeposit(&_POC_1_0_0.CallOpts)
}

// InitMinDeposit is a free data retrieval call binding the contract method 0x3913468e.
//
// Solidity: function initMinDeposit() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) InitMinDeposit() (*big.Int, error) {
	return _POC_1_0_0.Contract.InitMinDeposit(&_POC_1_0_0.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) MinDepositAmount(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "minDepositAmount")
	return *ret0, err
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) MinDepositAmount() (*big.Int, error) {
	return _POC_1_0_0.Contract.MinDepositAmount(&_POC_1_0_0.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) MinDepositAmount() (*big.Int, error) {
	return _POC_1_0_0.Contract.MinDepositAmount(&_POC_1_0_0.CallOpts)
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC_1_0_0 *POC_1_0_0Caller) MinerStatus(opts *bind.CallOptsWithNumber, _addr common.Address) (*big.Int, *big.Int, common.Address, error) {
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
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "minerStatus", _addr)
	return *ret0, *ret1, *ret2, err
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC_1_0_0 *POC_1_0_0Session) MinerStatus(_addr common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _POC_1_0_0.Contract.MinerStatus(&_POC_1_0_0.CallOpts, _addr)
}

// MinerStatus is a free data retrieval call binding the contract method 0x0c58ed8e.
//
// Solidity: function minerStatus(_addr address) constant returns(uint256, uint256, address)
func (_POC_1_0_0 *POC_1_0_0CallerSession) MinerStatus(_addr common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _POC_1_0_0.Contract.MinerStatus(&_POC_1_0_0.CallOpts, _addr)
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) NewOwnerEffectiveNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "newOwnerEffectiveNumber")
	return *ret0, err
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) NewOwnerEffectiveNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.NewOwnerEffectiveNumber(&_POC_1_0_0.CallOpts)
}

// NewOwnerEffectiveNumber is a free data retrieval call binding the contract method 0x4b58f36e.
//
// Solidity: function newOwnerEffectiveNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) NewOwnerEffectiveNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.NewOwnerEffectiveNumber(&_POC_1_0_0.CallOpts)
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Caller) WithdrawWaitNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _POC_1_0_0.contract.CallWithNumber(opts, out, "withdrawWaitNumber")
	return *ret0, err
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0Session) WithdrawWaitNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.WithdrawWaitNumber(&_POC_1_0_0.CallOpts)
}

// WithdrawWaitNumber is a free data retrieval call binding the contract method 0x14704970.
//
// Solidity: function withdrawWaitNumber() constant returns(uint256)
func (_POC_1_0_0 *POC_1_0_0CallerSession) WithdrawWaitNumber() (*big.Int, error) {
	return _POC_1_0_0.Contract.WithdrawWaitNumber(&_POC_1_0_0.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "changeOwner", _newOwner, _number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0Session) ChangeOwner(_newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.ChangeOwner(&_POC_1_0_0.TransactOpts, _newOwner, _number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x56df3db1.
//
// Solidity: function changeOwner(_newOwner address, _number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) ChangeOwner(_newOwner common.Address, _number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.ChangeOwner(&_POC_1_0_0.TransactOpts, _newOwner, _number)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) Deposit(opts *bind.TransactOpts, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "deposit", _r, _s, _v)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC_1_0_0 *POC_1_0_0Session) Deposit(_r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Deposit(&_POC_1_0_0.TransactOpts, _r, _s, _v)
}

// Deposit is a paid mutator transaction binding the contract method 0xf2a558f9.
//
// Solidity: function deposit(_r bytes32, _s bytes32, _v uint8) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) Deposit(_r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Deposit(&_POC_1_0_0.TransactOpts, _r, _s, _v)
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerEmptyBlackList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerEmptyBlackList")
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC_1_0_0 *POC_1_0_0Session) OwnerEmptyBlackList() (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerEmptyBlackList(&_POC_1_0_0.TransactOpts)
}

// OwnerEmptyBlackList is a paid mutator transaction binding the contract method 0x3853d0a1.
//
// Solidity: function ownerEmptyBlackList() returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerEmptyBlackList() (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerEmptyBlackList(&_POC_1_0_0.TransactOpts)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerPopBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerPopBlackList", minerAddress)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Session) OwnerPopBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPopBlackList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPopBlackList is a paid mutator transaction binding the contract method 0x1103ec96.
//
// Solidity: function ownerPopBlackList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerPopBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPopBlackList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerPopLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerPopLockList", minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Session) OwnerPopLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPopLockList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPopLockList is a paid mutator transaction binding the contract method 0x6fae4a5f.
//
// Solidity: function ownerPopLockList(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerPopLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPopLockList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerPushBlackList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerPushBlackList", minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) OwnerPushBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPushBlackList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPushBlackList is a paid mutator transaction binding the contract method 0x31f23668.
//
// Solidity: function ownerPushBlackList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerPushBlackList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPushBlackList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerPushLockList(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerPushLockList", minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) OwnerPushLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPushLockList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerPushLockList is a paid mutator transaction binding the contract method 0x478e7d6e.
//
// Solidity: function ownerPushLockList(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerPushLockList(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerPushLockList(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerSetInitBlockNumber is a paid mutator transaction binding the contract method 0xf30011b4.
//
// Solidity: function ownerSetInitBlockNumber(_number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerSetInitBlockNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerSetInitBlockNumber", _number)
}

// OwnerSetInitBlockNumber is a paid mutator transaction binding the contract method 0xf30011b4.
//
// Solidity: function ownerSetInitBlockNumber(_number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0Session) OwnerSetInitBlockNumber(_number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerSetInitBlockNumber(&_POC_1_0_0.TransactOpts, _number)
}

// OwnerSetInitBlockNumber is a paid mutator transaction binding the contract method 0xf30011b4.
//
// Solidity: function ownerSetInitBlockNumber(_number uint256) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerSetInitBlockNumber(_number *big.Int) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerSetInitBlockNumber(&_POC_1_0_0.TransactOpts, _number)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Transactor) OwnerStop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "ownerStop", minerAddress)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0Session) OwnerStop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerStop(&_POC_1_0_0.TransactOpts, minerAddress)
}

// OwnerStop is a paid mutator transaction binding the contract method 0x2988fcfd.
//
// Solidity: function ownerStop(minerAddress address) returns(int8)
func (_POC_1_0_0 *POC_1_0_0TransactorSession) OwnerStop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.OwnerStop(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) Start(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "start", minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) Start(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Start(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Start is a paid mutator transaction binding the contract method 0xdd0b281e.
//
// Solidity: function start(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) Start(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Start(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) Stop(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "stop", minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) Stop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Stop(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Stop is a paid mutator transaction binding the contract method 0x656cb0bc.
//
// Solidity: function stop(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) Stop(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Stop(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) Withdraw(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "withdraw", minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) Withdraw(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Withdraw(&_POC_1_0_0.TransactOpts, minerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) Withdraw(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.Withdraw(&_POC_1_0_0.TransactOpts, minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Transactor) WithdrawSurplus(opts *bind.TransactOpts, minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.contract.Transact(opts, "withdrawSurplus", minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0Session) WithdrawSurplus(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.WithdrawSurplus(&_POC_1_0_0.TransactOpts, minerAddress)
}

// WithdrawSurplus is a paid mutator transaction binding the contract method 0x65e73e32.
//
// Solidity: function withdrawSurplus(minerAddress address) returns()
func (_POC_1_0_0 *POC_1_0_0TransactorSession) WithdrawSurplus(minerAddress common.Address) (*types.Transaction, error) {
	return _POC_1_0_0.Contract.WithdrawSurplus(&_POC_1_0_0.TransactOpts, minerAddress)
}

// TribeChief_1_0_0ABI is the input ABI used to generate the binding from.
const TribeChief_1_0_0ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_nextRoundSignerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextRoundSignerList\",\"outputs\":[{\"name\":\"nextRoundSignerList\",\"type\":\"address[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_signerList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"baseAddress\",\"type\":\"address\"},{\"name\":\"pocAddress\",\"type\":\"address\"},{\"name\":\"startBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_1_0_0Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_1_0_0Bin = `0x60c0604052600560808190527f312e302e3000000000000000000000000000000000000000000000000000000060a0908152620000409160009190620003ad565b503480156200004e57600080fd5b5060405160608062001a3d833981018060405260608110156200007057600080fd5b508051602082015160409283015160018054600160a060020a031916600160a060020a03808616919091179182905585517ff09a40160000000000000000000000000000000000000000000000000000000081528185166004820152306024820152955194959394929391169163f09a40169160448082019260009290919082900301818387803b1580156200010557600080fd5b505af11580156200011a573d6000803e3d6000fd5b5050506004828155600154604080517ff0d55817000000000000000000000000000000000000000000000000000000008152905160609450600160a060020a039092169263f0d55817928282019260009290829003018186803b1580156200018157600080fd5b505afa15801562000196573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015620001c057600080fd5b810190808051640100000000811115620001d957600080fd5b82016020810184811115620001ed57600080fd5b81518560208202830111640100000000821117156200020b57600080fd5b5050805190945060001092506200022491505057600080fd5b6001600560008360008151811015156200023a57fe5b90602001906020020151600160a060020a0316600160a060020a031681526020019081526020016000208190555060028160008151811015156200027a57fe5b6020908102919091018101518254600181018455600093845291909220018054600160a060020a031916600160a060020a039092169190911790556002545b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200032457600080fd5b505afa15801562000339573d6000803e3d6000fd5b505050506040513d60208110156200035057600080fd5b5051811015620003a25760028054600181810183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018054600160a060020a031916905501620002b9565b505050505062000452565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003f057805160ff191683800117855562000420565b8280016001018555821562000420579182015b828111156200042057825182559160200191906001019062000403565b506200042e92915062000432565b5090565b6200044f91905b808211156200042e576000815560010162000439565b90565b6115db80620004626000396000f3fe608060405234801561001057600080fd5b50600436106100b85760003560e060020a9004806357e871e71161007557806357e871e71461042a578063757991a814610444578063961c5c7a1461044c578063b4deb2ca14610454578063d7ca4a1c14610471578063eb5c001114610519576100b8565b80631c1b8772146100bd5780631fcb9ffb146100e557806320c1a5181461011e5780634e69d5601461021157806353f197fe1461034a57806354fd4d50146103ad575b600080fd5b6100e3600480360360208110156100d357600080fd5b5035600160a060020a0316610521565b005b610102600480360360208110156100fb57600080fd5b50356108b2565b60408051600160a060020a039092168252519081900360200190f35b6101c16004803603602081101561013457600080fd5b81019060208101813564010000000081111561014f57600080fd5b82018360208201111561016157600080fd5b8035906020019184602083028401116401000000008311171561018357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108da945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101fd5781810151838201526020016101e5565b505050509050019250505060405180910390f35b6102196108e0565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b83811015610271578181015183820152602001610259565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b838110156102b0578181015183820152602001610298565b50505050905001858103835289818151815260200191508051906020019060200280838360005b838110156102ef5781810151838201526020016102d7565b50505050905001858103825288818151815260200191508051906020019060200280838360005b8381101561032e578181015183820152602001610316565b505050509050019a505050505050505050505060405180910390f35b610352610b1d565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b83811015610398578181015183820152602001610380565b50505050905001935050505060405180910390f35b6103b5610b86565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103ef5781810151838201526020016103d7565b50505050905090810190601f16801561041c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610432610c1d565b60408051918252519081900360200190f35b610432610c23565b610432610cb2565b6101026004803603602081101561046a57600080fd5b5035610d10565b610479610d1e565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b838110156104c35781810151838201526020016104ab565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156105025781810151838201526020016104ea565b505050509050019550505050505060405180910390f35b610432610d61565b3380151561052e57600080fd5b600160a060020a03811660009081526005602052604081205411806105e15750600154604080517fdb512e85000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163db512e85916024808301926020929190829003018186803b1580156105b457600080fd5b505afa1580156105c8573d6000803e3d6000fd5b505050506040513d60208110156105de57600080fd5b50515b15156105ec57600080fd5b436004908155600154604080517f6da508300000000000000000000000000000000000000000000000000000000081529051600093600160a060020a0390931692636da5083092808201926020929091829003018186803b15801561065057600080fd5b505afa158015610664573d6000803e3d6000fd5b505050506040513d602081101561067a57600080fd5b5051600454909150600090829081151561069057fe5b06905060006002828154811015156106a457fe5b6000918252602082200154600160a060020a0316915082111561087757600160a060020a038516156106d9576106d985610dbf565b600160a060020a038116158015906106fa575033600160a060020a03821614155b1561087757600154604080517f1dd47d3d000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000939290921691631dd47d3d9160248082019260209290919082900301818787803b15801561076a57600080fd5b505af115801561077e573d6000803e3d6000fd5b505050506040513d602081101561079457600080fd5b505160000b131561081f57600154604080517fc8ba1fd8000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163c8ba1fd891602480830192600092919082900301818387803b15801561080657600080fd5b505af115801561081a573d6000803e3d6000fd5b505050505b600160a060020a0381166000908152600560205260408120819055600280548490811061084857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b6001830382141561088a5761088a610ef5565b610892610c23565b4381151561089c57fe5b0615156108ab576108ab610f7a565b5050505050565b60038054829081106108c057fe5b600091825260209091200154600160a060020a0316905081565b50606090565b606080606080600080600280549050604051908082528060200260200182016040528015610918578160200160208202803883390190505b506002546040805182815260208084028201019091529195508015610947578160200160208202803883390190505b50925060005b6002548110156109ca576000858281518110151561096757fe5b60209081029091010152600280546005916000918490811061098557fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484518590839081106109b857fe5b6020908102909101015260010161094d565b50600160009054906101000a9004600160a060020a0316600160a060020a031663b175de176040518163ffffffff1660e060020a02815260040160006040518083038186803b158015610a1c57600080fd5b505afa158015610a30573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610a5957600080fd5b810190808051640100000000811115610a7157600080fd5b82016020810184811115610a8457600080fd5b8151856020820283011164010000000082111715610aa157600080fd5b505092919050505094506002805480602002602001604051908101604052809291908181526020018280548015610b0157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610ae3575b5050505050955060045490506003805490509150909192939495565b606060006003805480602002602001604051908101604052809291908181526020018280548015610b7757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b59575b50505050509150815190509091565b60008054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610c125780601f10610be757610100808354040283529160200191610c12565b820191906000526020600020905b815481529060010190602001808311610bf557829003601f168201915b505050505090505b90565b60045481565b600154604080517f4067f0c80000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691634067f0c8916004808301926020929190829003018186803b158015610c8157600080fd5b505afa158015610c95573d6000803e3d6000fd5b505050506040513d6020811015610cab57600080fd5b5051905090565b600154604080517f282f78f20000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163282f78f2916004808301926020929190829003018186803b158015610c8157600080fd5b60028054829081106108c057fe5b6060806000610d2b610b1d565b6040805182815260208084028201019091529194509150818015610d59578160200160208202803883390190505b509150909192565b600154604080517f6da508300000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691636da50830916004808301926020929190829003018186803b158015610c8157600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a031663282f78f26040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610e1057600080fd5b505afa158015610e24573d6000803e3d6000fd5b505050506040513d6020811015610e3a57600080fd5b50516003541015610ea257600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038316179055610ef2565b6040805160e560020a62461bcd02815260206004820152601a60248201527f6e65787420726f756e64207369676e657220746f6f206d756368000000000000604482015290519081900360640190fd5b50565b6000610eff610fe7565b9050600160a060020a0381161515610f61576040805160e560020a62461bcd02815260206004820152601860248201527f7369676e657230206d757374206e6f74206265207a65726f0000000000000000604482015290519081900360640190fd5b610f6a816110b8565b610f72611278565b610ef261139c565b600160009054906101000a9004600160a060020a0316600160a060020a0316637216276b6040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015610fcd57600080fd5b505af1158015610fe1573d6000803e3d6000fd5b50505050565b60008060026000815481101515610ffa57fe5b6000918252602082200154600160a060020a031691505b6002548110156110a45760006005600060028481548110151561103057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812091909155600280548390811061106657fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101611011565b5060006110b2600282611547565b50905090565b600154604080517ff0d558170000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163f0d55817916004808301926000929190829003018186803b15801561111657600080fd5b505afa15801561112a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561115357600080fd5b81019080805164010000000081111561116b57600080fd5b8201602081018481111561117e57600080fd5b815185602082028301116401000000008211171561119b57600080fd5b50909450600093505050505b815181101561123d57600082828151811015156111c057fe5b90602001906020020151905080600160a060020a031684600160a060020a031614156112345760018351038214156112185761121383600081518110151561120457fe5b90602001906020020151611406565b61122c565b61122c838360010181518110151561120457fe5b505050610ef2565b506001016111a7565b5060405160e560020a62461bcd02815260040180806020018281038252602181526020018061158f6021913960400191505060405180910390fd5b60005b6003548110156112bf57600060038281548110151561129657fe5b600091825260209091200154600160a060020a031690506112b681611406565b5060010161127b565b506002545b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561131557600080fd5b505afa158015611329573d6000803e3d6000fd5b505050506040513d602081101561133f57600080fd5b5051811015610ef25760028054600181810183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff19169055016112c4565b60005b6003548110156113f85760006003828154811015156113ba57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560010161139f565b506000610ef2600382611547565b600160009054906101000a9004600160a060020a0316600160a060020a0316636da508306040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561145757600080fd5b505afa15801561146b573d6000803e3d6000fd5b505050506040513d602081101561148157600080fd5b505160025410156114f7576002805460018082019092557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117909155600090815260056020526040902055610ef2565b6040805160e560020a62461bcd02815260206004820152600f60248201527f746f6f206d616e79207369676e65720000000000000000000000000000000000604482015290519081900360640190fd5b81548183558181111561156b5760008381526020902061156b918101908301611570565b505050565b610c1a91905b8082111561158a5760008155600101611576565b509056fe7369676e657230206d75737420657869737420696e206c6561646572206c697374a165627a7a723058204071393096383a2bc904cb8ebb6cbe106726ead9dbfc71ca9c9bb0bcfdc2df0d0029`

// DeployTribeChief_1_0_0 deploys a new Ethereum contract, binding an instance of TribeChief_1_0_0 to it.
func DeployTribeChief_1_0_0(auth *bind.TransactOpts, backend bind.ContractBackend, baseAddress common.Address, pocAddress common.Address, startBlockNumber *big.Int) (common.Address, *types.Transaction, *TribeChief_1_0_0, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_1_0_0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_1_0_0Bin), backend, baseAddress, pocAddress, startBlockNumber)
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

// NextRoundSignerList is a free data retrieval call binding the contract method 0x1fcb9ffb.
//
// Solidity: function _nextRoundSignerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) NextRoundSignerList(opts *bind.CallOptsWithNumber, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "_nextRoundSignerList", arg0)
	return *ret0, err
}

// NextRoundSignerList is a free data retrieval call binding the contract method 0x1fcb9ffb.
//
// Solidity: function _nextRoundSignerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) NextRoundSignerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.NextRoundSignerList(&_TribeChief_1_0_0.CallOpts, arg0)
}

// NextRoundSignerList is a free data retrieval call binding the contract method 0x1fcb9ffb.
//
// Solidity: function _nextRoundSignerList( uint256) constant returns(address)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) NextRoundSignerList(arg0 *big.Int) (common.Address, error) {
	return _TribeChief_1_0_0.Contract.NextRoundSignerList(&_TribeChief_1_0_0.CallOpts, arg0)
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

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Caller) GetNextRoundSignerList(opts *bind.CallOptsWithNumber) (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	ret := new(struct {
		NextRoundSignerList []common.Address
		Length              *big.Int
	})
	out := ret
	err := _TribeChief_1_0_0.contract.CallWithNumber(opts, out, "getNextRoundSignerList")
	return *ret, err
}

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0Session) GetNextRoundSignerList() (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetNextRoundSignerList(&_TribeChief_1_0_0.CallOpts)
}

// GetNextRoundSignerList is a free data retrieval call binding the contract method 0x53f197fe.
//
// Solidity: function getNextRoundSignerList() constant returns(nextRoundSignerList address[], length uint256)
func (_TribeChief_1_0_0 *TribeChief_1_0_0CallerSession) GetNextRoundSignerList() (struct {
	NextRoundSignerList []common.Address
	Length              *big.Int
}, error) {
	return _TribeChief_1_0_0.Contract.GetNextRoundSignerList(&_TribeChief_1_0_0.CallOpts)
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
