package day_01

import (
	"math"
)

func maxProduct(words []string) int {
	var field = make([]int64, len(words))
	var max float64
	for i, word := range words {
		for _, ch := range []byte(word) {
			field[i] |= 1 << (ch - 'a')
			// 1 << (ch - 'a') 字符为a，则第一个位是1，数值为1，这字符为b，数值为10
			// field[i] |= 1  等于 field[i] = field[i] ｜ 1 << (ch - 'a') ，如a|b = 1|10 = 11
		}
	}

	for i, _ := range words {
		for j := i + 1; j < len(words); j++ {
			if field[i]&field[j] == 0 {
				max = math.Max(max, float64(len(words[i])*len(words[j])))
			}
		}
	}
	return int(max)
}

func maxProductV1(words []string) int {
	var max float64
	var field = make([][26]bool, len(words))
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			field[i][word[j]-'a'] = true
		}
	}

	for i := 0; i < len(words)-1; i++ {

		for j := i + 1; j < len(words); j++ {
			var k = 0
			for ; k < 26; k++ {
				if field[i][k] && field[j][k] {
					break
				}
			}

			if k == 26 {
				t := float64(len(words[i]) * len(words[j]))
				max = math.Max(t, max)
			}

		}
	}
	return int(max)
}
