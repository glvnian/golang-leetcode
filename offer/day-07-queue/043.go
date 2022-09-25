package day_07_queue

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func TConstructor(root *TreeNode) CBTInserter {

	var parent = root
	var queue = []*TreeNode{}

	queue = append(queue, parent)

	for len(queue) != 0 && parent.Left != nil && parent.Right != nil {
		queue = append(queue, parent.Left)
		queue = append(queue, parent.Right)
		queue = queue[1:]
		parent = queue[0]
	}

	return CBTInserter{
		root:  root,
		queue: queue,
	}

}

func (this *CBTInserter) Insert(v int) int {
	var newNode = &TreeNode{Val: v}
	var parent *TreeNode

	parent = this.queue[0]
	if parent.Left == nil {
		parent.Left = newNode
	} else {
		parent.Right = newNode
		this.queue = append(this.queue, parent.Left)
		this.queue = append(this.queue, parent.Right)
		this.queue = this.queue[1:]
	}
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func (this *CBTInserter) InsertNil(v int) int {
	var newNode = &TreeNode{Val: v}
	var parent *TreeNode

	parent = this.queue[0]
	for parent == nil {
		this.queue = this.queue[1:]
		parent = this.queue[0]
	}
	if parent.Left == nil {
		parent.Left = newNode
	} else {
		parent.Right = newNode
		if parent.Left.Val == -1 {
			this.queue = append(this.queue, nil)
		} else {
			this.queue = append(this.queue, parent.Left)
		}
		if parent.Right.Val == -1 {
			this.queue = append(this.queue, nil)
		} else {
			this.queue = append(this.queue, parent.Right)
		}
		this.queue = this.queue[1:]
	}
	return parent.Val
}
