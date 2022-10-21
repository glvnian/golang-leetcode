package day_14

func uniquePaths(m int, n int) int {
	arr := [2][]int{}
	for i := 0; i < 2; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		arr[i%2] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				arr[i%2][j] = 1
			} else {
				arr[i%2][j] = arr[(i-1)%2][j] + arr[i%2][j-1]
			}
		}
	}
	return arr[(m-1)%2][n-1]
}

func uniquePathsV1(m int, n int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}
	}
	return arr[m-1][n-1]
}

// f(i,j)
// f(0,j) = 1
// f(i,0) = 1
// f(i,j) = f(i-1,j) + f(i,j-1)
