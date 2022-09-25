package sort

// 冒泡排序
//冒泡排序是入门级的算法，但也有一些有趣的玩法。通常来说，冒泡排序有三种写法：
//一边比较一边向后两两交换，将最大值 / 最小值冒泡到最后一位；
//经过优化的写法：使用一个变量记录当前轮次的比较是否发生过交换，如果没有发生交换表示已经有序，不再继续排序；
//进一步优化的写法：除了使用变量记录当前轮次是否发生交换外，再使用一个变量记录上次发生交换的位置，下一轮排序时到达上次交换的位置就停止比较。

func bubbleSortV1(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// 循环用于找到第i大数据
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func bubbleSortV2(nums []int) {
	// 初始时 swapped 为 true，否则排序过程无法启动
	var tag = true
	for i := 0; i < len(nums)-1; i++ {
		// 如果没有发生过交换，说明剩余部分已经有序，排序完成
		if !tag {
			break
		}

		// 设置 swapped 为 false，如果发生交换，则将其置为 true
		tag = false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				// 表示发生了交换
				tag = true
			}
		}
	}
}

func bubbleSortV3(nums []int) {
	var swapped = true
	// 最后一个没有经过排序的元素的下标
	var lastIndex = len(nums) - 1
	// 上次发生交换的位置
	var swappedIndex = -1

	for swapped {
		swapped = false
		for i := 0; i < lastIndex; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
				// 表示发生了交换
				swapped = true
				swappedIndex = i
			}
		}
		// 最后一个没有经过排序的元素的下标就是最后一次发生交换的位置
		lastIndex = swappedIndex

	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
