package sort

import "fmt"

//插入排序的思想非常简单，生活中有一个很常见的场景：在打扑克牌时，我们一边抓牌一边给扑克牌排序，
//每次摸一张牌，就将它插入手上已有的牌中合适的位置，逐渐完成整个排序。
//
//插入排序有两种写法：
//交换法：在新数字插入过程中，不断与前面的数字交换，直到找到自己合适的位置。
//移动法：在新数字插入过程中，与前面的数字不断比较，前面的数字不断向后挪出位置，当新数字找到自己的位置后，插入一次即可。

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i
		fmt.Println("num>>>", nums, i)
		for j >= 1 && nums[j] < nums[j-1] {
			swap(nums, j, j-1)
			fmt.Println("num>", nums)
			j--
		}
	}
}
