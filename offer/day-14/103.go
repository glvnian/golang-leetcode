package day_14

func coinChange(coins []int, amount int) int {
	maxCount := amount + 1
	dp := make([]int, maxCount)
	for i := range dp {
		dp[i] = maxCount
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				k := i / coins[j]
				dp[i] = min(dp[i], dp[i-coins[j]*k]+k)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinChangeV1(coins []int, amount int) int {
	maxCount := amount + 1
	dp := make([]int, maxCount)
	for i := range dp {
		dp[i] = maxCount
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// f(i,j)
// f(i,j) == min(f(i-1,j-k*nums[i-1])+k)
// f(0,j) == 0
// f(i,0) == 0
