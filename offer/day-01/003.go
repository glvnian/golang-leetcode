package day_01

func countBits(n int) []int {
	var res = make([]int, n+1)
	for i := 1; i <= n; i++ {
		//res[i] = res[i>>1] + (i & 1)
		res[i] = res[i/2] + (i & 1)
		// i>>1 = i/2
		//  (i & 1) 判断最后一位是否为1
	}
	return res
}

func countBitsV2(n int) []int {
	var res = make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func countBitsV1(n int) []int {
	var res = make([]int, n+1)
	for i := 0; i <= n; i++ {
		var tmp int
		for j := i; j > 0; {
			j = j & (j - 1)
			tmp++
		}
		res[i] = tmp
	}
	return res
}
