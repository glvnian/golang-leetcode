package day_14

import "fmt"

func minFlipsMonoIncr(s string) int {
	dp := [2][2]int{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 {
			if ch == '0' {
				dp[0][0] = 0
				dp[1][0] = 1
			}
			if ch == '1' {
				dp[0][0] = 1
				dp[1][0] = 0
			}
		}
		if i > 0 {
			if ch == '0' {
				dp[0][i%2] = dp[0][(i-1)%2]
				dp[1][i%2] = dp[1][(i-1)%2] + 1
			}
			if ch == '1' {
				dp[0][i%2] = dp[0][(i-1)%2] + 1
				dp[1][i%2] = min(dp[0][(i-1)%2], dp[1][(i-1)%2])
			}
		}
	}
	return min(dp[0][(len(s)-1)%2], dp[1][(len(s)-1)%2])
}

func minFlipsMonoIncrV1(s string) int {
	dp := [2][]int{}
	dp[0] = make([]int, len(s))
	dp[1] = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 {
			if ch == '0' {
				dp[0][i] = 0
				dp[1][i] = 1
			}
			if ch == '1' {
				dp[0][i] = 1
				dp[1][i] = 0
			}
		}
		if i > 0 {
			if ch == '0' {
				dp[0][i] = dp[0][i-1]
				dp[1][i] = dp[1][i-1] + 1
			}
			if ch == '1' {
				dp[0][i] = dp[0][i-1] + 1
				dp[1][i] = min(dp[0][i-1], dp[1][i-1])
			}
		}
	}
	fmt.Println(s)
	fmt.Println(dp[0])
	fmt.Println(dp[1])
	return min(dp[0][len(s)-1], dp[1][len(s)-1])
}

// f(n)  0 结尾。  0: f（n-1） ； 1: f(n-1) +1
// g(n). 1 结尾.   0: g(n-1) == f(n-1)  | 1: g(n-1)
// 结果：  min(f(n),g(n))

// dp[0][n] == f[n]
// dp[1][n] == g[n]
