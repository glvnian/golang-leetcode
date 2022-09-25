package day_08

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历-递归
func inorderTraversalDfs(root *TreeNode) []int {
	var nodes = make([]int, 0)
	inDfs(root, nodes)
	return nodes
}
func inDfs(root *TreeNode, nodes []int) {
	var node = root
	if node != nil {
		nodes = append(nodes, node.Val)
		inDfs(node.Left, nodes)
		inDfs(node.Right, nodes)
	}
}

// 中序遍历-迭代
func inorderTraversal(root *TreeNode) []int {
	var res = make([]int, 0)
	var stack = make([]*TreeNode, 0)
	var cur = root
	for cur != nil || len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}

// 前序遍历-递归
func preOrderTraversalDfs(root *TreeNode) []int {
	var nodes = make([]int, 0)
	preDfs(root, nodes)
	return nodes
}
func preDfs(root *TreeNode, nodes []int) {
	var node = root
	if node != nil {
		preDfs(node.Left, nodes)
		nodes = append(nodes, node.Val)
		preDfs(node.Right, nodes)
	}
}

// 前序遍历-递归
func preOrderTraversal(root *TreeNode) []int {
	var res = make([]int, 0)
	var stack = make([]*TreeNode, 0)
	var cur = root
	for cur != nil || len(stack) != 0 {

		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Right
	}

	return res
}

// 后序遍历-递归
func proOrderTraversalDfs(root *TreeNode) []int {
	var nodes = make([]int, 0)
	proDfs(root, nodes)
	return nodes
}

func proDfs(root *TreeNode, nodes []int) {
	var node = root
	if node != nil {
		proDfs(node.Left, nodes)
		proDfs(node.Right, nodes)
		nodes = append(nodes, node.Val)
	}
}

func proOrderTraversal(root *TreeNode) []int {
	var res = make([]int, 0)
	var stack = make([]*TreeNode, 0)
	var cur = root
	var prev *TreeNode

	for cur != nil || len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		if node.Right != nil && node.Right != prev {
			cur = node.Right
		} else {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			prev = node
			cur = nil
		}

	}

	return res
}
