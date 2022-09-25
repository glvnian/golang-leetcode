package day_12

import (
	"fmt"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	var arr1, arr2, r []int
	arr1 = []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 = []int{2, 1, 4, 3, 9, 6}

	r = relativeSortArray(arr1, arr2)
	fmt.Println(r)
}
