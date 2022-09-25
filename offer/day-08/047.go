package day_08

func pruneTreeV1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

//有 bug
func pruneTree(cur *TreeNode) *TreeNode {
	if cur == nil {
		return nil
	}
	var stack = make([]*TreeNode, 0)
	var first, prev *TreeNode
	first = cur
	for cur != nil && len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]

		if cur.Right != nil && cur.Right != prev {
			cur = cur.Right
		} else {
			// 右子树为空或已经搜索完右子树
			stack = stack[:len(stack)-1]
			if cur.Val == 0 && cur.Left == nil && cur.Right == nil {
				if len(stack) != 0 {
					//在这里进行剪枝操作
					if stack[len(stack)-1].Left == cur {
						stack[len(stack)-1].Left = nil
					} else if stack[len(stack)-1].Right == cur {
						stack[len(stack)-1].Right = nil
					}
				}
			}

			//if cur.Val == 0 && cur.Left == cur.Right {
			//	// 判断是否符合剪枝条件
			//	if len(stack) != 0 {
			//		// 在这里进行剪枝操作
			//		//if stack[len(stack)-1].Left == cur {
			//		//	stack[len(stack)-1].Left = nil
			//		//}
			//		//
			//		//if stack[len(stack)-1].Right == cur {
			//		//	stack[len(stack)-1].Right = nil
			//		//}
			//
			//	}
			//}

			prev = cur
			cur = nil
		}

	}

	return first
}
