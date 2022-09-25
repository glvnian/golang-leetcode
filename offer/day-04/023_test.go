package day_04

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	var nums1, nums2 []int
	var node *ListNode
	//nums = []int{3, 2, 0, -4}
	nums1 = []int{4, 1, 8, 4, 5}
	nums2 = []int{5, 0, 1, 8, 4, 5}
	var head1 = AddValList(nums1)
	var head2 = AddValList(nums2)
	fmt.Println(head1.Val)
	fmt.Println(head2.Val)

	//
	node = getIntersectionNode(head1, head2)
	fmt.Println(node)
}
