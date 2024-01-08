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
	Bin: "0x608060405234801561001057600080fd5b5061161c806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80637d2d9b63146100515780639ac4557214610066578063d17f1fb114610079578063dea2619d1461008c575b600080fd5b61006461005f366004610dd3565b61009f565b005b610064610074366004610dd3565b610191565b61006461008736600461107c565b610274565b61006461009a36600461107c565b61040f565b6100be6100af86868660016105ba565b6100b98385611146565b6106e4565b6040517fd882343600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86169063d8823436906101129087908790600401611173565b600060405180830381600087803b15801561012c57600080fd5b505af1158015610140573d6000803e3d6000fd5b50506040518535925086915073ffffffffffffffffffffffffffffffffffffffff8816907f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f90600090a45050505050565b6101a16100af86868660006105ba565b6040517f859c029d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86169063859c029d906101f59087908790600401611173565b600060405180830381600087803b15801561020f57600080fd5b505af1158015610223573d6000803e3d6000fd5b50506040518535925086915073ffffffffffffffffffffffffffffffffffffffff8816907ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca390600090a45050505050565b8151602080840151604080514681850152606088901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818301526054810194909452607484019190915260006094840152805160758185030181526095840182528051908301207f19457468657265756d205369676e6564204d6573736167653a0a33320000000060b585015260d1808501919091528151808503909101815260f19093019052815191012061032d905b826106e4565b6040517f9594dd6400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690639594dd649061037f908590600401611240565b600060405180830381600087803b15801561039957600080fd5b505af11580156103ad573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd846020015160405161040291815260200190565b60405180910390a3505050565b8151602080840151604080514681850152606088901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001681830152605481019490945260748401919091527f01000000000000000000000000000000000000000000000000000000000000006094840152805160758185030181526095840182528051908301207f19457468657265756d205369676e6564204d6573736167653a0a33320000000060b585015260d1808501919091528151808503909101815260f1909301905281519101206104e590610327565b6040517fe0c2b41800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063e0c2b41890610537908590600401611240565b600060405180830381600087803b15801561055157600080fd5b505af1158015610565573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5846020015160405161040291815260200190565b60008046868686356105d26040890160208a0161128b565b6105e260608a0160408b0161128b565b60405160200161062e93929190928352608091821b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009081166020850152911b16603082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261066e9493929188906020016112a6565b604051602081830303815290604052905080805190602001206040516020016106c391907f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b60405160208183030381529060405280519060200120915050949350505050565b6106f66106f183836106fa565b61084c565b5050565b6060815167ffffffffffffffff81111561071657610716610e9b565b60405190808252806020026020018201604052801561073f578160200160208202803683370190505b5090506000805b8351811015610844576000610774868684815181106107675761076761130d565b6020026020010151610aaa565b90508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16116107f8576040517f3d1fd1b000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b8084838151811061080b5761080b61130d565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015291508061083c8161136b565b915050610746565b505092915050565b600061100090506000611001905060008273ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cb91906113a3565b84519091506000805b828110156109b0578473ffffffffffffffffffffffffffffffffffffffff16639c50821988838151811061090a5761090a61130d565b6020026020010151866040518363ffffffff1660e01b815260040161095192919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b602060405180830381865afa15801561096e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099291906113a3565b61099c90836113bc565b9150806109a88161136b565b9150506108d4565b506040517f45367f230000000000000000000000000000000000000000000000000000000081526004810184905260009060649073ffffffffffffffffffffffffffffffffffffffff8716906345367f2390602401602060405180830381865afa158015610a22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4691906113a3565b610a519060336113d4565b610a5b9190611411565b905080821015610aa1576040517f09b3dd4a00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016107ef565b50505050505050565b6000806000610ab98585610c4d565b90925090506000816004811115610ad257610ad261144c565b03610adf57509050610c47565b6001816004811115610af357610af361144c565b03610b2c57836040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016107ef919061147b565b6002816004811115610b4057610b4061144c565b03610b7957836040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016107ef91906114cd565b6003816004811115610b8d57610b8d61144c565b03610bc657836040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016107ef919061151f565b6004816004811115610bda57610bda61144c565b03610c1357836040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016107ef9190611597565b836040517ff31b6ee50000000000000000000000000000000000000000000000000000000081526004016107ef919061147b565b92915050565b6000808251604103610c835760208301516040840151606085015160001a610c7787828585610c92565b94509450505050610c8b565b506000905060025b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610cc95750600090506003610da1565b8460ff16601b14158015610ce157508460ff16601c14155b15610cf25750600090506004610da1565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d46573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610d9a57600060019250925050610da1565b9150600090505b94509492505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610dce57600080fd5b919050565b600080600080600085870360c0811215610dec57600080fd5b610df587610daa565b95506020870135945060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082011215610e2e57600080fd5b5060408601925060a086013567ffffffffffffffff80821115610e5057600080fd5b818801915088601f830112610e6457600080fd5b813581811115610e7357600080fd5b8960208260051b8501011115610e8857600080fd5b9699959850939650602001949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610eed57610eed610e9b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610f3a57610f3a610e9b565b604052919050565b600082601f830112610f5357600080fd5b813567ffffffffffffffff811115610f6d57610f6d610e9b565b610f9e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610ef3565b818152846020838601011115610fb357600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff80841115610feb57610feb610e9b565b8360051b6020610ffc818301610ef3565b8681529350908401908084018783111561101557600080fd5b855b838110156110495780358581111561102f5760008081fd5b61103b8a828a01610f42565b835250908201908201611017565b50505050509392505050565b600082601f83011261106657600080fd5b61107583833560208501610fd0565b9392505050565b60008060006060848603121561109157600080fd5b61109a84610daa565b9250602084013567ffffffffffffffff808211156110b757600080fd5b9085019060a082880312156110cb57600080fd5b6110d3610eca565b8235815260208301356020820152604083013560408201526060830135606082015260808301358281111561110757600080fd5b61111389828601610f42565b6080830152509350604086013591508082111561112f57600080fd5b5061113c86828701611055565b9150509250925092565b6000611075368484610fd0565b80356fffffffffffffffffffffffffffffffff81168114610dce57600080fd5b60006080820190508382528235602083015261119160208401611153565b6fffffffffffffffffffffffffffffffff8082166040850152806111b760408701611153565b16606085015250509392505050565b60005b838110156111e15781810151838201526020016111c9565b838111156111f0576000848401525b50505050565b6000815180845261120e8160208601602086016111c6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152815160208201526020820151604082015260408201516060820152606082015160808201526000608083015160a08084015261128360c08401826111f6565b949350505050565b60006020828403121561129d57600080fd5b61107582611153565b8581527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b166020820152836034820152600083516112ef8160548501602088016111c6565b92151560f81b91909201605481019190915260550195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361139c5761139c61133c565b5060010190565b6000602082840312156113b557600080fd5b5051919050565b600082198211156113cf576113cf61133c565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561140c5761140c61133c565b500290565b600082611447577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60408152600061148e60408301846111f6565b8281036020840152601881527f45434453413a20696e76616c6964207369676e6174757265000000000000000060208201526040810191505092915050565b6040815260006114e060408301846111f6565b8281036020840152601f81527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060208201526040810191505092915050565b60408152600061153260408301846111f6565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60208201527f756500000000000000000000000000000000000000000000000000000000000060408201526060810191505092915050565b6040815260006115aa60408301846111f6565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60208201527f75650000000000000000000000000000000000000000000000000000000000006040820152606081019150509291505056fea164736f6c634300080f000a",
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
