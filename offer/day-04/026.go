package day_04

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	node1, node2 := splitList(head)
	node2 = reorderReverseList(node2)

	var prev = head
	for node1 != nil && node2 != nil { // 123  654
		var tmp = node1.Next // 23
		prev.Next = node1    // 1 ||
		node1.Next = node2   // 1 -> 6 ||
		prev = node2

		node1 = tmp        // 2
		node2 = node2.Next // 5
	}

	if node1 != nil {
		prev.Next = node1
	}
}

// 分为两部分
func splitList(head *ListNode) (head1, head2 *ListNode) {
	// 快慢指针分割
	var dummy = &ListNode{}
	dummy.Next = head

	var slow, fast *ListNode
	slow = dummy
	fast = dummy.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	head1 = head
	head2 = slow.Next
	slow.Next = nil
	return head1, head2
}

// 反转
func reorderReverseList(head *ListNode) *ListNode {
	var prev, cur, tmp *ListNode
	cur = head
	for cur != nil {
		tmp = cur.Next  // k  // cur = j cur.next = k  prev = i
		cur.Next = prev // i  // j > i

		prev = cur
		cur = tmp
	}
	return prev
}
