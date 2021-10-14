// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERC721HandlerABI is the input ABI used to generate the binding from.
const ERC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnableResouceIDs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"STATUS_BURN_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721HandlerBin = "0x60806040523480156200001157600080fd5b506040516200183238038062001832833981810160405260808110156200003757600080fd5b8101908080519060200190929190805160405193929190846401000000008211156200006257600080fd5b838201915060208201858111156200007957600080fd5b82518660208202830111640100000000821117156200009757600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000d0578082015181840152602081019050620000b3565b5050505090500160405260200180516040519392919084640100000000821115620000fa57600080fd5b838201915060208201858111156200011157600080fd5b82518660208202830111640100000000821117156200012f57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620001685780820151818401526020810190506200014b565b50505050905001604052602001805160405193929190846401000000008211156200019257600080fd5b83820191506020820185811115620001a957600080fd5b8251866020820283011164010000000082111715620001c757600080fd5b8083526020830192505050908051906020019060200280838360005b8381101562000200578082015181840152602081019050620001e3565b5050505090500160405250505083838383836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508151835114620002c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4c656e677468206d69736d61746368000000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b8351811015620003215762000313848281518110620002ea57fe5b6020026020010151848381518110620002ff57fe5b60200260200101516200037460201b60201c565b8080600101915050620002cf565b5060008090505b81518110156200036557620003578282815181106200034357fe5b60200260200101516200042560201b60201c565b808060010191505062000328565b5050505050505050506200044a565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060036000838152602001908152602001600020600090555050565b6001600360008381526020019081526020016000206000828254179250508190555050565b6113d8806200045a6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063b07e54bb11610071578063b07e54bb14610259578063b8fa3736146102fc578063c54c2a111461034a578063e248cff2146103b8578063e347ebcd1461043b578063ec97d3b414610459576100a9565b80631bc8a466146100ae578063318c136e14610131578063735429801461017b57806385e1fd34146101e957806389aa24531461022b575b600080fd5b61012f600480360360408110156100c457600080fd5b8101908080359060200190929190803590602001906401000000008111156100eb57600080fd5b8201836020820111156100fd57600080fd5b8035906020019184600183028401116401000000008311171561011f57600080fd5b90919293919293905050506104b1565b005b61013961072f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e76004803603606081101561019157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610754565b005b610215600480360360208110156101ff57600080fd5b8101908080359060200190929190505050610832565b6040518082815260200191505060405180910390f35b6102576004803603602081101561024157600080fd5b810190808035906020019092919050505061084a565b005b6102fa6004803603606081101561026f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156102b657600080fd5b8201836020820111156102c857600080fd5b803590602001918460018302840111640100000000831117156102ea57600080fd5b9091929391929390505050610918565b005b6103486004803603604081101561031257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b49565b005b6103766004803603602081101561036057600080fd5b8101908080359060200190929190505050610c19565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610439600480360360408110156103ce57600080fd5b8101908080359060200190929190803590602001906401000000008111156103f557600080fd5b82018360208201111561040757600080fd5b8035906020019184600183028401116401000000008311171561042957600080fd5b9091929391929390505050610c4c565b005b610443610f70565b6040518082815260200191505060405180910390f35b61049b6004803603602081101561046f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f75565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561064e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b60008060008585606081101561066357600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050509250925092506014821461071a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f496e76616c69642061646472657373206c656e6774680000000000000000000081525060200191505060405180910390fd5b61072684308386610f8d565b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561081457600080fd5b505af1158015610828573d6000803e3d6000fd5b5050505050505050565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461090c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b6109158161106c565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000828260006020808211156109ef57600080fd5b828111156109fc57600080fd5b6001820284019350818103925050506020811015610a1957600080fd5b8101908080359060200190929190505050905060006001600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b6000600160036000898152602001908152602001600020541614610b3457610b2f8183611091565b610b41565b610b4081863085611100565b5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b610c1582826111df565b5050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610de9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b600080600085856060811015610dfe57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505092509250925060148214610eb5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f496e76616c69642061646472657373206c656e6774680000000000000000000081525060200191505060405180910390fd5b60006001600360008a8152602001908152602001600020541614610f5a576060868660548180821115610ee757600080fd5b82811115610ef457600080fd5b6001820284019350818103925050508080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050610f5485838684611290565b50610f67565b610f6684308386610f8d565b5b50505050505050565b600181565b60026020528060005260406000206000915090505481565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561104d57600080fd5b505af1158015611061573d6000803e3d6000fd5b505050505050505050565b6001600360008381526020019081526020016000206000828254179250508190555050565b8173ffffffffffffffffffffffffffffffffffffffff166342966c68826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156110e457600080fd5b505af11580156110f8573d6000803e3d6000fd5b505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156111c057600080fd5b505af11580156111d4573d6000803e3d6000fd5b505050505050505050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060036000838152602001908152602001600020600090555050565b8373ffffffffffffffffffffffffffffffffffffffff166394d008ef8484846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561133657808201518184015260208101905061131b565b50505050905090810190601f1680156113635780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561138457600080fd5b505af1158015611398573d6000803e3d6000fd5b505050505050505056fea2646970667358221220fba2aa7cb39189070dc12548f0234bed7a2323f009cc16e1304382a70fb7726064736f6c63430006040033"

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721HandlerBin), backend, _bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC721Handler *ERC721HandlerCaller) STATUSBURNMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "STATUS_BURN_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC721Handler *ERC721HandlerSession) STATUSBURNMINT() (*big.Int, error) {
	return _ERC721Handler.Contract.STATUSBURNMINT(&_ERC721Handler.CallOpts)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_ERC721Handler *ERC721HandlerCallerSession) STATUSBURNMINT() (*big.Int, error) {
	return _ERC721Handler.Contract.STATUSBURNMINT(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.ContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.ContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC721Handler *ERC721HandlerCaller) Status(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_status", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC721Handler *ERC721HandlerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _ERC721Handler.Contract.Status(&_ERC721Handler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_ERC721Handler *ERC721HandlerCallerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _ERC721Handler.Contract.Status(&_ERC721Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) FundERC721(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "fundERC721", tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Release(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "release", resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Release(&_ERC721Handler.TransactOpts, resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Release(&_ERC721Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setBurnable", resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC721Handler *ERC721HandlerSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}
