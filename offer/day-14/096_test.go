package day_14

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"

	r := isInterleave(s1, s2, s3)
	fmt.Println(r)
}
