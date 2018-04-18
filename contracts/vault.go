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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestSpender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"removeSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isSpender\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"addSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"unapprove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Unapproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"AddedSpender\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"RemovedSpender\",\"type\":\"event\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `0x`

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token, user)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_Vault *VaultCaller) IsApproved(opts *bind.CallOpts, user common.Address, spender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isApproved", user, spender)
	return *ret0, err
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_Vault *VaultSession) IsApproved(user common.Address, spender common.Address) (bool, error) {
	return _Vault.Contract.IsApproved(&_Vault.CallOpts, user, spender)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(user address, spender address) constant returns(bool)
func (_Vault *VaultCallerSession) IsApproved(user common.Address, spender common.Address) (bool, error) {
	return _Vault.Contract.IsApproved(&_Vault.CallOpts, user, spender)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_Vault *VaultCaller) IsSpender(opts *bind.CallOpts, spender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isSpender", spender)
	return *ret0, err
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_Vault *VaultSession) IsSpender(spender common.Address) (bool, error) {
	return _Vault.Contract.IsSpender(&_Vault.CallOpts, spender)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(spender address) constant returns(bool)
func (_Vault *VaultCallerSession) IsSpender(spender common.Address) (bool, error) {
	return _Vault.Contract.IsSpender(&_Vault.CallOpts, spender)
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_Vault *VaultCaller) LatestSpender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "latestSpender")
	return *ret0, err
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_Vault *VaultSession) LatestSpender() (common.Address, error) {
	return _Vault.Contract.LatestSpender(&_Vault.CallOpts)
}

// LatestSpender is a free data retrieval call binding the contract method 0x6f362c2b.
//
// Solidity: function latestSpender() constant returns(address)
func (_Vault *VaultCallerSession) LatestSpender() (common.Address, error) {
	return _Vault.Contract.LatestSpender(&_Vault.CallOpts)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_Vault *VaultTransactor) AddSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addSpender", spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_Vault *VaultSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddSpender(&_Vault.TransactOpts, spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(spender address) returns()
