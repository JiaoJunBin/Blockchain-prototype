package main

import (
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