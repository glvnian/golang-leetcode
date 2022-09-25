package day_13

import "fmt"

func subsets(nums []int) [][]int {
	var subset = []int{}
	var result = make([][]int, 0)
	//helper(nums, 0, subset, &result)
	helperV2(nums, 0, subset, &result)
	fmt.Println(len(result))
	return result
}

func helperV2(nums []int, start int, subset []int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, subset...))
	}
	if start < len(nums) {
		// 选
		subset = append(subset, nums[start])
		helper(nums, start+1, subset, result)
		//// 不选
		subset = subset[:len(subset)-1]
		helper(nums, start+1, subset, result)
	}
}

func helper(nums []int, start int, subset []int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, subset...))
		//*result = append(*result, subset)
	} else {
		// 不选
		//[[],[3],[2],[2,3],[1],[1,3],[1,2],[1,2,3]]
		//[[],[3],[2],[2,3],[1],[1,3],[1,2],[1,2,3]]
		helper(nums, start+1, subset, result)
		// 选
		subset = append(subset, nums[start])
		helper(nums, start+1, subset, result)
		l := len(subset)
		subset = subset[:l-1]
	}
}

func subsetsV1(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

//
//func subsets(nums []int) [][]int {
//	var subset = []int{}
//	var result = make([][]int, 0)
//	var dfs func(start int)
//	dfs = func(start int) {
//		if start == len(nums) {
//			result = append(result, append([]int{}, subset...))
//			return
//		}
//		dfs(start + 1)
//		subset = append(subset, nums[start])
//		dfs(start + 1)
//		subset = subset[:len(subset)-1]
//	}
//	dfs(0)
//	return result
//}
