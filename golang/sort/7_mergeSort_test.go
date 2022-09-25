package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 4, 6, 9, 2, 5, 3, 11, 34, 5, 36}
	fmt.Println(nums)
	mergeSort(nums)
	fmt.Println(nums)

}
