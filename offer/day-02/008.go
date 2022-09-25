package day_02

import (
	"math"
	"sort"
)

func minSubArrayLenV1(target int, nums []int) int {
	var min = float64(math.MaxInt32)
	var i, j int
	var sum int
	for j = 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= target {
			min = math.Min(min, float64(j-i+1))
			sum -= nums[i]
			i++
		}
	}
	if min == float64(math.MaxInt32) {
		min = 0
	}
	return int(min)
}

// 前缀和 + 二分查找
// 为了使用二分查找，需要额外创建一个数组 sums 用于存储数组 nums 的前缀和，
// 其中 sums[i] 表示从 nums[0] 到 nums[i−1] 的元素和。
// 得到前缀和之后，对于每个开始下标 i，可通过二分查找得到大于或等于 i 的最小下标 bound，
// 使得 sums[bound]−sums[i−1]≥s，并更新子数组的最小长度（此时子数组的长度是 bound−(i−1)）。
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i <= n; i++ {
		// sums[bound] − sums[i−1] ≥ s
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
