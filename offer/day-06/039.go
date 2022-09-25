package day_06

import "fmt"

// 单调栈
func largestRectangleArea(heights []int) int {
	var maxArea int
	var stack = make([]int, 0)
	stack = append(stack, -1)

	for i := 0; i < len(heights); i++ {
		//fmt.Println(i, heights[i], stack[len(stack)-1])
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			weight := i - stack[len(stack)-1] - 1

			area := height * weight
			maxArea = Max(maxArea, area)
		}

		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		weight := len(heights) - stack[len(stack)-1] - 1
		area := height * weight
		maxArea = Max(maxArea, area)
	}

	return maxArea
}

// 分治法
func largestRectangleAreaV2(heights []int) int {
	return helper(heights, 0, len(heights))
}

func helper(heights []int, start, end int) int {
	fmt.Println(start, end)
	if start == end {
		return 0
	}
	if start+1 == end {
		return heights[start]
	}
	var minIndex, area, maxArea int
	minIndex = start

	for i := start; i < end; i++ {
		if heights[minIndex] > heights[i] {
			minIndex = i
		}
	}

	area = heights[minIndex] * (end - start)
	left := helper(heights, start, minIndex)
	right := helper(heights, minIndex+1, end)
	maxArea = Max(left, right)
	maxArea = Max(area, maxArea)
	return maxArea
}

// 暴力模式
func largestRectangleAreaV1(heights []int) int {
	fmt.Println(heights)
	var maxArea int
	for i := 0; i < len(heights); i++ {
		min := heights[i]
		fmt.Println(">>>>>>>", i)
		for j := i; j < len(heights); j++ {
			min = Min(min, heights[j])
			area := min * (j - i + 1)
			maxArea = Max(maxArea, area)
			fmt.Println(j+1, area, maxArea)
		}
	}
	return maxArea
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
