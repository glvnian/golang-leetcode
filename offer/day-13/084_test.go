package day_13

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2, 3}
	r := permuteUnique(nums)

	for _, item := range r {
		fmt.Println(item)
	}

	fmt.Println(len(r))
}
