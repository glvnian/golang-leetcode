package day_12

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	var intervals, r [][]int
	intervals = [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	intervals = [][]int{
		{1, 3},
		{2, 6},
		{9, 16},
		{15, 18},
	}
	fmt.Println(intervals)
	r = merge(intervals)

	fmt.Println(r)

}
