package day_14

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	return sumTarget(nums, target)
}

func sumTarget(nums []int, target int) bool {
	arr := make([][]bool, len(nums)+1)
	arr[0] = make([]bool, target+1)
	for i := 1; i <= len(nums); i++ {
		arr[i] = make([]bool, target+1)
		arr[i][0] = true
		for j := 1; j <= target; j++ {
			arr[i][j] = arr[i-1][j]
			if !arr[i][j] && j-nums[i-1] >= 0 {
				arr[i][j] = arr[i-1][j-nums[i-1]]
			}
		}
	}
	return arr[len(nums)][target]
}

// f(i,j)   f(i,target)
// f(i,0) == true
// i=0 && j > 0 ;f(0,j) == false
// f(i,j) ==  f(i-1,j-num[i-1]) || f(i-1,j)
