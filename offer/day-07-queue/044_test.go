package day_07_queue

import (
	"fmt"
	"testing"
)

func TestLargestValues(t *testing.T) {
	var root []int
	var res []int
	root = []int{1, 3, 2, 5, 3, -1, 9}
	node := createTreeNode(root)
	printTreeNode(node)
	res = largestValues(node)

	fmt.Println(res)
}

func createTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	obj := TConstructor(root)

	for _, num := range nums[1:] {
		//obj.Insert(num)
		obj.InsertNil(num)
	}

	return root
}
