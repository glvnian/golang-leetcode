package day_04

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddValList(nums []int) *ListNode {
	var node = &ListNode{}
	for _, val := range nums {
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
	for head != nil {
		fmt.Printf("%d", head.Val)
		//fmt.Printf("%d \n", head.Val)
		head = head.Next
		//time.Sleep(time.Second)
	}
	fmt.Println()
	fmt.Println("---------- PrintHead end ------------")
}
