// SPDX-License-Identifier: MIT

pragma solidity ^0.8.2;

import { OasysStateCommitmentChain } from "./OasysStateCommitmentChain.sol";
import { OasysL2OutputOracle } from "contracts/OasysL2OutputOracle.sol";

contract OasysStateCommitmentChainVerifier {
    event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot);

    event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot);

    event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot);

    event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot);

    struct ChainBatchHeader {
        uint256 batchIndex;
        bytes32 batchRoot;
        uint256 batchSize;
        uint256 prevTotalElements;
        bytes extraData;
    }

    struct OutputProposal {
        bytes32 outputRoot;
        uint128 timestamp;
        uint128 l2BlockNumber;
    }

    struct StateCommitmentChainAssertLog {
        address stateCommitmentChain;
        ChainBatchHeader batchHeader;
        bytes signatures;
        bool approve;
    }

    struct L2OutputOracleAssertLog {
        address l2OutputOracle;
        uint256 l2OutputIndex;
        OutputProposal l2Output;
        bytes signatures;
        bool approve;
    }

    struct L2OutputOracleSetting {
        bytes32 outputRoot;
        uint128 l2BlockNumber;
    }

    StateCommitmentChainAssertLog[] public sccAssertLogs;
    L2OutputOracleAssertLog[] public l2ooAssertLogs;
    L2OutputOracleSetting public l2ooSetting;

    function approve(
        address stateCommitmentChain,
        ChainBatchHeader memory batchHeader,
        bytes[] calldata signatures
    ) external {
        sccAssertLogs.push(
            StateCommitmentChainAssertLog(stateCommitmentChain, batchHeader, _joinSignatures(signatures), true)
        );

        OasysStateCommitmentChain(stateCommitmentChain).emitStateBatchVerified(
            batchHeader.batchIndex,
            batchHeader.batchRoot
        );

        emit StateBatchApproved(stateCommitmentChain, batchHeader.batchIndex, batchHeader.batchRoot);
    }

    function reject(
        address stateCommitmentChain,
        ChainBatchHeader memory batchHeader,
        bytes[] calldata signatures
    ) external {
        sccAssertLogs.push(
            StateCommitmentChainAssertLog(stateCommitmentChain, batchHeader, _joinSignatures(signatures), false)
        );

        OasysStateCommitmentChain(stateCommitmentChain).emitStateBatchFailed(
            batchHeader.batchIndex,
            batchHeader.batchRoot
        );

        emit StateBatchRejected(stateCommitmentChain, batchHeader.batchIndex, batchHeader.batchRoot);
    }

    function approve(
        address l2OutputOracle,
        uint256 l2OutputIndex,
        OutputProposal calldata l2Output,
        bytes[] calldata signatures
    ) external {
        l2ooAssertLogs.push(
            L2OutputOracleAssertLog(l2OutputOracle, l2OutputIndex, l2Output, _joinSignatures(signatures), true)
        );

        OasysL2OutputOracle(l2OutputOracle).emitOutputVerified(l2OutputIndex, l2ooSetting.outputRoot, l2ooSetting.l2BlockNumber);

        emit L2OutputApproved(l2OutputOracle, l2OutputIndex, l2Output.outputRoot);
    }

    function reject(
        address l2OutputOracle,
        uint256 l2OutputIndex,
        OutputProposal calldata l2Output,
        bytes[] calldata signatures
    ) external {
        l2ooAssertLogs.push(
            L2OutputOracleAssertLog(l2OutputOracle, l2OutputIndex, l2Output, _joinSignatures(signatures), false)
        );

        OasysL2OutputOracle(l2OutputOracle).emitOutputFailed(l2OutputIndex, l2ooSetting.outputRoot, l2ooSetting.l2BlockNumber);

        emit L2OutputRejected(l2OutputOracle, l2OutputIndex, l2Output.outputRoot);
    }

    function setL2ooSetting(L2OutputOracleSetting calldata _l2ooSetting) external {
        l2ooSetting = _l2ooSetting;
    }

    function _joinSignatures(bytes[] calldata signatures) internal pure returns (bytes memory joined) {
        for (uint256 i = 0; i < signatures.length; i++) {
            joined = abi.encodePacked(joined, signatures[i]);
        }
    }
}
