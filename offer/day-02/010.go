package day_02

func subarraySum(nums []int, k int) int {
	var count, sum int
	var m = make(map[int]int)
	// 初始化，默认情况一个数字都不添加，和为0的数量为1个
	m[0] = 1
	for _, num := range nums {
		sum += num

		if _, ok := m[sum-k]; ok {
			count += m[sum-k]
		}

		if _, ok := m[sum]; ok {
			m[sum] += 1
		} else {
			m[sum] = 1
		}

	}
	return count
}
