package day_08

import (
	"fmt"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	var nums = []int{1, 5, 9, 1, 5, 9}
	nums = []int{-3, 3, -6}
	var k = 2
	var t1 = 3
	fmt.Println("k: ", k, "t:", t1)
	b := containsNearbyAlmostDuplicate(nums, k, t1)
	fmt.Println(nums)
	fmt.Println(b)
}

//[1,5,9,1,5,9]
//2
//3
