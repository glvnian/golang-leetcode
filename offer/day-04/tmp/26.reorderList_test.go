package tmp

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestReorderList(t *testing.T) {
	fmt.Println("init")
	var valList = []int{1, 2, 3, 4, 5}
	head := addValList(valList)
	reorderList(head)
	PrintHead(head)
}

func TestRevertList(t *testing.T) {
	fmt.Println("init")
	var valList = []int{1, 2, 3, 4, 5}
	head := addValList(valList)
	newHead := revertList(head)
	PrintHead(newHead)
}
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	mid := getMiddle(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = revertList(l2)
	PrintHead(head)
	mergeList(l1, l2)
}

func getMiddle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func revertList(head *ListNode) *ListNode {
	var q *ListNode
	cur := head // 1.2.3.4.5
	for cur != nil {
		tmpPtr := cur.Next // 2.3.4.5 // 2.3.4
		cur.Next = q       // nil  // 2 > 1 > nil
		q = cur            // 1 > nil  //
		cur = tmpPtr       // 2.3.4.5
	}
	return q
}

func mergeList(l1, l2 *ListNode) { // 1.2.3 |6.5.4  >> 1.6.2.5.3.4
	fmt.Println(l1.Val)
	fmt.Println(l2.Val)
	var tmp1, tmp2 *ListNode
	for l1 != nil && l2 != nil { // 1>6>2>5>4
		tmp1 = l1.Next // 2>3 | 3
		tmp2 = l2.Next // 5>4 | 4
		l1.Next = l2   // 1 > 6 | 2>5
		l1 = tmp1      // 2>3 | 3
		l2.Next = l1   // 6 > 2 | 5 >4
		l2 = tmp2      // 5>4 |

		if l1 != nil {
			fmt.Println(l1.Val)
		}
		if l2 != nil {
			fmt.Println(l2.Val)
		}
	}
}
