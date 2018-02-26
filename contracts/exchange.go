// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"canTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getVolume\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"}],\"name\":\"Traded\",\"type\":\"event\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
const ExchangeBin = `0x`

// DeployExchange deploys a new Ethereum contract, binding an instance of Exchange to it.
func DeployExchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exchange, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// CanTrade is a free data retrieval call binding the contract method 0x0b98f9ad.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) constant returns(bool)
func (_Exchange *ExchangeCaller) CanTrade(opts *bind.CallOpts, addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "canTrade", addresses, values, v, r, s, amount, mode)
	return *ret0, err
}

// CanTrade is a free data retrieval call binding the contract method 0x0b98f9ad.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) constant returns(bool)
func (_Exchange *ExchangeSession) CanTrade(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (bool, error) {
	return _Exchange.Contract.CanTrade(&_Exchange.CallOpts, addresses, values, v, r, s, amount, mode)
}

// CanTrade is a free data retrieval call binding the contract method 0x0b98f9ad.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) constant returns(bool)
func (_Exchange *ExchangeCallerSession) CanTrade(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (bool, error) {
	return _Exchange.Contract.CanTrade(&_Exchange.CallOpts, addresses, values, v, r, s, amount, mode)
}

// Filled is a free data retrieval call binding the contract method 0xa3093e0f.
//
// Solidity: function filled(user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Filled(opts *bind.CallOpts, user common.Address, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "filled", user, hash)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0xa3093e0f.
//
// Solidity: function filled(user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Filled(user common.Address, hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, user, hash)
}

// Filled is a free data retrieval call binding the contract method 0xa3093e0f.
//
// Solidity: function filled(user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Filled(user common.Address, hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, user, hash)
}

// GetVolume is a free data retrieval call binding the contract method 0xa8aa5d34.
//
// Solidity: function getVolume(amountGet uint256, tokenGive address, amountGive uint256, user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) GetVolume(opts *bind.CallOpts, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, user common.Address, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getVolume", amountGet, tokenGive, amountGive, user, hash)
	return *ret0, err
}

// GetVolume is a free data retrieval call binding the contract method 0xa8aa5d34.
//
// Solidity: function getVolume(amountGet uint256, tokenGive address, amountGive uint256, user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) GetVolume(amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, user common.Address, hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetVolume(&_Exchange.CallOpts, amountGet, tokenGive, amountGive, user, hash)
}

// GetVolume is a free data retrieval call binding the contract method 0xa8aa5d34.
//
// Solidity: function getVolume(amountGet uint256, tokenGive address, amountGive uint256, user address, hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetVolume(amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, user common.Address, hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetVolume(&_Exchange.CallOpts, amountGet, tokenGive, amountGive, user, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0x93503e36.
//
// Solidity: function cancel(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, mode uint256) returns()
func (_Exchange *ExchangeTransactor) Cancel(opts *bind.TransactOpts, addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancel", addresses, values, v, r, s, mode)
}

// Cancel is a paid mutator transaction binding the contract method 0x93503e36.
//
// Solidity: function cancel(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, mode uint256) returns()
func (_Exchange *ExchangeSession) Cancel(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, addresses, values, v, r, s, mode)
}

// Cancel is a paid mutator transaction binding the contract method 0x93503e36.
//
// Solidity: function cancel(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, mode uint256) returns()
func (_Exchange *ExchangeTransactorSession) Cancel(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, addresses, values, v, r, s, mode)
}

// Trade is a paid mutator transaction binding the contract method 0x7e707789.
//
// Solidity: function trade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) returns()
func (_Exchange *ExchangeTransactor) Trade(opts *bind.TransactOpts, addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "trade", addresses, values, v, r, s, amount, mode)
}

// Trade is a paid mutator transaction binding the contract method 0x7e707789.
//
// Solidity: function trade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) returns()
func (_Exchange *ExchangeSession) Trade(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Trade(&_Exchange.TransactOpts, addresses, values, v, r, s, amount, mode)
}

// Trade is a paid mutator transaction binding the contract method 0x7e707789.
//
// Solidity: function trade(addresses address[3], values uint256[4], v uint8, r bytes32, s bytes32, amount uint256, mode uint256) returns()
func (_Exchange *ExchangeTransactorSession) Trade(addresses [3]common.Address, values [4]*big.Int, v uint8, r [32]byte, s [32]byte, amount *big.Int, mode *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Trade(&_Exchange.TransactOpts, addresses, values, v, r, s, amount, mode)
}

// ExchangeCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the Exchange contract.
type ExchangeCancelledIterator struct {
	Event *ExchangeCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeCancelled)
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
		it.Event = new(ExchangeCancelled)
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
func (it *ExchangeCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeCancelled represents a Cancelled event raised by the Exchange contract.
type ExchangeCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(hash indexed bytes32)
func (_Exchange *ExchangeFilterer) FilterCancelled(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Cancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeCancelledIterator{contract: _Exchange.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(hash indexed bytes32)
func (_Exchange *ExchangeFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Cancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeCancelled)
				if err := _Exchange.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ExchangeTradedIterator is returned from FilterTraded and is used to iterate over the raw logs and unpacked data for Traded events raised by the Exchange contract.
type ExchangeTradedIterator struct {
	Event *ExchangeTraded // Event containing the contract specifics and raw log

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
func (it *ExchangeTradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTraded)
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
		it.Event = new(ExchangeTraded)
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
func (it *ExchangeTradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTraded represents a Traded event raised by the Exchange contract.
type ExchangeTraded struct {
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
func (_Exchange *ExchangeFilterer) FilterTraded(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeTradedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Traded", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTradedIterator{contract: _Exchange.contract, event: "Traded", logs: logs, sub: sub}, nil
}

// WatchTraded is a free log subscription operation binding the contract event 0xe1d2889bf5062ca6cccab7b9d6f0548e654943875f2a9c45eaaef37b11d7f68c.
//
// Solidity: event Traded(hash indexed bytes32, tokenGive address, amountGive uint256, tokenGet address, amountGet uint256, maker address, taker address)
func (_Exchange *ExchangeFilterer) WatchTraded(opts *bind.WatchOpts, sink chan<- *ExchangeTraded, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Traded", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTraded)
				if err := _Exchange.contract.UnpackLog(event, "Traded", log); err != nil {
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
