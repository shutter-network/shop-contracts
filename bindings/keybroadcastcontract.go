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

// KeyBroadcastContractMetaData contains all meta data concerning the KeyBroadcastContract contract.
var KeyBroadcastContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyperSetManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyHaveKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"eon\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"EonKeyBroadcast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"eon\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"broadcastEonKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"eon\",\"type\":\"uint64\"}],\"name\":\"getEonKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106e53803806106e583398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610652806100936000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638a0b8b281461003b578063daade8e814610064575b600080fd5b61004e610049366004610308565b610079565b60405161005b9190610370565b60405180910390f35b610077610072366004610399565b610126565b005b67ffffffffffffffff811660009081526020819052604090208054606091906100a19061045b565b80601f01602080910402602001604051908101604052809291908181526020018280546100cd9061045b565b801561011a5780601f106100ef5761010080835404028352916020019161011a565b820191906000526020600020905b8154815290600101906020018083116100fd57829003601f168201915b50505050509050919050565b805160000361014857604051630eda9c3d60e31b815260040160405180910390fd5b67ffffffffffffffff82166000908152602081905260408120805461016c9061045b565b9050111561018d57604051632c6c0cd160e11b815260040160405180910390fd5b60015460405163f90f3bed60e01b815267ffffffffffffffff841660048201526001600160a01b039091169063f90f3bed90602401602060405180830381865afa1580156101df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102039190610495565b60405163cde1532d60e01b81523360048201526001600160a01b03919091169063cde1532d90602401602060405180830381865afa158015610249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026d91906104be565b61028a57604051631eb49d6d60e11b815260040160405180910390fd5b67ffffffffffffffff821660009081526020819052604090206102ad8282610531565b507ff0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad376882826040516102df9291906105f1565b60405180910390a15050565b803567ffffffffffffffff8116811461030357600080fd5b919050565b60006020828403121561031a57600080fd5b610323826102eb565b9392505050565b6000815180845260005b8181101561035057602081850181015186830182015201610334565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610323602083018461032a565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156103ac57600080fd5b6103b5836102eb565b9150602083013567ffffffffffffffff808211156103d257600080fd5b818501915085601f8301126103e657600080fd5b8135818111156103f8576103f8610383565b604051601f8201601f19908116603f0116810190838211818310171561042057610420610383565b8160405282815288602084870101111561043957600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600181811c9082168061046f57607f821691505b60208210810361048f57634e487b7160e01b600052602260045260246000fd5b50919050565b6000602082840312156104a757600080fd5b81516001600160a01b038116811461032357600080fd5b6000602082840312156104d057600080fd5b8151801515811461032357600080fd5b601f82111561052c576000816000526020600020601f850160051c810160208610156105095750805b601f850160051c820191505b8181101561052857828155600101610515565b5050505b505050565b815167ffffffffffffffff81111561054b5761054b610383565b61055f81610559845461045b565b846104e0565b602080601f831160018114610594576000841561057c5750858301515b600019600386901b1c1916600185901b178555610528565b600085815260208120601f198616915b828110156105c3578886015182559484019460019091019084016105a4565b50858210156105e15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff83168152604060208201526000610614604083018461032a565b94935050505056fea26469706673582212201261e516c15694af91d70196bef7e20320399412234a431ca4739ebcd565db1764736f6c63430008160033",
}

// KeyBroadcastContractABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyBroadcastContractMetaData.ABI instead.
var KeyBroadcastContractABI = KeyBroadcastContractMetaData.ABI

// KeyBroadcastContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyBroadcastContractMetaData.Bin instead.
var KeyBroadcastContractBin = KeyBroadcastContractMetaData.Bin

// DeployKeyBroadcastContract deploys a new Ethereum contract, binding an instance of KeyBroadcastContract to it.
func DeployKeyBroadcastContract(auth *bind.TransactOpts, backend bind.ContractBackend, keyperSetManagerAddress common.Address) (common.Address, *types.Transaction, *KeyBroadcastContract, error) {
	parsed, err := KeyBroadcastContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyBroadcastContractBin), backend, keyperSetManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyBroadcastContract{KeyBroadcastContractCaller: KeyBroadcastContractCaller{contract: contract}, KeyBroadcastContractTransactor: KeyBroadcastContractTransactor{contract: contract}, KeyBroadcastContractFilterer: KeyBroadcastContractFilterer{contract: contract}}, nil
}

// KeyBroadcastContract is an auto generated Go binding around an Ethereum contract.
type KeyBroadcastContract struct {
	KeyBroadcastContractCaller     // Read-only binding to the contract
	KeyBroadcastContractTransactor // Write-only binding to the contract
	KeyBroadcastContractFilterer   // Log filterer for contract events
}

// KeyBroadcastContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyBroadcastContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyBroadcastContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyBroadcastContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyBroadcastContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyBroadcastContractSession struct {
	Contract     *KeyBroadcastContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeyBroadcastContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyBroadcastContractCallerSession struct {
	Contract *KeyBroadcastContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// KeyBroadcastContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyBroadcastContractTransactorSession struct {
	Contract     *KeyBroadcastContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// KeyBroadcastContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyBroadcastContractRaw struct {
	Contract *KeyBroadcastContract // Generic contract binding to access the raw methods on
}

// KeyBroadcastContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyBroadcastContractCallerRaw struct {
	Contract *KeyBroadcastContractCaller // Generic read-only contract binding to access the raw methods on
}

// KeyBroadcastContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyBroadcastContractTransactorRaw struct {
	Contract *KeyBroadcastContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyBroadcastContract creates a new instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContract(address common.Address, backend bind.ContractBackend) (*KeyBroadcastContract, error) {
	contract, err := bindKeyBroadcastContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContract{KeyBroadcastContractCaller: KeyBroadcastContractCaller{contract: contract}, KeyBroadcastContractTransactor: KeyBroadcastContractTransactor{contract: contract}, KeyBroadcastContractFilterer: KeyBroadcastContractFilterer{contract: contract}}, nil
}

// NewKeyBroadcastContractCaller creates a new read-only instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractCaller(address common.Address, caller bind.ContractCaller) (*KeyBroadcastContractCaller, error) {
	contract, err := bindKeyBroadcastContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractCaller{contract: contract}, nil
}

// NewKeyBroadcastContractTransactor creates a new write-only instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyBroadcastContractTransactor, error) {
	contract, err := bindKeyBroadcastContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractTransactor{contract: contract}, nil
}

// NewKeyBroadcastContractFilterer creates a new log filterer instance of KeyBroadcastContract, bound to a specific deployed contract.
func NewKeyBroadcastContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyBroadcastContractFilterer, error) {
	contract, err := bindKeyBroadcastContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractFilterer{contract: contract}, nil
}

// bindKeyBroadcastContract binds a generic wrapper to an already deployed contract.
func bindKeyBroadcastContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyBroadcastContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyBroadcastContract *KeyBroadcastContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.KeyBroadcastContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyBroadcastContract *KeyBroadcastContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyBroadcastContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyBroadcastContract *KeyBroadcastContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyBroadcastContract *KeyBroadcastContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.contract.Transact(opts, method, params...)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractCaller) GetEonKey(opts *bind.CallOpts, eon uint64) ([]byte, error) {
	var out []interface{}
	err := _KeyBroadcastContract.contract.Call(opts, &out, "getEonKey", eon)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractSession) GetEonKey(eon uint64) ([]byte, error) {
	return _KeyBroadcastContract.Contract.GetEonKey(&_KeyBroadcastContract.CallOpts, eon)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_KeyBroadcastContract *KeyBroadcastContractCallerSession) GetEonKey(eon uint64) ([]byte, error) {
	return _KeyBroadcastContract.Contract.GetEonKey(&_KeyBroadcastContract.CallOpts, eon)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactor) BroadcastEonKey(opts *bind.TransactOpts, eon uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.contract.Transact(opts, "broadcastEonKey", eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.BroadcastEonKey(&_KeyBroadcastContract.TransactOpts, eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_KeyBroadcastContract *KeyBroadcastContractTransactorSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _KeyBroadcastContract.Contract.BroadcastEonKey(&_KeyBroadcastContract.TransactOpts, eon, key)
}

// KeyBroadcastContractEonKeyBroadcastIterator is returned from FilterEonKeyBroadcast and is used to iterate over the raw logs and unpacked data for EonKeyBroadcast events raised by the KeyBroadcastContract contract.
type KeyBroadcastContractEonKeyBroadcastIterator struct {
	Event *KeyBroadcastContractEonKeyBroadcast // Event containing the contract specifics and raw log

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
func (it *KeyBroadcastContractEonKeyBroadcastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyBroadcastContractEonKeyBroadcast)
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
		it.Event = new(KeyBroadcastContractEonKeyBroadcast)
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
func (it *KeyBroadcastContractEonKeyBroadcastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyBroadcastContractEonKeyBroadcastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyBroadcastContractEonKeyBroadcast represents a EonKeyBroadcast event raised by the KeyBroadcastContract contract.
type KeyBroadcastContractEonKeyBroadcast struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonKeyBroadcast is a free log retrieval operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) FilterEonKeyBroadcast(opts *bind.FilterOpts) (*KeyBroadcastContractEonKeyBroadcastIterator, error) {

	logs, sub, err := _KeyBroadcastContract.contract.FilterLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return &KeyBroadcastContractEonKeyBroadcastIterator{contract: _KeyBroadcastContract.contract, event: "EonKeyBroadcast", logs: logs, sub: sub}, nil
}

// WatchEonKeyBroadcast is a free log subscription operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) WatchEonKeyBroadcast(opts *bind.WatchOpts, sink chan<- *KeyBroadcastContractEonKeyBroadcast) (event.Subscription, error) {

	logs, sub, err := _KeyBroadcastContract.contract.WatchLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyBroadcastContractEonKeyBroadcast)
				if err := _KeyBroadcastContract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
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

// ParseEonKeyBroadcast is a log parse operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_KeyBroadcastContract *KeyBroadcastContractFilterer) ParseEonKeyBroadcast(log types.Log) (*KeyBroadcastContractEonKeyBroadcast, error) {
	event := new(KeyBroadcastContractEonKeyBroadcast)
	if err := _KeyBroadcastContract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
