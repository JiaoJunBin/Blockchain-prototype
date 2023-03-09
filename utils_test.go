package main

import (
	// bc "blockchain/prototype"
	"testing"
)

func TestStructToByte(t *testing.T) {
	a := struct {
		v string
	}{"hello"}
	data := StructToByte(a)
	if len(data) <= 0 {
		t.Errorf("data=%v\n", data)
	}
}
