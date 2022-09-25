package day_11

import (
	"fmt"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	var nums []int
	var r int
	nums = []int{1, 3, 5, 4, 2}
	nums = []int{0, 10, 5, 2}
	nums = []int{0, 10, 15, 112}
	r = peakIndexInMountainArray(nums)
	fmt.Println(r)
}
