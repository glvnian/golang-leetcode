package day_03

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return nil
	}

	// m key: 字符 value：字符出现的次数
	var m [26]int
	var i = 0
	for ; i < len(p); i++ {
		m[p[i]-'a']++
		m[s[i]-'a']--

	}
	if isAnagrams(m) {
		res = append(res, i-len(p))
	}
	// 3  --> 3-3
	for i = len(p); i < len(s); i++ {
		m[s[i]-'a']--
		m[s[i-len(p)]-'a']++
		if isAnagrams(m) {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}

func isAnagrams(m [26]int) bool {
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
