package day_02

type NumMatrix struct {
	//Row int
	//Col int
	//Sum int
	Sums [][]int
}

//[]int{3, 0, 1, 4, 2},
//[]int{5, 6, 3, 2, 1},
//[]int{1, 2, 0, 1, 5},
//[]int{4, 1, 0, 1, 7},
//[]int{1, 0, 3, 0, 5},
//[][]int
// sum[row][col] = int

func Constructor(matrix [][]int) NumMatrix {
	// 求坐标（0，0） 到任意一个坐标（i，j）的和

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	var n = NumMatrix{}
	//var dd [5][6]int
	var sums = make([][]int, len(matrix)+1)
	sums[0] = make([]int, len(matrix[0])+1)
	for i, lines := range matrix {
		sums[i+1] = make([]int, len(lines)+1)
		for j, col := range lines {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] + col - sums[i][j]
		}

	}
	n.Sums = sums
	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 2, 1, 4, 3
	// 左上角 坐标 [row1,col1]
	// 右上角 坐标 [row1,col2]
	// 左下角 坐标 [row2,col1]
	// 右下角 坐标 [row2,col2]
	// 面积 == [row2,col2] - [row1,col2] - [row2,col1] + [row1,col1]
	var row, col int
	var sums = this.Sums
	row = len(sums)
	if row == 0 {
		return 0
	}
	col = len(sums[0])
	if col == 0 {
		return 0
	}

	if row < row1 || row < row2 || col < col1 || col < col2 {
		return 0
	}
	return sums[row2+1][col2+1] - sums[row1][col2+1] - sums[row2+1][col1] + sums[row1][col1]
}
