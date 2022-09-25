package offer_02

func reversePairs(nums []int) int {

	var p1, count int
	for p1 = 0; p1 < len(nums)-1; p1++ {
		if nums[p1] > nums[p1+1] {
			count++
		}
	}
	return count
}

// 暴力
func reversePairsV1(nums []int) int {

	var p1, p2, count int
	for p1 = 0; p1 < len(nums); p1++ {
		for p2 = p1; p2 < len(nums); p2++ {
			if nums[p1] > nums[p2] {
				count++
			}
		}
	}
	return count
}
