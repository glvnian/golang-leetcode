package day_14

func minCut(s string) int {
	l := len(s)
	var isPal = make([][]bool, l)
	for i := 0; i < l; i++ {
		isPal[i] = make([]bool, l)
		for j := 0; j <= i; j++ {
			c1 := s[i]
			c2 := s[j]
			if c1 == c2 && (i-j <= 1 || isPal[j+1][i-1]) {
				isPal[j][i] = true
			}
		}
	}

	var dp = make([]int, l)
	for i := 0; i < l; i++ {
		if isPal[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j := 1; j <= i; j++ {
				if isPal[j][i] {
					dp[i] = min(dp[j-1]+1, dp[i])
				}
			}
		}
	}

	return dp[l-1]
}

//  如果 0~i 是回文, f(i) == 0
//  如果 0~i 不是回问，但 j~i 是回文，并且 0<j<i  那么 f(i) = f(j-1) +1
