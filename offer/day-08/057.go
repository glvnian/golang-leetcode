package day_08

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var numMap = map[int]int{}
	var bucketSize = t + 1
	for i, num := range nums {
		var id = getBucketID(num, bucketSize)
		if _, ok := numMap[id]; ok {
			fmt.Println(">", i, numMap[id], num)
			return true
		}
		if n, ok := numMap[id-1]; ok && n >= num-t {
			fmt.Println("-", n, num)
			return true
		}
		if n, ok := numMap[id+1]; ok && n <= num+t {
			fmt.Println("+", n, num)
			return true
		}

		numMap[id] = num
		if i >= k {
			rID := getBucketID(nums[i-k], bucketSize)
			fmt.Println("delete: ", i-k, nums[i-k])
			delete(numMap, rID)
		}
		fmt.Println(numMap)
	}
	return false
}

func getBucketID(n, bucketSize int) int {
	if n >= 0 {
		return n / bucketSize
	} else { // n = -1 ; return -1
		return (n+1)/bucketSize - 1
	}
}
