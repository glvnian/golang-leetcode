package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	var nums []int
	nums = []int{10, 2, 3, 24, 5, 0, 1}
	nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	heapSort(nums)
	fmt.Println("sort: ", nums)

}
