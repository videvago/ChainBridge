// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBridgeHandler

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

// IBridgeHandlerABI is the input ABI used to generate the binding from.
const IBridgeHandlerABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBridgeHandler is an auto generated Go binding around an Ethereum contract.
type IBridgeHandler struct {
	IBridgeHandlerCaller     // Read-only binding to the contract
	IBridgeHandlerTransactor // Write-only binding to the contract
	IBridgeHandlerFilterer   // Log filterer for contract events
}

// IBridgeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeHandlerSession struct {
	Contract     *IBridgeHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeHandlerCallerSession struct {
	Contract *IBridgeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IBridgeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeHandlerTransactorSession struct {
	Contract     *IBridgeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IBridgeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeHandlerRaw struct {
	Contract *IBridgeHandler // Generic contract binding to access the raw methods on
}

// IBridgeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeHandlerCallerRaw struct {
	Contract *IBridgeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeHandlerTransactorRaw struct {
	Contract *IBridgeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeHandler creates a new instance of IBridgeHandler, bound to a specific deployed contract.
func NewIBridgeHandler(address common.Address, backend bind.ContractBackend) (*IBridgeHandler, error) {
	contract, err := bindIBridgeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeHandler{IBridgeHandlerCaller: IBridgeHandlerCaller{contract: contract}, IBridgeHandlerTransactor: IBridgeHandlerTransactor{contract: contract}, IBridgeHandlerFilterer: IBridgeHandlerFilterer{contract: contract}}, nil
}

// NewIBridgeHandlerCaller creates a new read-only instance of IBridgeHandler, bound to a specific deployed contract.
func NewIBridgeHandlerCaller(address common.Address, caller bind.ContractCaller) (*IBridgeHandlerCaller, error) {
	contract, err := bindIBridgeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeHandlerCaller{contract: contract}, nil
}

// NewIBridgeHandlerTransactor creates a new write-only instance of IBridgeHandler, bound to a specific deployed contract.
func NewIBridgeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeHandlerTransactor, error) {
	contract, err := bindIBridgeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeHandlerTransactor{contract: contract}, nil
}

// NewIBridgeHandlerFilterer creates a new log filterer instance of IBridgeHandler, bound to a specific deployed contract.
func NewIBridgeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeHandlerFilterer, error) {
	contract, err := bindIBridgeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeHandlerFilterer{contract: contract}, nil
}

// bindIBridgeHandler binds a generic wrapper to an already deployed contract.
func bindIBridgeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeHandler *IBridgeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeHandler.Contract.IBridgeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeHandler *IBridgeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.IBridgeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeHandler *IBridgeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.IBridgeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeHandler *IBridgeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeHandler *IBridgeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeHandler *IBridgeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.Deposit(&_IBridgeHandler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.Deposit(&_IBridgeHandler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.ExecuteProposal(&_IBridgeHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.ExecuteProposal(&_IBridgeHandler.TransactOpts, resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactor) Release(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.contract.Transact(opts, "release", resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.Release(&_IBridgeHandler.TransactOpts, resourceID, data)
}

// Release is a paid mutator transaction binding the contract method 0x1bc8a466.
//
// Solidity: function release(bytes32 resourceID, bytes data) returns()
func (_IBridgeHandler *IBridgeHandlerTransactorSession) Release(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.Release(&_IBridgeHandler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_IBridgeHandler *IBridgeHandlerTransactor) SetBurnable(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _IBridgeHandler.contract.Transact(opts, "setBurnable", resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_IBridgeHandler *IBridgeHandlerSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.SetBurnable(&_IBridgeHandler.TransactOpts, resourceID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x89aa2453.
//
// Solidity: function setBurnable(bytes32 resourceID) returns()
func (_IBridgeHandler *IBridgeHandlerTransactorSession) SetBurnable(resourceID [32]byte) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.SetBurnable(&_IBridgeHandler.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_IBridgeHandler *IBridgeHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _IBridgeHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_IBridgeHandler *IBridgeHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.SetResource(&_IBridgeHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_IBridgeHandler *IBridgeHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _IBridgeHandler.Contract.SetResource(&_IBridgeHandler.TransactOpts, resourceID, contractAddress)
}
