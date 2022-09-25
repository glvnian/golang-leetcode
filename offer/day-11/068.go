package day_11

func searchInsert(nums []int, target int) int {

	var start = 0
	var end = len(nums) - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return len(nums)
}
