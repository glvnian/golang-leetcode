package sort

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int) {
	var start, end int
	start = 0
	end = len(nums) - 1
	QuickSort(nums, start, end)
}
func QuickSort(nums []int, start, end int) {
	if end > start {
		fmt.Println(start, end, nums)
		pivot := partition(nums, start, end)
		fmt.Println(pivot, nums)
		QuickSort(nums, start, pivot-1)
		QuickSort(nums, pivot+1, end)
		fmt.Println(">>", start, end, nums)
	}
}

func partition(nums []int, start, end int) int {
	// 选择分隔点
	pivot := rand.Intn(end-start+1) + start

	// 移动分割点到列表最后
	Swap(nums, pivot, end)

	// 指针1
	p1 := start - 1

	for p2 := start; p2 <= end; p2++ {
		if nums[p2] < nums[end] {
			p1++
			Swap(nums, p1, p2)
		}
	}

	p1++
	Swap(nums, p1, end)

	return p1
}

func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
