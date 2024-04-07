// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyperSet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_broadcaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eonKeyConfirmed\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishEonKey\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keyperId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EonVoteRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Bindings *BindingsCaller) EonKeyConfirmed(opts *bind.CallOpts, eonKey []byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "eonKeyConfirmed", eonKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Bindings *BindingsSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Bindings.Contract.EonKeyConfirmed(&_Bindings.CallOpts, eonKey)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Bindings *BindingsCallerSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Bindings.Contract.EonKeyConfirmed(&_Bindings.CallOpts, eonKey)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Bindings *BindingsTransactor) PublishEonKey(opts *bind.TransactOpts, eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "publishEonKey", eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Bindings *BindingsSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Bindings.Contract.PublishEonKey(&_Bindings.TransactOpts, eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Bindings *BindingsTransactorSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Bindings.Contract.PublishEonKey(&_Bindings.TransactOpts, eonKey, keyperId)
}

// BindingsEonVoteRegisteredIterator is returned from FilterEonVoteRegistered and is used to iterate over the raw logs and unpacked data for EonVoteRegistered events raised by the Bindings contract.
type BindingsEonVoteRegisteredIterator struct {
	Event *BindingsEonVoteRegistered // Event containing the contract specifics and raw log

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
func (it *BindingsEonVoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEonVoteRegistered)
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
		it.Event = new(BindingsEonVoteRegistered)
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
func (it *BindingsEonVoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEonVoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEonVoteRegistered represents a EonVoteRegistered event raised by the Bindings contract.
type BindingsEonVoteRegistered struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonVoteRegistered is a free log retrieval operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) FilterEonVoteRegistered(opts *bind.FilterOpts) (*BindingsEonVoteRegisteredIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return &BindingsEonVoteRegisteredIterator{contract: _Bindings.contract, event: "EonVoteRegistered", logs: logs, sub: sub}, nil
}

// WatchEonVoteRegistered is a free log subscription operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) WatchEonVoteRegistered(opts *bind.WatchOpts, sink chan<- *BindingsEonVoteRegistered) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEonVoteRegistered)
				if err := _Bindings.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
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

// ParseEonVoteRegistered is a log parse operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) ParseEonVoteRegistered(log types.Log) (*BindingsEonVoteRegistered, error) {
	event := new(BindingsEonVoteRegistered)
	if err := _Bindings.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
