package day_03

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var s string
	var b bool
	s = "A man, a plan, a canal: Panama"
	s = "race a car"
	b = isPalindrome(s)
	fmt.Println(b)
}

func TestPrintByte(t *testing.T) {
	fmt.Println(byte('a'))
	fmt.Println(byte('z'))
	fmt.Println(byte('0'))
	fmt.Println(byte('1'))
	fmt.Println(byte('9'))
}
