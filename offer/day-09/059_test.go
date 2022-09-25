package day_09

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	var k, r, val int
	var nums = []int{4, 5, 8, 2}
	k = 3
	obj := Constructor(k, nums)
	val = 3
	r = obj.Add(val)
	fmt.Println(r)

	val = 5
	r = obj.Add(val)
	fmt.Println(r)

	val = 10
	r = obj.Add(val)
	fmt.Println(r)

	val = 9
	r = obj.Add(val)
	fmt.Println(r)

	val = 4
	r = obj.Add(val)
	fmt.Println(r)
}
