package day_08

import (
	"fmt"
	"testing"
)

func TestIncreasingBST(t *testing.T) {

	var data string
	deser := Constructor()
	data = "5,1,#,#,7,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	newNode := increasingBST(node)
	data = deser.serialize(newNode)
	fmt.Println("new: ", data)
}
