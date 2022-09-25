package day_09

import (
	"fmt"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	var num1, num2 []int
	var k int
	num1 = []int{1, 7, 11}
	num2 = []int{2, 4, 6}
	k = 3
	//num1 = []int{1, 1, 2}
	//num2 = []int{1, 2, 3}
	//k = 2
	//
	//num1 = []int{1, 1, 2}
	//num2 = []int{1, 2, 3}
	//k = 2
	//num1 = []int{-10, -4, 0, 0, 6}
	//num2 = []int{3, 5, 6, 7, 8, 100}
	//k = 10
	r := kSmallestPairsV1(num1, num2, k)
	fmt.Println(r)
}
