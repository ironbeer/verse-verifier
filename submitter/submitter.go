package submitter

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/contract/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"golang.org/x/net/context"
)

const (
	maxDataSize = 120 * 1024 // 120KB
)

var (
	minStake   = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000))
	fakeTxOpts *bind.TransactOpts
)

func init() {
	fakeTxOpts = &bind.TransactOpts{
		NoSend:   true,
		Nonce:    common.Big1, // prevent `eth_getNonce`
		GasPrice: common.Big1, // prevent `eth_gasPrice`
		GasLimit: 21_000,      // prevent `eth_estimateGas`
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return t, nil
		},
	}
}

type Submitter struct {
	cfg          *config.Submitter
	db           *database.Database
	stakemanager *stakemanager.Cache
	tasks        *sync.Map
	log          log.Logger
}

func NewSubmitter(
	cfg *config.Submitter,
	db *database.Database,
	sm stakemanager.IStakeManager,
) *Submitter {
	return &Submitter{
		cfg:          cfg,
		db:           db,
		stakemanager: stakemanager.NewCache(sm),
		tasks:        &sync.Map{},
		log:          log.New("worker", "submitter"),
	}
}

func (w *Submitter) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.workLoop(ctx)
	}()

	w.log.Info("Worker started",
		"interval", w.cfg.Interval,
		"concurrency", w.cfg.Concurrency,
		"confirmations", w.cfg.Confirmations,
		"gas-multiplier", w.cfg.GasMultiplier,
		"batch-size", w.cfg.BatchSize,
		"verifier", w.cfg.VerifierAddress,
		"multicall2", w.cfg.Multicall2Address)

	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Submitter) workLoop(ctx context.Context) {
	wg := util.NewWorkerGroup(w.cfg.Concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			w.tasks.Range(func(key, value any) bool {
				l2ChainId, ok := key.(string)
				if !ok {
					return true
				}

				task, ok := value.(*Task)
				if !ok {
					return true
				}

				// deduplication
				if _, ok := running.Load(l2ChainId); ok {
					return true
				}
				running.Store(l2ChainId, 1)

				if !wg.Has(l2ChainId) {
					handler := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)

						if task, ok := data.(*Task); ok {
							w.work(ctx, task)
						}
					}
					wg.AddWorker(ctx, l2ChainId, handler)
				}

				wg.Enqueue(l2ChainId, task)
				return true
			})
		}
	}
}

// TODO
func (w *Submitter) HasTask(l2ChainID uint64, target common.Address) bool {
	_, ok := w.tasks.Load(strconv.FormatUint(l2ChainID, 10))
	return ok
}

func (w *Submitter) AddTask(l2ChainID uint64, task *Task) {
	w.tasks.LoadOrStore(strconv.FormatUint(l2ChainID, 10), task)
}

func (w *Submitter) work(ctx context.Context, task *Task) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	iter, logCtx, err := task.newIterator(ctx)
	if err != nil {
		w.log.Error(err.Error())
		return
	}

	var tx *types.Transaction
	if w.cfg.UseMulticall {
		tx, err = w.sendMulticallTx(ctx, logCtx, iter)
	} else {
		tx, err = w.sendNormalTx(ctx, logCtx, iter)
	}

	if err != nil {
		w.log.Error(err.Error(), logCtx...)
	} else if tx != nil {
		w.waitForCconfirmation(ctx, logCtx, task.l1Client, tx)
	}
}

func (w *Submitter) sendNormalTx(
	ctx context.Context,
	logCtx []interface{},
	iter iterator,
) (*types.Transaction, error) {
	// call eth_estimateGas
	opts := iter.task().l1Client.TransactOpts(ctx)
	opts.NoSend = true
	tx, _, err := iter.next(w, opts)
	if err != nil {
		return nil, err
	} else if tx == nil {
		w.log.Debug("No signatures", logCtx...)
		return nil, nil
	}

	// send transaction
	opts.NoSend = false
	opts.GasLimit = w.cfg.MultiplyGas(tx.Gas())
	if err := iter.task().l1Client.SendTransaction(ctx, tx); err != nil {
		return nil, err
	}

	w.log.Info(
		"Sent transaction",
		append(
			logCtx,
			"tx", tx.Hash().Hex(),
			"nonce", tx.Nonce(),
			"gas-limit", tx.Gas(),
			"gas-fee", tx.GasFeeCap(),
			"gas-tip", tx.GasTipCap(),
		)...)
	return tx, nil
}

