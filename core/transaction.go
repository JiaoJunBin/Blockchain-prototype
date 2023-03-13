package core

import "crypto/sha256"

type Transaction struct {
	ID hash // hash of {ID:nil, txinput, txoutput}
	TxInput
	Txoutput
}

// TODO:
type TxInput struct {
}

type Txoutput struct {
}

func (tx *Transaction) Hash() hash {
	b := ToByte(tx)
	return sha256.Sum256(b)
}
