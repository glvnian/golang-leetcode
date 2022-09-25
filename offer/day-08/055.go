package day_08

type BSTIterator struct {
	root  *TreeNode
	stack []*TreeNode
}

func BTConstructor(root *TreeNode) BSTIterator {

	return BSTIterator{
		root:  root,
		stack: make([]*TreeNode, 0),
	}
}

func (this *BSTIterator) Next() int {
	for this.root != nil {
		this.stack = append(this.stack, this.root)
		this.root = this.root.Left
	}

	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	this.root = node.Right
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.root != nil || len(this.stack) != 0 {
		return true
	}
	return false
}
