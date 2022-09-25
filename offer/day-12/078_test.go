package day_12

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var lists = [][]int{
		{1, 4, 5}, {1, 3, 4}, {2, 6},
	}
	var nodeLists []*ListNode
	for _, nums := range lists {
		var node *ListNode
		node = AddNode(nums)
		nodeLists = append(nodeLists, node)
		fmt.Println(printNode(node))

	}
	//fmt.Println(nodeLists)

	node := mergeKLists(nodeLists)
	rNode := printNode(node)
	fmt.Println(rNode)

}