func (w *Submitter) sendMulticallTx(
	ctx context.Context,
	logCtx []interface{},
	iter iterator,
) (*types.Transaction, error) {
	var calls []multicall2.Multicall2Call
	for i := 0; i < w.cfg.BatchSize; i++ {
		// build transaction (without sending it).
		rawTx, hasNext, err := iter.next(w, fakeTxOpts)
		if err != nil {
			return nil, err
		} else if rawTx == nil {
			break
		}

		call := multicall2.Multicall2Call{
			Target:   common.HexToAddress(w.cfg.VerifierAddress),
			CallData: rawTx.Data(),
		}

		rawTx, err = iter.task().multicall.TryAggregate(fakeTxOpts, true, append(calls, call))
		if err != nil {
			return nil, err
		} else if len(rawTx.Data()) > maxDataSize {
			w.log.Warn("Oversized", "size", len(rawTx.Data()), "len", i+1)
			break
		}

		calls = append(calls, call)
		if !hasNext {
			break
		}
	}
	if len(calls) == 0 {
		w.log.Info("No calldata", logCtx...)
		return nil, nil
	}

	// call eth_estimateGas
	opts := iter.task().l1Client.TransactOpts(ctx)
	opts.NoSend = true
	tx, err := iter.task().multicall.TryAggregate(opts, true, calls)
	if err != nil {
		return nil, err
	}

	// send transaction
	opts.NoSend = false
	opts.GasLimit = w.cfg.MultiplyGas(tx.Gas())
	tx, err = iter.task().multicall.TryAggregate(opts, true, calls)
	if err != nil {
		return nil, err
	}

	w.log.Info(
		"Sent transaction",
		append(
			logCtx,
			"call-size", len(calls),
			"tx", tx.Hash().Hex(),
			"nonce", tx.Nonce(),
			"gas-limit", tx.Gas(),
			"gas-fee", tx.GasFeeCap(),
			"gas-tip", tx.GasTipCap(),
		)...)
	return tx, nil
}

func (w *Submitter) waitForCconfirmation(
	ctx context.Context,
	logCtx []interface{},
	l1Client ethutil.WritableClient,
	tx *types.Transaction,
) {
	// wait for block to be validated
	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	if err != nil {
		w.log.Error("Failed to receive receipt", append(logCtx, "err", err)...)
		return
	}
	if receipt.Status != 1 {
		w.log.Error("Transaction reverted", logCtx...)
		return
	}

	// wait for confirmations
	confirmed := map[common.Hash]bool{receipt.BlockHash: true}
	for {
		remaining := w.cfg.Confirmations - len(confirmed)
		if remaining <= 0 {
			w.log.Info("Transaction succeeded", logCtx...)
			return
		}

		w.log.Info("Wait for confirmation", append(logCtx, "remaining", remaining)...)
		time.Sleep(time.Second)

		h, err := l1Client.HeaderByNumber(ctx, nil)
		if err != nil {
			w.log.Error("Failed to fetch block header", append(logCtx, "err", err)...)
			continue
		}
		confirmed[h.Hash()] = true
	}
}

func findStateBatchAppendedEvent(
	ctx context.Context,
	scc *scc.Scc,
	batchIndex uint64,
) (appended *scc.SccStateBatchAppended, err error) {
	opts := &bind.FilterOpts{Context: ctx}

	iter, err := scc.FilterStateBatchAppended(opts, []*big.Int{new(big.Int).SetUint64(batchIndex)})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	for {
		if iter.Next() {
			appended = iter.Event // returns the last event
		} else if err := iter.Error(); err != nil {
			return nil, err
		} else {
			break
		}
	}

	if appended == nil {
		err = errors.New("not found")
	}
	return appended, err
}

