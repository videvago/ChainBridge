// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnableResouceIDs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"STATUS_BURN_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b50604051620018b2380380620018b2833981810160405260808110156200003757600080fd5b8101908080519060200190929190805160405193929190846401000000008211156200006257600080fd5b838201915060208201858111156200007957600080fd5b82518660208202830111640100000000821117156200009757600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000d0578082015181840152602081019050620000b3565b5050505090500160405260200180516040519392919084640100000000821115620000fa57600080fd5b838201915060208201858111156200011157600080fd5b82518660208202830111640100000000821117156200012f57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620001685780820151818401526020810190506200014b565b50505050905001604052602001805160405193929190846401000000008211156200019257600080fd5b83820191506020820185811115620001a957600080fd5b8251866020820283011164010000000082111715620001c757600080fd5b8083526020830192505050908051906020019060200280838360005b8381101562000200578082015181840152602081019050620001e3565b5050505090500160405250505083838383836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508151835114620002c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4c656e677468206d69736d61746368000000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b8351811015620003215762000313848281518110620002ea57fe5b6020026020010151848381518110620002ff57fe5b60200260200101516200037460201b60201c565b8080600101915050620002cf565b5060008090505b81518110156200036557620003578282815181106200034357fe5b60200260200101516200042560201b60201c565b808060010191505062000328565b5050505050505050506200044a565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060036000838152602001908152602001600020600090555050565b6001600360008381526020019081526020016000206000828254179250508190555050565b611458806200045a6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063b07e54bb11610071578063b07e54bb14610259578063b8fa3736146102fc578063c54c2a111461034a578063e248cff2146103b8578063e347ebcd1461043b578063ec97d3b414610459576100a9565b80631bc8a466146100ae578063318c136e1461013157806385e1fd341461017b57806389aa2453146101bd57806395601f09146101eb575b600080fd5b61012f600480360360408110156100c457600080fd5b8101908080359060200190929190803590602001906401000000008111156100eb57600080fd5b8201836020820111156100fd57600080fd5b8035906020019184600183028401116401000000008311171561011f57600080fd5b90919293919293905050506104b1565b005b61013961072e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a76004803603602081101561019157600080fd5b8101908080359060200190929190505050610753565b6040518082815260200191505060405180910390f35b6101e9600480360360208110156101d357600080fd5b810190808035906020019092919050505061076b565b005b6102576004803603606081101561020157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610839565b005b6102fa6004803603606081101561026f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156102b657600080fd5b8201836020820111156102c857600080fd5b803590602001918460018302840111640100000000831117156102ea57600080fd5b9091929391929390505050610850565b005b6103486004803603604081101561031257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a82565b005b6103766004803603602081101561036057600080fd5b8101908080359060200190929190505050610b52565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610439600480360360408110156103ce57600080fd5b8101908080359060200190929190803590602001906401000000008111156103f557600080fd5b82018360208201111561040757600080fd5b8035906020019184600183028401116401000000008311171561042957600080fd5b9091929391929390505050610b85565b005b610443610e31565b6040518082815260200191505060405180910390f35b61049b6004803603602081101561046f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e36565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561064e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b60008060008585606081101561066357600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050509250925092506014821461071a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f496e76616c696420726563697069656e74206c656e677468000000000000000081525060200191505060405180910390fd5b610725848285610e4e565b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b61083681610e64565b50565b600083905061084a81843085610e89565b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610912576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008282600060208082111561092757600080fd5b8281111561093457600080fd5b600182028401935081810392505050602081101561095157600080fd5b8101908080359060200190929190505050905060006001600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a3f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b6000600160036000898152602001908152602001600020541614610a6d57610a68818684610f76565b610a7a565b610a7981863085611020565b5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b610b4e8282611038565b5050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d22576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b600080600085856060811015610d3757600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505092509250925060148214610dee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f496e76616c696420726563697069656e74206c656e677468000000000000000081525060200191505060405180910390fd5b60006001600360008a8152602001908152602001600020541614610e1c57610e178482856110e9565b610e28565b610e27848285610e4e565b5b50505050505050565b600181565b60026020528060005260406000206000915090505481565b6000839050610e5e818484611193565b50505050565b6001600360008381526020019081526020016000206000828254179250508190555050565b610f70846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061124b565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561100257600080fd5b505af1158015611016573d6000803e3d6000fd5b5050505050505050565b600084905061103181858585610e89565b5050505050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060036000838152602001908152602001600020600090555050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561117557600080fd5b505af1158015611189573d6000803e3d6000fd5b5050505050505050565b6112468363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061124b565b505050565b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b6020831061129a5780518252602082019150602081019050602083039250611277565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146112fc576040519150601f19603f3d011682016040523d82523d6000602084013e611301565b606091505b509150915081611379576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f45524332303a2063616c6c206661696c6564000000000000000000000000000081525060200191505060405180910390fd5b60008151111561141c5780806020019051602081101561139857600080fd5b810190808051906020019092919050505061141b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f45524332303a206f7065726174696f6e20646964206e6f74207375636365656481525060200191505060405180910390fd5b5b5050505056fea2646970667358221220232913bc4130906c3c050d8d32a10cfe52718852bb9d89595ff2c1c2838ad40f64736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, _bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) STATUSBURNMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "STATUS_BURN_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) STATUSBURNMINT() (*big.Int, error) {
	return _ERC20Handler.Contract.STATUSBURNMINT(&_ERC20Handler.CallOpts)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) STATUSBURNMINT() (*big.Int, error) {
	return _ERC20Handler.Contract.STATUSBURNMINT(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.ContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.ContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) Status(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_status", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _ERC20Handler.Contract.Status(&_ERC20Handler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _ERC20Handler.Contract.Status(&_ERC20Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Release(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "release", resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Release(&_ERC20Handler.TransactOpts, resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Release(&_ERC20Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}
