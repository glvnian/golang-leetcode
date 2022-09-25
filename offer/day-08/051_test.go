package day_08

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	var data string
	deser := Constructor()
	data = "-10,9,#,#,20,15,#,#,7,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	var count int
	count = maxPathSum(node)
	fmt.Println("count: ", count)
}
