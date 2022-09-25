package day_03

//func lengthOfLongestSubstring(s string) int {
//	var i = 0
//	var j = 0
//	var max = 0
//	var m [256]int
//	if len(s) == 0  {
//		return 0
//	}
//	for ; j < len(s); j++ {
//		m[s[j]-'a']++
//		for repeatChat(m) {
//			m[s[i]-'a']--
//			i++
//		}
//		max = Max(max, j-i+1)
//	}
//	return max
//}
//
//func repeatChat(m [256]int) bool {
//	for _, num := range m {
//		if num > 1 {
//			return true
//		}
//	}
//	return false
//}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	var i = 0
	var j = 0
	var max = 0
	var m [256]int
	if len(s) == 0 {
		return 0
	}
	for ; j < len(s); j++ {
		m[s[j]-'a']++
		for m[s[j]-'a'] > 1 {
			m[s[i]-'a']--
			i++
		}
		max = Max(max, j-i+1)
	}
	return max
}
