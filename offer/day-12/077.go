package day_12

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var head1 = head
	var head2 = split(head)

	head1 = sortList(head1)
	head2 = sortList(head2)

	return mergeListNode(head1, head2)
}

func mergeListNode(head1, head2 *ListNode) *ListNode {
	var node = &ListNode{}
	var cur = node

	for head1 != nil && head2 != nil {
		if head2 == nil || head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	if head2 == nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}

	return node.Next
}

func split(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	return second

}

func printNode(node *ListNode) []int {
	var r = make([]int, 0)
	for node != nil {
		r = append(r, node.Val)
		node = node.Next
	}
	return r
}
