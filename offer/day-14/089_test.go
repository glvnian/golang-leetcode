package day_14

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	nums = []int{2, 7, 9, 3, 1}
	//nums = []int{1, 1, 1}
	r := rob(nums)
	fmt.Println(r)
}
