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

// KeyperSetMetaData contains all meta data concerning the KeyperSet contract.
var KeyperSetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newMembers\",\"type\":\"address[]\"}],\"name\":\"addMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumMembers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"isAllowedToBroadcastEonKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFinalized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_broadcaster\",\"type\":\"address\"}],\"name\":\"setKeyBroadcaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_threshold\",\"type\":\"uint64\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106a7806100a56000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063715018a61161008c5780639eab5253116100665780639eab52531461019b578063cde1532d146101b0578063e75235b8146101d9578063f2fde38b146101eb57600080fd5b8063715018a6146101645780638d4e40831461016c5780638da5cb5b1461018a57600080fd5b806317c4de35146100d457806317d5430a146100e95780631de772531461010b5780632e8e6cad146101135780636a33e20e1461013e5780636f4d469b14610151575b600080fd5b6100e76100e236600461050a565b6101fe565b005b6001545b60405167ffffffffffffffff90911681526020015b60405180910390f35b6100e7610255565b61012661012136600461050a565b610272565b6040516001600160a01b039091168152602001610102565b6100e761014c36600461053b565b6102ac565b6100e761015f366004610564565b610310565b6100e76103d4565b600054600160a01b900460ff165b6040519015158152602001610102565b6000546001600160a01b0316610126565b6101a36103e8565b60405161010291906105d9565b61017a6101be36600461053b565b600254600160401b90046001600160a01b0390811691161490565b60025467ffffffffffffffff166100ed565b6100e76101f936600461053b565b61044a565b61020661048d565b600054600160a01b900460ff16156102315760405163475a253560e01b815260040160405180910390fd5b6002805467ffffffffffffffff191667ffffffffffffffff92909216919091179055565b61025d61048d565b6000805460ff60a01b1916600160a01b179055565b600060018267ffffffffffffffff168154811061029157610291610626565b6000918252602090912001546001600160a01b031692915050565b6102b461048d565b600054600160a01b900460ff16156102df5760405163475a253560e01b815260040160405180910390fd5b600280546001600160a01b03909216600160401b0268010000000000000000600160e01b0319909216919091179055565b61031861048d565b600054600160a01b900460ff16156103435760405163475a253560e01b815260040160405180910390fd5b60005b67ffffffffffffffff81168211156103cf57600183838367ffffffffffffffff1681811061037657610376610626565b905060200201602081019061038b919061053b565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055806103c78161063c565b915050610346565b505050565b6103dc61048d565b6103e660006104ba565b565b6060600180548060200260200160405190810160405280929190818152602001828054801561044057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610422575b5050505050905090565b61045261048d565b6001600160a01b03811661048157604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61048a816104ba565b50565b6000546001600160a01b031633146103e65760405163118cdaa760e01b8152336004820152602401610478565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561051c57600080fd5b813567ffffffffffffffff8116811461053457600080fd5b9392505050565b60006020828403121561054d57600080fd5b81356001600160a01b038116811461053457600080fd5b6000806020838503121561057757600080fd5b823567ffffffffffffffff8082111561058f57600080fd5b818501915085601f8301126105a357600080fd5b8135818111156105b257600080fd5b8660208260051b85010111156105c757600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b8181101561061a5783516001600160a01b0316835292840192918401916001016105f5565b50909695505050505050565b634e487b7160e01b600052603260045260246000fd5b600067ffffffffffffffff80831681810361066757634e487b7160e01b600052601160045260246000fd5b600101939250505056fea2646970667358221220ee8f22a9e07fcc081bc196818f44da5742af55729008eaeff013c4e2bbf44e2c64736f6c63430008160033",
}

// KeyperSetABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyperSetMetaData.ABI instead.
var KeyperSetABI = KeyperSetMetaData.ABI

// KeyperSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyperSetMetaData.Bin instead.
var KeyperSetBin = KeyperSetMetaData.Bin

