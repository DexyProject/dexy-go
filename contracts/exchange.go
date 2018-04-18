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
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isOrdered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"maxFillAmount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"canTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"}],\"name\":\"availableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[3]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[2]\"},{\"name\":\"values\",\"type\":\"uint256[4]\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"}],\"name\":\"Traded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Ordered\",\"type\":\"event\"}]"

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
	Contract     *Exchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
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

// AvailableAmount is a free data retrieval call binding the contract method 0x1a197588.
//
// Solidity: function availableAmount(addresses address[3], values uint256[4]) constant returns(uint256)
func (_Exchange *ExchangeCaller) AvailableAmount(opts *bind.CallOpts, addresses [3]common.Address, values [4]*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "availableAmount", addresses, values)
	return *ret0, err
}

// AvailableAmount is a free data retrieval call binding the contract method 0x1a197588.
//
// Solidity: function availableAmount(addresses address[3], values uint256[4]) constant returns(uint256)
func (_Exchange *ExchangeSession) AvailableAmount(addresses [3]common.Address, values [4]*big.Int) (*big.Int, error) {
	return _Exchange.Contract.AvailableAmount(&_Exchange.CallOpts, addresses, values)
}

// AvailableAmount is a free data retrieval call binding the contract method 0x1a197588.
//
// Solidity: function availableAmount(addresses address[3], values uint256[4]) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) AvailableAmount(addresses [3]common.Address, values [4]*big.Int) (*big.Int, error) {
	return _Exchange.Contract.AvailableAmount(&_Exchange.CallOpts, addresses, values)
}

// CanTrade is a free data retrieval call binding the contract method 0x08fa0e92.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], signature bytes) constant returns(bool)
func (_Exchange *ExchangeCaller) CanTrade(opts *bind.CallOpts, addresses [3]common.Address, values [4]*big.Int, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "canTrade", addresses, values, signature)
	return *ret0, err
}

// CanTrade is a free data retrieval call binding the contract method 0x08fa0e92.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], signature bytes) constant returns(bool)
func (_Exchange *ExchangeSession) CanTrade(addresses [3]common.Address, values [4]*big.Int, signature []byte) (bool, error) {
	return _Exchange.Contract.CanTrade(&_Exchange.CallOpts, addresses, values, signature)
}

// CanTrade is a free data retrieval call binding the contract method 0x08fa0e92.
//
// Solidity: function canTrade(addresses address[3], values uint256[4], signature bytes) constant returns(bool)
func (_Exchange *ExchangeCallerSession) CanTrade(addresses [3]common.Address, values [4]*big.Int, signature []byte) (bool, error) {
	return _Exchange.Contract.CanTrade(&_Exchange.CallOpts, addresses, values, signature)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Filled(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "filled", hash)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Filled(hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, hash)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(hash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Filled(hash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, hash)
}

// IsOrdered is a free data retrieval call binding the contract method 0x00f29d55.
//
// Solidity: function isOrdered(user address, hash bytes32) constant returns(bool)
func (_Exchange *ExchangeCaller) IsOrdered(opts *bind.CallOpts, user common.Address, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isOrdered", user, hash)
	return *ret0, err
}

// IsOrdered is a free data retrieval call binding the contract method 0x00f29d55.
//
// Solidity: function isOrdered(user address, hash bytes32) constant returns(bool)
func (_Exchange *ExchangeSession) IsOrdered(user common.Address, hash [32]byte) (bool, error) {
	return _Exchange.Contract.IsOrdered(&_Exchange.CallOpts, user, hash)
}

// IsOrdered is a free data retrieval call binding the contract method 0x00f29d55.
//
// Solidity: function isOrdered(user address, hash bytes32) constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsOrdered(user common.Address, hash [32]byte) (bool, error) {
	return _Exchange.Contract.IsOrdered(&_Exchange.CallOpts, user, hash)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Exchange *ExchangeCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Exchange *ExchangeSession) Vault() (common.Address, error) {
	return _Exchange.Contract.Vault(&_Exchange.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Exchange *ExchangeCallerSession) Vault() (common.Address, error) {
	return _Exchange.Contract.Vault(&_Exchange.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xb1c0e063.
//
// Solidity: function cancel(addresses address[3], values uint256[4]) returns()
func (_Exchange *ExchangeTransactor) Cancel(opts *bind.TransactOpts, addresses [3]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancel", addresses, values)
}

// Cancel is a paid mutator transaction binding the contract method 0xb1c0e063.
//
// Solidity: function cancel(addresses address[3], values uint256[4]) returns()
func (_Exchange *ExchangeSession) Cancel(addresses [3]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, addresses, values)
}

// Cancel is a paid mutator transaction binding the contract method 0xb1c0e063.
//
// Solidity: function cancel(addresses address[3], values uint256[4]) returns()
func (_Exchange *ExchangeTransactorSession) Cancel(addresses [3]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, addresses, values)
}

// Order is a paid mutator transaction binding the contract method 0xc6f54e62.
//
// Solidity: function order(addresses address[2], values uint256[4]) returns()
func (_Exchange *ExchangeTransactor) Order(opts *bind.TransactOpts, addresses [2]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "order", addresses, values)
}

