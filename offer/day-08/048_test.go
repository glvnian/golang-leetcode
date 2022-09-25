package day_08

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeSerialize(t *testing.T) {
	var data string
	//ser := Constructor()
	//data := ser.serialize(root)
	//
	deser := Constructor()
	data = "6,6,6,#,#,6,#,#,6,#,#"
	data = "1,2,3,#,#,4,5"
	data = ""
	node := deser.deserialize(data)
	fmt.Println("ori: ", data)
	ser := Constructor()
	data = ser.serialize(node)
	fmt.Println(data)
}

func TestName(t *testing.T) {
	a := []string{}
	fmt.Println(strings.Join(a, ","))
}
