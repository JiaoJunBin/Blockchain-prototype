package prototype

import (
	tx "blockchain/transaction"
)

type hash [32]byte

// a block in the blockchain
type Block struct {
	Header *BlockHeader
	tx     []*tx.Transaction
}

type BlockHeader struct {
	index          uint
	PrevBlockHash  hash
	CurBlockHash   hash
	MerkleRootHash hash
	TimeStamp      uint64 // the creation time of block (seconds from Unix Epoch)
	NBits          uint32 // difficulty
	Nonce          uint32
}
