package races

import (
	"fmt"
	"math"
	"sort"
)

func Min(a, b int) int {
	if a*a > b*b {
		return b
	} else if a*a == b*b {
		if a > b {
			return a
		}
	}
	return a
}

// [-4,-2,1,4,8]
// left = 0 ; r = 4 ; mid = 2
// left = 0 ; right = 2; mid = 1
// left = 1 ; right = 2; mid = 1
func findClosestNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	var left, mid int
	var right = len(nums) - 1
	for left <= right {
		fmt.Println(mid, left, right)
		if nums[mid] < 0 {
			left = mid + 1

		} else if nums[mid] > 0 {
			right = mid - 1

		} else {
			return nums[mid]
		}
		mid = (right + left) / 2
		fmt.Println(mid, left, right)
	}
	if left >= len(nums) {
		left = len(nums) - 1
	}
	if right < 0 {
		right = 0
	}
	return Min(nums[left], nums[right])
}

func findClosestNumberV1(nums []int) int {
	var min = math.MaxInt32
	for _, nums := range nums {
		min = Min(min, nums)
	}
	return min
}
