package day_01

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	works := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	a := maxProduct(works)
	fmt.Println(a)
	works = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	a = maxProduct(works)
	fmt.Println(a)
}
