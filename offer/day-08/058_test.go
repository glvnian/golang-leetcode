package day_08

import (
	"fmt"
	"testing"
)

func TestBookE1(t *testing.T) {
	var start, end int
	var b bool
	obj := MConstructor()
	start, end = 10, 20
	b = obj.Book(start, end)
	fmt.Println(b)

	start, end = 15, 30
	b = obj.Book(start, end)
	fmt.Println(b)

	start, end = 25, 30
	b = obj.Book(start, end)
	fmt.Println(b)
}
