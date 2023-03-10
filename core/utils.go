package core

import (
	"github.com/holiman/uint256"
	"math/big"
	"unsafe"
)

// type toByte interface {
// 	*Transaction | *BlockHeader | *MerkleNode
// }

type mockStruct struct {
	addr uintptr
	len  int
	cap  int
}

// convert a struct to []byte
func StructToByte[T any](anystruct T) (data []byte) {
	len := unsafe.Sizeof(anystruct)
	mockBytes := &mockStruct{
		addr: uintptr(unsafe.Pointer(&anystruct)),
		len:  int(len),
		cap:  int(len),
	}
	data = *(*[]byte)(unsafe.Pointer(mockBytes))
	return
}

// convert uint256.Int to big.Int
func ToBig(z *uint256.Int) *big.Int {
	b := new(big.Int)
	words := [4]big.Word{big.Word(z[0]), big.Word(z[1]), big.Word(z[2]), big.Word(z[3])}
	b.SetBits(words[:])
	return b
}
