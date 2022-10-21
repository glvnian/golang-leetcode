package day_14

import "fmt"

func rob(nums []int) int {
	dp := make([]int, 2)

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) >= 2 {
		dp[0] = nums[0]
		dp[1] = max(dp[0], nums[1])
	}

	for i := 2; i < len(nums); i++ {
		dp[i%2] = max(dp[(i-1)%2], dp[i%2]+nums[i])
	}
	return max(dp[0], dp[1])
}

func robV1(nums []int) int {
	dp := make([]int, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		dp[0] = nums[0]
		dp[1] = max(dp[0], nums[1])
		return dp[1]
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
