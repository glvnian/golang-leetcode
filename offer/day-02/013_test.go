package day_02

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	var matrix = [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)

	var row1, col1, row2, col2 = 2, 1, 4, 3
	row1, col1, row2, col2 = 1, 1, 2, 2
	//row1, col1, row2, col2 = 1, 2, 2, 4
	var res = obj.SumRegion(row1, col1, row2, col2)
	fmt.Println(res)
}

//["NumMatrix","sumRegion","sumRegion","sumRegion"]
//[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],
//[2,1,4,3],[1,1,2,2],[1,2,2,4]]
