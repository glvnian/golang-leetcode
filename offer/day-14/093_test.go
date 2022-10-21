package day_14

import (
	"fmt"
	"testing"
)

func TestLenLongestFibSubseq(t *testing.T) {
	arr := []int{1, 3, 7, 11, 12, 14, 18}
	r := lenLongestFibSubseq(arr)
	fmt.Println(r)
}
