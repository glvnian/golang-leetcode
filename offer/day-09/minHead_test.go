package day_09

import (
	"fmt"
	"testing"
)

func TestPQ_Len(t *testing.T) {
	var p = PQ{1, 2, 3, 4}
	fmt.Println(p.Len())
	fmt.Println(p)
	p.Swap(0, 1)
	fmt.Println("swap", p)
	p.Push(5)
	fmt.Println("push", p)
	p.Pop()
	fmt.Println("pop", p)
	fmt.Println(p.Len())
}
