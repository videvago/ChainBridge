// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnableResouceIDs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"STATUS_BURN_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b50604051620010b3380380620010b3833981810160405260808110156200003757600080fd5b8101908080519060200190929190805160405193929190846401000000008211156200006257600080fd5b838201915060208201858111156200007957600080fd5b82518660208202830111640100000000821117156200009757600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000d0578082015181840152602081019050620000b3565b5050505090500160405260200180516040519392919084640100000000821115620000fa57600080fd5b838201915060208201858111156200011157600080fd5b82518660208202830111640100000000821117156200012f57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620001685780820151818401526020810190506200014b565b50505050905001604052602001805160405193929190846401000000008211156200019257600080fd5b83820191506020820185811115620001a957600080fd5b8251866020820283011164010000000082111715620001c757600080fd5b8083526020830192505050908051906020019060200280838360005b8381101562000200578082015181840152602081019050620001e3565b5050505090500160405250505083838383836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508151835114620002c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4c656e677468206d69736d61746368000000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b8351811015620003215762000313848281518110620002ea57fe5b6020026020010151848381518110620002ff57fe5b60200260200101516200037460201b60201c565b8080600101915050620002cf565b5060008090505b81518110156200036557620003578282815181106200034357fe5b60200260200101516200042560201b60201c565b808060010191505062000328565b5050505050505050506200044a565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060036000838152602001908152602001600020600090555050565b6001600360008381526020019081526020016000206000828254179250508190555050565b610c59806200045a6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b8fa373611610066578063b8fa373614610283578063c54c2a11146102d1578063e248cff21461033f578063e347ebcd146103c2578063ec97d3b4146103e05761009e565b80631bc8a466146100a3578063318c136e1461012657806385e1fd341461017057806389aa2453146101b2578063b07e54bb146101e0575b600080fd5b610124600480360360408110156100b957600080fd5b8101908080359060200190929190803590602001906401000000008111156100e057600080fd5b8201836020820111156100f257600080fd5b8035906020019184600183028401116401000000008311171561011457600080fd5b9091929391929390505050610438565b005b61012e6105db565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61019c6004803603602081101561018657600080fd5b8101908080359060200190929190505050610600565b6040518082815260200191505060405180910390f35b6101de600480360360208110156101c857600080fd5b8101908080359060200190929190505050610618565b005b610281600480360360608110156101f657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561023d57600080fd5b82018360208201111561024f57600080fd5b8035906020019184600183028401116401000000008311171561027157600080fd5b90919293919293905050506106e6565b005b6102cf6004803603604081101561029957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061088a565b005b6102fd600480360360208110156102e757600080fd5b810190808035906020019092919050505061095a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103c06004803603604081101561035557600080fd5b81019080803590602001909291908035906020019064010000000081111561037c57600080fd5b82018360208201111561038e57600080fd5b803590602001918460018302840111640100000000831117156103b057600080fd5b909192939192939050505061098d565b005b6103ca610b30565b6040518082815260200191505060405180910390f35b610422600480360360208110156103f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b35565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156105d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b6106e381610b4d565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610883576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461094c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b6109568282610b72565b5050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f427269646765206f6e6c7900000000000000000000000000000000000000000081525060200191505060405180910390fd5b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b2a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636f757263654944206e6f74206d6170706564000000000000000000000081525060200191505060405180910390fd5b50505050565b600181565b60026020528060005260406000206000915090505481565b6001600360008381526020019081526020016000206000828254179250508190555050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506003600083815260200190815260200160002060009055505056fea26469706673582212203478f07e9cb7b2f74b1e6f3c3988c79200326a3d504572c3e3da0ec3ef462d6264736f6c63430006040033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, _bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_GenericHandler *GenericHandlerCaller) STATUSBURNMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "STATUS_BURN_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_GenericHandler *GenericHandlerSession) STATUSBURNMINT() (*big.Int, error) {
	return _GenericHandler.Contract.STATUSBURNMINT(&_GenericHandler.CallOpts)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) STATUSBURNMINT() (*big.Int, error) {
	return _GenericHandler.Contract.STATUSBURNMINT(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_GenericHandler *GenericHandlerCaller) Status(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_status", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_GenericHandler *GenericHandlerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _GenericHandler.Contract.Status(&_GenericHandler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _GenericHandler.Contract.Status(&_GenericHandler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address , bytes ) returns()
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", resourceID, arg1, arg2)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address , bytes ) returns()
func (_GenericHandler *GenericHandlerSession) Deposit(resourceID [32]byte, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, arg1, arg2)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address , bytes ) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(resourceID [32]byte, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, arg1, arg2)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeProposal", resourceID, arg1)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerSession) ExecuteProposal(resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, arg1)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, arg1)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerTransactor) Release(opts *bind.TransactOpts, resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "release", resourceID, arg1)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerSession) Release(resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Release(&_GenericHandler.TransactOpts, resourceID, arg1)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes ) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Release(resourceID [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Release(&_GenericHandler.TransactOpts, resourceID, arg1)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_GenericHandler *GenericHandlerTransactor) SetBurnable(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setBurnable", resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_GenericHandler *GenericHandlerSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetBurnable(&_GenericHandler.TransactOpts, resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetBurnable(&_GenericHandler.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress)
}
