package day_06

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	var heights []int
	var r int
	heights = []int{2, 4}
	//heights = []int{2, 1, 5, 6, 2, 3}
	fmt.Println(heights)
	r = largestRectangleArea(heights)
	fmt.Println(r)
}
