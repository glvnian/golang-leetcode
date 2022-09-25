package day_08

func sumNumbers(root *TreeNode) int {

	var path int
	return sumDfs(root, path)
}

func sumDfs(root *TreeNode, path int) int {
	if root == nil {
		return 0
	}

	path = path*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return path
	}

	return sumDfs(root.Left, path) + sumDfs(root.Right, path)
}
