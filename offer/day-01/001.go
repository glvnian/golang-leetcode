package day_01

import "math"

func divide(a int, b int) int {
	var result int
	var tag int

	if b == 0 {
		return 0
	}
	if b == 1 {
		return a
	}

	// 通过tag 判断 结果最终是正数还是复数
	if a > 0 {
		a = -a
		tag++
	}
	if b > 0 {
		b = -b
		tag++
	}

	//
	var (
		tmp   int
		value int
	)

	for a <= b {
		tmp = 1
		value = b
		for a <= value+value {
			tmp += tmp
			value += value
		}
		a = a - value
		result = result + tmp
	}

	if tag == 1 {
		result = -result
	}
	if result < -1<<31 || result > 1<<31-1 {
		result = math.MaxInt32
	}
	return result
}
