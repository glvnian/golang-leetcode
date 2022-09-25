package day_13

import (
	"fmt"
	"leetcode/util"
	"testing"
)

func TestSubsets(t *testing.T) {
	var nums []int
	nums = []int{
		1, 2, 3,
	}
	//nums = []int{
	//	0,
	//}
	//nums = []int{
	//	9, 0, 3, 5, 7,
	//}
	r := subsets(nums)
	fmt.Println(util.InterfaceToString(r))
}

func TestSub(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a[:len(a)-1])
}
