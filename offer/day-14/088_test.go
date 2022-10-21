package day_14

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	r := minCostClimbingStairs(cost)
	fmt.Println(r)
}
