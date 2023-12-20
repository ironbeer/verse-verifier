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

// OasysStateCommitmentChainVerifierChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type OasysStateCommitmentChainVerifierChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

// OasysStateCommitmentChainVerifierL2OutputOracleSetting is an auto generated low-level Go binding around an user-defined struct.
type OasysStateCommitmentChainVerifierL2OutputOracleSetting struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}

// OasysStateCommitmentChainVerifierOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type OasysStateCommitmentChainVerifierOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// SccverifierMetaData contains all meta data concerning the Sccverifier contract.
var SccverifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"l2ooAssertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ooSetting\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sccAssertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.L2OutputOracleSetting\",\"name\":\"_l2ooSetting\",\"type\":\"tuple\"}],\"name\":\"setL2ooSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b611dd4806100266000396000f3fe60806040523480156100115760006000fd5b506004361061008d5760003560e01c80639ac455721161005c5780639ac455721461011e578063a8d5c4d01461013a578063d17f1fb11461016d578063dea2619d146101895761008d565b80630bce93ac146100935780630c40acc4146100c757806374d5ecb7146100e65780637d2d9b63146101025761008d565b60006000fd5b6100ad60048036038101906100a891906114ad565b6101a5565b6040516100be959493929190611759565b60405180910390f35b6100cf610363565b6040516100dd9291906117d0565b60405180910390f35b61010060048036038101906100fb9190611457565b610397565b005b61011c600480360381019061011791906113d2565b6103b0565b005b610138600480360381019061013391906113d2565b610671565b005b610154600480360381019061014f91906114ad565b610932565b6040516101649493929190611705565b60405180910390f35b61018760048036038101906101829190611348565b610b0f565b005b6101a3600480360381019061019e9190611348565b610d56565b005b600160005081815481106101b857600080fd5b906000526020600020906006020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000505490806002016000506040518060600160405290816000820160005054600019166000191681526020016001820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016001820160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020015050908060040160005080546102cd90611a75565b80601f01602080910402602001604051908101604052809291908181526020018280546102f990611a75565b80156103465780601f1061031b57610100808354040283529160200191610346565b820191906000526020600020905b81548152906001019060200180831161032957829003601f168201915b5050505050908060050160009054906101000a900460ff16905085565b60026000508060000160005054908060010160009054906101000a90046fffffffffffffffffffffffffffffffff16905082565b80600260005081816103a99190611cfe565b9050505b50565b60016000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858036038101906103f39190611482565b81526020016104088585610f9d63ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050506060820151816004016000509080519060200190610555929190611040565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663931da5d785600260005060000160005054600260005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b81526004016105e493929190611824565b600060405180830381600087803b1580156105ff5760006000fd5b505af1158015610614573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f60405160405180910390a45b5050505050565b60016000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858036038101906106b49190611482565b81526020016106c98585610f9d63ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050506060820151816004016000509080519060200190610816929190611040565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663a511b01285600260005060000160005054600260005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b81526004016108a593929190611824565b600060405180830381600087803b1580156108c05760006000fd5b505af11580156108d5573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca360405160405180910390a45b5050505050565b6000600050818154811061094557600080fd5b906000526020600020906008020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000506040518060a00160405290816000820160005054815260200160018201600050546000191660001916815260200160028201600050548152602001600382016000505481526020016004820160005080546109e190611a75565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0d90611a75565b8015610a5a5780601f10610a2f57610100808354040283529160200191610a5a565b820191906000526020600020905b815481529060010190602001808311610a3d57829003601f168201915b50505050508152602001505090806006016000508054610a7990611a75565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa590611a75565b8015610af25780601f10610ac757610100808354040283529160200191610af2565b820191906000526020600020905b815481529060010190602001808311610ad557829003601f168201915b5050505050908060070160009054906101000a900460ff16905084565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001610b518585610f9d63ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610c3c929190611040565b5050506040820151816006016000509080519060200190610c5e929190611040565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663982bc5b0846000015185602001516040518363ffffffff1660e01b8152600401610cc49291906117fa565b600060405180830381600087803b158015610cdf5760006000fd5b505af1158015610cf4573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd8560200151604051610d4791906117b4565b60405180910390a35b50505050565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001610d988585610f9d63ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610e83929190611040565b5050506040820151816006016000509080519060200190610ea5929190611040565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663d1f93ca6846000015185602001516040518363ffffffff1660e01b8152600401610f0b9291906117fa565b600060405180830381600087803b158015610f265760006000fd5b505af1158015610f3b573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f58560200151604051610f8e91906117b4565b60405180910390a35b50505050565b60606000600090505b8383905081101561103857818484838181101515610fed577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002810190610fff919061185c565b604051602001611011939291906116de565b604051602081830303815290604052915081505b808061103090611adc565b915050610fa6565b505b92915050565b82805461104c90611a75565b90600052602060002090601f01602090048101928261106e57600085556110ba565b82601f1061108757805160ff19168380011785556110ba565b828001600101855582156110ba579182015b828111156110b95782518260005090905591602001919060010190611099565b5b5090506110c791906110cb565b5090565b6110d0565b808211156110ea57600081815060009055506001016110d0565b509056611d9d565b6000611105611100846118e1565b6118ba565b90508281526020810184848401111561111e5760006000fd5b6111298482856119e9565b505b9392505050565b60008135905061114181611d31565b5b92915050565b6000600083601f840112151561115e5760006000fd5b8235905067ffffffffffffffff8111156111785760006000fd5b6020830191508360208202830111156111915760006000fd5b5b9250929050565b6000813590506111a881611d4c565b5b92915050565b600082601f83011215156111c35760006000fd5b81356111d38482602086016110f2565b9150505b92915050565b600060a082840312156111f05760006000fd5b6111fa60a06118ba565b9050600061120a84828501611332565b600083015250602061121e84828501611199565b602083015250604061123284828501611332565b604083015250606061124684828501611332565b606083015250608082013567ffffffffffffffff8111156112675760006000fd5b611273848285016111af565b6080830152505b92915050565b6000604082840312156112935760006000fd5b8190505b92915050565b6000606082840312156112b05760006000fd5b8190505b92915050565b6000606082840312156112cd5760006000fd5b6112d760606118ba565b905060006112e784828501611199565b60008301525060206112fb8482850161131c565b602083015250604061130f8482850161131c565b6040830152505b92915050565b60008135905061132b81611d67565b5b92915050565b60008135905061134181611d82565b5b92915050565b6000600060006000606085870312156113615760006000fd5b600061136f87828801611132565b945050602085013567ffffffffffffffff81111561138d5760006000fd5b611399878288016111dd565b935050604085013567ffffffffffffffff8111156113b75760006000fd5b6113c387828801611148565b92509250505b92959194509250565b6000600060006000600060c086880312156113ed5760006000fd5b60006113fb88828901611132565b955050602061140c88828901611332565b945050604061141d8882890161129d565b93505060a086013567ffffffffffffffff81111561143b5760006000fd5b61144788828901611148565b92509250505b9295509295909350565b60006040828403121561146a5760006000fd5b600061147884828501611280565b9150505b92915050565b6000606082840312156114955760006000fd5b60006114a3848285016112ba565b9150505b92915050565b6000602082840312156114c05760006000fd5b60006114ce84828501611332565b9150505b92915050565b6114e18161194f565b82525b5050565b6114f181611962565b82525b5050565b6115018161196f565b82525b5050565b6115118161196f565b82525b5050565b60006115248385611943565b93506115318385846119e9565b82840190505b9392505050565b600061154982611913565b611553818561191f565b93506115638185602086016119f9565b61156c81611c36565b84019150505b92915050565b600061158382611913565b61158d8185611931565b935061159d8185602086016119f9565b6115a681611c36565b84019150505b92915050565b60006115bd82611913565b6115c78185611943565b93506115d78185602086016119f9565b8084019150505b92915050565b600060a0830160008301516115fc60008601826116be565b50602083015161160f60208601826114f8565b50604083015161162260408601826116be565b50606083015161163560608601826116be565b506080830151848203608086015261164d828261153e565b915050809150505b92915050565b60608201600082015161167160008501826114f8565b506020820151611684602085018261169e565b506040820151611697604085018261169e565b50505b5050565b6116a78161197a565b82525b5050565b6116b78161197a565b82525b5050565b6116c7816119b8565b82525b5050565b6116d7816119b8565b82525b5050565b60006116ea82866115b2565b91506116f7828486611518565b91508190505b949350505050565b600060808201905061171a60008301876114d8565b818103602083015261172c81866115e4565b905081810360408301526117408185611578565b905061174f60608301846114e8565b5b95945050505050565b600060e08201905061176e60008301886114d8565b61177b60208301876116ce565b611788604083018661165b565b81810360a083015261179a8185611578565b90506117a960c08301846114e8565b5b9695505050505050565b60006020820190506117c96000830184611508565b5b92915050565b60006040820190506117e56000830185611508565b6117f260208301846116ae565b5b9392505050565b600060408201905061180f60008301856116ce565b61181c6020830184611508565b5b9392505050565b600060608201905061183960008301866116ce565b6118466020830185611508565b61185360408301846116ae565b5b949350505050565b600060008335600160200384360303811215156118795760006000fd5b80840192508235915067ffffffffffffffff8211156118985760006000fd5b6020830192506001820236038313156118b15760006000fd5b505b9250929050565b60006118c46118d6565b90506118d08282611aaa565b5b919050565b600060405190505b90565b600067ffffffffffffffff8211156118fc576118fb611bb9565b5b61190582611c36565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b600061195a82611997565b90505b919050565b600081151590505b919050565b60008190505b919050565b60006fffffffffffffffffffffffffffffffff821690505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b60006119ce8261196f565b90505b919050565b60006119e18261197a565b90505b919050565b828183376000838301525b505050565b60005b83811015611a185780820151818401525b6020810190506119fc565b83811115611a27576000848401525b505b505050565b600081016000830180611a4081611c08565b9050611a4c8184611cda565b505050600181016020830180611a6181611c1f565b9050611a6d8184611d0d565b5050505b5050565b600060028204905060018216801515611a8f57607f821691505b60208210811415611aa357611aa2611b88565b5b505b919050565b611ab382611c36565b810181811067ffffffffffffffff82111715611ad257611ad1611bb9565b5b80604052505b5050565b6000611ae7826119b8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611b1a57611b19611b57565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000611bf582611c56565b90505b919050565b60008190505b919050565b60008135611c1581611d4c565b809150505b919050565b60008135611c2c81611d67565b809150505b919050565b6000601f19601f83011690505b919050565b60008160001b90505b919050565b60008160001c90505b919050565b60006fffffffffffffffffffffffffffffffff611c8084611c48565b935080198316925080841683179150505b92915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611cc384611c48565b935080198316925080841683179150505b92915050565b611ce3826119c3565b611cf6611cef82611bea565b8354611c97565b8255505b5050565b611d088282611a2e565b5b5050565b611d16826119d6565b611d29611d2282611bfd565b8354611c64565b8255505b5050565b611d3a8161194f565b81141515611d485760006000fd5b5b50565b611d558161196f565b81141515611d635760006000fd5b5b50565b611d708161197a565b81141515611d7e5760006000fd5b5b50565b611d8b816119b8565b81141515611d995760006000fd5b5b50565bfea2646970667358221220be77c07d3322bb1401e92416a567a9d3b0c33c4e93de7156963d0da98a38015464736f6c63430008020033",
}

// SccverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use SccverifierMetaData.ABI instead.
var SccverifierABI = SccverifierMetaData.ABI

// SccverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SccverifierMetaData.Bin instead.
var SccverifierBin = SccverifierMetaData.Bin

// DeploySccverifier deploys a new Ethereum contract, binding an instance of Sccverifier to it.
func DeploySccverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sccverifier, error) {
	parsed, err := SccverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SccverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sccverifier{SccverifierCaller: SccverifierCaller{contract: contract}, SccverifierTransactor: SccverifierTransactor{contract: contract}, SccverifierFilterer: SccverifierFilterer{contract: contract}}, nil
}

// Sccverifier is an auto generated Go binding around an Ethereum contract.
type Sccverifier struct {
	SccverifierCaller     // Read-only binding to the contract
	SccverifierTransactor // Write-only binding to the contract
	SccverifierFilterer   // Log filterer for contract events
}

// SccverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SccverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SccverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SccverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SccverifierSession struct {
	Contract     *Sccverifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SccverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SccverifierCallerSession struct {
	Contract *SccverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SccverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SccverifierTransactorSession struct {
	Contract     *SccverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SccverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SccverifierRaw struct {
	Contract *Sccverifier // Generic contract binding to access the raw methods on
}

// SccverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SccverifierCallerRaw struct {
	Contract *SccverifierCaller // Generic read-only contract binding to access the raw methods on
}

// SccverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SccverifierTransactorRaw struct {
	Contract *SccverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSccverifier creates a new instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifier(address common.Address, backend bind.ContractBackend) (*Sccverifier, error) {
	contract, err := bindSccverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sccverifier{SccverifierCaller: SccverifierCaller{contract: contract}, SccverifierTransactor: SccverifierTransactor{contract: contract}, SccverifierFilterer: SccverifierFilterer{contract: contract}}, nil
}

// NewSccverifierCaller creates a new read-only instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierCaller(address common.Address, caller bind.ContractCaller) (*SccverifierCaller, error) {
	contract, err := bindSccverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SccverifierCaller{contract: contract}, nil
}

