// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sccverifier

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

// LibOVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type LibOVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

// TypesOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// OasysRollupVerifierMetaData contains all meta data concerning the OasysRollupVerifier contract.
var OasysRollupVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidAddressSort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verified\",\"type\":\"uint256\"}],\"name\":\"StakeAmountShortage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611468806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80637d2d9b63146100515780639ac4557214610066578063d17f1fb114610079578063dea2619d1461008c575b600080fd5b61006461005f366004610e3c565b61009f565b005b610064610074366004610e3c565b61017f565b610064610087366004610ece565b61025f565b61006461009a366004610ece565b61034e565b6100ad848484600185610430565b6040517fd882343600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063d8823436906101019086908690600401610fb8565b600060405180830381600087803b15801561011b57600080fd5b505af115801561012f573d6000803e3d6000fd5b50506040518435925085915073ffffffffffffffffffffffffffffffffffffffff8716907f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f90600090a450505050565b61018d848484600085610430565b6040517f859c029d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063859c029d906101e19086908690600401610fb8565b600060405180830381600087803b1580156101fb57600080fd5b505af115801561020f573d6000803e3d6000fd5b50506040518435925085915073ffffffffffffffffffffffffffffffffffffffff8716907ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca390600090a450505050565b61026c838360008461051b565b6040517f9594dd6400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690639594dd64906102be908590600401611085565b600060405180830381600087803b1580156102d857600080fd5b505af11580156102ec573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd846020015160405161034191815260200190565b60405180910390a3505050565b61035b838360018461051b565b6040517fe0c2b41800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063e0c2b418906103ad908590600401611085565b600060405180830381600087803b1580156103c757600080fd5b505af11580156103db573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5846020015160405161034191815260200190565b600046868686356104476040890160208a016110d0565b61045760608a0160408b016110d0565b6040516020016104a393929190928352608091821b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009081166020850152911b16603082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526104e39493929188906020016110f2565b604051602081830303815290604052905060006105078280519060200120846105a7565b90506105128161088d565b50505050505050565b8251602080850151604080514681850152606089901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818301526054810194909452607484019190915284151560f81b60948401528051607581850301815260959093019052815191012060009061059590836105a7565b90506105a08161088d565b5050505050565b6060815167ffffffffffffffff8111156105c3576105c3610c68565b6040519080825280602002602001820160405280156105ec578160200160208202803683370190505b506040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101859052909150600090605c016040516020818303038152906040528051906020012090506000805b845181101561088457600085828151811061065f5761065f611159565b602002602001015190506000806106768684610ae2565b9092509050600181600481111561068f5761068f611188565b036106d157826040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016106c891906111b7565b60405180910390fd5b60028160048111156106e5576106e5611188565b0361071e57826040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016106c89190611209565b600381600481111561073257610732611188565b0361076b57826040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016106c8919061125b565b600481600481111561077f5761077f611188565b036107b857826040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016106c891906112d3565b8473ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1611610835576040517f3d1fd1b000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016106c8565b8187858151811061084857610848611159565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250925081905061087c8161137a565b915050610642565b50505092915050565b600061100090506000611001905060008273ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090c91906113b2565b84519091506000805b828110156109f1578473ffffffffffffffffffffffffffffffffffffffff16639c50821988838151811061094b5761094b611159565b6020026020010151866040518363ffffffff1660e01b815260040161099292919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b602060405180830381865afa1580156109af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d391906113b2565b6109dd90836113cb565b9150806109e98161137a565b915050610915565b506040517f45367f230000000000000000000000000000000000000000000000000000000081526004810184905260009060649073ffffffffffffffffffffffffffffffffffffffff8716906345367f2390602401602060405180830381865afa158015610a63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8791906113b2565b610a929060336113e3565b610a9c9190611420565b905080821015610512576040517f09b3dd4a00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016106c8565b6000808251604103610b185760208301516040840151606085015160001a610b0c87828585610b27565b94509450505050610b20565b506000905060025b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b5e5750600090506003610c36565b8460ff16601b14158015610b7657508460ff16601c14155b15610b875750600090506004610c36565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610bdb573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610c2f57600060019250925050610c36565b9150600090505b94509492505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c6357600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610cba57610cba610c68565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d0757610d07610c68565b604052919050565b600082601f830112610d2057600080fd5b813567ffffffffffffffff811115610d3a57610d3a610c68565b610d6b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610cc0565b818152846020838601011115610d8057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610dae57600080fd5b8135602067ffffffffffffffff80831115610dcb57610dcb610c68565b8260051b610dda838201610cc0565b9384528581018301938381019088861115610df457600080fd5b84880192505b85831015610e3057823584811115610e125760008081fd5b610e208a87838c0101610d0f565b8352509184019190840190610dfa565b98975050505050505050565b60008060008084860360c0811215610e5357600080fd5b610e5c86610c3f565b94506020860135935060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082011215610e9557600080fd5b5060408501915060a085013567ffffffffffffffff811115610eb657600080fd5b610ec287828801610d9d565b91505092959194509250565b600080600060608486031215610ee357600080fd5b610eec84610c3f565b9250602084013567ffffffffffffffff80821115610f0957600080fd5b9085019060a08288031215610f1d57600080fd5b610f25610c97565b82358152602083013560208201526040830135604082015260608301356060820152608083013582811115610f5957600080fd5b610f6589828601610d0f565b60808301525093506040860135915080821115610f8157600080fd5b50610f8e86828701610d9d565b9150509250925092565b80356fffffffffffffffffffffffffffffffff81168114610c6357600080fd5b600060808201905083825282356020830152610fd660208401610f98565b6fffffffffffffffffffffffffffffffff808216604085015280610ffc60408701610f98565b16606085015250509392505050565b60005b8381101561102657818101518382015260200161100e565b83811115611035576000848401525b50505050565b6000815180845261105381602086016020860161100b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152815160208201526020820151604082015260408201516060820152606082015160808201526000608083015160a0808401526110c860c084018261103b565b949350505050565b6000602082840312156110e257600080fd5b6110eb82610f98565b9392505050565b8581527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b1660208201528360348201526000835161113b81605485016020880161100b565b92151560f81b91909201605481019190915260550195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6040815260006111ca604083018461103b565b8281036020840152601881527f45434453413a20696e76616c6964207369676e6174757265000000000000000060208201526040810191505092915050565b60408152600061121c604083018461103b565b8281036020840152601f81527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060208201526040810191505092915050565b60408152600061126e604083018461103b565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60208201527f756500000000000000000000000000000000000000000000000000000000000060408201526060810191505092915050565b6040815260006112e6604083018461103b565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60208201527f756500000000000000000000000000000000000000000000000000000000000060408201526060810191505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113ab576113ab61134b565b5060010190565b6000602082840312156113c457600080fd5b5051919050565b600082198211156113de576113de61134b565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561141b5761141b61134b565b500290565b600082611456577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a",
}

// OasysRollupVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use OasysRollupVerifierMetaData.ABI instead.
var OasysRollupVerifierABI = OasysRollupVerifierMetaData.ABI

// OasysRollupVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OasysRollupVerifierMetaData.Bin instead.
var OasysRollupVerifierBin = OasysRollupVerifierMetaData.Bin

// DeployOasysRollupVerifier deploys a new Ethereum contract, binding an instance of OasysRollupVerifier to it.
func DeployOasysRollupVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OasysRollupVerifier, error) {
	parsed, err := OasysRollupVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OasysRollupVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OasysRollupVerifier{OasysRollupVerifierCaller: OasysRollupVerifierCaller{contract: contract}, OasysRollupVerifierTransactor: OasysRollupVerifierTransactor{contract: contract}, OasysRollupVerifierFilterer: OasysRollupVerifierFilterer{contract: contract}}, nil
}

// OasysRollupVerifier is an auto generated Go binding around an Ethereum contract.
type OasysRollupVerifier struct {
	OasysRollupVerifierCaller     // Read-only binding to the contract
	OasysRollupVerifierTransactor // Write-only binding to the contract
	OasysRollupVerifierFilterer   // Log filterer for contract events
}

// OasysRollupVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasysRollupVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysRollupVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasysRollupVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysRollupVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasysRollupVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysRollupVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasysRollupVerifierSession struct {
	Contract     *OasysRollupVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OasysRollupVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasysRollupVerifierCallerSession struct {
	Contract *OasysRollupVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OasysRollupVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasysRollupVerifierTransactorSession struct {
	Contract     *OasysRollupVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OasysRollupVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasysRollupVerifierRaw struct {
	Contract *OasysRollupVerifier // Generic contract binding to access the raw methods on
}

// OasysRollupVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasysRollupVerifierCallerRaw struct {
	Contract *OasysRollupVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// OasysRollupVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasysRollupVerifierTransactorRaw struct {
	Contract *OasysRollupVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasysRollupVerifier creates a new instance of OasysRollupVerifier, bound to a specific deployed contract.
func NewOasysRollupVerifier(address common.Address, backend bind.ContractBackend) (*OasysRollupVerifier, error) {
	contract, err := bindOasysRollupVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifier{OasysRollupVerifierCaller: OasysRollupVerifierCaller{contract: contract}, OasysRollupVerifierTransactor: OasysRollupVerifierTransactor{contract: contract}, OasysRollupVerifierFilterer: OasysRollupVerifierFilterer{contract: contract}}, nil
}

// NewOasysRollupVerifierCaller creates a new read-only instance of OasysRollupVerifier, bound to a specific deployed contract.
func NewOasysRollupVerifierCaller(address common.Address, caller bind.ContractCaller) (*OasysRollupVerifierCaller, error) {
	contract, err := bindOasysRollupVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierCaller{contract: contract}, nil
}

// NewOasysRollupVerifierTransactor creates a new write-only instance of OasysRollupVerifier, bound to a specific deployed contract.
func NewOasysRollupVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*OasysRollupVerifierTransactor, error) {
	contract, err := bindOasysRollupVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierTransactor{contract: contract}, nil
}

// NewOasysRollupVerifierFilterer creates a new log filterer instance of OasysRollupVerifier, bound to a specific deployed contract.
func NewOasysRollupVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*OasysRollupVerifierFilterer, error) {
	contract, err := bindOasysRollupVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierFilterer{contract: contract}, nil
}

