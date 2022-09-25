package day_08

// map  双指针
func findTarget(root *TreeNode, k int) bool {
	var start, end *TreeNode
	var iterNext = BSConstructor(root)
	var iterPrev = ReversedBSConstructor(root)
	if iterNext.HasNext() {
		start = iterNext.Next()
	}
	if iterPrev.HasPrev() {
		end = iterPrev.Prev()
	}

	for start != end {
		if start.Val+end.Val == k {
			return true
		} else if start.Val+end.Val > k {
			end = iterPrev.Prev()
		} else {
			start = iterNext.Next()
		}
	}

	return false
}

type BSIterator struct {
	stack []*TreeNode
	root  *TreeNode
}

func BSConstructor(root *TreeNode) BSIterator {

	return BSIterator{
		stack: make([]*TreeNode, 0),
		root:  root,
	}
}

func (this *BSIterator) HasNext() bool {
	if len(this.stack) != 0 || this.root != nil {
		return true
	}
	return false
}

func (this *BSIterator) Next() *TreeNode {
	for this.root != nil {
		this.stack = append(this.stack, this.root)
		this.root = this.root.Left
	}
	this.root = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	var node = this.root
	this.root = this.root.Right

	return node
}

type ReversedBSIterator struct {
	stack []*TreeNode
	root  *TreeNode
}

func ReversedBSConstructor(root *TreeNode) ReversedBSIterator {
	return ReversedBSIterator{
		stack: make([]*TreeNode, 0),
		root:  root,
	}
}

func (this *ReversedBSIterator) HasPrev() bool {
	if len(this.stack) != 0 || this.root != nil {
		return true
	}
	return false
}

func (this *ReversedBSIterator) Prev() *TreeNode {
	for this.root != nil {
		this.stack = append(this.stack, this.root)
		this.root = this.root.Right
	}
	this.root = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	var node = this.root
	this.root = this.root.Left
	return node
}

// map 广度遍历
func findTargetV2(root *TreeNode, k int) bool {
	var numMap = make(map[int]bool)
	var queque []*TreeNode

	queque = append(queque, root)

	for len(queque) != 0 {

		root = queque[0]
		queque = queque[1:]

		if _, ok := numMap[k-root.Val]; ok {
			return true
		}
		numMap[root.Val] = true

		if root.Left != nil {
			queque = append(queque, root.Left)
		}

		if root.Right != nil {
			queque = append(queque, root.Right)
		}
	}

	return false
}

// map 深度遍历
func findTargetV1(root *TreeNode, k int) bool {
	var numMap = make(map[int]bool)
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := numMap[k-root.Val]; ok {
			return true
		}
		numMap[root.Val] = true
		root = root.Right
	}

	return false
}