// NewSccverifierTransactor creates a new write-only instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SccverifierTransactor, error) {
	contract, err := bindSccverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SccverifierTransactor{contract: contract}, nil
}

// NewSccverifierFilterer creates a new log filterer instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SccverifierFilterer, error) {
	contract, err := bindSccverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SccverifierFilterer{contract: contract}, nil
}

// bindSccverifier binds a generic wrapper to an already deployed contract.
func bindSccverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SccverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sccverifier *SccverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sccverifier.Contract.SccverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sccverifier *SccverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sccverifier.Contract.SccverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sccverifier *SccverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sccverifier.Contract.SccverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sccverifier *SccverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sccverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sccverifier *SccverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sccverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sccverifier *SccverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sccverifier.Contract.contract.Transact(opts, method, params...)
}

// L2ooAssertLogs is a free data retrieval call binding the contract method 0x0bce93ac.
//
// Solidity: function l2ooAssertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCaller) L2ooAssertLogs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysStateCommitmentChainVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "l2ooAssertLogs", arg0)

	outstruct := new(struct {
		L2OutputOracle common.Address
		L2OutputIndex  *big.Int
		L2Output       OasysStateCommitmentChainVerifierOutputProposal
		Signatures     []byte
		Approve        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.L2OutputOracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.L2OutputIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2Output = *abi.ConvertType(out[2], new(OasysStateCommitmentChainVerifierOutputProposal)).(*OasysStateCommitmentChainVerifierOutputProposal)
	outstruct.Signatures = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Approve = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// L2ooAssertLogs is a free data retrieval call binding the contract method 0x0bce93ac.
//
// Solidity: function l2ooAssertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_Sccverifier *SccverifierSession) L2ooAssertLogs(arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysStateCommitmentChainVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	return _Sccverifier.Contract.L2ooAssertLogs(&_Sccverifier.CallOpts, arg0)
}

// L2ooAssertLogs is a free data retrieval call binding the contract method 0x0bce93ac.
//
// Solidity: function l2ooAssertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCallerSession) L2ooAssertLogs(arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysStateCommitmentChainVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	return _Sccverifier.Contract.L2ooAssertLogs(&_Sccverifier.CallOpts, arg0)
}

