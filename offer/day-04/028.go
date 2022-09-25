package day_04

type CNode struct {
	Val   int    `json:"val"`
	Prev  *CNode `json:"-"`
	Next  *CNode `json:"-"`
	Child *CNode `json:"-"`
}

// 递归
func flattenV1(root *CNode) *CNode {
	flattenTail(root)
	return root
}

func flattenTail(head *CNode) *CNode {

	var node = head
	var tail *CNode

	for node != nil {

		var next = node.Next

		if node.Child != nil {

			var child = node.Child
			// 获取子链条的最后一个节点
			var childTail = flattenTail(child)

			// 子链条设置为空
			node.Child = nil

			// next链条指向子链条
			node.Next = child
			child.Prev = node

			// 子链条的最后一个节点next 指向 next
			childTail.Next = next
			if next != nil {
				next.Prev = childTail
			}

			// 获取最后一个链条
			tail = childTail
		} else {
			tail = node
		}
		node = node.Next
	}

	return tail
}

// 迭代
func flatten(root *CNode) *CNode {
	var node = root

	for node != nil {

		if node.Child != nil {

			next := node.Next
			child := node.Child

			node.Next = child
			child.Prev = node
			node.Child = nil

			for child.Next != nil {
				child = child.Next
			}

			if next != nil {
				next.Prev = child
			}
			child.Next = next
		}
		node = node.Next
	}
	return root
}
