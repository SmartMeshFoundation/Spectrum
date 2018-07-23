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

// TribeChief_0_0_6ABI is the input ABI used to generate the binding from.
const TribeChief_0_0_6ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"},{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_signerLimit\",\"type\":\"uint256\"},{\"name\":\"_volunteerLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChief_0_0_6Bin is the compiled bytecode used for deploying new contracts.
const TribeChief_0_0_6Bin = `0x606060405260408051908101604052600581527f302e302e36000000000000000000000000000000000000000000000000000000602082015260009080516200004d929160200190620002ce565b5061181b60015560116002556104d260035534156200006b57600080fd5b60405162001a3a38038062001a3a83398101604052808051820191906020018051919060200180519190602001805191506000905080808080871115620000b25760018790555b6000861115620000c25760028690555b6000851115620000d25760038590555b60068054600160a060020a03191633600160a060020a0316179055875193506000841115620001b757600092505b83831015620001b1578783815181106200011657fe5b90602001906020020151600160a060020a0381166000908152600460205260409020805460ff1916600190811790915560078054929450919081016200015d838262000353565b5060009182526020909120018054600160a060020a031916600160a060020a038416179055821515620001a557620001a5826003640100000000620011286200026882021704565b60019092019162000100565b6200025a565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260046020527f052387a23a063359ef51016da56c6ef818568f4a38e27c2728fda35f7b8ae85e805460ff1916600190811790915560078054909181016200021a838262000353565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556200025a816003640100000000620002688102620011281704565b50505050505050506200039f565b6002546008541015620002ca57600880546001810162000289838262000353565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600b90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200031157805160ff191683800117855562000341565b8280016001018555821562000341579182015b828111156200034157825182559160200191906001019062000324565b506200034f9291506200037f565b5090565b8154818355818115116200037a576000838152602090206200037a9181019083016200037f565b505050565b6200039c91905b808211156200034f576000815560010162000386565b90565b61168b80620003af6000396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b8772811461009257806320c1a518146100b35780634e69d5601461015557806354fd4d5014610299578063757991a814610323578063961c5c7a14610348578063d7ca4a1c1461035b578063eb5c00111461040e575b600080fd5b341561009d57600080fd5b6100b1600160a060020a0360043516610421565b005b34156100be57600080fd5b61010260046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506104d895505050505050565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610141578082015183820152602001610129565b505050509050019250505060405180910390f35b341561016057600080fd5b6101686105eb565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b838110156101c05780820151838201526020016101a8565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b838110156101ff5780820151838201526020016101e7565b50505050905001858103835289818151815260200191508051906020019060200280838360005b8381101561023e578082015183820152602001610226565b50505050905001858103825288818151815260200191508051906020019060200280838360005b8381101561027d578082015183820152602001610265565b505050509050019a505050505050505050505060405180910390f35b34156102a457600080fd5b6102ac6107df565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102e85780820151838201526020016102d0565b50505050905090810190601f1680156103155780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561032e57600080fd5b610336610888565b60405190815260200160405180910390f35b341561035357600080fd5b61033661088e565b341561036657600080fd5b61036e610894565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b838110156103b85780820151838201526020016103a0565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156103f75780820151838201526020016103df565b505050509050019550505050505060405180910390f35b341561041957600080fd5b610336610998565b33600160a060020a038116151561043757600080fd5b600160a060020a0381166000908152600b60205260408120541161045a57600080fd5b4360058190556001549011801561047e575060015460055481151561047b57fe5b06155b1561048b5761048b61099e565b600160a060020a038216156104a5576104a58260056109f8565b60025460085410806104ba5750600854600954105b156104cc576104c7610b9d565b6104d4565b6104d4610d20565b5050565b6104e06115f1565b60008083516040518059106104f25750595b90808252806020026020018201604052506003546009549194509010156105e457600091505b83518210156105e45783828151811061052d57fe5b90602001906020020151600160a060020a0381166000908152600c60205260409020600101549091501580156105795750600160a060020a0381166000908152600d6020526040902054155b801561059e5750600160a060020a0381166000908152600b6020526040902060010154155b156105c05760018383815181106105b157fe5b602090810290910101526105d9565b60008383815181106105ce57fe5b602090810290910101525b600190910190610518565b5050919050565b6105f36115f1565b6105fb6115f1565b6106036115f1565b61060b6115f1565b60008060006008805490506040518059106106235750595b90808252806020026020018201604052506008549095506040518059106106475750595b90808252806020026020018201604052509350600090505b60085481101561071157600b600060088381548110151561067c57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548582815181106106ad57fe5b6020908102909101015260088054600b91600091849081106106cb57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600101548482815181106106ff57fe5b6020908102909101015260010161065f565b600954600a80549194509060208082020160405190810160405280929190818152602001828054801561076d57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161074f575b5050505050955060088054806020026020016040519081016040528092919081815260200182805480156107ca57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107ac575b50505050509650600554915050909192939495565b6107e76115f1565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561087d5780601f106108525761010080835404028352916020019161087d565b820191906000526020600020905b81548152906001019060200180831161086057829003601f168201915b505050505090505b90565b60015490565b60035490565b61089c6115f1565b6108a46115f1565b60095460009081906040518059106108b95750595b90808252806020026020018201604052509250600980548060200260200160405190810160405280929190818152602001828054801561092257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610904575b50505050509350600090505b60095481101561098e57600c600060098381548110151561094b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205483828151811061097c57fe5b6020908102909101015260010161092e565b8351915050909192565b60025490565b600a5460005b818110156109ec57600d6000600a838154811015156109bf57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016109a4565b6104d4600a6000611603565b6000811515610a6c576009546000901115610a5e575060005b600954811015610a5e576009805482908110610a2957fe5b600091825260209091200154600160a060020a0384811691161415610a5657610a5181610df7565b610a5e565b600101610a11565b610a6783610ed1565b610b98565b816005148015610a7f5750600354600954105b8015610aa45750600160a060020a0383166000908152600c6020526040902060010154155b8015610ac65750600160a060020a0383166000908152600d6020526040902054155b8015610aeb5750600160a060020a0383166000908152600b6020526040902060010154155b15610b47576009805460018101610b028382611621565b5060009182526020808320919091018054600160a060020a031916600160a060020a0387169081179091558252600c9052604090208281554360019190910155610b98565b600582108015610b705750600160a060020a0383166000908152600c6020526040812060010154115b15610b9857600160a060020a0383166000908152600c60205260409020828155436001909101555b505050565b600080610ba8610f52565b600854600554811515610bb757fe5b06915060046000600884815481101515610bcd57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161515610cc257600b6000600884815481101515610c0e57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600880549192509083908110610c4457fe5b60009182526020909120015433600160a060020a03908116911614610cbd5780546001901115610c835780546000190181556005546001820155610cb8565b610caf600883815481101515610c9557fe5b6000918252602082200154600160a060020a0316906109f8565b610cb882611054565b610cc2565b600381555b600254600854108015610cd85750600954600090115b156104d45760098054610d1091906000198101908110610cf457fe5b600091825260209091200154600160a060020a03166003611128565b6009546104d49060001901610df7565b60085460055460009081908390811515610d3657fe5b069150600882815481101515610d4857fe5b6000918252602082200154600160a060020a03169150821115610de45780600160a060020a031633600160a060020a0316141515610de457610d8b8160006109f8565b600160a060020a0381166000908152600b602052604081208181556001018190556008805484908110610dba57fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790555b60018303821415610b9857610b9861118b565b600954600081831015610b9857600c6000600985815481101515610e1757fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610eb8576009805460018301908110610e6157fe5b60009182526020909120015460098054600160a060020a039092169183908110610e8757fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101610e45565b600980546000190190610ecb9082611621565b50505050565b600160a060020a03811615801590610eff5750600160a060020a0381166000908152600d6020526040902054155b15610f4f57600a805460018101610f168382611621565b5060009182526020808320919091018054600160a060020a031916600160a060020a0385169081179091558252600d9052604090204390555b50565b60008080805b600854841015610fab57600880546000919086908110610f7457fe5b600091825260209091200154600160a060020a03161415610fa057610f9884611054565b600019909301925b600190930192610f58565b600092505b600854831015610ecb576008805484908110610fc857fe5b6000918252602080832090910154600160a060020a0316808352600c90915260408220600101549093501115611049575060005b60095481101561104957600980548290811061101457fe5b600091825260209091200154600160a060020a03838116911614156110415761103c81610df7565b611049565b600101610ffc565b600190920191610fb0565b600854600081831015610b9857600b600060088581548110151561107457fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b600182038110156111155760088054600183019081106110be57fe5b60009182526020909120015460088054600160a060020a0390921691839081106110e457fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556001016110a2565b600880546000190190610ecb9082611621565b60025460085410156104d45760088054600181016111468382611621565b5060009182526020808320919091018054600160a060020a038616600160a060020a031990911681179091558252600b90526040902081815543600191909101555050565b60006111956115f1565b60008060008060006111a56115f1565b600080600080600860008154811015156111bb57fe5b600091825260209091200154600854600160a060020a039091169c506040518059106111e45750595b90808252806020026020018201604052509a50600099505b8a518a101561125657600880548b90811061121357fe5b600091825260209091200154600160a060020a03168b8b8151811061123457fe5b600160a060020a039092166020928302909101909101526001909901986111fc565b8a5198505b60008911156113065760018903975061127388611054565b8a888151811061127f57fe5b9060200190602002015196506000881180156112a35750600160a060020a03871615155b156112fa57600160a060020a0387166000908152600c602052604090205415156112d2576112d28760056109f8565b600160a060020a0387166000908152600c60205260409020546112fa908890600019016109f8565b6000199098019761125b565b600160a060020a038c1660009081526004602052604090205460ff1680156113315750600754600190115b156113c957600095505b6007548610156113c4578b600160a060020a031660078781548110151561135e57fe5b600091825260209091200154600160a060020a031614156113b9576007546000190186141561139f5761139a60076000815481101515610cf457fe5b6113b4565b6113b4600787600101815481101515610cf457fe5b6113c4565b60019095019461133b565b6113dc565b6113dc60076000815481101515610cf457fe5b6002546040518059106113ec5750595b9080825280602002602001820160405250945060009350600092505b6009548310156114a7576002548410611420576114a7565b61145660098481548110151561143257fe5b600091825260209091200154600954600160a060020a039091169060001901611538565b9150611462858361159f565b1561146c5761149c565b61147e600983815481101515610cf457fe5b8185858151811061148b57fe5b602090810290910101526001909301925b600190920191611408565b60025484101561152a575060005b60095481101561152a57600b60006009838154811015156114d257fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154151561151257611512600982815481101515610cf457fe5b600254600854106115225761152a565b6001016114b5565b505050505050505050505050565b60008080831161154b5760009150611598565b834442604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390209050828181151561159457fe5b0691505b5092915050565b6000806000845111156115e7575060005b83518110156115e757828482815181106115c657fe5b9060200190602002015114156115df5760019150611598565b6001016115b0565b5060009392505050565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610f4f9190611641565b815481835581811511610b9857600083815260209020610b989181019083015b61088591905b8082111561165b5760008155600101611647565b50905600a165627a7a72305820c16ce7ea581cccea9bf68aa019a2aa46a6eb997273e77b7900b359b70559a2470029`

