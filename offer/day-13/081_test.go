package day_13

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var candidates = []int{2, 3, 6, 7}
	var target = 7

	r := combinationSum(candidates, target)
	fmt.Println(r)
}
