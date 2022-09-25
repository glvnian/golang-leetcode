package day_13

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	helperPermuteUnique(nums, 0, &result)
	return result
}

func helperPermuteUnique(nums []int, i int, result *[][]int) {

	if i == len(nums) {
		//desc := make([]int, len(nums))
		//copy(desc, nums)
		*result = append(*result, append([]int{}, nums...))
	} else {
		setMap := map[int]struct{}{}
		for j := i; j < len(nums); j++ {
			if _, ok := setMap[nums[j]]; !ok {
				setMap[nums[j]] = struct{}{}
				nums[i], nums[j] = nums[j], nums[i]
				helperPermuteUnique(nums, i+1, result)
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
