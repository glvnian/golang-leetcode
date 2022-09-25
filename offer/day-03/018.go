package day_03

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var start, end int
	start = 0
	end = len(s) - 1
	fmt.Println(s)
	for start <= end {
		if !isNumberOrChar(s[start]) {
			start++
			continue
		}
		if !isNumberOrChar(s[end]) {
			end--
			continue
		}
		//fmt.Println(start, end, s[start], s[end])
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isNumberOrChar(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}
