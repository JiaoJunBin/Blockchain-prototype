package core

import (
	"crypto/rand"
	"fmt"
	"log"
)


func NewCoinbaseTx(address string, data string) (t *Transaction) {
	//characteristics of mint coin
	//only have one input
	//Do not need transaction ID
	//Do not need index
	//Do not need sign, free to write data
	if data == "" {
		randData := make([]byte, 20)
		_, err := rand.Read(randData)
		if err != nil {
			log.Panic(err)
		}

		data = fmt.Sprintf("%x", randData)
	}
	//Coinbase Transaction: each block should contain a coinbase transaction at the very first place for transactions to mint 50 coins. 
	//const COINBASE_AMOUNT: number = 50;
	// const reward = 50
	//Input: the input of coinbase transaction is zero. 
	//input := TXInput{[]byte{}, -1, nil, []byte(data)}
	input := TXInput{ []byte{}, -1, nil, []byte(data)}
	//Output: the output consists of an address and the amount of minted coins. 
	//output := NewTXOutput(subsidy, to)
	output := NewTXOutput(subsidy, address)
	tx := Transaction{ [32]byte{}, []TXInput{input}, []TXOutput{*output}}
	//Transaction ID: the transaction ID of coinbase transaction is calculated by taking a hash of the transaction contents. 
	tx.ID = tx.Hash()

	return &tx

}
