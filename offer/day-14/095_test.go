package day_14

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	s1 := "abcde"
	s2 := "ace"
	s1 = "ezupkr"
	s2 = "ubmrapg"
	a := longestCommonSubsequence(s1, s2)

	fmt.Println(a)
}
