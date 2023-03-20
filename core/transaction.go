package core

import "crypto/sha256"

type Transaction struct {
	ID hash // hash of {ID:[32]byte{}, txinput, txoutput}
	TxInput
	Txoutput
}

// TODO:
type TxInput struct {
}

type Txoutput struct {
}

func (tx *Transaction) Hash() hash {
	tx.ID = [32]byte{}
	return sha256.Sum256(Serialize(tx))
}
