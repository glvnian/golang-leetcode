package day_10

import (
	"fmt"
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	var nums []int
	var r int
	nums = []int{3, 10, 5, 25, 2, 8}
	r = findMaximumXOR(nums)
	fmt.Println(r)
}
