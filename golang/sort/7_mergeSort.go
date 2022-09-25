package sort

// 迭代
func mergeSortV1(nums []int) {

	for seg := 1; seg < len(nums); seg += seg {

		for start := 0; start < len(nums); start += seg * 2 {

			var mid = Min(start+seg, len(nums))
			var end = Min(start+2*seg, len(nums))
			var i = start
			var j = mid
			var k = start

			for i < mid || j < end {
				if j == end || (i < mid && nums[i] < nums[j]) {
					nums[k] = nums[i]
					i++
				} else {
					nums[k] = nums[j]
					j++
				}
				k++
			}

		}

	}

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 递归
func mergeSort(nums []int) {
	var dst = make([]int, 0)
	dst = append(dst, nums...)
	_mergeSort(nums, dst, 0, len(nums))
}

func _mergeSort(src, dst []int, start, end int) {
	if start+1 >= end {
		return
	}
	var mid = (start + end) / 2
	_mergeSort(dst, src, start, mid)
	_mergeSort(dst, src, mid, end)

	var i = start
	var j = mid
	var k = start

	for i < mid || j < end {
		if j == end || (i < mid && src[i] < src[j]) {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
		k++
	}

}