// Order is a paid mutator transaction binding the contract method 0xc6f54e62.
//
// Solidity: function order(addresses address[2], values uint256[4]) returns()
func (_Exchange *ExchangeSession) Order(addresses [2]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, addresses, values)
}

// Order is a paid mutator transaction binding the contract method 0xc6f54e62.
//
// Solidity: function order(addresses address[2], values uint256[4]) returns()
func (_Exchange *ExchangeTransactorSession) Order(addresses [2]common.Address, values [4]*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Order(&_Exchange.TransactOpts, addresses, values)
}

// Trade is a paid mutator transaction binding the contract method 0x08218c98.
//
// Solidity: function trade(addresses address[3], values uint256[4], signature bytes, maxFillAmount uint256) returns()
func (_Exchange *ExchangeTransactor) Trade(opts *bind.TransactOpts, addresses [3]common.Address, values [4]*big.Int, signature []byte, maxFillAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "trade", addresses, values, signature, maxFillAmount)
}

// Trade is a paid mutator transaction binding the contract method 0x08218c98.
//
// Solidity: function trade(addresses address[3], values uint256[4], signature bytes, maxFillAmount uint256) returns()
func (_Exchange *ExchangeSession) Trade(addresses [3]common.Address, values [4]*big.Int, signature []byte, maxFillAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Trade(&_Exchange.TransactOpts, addresses, values, signature, maxFillAmount)
}

