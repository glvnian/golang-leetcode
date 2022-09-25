package day_08

import (
	"fmt"
	"testing"
)

func TestFindTarget(t *testing.T) {
	var data string
	deser := Constructor()
	data = "8,6,5,#,#,7,#,#,10,9,#,#,11,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	var bo bool
	var target int
	target = 12
	bo = findTarget(node, target)
	fmt.Println(bo)
}
