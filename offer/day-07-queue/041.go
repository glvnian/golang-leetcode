package day_07_queue

type MovingAverage struct {
	size int
	sum  int
	nums []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {

	return MovingAverage{
		size: size,
		nums: []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.size == 0 {
		return 0
	}

	if len(this.nums) == this.size {
		this.sum -= this.nums[0]
		this.nums = this.nums[1:]
	}
	this.sum += val
	this.nums = append(this.nums, val)
	return float64(float64(this.sum) / float64(len(this.nums)))
}
