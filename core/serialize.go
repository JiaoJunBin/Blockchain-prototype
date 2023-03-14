package core

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Serialize serializes the BlockHeader
func (bh *BlockHeader) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(bh)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a BlockHeader
func DeserializeBlockHeader(d []byte) *BlockHeader {
	var bh BlockHeader

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&bh)
	if err != nil {
		log.Panic(err)
	}

	return &bh
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// Serialize serializes the Transaction
func (b *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a Transaction
func DeserializeTx(d []byte) *Transaction {
	var tx Transaction

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&tx)
	if err != nil {
		log.Panic(err)
	}

	return &tx
}

// Serialize serializes the MerkleNode
func (n *MerkleNode) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(n)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeMerkleNode(d []byte) *MerkleNode {
	var n MerkleNode

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&n)
	if err != nil {
		log.Panic(err)
	}

	return &n
}
