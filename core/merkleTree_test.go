package core

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"unsafe"
)

func TestBuildTree(t *testing.T) {
	cases := []struct {
		id int
	}{
		{1000000},
		{0100000},
		{0010000},
		{0001000},
		{0000100},
		{0000010},
		{0000001},
	}
	hList := make([]hash, 0)
	for _, c := range cases {
		b := ToByte(c)
		hList = append(hList, sha256.Sum256(b))
	}
	root := buildTree(hList)
	fmt.Printf("root.Value= %x\n", root.Value)
}

// convert a struct to []byte in big endian
func ToByte[T any](anyType T) (data []byte) {
	len := unsafe.Sizeof(anyType)
	mockBytes := &mockStruct{
		addr: uintptr(unsafe.Pointer(&anyType)),
		len:  int(len),
		cap:  int(len),
	}
	data = *(*[]byte)(unsafe.Pointer(mockBytes))
	return
}