// L2ooSetting is a free data retrieval call binding the contract method 0x0c40acc4.
//
// Solidity: function l2ooSetting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_Sccverifier *SccverifierCaller) L2ooSetting(opts *bind.CallOpts) (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "l2ooSetting")

	outstruct := new(struct {
		OutputRoot    [32]byte
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutputRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// L2ooSetting is a free data retrieval call binding the contract method 0x0c40acc4.
//
// Solidity: function l2ooSetting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_Sccverifier *SccverifierSession) L2ooSetting() (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _Sccverifier.Contract.L2ooSetting(&_Sccverifier.CallOpts)
}

// L2ooSetting is a free data retrieval call binding the contract method 0x0c40acc4.
//
// Solidity: function l2ooSetting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_Sccverifier *SccverifierCallerSession) L2ooSetting() (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _Sccverifier.Contract.L2ooSetting(&_Sccverifier.CallOpts)
}

// SccAssertLogs is a free data retrieval call binding the contract method 0xa8d5c4d0.
//
// Solidity: function sccAssertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCaller) SccAssertLogs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "sccAssertLogs", arg0)

	outstruct := new(struct {
		StateCommitmentChain common.Address
		BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
		Signatures           []byte
		Approve              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateCommitmentChain = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BatchHeader = *abi.ConvertType(out[1], new(OasysStateCommitmentChainVerifierChainBatchHeader)).(*OasysStateCommitmentChainVerifierChainBatchHeader)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Approve = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// SccAssertLogs is a free data retrieval call binding the contract method 0xa8d5c4d0.
//
// Solidity: function sccAssertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierSession) SccAssertLogs(arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	return _Sccverifier.Contract.SccAssertLogs(&_Sccverifier.CallOpts, arg0)
}

// SccAssertLogs is a free data retrieval call binding the contract method 0xa8d5c4d0.
//
// Solidity: function sccAssertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCallerSession) SccAssertLogs(arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	return _Sccverifier.Contract.SccAssertLogs(&_Sccverifier.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Approve(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "approve", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Approve0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "approve0", stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Approve0(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve0(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Approve0(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve0(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Reject(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "reject", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject(&_Sccverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysStateCommitmentChainVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject(&_Sccverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Reject0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "reject0", stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Reject0(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject0(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Reject0(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject0(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_Sccverifier *SccverifierTransactor) SetL2ooSetting(opts *bind.TransactOpts, _l2ooSetting OasysStateCommitmentChainVerifierL2OutputOracleSetting) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "setL2ooSetting", _l2ooSetting)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_Sccverifier *SccverifierSession) SetL2ooSetting(_l2ooSetting OasysStateCommitmentChainVerifierL2OutputOracleSetting) (*types.Transaction, error) {
	return _Sccverifier.Contract.SetL2ooSetting(&_Sccverifier.TransactOpts, _l2ooSetting)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_Sccverifier *SccverifierTransactorSession) SetL2ooSetting(_l2ooSetting OasysStateCommitmentChainVerifierL2OutputOracleSetting) (*types.Transaction, error) {
	return _Sccverifier.Contract.SetL2ooSetting(&_Sccverifier.TransactOpts, _l2ooSetting)
}

