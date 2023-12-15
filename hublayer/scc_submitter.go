package hublayer

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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/scc"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"golang.org/x/net/context"
)

const (
	maxTxSize = 120 * 1024 // 120KB
	minTxGas  = 24871      // Multicall2 minimum required gas
)

var (
	minStake   = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000))
	mcall2Abi  *abi.ABI
	fakeSCCV   *sccverifier.OasysRollupVerifier
	fakeTxOpts *bind.TransactOpts
)

func init() {
	if abi, err := multicall2.Multicall2MetaData.GetAbi(); err != nil {
		panic(err)
	} else {
		mcall2Abi = abi
	}

	if sccv, err := sccverifier.NewOasysRollupVerifier(common.Address{}, nil); err != nil {
		panic(err)
	} else {
		fakeSCCV = sccv
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
}

type Submitter struct {
	cfg   *config.Submitter
	db    *database.Database
	vs    *ValidatorStakings
	tasks *sync.Map
	log   log.Logger
}

func NewSubmitter(cfg *config.Submitter, db *database.Database, sm StakeManager) *Submitter {
	return &Submitter{
		cfg:   cfg,
		db:    db,
		vs:    NewValidatorStakings(sm),
		tasks: &sync.Map{},
		log:   log.New("worker", "scc-submitter"),
	}
}

func (w *Submitter) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.stakeRefreshLoop(ctx)
	}()

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
		"max-gas", w.cfg.MaxGas,
		"verifier", w.cfg.VerifierAddress,
		"multicall2", w.cfg.Multicall2Address)

	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Submitter) stakeRefreshLoop(ctx context.Context) {
	tick := util.NewTicker(time.Hour, 1)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			if err := w.vs.Refresh(ctx); err != nil {
				w.log.Error("Failed to refresh stakes", "err", err)
			}
		}
	}
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
				id, ok := key.(string)
				if !ok {
					return true
				}

				task, ok := value.(submitTask)
				if !ok {
					return true
				}

				// deduplication
				if _, ok := running.Load(id); ok {
					return true
				}
				running.Store(id, 1)

				if !wg.Has(id) {
					handler := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)

						if task, ok := data.(submitTask); ok {
							w.work(ctx, task)
						}
					}
					wg.AddWorker(ctx, id, handler)
				}

				wg.Enqueue(id, task)
				return true
			})
		}
	}
}

func (w *Submitter) AddTask(task submitTask) {
	w.tasks.LoadOrStore(task.id(), task)
}

func (w *Submitter) work(ctx context.Context, task submitTask) {
	logCtx := []interface{}{}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	it, err := task.newIterator(ctx, w.db)
	if err != nil {
		w.log.Error(err.Error(), append(logCtx, "err", err)...)
		return
	}

	calls, err := w.consumeIterator(ctx, it)
	if err != nil {
		w.log.Error(err.Error(), append(logCtx, "err", err)...)
		return
	} else if len(calls) == 0 {
		w.log.Info("No signature", logCtx...)
		return
	}

	tx, err := w.sendTransaction(ctx, logCtx, task.client(), calls)
	if err != nil {
		return
	}

	w.waitForCconfirmation(ctx, append(logCtx, "tx", tx.Hash().String()), task.client(), tx)
}

func (w *Submitter) consumeIterator(ctx context.Context, it multicallIterator) (calls []multicall2.Multicall2Call, err error) {
	for it.hasNext() {
		sigs, err := it.signatures(ctx)
		if err != nil {
			return nil, err
		} else if sigs.Len() == 0 {
			break
		}

		topStakingSigs, first, err := getTopStakingSignatures(
			sigs, minStake, w.vs.TotalStake(), w.vs.StakeBySigner)
		if err != nil {
			return nil, err
		}

		tx, err := it.buildTransaction(ctx, first, topStakingSigs)
		if err != nil {
			return nil, err
		}

		appended := append(calls, multicall2.Multicall2Call{
			Target:   common.HexToAddress(w.cfg.VerifierAddress),
			CallData: tx.Data(),
		})
		if data, err := mcall2Abi.Pack("tryAggregate", true, appended); err != nil {
			return nil, fmt.Errorf("failed to pack multicall data: %w", err)
		} else if len(data) > maxTxSize {
			break
		}

		calls = appended
		it.next()
	}

	return calls, nil
}

