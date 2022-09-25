package day_02

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, right, count int
	var product = 1
	for right = 0; right < len(nums); right++ {

		product *= nums[right]
		fmt.Println(">>", left, right)
		for left <= right && k <= product {
			fmt.Println("---", left, right)
			product /= nums[left]
			left++

		}
		fmt.Println(">>>", left, right)

		count += right - left + 1

	}

	return count
}
