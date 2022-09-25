package day_04

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var nums1 []int
	var head *ListNode
	var b bool
	nums1 = []int{1, 2, 3, 4, 3, 2, 1}
	nums1 = []int{1, 2, 3, 1}
	fmt.Println(nums1)
	head = AddValList(nums1)
	b = isPalindrome(head)
	fmt.Println(b)
}

func TestIsPalindromeReverseList(t *testing.T) {
	nums1 := []int{1, 2, 3, 3, 2, 1}
	fmt.Println(nums1)
	head := AddValList(nums1)
	PrintHead(head)
	rhead := isPalindromeReverseList(head)
	PrintHead(rhead)
}
