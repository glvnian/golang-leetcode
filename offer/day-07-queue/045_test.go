package day_07_queue

import (
	"fmt"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	var root []int
	var res int
	root = []int{1, 3, 2, 5, 3, -1, 9}
	root = []int{2, 1, 3}
	root = []int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7}
	node := createTreeNode(root)
	printTreeNode(node)
	res = findBottomLeftValue(node)
	fmt.Println(res)

}
