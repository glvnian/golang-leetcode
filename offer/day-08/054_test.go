package day_08

import (
	"fmt"
	"testing"
)

func TestConvertBST(t *testing.T) {
	var data string
	deser := Constructor()
	data = "2,1,#,#,3,#,#"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	node = convertBST(node)
	data = deser.serialize(node)
	fmt.Println("convertBST: ", data)
}
