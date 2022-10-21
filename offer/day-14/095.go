package day_14

func longestCommonSubsequence(s1 string, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	dp := make([][]int, 2)
	dp[0] = make([]int, l2+1)
	for i := 0; i < l1; i++ {
		dp[(i+1)%2] = make([]int, l2+1)
		for j := 0; j < l2; j++ {
			c1 := s1[i]
			c2 := s2[j]
			if c1 == c2 {
				dp[(i+1)%2][j+1] = dp[i%2][j] + 1
				continue
			}
			// 不等
			dp[(i+1)%2][j+1] = max(dp[i%2][j+1], dp[(i+1)%2][j])
		}
	}
	return dp[l1%2][l2]
}

func longestCommonSubsequenceV1(s1 string, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	dp := make([][]int, l1+1, l1+1)
	dp[0] = make([]int, l2+1, l2+1)
	for i := 0; i < l1; i++ {
		dp[i+1] = make([]int, l2+1)
		for j := 0; j < l2; j++ {
			c1 := s1[i]
			c2 := s2[j]
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}
			// 不等
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
		}
	}
	return dp[l1][l2]
}

// f(i,j)   s1[0~i]  s2[0~j]
//  s1[i] == s2[j]  f(i,j) == f(i-1,j-1) + 1
//  s1[i] != s2[j]  f(i,j) == max(f(i-1,j) ,f(i,j-1))
// f(0,j)   s1[0] == s2[j]  = 1    | s1[0] != s2[j]  =  f(0,j-1)
// f(-1,-1) = 0
// f(-1,j) = 0
// f(i,-1) = 0
// 结果 f(len-1,len-1)
