package sort

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4, 5}
	swap(nums, 1, 2)
	fmt.Println(nums)
}

func TestBubbleSortV1(t *testing.T) {
	var nums []int
	nums = []int{10, 2, 3, 24, 5, 0, 1}
	fmt.Println(nums)
	//bubbleSortV1(nums)
	//bubbleSortV2(nums)
	bubbleSortV3(nums)
	fmt.Println(nums)
}
