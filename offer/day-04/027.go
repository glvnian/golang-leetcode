package day_04

import "fmt"

// 递归
func isPalindrome(head *ListNode) bool { // 1 2 1
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool

	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			fmt.Println("-->", curNode.Val, frontPointer.Val)
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			fmt.Println(curNode.Val, frontPointer.Val)
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

// 迭代
func isPalindromeV1(head *ListNode) bool {
	var node1, node2 *ListNode
	node1, node2 = splitListNode(head)
	node2 = isPalindromeReverseList(node2)
	var tag bool
	for node1 != nil || node2 != nil {
		if node2 != nil && node1.Val != node2.Val {
			return false
		}

		if node2 == nil && node1 != nil && !tag {
			tag = true
			node1 = node1.Next
			continue
		}

		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}

// 分割
func splitListNode(head *ListNode) (node1, node2 *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node1 = head
	node2 = slow.Next
	slow.Next = nil
	return node1, node2
}

// 反转
func isPalindromeReverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = head
	for cur != nil {
		next := cur.Next // cur.next = k ,cur = j , prev = i
		cur.Next = prev  // j -> i

		prev = cur //
		cur = next
	}

	return prev
}
