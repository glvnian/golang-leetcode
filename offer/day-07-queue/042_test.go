package day_07_queue

import (
	"fmt"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	var p, r int
	obj := RConstructor()
	p = 1
	r = obj.Ping(p)
	fmt.Println(r)

	p = 2
	r = obj.Ping(p)
	fmt.Println(r)

	p = 3001
	r = obj.Ping(p)
	fmt.Println(r)

	p = 3002
	r = obj.Ping(p)
	fmt.Println(r)

	p = 6002
	r = obj.Ping(p)
	fmt.Println(r)
}
