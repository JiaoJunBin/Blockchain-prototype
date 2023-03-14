package core

import (
	"log"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T){
	bh:=&BlockHeader{
		22,hash{},hash{},123,321,
	}
	// log.Printf("bh=%v\n",bh)
	d:=Serialize(bh)
	// log.Printf("d=%v\n",d)

	newbh:=&BlockHeader{}
	Deserialize(d,newbh)
	if !reflect.DeepEqual(bh,newbh) {
		log.Panicf("unequal after Deserialize")
	}
	// log.Printf("newbh=%v\n",newbh)
}