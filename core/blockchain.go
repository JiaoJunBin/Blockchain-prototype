package core

import (
	"bytes"
	"errors"
	"sort"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) GetIndex() uint {
	return uint(len(bc.Blocks))
}

// counting backward, GetPrevBlock(1) meaning get the prev 1 block before last block
func (bc *Blockchain) GetPrevBlock(i int) (b *Block, err error) {
	if len(bc.Blocks) == 0 {
		return nil, errors.New("GetPrevBlock(i int) error, No Block in Blockchain")
	}
	if i < 0 || i >= len(bc.Blocks) {
		return nil, errors.New("GetPrevBlock(i int) error, Index out of range")
	}
	return bc.Blocks[len(bc.Blocks)-i-1], nil
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
	lastBlock, err := bc.GetPrevBlock(1)
	if err != nil {
		return nil, err
	}
	nbits, err := bc.AdjustDifficulty()
	if err != nil {
		return nil, err
	}
	newHeader := &BlockHeader{
		Index:          bc.GetIndex(),
		PrevBlockHash:  lastBlock.CurHeaderHash,
		MerkleRootHash: root.Value,
		NBits:          nbits,
		Nonce:          0, // adjust Nonce in PoW()
	}

	nonce, hash := PoW(newHeader)
	newHeader.Nonce = nonce
	newBlock := &Block{
		BlockHeader:   newHeader,
		CurHeaderHash: hash,
		Tx:            txs,
		MerkleRoot:    root,
		TimeStamp:     int64(time.Now().Nanosecond()),
	}
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
	// in lexicographical order
	sort.Slice(txs, func(i, j int) bool {
		com := bytes.Compare(txs[i].ID[:], txs[j].ID[:])
		if com == -1 {
			return true
		} else {
			return false
		}
	})
	return txs, nil
}