func (w *Submitter) sendTransaction(
	ctx context.Context,
	logCtx []interface{},
	l1Client ethutil.WritableClient,
	calls []multicall2.Multicall2Call,
) (*types.Transaction, error) {
	mcall2, err := multicall2.NewMulticall2(common.HexToAddress(w.cfg.Multicall2Address), l1Client)
	if err != nil {
		w.log.Error("Failed to construct Multicall2 contract", "err", err)
		return nil, err
	}

	// to fit max gas
	opts := l1Client.TransactOpts(ctx)
	opts.NoSend = true
	tx, err := mcall2.TryAggregate(opts, true, calls[:1])
	if err != nil {
		w.log.Error("Failed to estimate gas", append(logCtx, "err", err)...)
		return nil, err
	}

	gasPerCall := int(tx.Gas() - minTxGas)
	end := len(calls)
	for ; end > 1 && end*gasPerCall > w.cfg.MaxGas; end-- {
	}
	calls = calls[:end]

	// estimate gas
	tx, err = mcall2.TryAggregate(opts, true, calls)
	if err != nil {
		w.log.Error("Failed to estimate gas", append(logCtx, "err", err)...)
		return nil, err
	}

	// send
	opts = l1Client.TransactOpts(ctx)
	opts.GasLimit = uint64(float64(tx.Gas()) * w.cfg.GasMultiplier)
	tx, err = mcall2.TryAggregate(opts, true, calls)
	if err != nil {
		w.log.Error("Failed to send transaction", append(logCtx, "err", err)...)
		return nil, err
	}

	w.log.Info(
		"Sent transaction",
		append(
			logCtx,
			"call-size", len(calls),
			"tx", tx.Hash().String(),
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

func (sigs optimismSignatures) Len() int                           { return len(sigs) }
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

func (sigs opstackSignatures) Len() int                           { return len(sigs) }
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

type multicallIterator interface {
	hasNext() bool
	next()
	signatures(ctx context.Context) (signatures, error)
	buildTransaction(ctx context.Context, firstRowIdx int, topStakingSignatures [][]byte) (*types.Transaction, error)
}

type sccMulticallIterator struct {
	db      *database.Database
	sccAddr common.Address
	scc     *scc.Scc
	index   uint64

	done bool
	rows []*database.OptimismSignature
}

func (it *sccMulticallIterator) hasNext() bool {
	return !it.done
}

func (it *sccMulticallIterator) next() {
	it.index++
}

func (it *sccMulticallIterator) signatures(ctx context.Context) (signatures, error) {
	rows, err := it.db.Optimism.FindSignatures(nil, nil, &it.sccAddr, &it.index, 1000, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find signatures(index=%d): %w", it.index, err)
	}

	it.done = len(rows) == 0
	it.rows = rows
	return optimismSignatures(it.rows), nil
}

func (it *sccMulticallIterator) buildTransaction(
	ctx context.Context,
	firstRowIdx int,
	topStakingSignatures [][]byte,
) (*types.Transaction, error) {
	// fetch the StateBatchAppended that matches the target batch index
	ev, err := findStateBatchAppendedEvent(ctx, it.scc, it.index)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the StateBatchAppended event(index=%d): %w", it.index, err)
	}
	batchHeader := sccverifier.LibOVMCodecChainBatchHeader{
		BatchIndex:        ev.BatchIndex,
		BatchRoot:         ev.BatchRoot,
		BatchSize:         ev.BatchSize,
		PrevTotalElements: ev.PrevTotalElements,
		ExtraData:         ev.ExtraData,
	}

	method := fakeSCCV.Approve0
	if !it.rows[firstRowIdx].Approved {
		method = fakeSCCV.Reject0
		it.done = true // if rejected, do not approve subsequent batch indexes
	}

	rawTx, err := method(fakeTxOpts, it.sccAddr, batchHeader, topStakingSignatures)
	if err != nil {
		return nil, fmt.Errorf("failed to construct transaction(index=%d): %w", it.index, err)
	}

	return rawTx, nil
}

type l2ooMulticallIterator struct {
	db       *database.Database
	l2ooAddr common.Address
	l2oo     *l2oo.OasysL2OutputOracle
	index    uint64

	done bool
	rows []*database.OpstackSignature
}

func (it *l2ooMulticallIterator) hasNext() bool {
	return !it.done
}

func (it *l2ooMulticallIterator) next() {
	it.index++
}

func (it *l2ooMulticallIterator) signatures(ctx context.Context) (signatures, error) {
	rows, err := it.db.OPStack.FindSignatures(nil, nil, &it.l2ooAddr, &it.index, 1000, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find signatures(index=%d): %w", it.index, err)
	}

	it.done = len(rows) == 0
	it.rows = rows
	return opstackSignatures(it.rows), nil
}

func (it *l2ooMulticallIterator) buildTransaction(
	ctx context.Context,
	firstRowIdx int,
	topStakingSignatures [][]byte,
) (*types.Transaction, error) {
	// fetch the OutputProposed that matches the target output index
	ev, err := findOutputProposed(ctx, it.l2oo, it.index)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the OutputProposed event(index=%d): %w", it.index, err)
	}
	l2Output := sccverifier.TypesOutputProposal{
		OutputRoot:    ev.OutputRoot,
		Timestamp:     ev.L1Timestamp,
		L2BlockNumber: ev.L2BlockNumber,
	}

	method := fakeSCCV.Approve
	if !it.rows[firstRowIdx].Approved {
		method = fakeSCCV.Reject
		it.done = true // if rejected, do not approve subsequent batch indexes
	}

	tx, err := method(fakeTxOpts, it.l2ooAddr, new(big.Int).SetUint64(it.index), l2Output, topStakingSignatures)
	if err != nil {
		return nil, fmt.Errorf("failed to construct transaction(index=%d): %w", it.index, err)
	}

	return tx, nil
}

type submitTask interface {
	id() string
	client() ethutil.WritableClient
	newIterator(ctx context.Context, db *database.Database) (multicallIterator, error)
}

type sccSubmitTask struct {
	l1Client ethutil.WritableClient
	sccAddr  common.Address
	scc      *scc.Scc
}

func (f *sccSubmitTask) id() string {
	return f.sccAddr.Hex()
}

func (f *sccSubmitTask) client() ethutil.WritableClient {
	return f.l1Client
}

func (f *sccSubmitTask) newIterator(ctx context.Context, db *database.Database) (multicallIterator, error) {
	// fetch the next index from hub-layer
	nextIndex, err := f.scc.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to call the SCC.nextIndex method: %w", err)
	}

	return &sccMulticallIterator{
		db:      db,
		sccAddr: f.sccAddr,
		scc:     f.scc,
		index:   nextIndex.Uint64(),
	}, nil
}

