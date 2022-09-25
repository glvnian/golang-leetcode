package day_13

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var subnet = make([]int, 0)
	var result = make([][]int, 0)
	combinationSumHelper(candidates, target, 0, subnet, &result)
	return result
}

func combinationSumHelper(nums []int, target int, start int, subnet []int, result *[][]int) {
	if len(nums) == start {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, subnet...))
		return
	}
	if target < 0 {
		return
	}
	fmt.Println(len(nums), start)
	// 忽略
	combinationSumHelper(nums, target, start+1, subnet, result)
	// 处理
	subnet = append(subnet, nums[start])
	combinationSumHelper(nums, target-nums[start], start, subnet, result)
}

func combinationSumV2(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
