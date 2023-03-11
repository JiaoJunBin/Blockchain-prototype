package core

import ()

type hash [32]byte

// a block in the blockchain
type Block struct {
	Header       *BlockHeader
	CurBlockHash hash
	Tx           []*Transaction // ascending order
	MerkleRoot   *MerkleTree
}

// all field should be used in mining
type BlockHeader struct {
	Index          uint
	PrevBlockHash  hash   // previous block header hash
	MerkleRootHash hash   //
	TimeStamp      uint64 // the creation time of block (seconds from Unix Epoch)
	NBits          uint32 // difficulty
	Nonce          uint32
}
