package day_11

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	var start, end, mid int
	start = 1
	end = x
	for start <= end {
		mid = (start + end) / 2
		fmt.Println(start, end, mid)
		if mid <= x/mid {
			if mid == 0 || (mid+1) > x/(mid+1) {
				return mid
			}

			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return mid
}
