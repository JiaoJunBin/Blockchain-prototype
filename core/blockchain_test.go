package core

import (
	"log"
	"testing"
)
func TestGetPrevBlock(t *testing.T)  {
	bc:=&Blockchain{}
	b,err:=bc.GetPrevBlock(-1)
	if err != nil {
		log.Printf("err=%+v",err)
	}
	log.Println(b)
}