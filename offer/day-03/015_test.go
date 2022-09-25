package day_03

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {

	var s, p string
	var bo []int
	p = "ab"
	s = "cbaooo"
	bo = findAnagrams(s, p)
	fmt.Println(bo)

	s = "cbaebabacd"
	p = "abc"
	bo = findAnagrams(s, p)
	fmt.Println(bo)
}