// DeployKeyperSet deploys a new Ethereum contract, binding an instance of KeyperSet to it.
func DeployKeyperSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyperSet, error) {
	parsed, err := KeyperSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyperSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyperSet{KeyperSetCaller: KeyperSetCaller{contract: contract}, KeyperSetTransactor: KeyperSetTransactor{contract: contract}, KeyperSetFilterer: KeyperSetFilterer{contract: contract}}, nil
}

// KeyperSet is an auto generated Go binding around an Ethereum contract.
type KeyperSet struct {
	KeyperSetCaller     // Read-only binding to the contract
	KeyperSetTransactor // Write-only binding to the contract
	KeyperSetFilterer   // Log filterer for contract events
}

// KeyperSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyperSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyperSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyperSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyperSetSession struct {
	Contract     *KeyperSet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyperSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyperSetCallerSession struct {
	Contract *KeyperSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KeyperSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyperSetTransactorSession struct {
	Contract     *KeyperSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KeyperSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyperSetRaw struct {
	Contract *KeyperSet // Generic contract binding to access the raw methods on
}

// KeyperSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyperSetCallerRaw struct {
	Contract *KeyperSetCaller // Generic read-only contract binding to access the raw methods on
}

// KeyperSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyperSetTransactorRaw struct {
	Contract *KeyperSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyperSet creates a new instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSet(address common.Address, backend bind.ContractBackend) (*KeyperSet, error) {
	contract, err := bindKeyperSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyperSet{KeyperSetCaller: KeyperSetCaller{contract: contract}, KeyperSetTransactor: KeyperSetTransactor{contract: contract}, KeyperSetFilterer: KeyperSetFilterer{contract: contract}}, nil
}

// NewKeyperSetCaller creates a new read-only instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetCaller(address common.Address, caller bind.ContractCaller) (*KeyperSetCaller, error) {
	contract, err := bindKeyperSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetCaller{contract: contract}, nil
}

// NewKeyperSetTransactor creates a new write-only instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyperSetTransactor, error) {
	contract, err := bindKeyperSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetTransactor{contract: contract}, nil
}

// NewKeyperSetFilterer creates a new log filterer instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyperSetFilterer, error) {
	contract, err := bindKeyperSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyperSetFilterer{contract: contract}, nil
}

