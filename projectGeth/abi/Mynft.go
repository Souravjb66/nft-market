// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Mycontract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MycontractMetaData contains all meta data concerning the Mycontract contract.
var MycontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"changeTokenOwner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIdValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalanceByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintLizToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MycontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MycontractMetaData.ABI instead.
var MycontractABI = MycontractMetaData.ABI

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MycontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// GetAddressById is a free data retrieval call binding the contract method 0x8d80c922.
//
// Solidity: function getAddressById(uint256 id) view returns(address)
func (_Mycontract *MycontractCaller) GetAddressById(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getAddressById", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressById is a free data retrieval call binding the contract method 0x8d80c922.
//
// Solidity: function getAddressById(uint256 id) view returns(address)
func (_Mycontract *MycontractSession) GetAddressById(id *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetAddressById(&_Mycontract.CallOpts, id)
}

// GetAddressById is a free data retrieval call binding the contract method 0x8d80c922.
//
// Solidity: function getAddressById(uint256 id) view returns(address)
func (_Mycontract *MycontractCallerSession) GetAddressById(id *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetAddressById(&_Mycontract.CallOpts, id)
}

// GetIdValue is a free data retrieval call binding the contract method 0x8aa374e3.
//
// Solidity: function getIdValue(uint256 id) view returns(string)
func (_Mycontract *MycontractCaller) GetIdValue(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getIdValue", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIdValue is a free data retrieval call binding the contract method 0x8aa374e3.
//
// Solidity: function getIdValue(uint256 id) view returns(string)
func (_Mycontract *MycontractSession) GetIdValue(id *big.Int) (string, error) {
	return _Mycontract.Contract.GetIdValue(&_Mycontract.CallOpts, id)
}

// GetIdValue is a free data retrieval call binding the contract method 0x8aa374e3.
//
// Solidity: function getIdValue(uint256 id) view returns(string)
func (_Mycontract *MycontractCallerSession) GetIdValue(id *big.Int) (string, error) {
	return _Mycontract.Contract.GetIdValue(&_Mycontract.CallOpts, id)
}

// GetTotalBalanceByAddress is a free data retrieval call binding the contract method 0x2e1578ab.
//
// Solidity: function getTotalBalanceByAddress() view returns(uint256)
func (_Mycontract *MycontractCaller) GetTotalBalanceByAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getTotalBalanceByAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalanceByAddress is a free data retrieval call binding the contract method 0x2e1578ab.
//
// Solidity: function getTotalBalanceByAddress() view returns(uint256)
func (_Mycontract *MycontractSession) GetTotalBalanceByAddress() (*big.Int, error) {
	return _Mycontract.Contract.GetTotalBalanceByAddress(&_Mycontract.CallOpts)
}

// GetTotalBalanceByAddress is a free data retrieval call binding the contract method 0x2e1578ab.
//
// Solidity: function getTotalBalanceByAddress() view returns(uint256)
func (_Mycontract *MycontractCallerSession) GetTotalBalanceByAddress() (*big.Int, error) {
	return _Mycontract.Contract.GetTotalBalanceByAddress(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCallerSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCallerSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCallerSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Mycontract *MycontractCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Mycontract *MycontractSession) TokenCounter() (*big.Int, error) {
	return _Mycontract.Contract.TokenCounter(&_Mycontract.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Mycontract *MycontractCallerSession) TokenCounter() (*big.Int, error) {
	return _Mycontract.Contract.TokenCounter(&_Mycontract.CallOpts)
}

// ChangeTokenOwner is a paid mutator transaction binding the contract method 0xd03efc01.
//
// Solidity: function changeTokenOwner(uint256 id) returns(string)
func (_Mycontract *MycontractTransactor) ChangeTokenOwner(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "changeTokenOwner", id)
}

// ChangeTokenOwner is a paid mutator transaction binding the contract method 0xd03efc01.
//
// Solidity: function changeTokenOwner(uint256 id) returns(string)
func (_Mycontract *MycontractSession) ChangeTokenOwner(id *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.ChangeTokenOwner(&_Mycontract.TransactOpts, id)
}

// ChangeTokenOwner is a paid mutator transaction binding the contract method 0xd03efc01.
//
// Solidity: function changeTokenOwner(uint256 id) returns(string)
func (_Mycontract *MycontractTransactorSession) ChangeTokenOwner(id *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.ChangeTokenOwner(&_Mycontract.TransactOpts, id)
}

// MintLizToken is a paid mutator transaction binding the contract method 0x8c2a73b4.
//
// Solidity: function mintLizToken(string tokenUri) returns(uint256)
func (_Mycontract *MycontractTransactor) MintLizToken(opts *bind.TransactOpts, tokenUri string) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "mintLizToken", tokenUri)
}

// MintLizToken is a paid mutator transaction binding the contract method 0x8c2a73b4.
//
// Solidity: function mintLizToken(string tokenUri) returns(uint256)
func (_Mycontract *MycontractSession) MintLizToken(tokenUri string) (*types.Transaction, error) {
	return _Mycontract.Contract.MintLizToken(&_Mycontract.TransactOpts, tokenUri)
}

// MintLizToken is a paid mutator transaction binding the contract method 0x8c2a73b4.
//
// Solidity: function mintLizToken(string tokenUri) returns(uint256)
func (_Mycontract *MycontractTransactorSession) MintLizToken(tokenUri string) (*types.Transaction, error) {
	return _Mycontract.Contract.MintLizToken(&_Mycontract.TransactOpts, tokenUri)
}
