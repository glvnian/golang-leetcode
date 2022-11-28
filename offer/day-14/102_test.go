package day_14

import (
	"fmt"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3

	nums = []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	target = 1
	r := findTargetSumWays(nums, target)
	fmt.Println(r)
}