func (_Vault *VaultTransactorSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddSpender(&_Vault.TransactOpts, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_Vault *VaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "approve", spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_Vault *VaultSession) Approve(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Approve(&_Vault.TransactOpts, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(spender address) returns()
func (_Vault *VaultTransactorSession) Approve(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Approve(&_Vault.TransactOpts, spender)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_Vault *VaultSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(token address, amount uint256) returns()
func (_Vault *VaultTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, token, amount)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_Vault *VaultTransactor) RemoveSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "removeSpender", spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_Vault *VaultSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveSpender(&_Vault.TransactOpts, spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(spender address) returns()
func (_Vault *VaultTransactorSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveSpender(&_Vault.TransactOpts, spender)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_Vault *VaultTransactor) TokenFallback(opts *bind.TransactOpts, from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "tokenFallback", from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_Vault *VaultSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.TokenFallback(&_Vault.TransactOpts, from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from address, value uint256, data bytes) returns()
func (_Vault *VaultTransactorSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.TokenFallback(&_Vault.TransactOpts, from, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", token, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_Vault *VaultSession) Transfer(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(token address, from address, to address, amount uint256) returns()
func (_Vault *VaultTransactorSession) Transfer(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, from, to, amount)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_Vault *VaultTransactor) Unapprove(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "unapprove", spender)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_Vault *VaultSession) Unapprove(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Unapprove(&_Vault.TransactOpts, spender)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfbf1f78a.
//
// Solidity: function unapprove(spender address) returns()
func (_Vault *VaultTransactorSession) Unapprove(spender common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Unapprove(&_Vault.TransactOpts, spender)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Vault *VaultSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(token address, amount uint256) returns()
func (_Vault *VaultTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, token, amount)
}

// VaultAddedSpenderIterator is returned from FilterAddedSpender and is used to iterate over the raw logs and unpacked data for AddedSpender events raised by the Vault contract.
type VaultAddedSpenderIterator struct {
	Event *VaultAddedSpender // Event containing the contract specifics and raw log

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
func (it *VaultAddedSpenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAddedSpender)
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
		it.Event = new(VaultAddedSpender)
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
func (it *VaultAddedSpenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAddedSpenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAddedSpender represents a AddedSpender event raised by the Vault contract.
type VaultAddedSpender struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedSpender is a free log retrieval operation binding the contract event 0x8fd571ab479506dd07023e78f221245916b6cb54285d954030be2cfb1674657a.
//
// Solidity: event AddedSpender(spender indexed address)
func (_Vault *VaultFilterer) FilterAddedSpender(opts *bind.FilterOpts, spender []common.Address) (*VaultAddedSpenderIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AddedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultAddedSpenderIterator{contract: _Vault.contract, event: "AddedSpender", logs: logs, sub: sub}, nil
}

// WatchAddedSpender is a free log subscription operation binding the contract event 0x8fd571ab479506dd07023e78f221245916b6cb54285d954030be2cfb1674657a.
//
// Solidity: event AddedSpender(spender indexed address)
func (_Vault *VaultFilterer) WatchAddedSpender(opts *bind.WatchOpts, sink chan<- *VaultAddedSpender, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AddedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAddedSpender)
				if err := _Vault.contract.UnpackLog(event, "AddedSpender", log); err != nil {
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

// VaultApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the Vault contract.
type VaultApprovedIterator struct {
	Event *VaultApproved // Event containing the contract specifics and raw log

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
func (it *VaultApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultApproved)
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
		it.Event = new(VaultApproved)
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
func (it *VaultApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultApproved represents a Approved event raised by the Vault contract.
type VaultApproved struct {
	User    common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xaad2833c9fd7a3de33f301e5186ee84d1a5753ce32de6b97baedaac4b92b55fc.
//
// Solidity: event Approved(user indexed address, spender indexed address)
func (_Vault *VaultFilterer) FilterApproved(opts *bind.FilterOpts, user []common.Address, spender []common.Address) (*VaultApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Approved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultApprovedIterator{contract: _Vault.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xaad2833c9fd7a3de33f301e5186ee84d1a5753ce32de6b97baedaac4b92b55fc.
//
// Solidity: event Approved(user indexed address, spender indexed address)
func (_Vault *VaultFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *VaultApproved, user []common.Address, spender []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Approved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultApproved)
				if err := _Vault.contract.UnpackLog(event, "Approved", log); err != nil {
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

// VaultDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Vault contract.
type VaultDepositedIterator struct {
	Event *VaultDeposited // Event containing the contract specifics and raw log

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
func (it *VaultDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposited)
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
		it.Event = new(VaultDeposited)
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
func (it *VaultDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposited represents a Deposited event raised by the Vault contract.
type VaultDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_Vault *VaultFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*VaultDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultDepositedIterator{contract: _Vault.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(user indexed address, token address, amount uint256)
func (_Vault *VaultFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *VaultDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposited)
				if err := _Vault.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// VaultRemovedSpenderIterator is returned from FilterRemovedSpender and is used to iterate over the raw logs and unpacked data for RemovedSpender events raised by the Vault contract.
type VaultRemovedSpenderIterator struct {
	Event *VaultRemovedSpender // Event containing the contract specifics and raw log

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
func (it *VaultRemovedSpenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRemovedSpender)
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
		it.Event = new(VaultRemovedSpender)
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
func (it *VaultRemovedSpenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRemovedSpenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRemovedSpender represents a RemovedSpender event raised by the Vault contract.
type VaultRemovedSpender struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedSpender is a free log retrieval operation binding the contract event 0x0e2fc808ab0ead56889f8ff2a8ea0841ba4c0b8311607a902eb24b834857e1b5.
//
// Solidity: event RemovedSpender(spender indexed address)
func (_Vault *VaultFilterer) FilterRemovedSpender(opts *bind.FilterOpts, spender []common.Address) (*VaultRemovedSpenderIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RemovedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultRemovedSpenderIterator{contract: _Vault.contract, event: "RemovedSpender", logs: logs, sub: sub}, nil
}

// WatchRemovedSpender is a free log subscription operation binding the contract event 0x0e2fc808ab0ead56889f8ff2a8ea0841ba4c0b8311607a902eb24b834857e1b5.
//
// Solidity: event RemovedSpender(spender indexed address)
func (_Vault *VaultFilterer) WatchRemovedSpender(opts *bind.WatchOpts, sink chan<- *VaultRemovedSpender, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RemovedSpender", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRemovedSpender)
				if err := _Vault.contract.UnpackLog(event, "RemovedSpender", log); err != nil {
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

// VaultUnapprovedIterator is returned from FilterUnapproved and is used to iterate over the raw logs and unpacked data for Unapproved events raised by the Vault contract.
type VaultUnapprovedIterator struct {
	Event *VaultUnapproved // Event containing the contract specifics and raw log

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
func (it *VaultUnapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnapproved)
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
		it.Event = new(VaultUnapproved)
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
func (it *VaultUnapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnapproved represents a Unapproved event raised by the Vault contract.
type VaultUnapproved struct {
	User    common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnapproved is a free log retrieval operation binding the contract event 0x1ab270601cc6b54dd5e8ce5c70dbac96a01ff12939e4e76488df62adc8e68373.
//
// Solidity: event Unapproved(user indexed address, spender indexed address)
func (_Vault *VaultFilterer) FilterUnapproved(opts *bind.FilterOpts, user []common.Address, spender []common.Address) (*VaultUnapprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unapproved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultUnapprovedIterator{contract: _Vault.contract, event: "Unapproved", logs: logs, sub: sub}, nil
}

// WatchUnapproved is a free log subscription operation binding the contract event 0x1ab270601cc6b54dd5e8ce5c70dbac96a01ff12939e4e76488df62adc8e68373.
//
// Solidity: event Unapproved(user indexed address, spender indexed address)
func (_Vault *VaultFilterer) WatchUnapproved(opts *bind.WatchOpts, sink chan<- *VaultUnapproved, user []common.Address, spender []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unapproved", userRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnapproved)
				if err := _Vault.contract.UnpackLog(event, "Unapproved", log); err != nil {
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

// VaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Vault contract.
type VaultWithdrawnIterator struct {
	Event *VaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdrawn)
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
		it.Event = new(VaultWithdrawn)
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
func (it *VaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdrawn represents a Withdrawn event raised by the Vault contract.
type VaultWithdrawn struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_Vault *VaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*VaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawnIterator{contract: _Vault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(user indexed address, token address, amount uint256)
func (_Vault *VaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *VaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdrawn)
				if err := _Vault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
