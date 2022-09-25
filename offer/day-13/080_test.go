package day_13

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	var n, k int
	n, k = 4, 2
	r := combine(n, k)
	fmt.Println(r)
}
