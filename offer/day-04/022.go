package day_04

import "fmt"

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	inLoop := getNodeInLoop(dummy)
	if inLoop == nil {
		return nil
	}
	fmt.Println("inLoop -->", inLoop.Val)
	fmt.Println("head -->", head.Val)

	var newNode = dummy
	for newNode != inLoop {
		newNode = newNode.Next
		inLoop = inLoop.Next
	}

	return newNode
}

func detectCycleV1(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	inLoop := getNodeInLoop(dummy)
	fmt.Println("getNodeInLoop", inLoop.Val)
	if inLoop == nil {
		return nil
	}
	// 判断环的大小
	var loopSize int = 1
	node := inLoop.Next
	for node != inLoop {
		node = node.Next
		loopSize++
	}
	fmt.Println("loopSize: ", loopSize)

	// 快忙指针查找入口
	slow := head
	fast := head
	for i := 0; i < loopSize; i++ {
		fast = fast.Next
	}

	for fast != slow {
		slow = slow.Next
		fast = fast.Next
		fmt.Println(">>")
	}

	return slow
}

// 判断链表是否为环
func getNodeInLoop(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
