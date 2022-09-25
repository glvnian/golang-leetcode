package day_10

type V1MagicDictionary struct {
	LenMap map[int][]string
}

/** Initialize your data structure here. */
func V1Constructor() V1MagicDictionary {
	return V1MagicDictionary{
		LenMap: map[int][]string{},
	}
}

func (this *V1MagicDictionary) BuildDict(dictionary []string) {
	// 只能换一个字符，那么必定是长度相等的两个字符串才可以
	// 根据字符串长度创建字典，长度相同的字符串写入一个列表
	for _, word := range dictionary {
		if _, ok := this.LenMap[len(word)]; ok {
			this.LenMap[len(word)] = append(this.LenMap[len(word)], word)
		} else {
			this.LenMap[len(word)] = []string{word}
		}
	}
}

func (this *V1MagicDictionary) Search(searchWord string) bool {
	if _, ok := this.LenMap[len(searchWord)]; !ok {
		return false
	}

	wordSlice := this.LenMap[len(searchWord)]

	for _, word := range wordSlice {
		// 对每个相同长度词进行比较进行比较，字符不同的数量超过 1 则不可能符合条件
		diffNum := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diffNum++
			}
			if diffNum > 1 { // 至少两个字符不相同
				break
			}
		}
		if diffNum == 1 { // 只有一个
			return true
		}
	}
	return false
}