// Trade is a paid mutator transaction binding the contract method 0x08218c98.
//
// Solidity: function trade(addresses address[3], values uint256[4], signature bytes, maxFillAmount uint256) returns()
func (_Exchange *ExchangeTransactorSession) Trade(addresses [3]common.Address, values [4]*big.Int, signature []byte, maxFillAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Trade(&_Exchange.TransactOpts, addresses, values, signature, maxFillAmount)
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

// ExchangeOrderedIterator is returned from FilterOrdered and is used to iterate over the raw logs and unpacked data for Ordered events raised by the Exchange contract.
type ExchangeOrderedIterator struct {
	Event *ExchangeOrdered // Event containing the contract specifics and raw log

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
func (it *ExchangeOrderedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrdered)
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
		it.Event = new(ExchangeOrdered)
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
func (it *ExchangeOrderedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrderedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrdered represents a Ordered event raised by the Exchange contract.
type ExchangeOrdered struct {
	Maker            common.Address
	MakerToken       common.Address
	TakerToken       common.Address
	MakerTokenAmount *big.Int
	TakerTokenAmount *big.Int
	Expires          *big.Int
	Nonce            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrdered is a free log retrieval operation binding the contract event 0x24ec4e2d3ad6fb01b5c3b3466504af096fcd9b951cd27a3a0e56225d39c17aa0.
//
// Solidity: event Ordered(maker address, makerToken address, takerToken address, makerTokenAmount uint256, takerTokenAmount uint256, expires uint256, nonce uint256)
func (_Exchange *ExchangeFilterer) FilterOrdered(opts *bind.FilterOpts) (*ExchangeOrderedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Ordered")
	if err != nil {
		return nil, err
	}
	return &ExchangeOrderedIterator{contract: _Exchange.contract, event: "Ordered", logs: logs, sub: sub}, nil
}

// WatchOrdered is a free log subscription operation binding the contract event 0x24ec4e2d3ad6fb01b5c3b3466504af096fcd9b951cd27a3a0e56225d39c17aa0.
//
// Solidity: event Ordered(maker address, makerToken address, takerToken address, makerTokenAmount uint256, takerTokenAmount uint256, expires uint256, nonce uint256)
func (_Exchange *ExchangeFilterer) WatchOrdered(opts *bind.WatchOpts, sink chan<- *ExchangeOrdered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Ordered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrdered)
				if err := _Exchange.contract.UnpackLog(event, "Ordered", log); err != nil {
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
	Hash             [32]byte
	MakerToken       common.Address
	MakerTokenAmount *big.Int
	TakerToken       common.Address
	TakerTokenAmount *big.Int
	Maker            common.Address
	Taker            common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTraded is a free log retrieval operation binding the contract event 0xe1d2889bf5062ca6cccab7b9d6f0548e654943875f2a9c45eaaef37b11d7f68c.
//
// Solidity: event Traded(hash indexed bytes32, makerToken address, makerTokenAmount uint256, takerToken address, takerTokenAmount uint256, maker address, taker address)
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
// Solidity: event Traded(hash indexed bytes32, makerToken address, makerTokenAmount uint256, takerToken address, takerTokenAmount uint256, maker address, taker address)
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

// VaultInterfaceABI is the input ABI used to generate the binding from.
const VaultInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestSpender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"removeSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isSpender\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"addSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"unapprove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Unapproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"AddedSpender\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"RemovedSpender\",\"type\":\"event\"}]"

// VaultInterfaceBin is the compiled bytecode used for deploying new contracts.
const VaultInterfaceBin = `0x`

// DeployVaultInterface deploys a new Ethereum contract, binding an instance of VaultInterface to it.
func DeployVaultInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultInterface{VaultInterfaceCaller: VaultInterfaceCaller{contract: contract}, VaultInterfaceTransactor: VaultInterfaceTransactor{contract: contract}, VaultInterfaceFilterer: VaultInterfaceFilterer{contract: contract}}, nil
}

// VaultInterface is an auto generated Go binding around an Ethereum contract.
type VaultInterface struct {
	VaultInterfaceCaller     // Read-only binding to the contract
	VaultInterfaceTransactor // Write-only binding to the contract
	VaultInterfaceFilterer   // Log filterer for contract events
}

// VaultInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultInterfaceSession struct {
	Contract     *VaultInterface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultInterfaceCallerSession struct {
	Contract *VaultInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VaultInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultInterfaceTransactorSession struct {
	Contract     *VaultInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VaultInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultInterfaceRaw struct {
	Contract *VaultInterface // Generic contract binding to access the raw methods on
}

// VaultInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultInterfaceCallerRaw struct {
	Contract *VaultInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// VaultInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultInterfaceTransactorRaw struct {
	Contract *VaultInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultInterface creates a new instance of VaultInterface, bound to a specific deployed contract.
func NewVaultInterface(address common.Address, backend bind.ContractBackend) (*VaultInterface, error) {
	contract, err := bindVaultInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultInterface{VaultInterfaceCaller: VaultInterfaceCaller{contract: contract}, VaultInterfaceTransactor: VaultInterfaceTransactor{contract: contract}, VaultInterfaceFilterer: VaultInterfaceFilterer{contract: contract}}, nil
}

// NewVaultInterfaceCaller creates a new read-only instance of VaultInterface, bound to a specific deployed contract.
func NewVaultInterfaceCaller(address common.Address, caller bind.ContractCaller) (*VaultInterfaceCaller, error) {
	contract, err := bindVaultInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceCaller{contract: contract}, nil
}

// NewVaultInterfaceTransactor creates a new write-only instance of VaultInterface, bound to a specific deployed contract.
func NewVaultInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultInterfaceTransactor, error) {
	contract, err := bindVaultInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceTransactor{contract: contract}, nil
}

// NewVaultInterfaceFilterer creates a new log filterer instance of VaultInterface, bound to a specific deployed contract.
func NewVaultInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultInterfaceFilterer, error) {
	contract, err := bindVaultInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceFilterer{contract: contract}, nil
}

// bindVaultInterface binds a generic wrapper to an already deployed contract.
func bindVaultInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultInterface *VaultInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultInterface.Contract.VaultInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultInterface *VaultInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultInterface.Contract.VaultInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultInterface *VaultInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultInterface.Contract.VaultInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultInterface *VaultInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultInterface *VaultInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultInterface *VaultInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultInterface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_VaultInterface *VaultInterfaceCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VaultInterface.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_VaultInterface *VaultInterfaceSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _VaultInterface.Contract.BalanceOf(&_VaultInterface.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_VaultInterface *VaultInterfaceCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _VaultInterface.Contract.BalanceOf(&_VaultInterface.CallOpts, token, user)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceCaller) IsApproved(opts *bind.CallOpts, user common.Address, spender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VaultInterface.contract.Call(opts, out, "isApproved", user, spender)
	return *ret0, err
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceSession) IsApproved(user common.Address, spender common.Address) (bool, error) {
	return _VaultInterface.Contract.IsApproved(&_VaultInterface.CallOpts, user, spender)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceCallerSession) IsApproved(user common.Address, spender common.Address) (bool, error) {
	return _VaultInterface.Contract.IsApproved(&_VaultInterface.CallOpts, user, spender)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceCaller) IsSpender(opts *bind.CallOpts, spender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VaultInterface.contract.Call(opts, out, "isSpender", spender)
	return *ret0, err
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceSession) IsSpender(spender common.Address) (bool, error) {
	return _VaultInterface.Contract.IsSpender(&_VaultInterface.CallOpts, spender)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_VaultInterface *VaultInterfaceCallerSession) IsSpender(spender common.Address) (bool, error) {
	return _VaultInterface.Contract.IsSpender(&_VaultInterface.CallOpts, spender)
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_VaultInterface *VaultInterfaceCaller) LatestSpender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VaultInterface.contract.Call(opts, out, "latestSpender")
	return *ret0, err
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_VaultInterface *VaultInterfaceSession) LatestSpender() (common.Address, error) {
	return _VaultInterface.Contract.LatestSpender(&_VaultInterface.CallOpts)
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_VaultInterface *VaultInterfaceCallerSession) LatestSpender() (common.Address, error) {
	return _VaultInterface.Contract.LatestSpender(&_VaultInterface.CallOpts)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactor) AddSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "addSpender", spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.AddSpender(&_VaultInterface.TransactOpts, spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.AddSpender(&_VaultInterface.TransactOpts, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "approve", spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_VaultInterface *VaultInterfaceSession) Approve(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.Approve(&_VaultInterface.TransactOpts, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) Approve(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.Approve(&_VaultInterface.TransactOpts, spender)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Deposit(&_VaultInterface.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Deposit(&_VaultInterface.TransactOpts, token, amount)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactor) RemoveSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "removeSpender", spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.RemoveSpender(&_VaultInterface.TransactOpts, spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.RemoveSpender(&_VaultInterface.TransactOpts, spender)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_VaultInterface *VaultInterfaceTransactor) TokenFallback(opts *bind.TransactOpts, from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "tokenFallback", from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_VaultInterface *VaultInterfaceSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VaultInterface.Contract.TokenFallback(&_VaultInterface.TransactOpts, from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VaultInterface.Contract.TokenFallback(&_VaultInterface.TransactOpts, from, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactor) Transfer(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "transfer", token, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceSession) Transfer(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Transfer(&_VaultInterface.TransactOpts, token, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) Transfer(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Transfer(&_VaultInterface.TransactOpts, token, from, to, amount)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactor) Unapprove(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "unapprove", spender)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_VaultInterface *VaultInterfaceSession) Unapprove(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.Unapprove(&_VaultInterface.TransactOpts, spender)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) Unapprove(spender common.Address) (*types.Transaction, error) {
	return _VaultInterface.Contract.Unapprove(&_VaultInterface.TransactOpts, spender)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Withdraw(&_VaultInterface.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_VaultInterface *VaultInterfaceTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultInterface.Contract.Withdraw(&_VaultInterface.TransactOpts, token, amount)
}

// VaultInterfaceAddedSpenderIterator is returned from FilterAddedSpender and is used to iterate over the raw logs and unpacked data for AddedSpender events raised by the VaultInterface contract.
type VaultInterfaceAddedSpenderIterator struct {
	Event *VaultInterfaceAddedSpender // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceAddedSpenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceAddedSpender)
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
		it.Event = new(VaultInterfaceAddedSpender)
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
func (it *VaultInterfaceAddedSpenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceAddedSpenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceAddedSpender represents a AddedSpender event raised by the VaultInterface contract.
type VaultInterfaceAddedSpender struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedSpender is a free log retrieval operation binding the contract event 0x8fd571ab479506dd07023e78f221245916b6cb54285d954030be2cfb1674657a.
//
// Solidity: event AddedSpender(spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) FilterAddedSpender(opts *bind.FilterOpts, spender []common.Address) (*VaultInterfaceAddedSpenderIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "AddedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceAddedSpenderIterator{contract: _VaultInterface.contract, event: "AddedSpender", logs: logs, sub: sub}, nil
}

// WatchAddedSpender is a free log subscription operation binding the contract event 0x8fd571ab479506dd07023e78f221245916b6cb54285d954030be2cfb1674657a.
//
// Solidity: event AddedSpender(spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) WatchAddedSpender(opts *bind.WatchOpts, sink chan<- *VaultInterfaceAddedSpender, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "AddedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceAddedSpender)
				if err := _VaultInterface.contract.UnpackLog(event, "AddedSpender", log); err != nil {
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

// VaultInterfaceApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the VaultInterface contract.
type VaultInterfaceApprovedIterator struct {
	Event *VaultInterfaceApproved // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceApproved)
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
		it.Event = new(VaultInterfaceApproved)
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
func (it *VaultInterfaceApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceApproved represents a Approved event raised by the VaultInterface contract.
type VaultInterfaceApproved struct {
	User    common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xaad2833c9fd7a3de33f301e5186ee84d1a5753ce32de6b97baedaac4b92b55fc.
//
// Solidity: event Approved(user indexed address, spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) FilterApproved(opts *bind.FilterOpts, user []common.Address, spender []common.Address) (*VaultInterfaceApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "Approved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceApprovedIterator{contract: _VaultInterface.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xaad2833c9fd7a3de33f301e5186ee84d1a5753ce32de6b97baedaac4b92b55fc.
//
// Solidity: event Approved(user indexed address, spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *VaultInterfaceApproved, user []common.Address, spender []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "Approved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceApproved)
				if err := _VaultInterface.contract.UnpackLog(event, "Approved", log); err != nil {
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

// VaultInterfaceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the VaultInterface contract.
type VaultInterfaceDepositedIterator struct {
	Event *VaultInterfaceDeposited // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceDeposited)
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
		it.Event = new(VaultInterfaceDeposited)
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
func (it *VaultInterfaceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceDeposited represents a Deposited event raised by the VaultInterface contract.
type VaultInterfaceDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_VaultInterface *VaultInterfaceFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*VaultInterfaceDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceDepositedIterator{contract: _VaultInterface.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_VaultInterface *VaultInterfaceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *VaultInterfaceDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceDeposited)
				if err := _VaultInterface.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// VaultInterfaceRemovedSpenderIterator is returned from FilterRemovedSpender and is used to iterate over the raw logs and unpacked data for RemovedSpender events raised by the VaultInterface contract.
type VaultInterfaceRemovedSpenderIterator struct {
	Event *VaultInterfaceRemovedSpender // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceRemovedSpenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceRemovedSpender)
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
		it.Event = new(VaultInterfaceRemovedSpender)
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
func (it *VaultInterfaceRemovedSpenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceRemovedSpenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceRemovedSpender represents a RemovedSpender event raised by the VaultInterface contract.
type VaultInterfaceRemovedSpender struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedSpender is a free log retrieval operation binding the contract event 0x0e2fc808ab0ead56889f8ff2a8ea0841ba4c0b8311607a902eb24b834857e1b5.
//
// Solidity: event RemovedSpender(spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) FilterRemovedSpender(opts *bind.FilterOpts, spender []common.Address) (*VaultInterfaceRemovedSpenderIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "RemovedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceRemovedSpenderIterator{contract: _VaultInterface.contract, event: "RemovedSpender", logs: logs, sub: sub}, nil
}

// WatchRemovedSpender is a free log subscription operation binding the contract event 0x0e2fc808ab0ead56889f8ff2a8ea0841ba4c0b8311607a902eb24b834857e1b5.
//
// Solidity: event RemovedSpender(spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) WatchRemovedSpender(opts *bind.WatchOpts, sink chan<- *VaultInterfaceRemovedSpender, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "RemovedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceRemovedSpender)
				if err := _VaultInterface.contract.UnpackLog(event, "RemovedSpender", log); err != nil {
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

// VaultInterfaceUnapprovedIterator is returned from FilterUnapproved and is used to iterate over the raw logs and unpacked data for Unapproved events raised by the VaultInterface contract.
type VaultInterfaceUnapprovedIterator struct {
	Event *VaultInterfaceUnapproved // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceUnapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceUnapproved)
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
		it.Event = new(VaultInterfaceUnapproved)
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
func (it *VaultInterfaceUnapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceUnapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceUnapproved represents a Unapproved event raised by the VaultInterface contract.
type VaultInterfaceUnapproved struct {
	User    common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnapproved is a free log retrieval operation binding the contract event 0x1ab270601cc6b54dd5e8ce5c70dbac96a01ff12939e4e76488df62adc8e68373.
//
// Solidity: event Unapproved(user indexed address, spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) FilterUnapproved(opts *bind.FilterOpts, user []common.Address, spender []common.Address) (*VaultInterfaceUnapprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "Unapproved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceUnapprovedIterator{contract: _VaultInterface.contract, event: "Unapproved", logs: logs, sub: sub}, nil
}

// WatchUnapproved is a free log subscription operation binding the contract event 0x1ab270601cc6b54dd5e8ce5c70dbac96a01ff12939e4e76488df62adc8e68373.
//
// Solidity: event Unapproved(user indexed address, spender indexed address)
func (_VaultInterface *VaultInterfaceFilterer) WatchUnapproved(opts *bind.WatchOpts, sink chan<- *VaultInterfaceUnapproved, user []common.Address, spender []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "Unapproved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceUnapproved)
				if err := _VaultInterface.contract.UnpackLog(event, "Unapproved", log); err != nil {
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

// VaultInterfaceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the VaultInterface contract.
type VaultInterfaceWithdrawnIterator struct {
	Event *VaultInterfaceWithdrawn // Event containing the contract specifics and raw log

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
func (it *VaultInterfaceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInterfaceWithdrawn)
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
		it.Event = new(VaultInterfaceWithdrawn)
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
func (it *VaultInterfaceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInterfaceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInterfaceWithdrawn represents a Withdrawn event raised by the VaultInterface contract.
type VaultInterfaceWithdrawn struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_VaultInterface *VaultInterfaceFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*VaultInterfaceWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VaultInterface.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultInterfaceWithdrawnIterator{contract: _VaultInterface.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_VaultInterface *VaultInterfaceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *VaultInterfaceWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VaultInterface.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInterfaceWithdrawn)
				if err := _VaultInterface.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
