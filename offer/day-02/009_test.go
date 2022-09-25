package day_02

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	var k = 100
	var nums = []int{10, 5, 2, 6}
	res := numSubarrayProductLessThanK(nums, k)
	fmt.Println(res)
}
