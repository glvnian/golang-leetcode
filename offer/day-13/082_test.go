package day_13

import (
	"fmt"
	"testing"
)

func TestCombinationSumTwo(t *testing.T) {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	//
	candidates = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	candidates = []int{3, 1, 3, 5, 1, 1}
	target = 8
	r := combinationSumTwo(candidates, target)
	fmt.Println(">>>>")
	fmt.Println(r)
}
