// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

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

// ExchangeInterfaceABI is the input ABI used to generate the binding from.

const ExchangeInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"canTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getVolume\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"}],\"name\":\"Traded\",\"type\":\"event\"}]"

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
	return address, tx, &ExchangeInterface{ExchangeInterfaceCaller: ExchangeInterfaceCaller{contract: contract}, ExchangeInterfaceTransactor: ExchangeInterfaceTransactor{contract: contract}, ExchangeInterfaceFilterer: ExchangeInterfaceFilterer{contract: contract}}, nil
}

// ExchangeInterface is an auto generated Go binding around an Ethereum contract.
type ExchangeInterface struct {
	ExchangeInterfaceCaller     // Read-only binding to the contract
	ExchangeInterfaceTransactor // Write-only binding to the contract
	ExchangeInterfaceFilterer   // Log filterer for contract events
}

// ExchangeInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeInterfaceFilterer struct {
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
	contract, err := bindExchangeInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterface{ExchangeInterfaceCaller: ExchangeInterfaceCaller{contract: contract}, ExchangeInterfaceTransactor: ExchangeInterfaceTransactor{contract: contract}, ExchangeInterfaceFilterer: ExchangeInterfaceFilterer{contract: contract}}, nil
}

// NewExchangeInterfaceCaller creates a new read-only instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ExchangeInterfaceCaller, error) {
	contract, err := bindExchangeInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceCaller{contract: contract}, nil
}

// NewExchangeInterfaceTransactor creates a new write-only instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeInterfaceTransactor, error) {
	contract, err := bindExchangeInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceTransactor{contract: contract}, nil
}

// NewExchangeInterfaceFilterer creates a new log filterer instance of ExchangeInterface, bound to a specific deployed contract.
func NewExchangeInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeInterfaceFilterer, error) {
	contract, err := bindExchangeInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceFilterer{contract: contract}, nil
}

// bindExchangeInterface binds a generic wrapper to an already deployed contract.
func bindExchangeInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ExchangeInterfaceCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the ExchangeInterface contract.
type ExchangeInterfaceCancelledIterator struct {
	Event *ExchangeInterfaceCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeInterfaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInterfaceCancelled)
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
		it.Event = new(ExchangeInterfaceCancelled)
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
func (it *ExchangeInterfaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInterfaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInterfaceCancelled represents a Cancelled event raised by the ExchangeInterface contract.
type ExchangeInterfaceCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(hash indexed bytes32)
func (_ExchangeInterface *ExchangeInterfaceFilterer) FilterCancelled(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeInterfaceCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ExchangeInterface.contract.FilterLogs(opts, "Cancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceCancelledIterator{contract: _ExchangeInterface.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(hash indexed bytes32)
func (_ExchangeInterface *ExchangeInterfaceFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeInterfaceCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ExchangeInterface.contract.WatchLogs(opts, "Cancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInterfaceCancelled)
				if err := _ExchangeInterface.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ExchangeInterfaceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ExchangeInterface contract.
type ExchangeInterfaceDepositedIterator struct {
	Event *ExchangeInterfaceDeposited // Event containing the contract specifics and raw log

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
func (it *ExchangeInterfaceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInterfaceDeposited)
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
		it.Event = new(ExchangeInterfaceDeposited)
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
func (it *ExchangeInterfaceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInterfaceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInterfaceDeposited represents a Deposited event raised by the ExchangeInterface contract.
type ExchangeInterfaceDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_ExchangeInterface *ExchangeInterfaceFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*ExchangeInterfaceDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangeInterface.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceDepositedIterator{contract: _ExchangeInterface.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_ExchangeInterface *ExchangeInterfaceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ExchangeInterfaceDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangeInterface.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInterfaceDeposited)
				if err := _ExchangeInterface.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ExchangeInterfaceTradedIterator is returned from FilterTraded and is used to iterate over the raw logs and unpacked data for Traded events raised by the ExchangeInterface contract.
type ExchangeInterfaceTradedIterator struct {
	Event *ExchangeInterfaceTraded // Event containing the contract specifics and raw log

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
func (it *ExchangeInterfaceTradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInterfaceTraded)
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
		it.Event = new(ExchangeInterfaceTraded)
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
func (it *ExchangeInterfaceTradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInterfaceTradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInterfaceTraded represents a Traded event raised by the ExchangeInterface contract.
type ExchangeInterfaceTraded struct {
	Hash       [32]byte
	TokenGive  common.Address
	AmountGive *big.Int
	TokenGet   common.Address
	AmountGet  *big.Int
	Maker      common.Address
	Taker      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTraded is a free log retrieval operation binding the contract event 0xe1d2889bf5062ca6cccab7b9d6f0548e654943875f2a9c45eaaef37b11d7f68c.
//
// Solidity: event Traded(hash indexed bytes32, tokenGive address, amountGive uint256, tokenGet address, amountGet uint256, maker address, taker address)
func (_ExchangeInterface *ExchangeInterfaceFilterer) FilterTraded(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeInterfaceTradedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ExchangeInterface.contract.FilterLogs(opts, "Traded", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceTradedIterator{contract: _ExchangeInterface.contract, event: "Traded", logs: logs, sub: sub}, nil
}

// WatchTraded is a free log subscription operation binding the contract event 0xe1d2889bf5062ca6cccab7b9d6f0548e654943875f2a9c45eaaef37b11d7f68c.
//
// Solidity: event Traded(hash indexed bytes32, tokenGive address, amountGive uint256, tokenGet address, amountGet uint256, maker address, taker address)
func (_ExchangeInterface *ExchangeInterfaceFilterer) WatchTraded(opts *bind.WatchOpts, sink chan<- *ExchangeInterfaceTraded, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ExchangeInterface.contract.WatchLogs(opts, "Traded", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInterfaceTraded)
				if err := _ExchangeInterface.contract.UnpackLog(event, "Traded", log); err != nil {
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

// ExchangeInterfaceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ExchangeInterface contract.
type ExchangeInterfaceWithdrawnIterator struct {
	Event *ExchangeInterfaceWithdrawn // Event containing the contract specifics and raw log

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
func (it *ExchangeInterfaceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInterfaceWithdrawn)
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
		it.Event = new(ExchangeInterfaceWithdrawn)
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
func (it *ExchangeInterfaceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInterfaceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInterfaceWithdrawn represents a Withdrawn event raised by the ExchangeInterface contract.
type ExchangeInterfaceWithdrawn struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_ExchangeInterface *ExchangeInterfaceFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*ExchangeInterfaceWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangeInterface.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeInterfaceWithdrawnIterator{contract: _ExchangeInterface.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_ExchangeInterface *ExchangeInterfaceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ExchangeInterfaceWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangeInterface.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInterfaceWithdrawn)
				if err := _ExchangeInterface.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