// bindKeyperSet binds a generic wrapper to an already deployed contract.
func bindKeyperSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyperSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSet *KeyperSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSet.Contract.KeyperSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSet *KeyperSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.Contract.KeyperSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSet *KeyperSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSet.Contract.KeyperSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSet *KeyperSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSet *KeyperSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSet *KeyperSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSet.Contract.contract.Transact(opts, method, params...)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetCaller) GetMember(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getMember", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetSession) GetMember(index uint64) (common.Address, error) {
	return _KeyperSet.Contract.GetMember(&_KeyperSet.CallOpts, index)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetCallerSession) GetMember(index uint64) (common.Address, error) {
	return _KeyperSet.Contract.GetMember(&_KeyperSet.CallOpts, index)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetSession) GetMembers() ([]common.Address, error) {
	return _KeyperSet.Contract.GetMembers(&_KeyperSet.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetCallerSession) GetMembers() ([]common.Address, error) {
	return _KeyperSet.Contract.GetMembers(&_KeyperSet.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetCaller) GetNumMembers(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getNumMembers")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetSession) GetNumMembers() (uint64, error) {
	return _KeyperSet.Contract.GetNumMembers(&_KeyperSet.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetCallerSession) GetNumMembers() (uint64, error) {
	return _KeyperSet.Contract.GetNumMembers(&_KeyperSet.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetCaller) GetThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetSession) GetThreshold() (uint64, error) {
	return _KeyperSet.Contract.GetThreshold(&_KeyperSet.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetCallerSession) GetThreshold() (uint64, error) {
	return _KeyperSet.Contract.GetThreshold(&_KeyperSet.CallOpts)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetCaller) IsAllowedToBroadcastEonKey(opts *bind.CallOpts, a common.Address) (bool, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "isAllowedToBroadcastEonKey", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _KeyperSet.Contract.IsAllowedToBroadcastEonKey(&_KeyperSet.CallOpts, a)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetCallerSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _KeyperSet.Contract.IsAllowedToBroadcastEonKey(&_KeyperSet.CallOpts, a)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetSession) IsFinalized() (bool, error) {
	return _KeyperSet.Contract.IsFinalized(&_KeyperSet.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetCallerSession) IsFinalized() (bool, error) {
	return _KeyperSet.Contract.IsFinalized(&_KeyperSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetSession) Owner() (common.Address, error) {
	return _KeyperSet.Contract.Owner(&_KeyperSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetCallerSession) Owner() (common.Address, error) {
	return _KeyperSet.Contract.Owner(&_KeyperSet.CallOpts)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetTransactor) AddMembers(opts *bind.TransactOpts, newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "addMembers", newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.AddMembers(&_KeyperSet.TransactOpts, newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetTransactorSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.AddMembers(&_KeyperSet.TransactOpts, newMembers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyperSet.Contract.RenounceOwnership(&_KeyperSet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyperSet.Contract.RenounceOwnership(&_KeyperSet.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetTransactor) SetFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setFinalized")
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetSession) SetFinalized() (*types.Transaction, error) {
	return _KeyperSet.Contract.SetFinalized(&_KeyperSet.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetTransactorSession) SetFinalized() (*types.Transaction, error) {
	return _KeyperSet.Contract.SetFinalized(&_KeyperSet.TransactOpts)
}

// SetKeyBroadcaster is a paid mutator transaction binding the contract method 0x6a33e20e.
//
// Solidity: function setKeyBroadcaster(address _broadcaster) returns()
func (_KeyperSet *KeyperSetTransactor) SetKeyBroadcaster(opts *bind.TransactOpts, _broadcaster common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setKeyBroadcaster", _broadcaster)
}

// SetKeyBroadcaster is a paid mutator transaction binding the contract method 0x6a33e20e.
//
// Solidity: function setKeyBroadcaster(address _broadcaster) returns()
func (_KeyperSet *KeyperSetSession) SetKeyBroadcaster(_broadcaster common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetKeyBroadcaster(&_KeyperSet.TransactOpts, _broadcaster)
}

// SetKeyBroadcaster is a paid mutator transaction binding the contract method 0x6a33e20e.
//
// Solidity: function setKeyBroadcaster(address _broadcaster) returns()
func (_KeyperSet *KeyperSetTransactorSession) SetKeyBroadcaster(_broadcaster common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetKeyBroadcaster(&_KeyperSet.TransactOpts, _broadcaster)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetTransactor) SetThreshold(opts *bind.TransactOpts, _threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetThreshold(&_KeyperSet.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetTransactorSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetThreshold(&_KeyperSet.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.TransferOwnership(&_KeyperSet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.TransferOwnership(&_KeyperSet.TransactOpts, newOwner)
}

// KeyperSetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyperSet contract.
type KeyperSetOwnershipTransferredIterator struct {
	Event *KeyperSetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyperSetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetOwnershipTransferred)
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
		it.Event = new(KeyperSetOwnershipTransferred)
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
func (it *KeyperSetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetOwnershipTransferred represents a OwnershipTransferred event raised by the KeyperSet contract.
type KeyperSetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyperSet *KeyperSetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyperSetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyperSet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSetOwnershipTransferredIterator{contract: _KeyperSet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyperSet *KeyperSetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyperSetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyperSet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetOwnershipTransferred)
				if err := _KeyperSet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeyperSet *KeyperSetFilterer) ParseOwnershipTransferred(log types.Log) (*KeyperSetOwnershipTransferred, error) {
	event := new(KeyperSetOwnershipTransferred)
	if err := _KeyperSet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
