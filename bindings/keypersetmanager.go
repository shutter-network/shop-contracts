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

// KeyperSetManagerMetaData contains all meta data concerning the KeyperSetManager contract.
var KeyperSetManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminRoleAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyHaveKeyperSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyperSetNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveKeyperSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"activationBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keyperSetContract\",\"type\":\"address\"}],\"name\":\"KeyperSetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"activationBlock\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"keyperSetContract\",\"type\":\"address\"}],\"name\":\"addKeyperSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getKeyperSetActivationBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getKeyperSetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"getKeyperSetIndexByBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumKeyperSets\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b97380380610b9783398101604081905261002f916100f9565b6001805460ff191690558061004560008261004d565b505050610129565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166100ef576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556100a73390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016100f3565b5060005b92915050565b60006020828403121561010b57600080fd5b81516001600160a01b038116811461012257600080fd5b9392505050565b610a5f806101386000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638456cb5911610097578063d547741f11610066578063d547741f1461020e578063e63ab1e914610221578063f2e6100a14610248578063f90f3bed1461025057600080fd5b80638456cb59146101d857806391d14854146101e0578063a217fddf146101f3578063d3877c43146101fb57600080fd5b806336568abe116100d357806336568abe1461019f5780633f4ba83a146101b25780635c975abb146101ba578063636df979146101c557600080fd5b806301ffc9a714610105578063035cef151461012d578063248a9ca3146101595780632f2ff15d1461018a575b600080fd5b6101186101133660046108b6565b61027b565b60405190151581526020015b60405180910390f35b61014061013b3660046108fd565b6102b2565b60405167ffffffffffffffff9091168152602001610124565b61017c610167366004610918565b60009081526020819052604090206001015490565b604051908152602001610124565b61019d610198366004610948565b610340565b005b61019d6101ad366004610948565b61036b565b61019d6103a3565b60015460ff16610118565b6101406101d33660046108fd565b6103b9565b61019d6103f4565b6101186101ee366004610948565b610426565b61017c600081565b61019d610209366004610974565b61044f565b61019d61021c366004610948565b61061c565b61017c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b600254610140565b61026361025e3660046108fd565b610641565b6040516001600160a01b039091168152602001610124565b60006001600160e01b03198216637965db0b60e01b14806102ac57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6002546000905b80156103265767ffffffffffffffff831660026102d76001846109b4565b815481106102e7576102e76109c7565b60009182526020909120015467ffffffffffffffff16116103145761030d6001826109b4565b9392505050565b8061031e816109dd565b9150506102b9565b506040516367c9fd1d60e11b815260040160405180910390fd5b60008281526020819052604090206001015461035b81610682565b610365838361068c565b50505050565b6001600160a01b03811633146103945760405163334bd91960e11b815260040160405180910390fd5b61039e828261071e565b505050565b60006103ae81610682565b6103b6610789565b50565b600060028267ffffffffffffffff16815481106103d8576103d86109c7565b60009182526020909120015467ffffffffffffffff1692915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61041e81610682565b6103b66107db565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b600061045a81610682565b600254158015906104be5750600280546104b1919061047b906001906109b4565b8154811061048b5761048b6109c7565b60009182526020909120015467ffffffffffffffff166104ac4360016109f4565b610816565b8367ffffffffffffffff16105b156104dc576040516361cb74ab60e11b815260040160405180910390fd5b816001600160a01b0316638d4e40836040518163ffffffff1660e01b8152600401602060405180830381865afa15801561051a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053e9190610a07565b61055b5760405163756318d560e11b815260040160405180910390fd5b60408051808201825267ffffffffffffffff8581168083526001600160a01b0386811660208086018281526002805460018101825560009190915296517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90970180549151909416600160401b026001600160e01b031990911696909516959095179390931790558351908152918201527f6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b910160405180910390a1505050565b60008281526020819052604090206001015461063781610682565b610365838361071e565b600060028267ffffffffffffffff1681548110610660576106606109c7565b600091825260209091200154600160401b90046001600160a01b031692915050565b6103b6813361082c565b60006106988383610426565b610716576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556106ce3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102ac565b5060006102ac565b600061072a8383610426565b15610716576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016102ac565b61079161086d565b6001805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6107e3610892565b6001805460ff1916811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336107be565b6000818311610825578161030d565b5090919050565b6108368282610426565b6108695760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b5050565b60015460ff1661089057604051638dfc202b60e01b815260040160405180910390fd5b565b60015460ff16156108905760405163d93c066560e01b815260040160405180910390fd5b6000602082840312156108c857600080fd5b81356001600160e01b03198116811461030d57600080fd5b803567ffffffffffffffff811681146108f857600080fd5b919050565b60006020828403121561090f57600080fd5b61030d826108e0565b60006020828403121561092a57600080fd5b5035919050565b80356001600160a01b03811681146108f857600080fd5b6000806040838503121561095b57600080fd5b8235915061096b60208401610931565b90509250929050565b6000806040838503121561098757600080fd5b610990836108e0565b915061096b60208401610931565b634e487b7160e01b600052601160045260246000fd5b818103818111156102ac576102ac61099e565b634e487b7160e01b600052603260045260246000fd5b6000816109ec576109ec61099e565b506000190190565b808201808211156102ac576102ac61099e565b600060208284031215610a1957600080fd5b8151801515811461030d57600080fdfea2646970667358221220be728f61c0f5a9b1bd9fa50bb1bd41d4484ba32e65526442dc09b94ab5da0bd264736f6c63430008160033",
}

// KeyperSetManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyperSetManagerMetaData.ABI instead.
var KeyperSetManagerABI = KeyperSetManagerMetaData.ABI

// KeyperSetManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyperSetManagerMetaData.Bin instead.
var KeyperSetManagerBin = KeyperSetManagerMetaData.Bin

// DeployKeyperSetManager deploys a new Ethereum contract, binding an instance of KeyperSetManager to it.
func DeployKeyperSetManager(auth *bind.TransactOpts, backend bind.ContractBackend, _adminRoleAddress common.Address) (common.Address, *types.Transaction, *KeyperSetManager, error) {
	parsed, err := KeyperSetManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyperSetManagerBin), backend, _adminRoleAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyperSetManager{KeyperSetManagerCaller: KeyperSetManagerCaller{contract: contract}, KeyperSetManagerTransactor: KeyperSetManagerTransactor{contract: contract}, KeyperSetManagerFilterer: KeyperSetManagerFilterer{contract: contract}}, nil
}

// KeyperSetManager is an auto generated Go binding around an Ethereum contract.
type KeyperSetManager struct {
	KeyperSetManagerCaller     // Read-only binding to the contract
	KeyperSetManagerTransactor // Write-only binding to the contract
	KeyperSetManagerFilterer   // Log filterer for contract events
}

// KeyperSetManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyperSetManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyperSetManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyperSetManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyperSetManagerSession struct {
	Contract     *KeyperSetManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyperSetManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyperSetManagerCallerSession struct {
	Contract *KeyperSetManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KeyperSetManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyperSetManagerTransactorSession struct {
	Contract     *KeyperSetManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KeyperSetManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyperSetManagerRaw struct {
	Contract *KeyperSetManager // Generic contract binding to access the raw methods on
}

// KeyperSetManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyperSetManagerCallerRaw struct {
	Contract *KeyperSetManagerCaller // Generic read-only contract binding to access the raw methods on
}

// KeyperSetManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyperSetManagerTransactorRaw struct {
	Contract *KeyperSetManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyperSetManager creates a new instance of KeyperSetManager, bound to a specific deployed contract.
func NewKeyperSetManager(address common.Address, backend bind.ContractBackend) (*KeyperSetManager, error) {
	contract, err := bindKeyperSetManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManager{KeyperSetManagerCaller: KeyperSetManagerCaller{contract: contract}, KeyperSetManagerTransactor: KeyperSetManagerTransactor{contract: contract}, KeyperSetManagerFilterer: KeyperSetManagerFilterer{contract: contract}}, nil
}

// NewKeyperSetManagerCaller creates a new read-only instance of KeyperSetManager, bound to a specific deployed contract.
func NewKeyperSetManagerCaller(address common.Address, caller bind.ContractCaller) (*KeyperSetManagerCaller, error) {
	contract, err := bindKeyperSetManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerCaller{contract: contract}, nil
}

// NewKeyperSetManagerTransactor creates a new write-only instance of KeyperSetManager, bound to a specific deployed contract.
func NewKeyperSetManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyperSetManagerTransactor, error) {
	contract, err := bindKeyperSetManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerTransactor{contract: contract}, nil
}

// NewKeyperSetManagerFilterer creates a new log filterer instance of KeyperSetManager, bound to a specific deployed contract.
func NewKeyperSetManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyperSetManagerFilterer, error) {
	contract, err := bindKeyperSetManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerFilterer{contract: contract}, nil
}

