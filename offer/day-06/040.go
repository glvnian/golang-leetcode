package day_06

func maximalRectangle(matrix []string) int {
	var maxRectangle int
	if len(matrix) == 0 {
		return 0
	}
	var nums = make([]int, len(matrix[0]))
	for _, str := range matrix {

		for i, ch := range str {
			num := ch - '0'
			if num == 0 {
				nums[i] = 0
			} else {
				nums[i] = nums[i] + int(num)
			}
		}
		maxRectangle = Max(maxRectangle, LargestRectangleArea(nums))
	}

	return maxRectangle
}

// 单调栈
func LargestRectangleArea(heights []int) int {
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
