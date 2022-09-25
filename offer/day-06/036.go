package day_06

import (
	"fmt"
	"strconv"
)

// 优化
func evalRPN(tokens []string) int {
	// 数字
	l := (len(tokens) + 1) / 2
	var stack = make([]int, l)
	var i int = -1
	for _, str := range tokens {
		switch str {
		case "+":
			i--
			stack[i] = stack[i] + stack[i+1]
		case "-":
			i--
			stack[i] = stack[i] - stack[i+1]
		case "*":
			i--
			stack[i] = stack[i] * stack[i+1]
		case "/":
			i--
			stack[i] = stack[i] / stack[i+1]
		default:
			val, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err.Error())
			}
			i++
			stack[i] = val
		}
	}
	return stack[0]
}

func evalRPNV1(tokens []string) int {
	// 数字
	l := (len(tokens) + 1) / 2
	var stack = make([]int, l)
	var i int = -1
	for _, str := range tokens {
		val, err := strconv.Atoi(str)
		// 字符串
		if err != nil {
			switch str {
			case "+":
				num2 := stack[i]
				i--
				num1 := stack[i]
				stack[i] = num1 + num2
			case "-":
				num2 := stack[i]
				i--
				num1 := stack[i]
				stack[i] = num1 - num2
			case "*":
				num2 := stack[i]
				i--
				num1 := stack[i]
				stack[i] = num1 * num2
			case "/":
				num2 := stack[i]
				i--
				num1 := stack[i]
				stack[i] = num1 / num2
			default:
				fmt.Printf("str:%s \n", str)
			}

		} else {
			// 数字
			i++
			stack[i] = val
		}
	}

	return stack[0]
}
