package day_03

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		// 如果i是奇数 l=i/2 r=1/2+1  ；
		// 通过i%2 可以确定i是否为奇数。
		// 如果i是偶数。l和r都是i/2
		l, r := i/2, i/2+i%2

		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func countSubstringsV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(s); i++ {
		count += checkPalindrome(s, i, i)
		count += checkPalindrome(s, i, i+1)
	}
	return count
}

func checkPalindrome(s string, start, end int) int {
	var count int
	for start >= 0 && end < len(s) && s[start] == s[end] {
		fmt.Println(start, end, s[start:end+1])
		count++
		start--
		end++
	}
	return count
}

// a = [a,a,a]
// a1 = a,aa,aaa
// a2 = a,aa,
// a3 = a
// count = 6
//func countSubstrings(s string) int {
//	var start, end, count int
//	for ; start < len(s); start++ {
//		for end = start; end < len(s); end++ {
//			if checkPalindrome(s, start, end) {
//				fmt.Println(start, end, s[start:end+1])
//				count++
//			}
//		}
//	}
//	return count
//}
//
//func checkPalindrome(s string, start, end int) bool {
//	for start < end {
//		if s[start] != s[end] {
//			return false
//		}
//		start++
//		end--
//	}
//	return true
//}
