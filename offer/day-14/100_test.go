package day_14

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	triangle = [][]int{{-1}, {2, 3}, {1, -1, -3}}
	r := minimumTotal(triangle)
	fmt.Println(r)

}
