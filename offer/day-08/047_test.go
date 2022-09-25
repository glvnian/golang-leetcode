package day_08

import (
	"fmt"
	"testing"
)

//func TestTreeSerialize(t *testing.T) {
//	var data string
//	//ser := Constructor()
//	//data := ser.serialize(root)
//	//
//	deser := Constructor()
//	data = "6,6,6,#,#,6,#,#,6,#,#"
//	data = "1,2,3,#,#,4,5"
//	data = ""
//	node := deser.deserialize(data)
//	fmt.Println("ori: ", data)
//	ser := Constructor()
//	data = ser.serialize(node)
//	fmt.Println(data)
//}

func TestPruneTree(t *testing.T) {
	var data string
	var node *TreeNode
	deser := Constructor()
	data = "0,0,0,#,#,0,#,#,0,0,#,#,1,#,#"
	node = deser.deserialize(data)
	fmt.Println("ori: ", data)
	node = pruneTree(node)
	data = deser.serialize(node)
	fmt.Println(data)
}
