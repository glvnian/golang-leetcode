package day_04

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var nums []int
	var head, res *ListNode
	nums = []int{5, 0, 1, 8, 4, 5}
	head = AddValList(nums)
	PrintHead(head)
	res = reverseList(head)
	_ = res
	PrintHead(res)
}
