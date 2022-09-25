package day_11

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	fmt.Println(arr)
	var start = 0
	// 这个是重点
	var end = len(arr) - 2
	var mid int
	for start <= end {
		mid = (start + end) / 2
		fmt.Println(start, end, mid)

		if arr[mid] > arr[mid+1] {
			if mid == 0 || arr[mid] > arr[mid-1] {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return len(arr)
}
