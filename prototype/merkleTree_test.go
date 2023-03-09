package prototype

import (
	"blockchain/utils"
	"crypto/sha256"
	"fmt"
	"testing"
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
		b := utils.StructToByte(c)
		hList = append(hList, sha256.Sum256(b))
	}
	root := buildTree(hList)
	fmt.Printf("root.Value= %v\n", root.Value)
}
