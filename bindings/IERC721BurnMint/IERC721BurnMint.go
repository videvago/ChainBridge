// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC721BurnMint

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

// IERC721BurnMintABI is the input ABI used to generate the binding from.
const IERC721BurnMintABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721BurnMint is an auto generated Go binding around an Ethereum contract.
type IERC721BurnMint struct {
	IERC721BurnMintCaller     // Read-only binding to the contract
	IERC721BurnMintTransactor // Write-only binding to the contract
	IERC721BurnMintFilterer   // Log filterer for contract events
}

// IERC721BurnMintCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721BurnMintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BurnMintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721BurnMintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BurnMintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721BurnMintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BurnMintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721BurnMintSession struct {
	Contract     *IERC721BurnMint  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721BurnMintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721BurnMintCallerSession struct {
	Contract *IERC721BurnMintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721BurnMintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721BurnMintTransactorSession struct {
	Contract     *IERC721BurnMintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721BurnMintRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721BurnMintRaw struct {
	Contract *IERC721BurnMint // Generic contract binding to access the raw methods on
}

// IERC721BurnMintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721BurnMintCallerRaw struct {
	Contract *IERC721BurnMintCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721BurnMintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721BurnMintTransactorRaw struct {
	Contract *IERC721BurnMintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721BurnMint creates a new instance of IERC721BurnMint, bound to a specific deployed contract.
func NewIERC721BurnMint(address common.Address, backend bind.ContractBackend) (*IERC721BurnMint, error) {
	contract, err := bindIERC721BurnMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMint{IERC721BurnMintCaller: IERC721BurnMintCaller{contract: contract}, IERC721BurnMintTransactor: IERC721BurnMintTransactor{contract: contract}, IERC721BurnMintFilterer: IERC721BurnMintFilterer{contract: contract}}, nil
}

// NewIERC721BurnMintCaller creates a new read-only instance of IERC721BurnMint, bound to a specific deployed contract.
func NewIERC721BurnMintCaller(address common.Address, caller bind.ContractCaller) (*IERC721BurnMintCaller, error) {
	contract, err := bindIERC721BurnMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintCaller{contract: contract}, nil
}

// NewIERC721BurnMintTransactor creates a new write-only instance of IERC721BurnMint, bound to a specific deployed contract.
func NewIERC721BurnMintTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721BurnMintTransactor, error) {
	contract, err := bindIERC721BurnMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintTransactor{contract: contract}, nil
}

// NewIERC721BurnMintFilterer creates a new log filterer instance of IERC721BurnMint, bound to a specific deployed contract.
func NewIERC721BurnMintFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721BurnMintFilterer, error) {
	contract, err := bindIERC721BurnMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintFilterer{contract: contract}, nil
}

// bindIERC721BurnMint binds a generic wrapper to an already deployed contract.
func bindIERC721BurnMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721BurnMintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721BurnMint *IERC721BurnMintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721BurnMint.Contract.IERC721BurnMintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721BurnMint *IERC721BurnMintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.IERC721BurnMintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721BurnMint *IERC721BurnMintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.IERC721BurnMintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721BurnMint *IERC721BurnMintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721BurnMint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721BurnMint *IERC721BurnMintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721BurnMint *IERC721BurnMintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721BurnMint *IERC721BurnMintCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721BurnMint.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721BurnMint *IERC721BurnMintSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721BurnMint.Contract.BalanceOf(&_IERC721BurnMint.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721BurnMint *IERC721BurnMintCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721BurnMint.Contract.BalanceOf(&_IERC721BurnMint.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721BurnMint *IERC721BurnMintCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721BurnMint.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721BurnMint *IERC721BurnMintSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721BurnMint.Contract.GetApproved(&_IERC721BurnMint.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721BurnMint *IERC721BurnMintCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721BurnMint.Contract.GetApproved(&_IERC721BurnMint.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721BurnMint.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721BurnMint.Contract.IsApprovedForAll(&_IERC721BurnMint.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721BurnMint.Contract.IsApprovedForAll(&_IERC721BurnMint.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721BurnMint *IERC721BurnMintCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721BurnMint.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721BurnMint *IERC721BurnMintSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721BurnMint.Contract.OwnerOf(&_IERC721BurnMint.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721BurnMint *IERC721BurnMintCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721BurnMint.Contract.OwnerOf(&_IERC721BurnMint.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721BurnMint.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721BurnMint.Contract.SupportsInterface(&_IERC721BurnMint.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721BurnMint *IERC721BurnMintCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721BurnMint.Contract.SupportsInterface(&_IERC721BurnMint.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Approve(&_IERC721BurnMint.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Approve(&_IERC721BurnMint.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Burn(&_IERC721BurnMint.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Burn(&_IERC721BurnMint.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address recipient, uint256 tokenID, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "mint", recipient, tokenID, data)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address recipient, uint256 tokenID, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) Mint(recipient common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Mint(&_IERC721BurnMint.TransactOpts, recipient, tokenID, data)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address recipient, uint256 tokenID, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) Mint(recipient common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.Mint(&_IERC721BurnMint.TransactOpts, recipient, tokenID, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SafeTransferFrom(&_IERC721BurnMint.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SafeTransferFrom(&_IERC721BurnMint.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SafeTransferFrom0(&_IERC721BurnMint.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SafeTransferFrom0(&_IERC721BurnMint.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SetApprovalForAll(&_IERC721BurnMint.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.SetApprovalForAll(&_IERC721BurnMint.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.TransferFrom(&_IERC721BurnMint.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721BurnMint *IERC721BurnMintTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721BurnMint.Contract.TransferFrom(&_IERC721BurnMint.TransactOpts, from, to, tokenId)
}

// IERC721BurnMintApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721BurnMint contract.
type IERC721BurnMintApprovalIterator struct {
	Event *IERC721BurnMintApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC721BurnMintApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BurnMintApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC721BurnMintApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC721BurnMintApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BurnMintApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BurnMintApproval represents a Approval event raised by the IERC721BurnMint contract.
type IERC721BurnMintApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721BurnMintApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintApprovalIterator{contract: _IERC721BurnMint.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721BurnMintApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BurnMintApproval)
				if err := _IERC721BurnMint.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) ParseApproval(log types.Log) (*IERC721BurnMintApproval, error) {
	event := new(IERC721BurnMintApproval)
	if err := _IERC721BurnMint.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BurnMintApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721BurnMint contract.
type IERC721BurnMintApprovalForAllIterator struct {
	Event *IERC721BurnMintApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC721BurnMintApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BurnMintApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC721BurnMintApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC721BurnMintApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BurnMintApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BurnMintApprovalForAll represents a ApprovalForAll event raised by the IERC721BurnMint contract.
type IERC721BurnMintApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721BurnMint *IERC721BurnMintFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721BurnMintApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintApprovalForAllIterator{contract: _IERC721BurnMint.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721BurnMint *IERC721BurnMintFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721BurnMintApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BurnMintApprovalForAll)
				if err := _IERC721BurnMint.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721BurnMint *IERC721BurnMintFilterer) ParseApprovalForAll(log types.Log) (*IERC721BurnMintApprovalForAll, error) {
	event := new(IERC721BurnMintApprovalForAll)
	if err := _IERC721BurnMint.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BurnMintTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721BurnMint contract.
type IERC721BurnMintTransferIterator struct {
	Event *IERC721BurnMintTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC721BurnMintTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BurnMintTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC721BurnMintTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC721BurnMintTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BurnMintTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BurnMintTransfer represents a Transfer event raised by the IERC721BurnMint contract.
type IERC721BurnMintTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721BurnMintTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BurnMintTransferIterator{contract: _IERC721BurnMint.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721BurnMintTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721BurnMint.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BurnMintTransfer)
				if err := _IERC721BurnMint.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721BurnMint *IERC721BurnMintFilterer) ParseTransfer(log types.Log) (*IERC721BurnMintTransfer, error) {
	event := new(IERC721BurnMintTransfer)
	if err := _IERC721BurnMint.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
