package day_14

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m := 3
	n := 7

	r := uniquePaths(m, n)
	fmt.Println(r)
}
