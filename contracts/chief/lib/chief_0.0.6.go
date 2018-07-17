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
const TribeChief_0_0_6Bin = `0x60606040526020604051908101604052600080825290805162000027929160200190620002f2565b5060408051908101604052600581527f302e302e360000000000000000000000000000000000000000000000000000006020820152600190805162000071929160200190620002f2565b5061181b60025560116003556104d260045534156200008f57600080fd5b60405162001a5e38038062001a5e83398101604052808051820191906020018051919060200180519190602001805191506000905080808080871115620000d65760028790555b6000861115620000e65760038690555b6000851115620000f65760048590555b60078054600160a060020a03191633600160a060020a0316179055875193506000841115620001db57600092505b83831015620001d5578783815181106200013a57fe5b90602001906020020151600160a060020a0381166000908152600560205260409020805460ff19166001908117909155600880549294509190810162000181838262000377565b5060009182526020909120018054600160a060020a031916600160a060020a038416179055821515620001c957620001c9826003640100000000620010ae6200028c82021704565b60019092019162000124565b6200027e565b50734110bd1ff0b73fa12c259acf39c950277f266787600081905260056020527fe972f0cd07b79b44c16c7f378a5866e05e66a2c3fdcdb77d23dbf43a635beeee805460ff1916600190811790915560088054909181016200023e838262000377565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556200027e8160036401000000006200028c8102620010ae1704565b5050505050505050620003c3565b6003546009541015620002ee576009805460018101620002ad838262000377565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600c90526040902081815543600191909101555b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200033557805160ff191683800117855562000365565b8280016001018555821562000365579182015b828111156200036557825182559160200191906001019062000348565b5062000373929150620003a3565b5090565b8154818355818115116200039e576000838152602090206200039e918101908301620003a3565b505050565b620003c091905b80821115620003735760008155600101620003aa565b90565b61168b80620003d36000396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b8772811461009257806320c1a518146100b35780634e69d5601461015557806354fd4d5014610299578063757991a814610323578063961c5c7a14610348578063d7ca4a1c1461035b578063eb5c00111461040e575b600080fd5b341561009d57600080fd5b6100b1600160a060020a0360043516610421565b005b34156100be57600080fd5b610102600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061051895505050505050565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610141578082015183820152602001610129565b505050509050019250505060405180910390f35b341561016057600080fd5b61016861062b565b604051808060200180602001806020018060200187815260200186815260200185810385528b818151815260200191508051906020019060200280838360005b838110156101c05780820151838201526020016101a8565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b838110156101ff5780820151838201526020016101e7565b50505050905001858103835289818151815260200191508051906020019060200280838360005b8381101561023e578082015183820152602001610226565b50505050905001858103825288818151815260200191508051906020019060200280838360005b8381101561027d578082015183820152602001610265565b505050509050019a505050505050505050505060405180910390f35b34156102a457600080fd5b6102ac61081f565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102e85780820151838201526020016102d0565b50505050905090810190601f1680156103155780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561032e57600080fd5b6103366108c8565b60405190815260200160405180910390f35b341561035357600080fd5b6103366108ce565b341561036657600080fd5b61036e6108d4565b604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b838110156103b85780820151838201526020016103a0565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156103f75780820151838201526020016103df565b505050509050019550505050505060405180910390f35b341561041957600080fd5b6103366109d8565b33600160a060020a038116151561043757600080fd5b600160a060020a0381166000908152600c60205260408120541161045a57600080fd5b60408051908101604052600681527f7570646174650000000000000000000000000000000000000000000000000000602082015260009080516104a1929160200190611577565b50600254431180156104be5750600254438115156104bb57fe5b06155b156104cb576104cb6109de565b600160a060020a038216156104e5576104e5826005610a38565b60035460095410806104fa5750600954600a54105b1561050c57610507610bdd565b610514565b610514610da4565b5050565b6105206115f5565b60008083516040518059106105325750595b9080825280602002602001820160405250600454600a5491945090101561062457600091505b83518210156106245783828151811061056d57fe5b90602001906020020151600160a060020a0381166000908152600d60205260409020600101549091501580156105b95750600160a060020a0381166000908152600e6020526040902054155b80156105de5750600160a060020a0381166000908152600c6020526040902060010154155b156106005760018383815181106105f157fe5b60209081029091010152610619565b600083838151811061060e57fe5b602090810290910101525b600190910190610558565b5050919050565b6106336115f5565b61063b6115f5565b6106436115f5565b61064b6115f5565b60008060006009805490506040518059106106635750595b90808252806020026020018201604052506009549095506040518059106106875750595b90808252806020026020018201604052509350600090505b60095481101561075157600c60006009838154811015156106bc57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548582815181106106ed57fe5b6020908102909101015260098054600c916000918490811061070b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015484828151811061073f57fe5b6020908102909101015260010161069f565b600a54600b8054919450906020808202016040519081016040528092919081815260200182805480156107ad57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161078f575b50505050509550600980548060200260200160405190810160405280929190818152602001828054801561080a57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107ec575b50505050509650600654915050909192939495565b6108276115f5565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108bd5780601f10610892576101008083540402835291602001916108bd565b820191906000526020600020905b8154815290600101906020018083116108a057829003601f168201915b505050505090505b90565b60025490565b60045490565b6108dc6115f5565b6108e46115f5565b600a5460009081906040518059106108f95750595b90808252806020026020018201604052509250600a80548060200260200160405190810160405280929190818152602001828054801561096257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610944575b50505050509350600090505b600a548110156109ce57600d6000600a8381548110151561098b57fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020548382815181106109bc57fe5b6020908102909101015260010161096e565b8351915050909192565b60035490565b600b5460005b81811015610a2c57600e6000600b838154811015156109ff57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120556001016109e4565b610514600b6000611607565b6000811515610aac57600a546000901115610a9e575060005b600a54811015610a9e57600a805482908110610a6957fe5b600091825260209091200154600160a060020a0384811691161415610a9657610a9181610e7f565b610a9e565b600101610a51565b610aa783610f59565b610bd8565b816005148015610abf5750600454600a54105b8015610ae45750600160a060020a0383166000908152600d6020526040902060010154155b8015610b065750600160a060020a0383166000908152600e6020526040902054155b8015610b2b5750600160a060020a0383166000908152600c6020526040902060010154155b15610b8757600a805460018101610b428382611625565b5060009182526020808320919091018054600160a060020a031916600160a060020a0387169081179091558252600d9052604090208281554360019190910155610bd8565b600582108015610bb05750600160a060020a0383166000908152600d6020526040812060010154115b15610bd857600160a060020a0383166000908152600d60205260409020828155436001909101555b505050565b60008060408051908101604052600b81527f75706461746552756c653100000000000000000000000000000000000000000060208201526000908051610c27929160200190611577565b5043600681905560095490811515610c3b57fe5b06915060056000600984815481101515610c5157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161515610d4657600c6000600984815481101515610c9257fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600980549192509083908110610cc857fe5b60009182526020909120015433600160a060020a03908116911614610d415780546001901115610d075780546000190181556006546001820155610d3c565b610d33600983815481101515610d1957fe5b6000918252602082200154600160a060020a031690610a38565b610d3c82610fda565b610d46565b600381555b600354600954108015610d5c5750600a54600090115b1561051457600a8054610d9491906000198101908110610d7857fe5b600091825260209091200154600160a060020a031660036110ae565b600a546105149060001901610e7f565b4360068190556009549060009081908390811515610dbe57fe5b069150600982815481101515610dd057fe5b6000918252602082200154600160a060020a03169150821115610e6c5780600160a060020a031633600160a060020a0316141515610e6c57610e13816000610a38565b600160a060020a0381166000908152600c602052604081208181556001018190556009805484908110610e4257fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790555b60018303821415610bd857610bd8611111565b600a54600081831015610bd857600d6000600a85815481101515610e9f57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b60018203811015610f4057600a805460018301908110610ee957fe5b600091825260209091200154600a8054600160a060020a039092169183908110610f0f57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101610ecd565b600a80546000190190610f539082611625565b50505050565b600160a060020a03811615801590610f875750600160a060020a0381166000908152600e6020526040902054155b15610fd757600b805460018101610f9e8382611625565b5060009182526020808320919091018054600160a060020a031916600160a060020a0385169081179091558252600e9052604090204390555b50565b600954600081831015610bd857600c6000600985815481101515610ffa57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181208181556001015550815b6001820381101561109b57600980546001830190811061104457fe5b60009182526020909120015460098054600160a060020a03909216918390811061106a57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600101611028565b600980546000190190610f539082611625565b60035460095410156105145760098054600181016110cc8382611625565b5060009182526020808320919091018054600160a060020a038616600160a060020a031990911681179091558252600c90526040902081815543600191909101555050565b600061111b6115f5565b600080600080600061112b6115f5565b6000806000806009600081548110151561114157fe5b600091825260209091200154600954600160a060020a039091169c5060405180591061116a5750595b90808252806020026020018201604052509a50600099505b8a518a10156111dc57600980548b90811061119957fe5b600091825260209091200154600160a060020a03168b8b815181106111ba57fe5b600160a060020a03909216602092830290910190910152600190990198611182565b8a5198505b600089111561128c576001890397506111f988610fda565b8a888151811061120557fe5b9060200190602002015196506000881180156112295750600160a060020a03871615155b1561128057600160a060020a0387166000908152600d6020526040902054151561125857611258876005610a38565b600160a060020a0387166000908152600d602052604090205461128090889060001901610a38565b600019909801976111e1565b600160a060020a038c1660009081526005602052604090205460ff1680156112b75750600854600190115b1561134f57600095505b60085486101561134a578b600160a060020a03166008878154811015156112e457fe5b600091825260209091200154600160a060020a0316141561133f57600854600019018614156113255761132060086000815481101515610d7857fe5b61133a565b61133a600887600101815481101515610d7857fe5b61134a565b6001909501946112c1565b611362565b61136260086000815481101515610d7857fe5b6003546040518059106113725750595b9080825280602002602001820160405250945060009350600092505b600a5483101561142d5760035484106113a65761142d565b6113dc600a848154811015156113b857fe5b600091825260209091200154600a54600160a060020a0390911690600019016114be565b91506113e88583611525565b156113f257611422565b611404600a83815481101515610d7857fe5b8185858151811061141157fe5b602090810290910101526001909301925b60019092019161138e565b6003548410156114b0575060005b600a548110156114b057600c6000600a8381548110151561145857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154151561149857611498600a82815481101515610d7857fe5b600354600954106114a8576114b0565b60010161143b565b505050505050505050505050565b6000808083116114d1576000915061151e565b834442604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390209050828181151561151a57fe5b0691505b5092915050565b60008060008451111561156d575060005b835181101561156d578284828151811061154c57fe5b906020019060200201511415611565576001915061151e565b600101611536565b5060009392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115b857805160ff19168380011785556115e5565b828001600101855582156115e5579182015b828111156115e55782518255916020019190600101906115ca565b506115f1929150611645565b5090565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610fd79190611645565b815481835581811511610bd857600083815260209020610bd89181019083015b6108c591905b808211156115f1576000815560010161164b5600a165627a7a7230582097bcda1b95cd5de76b301978a976d8a2d462259d9dba68712c3bc6fa8f14a29b0029`

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
