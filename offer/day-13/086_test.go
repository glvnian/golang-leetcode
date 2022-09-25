package day_13

import (
	"fmt"
	"testing"
)

func TestSplitString(t *testing.T) {
	a := "cbbbcc"
	r := partition(a)
	fmt.Println(r)
	for _, item := range r {
		fmt.Println(">>", item)
	}
}

func TestPartition(t *testing.T) {
	a := "1233421"
	fmt.Println(isPartition(a, 0, 0))
}
