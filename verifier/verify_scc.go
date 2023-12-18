package verifier

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var (
	// See: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L29-L46
	merkleDefaultBytes [][32]byte
	merkleDefaultHexs  = []string{
		"0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
		"0x633dc4d7da7256660a892f8f1604a44b5432649cc8ec5cb3ced4c4e6ac94dd1d",
		"0x890740a8eb06ce9be422cb8da5cdafc2b58c0a5e24036c578de2a433c828ff7d",
		"0x3b8ec09e026fdc305365dfc94e189a81b38c7597b3d941c279f042e8206e0bd8",
		"0xecd50eee38e386bd62be9bedb990706951b65fe053bd9d8a521af753d139e2da",
		"0xdefff6d330bb5403f63b14f33b578274160de3a50df4efecf0e0db73bcdd3da5",
		"0x617bdd11f7c0a11f49db22f629387a12da7596f9d1704d7465177c63d88ec7d7",
		"0x292c23a9aa1d8bea7e2435e555a4a60e379a5a35f3f452bae60121073fb6eead",
		"0xe1cea92ed99acdcb045a6726b2f87107e8a61620a232cf4d7d5b5766b3952e10",
		"0x7ad66c0a68c72cb89e4fb4303841966e4062a76ab97451e3b9fb526a5ceb7f82",
		"0xe026cc5a4aed3c22a58cbd3d2ac754c9352c5436f638042dca99034e83636516",
		"0x3d04cffd8b46a874edf5cfae63077de85f849a660426697b06a829c70dd1409c",
		"0xad676aa337a485e4728a0b240d92b3ef7b3c372d06d189322bfd5f61f1e7203e",
		"0xa2fca4a49658f9fab7aa63289c91b7c7b6c832a6d0e69334ff5b0a3483d09dab",
		"0x4ebfd9cd7bca2505f7bef59cc1c12ecc708fff26ae4af19abe852afe9e20c862",
		"0x2def10d13dd169f550f578bda343d9717a138562e0093b380a1120789d53cf10",
	}
)

type sccInterface interface {
	NextIndex(*bind.CallOpts) (*big.Int, error)
}

func init() {
	merkleDefaultBytes = make([][32]byte, len(merkleDefaultHexs))
	for i, hex := range merkleDefaultHexs {
		merkleDefaultBytes[i] = util.BytesToBytes32(common.FromHex(hex))
	}
}

type sccVerifyWorker struct {
	l2Client ethutil.ReadOnlyClient
	sccAddr  common.Address
	scc      sccInterface
}

func NewSccVerifyWorker(
	l2Client ethutil.ReadOnlyClient,
	sccAddr common.Address,
	scc sccInterface,
) verifyWorker {
	return &sccVerifyWorker{
		l2Client: l2Client,
		sccAddr:  sccAddr,
		scc:      scc,
	}
}

func (w *sccVerifyWorker) id() string {
	return w.sccAddr.Hex()
}

func (w *sccVerifyWorker) rpc() string {
	return w.l2Client.URL()
}

func (w *sccVerifyWorker) work(wc *verifyWorkerContext, ctx context.Context) {
	wc.log = wc.log.New("scc", w.sccAddr.Hex())

	// fetch the next index from hub-layer
	nextIndex, err := w.scc.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		wc.log.Error("Failed to call the SCC.nextIndex method", "err", err)
		return
	}

	// verify the signature that match the nextIndex
	// and delete after signatures if there is a problem.
	// Prevent getting stuck indefinitely in the Verify waiting
	// state due to a bug in the signature creation process.
	w.deleteInvalidSignature(wc, nextIndex.Uint64())

	// run verification tasks until time out
	ctx, cancel := context.WithTimeout(ctx, wc.cfg.StateCollectTimeout)
	defer cancel()

	for i := nextIndex.Uint64(); ; i++ {
		states, err := wc.db.Optimism.FindVerificationWaitingStates(
			wc.signer.Signer, w.sccAddr, i, 1)
		if err != nil {
			wc.log.Error("Failed to find states", "err", err)
			return
		} else if len(states) == 0 {
			wc.log.Debug("Wait for new state")
			return
		}

		state := states[0]
		logCtx := []interface{}{"index", state.BatchIndex}

		wc.log.Info("Start state verification", logCtx...)
		approved, sig, err := w.verifyState(wc, ctx, state)
		if err != nil {
			return
		}

		row, err := wc.db.Optimism.SaveSignature(
			nil, nil,
			wc.signer.Signer, state.OptimismScc.Address,
			state.BatchIndex, state.BatchRoot, state.BatchSize,
			state.PrevTotalElements, state.ExtraData,
			approved, sig)
		if err != nil {
			wc.log.Error("Failed to save signature", append(logCtx, "err", err)...)
			return
		}

		wc.topic.Publish(row)
		wc.log.Info("State verification completed", append(logCtx, "approved", approved)...)
	}
}

