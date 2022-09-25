package day_04

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	var n int
	var node *Node
	n = 1
	node = insert(nil, n)
	PrintNode(node)
	n = 3
	node = insert(node, n)
	PrintNode(node)
	n = 5
	node = insert(node, n)
	PrintNode(node)
	n = 0
	fmt.Println("add - : ", n)
	node = insert(node, n)
	PrintNode(node)
}

func PrintNode(head *Node) {
	node := head.Next
	for node != head {
		fmt.Printf("%d \n", node.Val)
		//fmt.Printf("%d \n", head.Val)
		node = node.Next
		//time.Sleep(time.Second)
	}
	fmt.Println(">n", node.Val)
	fmt.Println("---------- PrintHead end ------------")
}
