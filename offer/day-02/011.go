package day_02

//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
// 由于0 和 1 的数量相同 等价于 1 的数量减去 0 的数量等于 0，
// 我们可以将数组中的 0 视作 −1，则原问题转换成「求最长的连续子数组，其元素和为 0。
func findMaxLength(nums []int) int {
	var k = 0
	var sum, max int

	// m 的key和0到i的和，value为当前的坐标i
	var m = make(map[int]int)
	// 初始化，默认情况一个数字都不添加，和为0的当前坐标为-1
	m[0] = -1
	for i, num := range nums {
		// 把0替换为-1，那么数组中就只存在1，和-1。他们直接的和为0
		if num == 0 {
			num = -1
		}

		sum += num
		if _, ok := m[sum-k]; ok {
			max = Max(max, i-m[sum-k])
			//fmt.Println("max ", max, i-m[sum-k])
		}

		if _, ok := m[sum]; !ok {
			m[sum] = i
		}

	}
	return max
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
