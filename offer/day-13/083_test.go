package day_13

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	r := permute(nums)

	for _, item := range r {
		fmt.Println(item)
	}
}
