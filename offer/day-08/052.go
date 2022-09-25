package day_08

func increasingBST(root *TreeNode) *TreeNode {
	var first, prev *TreeNode
	var stack = make([]*TreeNode, 0)
	var cur = root
	for cur != nil || len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Right = cur
		} else {
			first = cur
		}
		prev = cur
		cur.Left = nil
		cur = cur.Right
	}

	return first
}
