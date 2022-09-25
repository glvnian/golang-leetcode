package day_01

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var res string
	i := len(a) - 1
	j := len(b) - 1
	var carry uint8
	for i >= 0 || j >= 0 || carry > 0 {
		var ta, tb, ts uint8
		if i >= 0 {
			ta = a[i] - '0'
		}
		if j >= 0 {
			tb = b[j] - '0'
		}
		ts = ta + tb + carry
		carry = 0
		if ts > 1 {
			carry = 1
			ts -= 2
		}
		res = fmt.Sprintf("%d%s", ts, res)
		i--
		j--
	}

	return res
}

// 反转
func addBinaryV1(a string, b string) string {
	a = reverseString(a)
	b = reverseString(b)
	var i, j int
	var quo uint8 = 0
	var sum string
	for i < len(a) || j < len(b) || quo > 0 {
		var t1, t2 uint8
		if i >= len(a) {
			t1 = 0
		} else {
			t1 = a[i] - '0'
		}
		if j >= len(b) {
			t2 = 0
		} else {
			t2 = b[j] - '0'
		}

		ts := t1 + t2 + quo
		quo = 0
		if ts > 1 {
			quo = 1
			ts = ts - 2
		}
		sum = fmt.Sprintf("%s%d", sum, ts)
		i++
		j++
	}
	return reverseString(sum)
}

func reverseString(s string) string {
	var r string
	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}
	return r
}
