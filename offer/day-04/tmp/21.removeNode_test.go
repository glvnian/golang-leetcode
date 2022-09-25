package tmp

import (
	"fmt"
	"testing"
)

func TestRemoveNthFormEnd(t *testing.T) {
	fmt.Println("init")
	var valList = []int{1}
	head := addValList(valList)
	removeNthFormEnd(head, 1)
	//removeNthFormEndV2(head, 2)
}

func removeNthFormEnd(head *ListNode, n int) *ListNode {
	dummp := &ListNode{Val: -1}
	dummp.Next = head

	slow := dummp
	fast := head

	for i := 0; i < n-1; i++ {
		if fast.Next == nil {
			return head
		}
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		fmt.Println("slow: ", slow.Val)
		fmt.Println("fast: ", fast.Val)
	}

	slow.Next = slow.Next.Next
	newNode := dummp.Next
	for newNode != nil {
		fmt.Println("newNode: ", newNode.Val)
		newNode = newNode.Next
	}

	return dummp.Next
}

func removeNthFormEndV2(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummp := &ListNode{Val: -1}
	dummp.Next = head

	slow := dummp
	fast := head

	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		fmt.Println("slow: ", slow.Val)
		fmt.Println("fast: ", fast.Val)
	}

	slow.Next = slow.Next.Next
	newNode := dummp.Next
	for newNode != nil {
		fmt.Println("newNode: ", newNode.Val)
		newNode = newNode.Next
	}

	return dummp.Next
}
