package day_12

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists))
}

func mergeLists(lists []*ListNode, start, end int) *ListNode {
	if start+1 == end {
		return lists[start]
	}
	var mid = (start + end) / 2
	var head1 = mergeLists(lists, start, mid)
	var head2 = mergeLists(lists, mid, end)
	return mergeNode(head1, head2)
}

func mergeNode(head1, head2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var cur = dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}

	return dummy.Next
}

// 最小堆
