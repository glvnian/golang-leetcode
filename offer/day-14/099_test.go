package day_14

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	//grid = [[1,3,1],[1,5,1],[4,2,1]]
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	r := minPathSum(grid)
	fmt.Println(r)
}
