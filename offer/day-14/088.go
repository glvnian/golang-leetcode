package day_14

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		fmt.Println(i, dp[i-1], dp[i-2])
		dp[i%2] = min(dp[0], dp[1]) + cost[i]
	}
	return min(dp[0], dp[1])
}

func helper() {

}

func minCostClimbingStairsV2(cost []int) int {
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		fmt.Println(i, dp[i-1], dp[i-2])
		dp[i%2] = min(dp[0], dp[1]) + cost[i]
	}
	return min(dp[0], dp[1])
}

func MinCostClimbingStairsV1(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// f(0) == cast[0]
// f(1) == min(f(0),cast[1])
