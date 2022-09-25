package day_04

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	var node *ListNode
	var nums []int
	var n int
	//nums = []int{3, 2, 0, -4}
	nums = []int{1, 2, 3, 4, 5, 6}
	n = 3
	var head = AddCycleValList(nums, n)
	fmt.Println(head.Val)

	node = detectCycleV1(head)
	fmt.Println(node.Val)
	//PrintHead(head)
}

func AddCycleValList(nums []int, cycleVal int) *ListNode {
	var node = &ListNode{}
	var cycleNode *ListNode

	for _, val := range nums {
		if node.Val == 0 {
			node = addHead(nil, val)
			continue
		}
		node = addHead(node, val)
	}
	head := node

	for node.Next != nil && node.Next != head {
		fmt.Println("--", node.Val)
		if node.Next.Val == cycleVal {
			cycleNode = node.Next
			break
		}
		node = node.Next
	}

	for node.Next != nil {
		node = node.Next
	}

	node.Next = cycleNode
	return head
}
