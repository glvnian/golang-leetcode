package day_01

func singleNumber(nums []int) int {
	var m = make([]int, 32)
	var res int32
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			m[i] += num >> (31 - i) & 1
		}
	}
	for i := 0; i < 32; i++ {
		res = res<<1 + int32(m[i]%3)
	}
	return int(res)
}

func singleNumberV1(nums []int) int {
	var m = make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 0
		}
		m[n] += 1
	}

	for _, n := range nums {
		if m[n] == 1 {
			return n
		}
	}
	return 0
}
