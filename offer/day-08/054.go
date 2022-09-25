package day_08

func convertBST(root *TreeNode) *TreeNode {

	var stack = make([]*TreeNode, 0)
	var cur = root
	var sum int
	for cur != nil || len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += cur.Val
		cur.Val = sum

		cur = cur.Left
	}

	return root
}