// DeployTribeChief_0_0_6 deploys a new Ethereum contract, binding an instance of TribeChief_0_0_6 to it.
func DeployTribeChief_0_0_6(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address, _epoch *big.Int, _signerLimit *big.Int, _volunteerLimit *big.Int) (common.Address, *types.Transaction, *TribeChief_0_0_6, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_6ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChief_0_0_6Bin), backend, genesisSigners, _epoch, _signerLimit, _volunteerLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief_0_0_6{TribeChief_0_0_6Caller: TribeChief_0_0_6Caller{contract: contract}, TribeChief_0_0_6Transactor: TribeChief_0_0_6Transactor{contract: contract}}, nil
}

// TribeChief_0_0_6 is an auto generated Go binding around an Ethereum contract.
type TribeChief_0_0_6 struct {
	TribeChief_0_0_6Caller     // Read-only binding to the contract
	TribeChief_0_0_6Transactor // Write-only binding to the contract
}

// TribeChief_0_0_6Caller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChief_0_0_6Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_6Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChief_0_0_6Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChief_0_0_6Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChief_0_0_6Session struct {
	Contract     *TribeChief_0_0_6       // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribeChief_0_0_6CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChief_0_0_6CallerSession struct {
	Contract *TribeChief_0_0_6Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribeChief_0_0_6TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChief_0_0_6TransactorSession struct {
	Contract     *TribeChief_0_0_6Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TribeChief_0_0_6Raw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChief_0_0_6Raw struct {
	Contract *TribeChief_0_0_6 // Generic contract binding to access the raw methods on
}

// TribeChief_0_0_6CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChief_0_0_6CallerRaw struct {
	Contract *TribeChief_0_0_6Caller // Generic read-only contract binding to access the raw methods on
}

