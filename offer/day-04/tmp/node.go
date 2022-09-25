package tmp

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var node = &ListNode{}

func addValList(nodeValList []int) *ListNode {
	for _, val := range nodeValList {
		if node.Val == 0 {
			node = addHead(nil, val)
			continue
		}
		node = addHead(node, val)
	}
	return node
}

func addHead(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{0, nil}
	dummy.Next = head

	node := dummy
	var newNode = &ListNode{val, nil}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode

	return dummy.Next
}

func PrintHead(head *ListNode) {
	fmt.Println("---------- PrintHead ------------")

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println("---------- PrintHead end------------")

}
