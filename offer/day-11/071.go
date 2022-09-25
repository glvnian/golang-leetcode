package day_11

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
	sum  int
}

func Constructor(w []int) Solution {

	var nums = make([]int, 0)
	var sum int
	for i := 0; i < len(w); i++ {
		sum += w[i]
		nums = append(nums, sum)
	}
	fmt.Println(sum)
	fmt.Println(nums)
	return Solution{
		nums: nums,
		sum:  sum,
	}
}

func (this *Solution) PickIndex() int {
	num := rand.Intn(this.sum)
	return pick(this.nums, num)
}

func pick(nums []int, target int) int {
	var start, end, mid int
	start = 0
	end = len(nums)
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			if mid == 0 || nums[mid-1] <= target {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
