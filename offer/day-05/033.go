package day_05

import (
	"fmt"
	"sort"
	"strings"
)

// 素数
func groupAnagramsV1(strs []string) [][]string {
	var hash = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	fmt.Println("hash len:", len(hash))

	var m = map[int][]string{}
	for _, str := range strs {
		var num = 1
		for _, ch := range []byte(str) {
			var site = int(ch - 'a')
			num *= hash[site]
		}
		m[num] = append(m[num], str)
	}

	var res = [][]string{}
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

// 排序
func groupAnagramsV2(strs []string) [][]string {
	var res = [][]string{}
	var m = map[string][]string{}
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		t := strings.Join(s, "")
		if _, ok := m[t]; !ok {
			m[t] = []string{}
		}
		m[t] = append(m[t], str)
	}

	for _, val := range m {
		res = append(res, val)
	}
	return res
}

// 字符计数
func groupAnagrams(strs []string) [][]string {
	var res = [][]string{}
	var m = map[[26]int][]string{}
	for _, str := range strs {
		var cnt [26]int
		for _, c := range str {
			cnt[c-'a']++
		}
		fmt.Println(cnt)
		m[cnt] = append(m[cnt], str)
	}

	for _, val := range m {
		res = append(res, val)
	}
	return res
}
