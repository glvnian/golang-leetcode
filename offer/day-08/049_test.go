package day_08

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	var data string
	deser := Constructor()
	data = "1,2,#,#,3"
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	data = deser.serialize(node)
	fmt.Println("new: ", data)

	var path int
	path = sumNumbers(node)
	fmt.Println("path: ", path)
}
