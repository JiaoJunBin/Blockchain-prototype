package core

import (
	"errors"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) GetIndex() uint {
	return uint(len(bc.Blocks))
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) GenerateBlock(memPool map[hash]*Transaction) (b *Block, err error) {

	txs, err := preparetxs(memPool)
	if err != nil {
		return nil, err
	}

	root, err := NewMerkleTree(txs)
	if err != nil {
		return nil, err
	}

	newHeader := &BlockHeader{
		Index:          bc.GetIndex(),
		PrevBlockHash:  bc.GetLastBlock().CurBlockHash,
		MerkleRootHash: root.Value,
		TimeStamp:      uint64(time.Now().Nanosecond()),
		NBits:          bc.GetNBits(),
		Nonce:          0, // adjust Nonce in PoW()
	}

	newBlock := PoW(newHeader)

	return newBlock, nil
}

// prepare txs for mining
func preparetxs(memPool map[hash]*Transaction) (txs []*Transaction, err error) {
	if len(memPool) < MIN_TRANSACTIONS_PER_BLOCK {
		err = errors.New("not enough txs in MemPool to generate block")
		return
	}
	txs = make([]*Transaction, 0)
	cnt := 0
	for _, tx := range memPool {
		cnt += 1
		if cnt > MAX_TRANSACTIONS_PER_BLOCK {
			break
		}
		txs = append(txs, tx)
	}
	return
}

func (bc *Blockchain) GetNBits() (nBits uint32) {
	//TODO:DIFFICULTY_ADJUSTMENT
	if bc.GetIndex()%DIFFICULTY_ADJUSTMENT_INTERVAL == 0 {

	}
	return bc.GetLastBlock().Header.NBits
}
