package day_08

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	var data string
	deser := Constructor()
	data = "10,5,3,3,#,#,-2,#,#,2,#,1,#,#,-3,#,11,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	var count int
	var target int = 8
	count = pathSum(node, target)
	fmt.Println("count: ", count)
}
