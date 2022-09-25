package day_04

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var nums1, nums2 []int
	var head1, head2, res *ListNode
	nums1 = []int{1, 2, 3, 4, 5, 6}
	nums2 = []int{1, 2, 3, 4}
	head1 = AddValList(nums1)
	head2 = AddValList(nums2)
	PrintHead(head1)
	PrintHead(head2)
	res = addTwoNumbers(head1, head2)
	_ = res
	PrintHead(res)
	fmt.Println(654321 + 4321)
}

func TestAdd(t *testing.T) {

}
