package core

import "time"

// block
const (
	MIN_TRANSACTIONS_PER_BLOCK = 2  // txs
	MAX_TRANSACTIONS_PER_BLOCK = 10 // txs
)

// mining
const (
	// compact version of 0x00000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
	GENESIS_NBITS = 0x1d00ffff

	BLOCK_GENERATION_INTERVAL      = 3 * time.Second // seconds
	DIFFICULTY_ADJUSTMENT_INTERVAL = 3               // blocks
)
