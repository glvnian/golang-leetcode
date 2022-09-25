package day_03

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	var s1, s2 string
	var bo bool
	s1 = "ab"
	s2 = "eidbaooo"
	bo = checkInclusion(s1, s2)
	fmt.Println(bo)

	s1 = "abc"
	s2 = "eidbaooo"
	bo = checkInclusion(s1, s2)
	fmt.Println(bo)
}
