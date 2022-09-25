package day_07_queue

import (
	"fmt"
	"testing"
)

func TestCBTInserter(t *testing.T) {
	var root *TreeNode
	var r, v int

	root = &TreeNode{
		Val: 0,
	}
	obj := TConstructor(root)

	v = 1
	r = obj.Insert(v)
	fmt.Println(r)

	v = 2
	r = obj.Insert(v)
	fmt.Println(r)

	v = 3
	r = obj.Insert(v)
	fmt.Println(r)

	v = 4
	r = obj.Insert(v)
	fmt.Println(r)
}
