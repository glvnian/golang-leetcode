package day_03

// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
// https://leetcode-cn.com/problems/MPnaiL/
// 双子针
// 首先初始化s1。新增字符+1
// s2： 添加字符-1，删除字符+1
// s1: a,b,c
// s2: a,d,c,b,a
func checkInclusionV1(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	// m key: 字符 value：字符出现的次数
	var m [26]int
	var i = 0
	for ; i < len(s1); i++ {
		m[s1[i]-'a']++
		m[s2[i]-'a']--
	}

	if isInclusionV1(m) {
		return true
	}
	// 3  --> 3-3
	//
	for i = len(s1); i < len(s2); i++ {
		m[s2[i]-'a']--
		m[s2[i-len(s1)]-'a']++
		if isInclusionV1(m) {
			return true
		}
	}
	return false
}

func isInclusionV1(m [26]int) bool {
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func isInclusion(m map[byte]int) bool {
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	// m key: 字符 value：字符出现的次数
	var m = make(map[byte]int)
	var i = 0
	for ; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}

	if isInclusion(m) {
		return true
	}
	// 3  --> 3-3
	//
	for i = len(s1); i < len(s2); i++ {
		m[s2[i]]--
		m[s2[i-len(s1)]]++
		if isInclusion(m) {
			return true
		}
	}
	return false
}
