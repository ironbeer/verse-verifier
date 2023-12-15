// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2oo

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

// TypesOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// OasysL2OutputOracleMetaData contains all meta data concerning the OasysL2OutputOracle contract.
var OasysL2OutputOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"L2OutputFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"L2OutputVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevNextOutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newNextOutputIndex\",\"type\":\"uint256\"}],\"name\":\"OutputsDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHALLENGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"deleteL2Outputs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"outputRootProof\",\"type\":\"tuple\"}],\"name\":\"failVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationPeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputAfter\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputIndexAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"isOutputFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextVerifyIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l1BlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l1BlockNumber\",\"type\":\"uint256\"}],\"name\":\"proposeL2Output\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"outputRootProof\",\"type\":\"tuple\"}],\"name\":\"succeedVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200214738038062002147833981016040819052620000359162000370565b8686868686868660008611620000b85760405162461bcd60e51b815260206004820152603460248201527f4c324f75747075744f7261636c653a204c3220626c6f636b2074696d65206d7560448201527f73742062652067726561746572207468616e203000000000000000000000000060648201526084015b60405180910390fd5b60008711620001305760405162461bcd60e51b815260206004820152603a60248201527f4c324f75747075744f7261636c653a207375626d697373696f6e20696e74657260448201527f76616c206d7573742062652067726561746572207468616e20300000000000006064820152608401620000af565b608087905260a08690526001600160a01b0380841660e052821660c0526101008190526200015f858562000173565b5050505050505050505050505050620003d8565b6200018a82826200018e60201b620014c71760201c565b5050565b600054610100900460ff1615808015620001af5750600054600160ff909116105b80620001df5750620001cc306200034460201b620017141760201c565b158015620001df575060005460ff166001145b620002445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620000af565b6000805460ff19166001179055801562000268576000805461ff0019166101001790555b42821115620002ee5760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201526374696d6560e01b608482015260a401620000af565b6002829055600183905580156200033f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b80516001600160a01b03811681146200036b57600080fd5b919050565b600080600080600080600060e0888a0312156200038c57600080fd5b87519650602088015195506040880151945060608801519350620003b36080890162000353565b9250620003c360a0890162000353565b915060c0880151905092959891949750929550565b60805160a05160c05160e05161010051611cdd6200046a6000396000818161057f0152818161065d01526117a20152600081816105180152818161054e0152610f09015260008181610265015281816103890152610c280152600081816101d801528181610456015261144001526000818161023401528181610607015281816106f601526114880152611cdd6000f3fe6080604052600436106101c15760003560e01c806388786272116100f7578063bffa7f0f11610095578063dcec334811610064578063dcec3348146105e3578063e1a41bcf146105f8578063e4a301161461062b578063f4daa2911461064b57600080fd5b8063bffa7f0f1461053c578063ce5db8d614610570578063cf8e5cf0146105a3578063d1de856c146105c357600080fd5b806399f10c19116100d157806399f10c191461047a5780639aaab6481461049a578063a25ae557146104ad578063a8e4fb901461050957600080fd5b8063887862721461041157806389c44cbb1461042757806393991af31461044757600080fd5b806369f16eec116101645780636b4d98dd1161013e5780636b4d98dd146103775780636dbffb78146103ab57806370872aa5146103db5780637f006420146103f157600080fd5b806369f16eec146103375780636abcf5631461034c5780636b405db01461036157600080fd5b8063534db0e2116101a0578063534db0e21461025657806354fd4d50146102aa578063594098c614610300578063657c10ff1461031557600080fd5b80622134cc146101c65780634599c7881461020d578063529933df14610222575b600080fd5b3480156101d257600080fd5b506101fa7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b34801561021957600080fd5b506101fa61067f565b34801561022e57600080fd5b506101fa7f000000000000000000000000000000000000000000000000000000000000000081565b34801561026257600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610204565b3480156102b657600080fd5b506102f36040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161020491906119d0565b34801561030c57600080fd5b506101fa6106f2565b34801561032157600080fd5b50610335610330366004611a43565b610722565b005b34801561034357600080fd5b506101fa6109d9565b34801561035857600080fd5b506003546101fa565b34801561036d57600080fd5b506101fa60045481565b34801561038357600080fd5b506102857f000000000000000000000000000000000000000000000000000000000000000081565b3480156103b757600080fd5b506103cb6103c6366004611a99565b6109eb565b6040519015158152602001610204565b3480156103e757600080fd5b506101fa60015481565b3480156103fd57600080fd5b506101fa61040c366004611a99565b6109fc565b34801561041d57600080fd5b506101fa60025481565b34801561043357600080fd5b50610335610442366004611a99565b610c10565b34801561045357600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101fa565b34801561048657600080fd5b50610335610495366004611a43565b610ce1565b6103356104a8366004611ab2565b610ef1565b3480156104b957600080fd5b506104cd6104c8366004611a99565b611370565b60408051825181526020808401516fffffffffffffffffffffffffffffffff908116918301919091529282015190921690820152606001610204565b34801561051557600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610285565b34801561054857600080fd5b506102857f000000000000000000000000000000000000000000000000000000000000000081565b34801561057c57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101fa565b3480156105af57600080fd5b506104cd6105be366004611a99565b611404565b3480156105cf57600080fd5b506101fa6105de366004611a99565b61143c565b3480156105ef57600080fd5b506101fa611484565b34801561060457600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101fa565b34801561063757600080fd5b50610335610646366004611ae4565b6114b9565b34801561065757600080fd5b506101fa7f000000000000000000000000000000000000000000000000000000000000000081565b600354600090156106e9576003805461069a90600190611b35565b815481106106aa576106aa611b4c565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60007f00000000000000000000000000000000000000000000000000000000000000006004546106ed9190611b7b565b33735200000000000000000000000000000000000014146107ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4c324f75747075744f7261636c653a20496e76616c6964206d6573736167652060448201527f73656e646572000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000600383815481106107df576107df611b4c565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff80821694840194909452700100000000000000000000000000000000900490921691810191909152905061085661085136849003840184611bb8565b611730565b8151146108e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4c324f75747075744f7261636c653a20696e76616c6964206f7574707574207260448201527f6f6f742070726f6f66000000000000000000000000000000000000000000000060648201526084016107c1565b6004548314610976576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4c324f75747075744f7261636c653a20496e76616c6964204c32206f7574707560448201527f7420696e6465780000000000000000000000000000000000000000000000000060648201526084016107c1565b6004805490600061098683611c45565b919050555080604001516fffffffffffffffffffffffffffffffff168160000151847f43559e36255b2cd130c34b2551834b85ae8fb206cbe183be98ff8c8c44c2250a60405160405180910390a4505050565b6003546000906106ed90600190611b35565b60006109f68261178c565b92915050565b6000610a0661067f565b821115610abb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a4016107c1565b600354610b70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e2070726f706f7360648201527f6564207965740000000000000000000000000000000000000000000000000000608482015260a4016107c1565b6003546000905b80821015610c095760006002610b8d8385611c7d565b610b979190611c95565b90508460038281548110610bad57610bad611b4c565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff161015610bff57610bf8816001611c7d565b9250610c03565b8091505b50610b77565b5092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206368616c6c656e60448201527f67657220616464726573732063616e2064656c657465206f757470757473000060648201526084016107c1565b610cde8161181a565b50565b3373520000000000000000000000000000000000001414610d84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4c324f75747075744f7261636c653a20496e76616c6964206d6573736167652060448201527f73656e646572000000000000000000000000000000000000000000000000000060648201526084016107c1565b600060038381548110610d9957610d99611b4c565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff808216948401949094527001000000000000000000000000000000009004909216918101919091529050610e0b61085136849003840184611bb8565b815114610e9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4c324f75747075744f7261636c653a20696e76616c6964206f7574707574207260448201527f6f6f742070726f6f66000000000000000000000000000000000000000000000060648201526084016107c1565b610ea38361181a565b80604001516fffffffffffffffffffffffffffffffff168160000151847fbecbaa918310b9e773752f3b22d0cd05e91d296a3ca04a065d1cc0111fba2bca60405160405180910390a4505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610fdc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c79207468652070726f706f736560448201527f7220616464726573732063616e2070726f706f7365206e6577206f757470757460648201527f7300000000000000000000000000000000000000000000000000000000000000608482015260a4016107c1565b610fe4611484565b8314611098576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a4016107c1565b426110a28461143c565b1061112f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324f75747075744f7261636c653a2063616e6e6f742070726f706f7365204c60448201527f32206f757470757420696e20746865206675747572650000000000000000000060648201526084016107c1565b836111bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c324f75747075744f7261636c653a204c32206f75747075742070726f706f7360448201527f616c2063616e6e6f7420626520746865207a65726f206861736800000000000060648201526084016107c1565b81156112785781814014611278576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a4016107c1565b8261128260035490565b857fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2426040516112b491815260200190565b60405180910390a45050604080516060810182529283526fffffffffffffffffffffffffffffffff4281166020850190815292811691840191825260038054600181018255600091909152935160029094027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810194909455915190518216700100000000000000000000000000000000029116177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c90910155565b6040805160608101825260008082526020820181905291810191909152600382815481106113a0576113a0611b4c565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b6040805160608101825260008082526020820181905291810191909152600361142c836109fc565b815481106113a0576113a0611b4c565b60007f00000000000000000000000000000000000000000000000000000000000000006001548361146d9190611b35565b6114779190611b7b565b6002546109f69190611c7d565b60007f00000000000000000000000000000000000000000000000000000000000000006114af61067f565b6106ed9190611c7d565b6114c382826114c7565b5050565b600054610100900460ff16158080156114e75750600054600160ff909116105b806115015750303b158015611501575060005460ff166001145b61158d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016107c1565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156115eb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156116a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a4016107c1565b60028290556001839055801561170f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000816000015182602001518360400151846060015160405160200161176f949392919093845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b60006004548210156117a057506001919050565b7f0000000000000000000000000000000000000000000000000000000000000000600383815481106117d4576117d4611b4c565b6000918252602090912060016002909202010154611804906fffffffffffffffffffffffffffffffff1642611b35565b111561181257506001919050565b506000919050565b60035481106118d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f747075747320616674657220746865206c6174657374206f757470757420696e60648201527f6465780000000000000000000000000000000000000000000000000000000000608482015260a4016107c1565b6118da8161178c565b1561198d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f74707574732074686174206861766520616c7265616479206265656e2066696e60648201527f616c697a65640000000000000000000000000000000000000000000000000000608482015260a4016107c1565b600061199860035490565b90508160035581817f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35050565b600060208083528351808285015260005b818110156119fd578581018301518582016040015282016119e1565b81811115611a0f576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008082840360a0811215611a5757600080fd5b8335925060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082011215611a8b57600080fd5b506020830190509250929050565b600060208284031215611aab57600080fd5b5035919050565b60008060008060808587031215611ac857600080fd5b5050823594602084013594506040840135936060013592509050565b60008060408385031215611af757600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611b4757611b47611b06565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611bb357611bb3611b06565b500290565b600060808284031215611bca57600080fd5b6040516080810181811067ffffffffffffffff82111715611c14577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c7657611c76611b06565b5060010190565b60008219821115611c9057611c90611b06565b500190565b600082611ccb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a",
}

