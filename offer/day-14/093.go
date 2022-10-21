package day_14

func lenLongestFibSubseq(arr []int) int {
	numMap := map[int]int{}
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = i
	}

	dp := make([][]int, len(arr))
	maxSize := 0
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = i
		dp[i] = make([]int, len(arr))
		for j := i - 1; j >= 0; j-- {
			if j == 0 {
				dp[i][j] = 2
			} else {
				target := arr[i] - arr[j]
				if k, ok := numMap[target]; ok && i > j && j > k {
					dp[i][j] = max(dp[j][k]+1, dp[i][j])
				} else {
					dp[i][j] = 2
				}
			}
			maxSize = max(maxSize, dp[i][j])

		}
	}

	if maxSize == 2 {
		return 0
	}
	return maxSize
}

// i > j > k > 0
// A[i] = A[j] + A[k]  ===> dp[i][j] = dp[j][k] + 1
// [1,3,7,11,12,14,18]
// 1 + 11 = 12 | 3 + 11 = 14
