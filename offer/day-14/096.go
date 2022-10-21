package day_14

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if l1+l2 != l3 {
		return false
	}
	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}

	dp := make([][]bool, l1+1)
	dp[0] = make([]bool, l2+1)
	dp[0][0] = true
	for i := 0; i < l1; i++ {
		dp[i+1] = make([]bool, l2+1)
		dp[i+1][0] = s1[i] == s3[i] && dp[i][0]
	}
	for j := 0; j < l2; j++ {
		dp[0][j+1] = s2[j] == s3[j] && dp[0][j]
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			c1 := s1[i]
			c2 := s2[j]
			c3 := s3[i+j+1]

			dp[i+1][j+1] = (c1 == c3 && dp[i][j+1]) || (c2 == c3 && dp[i+1][j])

		}
	}

	return dp[l1][l2]
}

// len(s1)=5,len(s2) = 6 , len(s3) = 11 ,   f(4,5)
// f(i,j)
// 如果s1[i] == s3[i+j+1]   :  f(i,j) = f(i-1,j)
// 如果s2[j] == s3[i+j+1]  : f(i,j) = f(i,j-i)
// 如果s1[i] != s3[i+j+1] && s2[j] != s3[i+j+1]  false
