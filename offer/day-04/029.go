package day_04

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, x int) *Node {

	newNode := &Node{Val: x}
	if head == nil {
		head = newNode
		head.Next = head
		return head
	}

	if head.Next == head {
		head.Next = newNode
		newNode.Next = head
		return head
	}

	insertNode(head, newNode)
	return head
}

func insertNode(head, node *Node) {
	var cur = head
	var next = head.Next
	var biggest = head
	for !(cur.Val <= node.Val && node.Val <= next.Val) && next != head {
		fmt.Println(">>", cur.Val, node.Val, next.Val)
		cur = next
		next = next.Next

		if cur.Val >= biggest.Val {
			biggest = cur.Next
		}
	}

	if cur.Val <= node.Val && node.Val <= next.Val {
		cur.Next = node
		node.Next = next
	} else {
		node.Next = biggest.Next // biggest.Next  为最小节点
		biggest.Next = node      // biggest 是最大节点

	}
}
