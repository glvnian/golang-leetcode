package day_05

import (
	"math/rand"
)

type RandomizedSet struct {
	Nums []int
	Map  map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		Nums: make([]int, 0),
		Map:  make(map[int]int), // key: 是需要保存的内容，val: 保存在nums上对应的位置
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; !ok {
		this.Map[val] = len(this.Nums)
		this.Nums = append(this.Nums, val)
		return true
	}

	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {

	if _, ok := this.Map[val]; !ok {
		return false
	}

	var location = this.Map[val]
	var last = len(this.Nums) - 1

	this.Nums[location] = this.Nums[last]
	this.Map[this.Nums[location]] = location

	this.Nums = this.Nums[:len(this.Nums)]
	delete(this.Map, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.Nums[rand.Intn(len(this.Nums))]
}
