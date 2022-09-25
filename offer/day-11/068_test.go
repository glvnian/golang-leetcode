package day_11

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var nums []int
	var target, r int
	nums = []int{1, 3, 5, 6}
	target = 5

	nums = []int{1, 3, 5, 6}
	target = 9

	r = searchInsert(nums, target)
	fmt.Println(r)
}
