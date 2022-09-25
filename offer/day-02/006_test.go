package day_02

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 4, 6, 10}
	target := 8
	r := twoSum(numbers, target)
	fmt.Println(r)

	numbers = []int{1, 2, 3, 4, 6, 10}
	r = twoSum(numbers, target)
	fmt.Println(r)
}
