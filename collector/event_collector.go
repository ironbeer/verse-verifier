package collector

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type logProcessor struct {
	abi          *abi.ABI
	name         string
	log2event    func(log types.Log) interface{}
	eventHandler func(w *EventCollector, tx *database.Database, event interface{}) error
}

var (
	filterTopics [][]common.Hash
	processors   []*logProcessor
)

func init() {
	// Events of legacy optimism
	if parsed, err := abi.JSON(strings.NewReader(scc.SccABI)); err != nil {
		panic(err)
	} else {
		processors = append(processors,
			&logProcessor{
				abi:          &parsed,
				name:         "StateBatchAppended",
				log2event:    func(log types.Log) interface{} { return &scc.SccStateBatchAppended{Raw: log} },
				eventHandler: handleStateBatchAppendedEvent,
			},
			&logProcessor{
				abi:          &parsed,
				name:         "StateBatchDeleted",
				log2event:    func(log types.Log) interface{} { return &scc.SccStateBatchDeleted{Raw: log} },
				eventHandler: handleStateBatchDeletedEvent,
			},
			&logProcessor{
				abi:          &parsed,
				name:         "StateBatchVerified",
				log2event:    func(log types.Log) interface{} { return &scc.SccStateBatchVerified{Raw: log} },
				eventHandler: handleStateBatchVerifiedEvent,
			})
	}

	// Events of opstack
	if parsed, err := abi.JSON(strings.NewReader(l2oo.OasysL2OutputOracleABI)); err != nil {
		panic(err)
	} else {
		processors = append(processors,
			&logProcessor{
				abi:          &parsed,
				name:         "OutputProposed",
				log2event:    func(log types.Log) interface{} { return &l2oo.OasysL2OutputOracleOutputProposed{Raw: log} },
				eventHandler: handleOutputProposedEvent,
			},
			&logProcessor{
				abi:          &parsed,
				name:         "OutputsDeleted",
				log2event:    func(log types.Log) interface{} { return &l2oo.OasysL2OutputOracleOutputsDeleted{Raw: log} },
				eventHandler: handleOutputsDeletedEvent,
			},
			&logProcessor{
				abi:          &parsed,
				name:         "OutputVerified",
				log2event:    func(log types.Log) interface{} { return &l2oo.OasysL2OutputOracleOutputVerified{Raw: log} },
				eventHandler: handleOutputVerifiedEvent,
			})
	}

	filterTopics = append(filterTopics, []common.Hash{})
	for _, p := range processors {
		if ae, ok := p.abi.Events[p.name]; !ok {
			panic(fmt.Sprintf("Failed to get event topic(name=%s) from ABI.", p.name))
		} else {
			filterTopics[0] = append(filterTopics[0], ae.ID)
		}
	}
}

// Worker to collect events for OasysStateCommitmentChain.
type EventCollector struct {
	cfg    *config.Verifier
	db     *database.Database
	hub    ethutil.ReadOnlyClient
	signer common.Address
	log    log.Logger
}

func NewEventCollector(
	cfg *config.Verifier,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
	signer common.Address,
) *EventCollector {
	return &EventCollector{
		cfg:    cfg,
		db:     db,
		hub:    hub,
		signer: signer,
		log:    log.New("worker", "event-collector"),
	}
}

func (w *EventCollector) Start(ctx context.Context) {
	w.log.Info("Worker started",
		"interval", w.cfg.Interval, "event-filter-limit", w.cfg.EventFilterLimit)

	ticker := time.NewTicker(w.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-ticker.C:
			w.work(ctx)
		}
	}
}

func (w *EventCollector) work(ctx context.Context) {
	for {
		// get new blocks from database
		blocks, err := w.db.Block.FindUncollecteds(w.cfg.EventFilterLimit)
		if err != nil && !errors.Is(err, database.ErrNotFound) {
			w.log.Error("Failed to find uncollected blocks", "err", err)
			return
		} else if len(blocks) == 0 {
			w.log.Debug("Wait for new block")
			return
		}

		// collect event logs from hub-layer
		start, end := blocks[0].Number, blocks[len(blocks)-1].Number
		filter := ethereum.FilterQuery{
			Topics:    filterTopics,
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
		}
		logs, err := w.hub.FilterLogs(ctx, filter)
		if err != nil {
			w.log.Error("Failed to fetch event logs from hub-layer",
				"start", start, "end", end, "err", err)
			return
		}

		if err = w.db.Transaction(func(tx *database.Database) error {
			if len(logs) == 0 {
				w.log.Debug("No event log", "start", start, "end", end)
			} else {
				for _, log := range logs {
					if err := w.processLog(tx, log); err != nil {
						return err
					}
				}
			}

			return w.saveLogCollectedBlocks(tx, start, end)
		}); err != nil {
			return
		}
	}
}

// Parse event logs and save to database.
func (w *EventCollector) processLog(tx *database.Database, log types.Log) error {
	var processor *logProcessor
	for _, p := range processors {
		if p.abi.Events[p.name].ID == log.Topics[0] {
			processor = p
			break
		}
	}
	if processor == nil {
		return fmt.Errorf("unknown log topic: %s", log.Topics[0])
	}

	event := processor.log2event(log)
	if err := processor.abi.UnpackIntoInterface(event, processor.name, log.Data); err != nil {
		return fmt.Errorf("failed to unpack log data: %w", err)
	}

	var indexed abi.Arguments
	for _, arg := range processor.abi.Events[processor.name].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if err := abi.ParseTopics(event, indexed, log.Topics[1:]); err != nil {
		return fmt.Errorf("failed to parse indexed log data: %w", err)
	}

	return processor.eventHandler(w, tx, event)
}

