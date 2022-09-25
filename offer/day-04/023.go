package day_04

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}

	}

	return pa
}

func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	var countA, countB, delte int
	var longest, shorter *ListNode
	// headA 首尾相接
	countA = countListNode(headA)
	countB = countListNode(headB)
	fmt.Println("size: ", countA, countB)
	if countA > countB {
		delte = countA - countB
		longest, shorter = headA, headB

	} else {
		delte = countB - countA
		longest, shorter = headB, headA
	}
	node1 := longest
	node2 := shorter

	for i := 0; i < delte; i++ {
		node1 = node1.Next
		fmt.Println(">>", i)
	}
	fmt.Println("size: node1 ", node1.Val)

	for node1 != node2 {
		fmt.Println("-node1:", node1.Val)
		fmt.Println("node2:", node2.Val)
		node1 = node1.Next
		node2 = node2.Next
	}
	fmt.Println(node1)
	return node1
}

func countListNode(head *ListNode) int {
	var count int
	node := head
	for node != nil {
		node = node.Next
		count++
	}
	return count
}
