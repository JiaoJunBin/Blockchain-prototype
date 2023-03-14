package core

import (
	"crypto/sha256"
)

// import "time"

type hash [32]byte

// a block in the blockchain
type Block struct {
	*BlockHeader
	CurHeaderHash hash
	Tx            []*Transaction // ascending order
	MerkleRoot    *MerkleNode    // doesn't contain txs
	TimeStamp     int64          // the creation time of block (seconds from Unix Epoch)
}

// all field should be used in mining
type BlockHeader struct {
	Index          uint
	PrevBlockHash  hash   // previous block header hash
	MerkleRootHash hash   //
	NBits          uint32 // difficulty
	Nonce          uint32
}

func (b *Block) IsValid() bool {
	headerHash := b.BlockHeader.Serialize()
	return b.CurHeaderHash == sha256.Sum256(headerHash)
}

