package day_09

import "container/heap"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 最大堆
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res = [][]int{}
	var p = &MaxPQ{}

	heap.Init(p)
	l1 := Min(k, len(nums1))
	l2 := Min(k, len(nums2))
	var num1, num2 int
	for i := 0; i < l1; i++ {
		num1 = nums1[i]
		for j := 0; j < l2; j++ {
			num2 = nums2[j]
			if p.Len() < k || (*p)[0][0]+(*p)[0][1] > num1+num2 {
				heap.Push(p, [2]int{num1, num2})
			}
			if p.Len() > k {
				heap.Pop(p)
			}
		}
	}

	for p.Len() > 0 {
		item := p.Pop().([2]int)
		res = append(res, []int{item[0], item[1]})
	}

	return res
}

type MaxPQ [][2]int

func (p MaxPQ) Len() int {
	return len(p)
}

// 最大dui
func (p MaxPQ) Less(i, j int) bool {
	return p[i][0]+p[i][1] > p[j][0]+p[j][1]
}

func (p MaxPQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *MaxPQ) Push(val interface{}) {
	*p = append(*p, val.([2]int))
}

func (p *MaxPQ) Pop() interface{} {
	n := p.Len() - 1
	old := (*p)[n]
	*p = (*p)[:n]
	return old
}
