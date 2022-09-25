package day_01

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{0, 1, 0, 1, 0, 1, 100}
	a := singleNumber(nums)
	fmt.Println(nums)
	fmt.Println(a)
	nums = []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
	a = singleNumber(nums)
	fmt.Println(nums)
	fmt.Println(a)
}

func TestNum(t *testing.T) {
	num := -2
	var a [32]int
	for i := 0; i < 32; i++ {
		a[i] = num >> (31 - i) & 1
	}
	fmt.Println(a)
}
