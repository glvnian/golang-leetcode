package sort

import "fmt"

func heapSort(nums []int) {
	initHeap(nums)
	fmt.Println("init maxHeap: ", nums)
	for i := len(nums) - 1; i > 0; i-- {
		// 将最大值交换到数组最后
		swap(nums, 0, i)
		// 调整剩余数组，使其满足大顶堆
		maxHeap(nums, 0, i)
	}
}

// 初始化最大堆
func initHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		fmt.Println("> ", i, nums)
		maxHeap(nums, i, len(nums))
	}
}

func maxHeap(nums []int, start, heapSize int) {
	fmt.Println(start, nums)
	// 左子结点下标
	var l = 2*start + 1
	// 右子结点下标
	var r = l + 1
	// 记录根结点、左子树结点、右子树结点三者中的最大值下标
	var largest = start
	// 与左子树结点比较
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	// 与右子树结点比较
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}

	if largest != start {
		fmt.Println(">>> - swap", start, largest, nums)
		// 将最大值交换为根结点
		swap(nums, start, largest)
		fmt.Println(">>> swap", start, largest, nums)

		// 再次调整交换数字后的大顶堆
		maxHeap(nums, largest, heapSize)
		fmt.Println("swap", nums)
	}
}