// Event handlers for StateCommitmentChain of Legacy Optimism
func handleStateBatchAppendedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*scc.SccStateBatchAppended)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	var (
		address    = e.Raw.Address
		batchIndex = e.BatchIndex.Uint64()
		logCtx     = []interface{}{
			"block", e.Raw.BlockNumber,
			"scc", address.Hex(),
			"index", batchIndex,
		}
	)
	w.log.Info("New SCC.StateBatchAppended event", logCtx...)

	// delete the `OptimismState` records in consideration of chain reorganization
	if rows, err := tx.Optimism.DeleteStates(address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized states", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized states", append(logCtx, "rows", rows)...)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	if rows, err := tx.Optimism.DeleteSignatures(w.signer, address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures",
			append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	// save new state
	if _, err := tx.Optimism.SaveState(e); err != nil {
		w.log.Error("Failed to save SCC.StateBatchAppended event", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

func handleStateBatchDeletedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*scc.SccStateBatchDeleted)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	var (
		address    = e.Raw.Address
		batchIndex = e.BatchIndex.Uint64()
		logCtx     = []interface{}{
			"block", e.Raw.BlockNumber,
			"scc", address.Hex(),
			"index", batchIndex,
		}
	)
	w.log.Info("New SCC.StateBatchDeleted event", logCtx...)

	// delete `OptimismState` records after target batchIndex
	if rows, err := tx.Optimism.DeleteStates(address, batchIndex); err != nil {
		w.log.Error("Failed to delete states", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted states", append(logCtx, "rows", rows)...)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	if rows, err := tx.Optimism.DeleteSignatures(w.signer, address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	return nil
}

func handleStateBatchVerifiedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*scc.SccStateBatchVerified)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	nextIndex := e.BatchIndex.Uint64() + 1

	logCtx := []interface{}{
		"block", e.Raw.BlockNumber,
		"scc", e.Raw.Address.Hex(),
		"next_index", nextIndex,
	}
	w.log.Info("New SCC.StateBatchVerified event", logCtx...)

	if err := tx.Optimism.SaveNextIndex(e.Raw.Address, nextIndex); err != nil {
		w.log.Error("Failed to save next index", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

// Event handlers for L2OutputOracle of OPStack
func handleOutputProposedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*l2oo.OasysL2OutputOracleOutputProposed)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	var (
		address     = e.Raw.Address
		outputIndex = e.L2OutputIndex.Uint64()
		logCtx      = []interface{}{
			"block", e.Raw.BlockNumber,
			"l2oo", address.Hex(),
			"index", outputIndex,
		}
	)
	w.log.Info("New L2OO.OutputProposed event", logCtx...)

	// delete the `OPStackProposal` records in consideration of chain reorganization
	if rows, err := tx.OPStack.DeleteProposals(address, outputIndex); err != nil {
		w.log.Error("Failed to delete reorganized proposals", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized proposals", append(logCtx, "rows", rows)...)
	}

	// delete the `OPStackSignature` records in consideration of chain reorganization
	if rows, err := tx.OPStack.DeleteSignatures(w.signer, address, outputIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures",
			append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	// save new state
	if _, err := tx.OPStack.SaveProposal(e); err != nil {
		w.log.Error("Failed to save L2OO.OutputProposed event", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

func handleOutputsDeletedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*l2oo.OasysL2OutputOracleOutputsDeleted)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	var (
		address     = e.Raw.Address
		outputIndex = e.NewNextOutputIndex.Uint64()
		logCtx      = []interface{}{
			"block", e.Raw.BlockNumber,
			"l2oo", address.Hex(),
			"index", outputIndex,
		}
	)
	w.log.Info("New L2OO.OutputsDeleted event", logCtx...)

	// delete `OPStackProposal` records after target outputIndex
	if rows, err := tx.OPStack.DeleteProposals(address, outputIndex); err != nil {
		w.log.Error("Failed to delete states", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted states", append(logCtx, "rows", rows)...)
	}

	// delete the `OPStackSignature` records in consideration of chain reorganization
	if rows, err := tx.OPStack.DeleteSignatures(w.signer, address, outputIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	return nil
}

func handleOutputVerifiedEvent(w *EventCollector, tx *database.Database, event interface{}) error {
	e, ok := event.(*l2oo.OasysL2OutputOracleOutputVerified)
	if !ok {
		return fmt.Errorf("event type mismatch(%v)", event)
	}

	nextVerifyIndex := e.L2OutputIndex.Uint64() + 1

	logCtx := []interface{}{
		"block", e.Raw.BlockNumber,
		"l2oo", e.Raw.Address.Hex(),
		"next-verify-index", nextVerifyIndex,
	}
	w.log.Info("New L2OO.OutputVerified event", logCtx...)

	if err := tx.OPStack.SaveNextVerifyIndex(e.Raw.Address, nextVerifyIndex); err != nil {
		w.log.Error("Failed to save next index", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

func (w *EventCollector) saveLogCollectedBlocks(tx *database.Database, start, end uint64) error {
	// save collected blocks
	for number := start; number <= end; number++ {
		if err := tx.Block.SaveLogCollected(number); err != nil {
			w.log.Error("Failed to save collected block", "number", number, "err", err)
			return err
		}
	}
	return nil
}
