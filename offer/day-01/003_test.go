package day_01

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {
	n := 2
	a := countBits(n)
	fmt.Println(n, a)
	n = 5
	a = countBits(n)
	fmt.Println(n, a)
}
