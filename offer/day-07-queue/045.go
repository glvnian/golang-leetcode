package day_07_queue

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftNum int
	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)
	leftNum = queue1[0].Val
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
			queue1 = queue2
			queue2 = nil
			if len(queue1) != 0 {
				leftNum = queue1[0].Val
			}
		}
	}

	return leftNum
}
