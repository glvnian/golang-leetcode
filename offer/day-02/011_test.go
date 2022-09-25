package day_02

import (
	"fmt"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	var nums = []int{0, 1, 0, 1}
	var res = findMaxLength(nums)
	fmt.Println(res)

	nums = []int{0, 1}
	res = findMaxLength(nums)
	fmt.Println(res)
}
