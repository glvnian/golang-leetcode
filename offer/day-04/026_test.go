package day_04

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	var nums1 []int
	var head1 *ListNode
	nums1 = []int{1, 2, 3, 4, 5, 6}
	nums1 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums1)
	head1 = AddValList(nums1)
	reorderList(head1)
}
