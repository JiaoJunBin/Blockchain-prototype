package core

import (
	"log"
	"testing"

	"github.com/holiman/uint256"
)

func TestCompactToBig(t *testing.T) {
	num := 0x181bc330
	rst := CompactToBig(uint32(num))
	log.Printf("rst= %d", rst)
}
func TestBigToCompact(t *testing.T) {
	z, err := uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF") // 64-8 bits
	if err != nil {
		log.Fatalf("err=%v\n", err)
	}

	b := ToBig(z)
	// b:=uint256.ToBig(z)

	rst := BigToCompact(b)
	log.Printf("rst= %x\n", rst)
}
