// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeHandler

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

// BridgeHandlerABI is the input ABI used to generate the binding from.
const BridgeHandlerABI = "[{\"inputs\":[],\"name\":\"STATUS_BURN_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnableResouceIDs\",\"type\":\"bytes32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeHandler is an auto generated Go binding around an Ethereum contract.
type BridgeHandler struct {
	BridgeHandlerCaller     // Read-only binding to the contract
	BridgeHandlerTransactor // Write-only binding to the contract
	BridgeHandlerFilterer   // Log filterer for contract events
}

// BridgeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeHandlerSession struct {
	Contract     *BridgeHandler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeHandlerCallerSession struct {
	Contract *BridgeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeHandlerTransactorSession struct {
	Contract     *BridgeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeHandlerRaw struct {
	Contract *BridgeHandler // Generic contract binding to access the raw methods on
}

// BridgeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeHandlerCallerRaw struct {
	Contract *BridgeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeHandlerTransactorRaw struct {
	Contract *BridgeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeHandler creates a new instance of BridgeHandler, bound to a specific deployed contract.
func NewBridgeHandler(address common.Address, backend bind.ContractBackend) (*BridgeHandler, error) {
	contract, err := bindBridgeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeHandler{BridgeHandlerCaller: BridgeHandlerCaller{contract: contract}, BridgeHandlerTransactor: BridgeHandlerTransactor{contract: contract}, BridgeHandlerFilterer: BridgeHandlerFilterer{contract: contract}}, nil
}

// NewBridgeHandlerCaller creates a new read-only instance of BridgeHandler, bound to a specific deployed contract.
func NewBridgeHandlerCaller(address common.Address, caller bind.ContractCaller) (*BridgeHandlerCaller, error) {
	contract, err := bindBridgeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeHandlerCaller{contract: contract}, nil
}

// NewBridgeHandlerTransactor creates a new write-only instance of BridgeHandler, bound to a specific deployed contract.
func NewBridgeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeHandlerTransactor, error) {
	contract, err := bindBridgeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeHandlerTransactor{contract: contract}, nil
}

// NewBridgeHandlerFilterer creates a new log filterer instance of BridgeHandler, bound to a specific deployed contract.
func NewBridgeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeHandlerFilterer, error) {
	contract, err := bindBridgeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeHandlerFilterer{contract: contract}, nil
}

// bindBridgeHandler binds a generic wrapper to an already deployed contract.
func bindBridgeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeHandler *BridgeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeHandler.Contract.BridgeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeHandler *BridgeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHandler.Contract.BridgeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeHandler *BridgeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeHandler.Contract.BridgeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeHandler *BridgeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeHandler *BridgeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeHandler *BridgeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeHandler.Contract.contract.Transact(opts, method, params...)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_BridgeHandler *BridgeHandlerCaller) STATUSBURNMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeHandler.contract.Call(opts, &out, "STATUS_BURN_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_BridgeHandler *BridgeHandlerSession) STATUSBURNMINT() (*big.Int, error) {
	return _BridgeHandler.Contract.STATUSBURNMINT(&_BridgeHandler.CallOpts)
}

// STATUSBURNMINT is a free data retrieval call binding the contract method 0xe347ebcd.
//
// Solidity: function STATUS_BURN_MINT() view returns(uint256)
func (_BridgeHandler *BridgeHandlerCallerSession) STATUSBURNMINT() (*big.Int, error) {
	return _BridgeHandler.Contract.STATUSBURNMINT(&_BridgeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BridgeHandler *BridgeHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BridgeHandler *BridgeHandlerSession) BridgeAddress() (common.Address, error) {
	return _BridgeHandler.Contract.BridgeAddress(&_BridgeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BridgeHandler *BridgeHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _BridgeHandler.Contract.BridgeAddress(&_BridgeHandler.CallOpts)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_BridgeHandler *BridgeHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_BridgeHandler *BridgeHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _BridgeHandler.Contract.ContractAddressToResourceID(&_BridgeHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_BridgeHandler *BridgeHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _BridgeHandler.Contract.ContractAddressToResourceID(&_BridgeHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_BridgeHandler *BridgeHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BridgeHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_BridgeHandler *BridgeHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _BridgeHandler.Contract.ResourceIDToContractAddress(&_BridgeHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_BridgeHandler *BridgeHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _BridgeHandler.Contract.ResourceIDToContractAddress(&_BridgeHandler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_BridgeHandler *BridgeHandlerCaller) Status(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BridgeHandler.contract.Call(opts, &out, "_status", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_BridgeHandler *BridgeHandlerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _BridgeHandler.Contract.Status(&_BridgeHandler.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x85e1fd34.
//
// Solidity: function _status(bytes32 ) view returns(uint256)
func (_BridgeHandler *BridgeHandlerCallerSession) Status(arg0 [32]byte) (*big.Int, error) {
	return _BridgeHandler.Contract.Status(&_BridgeHandler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_BridgeHandler *BridgeHandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Deposit(&_BridgeHandler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Deposit(&_BridgeHandler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.ExecuteProposal(&_BridgeHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.ExecuteProposal(&_BridgeHandler.TransactOpts, resourceID, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x6b392b20.
//
// Solidity: function initialize(address bridgeAddress, bytes32[] initialResourceIDs, address[] initialContractAddresses, bytes32[] burnableResouceIDs) returns()
func (_BridgeHandler *BridgeHandlerTransactor) Initialize(opts *bind.TransactOpts, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "initialize", bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
}

// Initialize is a paid mutator transaction binding the contract method 0x6b392b20.
//
// Solidity: function initialize(address bridgeAddress, bytes32[] initialResourceIDs, address[] initialContractAddresses, bytes32[] burnableResouceIDs) returns()
func (_BridgeHandler *BridgeHandlerSession) Initialize(bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Initialize(&_BridgeHandler.TransactOpts, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
}

// Initialize is a paid mutator transaction binding the contract method 0x6b392b20.
//
// Solidity: function initialize(address bridgeAddress, bytes32[] initialResourceIDs, address[] initialContractAddresses, bytes32[] burnableResouceIDs) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) Initialize(bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableResouceIDs [][32]byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Initialize(&_BridgeHandler.TransactOpts, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableResouceIDs)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactor) Release(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "release", resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Release(&_BridgeHandler.TransactOpts, resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.Release(&_BridgeHandler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_BridgeHandler *BridgeHandlerTransactor) SetBurnable(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "setBurnable", resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_BridgeHandler *BridgeHandlerSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.SetBurnable(&_BridgeHandler.TransactOpts, resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _BridgeHandler.Contract.SetBurnable(&_BridgeHandler.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_BridgeHandler *BridgeHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_BridgeHandler *BridgeHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeHandler.Contract.SetResource(&_BridgeHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_BridgeHandler *BridgeHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeHandler.Contract.SetResource(&_BridgeHandler.TransactOpts, resourceID, contractAddress)
}
