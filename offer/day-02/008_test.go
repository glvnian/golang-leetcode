package day_02

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var target = 7
	var nums = []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)

	target = 4
	nums = []int{1, 4, 4}
	res = minSubArrayLen(target, nums)
	fmt.Println(res)
}

func TestSearchInts(t *testing.T) {
	var target = 5
	//var nums = []int{2, 3, 1, 2, 4, 3}
	var nums = []int{1, 7, 2, 3, 4, 3, 6, 1}
	var b = sort.SearchInts(nums, target)
	fmt.Println(b)
}
