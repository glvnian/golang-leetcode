package offer_02

import (
	"fmt"
	"testing"
)

func TestReversePairs(t *testing.T) {
	var nums []int
	var r int
	nums = []int{7, 5, 6, 4}
	fmt.Println(nums)
	r = reversePairs(nums)
	fmt.Println(r)

}
