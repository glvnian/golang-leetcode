package day_04

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseListV1(head *ListNode) *ListNode {
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