func (w *sccVerifyWorker) verifyState(
	wc *verifyWorkerContext,
	ctx context.Context,
	state *database.OptimismState,
) (bool, database.Signature, error) {
	logCtx := []interface{}{"index", state.BatchIndex}

	// collect block headers from verse-layer
	var (
		start   = state.PrevTotalElements + 1
		end     = start + state.BatchSize - 1
		headers []*types.Header
		err     error
	)

	bc, err := w.l2Client.NewBatchHeaderClient()
	if err != nil {
		wc.log.Error("Failed to construct batch client", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}

	bi := ethutil.NewBatchHeaderIterator(bc, start, end, wc.cfg.StateCollectLimit)
	defer bi.Close()

	st := time.Now()
	logCtx = append(logCtx, "start", start, "end", end, "batch-size", state.BatchSize)
	for {
		hs, err := bi.Next(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				wc.log.Warn("Time up", logCtx...)
			} else {
				wc.log.Error("Failed to collect state roots", append(logCtx, "err", err)...)
			}
			return false, database.Signature{}, err
		} else if len(hs) == 0 {
			break
		}
		headers = append(headers, hs...)
	}

	wc.log.Info("Collected state roots", append(logCtx, "elapsed", time.Since(st))...)

	// calc and compare state root
	elements := make([][32]byte, len(headers))
	for i, header := range headers {
		elements[i] = header.Root
	}
	merkleRoot, err := calcMerkleRoot(elements)
	if err != nil {
		wc.log.Error("Failed to calculate merkle root", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}
	approved := bytes.Equal(state.BatchRoot[:], merkleRoot[:])

	// calc and save signature
	msg := NewSccMessage(
		wc.signer.ChainID,
		state.OptimismScc.Address,
		new(big.Int).SetUint64(state.BatchIndex),
		state.BatchRoot,
		approved,
	)
	if sig, err := msg.Signature(wc.signer.SignData); err == nil {
		return approved, sig, nil
	} else {
		wc.log.Error("Failed to calculate signature", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}
}

func (w *sccVerifyWorker) deleteInvalidSignature(wc *verifyWorkerContext, nextIndex uint64) {
	logCtx := []interface{}{"next-index", nextIndex}

	signer := wc.signer.Signer
	sigs, err := wc.db.Optimism.FindSignatures(nil, &signer, &w.sccAddr, &nextIndex, 1, 0)
	if err != nil {
		wc.log.Error("Unable to find signatures", append(logCtx, "err", err)...)
		return
	} else if len(sigs) == 0 {
		wc.log.Debug("No invalid signature", logCtx...)
		return
	}

	msg := NewSccMessage(
		wc.signer.ChainID,
		sigs[0].OptimismScc.Address,
		new(big.Int).SetUint64(sigs[0].BatchIndex),
		sigs[0].BatchRoot,
		sigs[0].Approved)
	if match, err := msg.VerifySigner(sigs[0].Signature[:], signer); err == nil && match {
		wc.log.Debug("No invalid signature", logCtx...)
		return
	} else if err != nil {
		wc.log.Error("Unable to verify signature", append(logCtx, "err", err)...)
	}

	wc.log.Warn("Found invalid signature", append(logCtx, "signature", sigs[0].Signature.Hex())...)

	if rows, err := wc.db.Optimism.DeleteSignatures(signer, w.sccAddr, nextIndex); err != nil {
		wc.log.Error("Unable to delete signatures", append(logCtx, "err", err)...)
	} else {
		wc.log.Warn("Deleted invalid signature", append(logCtx, "delete-rows", rows)...)
	}
}

// Calculates a merkle root for a list of 32-byte leaf hashes.
// see: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L22
func calcMerkleRoot(elements [][32]byte) ([32]byte, error) {
	if len(elements) == 0 {
		return [32]byte{}, errors.New("must provide at least one leaf hash")
	}
	if len(elements) == 1 {
		return elements[0], nil
	}

	rowSize := len(elements)
	depth := 0

	for rowSize > 1 {
		halfRowSize := rowSize / 2
		rowSizeIsOdd := rowSize%2 == 1

		for i := 0; i < halfRowSize; i++ {
			leftSibling := elements[(2 * i)][:]
			rightSibling := elements[(2*i)+1][:]
			elements[i] = util.BytesToBytes32(
				crypto.Keccak256(bytes.Join([][]byte{leftSibling, rightSibling}, []byte(""))),
			)
		}

		if rowSizeIsOdd {
			leftSibling := elements[rowSize-1][:]
			rightSibling := merkleDefaultBytes[depth][:]
			elements[halfRowSize] = util.BytesToBytes32(
				crypto.Keccak256(bytes.Join([][]byte{leftSibling, rightSibling}, []byte(""))),
			)
		}

		rowSize = halfRowSize
		if rowSizeIsOdd {
			rowSize++
		}
		depth++
	}

	return elements[0], nil
}
