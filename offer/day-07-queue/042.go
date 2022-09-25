package day_07_queue

type RecentCounter struct {
	size int
	nums []int
}

func RConstructor() RecentCounter {
	return RecentCounter{
		size: 3000,
		nums: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.nums) != 0 && this.nums[0]+this.size < t {
		this.nums = this.nums[1:]
	}
	this.nums = append(this.nums, t)
	return len(this.nums)
}
