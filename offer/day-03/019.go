package day_03

func validPalindrome(s string) bool {

	var start, end int
	start = 0
	end = len(s) - 1

	for start <= end {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}

	if start == len(s)/2 || IsPalindrome(start, end-1, s) || IsPalindrome(start+1, end, s) {
		return true
	}

	return false
}

func IsPalindrome(start, end int, s string) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
