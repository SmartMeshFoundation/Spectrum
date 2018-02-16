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
const TribeChiefABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"int256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChiefBin is the compiled bytecode used for deploying new contracts.
const TribeChiefBin = `0x60606040526005600155600f60025534156200001a57600080fd5b60405162000c7d38038062000c7d8339810160405280805160008054600160a060020a03191633600160a060020a031617815592019190508080808080865195506000861115620000e257600094505b85851015620000dc578685815181106200008057fe5b90602001906020020151600160a060020a0381166000908152600360208190526040909120805460ff19166001179055909450620000ce908590640100000000620008326200022782021704565b50506001909401936200006a565b6200021a565b5050600360208190527f39948882bf66b32f56593f076dc2f163de9642d786281d35732f0aa288e1a3b48054600160ff1991821681179092557fe8505a5ac9cb1417183861875ba9240bf4785e37ddba4197fc18d4309f87a24f805482168317905573b94b3aa41609e3f59cbaff3c2c298c6cc4c50b8160008190527f2bc1fab0e28baeabdf97c13de3f177853d63c25aed1eedd8ebec32fac85069a380549092169092179055734110bd1ff0b73fa12c259acf39c950277f266787925073adb3ea3ad356199206ca817b04fd668cc5196df291620001d1908490640100000000620002278102620008321704565b5050620001f4826003620002276401000000000262000832176401000000009004565b505062000217816003620002276401000000000262000832176401000000009004565b50505b5050505050505062000347565b600680546000918291606f9010620002465760009250829150620002ef565b80548190600181016200025a8382620002f7565b5060009182526020909120018054600160a060020a031916600160a060020a0387161790556001818101805490918101620002968382620002f7565b50600091825260209091200184905560038101805460018101620002bb8382620002f7565b506000918252602080832043920191909155600160a060020a03871682526002830190526040902084905580549250600191505b509250929050565b8154818355818115116200031e576000838152602090206200031e91810190830162000323565b505050565b6200034491905b808211156200034057600081556001016200032a565b5090565b90565b61092680620003576000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b877281146100505780634e69d56014610071575b600080fd5b341561005b57600080fd5b61006f600160a060020a03600435166101ae565b005b341561007c57600080fd5b610084610259565b604051808060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019060200280838360005b838110156100d65780820151838201526020016100be565b50505050905001858103845289818151815260200191508051906020019060200280838360005b838110156101155780820151838201526020016100fd565b50505050905001858103835288818151815260200191508051906020019060200280838360005b8381101561015457808201518382015260200161013c565b50505050905001858103825287818151815260200191508051906020019060200280838360005b8381101561019357808201518382015260200161017b565b50505050905001995050505050505050505060405180910390f35b436005556000600160a060020a038216158015906101d057506101d0826103f3565b156101e1576101de826103f9565b50505b6007805460009081106101f057fe5b6000918252602082200154915081131561024b576007805460001990920191829190600090811061021d57fe5b600091825260208220019190915560098054439290811061023a57fe5b600091825260209091200155610255565b61025560006104a1565b5050565b6102616107e0565b6102696107e0565b6102716107e0565b6102796107e0565b6000600a6000018054806020026020016040519081016040528092919081815260200182805480156102d457602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116102b6575b50505050509450600660000180548060200260200160405190810160405280929190818152602001828054801561033457602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610316575b50505050509350600660010180548060200260200160405190810160405280929190818152602001828054801561038a57602002820191906000526020600020905b815481526020019060010190808311610376575b5050505050925060066003018054806020026020016040519081016040528092919081815260200182805480156103e057602002820191906000526020600020905b8154815260200190600101908083116103cc575b5050505050915060055490509091929394565b50600190565b600a54600090819061014d90106104155750600090508061049c565b61041e83610676565b600a80546001810161043083826107f2565b5060009182526020909120018054600160a060020a031916600160a060020a038516179055600b80546001810161046783826107f2565b506000918252602080832043920191909155600a54600160a060020a0386168352600c90915260409091208190559150600190505b915091565b60065460008183106104b257610671565b60068054600891600091869081106104c657fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b600182038110156105cf57600680546001830190811061050a57fe5b60009182526020909120015460068054600160a060020a03909216918390811061053057fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600780546001830190811061056a57fe5b600091825260209091200154600780548390811061058457fe5b60009182526020909120015560098054600183019081106105a157fe5b60009182526020909120015460098054839081106105bb57fe5b6000918252602090912001556001016104ee565b6006805460001984019081106105e157fe5b60009182526020909120018054600160a060020a031916905560078054600019840190811061060c57fe5b600091825260208220015560098054600019840190811061062957fe5b600091825260208220015560068054600019019061064790826107f2565b5060078054600019019061065b90826107f2565b5060098054600019019061066f90826107f2565b505b505050565b600a54600160a060020a0382166000908152600c60205260408120549080808311156107d95750506000198101805b6001840381101561074f57600a8054600183019081106106c157fe5b600091825260209091200154600a8054600160a060020a0390921691839081106106e757fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600b80546001830190811061072157fe5b600091825260209091200154600b80548390811061073b57fe5b6000918252602090912001556001016106a5565b600a8054600019860190811061076157fe5b60009182526020909120018054600160a060020a0319169055600b8054600019860190811061078c57fe5b6000918252602082200155600a805460001901906107aa90826107f2565b50600b805460001901906107be90826107f2565b50600160a060020a0385166000908152600c60205260408120555b5050505050565b60206040519081016040526000815290565b8154818355818115116106715760008381526020902061067191810190830161082f91905b8082111561082b5760008155600101610817565b5090565b90565b600680546000918291606f901061084f57600092508291506108f2565b805481906001810161086183826107f2565b5060009182526020909120018054600160a060020a031916600160a060020a038716179055600181810180549091810161089b83826107f2565b506000918252602090912001849055600381018054600181016108be83826107f2565b506000918252602080832043920191909155600160a060020a03871682526002830190526040902084905580549250600191505b5092509290505600a165627a7a72305820b348e6c2155bf83e44a6233d29354e829f334e69efbc1662a1c73db6863b990d0029`

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
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCaller) GetStatus(opts *bind.CallOpts) (struct {
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
	err := _TribeChief.contract.Call(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
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
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
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
