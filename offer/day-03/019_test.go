package day_03

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	var s string
	var b bool
	s = "abac"
	//s = "race a car"
	b = validPalindrome(s)
	fmt.Println(b)
}

func TestChar(t *testing.T) {
	var s = "abac"
	fmt.Println(">>", s[0:0+1])
}
