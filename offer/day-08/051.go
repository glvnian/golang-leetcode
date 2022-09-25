package day_08

import "math"

func maxPathSum(root *TreeNode) int {
	var maxSum = []int{math.MinInt32}
	maxPathSumDfs(root, maxSum)
	return maxSum[0]
}

func maxPathSumDfs(root *TreeNode, maxSum []int) int {
	if root == nil {
		return 0
	}

	var maxSumLeft = []int{math.MinInt32}
	var left = Max(0, maxPathSumDfs(root.Left, maxSumLeft))

	var maxSumRight = []int{math.MinInt32}
	var right = Max(0, maxPathSumDfs(root.Right, maxSumRight))

	maxSum[0] = Max(maxSumLeft[0], maxSumRight[0])
	maxSum[0] = Max(maxSum[0], root.Val+left+right)

	return root.Val + Max(left, right)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