// bindKeyperSetManager binds a generic wrapper to an already deployed contract.
func bindKeyperSetManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyperSetManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSetManager *KeyperSetManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSetManager.Contract.KeyperSetManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSetManager *KeyperSetManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.KeyperSetManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSetManager *KeyperSetManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.KeyperSetManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSetManager *KeyperSetManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSetManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSetManager *KeyperSetManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSetManager *KeyperSetManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KeyperSetManager.Contract.DEFAULTADMINROLE(&_KeyperSetManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KeyperSetManager.Contract.DEFAULTADMINROLE(&_KeyperSetManager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerSession) PAUSERROLE() ([32]byte, error) {
	return _KeyperSetManager.Contract.PAUSERROLE(&_KeyperSetManager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _KeyperSetManager.Contract.PAUSERROLE(&_KeyperSetManager.CallOpts)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCaller) GetKeyperSetActivationBlock(opts *bind.CallOpts, index uint64) (uint64, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "getKeyperSetActivationBlock", index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _KeyperSetManager.Contract.GetKeyperSetActivationBlock(&_KeyperSetManager.CallOpts, index)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCallerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _KeyperSetManager.Contract.GetKeyperSetActivationBlock(&_KeyperSetManager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_KeyperSetManager *KeyperSetManagerCaller) GetKeyperSetAddress(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "getKeyperSetAddress", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_KeyperSetManager *KeyperSetManagerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _KeyperSetManager.Contract.GetKeyperSetAddress(&_KeyperSetManager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_KeyperSetManager *KeyperSetManagerCallerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _KeyperSetManager.Contract.GetKeyperSetAddress(&_KeyperSetManager.CallOpts, index)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCaller) GetKeyperSetIndexByBlock(opts *bind.CallOpts, blockNumber uint64) (uint64, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "getKeyperSetIndexByBlock", blockNumber)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _KeyperSetManager.Contract.GetKeyperSetIndexByBlock(&_KeyperSetManager.CallOpts, blockNumber)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCallerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _KeyperSetManager.Contract.GetKeyperSetIndexByBlock(&_KeyperSetManager.CallOpts, blockNumber)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCaller) GetNumKeyperSets(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "getNumKeyperSets")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerSession) GetNumKeyperSets() (uint64, error) {
	return _KeyperSetManager.Contract.GetNumKeyperSets(&_KeyperSetManager.CallOpts)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_KeyperSetManager *KeyperSetManagerCallerSession) GetNumKeyperSets() (uint64, error) {
	return _KeyperSetManager.Contract.GetNumKeyperSets(&_KeyperSetManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KeyperSetManager.Contract.GetRoleAdmin(&_KeyperSetManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeyperSetManager *KeyperSetManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KeyperSetManager.Contract.GetRoleAdmin(&_KeyperSetManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KeyperSetManager.Contract.HasRole(&_KeyperSetManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KeyperSetManager.Contract.HasRole(&_KeyperSetManager.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyperSetManager *KeyperSetManagerSession) Paused() (bool, error) {
	return _KeyperSetManager.Contract.Paused(&_KeyperSetManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCallerSession) Paused() (bool, error) {
	return _KeyperSetManager.Contract.Paused(&_KeyperSetManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KeyperSetManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KeyperSetManager.Contract.SupportsInterface(&_KeyperSetManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KeyperSetManager *KeyperSetManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KeyperSetManager.Contract.SupportsInterface(&_KeyperSetManager.CallOpts, interfaceId)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) AddKeyperSet(opts *bind.TransactOpts, activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "addKeyperSet", activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_KeyperSetManager *KeyperSetManagerSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.AddKeyperSet(&_KeyperSetManager.TransactOpts, activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.AddKeyperSet(&_KeyperSetManager.TransactOpts, activationBlock, keyperSetContract)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.GrantRole(&_KeyperSetManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.GrantRole(&_KeyperSetManager.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyperSetManager *KeyperSetManagerSession) Pause() (*types.Transaction, error) {
	return _KeyperSetManager.Contract.Pause(&_KeyperSetManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _KeyperSetManager.Contract.Pause(&_KeyperSetManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_KeyperSetManager *KeyperSetManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.RenounceRole(&_KeyperSetManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.RenounceRole(&_KeyperSetManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.RevokeRole(&_KeyperSetManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeyperSetManager.Contract.RevokeRole(&_KeyperSetManager.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyperSetManager *KeyperSetManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSetManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyperSetManager *KeyperSetManagerSession) Unpause() (*types.Transaction, error) {
	return _KeyperSetManager.Contract.Unpause(&_KeyperSetManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyperSetManager *KeyperSetManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _KeyperSetManager.Contract.Unpause(&_KeyperSetManager.TransactOpts)
}

// KeyperSetManagerKeyperSetAddedIterator is returned from FilterKeyperSetAdded and is used to iterate over the raw logs and unpacked data for KeyperSetAdded events raised by the KeyperSetManager contract.
type KeyperSetManagerKeyperSetAddedIterator struct {
	Event *KeyperSetManagerKeyperSetAdded // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerKeyperSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerKeyperSetAdded)
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
		it.Event = new(KeyperSetManagerKeyperSetAdded)
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
func (it *KeyperSetManagerKeyperSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerKeyperSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerKeyperSetAdded represents a KeyperSetAdded event raised by the KeyperSetManager contract.
type KeyperSetManagerKeyperSetAdded struct {
	ActivationBlock   uint64
	KeyperSetContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterKeyperSetAdded is a free log retrieval operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterKeyperSetAdded(opts *bind.FilterOpts) (*KeyperSetManagerKeyperSetAddedIterator, error) {

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerKeyperSetAddedIterator{contract: _KeyperSetManager.contract, event: "KeyperSetAdded", logs: logs, sub: sub}, nil
}

// WatchKeyperSetAdded is a free log subscription operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchKeyperSetAdded(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerKeyperSetAdded) (event.Subscription, error) {

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerKeyperSetAdded)
				if err := _KeyperSetManager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
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

// ParseKeyperSetAdded is a log parse operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParseKeyperSetAdded(log types.Log) (*KeyperSetManagerKeyperSetAdded, error) {
	event := new(KeyperSetManagerKeyperSetAdded)
	if err := _KeyperSetManager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSetManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KeyperSetManager contract.
type KeyperSetManagerPausedIterator struct {
	Event *KeyperSetManagerPaused // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerPaused)
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
		it.Event = new(KeyperSetManagerPaused)
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
func (it *KeyperSetManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerPaused represents a Paused event raised by the KeyperSetManager contract.
type KeyperSetManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*KeyperSetManagerPausedIterator, error) {

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerPausedIterator{contract: _KeyperSetManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerPaused) (event.Subscription, error) {

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerPaused)
				if err := _KeyperSetManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParsePaused(log types.Log) (*KeyperSetManagerPaused, error) {
	event := new(KeyperSetManagerPaused)
	if err := _KeyperSetManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSetManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the KeyperSetManager contract.
type KeyperSetManagerRoleAdminChangedIterator struct {
	Event *KeyperSetManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerRoleAdminChanged)
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
		it.Event = new(KeyperSetManagerRoleAdminChanged)
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
func (it *KeyperSetManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerRoleAdminChanged represents a RoleAdminChanged event raised by the KeyperSetManager contract.
type KeyperSetManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KeyperSetManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerRoleAdminChangedIterator{contract: _KeyperSetManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerRoleAdminChanged)
				if err := _KeyperSetManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParseRoleAdminChanged(log types.Log) (*KeyperSetManagerRoleAdminChanged, error) {
	event := new(KeyperSetManagerRoleAdminChanged)
	if err := _KeyperSetManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSetManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KeyperSetManager contract.
type KeyperSetManagerRoleGrantedIterator struct {
	Event *KeyperSetManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerRoleGranted)
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
		it.Event = new(KeyperSetManagerRoleGranted)
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
func (it *KeyperSetManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerRoleGranted represents a RoleGranted event raised by the KeyperSetManager contract.
type KeyperSetManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeyperSetManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerRoleGrantedIterator{contract: _KeyperSetManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerRoleGranted)
				if err := _KeyperSetManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParseRoleGranted(log types.Log) (*KeyperSetManagerRoleGranted, error) {
	event := new(KeyperSetManagerRoleGranted)
	if err := _KeyperSetManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSetManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KeyperSetManager contract.
type KeyperSetManagerRoleRevokedIterator struct {
	Event *KeyperSetManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerRoleRevoked)
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
		it.Event = new(KeyperSetManagerRoleRevoked)
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
func (it *KeyperSetManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerRoleRevoked represents a RoleRevoked event raised by the KeyperSetManager contract.
type KeyperSetManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeyperSetManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerRoleRevokedIterator{contract: _KeyperSetManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerRoleRevoked)
				if err := _KeyperSetManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParseRoleRevoked(log types.Log) (*KeyperSetManagerRoleRevoked, error) {
	event := new(KeyperSetManagerRoleRevoked)
	if err := _KeyperSetManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyperSetManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KeyperSetManager contract.
type KeyperSetManagerUnpausedIterator struct {
	Event *KeyperSetManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *KeyperSetManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetManagerUnpaused)
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
		it.Event = new(KeyperSetManagerUnpaused)
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
func (it *KeyperSetManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetManagerUnpaused represents a Unpaused event raised by the KeyperSetManager contract.
type KeyperSetManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KeyperSetManagerUnpausedIterator, error) {

	logs, sub, err := _KeyperSetManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KeyperSetManagerUnpausedIterator{contract: _KeyperSetManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeyperSetManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _KeyperSetManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetManagerUnpaused)
				if err := _KeyperSetManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KeyperSetManager *KeyperSetManagerFilterer) ParseUnpaused(log types.Log) (*KeyperSetManagerUnpaused, error) {
	event := new(KeyperSetManagerUnpaused)
	if err := _KeyperSetManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
