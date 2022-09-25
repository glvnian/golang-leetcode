package day_11

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	var min, max, mid int
	min = 1
	for _, item := range piles {
		max = Max(max, item)
	}

	for min <= max {
		mid = (min + max) / 2

		fmt.Println(min, max, mid, ">>", eatHours(piles, mid), eatHours(piles, mid+1))

		if eatHours(piles, mid) <= h {
			//吃快了
			if eatHours(piles, mid-1) > h {
				return mid
			}

			max = mid - 1
		} else {
			min = mid + 1
		}

	}

	return mid
}

func eatHours(piles []int, speed int) int {
	var sum int
	for _, item := range piles {
		sum += (item + speed - 1) / speed
	}
	return sum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 8 2 = 4  || 8 + 2 -1 / 4
// 9 2 = 5  || 9 + 2 -1 / 5
// 7 2 = 3   || 7 +2 -1 / 4
//
