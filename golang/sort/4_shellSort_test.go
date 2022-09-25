package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	var nums []int
	nums = []int{10, 2, 3, 24, 5, 0, 1}
	fmt.Println(nums)
	shellSort(nums)
	fmt.Println(nums)
}
