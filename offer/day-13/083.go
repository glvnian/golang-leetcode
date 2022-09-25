package day_13

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	for i := range nums {
		fmt.Println(i)
		helperPermute(nums, i, &result)
	}
	return result
}

func helperPermute(nums []int, i int, result *[][]int) {
	if i == len(nums) {
		var p = make([]int, len(nums))
		copy(p, nums)
		*result = append(*result, p)
	} else {
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			helperPermute(nums, i+1, result)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func permuteV2(nums []int) [][]int {
	ans, cur, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var dfs func(c int)
	dfs = func(c int) {
		if c == len(nums) {
			ans = append(ans, append([]int(nil), cur...))
			return
		}
		for i, isUsed := range used {
			if isUsed {
				continue
			}
			cur = append(cur, nums[i])
			used[i] = true
			dfs(c + 1)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	dfs(0)
	return ans
}