func findOutputProposed(
	ctx context.Context,
	l2oo *l2oo.OasysL2OutputOracle,
	l2OutputIndex uint64,
) (proposed *l2oo.OasysL2OutputOracleOutputProposed, err error) {
	opts := &bind.FilterOpts{Context: ctx}

	iter, err := l2oo.FilterOutputProposed(opts, nil, []*big.Int{new(big.Int).SetUint64(l2OutputIndex)}, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	for {
		if iter.Next() {
			proposed = iter.Event // returns the last event
		} else if err := iter.Error(); err != nil {
			return nil, err
		} else {
			break
		}
	}

	if proposed == nil {
		err = errors.New("not found")
	}
	return proposed, err
}

type signatures interface {
	Len() int
	Get(i int) interface{}
	Key(i int) string
	Signer(i int) common.Address
	Signature(i int) database.Signature
}

type optimismSignatures []*database.OptimismSignature

func (sigs optimismSignatures) Len() int { return len(sigs) }
func (sigs optimismSignatures) Less(i, j int) bool {
	return bytes.Compare(sigs.Signer(i).Bytes(), sigs.Signer(j).Bytes()) == -1
}
func (sigs optimismSignatures) Swap(i, j int)                      { sigs[i], sigs[j] = sigs[j], sigs[i] }
func (sigs optimismSignatures) Get(i int) interface{}              { return sigs[i] }
func (sigs optimismSignatures) Signer(i int) common.Address        { return sigs[i].Signer.Address }
func (sigs optimismSignatures) Signature(i int) database.Signature { return sigs[i].Signature }
func (sigs optimismSignatures) Key(i int) string {
	row := sigs[i]
	return strings.Join([]string{
		row.BatchRoot.Hex(),
		fmt.Sprintf("%t", row.Approved),
	}, ":")
}

type opstackSignatures []*database.OpstackSignature

func (sigs opstackSignatures) Len() int { return len(sigs) }
func (sigs opstackSignatures) Less(i, j int) bool {
	return bytes.Compare(sigs.Signer(i).Bytes(), sigs.Signer(j).Bytes()) == -1
}
func (sigs opstackSignatures) Swap(i, j int)                      { sigs[i], sigs[j] = sigs[j], sigs[i] }
func (sigs opstackSignatures) Get(i int) interface{}              { return sigs[i] }
func (sigs opstackSignatures) Signer(i int) common.Address        { return sigs[i].Signer.Address }
func (sigs opstackSignatures) Signature(i int) database.Signature { return sigs[i].Signature }
func (sigs opstackSignatures) Key(i int) string {
	row := sigs[i]
	return strings.Join([]string{
		row.OutputRoot.Hex(),
		strconv.FormatUint(row.L2BlockNumber, 10),
		strconv.FormatUint(row.L1Timestamp, 10),
		fmt.Sprintf("%v", row.Approved),
	}, ":")
}

func getTopStakingSignatures(
	rows signatures,
	minStake, totalStake *big.Int,
	stakeBySigner func(common.Address) *big.Int,
) (signatures [][]byte, firstRowIdx int, err error) {
	type group struct {
		stake   *big.Int
		indexes []int
	}
	groups := map[string]*group{}

	// group total staked amounts and indices by key.
	for rowIdx := 0; rowIdx < rows.Len(); rowIdx++ {
		stake := stakeBySigner(rows.Signer(rowIdx))
		if stake.Cmp(minStake) == -1 {
			continue
		}

		key := rows.Key(rowIdx)
		if _, ok := groups[key]; !ok {
			groups[key] = &group{stake: new(big.Int)}
		}
		groups[key].indexes = append(groups[key].indexes, rowIdx)
		groups[key].stake = new(big.Int).Add(groups[key].stake, stake)
	}
	if len(groups) == 0 {
		return nil, 0, nil
	}

	var highest *group
	for k := range groups {
		if highest == nil || groups[k].stake.Cmp(highest.stake) == 1 {
			highest = groups[k]
		}
	}

	// check over half
	required := new(big.Int).Mul(new(big.Int).Div(totalStake, big.NewInt(100)), big.NewInt(51))
	if highest.stake.Cmp(required) == -1 {
		return nil, 0, fmt.Errorf("stake amount shortage(required=%s actual=%s)",
			fromWei(required), fromWei(highest.stake))
	}

	// sort by signer address
	sort.Slice(highest.indexes, func(i, j int) bool {
		a, b := rows.Signer(highest.indexes[i]), rows.Signer(highest.indexes[j])
		return bytes.Compare(a.Bytes(), b.Bytes()) == -1
	})

	signatures = make([][]byte, len(highest.indexes))
	for i, rowIdx := range highest.indexes {
		signatures[i] = rows.Signature(rowIdx).Bytes()
	}

	return signatures, highest.indexes[0], nil
}

func fromWei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}

type iterator interface {
	task() *Task
	next(s *Submitter, opts *bind.TransactOpts) (tx *types.Transaction, hasNext bool, err error)
}

type sccIterator struct {
	t     *Task
	scc   *scc.Scc
	index uint64
	done  bool
}

func (it *sccIterator) task() *Task {
	return it.t
}

