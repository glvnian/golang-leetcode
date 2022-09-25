package day_03

func minWindow(s string, t string) string {
	var count, start, end, minStart, minEnd, minLength int
	minLength = len(s) + 1
	var m = make(map[byte]int)
	for _, ch := range []byte(t) {
		if _, ok := m[ch]; !ok {
			m[ch] = 0
		}
		m[ch]++
	}
	count = len(m)

	for end < len(s) || (count == 0 && start < len(s)) {
		if count > 0 {
			if _, ok := m[s[end]]; ok {
				m[s[end]]--
				if m[s[end]] == 0 { // todo:
					count--
				}
			}
			end++
		} else {
			// count = 0
			// 此时可获取符合条件的子字符串。如果
			if (end - start) < minLength {
				minLength = end - start
				minStart = start
				minEnd = end
			}

			if _, ok := m[s[start]]; ok {
				m[s[start]]++
				if m[s[start]] > 0 {
					count++
				}
			}
			start++
		}

	}
	if minLength == len(s)+1 {
		return ""
	}
	return s[minStart:minEnd]
}

func minWindowV2(s string, t string) string {
	var count, start, end, minStart, minEnd, minLength int
	minLength = len(s) + 1
	var m = make(map[byte]int)
	for _, ch := range []byte(t) {
		if _, ok := m[ch]; !ok {
			m[ch] = 0
		}
		m[ch]++
	}
	count = len(m)

	for end < len(s) || (count == 0 && start < len(s)) {
		if count > 0 {
			if _, ok := m[s[end]]; ok {
				m[s[end]]--
				if m[s[end]] == 0 { // todo:
					count--
				}
			}
			end++
		} else {
			// count = 0
			// 此时可获取符合条件的子字符串。如果
			if (end - start) < minLength {
				minLength = end - start
				minStart = start
				minEnd = end
			}

			if _, ok := m[s[start]]; ok {
				m[s[start]]++
				if m[s[start]] > 0 {
					count++
				}
			}
			start++
		}

	}
	if minLength == len(s)+1 {
		return ""
	}
	return s[minStart:minEnd]
}

func minWindowV1(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	var res = ""
	var m = make(map[byte]int)
	for _, ch := range []byte(t) {
		if _, ok := m[ch]; !ok {
			m[ch] = 0
		}
		m[ch]++
	}

	var i, j int
	for ; j < len(s); j++ {

		//
		if _, ok := m[s[j]]; !ok {
			continue
		}
		m[s[j]]--
		for ; isFoundDown(m); i++ {
			if len(res) > j-i+1 || res == "" {
				res = s[i : j+1]
			}

			if _, ok := m[s[i]]; !ok {
				continue
			}
			m[s[i]]++
		}

	}
	return res
}

// false : 代表还有字符未找到
func isFoundDown(m map[byte]int) bool {
	for _, num := range m {
		if num > 0 {
			return false
		}
	}
	return true
}
