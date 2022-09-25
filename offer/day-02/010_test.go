package day_02

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	var k = 2
	var nums = []int{1, 1, 1}
	var res = subarraySum(nums, k)
	fmt.Println(res)
}
