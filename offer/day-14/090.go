package day_14

func ring(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(ringRob(nums[:len(nums)-1]), ringRob(nums[1:]))
}

func ringRob(nums []int) int {
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
