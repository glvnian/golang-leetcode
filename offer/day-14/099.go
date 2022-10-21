package day_14

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	arr := make([][]int, 2)
	for i := 0; i < 2; i++ {
		arr[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				arr[0][0] = grid[0][0]
				continue
			}
			if i == 0 {
				arr[0][j] = arr[0][j-1] + grid[0][j]
				continue
			}
			if j == 0 {
				arr[i%2][0] = arr[(i-1)%2][0] + grid[i][0]
				continue
			}
			arr[i%2][j] = min(arr[(i-1)%2][j], arr[i%2][j-1]) + grid[i][j]
		}

	}

	return arr[(row-1)%2][col-1]
}

func minPathSumV1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	arr := make([][]int, row)
	for i := 0; i < row; i++ {
		arr[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				arr[i][j] = grid[0][0]
				continue
			}
			if i == 0 {
				arr[0][j] = arr[0][j-1] + grid[0][j]
				continue
			}
			if j == 0 {
				arr[i][0] = arr[i-1][0] + grid[i][0]
				continue
			}
			arr[i][j] = min(arr[i-1][j], arr[i][j-1]) + grid[i][j]
		}

	}

	return arr[row-1][col-1]
}

// f(i,j)
// f(0,j) = f(0,j-1) + g[0][j]
// f(i,0) = f(i-1,0) + g[i][0]
// f(i,j) = min(f(i-1,j),f(i,j-1)) + g[i][j]
