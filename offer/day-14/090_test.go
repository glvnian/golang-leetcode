package day_14

import (
	"fmt"
	"testing"
)

func TestRingRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	r := ring(nums)
	fmt.Println(r)
}
