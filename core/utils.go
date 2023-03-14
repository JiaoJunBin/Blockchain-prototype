package core

import (
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