// TribeChief_0_0_6TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChief_0_0_6TransactorRaw struct {
	Contract *TribeChief_0_0_6Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief_0_0_6 creates a new instance of TribeChief_0_0_6, bound to a specific deployed contract.
func NewTribeChief_0_0_6(address common.Address, backend bind.ContractBackend) (*TribeChief_0_0_6, error) {
	contract, err := bindTribeChief_0_0_6(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_6{TribeChief_0_0_6Caller: TribeChief_0_0_6Caller{contract: contract}, TribeChief_0_0_6Transactor: TribeChief_0_0_6Transactor{contract: contract}}, nil
}

// NewTribeChief_0_0_6Caller creates a new read-only instance of TribeChief_0_0_6, bound to a specific deployed contract.
func NewTribeChief_0_0_6Caller(address common.Address, caller bind.ContractCaller) (*TribeChief_0_0_6Caller, error) {
	contract, err := bindTribeChief_0_0_6(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_6Caller{contract: contract}, nil
}

// NewTribeChief_0_0_6Transactor creates a new write-only instance of TribeChief_0_0_6, bound to a specific deployed contract.
func NewTribeChief_0_0_6Transactor(address common.Address, transactor bind.ContractTransactor) (*TribeChief_0_0_6Transactor, error) {
	contract, err := bindTribeChief_0_0_6(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChief_0_0_6Transactor{contract: contract}, nil
}

// bindTribeChief_0_0_6 binds a generic wrapper to an already deployed contract.
func bindTribeChief_0_0_6(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChief_0_0_6ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_6 *TribeChief_0_0_6Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_6.Contract.TribeChief_0_0_6Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_6 *TribeChief_0_0_6Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.TribeChief_0_0_6Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_6 *TribeChief_0_0_6Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.TribeChief_0_0_6Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _TribeChief_0_0_6.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief_0_0_6 *TribeChief_0_0_6TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief_0_0_6 *TribeChief_0_0_6TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.contract.Transact(opts, method, params...)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) FilterVolunteer(opts *bind.CallOptsWithNumber, volunteers []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "filterVolunteer", volunteers)
	return *ret0, err
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_0_0_6.Contract.FilterVolunteer(&_TribeChief_0_0_6.CallOpts, volunteers)
}

// FilterVolunteer is a free data retrieval call binding the contract method 0x20c1a518.
//
// Solidity: function filterVolunteer(volunteers address[]) constant returns(result uint256[])
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) FilterVolunteer(volunteers []common.Address) ([]*big.Int, error) {
	return _TribeChief_0_0_6.Contract.FilterVolunteer(&_TribeChief_0_0_6.CallOpts, volunteers)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) GetEpoch(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetEpoch(&_TribeChief_0_0_6.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) GetEpoch() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetEpoch(&_TribeChief_0_0_6.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) GetSignerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "getSignerLimit")
	return *ret0, err
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetSignerLimit(&_TribeChief_0_0_6.CallOpts)
}

