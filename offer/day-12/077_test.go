package day_12

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4, 6, 1, 4}
	fmt.Println(nums)
	node := AddNode(nums)
	pNode := printNode(node)
	fmt.Println(pNode)

	rNode := sortList(node)
	rNums := printNode(rNode)
	fmt.Println(rNums)

}

func AddNode(nums []int) *ListNode {
	var node = &ListNode{Val: nums[0]}
	for _, num := range nums[1:] {
		node = addNode(node, num)
	}
	return node
}

func addNode(node *ListNode, val int) *ListNode {
	var dummy = &ListNode{}

	if node == nil {
		node = &ListNode{
			Val: val,
		}
		return node
	}

	dummy.Next = node

	for node.Next != nil {
		node = node.Next
	}
	node.Next = &ListNode{Val: val}

	return dummy.Next
}
