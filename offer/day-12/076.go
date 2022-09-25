package day_12

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargestV1(nums []int, k int) int {
	fmt.Println(nums)
	var pivot, target, start, end int
	start = 0
	end = len(nums) - 1
	target = len(nums) - k

	pivot = partition(nums, start, end)
	fmt.Println(target, pivot, nums[pivot], start, end, nums)
	for pivot != target {
		if pivot > target {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
		fmt.Println(pivot, nums[pivot], start, end, nums)
		pivot = partition(nums, start, end)
		fmt.Println("---", pivot, nums[pivot], start, end, nums)
	}
	return nums[pivot]
}

func partition(nums []int, start, end int) int {
	var pivot, small, i int
	rand.Seed(time.Now().Unix())
	pivot = (rand.Intn(end-start+1) + start)
	fmt.Println(">>", end, start, end-start, pivot)
	swap(nums, pivot, end)

	small = start - 1
	for i = start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			swap(nums, i, small)
		}
	}

	small++
	swap(nums, small, end)
	return small

}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func findKthLargest(nums []int, k int) int {
	var heapSize = len(nums)
	buildHeap(nums, heapSize)

	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 最大堆
func buildHeap(nums []int, heapSize int) {
	for start := heapSize / 2; start >= 0; start-- {
		maxHeapify(nums, start, heapSize)
	}
}

func maxHeapify(nums []int, start, heapSize int) {
	largest, left, right := start, 2*start+1, 2*start+2

	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}

	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if largest != start {
		nums[start], nums[largest] = nums[largest], nums[start]
		maxHeapify(nums, largest, heapSize)
	}
}
