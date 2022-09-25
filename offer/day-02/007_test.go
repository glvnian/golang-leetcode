package day_02

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	r := threeSum(nums)
	fmt.Println(r)

	nums = []int{0, 0, 0, 0, 0}
	r = threeSum(nums)
	fmt.Println(r)
}

func TestSort(t *testing.T) {
	// []int排序
	nums := []int{2, 31, 5, 6, 3}
	//顺序
	sort.Ints(nums)
	fmt.Println("1: ", nums)
	//使用 sort.Reverse 进行逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("2: ", nums)

	// []string字符串排序
	names := []string{"abc", "12", "kk", "Jordan", "Ko", "DD"}
	// 顺序
	sort.Strings(names)
	fmt.Println("3: ", names)
	// 逆序
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Printf("4: After Reverse: %#v\n", names)

	//查找
	// [0,100)
	// 二分查找
	nums = []int{1, 3, 5, 7, 9}
	fmt.Println("5: ", nums[sort.SearchInts(nums, 8)] == 8)
	fmt.Println("6: ", nums[sort.SearchInts(nums, 5)] == 5)
}
