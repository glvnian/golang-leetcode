package day_03

import (
	"fmt"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	var s string
	var c int
	//s = "abc"
	s = "aaa"
	//s = "race a car"
	c = countSubstrings(s)
	//c = countSubstringsV1(s)
	fmt.Println(c)
}

func TestCCC(t *testing.T) {
	// a = [a,a,a]
	// a1 = a,aa,aaa
	// a2 = a,aa,
	// a3 = a
}
