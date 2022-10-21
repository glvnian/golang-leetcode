package day_14

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	if len(costs) == 1 {
		return min(min(costs[0][0], costs[0][1]), costs[0][2])
	}
	row := len(costs)
	col := len(costs[0])
	var dp = make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < 3; j++ {
			if i == 0 {
				dp[0][j] = costs[0][j]
			} else {
				dp[i%2][j%3] = min(dp[(i-1)%2][(j+1)%3], dp[(i-1)%2][(j+2)%3]) + costs[i][j]
			}

		}
	}
	return min(min(dp[(row-1)%2][0], dp[(row-1)%2][1]), dp[(row-1)%2][2])
}

//
