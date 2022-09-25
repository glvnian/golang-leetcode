package day_11

import "fmt"

func singleNonDuplicate(nums []int) int {
	var start, end, group, mid int
	group = len(nums) / 2 //4
	start = 0
	end = group - 1 // 4

	for start <= end {
		mid = (start + end) / 2 // 2
		fmt.Println(start, end, mid)
		if nums[2*mid] != nums[2*mid+1] {
			if mid == 0 || nums[2*(mid-1)] == nums[2*(mid-1)+1] {
				return nums[mid*2]
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return nums[len(nums)-1]

}

// 1,1,2,3,3,4,4,8,8