// OasysL2OutputOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OasysL2OutputOracleMetaData.ABI instead.
var OasysL2OutputOracleABI = OasysL2OutputOracleMetaData.ABI

// OasysL2OutputOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OasysL2OutputOracleMetaData.Bin instead.
var OasysL2OutputOracleBin = OasysL2OutputOracleMetaData.Bin

// DeployOasysL2OutputOracle deploys a new Ethereum contract, binding an instance of OasysL2OutputOracle to it.
func DeployOasysL2OutputOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _submissionInterval *big.Int, _l2BlockTime *big.Int, _startingBlockNumber *big.Int, _startingTimestamp *big.Int, _proposer common.Address, _challenger common.Address, _finalizationPeriodSeconds *big.Int) (common.Address, *types.Transaction, *OasysL2OutputOracle, error) {
	parsed, err := OasysL2OutputOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OasysL2OutputOracleBin), backend, _submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _proposer, _challenger, _finalizationPeriodSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OasysL2OutputOracle{OasysL2OutputOracleCaller: OasysL2OutputOracleCaller{contract: contract}, OasysL2OutputOracleTransactor: OasysL2OutputOracleTransactor{contract: contract}, OasysL2OutputOracleFilterer: OasysL2OutputOracleFilterer{contract: contract}}, nil
}

