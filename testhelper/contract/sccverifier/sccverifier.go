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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"l2ooAssertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ooAssertLogsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ooSetting\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sccAssertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sccAssertLogsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.L2OutputOracleSetting\",\"name\":\"_l2ooSetting\",\"type\":\"tuple\"}],\"name\":\"setL2ooSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b611e6c806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100a35760003560e01c80639ac45572116100675780639ac4557214610152578063a8d5c4d01461016e578063b32213bf146101a1578063d17f1fb1146101bf578063dea2619d146101db576100a3565b80630bce93ac146100a95780630c40acc4146100dd578063358b3cbb146100fc57806374d5ecb71461011a5780637d2d9b6314610136576100a3565b60006000fd5b6100c360048036038101906100be9190611529565b6101f7565b6040516100d49594939291906117d5565b60405180910390f35b6100e56103b5565b6040516100f392919061184c565b60405180910390f35b6101046103e9565b6040516101119190611876565b60405180910390f35b610134600480360381019061012f91906114d3565b6103fe565b005b610150600480360381019061014b919061144e565b610417565b005b61016c6004803603810190610167919061144e565b6106d8565b005b61018860048036038101906101839190611529565b610999565b6040516101989493929190611781565b60405180910390f35b6101a9610b76565b6040516101b69190611876565b60405180910390f35b6101d960048036038101906101d491906113c4565b610b8b565b005b6101f560048036038101906101f091906113c4565b610dd2565b005b6001600050818154811061020a57600080fd5b906000526020600020906006020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000505490806002016000506040518060600160405290816000820160005054600019166000191681526020016001820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016001820160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200150509080600401600050805461031f90611b0d565b80601f016020809104026020016040519081016040528092919081815260200182805461034b90611b0d565b80156103985780601f1061036d57610100808354040283529160200191610398565b820191906000526020600020905b81548152906001019060200180831161037b57829003601f168201915b5050505050908060050160009054906101000a900460ff16905085565b60026000508060000160005054908060010160009054906101000a90046fffffffffffffffffffffffffffffffff16905082565b600060016000508054905090506103fb565b90565b80600260005081816104109190611d96565b9050505b50565b60016000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018580360381019061045a91906114fe565b815260200161046f858561101963ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550505060608201518160040160005090805190602001906105bc9291906110bc565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663931da5d785600260005060000160005054600260005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b815260040161064b939291906118bc565b600060405180830381600087803b1580156106665760006000fd5b505af115801561067b573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f60405160405180910390a45b5050505050565b60016000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018580360381019061071b91906114fe565b8152602001610730858561101963ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055505050606082015181600401600050908051906020019061087d9291906110bc565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663a511b01285600260005060000160005054600260005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b815260040161090c939291906118bc565b600060405180830381600087803b1580156109275760006000fd5b505af115801561093c573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca360405160405180910390a45b5050505050565b600060005081815481106109ac57600080fd5b906000526020600020906008020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000506040518060a0016040529081600082016000505481526020016001820160005054600019166000191681526020016002820160005054815260200160038201600050548152602001600482016000508054610a4890611b0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7490611b0d565b8015610ac15780601f10610a9657610100808354040283529160200191610ac1565b820191906000526020600020905b815481529060010190602001808311610aa457829003601f168201915b50505050508152602001505090806006016000508054610ae090611b0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0c90611b0d565b8015610b595780601f10610b2e57610100808354040283529160200191610b59565b820191906000526020600020905b815481529060010190602001808311610b3c57829003601f168201915b5050505050908060070160009054906101000a900460ff16905084565b60006000600050805490509050610b88565b90565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001610bcd858561101963ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610cb89291906110bc565b5050506040820151816006016000509080519060200190610cda9291906110bc565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663982bc5b0846000015185602001516040518363ffffffff1660e01b8152600401610d40929190611892565b600060405180830381600087803b158015610d5b5760006000fd5b505af1158015610d70573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd8560200151604051610dc39190611830565b60405180910390a35b50505050565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001610e14858561101963ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610eff9291906110bc565b5050506040820151816006016000509080519060200190610f219291906110bc565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663d1f93ca6846000015185602001516040518363ffffffff1660e01b8152600401610f87929190611892565b600060405180830381600087803b158015610fa25760006000fd5b505af1158015610fb7573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5856020015160405161100a9190611830565b60405180910390a35b50505050565b60606000600090505b838390508110156110b457818484838181101515611069577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200281019061107b91906118f4565b60405160200161108d9392919061175a565b604051602081830303815290604052915081505b80806110ac90611b74565b915050611022565b505b92915050565b8280546110c890611b0d565b90600052602060002090601f0160209004810192826110ea5760008555611136565b82601f1061110357805160ff1916838001178555611136565b82800160010185558215611136579182015b828111156111355782518260005090905591602001919060010190611115565b5b5090506111439190611147565b5090565b61114c565b80821115611166576000818150600090555060010161114c565b509056611e35565b600061118161117c84611979565b611952565b90508281526020810184848401111561119a5760006000fd5b6111a5848285611a81565b505b9392505050565b6000813590506111bd81611dc9565b5b92915050565b6000600083601f84011215156111da5760006000fd5b8235905067ffffffffffffffff8111156111f45760006000fd5b60208301915083602082028301111561120d5760006000fd5b5b9250929050565b60008135905061122481611de4565b5b92915050565b600082601f830112151561123f5760006000fd5b813561124f84826020860161116e565b9150505b92915050565b600060a0828403121561126c5760006000fd5b61127660a0611952565b90506000611286848285016113ae565b600083015250602061129a84828501611215565b60208301525060406112ae848285016113ae565b60408301525060606112c2848285016113ae565b606083015250608082013567ffffffffffffffff8111156112e35760006000fd5b6112ef8482850161122b565b6080830152505b92915050565b60006040828403121561130f5760006000fd5b8190505b92915050565b60006060828403121561132c5760006000fd5b8190505b92915050565b6000606082840312156113495760006000fd5b6113536060611952565b9050600061136384828501611215565b600083015250602061137784828501611398565b602083015250604061138b84828501611398565b6040830152505b92915050565b6000813590506113a781611dff565b5b92915050565b6000813590506113bd81611e1a565b5b92915050565b6000600060006000606085870312156113dd5760006000fd5b60006113eb878288016111ae565b945050602085013567ffffffffffffffff8111156114095760006000fd5b61141587828801611259565b935050604085013567ffffffffffffffff8111156114335760006000fd5b61143f878288016111c4565b92509250505b92959194509250565b6000600060006000600060c086880312156114695760006000fd5b6000611477888289016111ae565b9550506020611488888289016113ae565b945050604061149988828901611319565b93505060a086013567ffffffffffffffff8111156114b75760006000fd5b6114c3888289016111c4565b92509250505b9295509295909350565b6000604082840312156114e65760006000fd5b60006114f4848285016112fc565b9150505b92915050565b6000606082840312156115115760006000fd5b600061151f84828501611336565b9150505b92915050565b60006020828403121561153c5760006000fd5b600061154a848285016113ae565b9150505b92915050565b61155d816119e7565b82525b5050565b61156d816119fa565b82525b5050565b61157d81611a07565b82525b5050565b61158d81611a07565b82525b5050565b60006115a083856119db565b93506115ad838584611a81565b82840190505b9392505050565b60006115c5826119ab565b6115cf81856119b7565b93506115df818560208601611a91565b6115e881611cce565b84019150505b92915050565b60006115ff826119ab565b61160981856119c9565b9350611619818560208601611a91565b61162281611cce565b84019150505b92915050565b6000611639826119ab565b61164381856119db565b9350611653818560208601611a91565b8084019150505b92915050565b600060a083016000830151611678600086018261173a565b50602083015161168b6020860182611574565b50604083015161169e604086018261173a565b5060608301516116b1606086018261173a565b50608083015184820360808601526116c982826115ba565b915050809150505b92915050565b6060820160008201516116ed6000850182611574565b506020820151611700602085018261171a565b506040820151611713604085018261171a565b50505b5050565b61172381611a12565b82525b5050565b61173381611a12565b82525b5050565b61174381611a50565b82525b5050565b61175381611a50565b82525b5050565b6000611766828661162e565b9150611773828486611594565b91508190505b949350505050565b60006080820190506117966000830187611554565b81810360208301526117a88186611660565b905081810360408301526117bc81856115f4565b90506117cb6060830184611564565b5b95945050505050565b600060e0820190506117ea6000830188611554565b6117f7602083018761174a565b61180460408301866116d7565b81810360a083015261181681856115f4565b905061182560c0830184611564565b5b9695505050505050565b60006020820190506118456000830184611584565b5b92915050565b60006040820190506118616000830185611584565b61186e602083018461172a565b5b9392505050565b600060208201905061188b600083018461174a565b5b92915050565b60006040820190506118a7600083018561174a565b6118b46020830184611584565b5b9392505050565b60006060820190506118d1600083018661174a565b6118de6020830185611584565b6118eb604083018461172a565b5b949350505050565b600060008335600160200384360303811215156119115760006000fd5b80840192508235915067ffffffffffffffff8211156119305760006000fd5b6020830192506001820236038313156119495760006000fd5b505b9250929050565b600061195c61196e565b90506119688282611b42565b5b919050565b600060405190505b90565b600067ffffffffffffffff82111561199457611993611c51565b5b61199d82611cce565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60006119f282611a2f565b90505b919050565b600081151590505b919050565b60008190505b919050565b60006fffffffffffffffffffffffffffffffff821690505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6000611a6682611a07565b90505b919050565b6000611a7982611a12565b90505b919050565b828183376000838301525b505050565b60005b83811015611ab05780820151818401525b602081019050611a94565b83811115611abf576000848401525b505b505050565b600081016000830180611ad881611ca0565b9050611ae48184611d72565b505050600181016020830180611af981611cb7565b9050611b058184611da5565b5050505b5050565b600060028204905060018216801515611b2757607f821691505b60208210811415611b3b57611b3a611c20565b5b505b919050565b611b4b82611cce565b810181811067ffffffffffffffff82111715611b6a57611b69611c51565b5b80604052505b5050565b6000611b7f82611a50565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611bb257611bb1611bef565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000611c8d82611cee565b90505b919050565b60008190505b919050565b60008135611cad81611de4565b809150505b919050565b60008135611cc481611dff565b809150505b919050565b6000601f19601f83011690505b919050565b60008160001b90505b919050565b60008160001c90505b919050565b60006fffffffffffffffffffffffffffffffff611d1884611ce0565b935080198316925080841683179150505b92915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611d5b84611ce0565b935080198316925080841683179150505b92915050565b611d7b82611a5b565b611d8e611d8782611c82565b8354611d2f565b8255505b5050565b611da08282611ac6565b5b5050565b611dae82611a6e565b611dc1611dba82611c95565b8354611cfc565b8255505b5050565b611dd2816119e7565b81141515611de05760006000fd5b5b50565b611ded81611a07565b81141515611dfb5760006000fd5b5b50565b611e0881611a12565b81141515611e165760006000fd5b5b50565b611e2381611a50565b81141515611e315760006000fd5b5b50565bfea2646970667358221220eba969a55a7f4987a44f6d1153d2dac752af69cfb02adae7bc9ec05876ce234e64736f6c63430008020033",
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

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCaller) L2ooAssertLogsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "l2ooAssertLogsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierSession) L2ooAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.L2ooAssertLogsLen(&_Sccverifier.CallOpts)
}

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCallerSession) L2ooAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.L2ooAssertLogsLen(&_Sccverifier.CallOpts)
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

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCaller) SccAssertLogsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "sccAssertLogsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierSession) SccAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.SccAssertLogsLen(&_Sccverifier.CallOpts)
}

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCallerSession) SccAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.SccAssertLogsLen(&_Sccverifier.CallOpts)
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
