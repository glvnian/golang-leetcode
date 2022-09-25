package day_13

func combine(n int, k int) [][]int {
	subnet := make([]int, 0)
	result := make([][]int, 0)

	combineHelper(n, k, 1, subnet, &result)
	return result
}

func combineHelper(n, k int, start int, subnet []int, result *[][]int) {
	if len(subnet) == k {
		*result = append(*result, append([]int{}, subnet...))
	}
	if n >= start {
		// 忽略
		combineHelper(n, k, start+1, subnet, result)
		// 添加
		subnet = append(subnet, start)
		combineHelper(n, k, start+1, subnet, result)
		subnet = subnet[:len(subnet)-1]
	}
}

func combineV2(n int, k int) [][]int {
	subnet := make([]int, 0)
	result := make([][]int, 0)

	var dfs func(start int)

	dfs = func(start int) {
		if n < start-1 {
			return
		}
		if len(subnet) == k {
			result = append(result, append([]int{}, subnet...))
			return
		}

		// 忽视
		dfs(start + 1)

		//新增
		subnet = append(subnet, start)
		dfs(start + 1)
		subnet = subnet[:len(subnet)-1]
	}
	dfs(1)
	return result
}
