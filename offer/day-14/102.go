package day_14

import "fmt"

func findTargetSumWaysV2(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	fmt.Println("sum:", sum)
	fmt.Println("target:", target)
	if (sum-target)%2 == 1 || sum < target {
		return 0
	}

	neg := (sum - target) / 2
	fmt.Println("neg: ", neg)
	fmt.Println(nums)

	arr := make([][]int, len(nums)+1)
	arr[0] = make([]int, neg+1)
	arr[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		arr[i] = make([]int, neg+1)
		for j := 0; j <= neg; j++ {
			arr[i][j] = arr[i-1][j]
			if j-nums[i-1] >= 0 {
				arr[i][j] = arr[i][j] + arr[i-1][j-nums[i-1]]
			}
		}
	}
	return arr[len(nums)][neg]
}
