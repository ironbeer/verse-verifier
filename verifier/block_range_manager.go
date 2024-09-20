package verifier

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

var ErrStartBlockIsTooLarge = errors.New("start block is too large")

type EventFetchingBlockRangeManager struct {
	l1Signer         ethutil.SignableClient
	maxRange         uint64
	startBlockOffset uint64

	lastStart                uint64
	nextStart                uint64
	startTooLargeCheckPassed bool
}

func NewEventFetchingBlockRangeManager(l1Signer ethutil.SignableClient, maxRange, startBlockOffset uint64) *EventFetchingBlockRangeManager {
	return &EventFetchingBlockRangeManager{
		l1Signer:         l1Signer,
		maxRange:         maxRange,
		startBlockOffset: startBlockOffset,
	}
}

// Check if the start block number is too large.
// If the rollup index of the first event is greater than the next index, the start block number is too large.
// Move back the start block number to ensure the nex index is correctly verified.
func (m *EventFetchingBlockRangeManager) CheckIfStartTooLarge(nextRollupIndex *big.Int, firstEventRollupIndex uint64) error {
	if m.startTooLargeCheckPassed {
		return nil
	}
	// OK if the first event rollup index is less than the next index.
	if firstEventRollupIndex <= nextRollupIndex.Uint64() {
		m.startTooLargeCheckPassed = true
		return nil
	}
	// NG if the first event rollup index is greater than the next index.
	nextStart := m.nextStart - m.maxRange*4
	if m.nextStart < m.maxRange*4 {
		nextStart = 0
	}
	m.nextStart = nextStart
	return fmt.Errorf("the first event rollup index(%d) is greater than the next index(%d), %w", firstEventRollupIndex, nextRollupIndex, ErrStartBlockIsTooLarge)
}

func (m *EventFetchingBlockRangeManager) GetBlockRange(ctx context.Context) (start, end uint64, skipFetchlog bool, err error) {
	// Basically, the end block number is the latest block number.
	if end, err = m.l1Signer.BlockNumber(ctx); err != nil {
		err = fmt.Errorf("failed to fetch the latest block number: %w", err)
		return
	}

	if m.nextStart == 0 {
		// The first time this function is called
		offset := m.startBlockOffset
		if end < offset {
			start = 0
		} else {
			start = end - offset
		}
	} else {
		// The more than once this function is called
		start = m.nextStart
		if start > end {
			// Block number is not updated yet.
			skipFetchlog = true
		}
	}
	// If the range is too wide, divide it into smaller blocks.
	if start < end && m.maxRange < end-start {
		end = start + m.maxRange
	}

	// Update the last start block number.
	m.lastStart = start

	// Next start block is the current end block + 1
	m.nextStart = end + 1

	return
}

func (m *EventFetchingBlockRangeManager) RestoreNextStart() {
	m.nextStart = m.lastStart
}
