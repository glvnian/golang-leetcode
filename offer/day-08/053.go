package day_08

import "fmt"

// 遍历
func inorderSuccessorV1(root *TreeNode, p *TreeNode) *TreeNode {
	var found bool
	var stack = make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if found {
			return root
		}
		//if root.Val == p.Val {
		if root == p {
			fmt.Println("found")
			found = true
		}

		root = root.Right
	}

	return nil

}

// 二叉树遍历
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var r *TreeNode
	for root != nil {

		// 如果当前数大于p。那边它有可能就是当前要找的值
		// 也可能不是,但需要找的一定在它的左边
		if root.Val > p.Val {
			r = root
			root = root.Left
		} else {
			root = root.Right
		}

	}

	return r
}