// OasysL2OutputOracle is an auto generated Go binding around an Ethereum contract.
type OasysL2OutputOracle struct {
	OasysL2OutputOracleCaller     // Read-only binding to the contract
	OasysL2OutputOracleTransactor // Write-only binding to the contract
	OasysL2OutputOracleFilterer   // Log filterer for contract events
}

// OasysL2OutputOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasysL2OutputOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasysL2OutputOracleSession struct {
	Contract     *OasysL2OutputOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasysL2OutputOracleCallerSession struct {
	Contract *OasysL2OutputOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OasysL2OutputOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasysL2OutputOracleTransactorSession struct {
	Contract     *OasysL2OutputOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasysL2OutputOracleRaw struct {
	Contract *OasysL2OutputOracle // Generic contract binding to access the raw methods on
}

// OasysL2OutputOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleCallerRaw struct {
	Contract *OasysL2OutputOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OasysL2OutputOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleTransactorRaw struct {
	Contract *OasysL2OutputOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasysL2OutputOracle creates a new instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracle(address common.Address, backend bind.ContractBackend) (*OasysL2OutputOracle, error) {
	contract, err := bindOasysL2OutputOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracle{OasysL2OutputOracleCaller: OasysL2OutputOracleCaller{contract: contract}, OasysL2OutputOracleTransactor: OasysL2OutputOracleTransactor{contract: contract}, OasysL2OutputOracleFilterer: OasysL2OutputOracleFilterer{contract: contract}}, nil
}

// NewOasysL2OutputOracleCaller creates a new read-only instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleCaller(address common.Address, caller bind.ContractCaller) (*OasysL2OutputOracleCaller, error) {
	contract, err := bindOasysL2OutputOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleCaller{contract: contract}, nil
}

// NewOasysL2OutputOracleTransactor creates a new write-only instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OasysL2OutputOracleTransactor, error) {
	contract, err := bindOasysL2OutputOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleTransactor{contract: contract}, nil
}

// NewOasysL2OutputOracleFilterer creates a new log filterer instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OasysL2OutputOracleFilterer, error) {
	contract, err := bindOasysL2OutputOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleFilterer{contract: contract}, nil
}