// GetSignerLimit is a free data retrieval call binding the contract method 0xeb5c0011.
//
// Solidity: function getSignerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) GetSignerLimit() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetSignerLimit(&_TribeChief_0_0_6.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) GetStatus(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_0_0_6.Contract.GetStatus(&_TribeChief_0_0_6.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(signerList address[], blackList address[], scoreList uint256[], numberList uint256[], totalVolunteer uint256, number uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) GetStatus() (struct {
	SignerList     []common.Address
	BlackList      []common.Address
	ScoreList      []*big.Int
	NumberList     []*big.Int
	TotalVolunteer *big.Int
	Number         *big.Int
}, error) {
	return _TribeChief_0_0_6.Contract.GetStatus(&_TribeChief_0_0_6.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) GetVolunteerLimit(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "getVolunteerLimit")
	return *ret0, err
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetVolunteerLimit(&_TribeChief_0_0_6.CallOpts)
}

// GetVolunteerLimit is a free data retrieval call binding the contract method 0x961c5c7a.
//
// Solidity: function getVolunteerLimit() constant returns(uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) GetVolunteerLimit() (*big.Int, error) {
	return _TribeChief_0_0_6.Contract.GetVolunteerLimit(&_TribeChief_0_0_6.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) GetVolunteers(opts *bind.CallOptsWithNumber) (struct {
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
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "getVolunteers")
	return *ret, err
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_0_0_6.Contract.GetVolunteers(&_TribeChief_0_0_6.CallOpts)
}

// GetVolunteers is a free data retrieval call binding the contract method 0xd7ca4a1c.
//
// Solidity: function getVolunteers() constant returns(volunteerList address[], weightList uint256[], length uint256)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) GetVolunteers() (struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}, error) {
	return _TribeChief_0_0_6.Contract.GetVolunteers(&_TribeChief_0_0_6.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Caller) Version(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief_0_0_6.contract.CallWithNumber(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) Version() (string, error) {
	return _TribeChief_0_0_6.Contract.Version(&_TribeChief_0_0_6.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief_0_0_6 *TribeChief_0_0_6CallerSession) Version() (string, error) {
	return _TribeChief_0_0_6.Contract.Version(&_TribeChief_0_0_6.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_6 *TribeChief_0_0_6Transactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_6.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_6 *TribeChief_0_0_6Session) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.Update(&_TribeChief_0_0_6.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief_0_0_6 *TribeChief_0_0_6TransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief_0_0_6.Contract.Update(&_TribeChief_0_0_6.TransactOpts, volunteer)
}
