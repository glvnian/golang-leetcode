package day_07_queue

import (
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var max = math.MinInt32
	var queue []*TreeNode
	var count, next int

	queue = append(queue, root)
	count = 1
	for len(queue) != 0 {
		node := queue[0]
		max = Max(max, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
			next++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			next++
		}
		queue = queue[1:]
		count--

		if count == 0 {
			res = append(res, max)
			max = math.MinInt32
			count = next
			next = 0
		}
	}

	return res
}

func largestValuesV1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var max = math.MinInt32
	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)

	for len(queue1) != 0 {
		node := queue1[0]
		max = Max(max, node.Val)

		if node.Left != nil {
			queue2 = append(queue2, node.Left)
		}
		if node.Right != nil {
			queue2 = append(queue2, node.Right)
		}

		queue1 = queue1[1:]

		if len(queue1) == 0 {
			res = append(res, max)
			queue1 = queue2
			queue2 = nil
			max = math.MinInt32
		}
	}

	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
