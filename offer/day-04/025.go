package day_04

// 还可以用栈

//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r1, r2 *ListNode

	r1 = NumReverseList(l1)
	r2 = NumReverseList(l2)
	var carry int
	var prev = &ListNode{}
	res := prev
	for r1 != nil || r2 != nil {
		var num1, num2, sum int
		if r1 != nil {
			num1 = r1.Val
			r1 = r1.Next
		}
		if r2 != nil {
			num2 = r2.Val
			r2 = r2.Next
		}

		sum = num1 + num2 + carry
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		newNode := &ListNode{Val: sum}
		prev.Next = newNode
		prev = prev.Next
	}

	if carry == 1 {
		newNode := &ListNode{Val: carry}
		prev.Next = newNode
	}
	return NumReverseList(res.Next)
}

func NumReverseList(l1 *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = l1
	for cur != nil {
		next := cur.Next // cur.next = k
		cur.Next = prev  // cur = j ; prev = i //
		prev = cur       // j
		cur = next       // k
	}
	return prev
}
