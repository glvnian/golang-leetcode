package day_04

import (
	"testing"
)

//head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]

func TestRemoveNthFromEnd(t *testing.T) {
	var res *ListNode
	var nums []int
	var n int
	nums = []int{1, 2, 3, 4, 5}
	n = 1
	var head = AddValList(nums)
	PrintHead(head)

	res = removeNthFromEnd(head, n)

	PrintHead(res)
}
