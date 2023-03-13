package core

import (
	// bc "blockchain/prototype"
	"log"
	"testing"
)

func TestToByte(t *testing.T) {
	// a := struct {
	// 	v string
	// }{"hello"}
	a := 1
	data := ToByte(a)
	log.Printf("a=%d", data)
}