// bindOasysL2OutputOracle binds a generic wrapper to an already deployed contract.
func bindOasysL2OutputOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OasysL2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.contract.Transact(opts, method, params...)
}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) CHALLENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "CHALLENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) CHALLENGER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.CHALLENGER(&_OasysL2OutputOracle.CallOpts)
}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) CHALLENGER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.CHALLENGER(&_OasysL2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_OasysL2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_OasysL2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) L2BLOCKTIME() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BLOCKTIME(&_OasysL2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BLOCKTIME(&_OasysL2OutputOracle.CallOpts)
}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) PROPOSER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "PROPOSER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) PROPOSER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.PROPOSER(&_OasysL2OutputOracle.CallOpts)
}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) PROPOSER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.PROPOSER(&_OasysL2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SUBMISSIONINTERVAL(&_OasysL2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SUBMISSIONINTERVAL(&_OasysL2OutputOracle.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Challenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "challenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Challenger() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Challenger(&_OasysL2OutputOracle.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Challenger() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Challenger(&_OasysL2OutputOracle.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.ComputeL2Timestamp(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.ComputeL2Timestamp(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) FinalizationPeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "finalizationPeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FinalizationPeriodSeconds(&_OasysL2OutputOracle.CallOpts)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FinalizationPeriodSeconds(&_OasysL2OutputOracle.CallOpts)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2Output(&_OasysL2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2Output(&_OasysL2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2OutputAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2OutputAfter", _l2BlockNumber)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2OutputIndexAfter", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputIndexAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputIndexAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) IsOutputFinalized(opts *bind.CallOpts, l2OutputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "isOutputFinalized", l2OutputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) IsOutputFinalized(l2OutputIndex *big.Int) (bool, error) {
	return _OasysL2OutputOracle.Contract.IsOutputFinalized(&_OasysL2OutputOracle.CallOpts, l2OutputIndex)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) IsOutputFinalized(l2OutputIndex *big.Int) (bool, error) {
	return _OasysL2OutputOracle.Contract.IsOutputFinalized(&_OasysL2OutputOracle.CallOpts, l2OutputIndex)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) L2BlockTime() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BlockTime(&_OasysL2OutputOracle.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) L2BlockTime() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BlockTime(&_OasysL2OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) LatestBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) LatestOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextVerifyIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextVerifyIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextVerifyIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextVerifyIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextVerifyIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextVerifyIndex(&_OasysL2OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Proposer() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Proposer(&_OasysL2OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Proposer() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Proposer(&_OasysL2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) StartingBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) StartingTimestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingTimestamp(&_OasysL2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) StartingTimestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingTimestamp(&_OasysL2OutputOracle.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) SubmissionInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "submissionInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SubmissionInterval() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SubmissionInterval(&_OasysL2OutputOracle.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) SubmissionInterval() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SubmissionInterval(&_OasysL2OutputOracle.CallOpts)
}

// VerifiedBlockNumber is a free data retrieval call binding the contract method 0x594098c6.
//
// Solidity: function verifiedBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) VerifiedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "verifiedBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifiedBlockNumber is a free data retrieval call binding the contract method 0x594098c6.
//
// Solidity: function verifiedBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) VerifiedBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.VerifiedBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// VerifiedBlockNumber is a free data retrieval call binding the contract method 0x594098c6.
//
// Solidity: function verifiedBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) VerifiedBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.VerifiedBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Version() (string, error) {
	return _OasysL2OutputOracle.Contract.Version(&_OasysL2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Version() (string, error) {
	return _OasysL2OutputOracle.Contract.Version(&_OasysL2OutputOracle.CallOpts)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) DeleteL2Outputs(opts *bind.TransactOpts, l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "deleteL2Outputs", l2OutputIndex)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) DeleteL2Outputs(l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.DeleteL2Outputs(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) DeleteL2Outputs(l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.DeleteL2Outputs(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex)
}

// FailVerification is a paid mutator transaction binding the contract method 0x99f10c19.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) FailVerification(opts *bind.TransactOpts, l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "failVerification", l2OutputIndex, outputRootProof)
}

// FailVerification is a paid mutator transaction binding the contract method 0x99f10c19.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FailVerification(l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.FailVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, outputRootProof)
}

// FailVerification is a paid mutator transaction binding the contract method 0x99f10c19.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) FailVerification(l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.FailVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, outputRootProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) Initialize(opts *bind.TransactOpts, _startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "initialize", _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.Initialize(&_OasysL2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.Initialize(&_OasysL2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) ProposeL2Output(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "proposeL2Output", _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.ProposeL2Output(&_OasysL2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.ProposeL2Output(&_OasysL2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0x657c10ff.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) SucceedVerification(opts *bind.TransactOpts, l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "succeedVerification", l2OutputIndex, outputRootProof)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0x657c10ff.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SucceedVerification(l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.SucceedVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, outputRootProof)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0x657c10ff.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) outputRootProof) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) SucceedVerification(l2OutputIndex *big.Int, outputRootProof TypesOutputRootProof) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.SucceedVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, outputRootProof)
}

// OasysL2OutputOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleInitializedIterator struct {
	Event *OasysL2OutputOracleInitialized // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleInitialized)
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
		it.Event = new(OasysL2OutputOracleInitialized)
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
func (it *OasysL2OutputOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleInitialized represents a Initialized event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*OasysL2OutputOracleInitializedIterator, error) {

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleInitializedIterator{contract: _OasysL2OutputOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleInitialized)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseInitialized(log types.Log) (*OasysL2OutputOracleInitialized, error) {
	event := new(OasysL2OutputOracleInitialized)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleL2OutputFailedIterator is returned from FilterL2OutputFailed and is used to iterate over the raw logs and unpacked data for L2OutputFailed events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleL2OutputFailedIterator struct {
	Event *OasysL2OutputOracleL2OutputFailed // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleL2OutputFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleL2OutputFailed)
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
		it.Event = new(OasysL2OutputOracleL2OutputFailed)
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
func (it *OasysL2OutputOracleL2OutputFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleL2OutputFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleL2OutputFailed represents a L2OutputFailed event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleL2OutputFailed struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL2OutputFailed is a free log retrieval operation binding the contract event 0xbecbaa918310b9e773752f3b22d0cd05e91d296a3ca04a065d1cc0111fba2bca.
//
// Solidity: event L2OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterL2OutputFailed(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*OasysL2OutputOracleL2OutputFailedIterator, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "L2OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleL2OutputFailedIterator{contract: _OasysL2OutputOracle.contract, event: "L2OutputFailed", logs: logs, sub: sub}, nil
}

// WatchL2OutputFailed is a free log subscription operation binding the contract event 0xbecbaa918310b9e773752f3b22d0cd05e91d296a3ca04a065d1cc0111fba2bca.
//
// Solidity: event L2OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchL2OutputFailed(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleL2OutputFailed, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "L2OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleL2OutputFailed)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "L2OutputFailed", log); err != nil {
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

// ParseL2OutputFailed is a log parse operation binding the contract event 0xbecbaa918310b9e773752f3b22d0cd05e91d296a3ca04a065d1cc0111fba2bca.
//
// Solidity: event L2OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseL2OutputFailed(log types.Log) (*OasysL2OutputOracleL2OutputFailed, error) {
	event := new(OasysL2OutputOracleL2OutputFailed)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "L2OutputFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleL2OutputVerifiedIterator is returned from FilterL2OutputVerified and is used to iterate over the raw logs and unpacked data for L2OutputVerified events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleL2OutputVerifiedIterator struct {
	Event *OasysL2OutputOracleL2OutputVerified // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleL2OutputVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleL2OutputVerified)
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
		it.Event = new(OasysL2OutputOracleL2OutputVerified)
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
func (it *OasysL2OutputOracleL2OutputVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleL2OutputVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleL2OutputVerified represents a L2OutputVerified event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleL2OutputVerified struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL2OutputVerified is a free log retrieval operation binding the contract event 0x43559e36255b2cd130c34b2551834b85ae8fb206cbe183be98ff8c8c44c2250a.
//
// Solidity: event L2OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterL2OutputVerified(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*OasysL2OutputOracleL2OutputVerifiedIterator, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "L2OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleL2OutputVerifiedIterator{contract: _OasysL2OutputOracle.contract, event: "L2OutputVerified", logs: logs, sub: sub}, nil
}

// WatchL2OutputVerified is a free log subscription operation binding the contract event 0x43559e36255b2cd130c34b2551834b85ae8fb206cbe183be98ff8c8c44c2250a.
//
// Solidity: event L2OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchL2OutputVerified(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleL2OutputVerified, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "L2OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleL2OutputVerified)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "L2OutputVerified", log); err != nil {
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

// ParseL2OutputVerified is a log parse operation binding the contract event 0x43559e36255b2cd130c34b2551834b85ae8fb206cbe183be98ff8c8c44c2250a.
//
// Solidity: event L2OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseL2OutputVerified(log types.Log) (*OasysL2OutputOracleL2OutputVerified, error) {
	event := new(OasysL2OutputOracleL2OutputVerified)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "L2OutputVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputProposedIterator struct {
	Event *OasysL2OutputOracleOutputProposed // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputProposed)
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
		it.Event = new(OasysL2OutputOracleOutputProposed)
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
func (it *OasysL2OutputOracleOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputProposed represents a OutputProposed event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*OasysL2OutputOracleOutputProposedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputProposedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputProposed)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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

// ParseOutputProposed is a log parse operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputProposed(log types.Log) (*OasysL2OutputOracleOutputProposed, error) {
	event := new(OasysL2OutputOracleOutputProposed)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputsDeletedIterator is returned from FilterOutputsDeleted and is used to iterate over the raw logs and unpacked data for OutputsDeleted events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputsDeletedIterator struct {
	Event *OasysL2OutputOracleOutputsDeleted // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputsDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputsDeleted)
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
		it.Event = new(OasysL2OutputOracleOutputsDeleted)
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
func (it *OasysL2OutputOracleOutputsDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputsDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputsDeleted represents a OutputsDeleted event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputsDeleted struct {
	PrevNextOutputIndex *big.Int
	NewNextOutputIndex  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOutputsDeleted is a free log retrieval operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputsDeleted(opts *bind.FilterOpts, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (*OasysL2OutputOracleOutputsDeletedIterator, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputsDeletedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputsDeleted", logs: logs, sub: sub}, nil
}

// WatchOutputsDeleted is a free log subscription operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputsDeleted(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputsDeleted, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (event.Subscription, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputsDeleted)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
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

// ParseOutputsDeleted is a log parse operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputsDeleted(log types.Log) (*OasysL2OutputOracleOutputsDeleted, error) {
	event := new(OasysL2OutputOracleOutputsDeleted)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
