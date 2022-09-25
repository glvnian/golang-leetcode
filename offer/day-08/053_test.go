package day_08

import (
	"fmt"
	"testing"
)

func TestInorderSuccessor(t *testing.T) {
	var data string
	deser := Constructor()
	data = "2,1,#,#,3,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	p := &TreeNode{Val: 1}
	newNode := inorderSuccessor(node, p)
	if newNode != nil {
		fmt.Println(newNode.Val)
	} else {
		fmt.Println("not found")
	}

}