// bindOasysRollupVerifier binds a generic wrapper to an already deployed contract.
func bindOasysRollupVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OasysRollupVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysRollupVerifier *OasysRollupVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysRollupVerifier.Contract.OasysRollupVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysRollupVerifier *OasysRollupVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.OasysRollupVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysRollupVerifier *OasysRollupVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.OasysRollupVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysRollupVerifier *OasysRollupVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysRollupVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysRollupVerifier *OasysRollupVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysRollupVerifier *OasysRollupVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.contract.Transact(opts, method, params...)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactor) Approve(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.contract.Transact(opts, "approve", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Approve(&_OasysRollupVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactorSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Approve(&_OasysRollupVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactor) Approve0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.contract.Transact(opts, "approve0", stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierSession) Approve0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Approve0(&_OasysRollupVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactorSession) Approve0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Approve0(&_OasysRollupVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactor) Reject(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.contract.Transact(opts, "reject", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Reject(&_OasysRollupVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactorSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Reject(&_OasysRollupVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactor) Reject0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.contract.Transact(opts, "reject0", stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierSession) Reject0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Reject0(&_OasysRollupVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysRollupVerifier *OasysRollupVerifierTransactorSession) Reject0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysRollupVerifier.Contract.Reject0(&_OasysRollupVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// OasysRollupVerifierL2OutputApprovedIterator is returned from FilterL2OutputApproved and is used to iterate over the raw logs and unpacked data for L2OutputApproved events raised by the OasysRollupVerifier contract.
type OasysRollupVerifierL2OutputApprovedIterator struct {
	Event *OasysRollupVerifierL2OutputApproved // Event containing the contract specifics and raw log

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
func (it *OasysRollupVerifierL2OutputApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysRollupVerifierL2OutputApproved)
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
		it.Event = new(OasysRollupVerifierL2OutputApproved)
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
func (it *OasysRollupVerifierL2OutputApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysRollupVerifierL2OutputApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysRollupVerifierL2OutputApproved represents a L2OutputApproved event raised by the OasysRollupVerifier contract.
type OasysRollupVerifierL2OutputApproved struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputApproved is a free log retrieval operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) FilterL2OutputApproved(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*OasysRollupVerifierL2OutputApprovedIterator, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.FilterLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierL2OutputApprovedIterator{contract: _OasysRollupVerifier.contract, event: "L2OutputApproved", logs: logs, sub: sub}, nil
}

// WatchL2OutputApproved is a free log subscription operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) WatchL2OutputApproved(opts *bind.WatchOpts, sink chan<- *OasysRollupVerifierL2OutputApproved, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.WatchLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysRollupVerifierL2OutputApproved)
				if err := _OasysRollupVerifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
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

// ParseL2OutputApproved is a log parse operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) ParseL2OutputApproved(log types.Log) (*OasysRollupVerifierL2OutputApproved, error) {
	event := new(OasysRollupVerifierL2OutputApproved)
	if err := _OasysRollupVerifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysRollupVerifierL2OutputRejectedIterator is returned from FilterL2OutputRejected and is used to iterate over the raw logs and unpacked data for L2OutputRejected events raised by the OasysRollupVerifier contract.
type OasysRollupVerifierL2OutputRejectedIterator struct {
	Event *OasysRollupVerifierL2OutputRejected // Event containing the contract specifics and raw log

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
func (it *OasysRollupVerifierL2OutputRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysRollupVerifierL2OutputRejected)
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
		it.Event = new(OasysRollupVerifierL2OutputRejected)
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
func (it *OasysRollupVerifierL2OutputRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysRollupVerifierL2OutputRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysRollupVerifierL2OutputRejected represents a L2OutputRejected event raised by the OasysRollupVerifier contract.
type OasysRollupVerifierL2OutputRejected struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputRejected is a free log retrieval operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) FilterL2OutputRejected(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*OasysRollupVerifierL2OutputRejectedIterator, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.FilterLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierL2OutputRejectedIterator{contract: _OasysRollupVerifier.contract, event: "L2OutputRejected", logs: logs, sub: sub}, nil
}

// WatchL2OutputRejected is a free log subscription operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) WatchL2OutputRejected(opts *bind.WatchOpts, sink chan<- *OasysRollupVerifierL2OutputRejected, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.WatchLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysRollupVerifierL2OutputRejected)
				if err := _OasysRollupVerifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
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

// ParseL2OutputRejected is a log parse operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) ParseL2OutputRejected(log types.Log) (*OasysRollupVerifierL2OutputRejected, error) {
	event := new(OasysRollupVerifierL2OutputRejected)
	if err := _OasysRollupVerifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysRollupVerifierStateBatchApprovedIterator is returned from FilterStateBatchApproved and is used to iterate over the raw logs and unpacked data for StateBatchApproved events raised by the OasysRollupVerifier contract.
type OasysRollupVerifierStateBatchApprovedIterator struct {
	Event *OasysRollupVerifierStateBatchApproved // Event containing the contract specifics and raw log

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
func (it *OasysRollupVerifierStateBatchApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysRollupVerifierStateBatchApproved)
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
		it.Event = new(OasysRollupVerifierStateBatchApproved)
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
func (it *OasysRollupVerifierStateBatchApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysRollupVerifierStateBatchApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysRollupVerifierStateBatchApproved represents a StateBatchApproved event raised by the OasysRollupVerifier contract.
type OasysRollupVerifierStateBatchApproved struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchApproved is a free log retrieval operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) FilterStateBatchApproved(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*OasysRollupVerifierStateBatchApprovedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.FilterLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierStateBatchApprovedIterator{contract: _OasysRollupVerifier.contract, event: "StateBatchApproved", logs: logs, sub: sub}, nil
}

// WatchStateBatchApproved is a free log subscription operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) WatchStateBatchApproved(opts *bind.WatchOpts, sink chan<- *OasysRollupVerifierStateBatchApproved, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.WatchLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysRollupVerifierStateBatchApproved)
				if err := _OasysRollupVerifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
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

// ParseStateBatchApproved is a log parse operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) ParseStateBatchApproved(log types.Log) (*OasysRollupVerifierStateBatchApproved, error) {
	event := new(OasysRollupVerifierStateBatchApproved)
	if err := _OasysRollupVerifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysRollupVerifierStateBatchRejectedIterator is returned from FilterStateBatchRejected and is used to iterate over the raw logs and unpacked data for StateBatchRejected events raised by the OasysRollupVerifier contract.
type OasysRollupVerifierStateBatchRejectedIterator struct {
	Event *OasysRollupVerifierStateBatchRejected // Event containing the contract specifics and raw log

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
func (it *OasysRollupVerifierStateBatchRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysRollupVerifierStateBatchRejected)
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
		it.Event = new(OasysRollupVerifierStateBatchRejected)
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
func (it *OasysRollupVerifierStateBatchRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysRollupVerifierStateBatchRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysRollupVerifierStateBatchRejected represents a StateBatchRejected event raised by the OasysRollupVerifier contract.
type OasysRollupVerifierStateBatchRejected struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchRejected is a free log retrieval operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) FilterStateBatchRejected(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*OasysRollupVerifierStateBatchRejectedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.FilterLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysRollupVerifierStateBatchRejectedIterator{contract: _OasysRollupVerifier.contract, event: "StateBatchRejected", logs: logs, sub: sub}, nil
}

// WatchStateBatchRejected is a free log subscription operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) WatchStateBatchRejected(opts *bind.WatchOpts, sink chan<- *OasysRollupVerifierStateBatchRejected, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysRollupVerifier.contract.WatchLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysRollupVerifierStateBatchRejected)
				if err := _OasysRollupVerifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
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

// ParseStateBatchRejected is a log parse operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysRollupVerifier *OasysRollupVerifierFilterer) ParseStateBatchRejected(log types.Log) (*OasysRollupVerifierStateBatchRejected, error) {
	event := new(OasysRollupVerifierStateBatchRejected)
	if err := _OasysRollupVerifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
