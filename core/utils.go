package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/big"

	"github.com/holiman/uint256"
)

// type toByte interface {
// 	*Transaction | *BlockHeader | *MerkleNode
// }

type mockStruct struct {
	addr uintptr
	len  int
	cap  int
}

// convert uint256.Int to big.Int
func ToBig(z *uint256.Int) *big.Int {
	b := new(big.Int)
	words := [4]big.Word{big.Word(z[0]), big.Word(z[1]), big.Word(z[2]), big.Word(z[3])}
	b.SetBits(words[:])
	return b
}

// Serialize serializes the BlockHeader
func Serialize[T any](anyType T) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(anyType)
	if err != nil {
		log.Panicln(err)
	}
	return result.Bytes()
}

// DeserializeBlock deserializes a BlockHeader
func Deserialize[T any](d []byte, anyType *T) {

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(anyType)
	if err != nil {
		log.Panicln(err)
	}
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