func (it *sccIterator) next(s *Submitter, opts *bind.TransactOpts) (tx *types.Transaction, hasNext bool, err error) {
	if it.done {
		return nil, false, nil
	}

	rows, err := s.db.Optimism.FindSignatures(nil, nil, &it.t.target, &it.index, 1000, 0)
	if err != nil {
		return nil, false, fmt.Errorf("failed to find signatures(index=%d): %w", it.index, err)
	}

	signatures, first, err := getTopStakingSignatures(
		optimismSignatures(rows), minStake, s.stakemanager.TotalStake(), s.stakemanager.StakeBySigner)
	if err != nil {
		return nil, false, err
	} else if len(signatures) == 0 {
		it.done = true
		return nil, false, nil
	}

	// fetch the StateBatchAppended that matches the target batch index
	ev, err := findStateBatchAppendedEvent(opts.Context, it.scc, it.index)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch the StateBatchAppended event(index=%d): %w", it.index, err)
	}

	method := it.t.verifier.Approve0
	if !rows[first].Approved {
		method = it.t.verifier.Reject0
		it.done = true // if rejected, do not approve subsequent batch indexes
	}

	tx, err = method(
		opts,
		it.t.target,
		sccverifier.LibOVMCodecChainBatchHeader{
			BatchIndex:        ev.BatchIndex,
			BatchRoot:         ev.BatchRoot,
			BatchSize:         ev.BatchSize,
			PrevTotalElements: ev.PrevTotalElements,
			ExtraData:         ev.ExtraData,
		},
		signatures)

	it.index++
	return tx, !it.done, err
}

type l2ooIterator struct {
	t     *Task
	l2oo  *l2oo.OasysL2OutputOracle
	index uint64
	done  bool
}

func (it *l2ooIterator) task() *Task {
	return it.t
}

func (it *l2ooIterator) next(s *Submitter, opts *bind.TransactOpts) (tx *types.Transaction, hasNext bool, err error) {
	if it.done {
		return nil, false, nil
	}

	rows, err := s.db.OPStack.FindSignatures(nil, nil, &it.t.target, &it.index, 1000, 0)
	if err != nil {
		return nil, false, fmt.Errorf("failed to find signatures(index=%d): %w", it.index, err)
	}

	signatures, first, err := getTopStakingSignatures(
		opstackSignatures(rows), minStake, s.stakemanager.TotalStake(), s.stakemanager.StakeBySigner)
	if err != nil {
		return nil, false, err
	} else if len(signatures) == 0 {
		it.done = true
		return nil, false, nil
	}

	// fetch the OutputProposed that matches the target output index
	ev, err := findOutputProposed(opts.Context, it.l2oo, it.index)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch the OutputProposed event(index=%d): %w", it.index, err)
	}

	method := it.t.verifier.Approve
	if !rows[first].Approved {
		method = it.t.verifier.Reject
		it.done = true // if rejected, do not approve subsequent output indexes
	}

	tx, err = method(
		opts,
		it.t.target,
		new(big.Int).SetUint64(it.index),
		sccverifier.TypesOutputProposal{
			OutputRoot:    ev.OutputRoot,
			Timestamp:     ev.L1Timestamp,
			L2BlockNumber: ev.L2BlockNumber,
		},
		signatures)

	it.index++
	return tx, !it.done, err
}

type Task struct {
	l1Client    ethutil.WritableClient
	target      common.Address
	verifier    *sccverifier.OasysRollupVerifier
	multicall   *multicall2.Multicall2
	newIterator func(ctx context.Context) (it iterator, logCtx []interface{}, err error)
}

func NewTask(
	l1Client ethutil.WritableClient,
	target common.Address,
	contract interface{},
	verifier *sccverifier.OasysRollupVerifier,
	multicall *multicall2.Multicall2,
) (*Task, error) {
	task := &Task{
		l1Client:  l1Client,
		target:    target,
		verifier:  verifier,
		multicall: multicall,
	}

	switch t := contract.(type) {
	case *scc.Scc:
		task.newIterator = func(ctx context.Context) (it iterator, logCtx []interface{}, err error) {
			// fetch the next index from hub-layer
			nextIndex, err := t.NextIndex(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to call the SCC.nextIndex method: %w", err)
			}

			return &sccIterator{
					t:     task,
					scc:   t,
					index: nextIndex.Uint64(),
				}, []interface{}{
					"scc", target,
					"from-index", nextIndex.Uint64(),
				}, nil
		}
	case *l2oo.OasysL2OutputOracle:
		task.newIterator = func(ctx context.Context) (it iterator, logCtx []interface{}, err error) {
			// fetch the next index from hub-layer
			nextIndex, err := t.NextVerifyIndex(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to call the L2OO.nextVerifyIndex method: %w", err)
			}

			return &l2ooIterator{
					t:     task,
					l2oo:  t,
					index: nextIndex.Uint64(),
				}, []interface{}{
					"l2oo", target,
					"from-index", nextIndex.Uint64(),
				}, nil
		}
	default:
		return nil, fmt.Errorf("unknown contract")
	}

	return task, nil
}
