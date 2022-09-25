package day_06

func dailyTemperatures(temperatures []int) []int {
	var res []int = make([]int, len(temperatures))
	var stack []int = make([]int, 0) // val 为temperatures 位置

	for i, _ := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			var prev = stack[len(stack)-1]
			res[prev] = i - prev
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
