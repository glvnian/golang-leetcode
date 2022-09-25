package day_06

import (
	"fmt"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	var matrix []string
	var r int
	matrix = []string{"10100", "10111", "11111", "10010"}
	r = maximalRectangle(matrix)
	fmt.Println(r)
}
