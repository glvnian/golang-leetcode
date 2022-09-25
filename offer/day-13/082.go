package day_13

import (
	"fmt"
	"sort"
)

func combinationSumTwo(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	subset := make([]int, 0)
	startIndex := 0
	sort.Ints(candidates)
	fmt.Println(candidates)
	HelperCombinationSumTwo(candidates, target, startIndex, subset, &result)
	return result
}

func HelperCombinationSumTwo(candidates []int, target, startIndex int, subset []int, result *[][]int) {
	if target == 0 {
		desc := make([]int, len(subset))
		copy(desc, subset)
		*result = append(*result, desc)
	}

	if target > 0 && startIndex < len(candidates) {
		// 不选择 startIndex
		HelperCombinationSumTwo(candidates, target, getNextIndex(candidates, startIndex), subset, result)
		// 选择 startIndex
		subset = append(subset, candidates[startIndex])
		HelperCombinationSumTwo(candidates, target-candidates[startIndex], startIndex+1, subset, result)

	}
}

func getNextIndex(nums []int, index int) int {
	next := index
	for next < len(nums) && nums[next] == nums[index] {
		next++
	}
	return next
}
