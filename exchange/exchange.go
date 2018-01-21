// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ExchangeInterfaceABI is the input ABI used to generate the binding from.
const ExchangeInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"canTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Traded\",\"type\":\"event\"}]"

// ExchangeInterfaceBin is the compiled bytecode used for deploying new contracts.
const ExchangeInterfaceBin = `0x`

// DeployExchangeInterface deploys a new Ethereum contract, binding an instance of ExchangeInterface to it.
func DeployExchangeInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExchangeInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExchangeInterface{ExchangeInterfaceCaller: ExchangeInterfaceCaller{contract: contract}, ExchangeInterfaceTransactor: ExchangeInterfaceTransactor{contract: contract}}, nil
}

// ExchangeInterface is an auto generated Go binding around an Ethereum contract.
type ExchangeInterface struct {
	ExchangeInterfaceCaller     // Read-only binding to the contract
	ExchangeInterfaceTransactor // Write-only binding to the contract
}

// ExchangeInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeInterfaceSession struct {
	Contract     *ExchangeInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExchangeInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeInterfaceCallerSession struct {
	Contract *ExchangeInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExchangeInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeInterfaceTransactorSession struct {
	Contract     *ExchangeInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExchangeInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeInterfaceRaw struct {
	Contract *ExchangeInterface // Generic contract binding to access the raw methods on
}

// ExchangeInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeInterfaceCallerRaw struct {
	Contract *ExchangeInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeInterfaceTransactorRaw struct {
	Contract *ExchangeInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeInterface creates a new instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterface(address common.Address, backend bind.ContractBackend) (*ExchangeInterface, error) {
	contract, err := bindExchangeInterface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterface{ExchangeInterfaceCaller: ExchangeInterfaceCaller{contract: contract}, ExchangeInterfaceTransactor: ExchangeInterfaceTransactor{contract: contract}}, nil
}

// NewExchangeInterfaceCaller creates a new read-only instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ExchangeInterfaceCaller, error) {
	contract, err := bindExchangeInterface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceCaller{contract: contract}, nil
}

// NewExchangeInterfaceTransactor creates a new write-only instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeInterfaceTransactor, error) {
	contract, err := bindExchangeInterface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceTransactor{contract: contract}, nil
}

// bindExchangeInterface binds a generic wrapper to an already deployed contract.
func bindExchangeInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeInterface *ExchangeInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeInterface.Contract.ExchangeInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeInterface *ExchangeInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.ExchangeInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeInterface *ExchangeInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.ExchangeInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeInterface *ExchangeInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeInterface *ExchangeInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeInterface *ExchangeInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_ExchangeInterface *ExchangeInterfaceCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeInterface.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_ExchangeInterface *ExchangeInterfaceSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _ExchangeInterface.Contract.BalanceOf(&_ExchangeInterface.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_ExchangeInterface *ExchangeInterfaceCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _ExchangeInterface.Contract.BalanceOf(&_ExchangeInterface.CallOpts, token, user)
}

// CanTrade is a free data retrieval call binding the contract method 0xd4e0b0be.
//
// Solidity: function canTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, hash bytes32) constant returns(bool)
func (_ExchangeInterface *ExchangeInterfaceCaller) CanTrade(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeInterface.contract.Call(opts, out, "canTrade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, hash)
	return *ret0, err
}

// CanTrade is a free data retrieval call binding the contract method 0xd4e0b0be.
//
// Solidity: function canTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, hash bytes32) constant returns(bool)
func (_ExchangeInterface *ExchangeInterfaceSession) CanTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, hash [32]byte) (bool, error) {
	return _ExchangeInterface.Contract.CanTrade(&_ExchangeInterface.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, hash)
}

// CanTrade is a free data retrieval call binding the contract method 0xd4e0b0be.
//
// Solidity: function canTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, hash bytes32) constant returns(bool)
func (_ExchangeInterface *ExchangeInterfaceCallerSession) CanTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, hash [32]byte) (bool, error) {
	return _ExchangeInterface.Contract.CanTrade(&_ExchangeInterface.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0x406edbd3.
//
// Solidity: function cancel(expires uint256, amountGive uint256, amountGet uint256, tokenGet address, tokenGive address, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactor) Cancel(opts *bind.TransactOpts, expires *big.Int, amountGive *big.Int, amountGet *big.Int, tokenGet common.Address, tokenGive common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ExchangeInterface.contract.Transact(opts, "cancel", expires, amountGive, amountGet, tokenGet, tokenGive, nonce, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x406edbd3.
//
// Solidity: function cancel(expires uint256, amountGive uint256, amountGet uint256, tokenGet address, tokenGive address, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_ExchangeInterface *ExchangeInterfaceSession) Cancel(expires *big.Int, amountGive *big.Int, amountGet *big.Int, tokenGet common.Address, tokenGive common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Cancel(&_ExchangeInterface.TransactOpts, expires, amountGive, amountGet, tokenGet, tokenGive, nonce, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x406edbd3.
//
// Solidity: function cancel(expires uint256, amountGive uint256, amountGet uint256, tokenGet address, tokenGive address, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactorSession) Cancel(expires *big.Int, amountGive *big.Int, amountGet *big.Int, tokenGet common.Address, tokenGive common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Cancel(&_ExchangeInterface.TransactOpts, expires, amountGive, amountGet, tokenGet, tokenGive, nonce, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Deposit(&_ExchangeInterface.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Deposit(&_ExchangeInterface.TransactOpts, token, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactor) Trade(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.contract.Transact(opts, "trade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Trade(&_ExchangeInterface.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactorSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Trade(&_ExchangeInterface.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Withdraw(&_ExchangeInterface.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_ExchangeInterface *ExchangeInterfaceTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeInterface.Contract.Withdraw(&_ExchangeInterface.TransactOpts, token, amount)
}
