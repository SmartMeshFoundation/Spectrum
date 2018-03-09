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
const TribeChiefABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChiefBin is the compiled bytecode used for deploying new contracts.
const TribeChiefBin = `0x606060405260408051908101604052600581527f302e302e32000000000000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000246565b5061012c60015534156200006057600080fd5b60405162000d9438038062000d94833981016040528080519091019050600080808080808651955060008611156200010957600094505b858510156200010357868581518110620000ad57fe5b90602001906020020151600160a060020a0381166000908152600260205260409020805460ff191660011790559350620000f784600364010000000062000888620001e082021704565b60019094019362000097565b620001d3565b5050734110bd1ff0b73fa12c259acf39c950277f266787600081905260026020527ff0054db7accc4f10966f99d7cf81a202687bc788eef5a559cb9a99f0ac4cb305805460ff19166001179055905073ca012e2facf405885293d8d3ba3f14fae1e58b7173adb3ea3ad356199206ca817b04fd668cc5196df26200019d83600364010000000062000888620001e082021704565b620001b882600364010000000062000888620001e082021704565b620001d381600364010000000062000888620001e082021704565b5050505050505062000317565b600454601e90101562000242576004805460018101620002018382620002cb565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600690526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028957805160ff1916838001178555620002b9565b82800160010185558215620002b9579182015b82811115620002b95782518255916020019190600101906200029c565b50620002c7929150620002f7565b5090565b815481835581811511620002f257600083815260209020620002f2918101908301620002f7565b505050565b6200031491905b80821115620002c75760008155600101620002fe565b90565b610a6d80620003276000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b8772811461005b5780634e69d5601461007c57806354fd4d50146101b9575b600080fd5b341561006657600080fd5b61007a600160a060020a0360043516610243565b005b341561008757600080fd5b61008f61045f565b604051808060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019060200280838360005b838110156100e15780820151838201526020016100c9565b50505050905001858103845289818151815260200191508051906020019060200280838360005b83811015610120578082015183820152602001610108565b50505050905001858103835288818151815260200191508051906020019060200280838360005b8381101561015f578082015183820152602001610147565b50505050905001858103825287818151815260200191508051906020019060200280838360005b8381101561019e578082015183820152602001610186565b50505050905001995050505050505050505060405180910390f35b34156101c457600080fd5b6101cc610649565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102085780820151838201526020016101f0565b50505050905090810190601f1680156102355780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b33600160a060020a038116600090815260066020526040812060010154909182918291829182901161027457600080fd5b4360038190556001549011801561029657506001544381151561029357fe5b06155b156102fd576005549450600093505b848410156102f157600760006005868154811015156102c057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812055600193909301926102a5565b6102fd600560006109d3565b60045460035481151561030c57fe5b0692506002600060048581548110151561032257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615156103eb576006600060048581548110151561036357fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060048054919350908490811061039957fe5b60009182526020909120015433600160a060020a039081169116146103e657815460019011156103d857815460001901825560035460018301556103e1565b6103e1836106f2565b6103eb565b600382555b600160a060020a0386161561040357610403866107da565b600454601e901080156104195750600554600090115b156104575761044d6005600081548110151561043157fe5b600091825260209091200154600160a060020a03166003610888565b61045760006108f8565b505050505050565b6104676109f1565b61046f6109f1565b6104776109f1565b61047f6109f1565b60045460009081906040518059106104945750595b90808252806020026020018201604052506004549094506040518059106104b85750595b90808252806020026020018201604052509250600090505b60045481101561058257600660006004838154811015156104ed57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205484828151811061051e57fe5b60209081029091010152600480546006916000918490811061053c57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015483828151811061057057fe5b602090810290910101526001016104d0565b60058054806020026020016040519081016040528092919081815260200182805480156105d857602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105ba575b50505050509550600480548060200260200160405190810160405280929190818152602001828054801561063557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610617575b505050505094506003549150509091929394565b6106516109f1565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106e75780601f106106bc576101008083540402835291602001916106e7565b820191906000526020600020905b8154815290600101906020018083116106ca57829003601f168201915b505050505090505b90565b6004546000818310156107d5576006600060048581548110151561071257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b600182038110156107c057600480546001830190811061075c57fe5b60009182526020909120015460048054600160a060020a03909216918390811061078257fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610740565b6004805460001901906107d39082610a03565b505b505050565b6005546004901080156108035750600160a060020a038116600090815260076020526040902054155b80156108285750600160a060020a038116600090815260066020526040902060010154155b1561088557600580546001810161083f8382610a03565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038516908117909155825260079052604090204390555b50565b600454601e9010156108f45760048054600181016108a68382610a03565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558252600690526040902081815543600191909101555b5050565b6005546000818310156107d5576007600060058581548110151561091857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b600182038110156109c057600580546001830190811061095c57fe5b60009182526020909120015460058054600160a060020a03909216918390811061098257fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610940565b6005805460001901906107d39082610a03565b50805460008255906000526020600020908101906108859190610a23565b60206040519081016040526000815290565b8154818355818115116107d5576000838152602090206107d59181019083015b6106ef91905b80821115610a3d5760008155600101610a29565b50905600a165627a7a72305820b2c19c242d3ef4e0491ac65d4b1ea07059ddc52cacb707de957c85f9d58ad7110029`

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