// SccverifierL2OutputApprovedIterator is returned from FilterL2OutputApproved and is used to iterate over the raw logs and unpacked data for L2OutputApproved events raised by the Sccverifier contract.
type SccverifierL2OutputApprovedIterator struct {
	Event *SccverifierL2OutputApproved // Event containing the contract specifics and raw log

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
func (it *SccverifierL2OutputApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierL2OutputApproved)
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
		it.Event = new(SccverifierL2OutputApproved)
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
func (it *SccverifierL2OutputApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierL2OutputApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierL2OutputApproved represents a L2OutputApproved event raised by the Sccverifier contract.
type SccverifierL2OutputApproved struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputApproved is a free log retrieval operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_Sccverifier *SccverifierFilterer) FilterL2OutputApproved(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*SccverifierL2OutputApprovedIterator, error) {

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

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierL2OutputApprovedIterator{contract: _Sccverifier.contract, event: "L2OutputApproved", logs: logs, sub: sub}, nil
}

// WatchL2OutputApproved is a free log subscription operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_Sccverifier *SccverifierFilterer) WatchL2OutputApproved(opts *bind.WatchOpts, sink chan<- *SccverifierL2OutputApproved, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierL2OutputApproved)
				if err := _Sccverifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
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
func (_Sccverifier *SccverifierFilterer) ParseL2OutputApproved(log types.Log) (*SccverifierL2OutputApproved, error) {
	event := new(SccverifierL2OutputApproved)
	if err := _Sccverifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccverifierL2OutputRejectedIterator is returned from FilterL2OutputRejected and is used to iterate over the raw logs and unpacked data for L2OutputRejected events raised by the Sccverifier contract.
type SccverifierL2OutputRejectedIterator struct {
	Event *SccverifierL2OutputRejected // Event containing the contract specifics and raw log

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
func (it *SccverifierL2OutputRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierL2OutputRejected)
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
		it.Event = new(SccverifierL2OutputRejected)
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
func (it *SccverifierL2OutputRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierL2OutputRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierL2OutputRejected represents a L2OutputRejected event raised by the Sccverifier contract.
type SccverifierL2OutputRejected struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputRejected is a free log retrieval operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_Sccverifier *SccverifierFilterer) FilterL2OutputRejected(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*SccverifierL2OutputRejectedIterator, error) {

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

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierL2OutputRejectedIterator{contract: _Sccverifier.contract, event: "L2OutputRejected", logs: logs, sub: sub}, nil
}

// WatchL2OutputRejected is a free log subscription operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_Sccverifier *SccverifierFilterer) WatchL2OutputRejected(opts *bind.WatchOpts, sink chan<- *SccverifierL2OutputRejected, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierL2OutputRejected)
				if err := _Sccverifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
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
func (_Sccverifier *SccverifierFilterer) ParseL2OutputRejected(log types.Log) (*SccverifierL2OutputRejected, error) {
	event := new(SccverifierL2OutputRejected)
	if err := _Sccverifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccverifierStateBatchApprovedIterator is returned from FilterStateBatchApproved and is used to iterate over the raw logs and unpacked data for StateBatchApproved events raised by the Sccverifier contract.
type SccverifierStateBatchApprovedIterator struct {
	Event *SccverifierStateBatchApproved // Event containing the contract specifics and raw log

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
func (it *SccverifierStateBatchApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierStateBatchApproved)
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
		it.Event = new(SccverifierStateBatchApproved)
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
func (it *SccverifierStateBatchApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierStateBatchApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierStateBatchApproved represents a StateBatchApproved event raised by the Sccverifier contract.
type SccverifierStateBatchApproved struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchApproved is a free log retrieval operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) FilterStateBatchApproved(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*SccverifierStateBatchApprovedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierStateBatchApprovedIterator{contract: _Sccverifier.contract, event: "StateBatchApproved", logs: logs, sub: sub}, nil
}

// WatchStateBatchApproved is a free log subscription operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) WatchStateBatchApproved(opts *bind.WatchOpts, sink chan<- *SccverifierStateBatchApproved, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierStateBatchApproved)
				if err := _Sccverifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
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
func (_Sccverifier *SccverifierFilterer) ParseStateBatchApproved(log types.Log) (*SccverifierStateBatchApproved, error) {
	event := new(SccverifierStateBatchApproved)
	if err := _Sccverifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccverifierStateBatchRejectedIterator is returned from FilterStateBatchRejected and is used to iterate over the raw logs and unpacked data for StateBatchRejected events raised by the Sccverifier contract.
type SccverifierStateBatchRejectedIterator struct {
	Event *SccverifierStateBatchRejected // Event containing the contract specifics and raw log

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
func (it *SccverifierStateBatchRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierStateBatchRejected)
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
		it.Event = new(SccverifierStateBatchRejected)
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
func (it *SccverifierStateBatchRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierStateBatchRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierStateBatchRejected represents a StateBatchRejected event raised by the Sccverifier contract.
type SccverifierStateBatchRejected struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchRejected is a free log retrieval operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) FilterStateBatchRejected(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*SccverifierStateBatchRejectedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierStateBatchRejectedIterator{contract: _Sccverifier.contract, event: "StateBatchRejected", logs: logs, sub: sub}, nil
}

// WatchStateBatchRejected is a free log subscription operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) WatchStateBatchRejected(opts *bind.WatchOpts, sink chan<- *SccverifierStateBatchRejected, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierStateBatchRejected)
				if err := _Sccverifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
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
func (_Sccverifier *SccverifierFilterer) ParseStateBatchRejected(log types.Log) (*SccverifierStateBatchRejected, error) {
	event := new(SccverifierStateBatchRejected)
	if err := _Sccverifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
