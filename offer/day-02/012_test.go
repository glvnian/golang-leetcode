package day_02

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	var nums = []int{1, 7, 3, 6, 6, 5, 6}
	var res = pivotIndex(nums)
	fmt.Println(res)

	// [2, 1, -1]
	//nums = []int{2, 1, -1}
	nums = []int{-1, 1, 2}
	res = pivotIndex(nums)
	fmt.Println(res)
}
