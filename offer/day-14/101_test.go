package day_14

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	r := canPartition(nums)
	fmt.Println(r)
}
