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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMembers\",\"inputs\":[{\"name\":\"newMembers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMember\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublisher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedToBroadcastEonKey\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublisher\",\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
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

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Bindings *BindingsCaller) GetMember(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMember", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Bindings *BindingsSession) GetMember(index uint64) (common.Address, error) {
	return _Bindings.Contract.GetMember(&_Bindings.CallOpts, index)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Bindings *BindingsCallerSession) GetMember(index uint64) (common.Address, error) {
	return _Bindings.Contract.GetMember(&_Bindings.CallOpts, index)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Bindings *BindingsCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Bindings *BindingsSession) GetMembers() ([]common.Address, error) {
	return _Bindings.Contract.GetMembers(&_Bindings.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Bindings *BindingsCallerSession) GetMembers() ([]common.Address, error) {
	return _Bindings.Contract.GetMembers(&_Bindings.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Bindings *BindingsCaller) GetNumMembers(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getNumMembers")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Bindings *BindingsSession) GetNumMembers() (uint64, error) {
	return _Bindings.Contract.GetNumMembers(&_Bindings.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Bindings *BindingsCallerSession) GetNumMembers() (uint64, error) {
	return _Bindings.Contract.GetNumMembers(&_Bindings.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Bindings *BindingsCaller) GetPublisher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getPublisher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Bindings *BindingsSession) GetPublisher() (common.Address, error) {
	return _Bindings.Contract.GetPublisher(&_Bindings.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Bindings *BindingsCallerSession) GetPublisher() (common.Address, error) {
	return _Bindings.Contract.GetPublisher(&_Bindings.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Bindings *BindingsCaller) GetThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Bindings *BindingsSession) GetThreshold() (uint64, error) {
	return _Bindings.Contract.GetThreshold(&_Bindings.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Bindings *BindingsCallerSession) GetThreshold() (uint64, error) {
	return _Bindings.Contract.GetThreshold(&_Bindings.CallOpts)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Bindings *BindingsCaller) IsAllowedToBroadcastEonKey(opts *bind.CallOpts, a common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isAllowedToBroadcastEonKey", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Bindings *BindingsSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _Bindings.Contract.IsAllowedToBroadcastEonKey(&_Bindings.CallOpts, a)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Bindings *BindingsCallerSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _Bindings.Contract.IsAllowedToBroadcastEonKey(&_Bindings.CallOpts, a)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Bindings *BindingsCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Bindings *BindingsSession) IsFinalized() (bool, error) {
	return _Bindings.Contract.IsFinalized(&_Bindings.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Bindings *BindingsCallerSession) IsFinalized() (bool, error) {
	return _Bindings.Contract.IsFinalized(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Bindings *BindingsTransactor) AddMembers(opts *bind.TransactOpts, newMembers []common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addMembers", newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Bindings *BindingsSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddMembers(&_Bindings.TransactOpts, newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Bindings *BindingsTransactorSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddMembers(&_Bindings.TransactOpts, newMembers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Bindings *BindingsTransactor) SetFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setFinalized")
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Bindings *BindingsSession) SetFinalized() (*types.Transaction, error) {
	return _Bindings.Contract.SetFinalized(&_Bindings.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Bindings *BindingsTransactorSession) SetFinalized() (*types.Transaction, error) {
	return _Bindings.Contract.SetFinalized(&_Bindings.TransactOpts)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Bindings *BindingsTransactor) SetPublisher(opts *bind.TransactOpts, _publisher common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPublisher", _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Bindings *BindingsSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPublisher(&_Bindings.TransactOpts, _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Bindings *BindingsTransactorSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPublisher(&_Bindings.TransactOpts, _publisher)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Bindings *BindingsTransactor) SetThreshold(opts *bind.TransactOpts, _threshold uint64) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Bindings *BindingsSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _Bindings.Contract.SetThreshold(&_Bindings.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Bindings *BindingsTransactorSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _Bindings.Contract.SetThreshold(&_Bindings.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
