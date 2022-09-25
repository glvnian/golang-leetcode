package day_06

import "fmt"

func asteroidCollision(asteroids []int) []int {
	var stack []int
	stack = make([]int, 0)
	fmt.Println(asteroids)
	for _, num := range asteroids {
		for len(stack) != 0 && stack[len(stack)-1] > 0 && -num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 && stack[len(stack)-1] > 0 && -num == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] < 0 || num > 0 {
			stack = append(stack, num)
		}

	}
	return stack
}

func asteroidCollisionV1(asteroids []int) []int {
	var stack []int
	stack = make([]int, 0)
	for _, num := range asteroids {
		// stack 存在数据
		// num 往左 stack最后一个数据往右，并且num的绝对值大于stack最后一个数据。这个时候取消 stack 最后一位
		for len(stack) != 0 && stack[len(stack)-1] > 0 && -num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		// stack 存在数据  num向左，
		// 并且stack最后一个数据和当前num绝对值相等. 不添加num，同时把数组最后一个数拿掉
		if len(stack) != 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] == -num {
			stack = stack[:len(stack)-1]
		} else if len(stack) == 0 || stack[len(stack)-1] < 0 || num > 0 {
			// 当stack为空，或者最后stack最后一个数字为复数，向左
			// 或者当前num数字大于0，但比stack最后一个数字的绝对值大。或者同方向
			stack = append(stack, num)
		}

	}
	return stack
}
