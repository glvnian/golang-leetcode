package day_05

func isAnagram(s string, t string) bool {

	if len(s) != len(t) || s == t {
		return false
	}

	var m = map[byte]int{}
	for _, ch := range []byte(s) {
		m[ch]++
	}

	for _, ch := range []byte(t) {
		m[ch]--
		if m[ch] == 0 {
			delete(m, ch)
		}
	}

	if len(m) != 0 {
		return false
	}

	return true
}

func isAnagramV1(s string, t string) bool {

	if len(s) != len(t) || s == t {
		return false
	}

	var m [26]int
	for _, ch := range []byte(s) {
		m[ch-'a']++
	}

	for _, ch := range []byte(t) {
		m[ch-'a']--
	}

	for _, num := range m {
		if num != 0 {
			return false
		}
	}

	return true
}
