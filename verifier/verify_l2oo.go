package verifier

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type l2ooVerifyWorker struct {
	l2Client ethutil.ReadOnlyClient
	l2ooAddr common.Address
	l2oo     *l2oo.OasysL2OutputOracle
}

func NewL2OOVerifyWorker(
	l2Client ethutil.ReadOnlyClient,
	l2ooAddr common.Address,
	l2oo *l2oo.OasysL2OutputOracle,
) verifyWorker {
	return &l2ooVerifyWorker{
		l2Client: l2Client,
		l2ooAddr: l2ooAddr,
		l2oo:     l2oo,
	}
}

func (w *l2ooVerifyWorker) id() string {
	return w.l2ooAddr.Hex()
}

func (w *l2ooVerifyWorker) rpc() string {
	return w.l2Client.URL()
}

func (w *l2ooVerifyWorker) work(wc *verifyWorkerContext, ctx context.Context) {
	wc.log = wc.log.New("l2oo", w.l2ooAddr.Hex())

	// fetch the next index from hub-layer
	nextVerifyIndex, err := w.l2oo.NextVerifyIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		wc.log.Error("Failed to call L2OO.nextVerifyIndex method", "err", err)
		return
	}

	// verify the signature that match the nextVerifyIndex
	// and delete after signatures if there is a problem.
	// Prevent getting stuck indefinitely in the Verify waiting
	// state due to a bug in the signature creation process.
	w.deleteInvalidSignature(wc, nextVerifyIndex.Uint64())

	// run verification tasks until time out
	ctx, cancel := context.WithTimeout(ctx, wc.cfg.StateCollectTimeout)
	defer cancel()

	for i := nextVerifyIndex.Uint64(); ; i++ {
		proposals, err := wc.db.OPStack.FindVerificationWaitingProposals(
			wc.signerCtx.Signer, w.l2ooAddr, i, 1)
		if err != nil {
			wc.log.Error("Failed to find proposals", "err", err)
			return
		} else if len(proposals) == 0 {
			wc.log.Debug("Wait for new state")
			return
		}

		proposal := proposals[0]
		logCtx := []interface{}{"index", proposal.L2OutputIndex}

		wc.log.Info("Start proposal verification", logCtx...)
		approved, sig, err := w.verifyProposal(wc, ctx, proposal)
		if err != nil {
			return
		}

		row, err := wc.db.OPStack.SaveSignature(
			nil, nil,
			wc.signerCtx.Signer,
			proposal.OpstackL2OutputOracle.Address,
			proposal.L2OutputIndex,
			proposal.OutputRoot,
			proposal.L2BlockNumber,
			proposal.L1Timestamp,
			approved, sig)
		if err != nil {
			wc.log.Error("Failed to save signature", append(logCtx, "err", err)...)
			return
		}

		wc.topic.Publish(row)
		wc.log.Info("State verification completed", append(logCtx, "approved", approved)...)
	}
}

func (w *l2ooVerifyWorker) verifyProposal(
	wc *verifyWorkerContext,
	ctx context.Context,
	proposal *database.OpstackProposal,
) (bool, database.Signature, error) {
	logCtx := []interface{}{
		"l2oo", proposal.OpstackL2OutputOracle.Address.Hex(),
		"index", proposal.L2OutputIndex,
	}

	// verify storage proof of L2ToL1MessagePasser
	output, err := ethutil.GetOpstackOutputV0(ctx, w.l2Client,
		ethutil.OpstackPredeploys.L2ToL1MessagePasser, []string{}, proposal.L2BlockNumber)
	if err != nil {
		return false, database.Signature{}, err
	}

	approved := bytes.Equal(proposal.OutputRoot[:], output.OutputRoot().Bytes())

	// calc and save signature
	msg := NewL2ooMessage(
		wc.signerCtx.ChainID,
		proposal.OpstackL2OutputOracle.Address,
		new(big.Int).SetUint64(proposal.L2OutputIndex),
		proposal.OutputRoot,
		new(big.Int).SetUint64(proposal.L1Timestamp),
		new(big.Int).SetUint64(proposal.L2BlockNumber),
		approved,
	)
	if sig, err := msg.Signature(wc.signerCtx.SignData); err == nil {
		return approved, sig, nil
	} else {
		wc.log.Error("Failed to calculate signature", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}
}

func (w *l2ooVerifyWorker) deleteInvalidSignature(wc *verifyWorkerContext, nextVerifyIndex uint64) {
	logCtx := []interface{}{"next-verify-index", nextVerifyIndex}

	sigs, err := wc.db.OPStack.FindSignatures(nil, &wc.signerCtx.Signer, &w.l2ooAddr, &nextVerifyIndex, 1, 0)
	if err != nil {
		wc.log.Error("Unable to find signatures", append(logCtx, "err", err)...)
		return
	} else if len(sigs) == 0 {
		wc.log.Debug("No invalid signature", logCtx...)
		return
	}

	msg := NewL2ooMessage(
		wc.signerCtx.ChainID,
		sigs[0].OpstackL2OutputOracle.Address,
		new(big.Int).SetUint64(sigs[0].L2OutputIndex),
		sigs[0].OutputRoot,
		new(big.Int).SetUint64(sigs[0].L1Timestamp),
		new(big.Int).SetUint64(sigs[0].L2BlockNumber),
		sigs[0].Approved)
	if match, err := msg.VerifySigner(sigs[0].Signature[:], wc.signerCtx.Signer); err == nil && match {
		wc.log.Debug("No invalid signature", logCtx...)
		return
	} else if err != nil {
		wc.log.Error("Unable to verify signature", append(logCtx, "err", err)...)
	}

	wc.log.Warn("Found invalid signature", append(logCtx, "signature", sigs[0].Signature.Hex())...)

	if rows, err := wc.db.OPStack.DeleteSignatures(wc.signerCtx.Signer, w.l2ooAddr, nextVerifyIndex); err != nil {
		wc.log.Error("Unable to delete signatures", append(logCtx, "err", err)...)
	} else {
		wc.log.Warn("Deleted invalid signature", append(logCtx, "delete-rows", rows)...)
	}
}
