package day_02

func pivotIndex(nums []int) int {

	var total, sum int
	for _, num := range nums {
		total += num
	}

	for i, num := range nums {
		sum += num
		if sum-num == total-sum {
			return i
		}
	}
	return -1
}
