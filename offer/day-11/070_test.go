package day_11

import (
	"fmt"
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	var nums []int
	var r int
	nums = []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	nums = []int{3, 3, 7, 7, 10, 11, 11}
	nums = []int{1, 3, 3, 7, 7, 11, 11, 12, 12}
	fmt.Println(nums)
	r = singleNonDuplicate(nums)
	fmt.Println(r)
}
