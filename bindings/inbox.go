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

// InboxTransaction is an auto generated low-level Go binding around an user-defined struct.
type InboxTransaction struct {
	EncryptedTransaction []byte
	GasLimit             uint64
	CumulativeGasLimit   uint64
}

// InboxMetaData contains all meta data concerning the Inbox contract.
var InboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_blockGasLimit\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockGasLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferEtherFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedTransaction\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"EncryptedTransactionSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_GAS_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"getTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"encryptedTransaction\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structInbox.Transaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newBlockGasLimit\",\"type\":\"uint64\"}],\"name\":\"setBlockGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"encryptedTransaction\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"excessFeeRecipient\",\"type\":\"address\"}],\"name\":\"submitEncryptedTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161132638038061132683398101604081905261002f91610116565b6001805460ff1916905561004460003361006a565b50600380546001600160401b0319166001600160401b0392909216919091179055610146565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1661010c576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556100c43390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610110565b5060005b92915050565b60006020828403121561012857600080fd5b81516001600160401b038116811461013f57600080fd5b9392505050565b6111d1806101556000396000f3fe60806040526004361061011f5760003560e01c80635c975abb116100a0578063a44e7cfe11610064578063a44e7cfe14610342578063ae0dad6014610355578063d547741f14610375578063e02023a114610395578063e63ab1e9146103c957600080fd5b80635c975abb146102ac5780638456cb59146102c457806391d14854146102d957806391de737e146102f9578063a217fddf1461032d57600080fd5b80633f4ba83a116100e75780633f4ba83a146102015780634842855c1461021657806351cff8d91461024a57806352efea6e1461026a578063552fd4aa1461027f57600080fd5b806301ffc9a714610124578063248a9ca3146101595780632cc8377d146101975780632f2ff15d146101bf57806336568abe146101e1575b600080fd5b34801561013057600080fd5b5061014461013f366004610cd1565b6103fd565b60405190151581526020015b60405180910390f35b34801561016557600080fd5b50610189610174366004610d02565b60009081526020819052604090206001015490565b604051908152602001610150565b3480156101a357600080fd5b506003546040516001600160401b039091168152602001610150565b3480156101cb57600080fd5b506101df6101da366004610d37565b610434565b005b3480156101ed57600080fd5b506101df6101fc366004610d37565b61045f565b34801561020d57600080fd5b506101df610497565b34801561022257600080fd5b506101897fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c881565b34801561025657600080fd5b506101df610265366004610d63565b6104ad565b34801561027657600080fd5b506101df610521565b34801561028b57600080fd5b5061029f61029a366004610d95565b61056c565b6040516101509190610df6565b3480156102b857600080fd5b5060015460ff16610144565b3480156102d057600080fd5b506101df61069d565b3480156102e557600080fd5b506101446102f4366004610d37565b6106cf565b34801561030557600080fd5b506101897ff265e6f4485879ff06917e0b78d882e64c54d090917c00d06bf8720a31ba125681565b34801561033957600080fd5b50610189600081565b6101df610350366004610e9d565b6106f8565b34801561036157600080fd5b506101df610370366004610d95565b610932565b34801561038157600080fd5b506101df610390366004610d37565b610980565b3480156103a157600080fd5b506101897f5d8e12c39142ff96d79d04d15d1ba1269e4fe57bb9d26f43523628b34ba108ec81565b3480156103d557600080fd5b506101897f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60006001600160e01b03198216637965db0b60e01b148061042e57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60008281526020819052604090206001015461044f816109a5565b61045983836109af565b50505050565b6001600160a01b03811633146104885760405163334bd91960e11b815260040160405180910390fd5b6104928282610a41565b505050565b60006104a2816109a5565b6104aa610aac565b50565b7f5d8e12c39142ff96d79d04d15d1ba1269e4fe57bb9d26f43523628b34ba108ec6104d7816109a5565b6104e18247610afe565b6040516001600160a01b03831681527f792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c79060200160405180910390a15050565b7fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c861054b816109a5565b6001600160401b03431660009081526002602052604081206104aa91610c37565b6001600160401b0381166000908152600260209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561069257838290600052602060002090600202016040518060600160405290816000820180546105d790610f7f565b80601f016020809104026020016040519081016040528092919081815260200182805461060390610f7f565b80156106505780601f1061062557610100808354040283529160200191610650565b820191906000526020600020905b81548152906001019060200180831161063357829003601f168201915b50505091835250506001918201546001600160401b03808216602080850191909152600160401b909204166040909201919091529183529290920191016105a4565b505050509050919050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6106c7816109a5565b6104aa610b72565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b610700610bad565b436001600160401b0316846001600160401b031611610732576040516351c9b4cf60e11b815260040160405180910390fd5b6000610747486001600160401b038516610fcf565b90508034101561076a5760405163356680b760e01b815260040160405180910390fd5b6001600160401b0380861660009081526002602052604081208054909290919082166107975760006107e2565b826107a3600184610fe6565b6001600160401b0316815481106107bc576107bc61100d565b906000526020600020906002020160010160089054906101000a90046001600160401b03165b6003549091506001600160401b03166107fb8783611023565b6001600160401b0316111561082357604051635013664360e11b815260040160405180910390fd5b604080516060810182528881526001600160401b03881660208201528491810161084d8985611023565b6001600160401b03169052815460018101835560009283526020909220815191926002020190819061087f9082611093565b506020820151600190910180546040909301516001600160401b03908116600160401b026001600160801b031990941692811692909217929092179091558881169083167fa7225f029f7fe3ce824a1ea993c91894e75616de5ac949a7b8c85a4d9531200489896108f08187611023565b896040516109019493929190611152565b60405180910390a360006109158534611188565b90508015610927576109278682610afe565b505050505050505050565b7ff265e6f4485879ff06917e0b78d882e64c54d090917c00d06bf8720a31ba125661095c816109a5565b506003805467ffffffffffffffff19166001600160401b0392909216919091179055565b60008281526020819052604090206001015461099b816109a5565b6104598383610a41565b6104aa8133610bd3565b60006109bb83836106cf565b610a39576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556109f13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161042e565b50600061042e565b6000610a4d83836106cf565b15610a39576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161042e565b610ab4610c14565b6001805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610b4b576040519150601f19603f3d011682016040523d82523d6000602084013e610b50565b606091505b505090508061049257604051631583ef2560e21b815260040160405180910390fd5b610b7a610bad565b6001805460ff1916811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610ae1565b60015460ff1615610bd15760405163d93c066560e01b815260040160405180910390fd5b565b610bdd82826106cf565b610c105760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610bd157604051638dfc202b60e01b815260040160405180910390fd5b50805460008255600202906000526020600020908101906104aa91905b80821115610c83576000610c688282610c87565b506001810180546001600160801b0319169055600201610c54565b5090565b508054610c9390610f7f565b6000825580601f10610ca3575050565b601f0160209004906000526020600020908101906104aa91905b80821115610c835760008155600101610cbd565b600060208284031215610ce357600080fd5b81356001600160e01b031981168114610cfb57600080fd5b9392505050565b600060208284031215610d1457600080fd5b5035919050565b80356001600160a01b0381168114610d3257600080fd5b919050565b60008060408385031215610d4a57600080fd5b82359150610d5a60208401610d1b565b90509250929050565b600060208284031215610d7557600080fd5b610cfb82610d1b565b80356001600160401b0381168114610d3257600080fd5b600060208284031215610da757600080fd5b610cfb82610d7e565b6000815180845260005b81811015610dd657602081850181015186830182015201610dba565b506000602082860101526020601f19601f83011685010191505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015610e7957603f19898403018552815160608151818652610e4582870182610db0565b838b01516001600160401b03908116888d0152938a0151909316958901959095525094870194925090860190600101610e1f565b509098975050505050505050565b634e487b7160e01b600052604160045260246000fd5b60008060008060808587031215610eb357600080fd5b610ebc85610d7e565b935060208501356001600160401b0380821115610ed857600080fd5b818701915087601f830112610eec57600080fd5b813581811115610efe57610efe610e87565b604051601f8201601f19908116603f01168101908382118183101715610f2657610f26610e87565b816040528281528a6020848701011115610f3f57600080fd5b826020860160208301376000602084830101528097505050505050610f6660408601610d7e565b9150610f7460608601610d1b565b905092959194509250565b600181811c90821680610f9357607f821691505b602082108103610fb357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761042e5761042e610fb9565b6001600160401b0382811682821603908082111561100657611006610fb9565b5092915050565b634e487b7160e01b600052603260045260246000fd5b6001600160401b0381811683821601908082111561100657611006610fb9565b601f821115610492576000816000526020600020601f850160051c8101602086101561106c5750805b601f850160051c820191505b8181101561108b57828155600101611078565b505050505050565b81516001600160401b038111156110ac576110ac610e87565b6110c0816110ba8454610f7f565b84611043565b602080601f8311600181146110f557600084156110dd5750858301515b600019600386901b1c1916600185901b17855561108b565b600085815260208120601f198616915b8281101561112457888601518255948401946001909101908401611105565b50858210156111425787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080815260006111656080830187610db0565b6001600160401b0395861660208401529390941660408201526060015292915050565b8181038181111561042e5761042e610fb956fea26469706673582212205d3f5a9d3f26881624dc27a9f346327edc36a9fa84ac7304394205a273004c6564736f6c63430008160033",
}

// InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxMetaData.ABI instead.
var InboxABI = InboxMetaData.ABI

// InboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxMetaData.Bin instead.
var InboxBin = InboxMetaData.Bin

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _blockGasLimit uint64) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxBin), backend, _blockGasLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) BLOCKGASLIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "BLOCK_GAS_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) SEQUENCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "SEQUENCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCaller) GetBlockGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getBlockGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCallerSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,uint64,uint64)[])
func (_Inbox *InboxCaller) GetTransactions(opts *bind.CallOpts, blockNumber uint64) ([]InboxTransaction, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getTransactions", blockNumber)

	if err != nil {
		return *new([]InboxTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]InboxTransaction)).(*[]InboxTransaction)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,uint64,uint64)[])
func (_Inbox *InboxSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,uint64,uint64)[])
func (_Inbox *InboxCallerSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCallerSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactorSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactor) SetBlockGasLimit(opts *bind.TransactOpts, newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setBlockGasLimit", newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactorSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "submitEncryptedTransaction", blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactorSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// InboxEncryptedTransactionSubmittedIterator is returned from FilterEncryptedTransactionSubmitted and is used to iterate over the raw logs and unpacked data for EncryptedTransactionSubmitted events raised by the Inbox contract.
type InboxEncryptedTransactionSubmittedIterator struct {
	Event *InboxEncryptedTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *InboxEncryptedTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxEncryptedTransactionSubmitted)
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
		it.Event = new(InboxEncryptedTransactionSubmitted)
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
func (it *InboxEncryptedTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxEncryptedTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxEncryptedTransactionSubmitted represents a EncryptedTransactionSubmitted event raised by the Inbox contract.
type InboxEncryptedTransactionSubmitted struct {
	Index                uint64
	Block                uint64
	EncryptedTransaction []byte
	GasLimit             uint64
	CumulativeGasLimit   uint64
	Fee                  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEncryptedTransactionSubmitted is a free log retrieval operation binding the contract event 0xa7225f029f7fe3ce824a1ea993c91894e75616de5ac949a7b8c85a4d95312004.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) FilterEncryptedTransactionSubmitted(opts *bind.FilterOpts, index []uint64, block []uint64) (*InboxEncryptedTransactionSubmittedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return &InboxEncryptedTransactionSubmittedIterator{contract: _Inbox.contract, event: "EncryptedTransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchEncryptedTransactionSubmitted is a free log subscription operation binding the contract event 0xa7225f029f7fe3ce824a1ea993c91894e75616de5ac949a7b8c85a4d95312004.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) WatchEncryptedTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *InboxEncryptedTransactionSubmitted, index []uint64, block []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxEncryptedTransactionSubmitted)
				if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
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

// ParseEncryptedTransactionSubmitted is a log parse operation binding the contract event 0xa7225f029f7fe3ce824a1ea993c91894e75616de5ac949a7b8c85a4d95312004.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) ParseEncryptedTransactionSubmitted(log types.Log) (*InboxEncryptedTransactionSubmitted, error) {
	event := new(InboxEncryptedTransactionSubmitted)
	if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the Inbox contract.
type InboxFeesWithdrawnIterator struct {
	Event *InboxFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *InboxFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxFeesWithdrawn)
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
		it.Event = new(InboxFeesWithdrawn)
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
func (it *InboxFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxFeesWithdrawn represents a FeesWithdrawn event raised by the Inbox contract.
type InboxFeesWithdrawn struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*InboxFeesWithdrawnIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &InboxFeesWithdrawnIterator{contract: _Inbox.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *InboxFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxFeesWithdrawn)
				if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

// ParseFeesWithdrawn is a log parse operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) ParseFeesWithdrawn(log types.Log) (*InboxFeesWithdrawn, error) {
	event := new(InboxFeesWithdrawn)
	if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Inbox contract.
type InboxPausedIterator struct {
	Event *InboxPaused // Event containing the contract specifics and raw log

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
func (it *InboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPaused)
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
		it.Event = new(InboxPaused)
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
func (it *InboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPaused represents a Paused event raised by the Inbox contract.
type InboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) FilterPaused(opts *bind.FilterOpts) (*InboxPausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InboxPausedIterator{contract: _Inbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InboxPaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPaused)
				if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParsePaused(log types.Log) (*InboxPaused, error) {
	event := new(InboxPaused)
	if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Inbox contract.
type InboxRoleAdminChangedIterator struct {
	Event *InboxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *InboxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleAdminChanged)
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
		it.Event = new(InboxRoleAdminChanged)
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
func (it *InboxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleAdminChanged represents a RoleAdminChanged event raised by the Inbox contract.
type InboxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*InboxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleAdminChangedIterator{contract: _Inbox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *InboxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleAdminChanged)
				if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleAdminChanged(log types.Log) (*InboxRoleAdminChanged, error) {
	event := new(InboxRoleAdminChanged)
	if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Inbox contract.
type InboxRoleGrantedIterator struct {
	Event *InboxRoleGranted // Event containing the contract specifics and raw log

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
func (it *InboxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleGranted)
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
		it.Event = new(InboxRoleGranted)
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
func (it *InboxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleGranted represents a RoleGranted event raised by the Inbox contract.
type InboxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleGrantedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleGrantedIterator{contract: _Inbox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *InboxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleGranted)
				if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleGranted(log types.Log) (*InboxRoleGranted, error) {
	event := new(InboxRoleGranted)
	if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Inbox contract.
type InboxRoleRevokedIterator struct {
	Event *InboxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *InboxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleRevoked)
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
		it.Event = new(InboxRoleRevoked)
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
func (it *InboxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleRevoked represents a RoleRevoked event raised by the Inbox contract.
type InboxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleRevokedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleRevokedIterator{contract: _Inbox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *InboxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleRevoked)
				if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleRevoked(log types.Log) (*InboxRoleRevoked, error) {
	event := new(InboxRoleRevoked)
	if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Inbox contract.
type InboxUnpausedIterator struct {
	Event *InboxUnpaused // Event containing the contract specifics and raw log

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
func (it *InboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxUnpaused)
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
		it.Event = new(InboxUnpaused)
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
func (it *InboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxUnpaused represents a Unpaused event raised by the Inbox contract.
type InboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InboxUnpausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InboxUnpausedIterator{contract: _Inbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxUnpaused)
				if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseUnpaused(log types.Log) (*InboxUnpaused, error) {
	event := new(InboxUnpaused)
	if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
