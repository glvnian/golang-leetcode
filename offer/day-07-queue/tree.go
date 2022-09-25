package day_07_queue

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printQueue(queue []*TreeNode) {
	for _, q := range queue {
		fmt.Println(q.Val)
	}
}

func printTreeNode(node *TreeNode) {
	var queue []*TreeNode
	queue = append(queue, node)

	for len(queue) != 0 {
		parent := queue[0]
		fmt.Printf("%d ", parent.Val)
		if parent.Left != nil {
			queue = append(queue, parent.Left)
		}
		if parent.Right != nil {
			queue = append(queue, parent.Right)
			parent = queue[0]
		}
		queue = queue[1:]
	}
	fmt.Println()
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	dfs(node.Left)
	dfs(node.Right)
}
