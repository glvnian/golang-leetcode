package day_05

import (
	"fmt"
	"testing"
)

func TestIsAlienSorted(t *testing.T) {
	var words []string
	var order string
	var b bool
	//words = []string{"hello", "leetcode"}
	//order = "hlabcdefgijkmnopqrstuvwxyz"
	words = []string{"word", "world", "row"}
	order = "worldabcefghijkmnpqstuvxyz"
	b = isAlienSorted(words, order)
	fmt.Println(b)
}