type l2ooSubmitTask struct {
	l1Client ethutil.WritableClient
	l2ooAddr common.Address
	l2oo     *l2oo.OasysL2OutputOracle
}

func (f *l2ooSubmitTask) id() string {
	return f.l2ooAddr.Hex()
}

func (f *l2ooSubmitTask) client() ethutil.WritableClient {
	return f.l1Client
}

func (f *l2ooSubmitTask) newIterator(ctx context.Context, db *database.Database) (multicallIterator, error) {
	// fetch the next index from hub-layer
	nextIndex, err := f.l2oo.NextVerifyIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to call the SCC.nextIndex method: %w", err)
	}

	return &l2ooMulticallIterator{
		db:       db,
		l2ooAddr: f.l2ooAddr,
		l2oo:     f.l2oo,
		index:    nextIndex.Uint64(),
	}, nil
}

func NewSccSubmitTask(
	l1Client ethutil.WritableClient,
	sccAddr common.Address,
	scc *scc.Scc,
) submitTask {
	return &sccSubmitTask{
		l1Client: l1Client,
		sccAddr:  sccAddr,
		scc:      scc,
	}
}

func NewL2OOSubmitTask(
	l1Client ethutil.WritableClient,
	l2ooAddr common.Address,
	l2oo *l2oo.OasysL2OutputOracle,
) submitTask {
	return &l2ooSubmitTask{
		l1Client: l1Client,
		l2ooAddr: l2ooAddr,
		l2oo:     l2oo,
	}
}
