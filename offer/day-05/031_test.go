package day_05

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	var obj LRUCache
	var capacity = 4
	var key, val int
	obj = MConstructor(capacity)
	key = 1
	val = 1
	val = obj.Get(key)
	fmt.Println("val :", val)
	val = 2
	obj.Put(key, val)
}
