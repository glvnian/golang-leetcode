package day_02

import (
	"fmt"
	"sort"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/1fGaJU
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 排序之后，如果nums[i]大于0，那么后面的其他数一定都大于0。
		if nums[i] > 0 {
			break
		}

		var j = i + 1
		var k = len(nums) - 1
		var targets = -nums[i]

		// 这个是排除重复的数字，如果把这个条件拿掉，那会存在相同的接口在数组上
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			// 这个是排除重复的数字，如果把这个条件拿掉，那会存在相同的接口在数组上
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			// 这个是排除重复的数字，如果把这个条件拿掉，那会存在相同的接口在数组上
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[i]+nums[j] > 0 {
				break
			}
			if targets > nums[j]+nums[k] {
				j++
				continue
			}
			if targets < nums[j]+nums[k] {
				k--
				continue
			}
			if targets == nums[j]+nums[k] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				k--
				continue
			}

		}

	}
	return res

}

func threeSumOffice(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		j := i + 1
		k := n - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j] > 0 {
				break
			}
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}
	return res
}

func threeSumV0(nums []int) [][]int {
	fmt.Println(nums)
	var res = make([][]int, 0)
	var m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], 0 - nums[j] - nums[k]})
				}
			}
		}
	}
	return res
}

func TwoSum(numbers []int, target int) []int {

	var i = 0
	var j = len(numbers) - 1

	for i < j && i < len(numbers) {

		if target == numbers[i]+numbers[j] {
			return []int{i, j}
		}
		if target > numbers[i]+numbers[j] {
			i++
		}
		if target < numbers[i]+numbers[j] {
			j--
		}
	}
	return nil
}

// a+b+c = 0
// c = a+b
