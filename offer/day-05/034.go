package day_05

import "fmt"

func isAlienSorted(words []string, order string) bool {

	var sort [26]int
	for i, ch := range order {
		sort[ch-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !Sort(words[i], words[i+1], sort) {
			fmt.Println(words[i], words[i+1], Sort(words[i], words[i+1], sort))
			return false
		}
	}
	return true
}

func Sort(str1, str2 string, sort [26]int) bool {
	var j int
	for j < len(str1) && j < len(str2) {
		var ch1 = str1[j]
		var ch2 = str2[j]
		if sort[ch1-'a'] < sort[ch2-'a'] {
			return true
		}
		if sort[ch1-'a'] > sort[ch2-'a'] {
			return false
		}
		j++
	}

	// 相同字符为true
	return j == len(str1)
}
