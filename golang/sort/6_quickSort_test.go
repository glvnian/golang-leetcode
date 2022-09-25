package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 4, 6, 9, 2, 5, 3, 11, 34, 5, 36}
	fmt.Println(nums)
	quickSort(nums)
	fmt.Println(nums)
}
