package day_09

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var nums = []int{1, 1, 1, 2, 2, 3}
	var k = 2
	r := topKFrequentV1(nums, k)
	fmt.Println(r)
}
