package day_07_queue

func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)
	for len(queue1) != 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		if node.Left != nil {
			queue2 = append(queue2, node.Left)
		}
		if node.Right != nil {
			queue2 = append(queue2, node.Right)
		}

		if len(queue1) == 0 {
			res = append(res, node.Val)
			queue1 = queue2
			queue2 = nil
		}
	}

	return res
}

func rightSideViewV1(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var rightNum int
	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)
	rightNum = queue1[0].Val

	res = append(res, rightNum)

	for len(queue1) != 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		if node.Left != nil {
			queue2 = append(queue2, node.Left)
			rightNum = node.Left.Val
		}
		if node.Right != nil {
			queue2 = append(queue2, node.Right)
			rightNum = node.Right.Val
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = nil
			if len(queue1) != 0 {
				res = append(res, rightNum)
			}
		}
	}

	return res
}
