package day_14

import (
	"fmt"
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 4
	r := combinationSum4(nums, target)
	fmt.Println(r)
}

// dp[0] = 1
// dp[1] = dp[1] + dp[0] = 0 + 1 = 1
// dp[2] = dp[2] + dp[1]  1| dp[2] + dp[0] 2 = 3
// dp[3] = dp[]
