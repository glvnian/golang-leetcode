package day_14

import (
	"fmt"
	"testing"
)

func TestMinCost(t *testing.T) {
	// costs = [[17,2,17],[16,16,5],[14,3,19]]
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}
	costs = [][]int{
		{17, 2, 17},
	}

	// [[20,18,4],[9,9,10]]
	costs = [][]int{
		{20, 18, 4},
		{9, 9, 10},
	}
	r := minCost(costs)
	fmt.Println(r)
}
