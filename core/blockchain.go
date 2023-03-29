package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"sort"
	"time"

	"github.com/pkg/errors"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) GetIndex() uint {
	return uint(len(bc.Blocks))
}

func (bc *Blockchain) GetLastBlock() (b *Block, err error) {
	return bc.GetPrevBlock(0)
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

// args: coinbaase Transaction
func GenerateGenisesBlock(cbtx *Transaction) (b *Block, err error) {
	root, err := NewMerkleTree([]*Transaction{cbtx})
	if err != nil {
		return nil, err
	}
	newHeader := &BlockHeader{
		Index:          0,
		PrevBlockHash:  [32]byte{},
		MerkleRootHash: cbtx.ID,
		NBits:          GENESIS_NBITS,
		Nonce:          0,
	}
	genesisBlock := &Block{
		BlockHeader:   newHeader,
		CurHeaderHash: sha256.Sum256(Serialize(newHeader)),
		Tx:            []*Transaction{cbtx},
		MerkleRoot:    root,
		TimeStamp:     int64(time.Now().Nanosecond()),
	}
	return genesisBlock, nil
}

func (bc *Blockchain) GenerateBlock(txs []*Transaction) (b *Block, err error) {

	// txs, err := preparetxs(memPool)
	// if err != nil {
	// 	return nil, err
	// }
	root, err := NewMerkleTree(txs)
	if err != nil {
		return nil, err
	}
	lastBlock, err := bc.GetLastBlock()
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
func Preparetxs(memPool map[hash]*Transaction) (txs []*Transaction, err error) {
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

func (bc *Blockchain) ValidateNewBlock(b *Block) (bool, error) {
	lastBlock, err := bc.GetLastBlock()
	if err != nil {
		return false, err
	}
	if b.PrevBlockHash != lastBlock.CurHeaderHash ||
		b.Index != lastBlock.Index+1 ||
		b.TimeStamp <= lastBlock.TimeStamp {
		return false, nil
	}
	return b.IsValid(), nil
}

func (bc *Blockchain) AppendBlock(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *Blockchain) FindUTXO() map[string]TXOutputs {
	UTXO := make(map[string]TXOutputs)
	spentTXOs := make(map[string][]int)
	for _, block := range bc.Blocks {

		for _, tx := range block.Tx {
			txID := hex.EncodeToString(tx.ID[:])

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

// FindTransaction finds a transaction by its ID
func (bc *Blockchain) FindTransaction(ID []byte) (Transaction, error) {

	for _, block := range bc.Blocks {

		for _, tx := range block.Tx {
			if bytes.Compare(tx.ID[:], ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
}

// SignTransaction signs inputs of a Transaction
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID[:])] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}
