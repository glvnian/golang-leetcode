package day_07_queue

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T) {

	var root []int
	var res []int
	root = []int{1, 2, 3, -1, 5, -1, 4}
	root = []int{1, -1, 4, -1, 3}
	root = []int{1, 2}
	node := createTreeNode(root)
	printTreeNode(node)
	res = rightSideView(node)

	fmt.Println(res)
}
