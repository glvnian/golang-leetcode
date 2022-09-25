package day_10

import (
	"fmt"
	"testing"
)

func TestMapSum(t *testing.T) {
	var key, prefix string
	var val, r int
	obj := SConstructor()
	key = "a"
	val = 3
	obj.Insert(key, val)
	prefix = "ap"
	r = obj.Sum(prefix)
	fmt.Println(r)

	key = "b"
	val = 2
	obj.Insert(key, val)

	prefix = "a"
	r = obj.Sum(prefix)
	fmt.Println(r)
}
