package day_09

import (
	"container/heap"
	"fmt"
)

// 最大堆
func kSmallestPairsV1(nums1 []int, nums2 []int, k int) [][]int {
	var res = [][]int{}
	var p = &MinPQ{
		PQ:   &[][2]int{},
		num1: nums1,
		num2: nums2,
	}

	for i := 0; i < Min(k, len(nums1)); i++ {
		heap.Push(p, [2]int{i, 0})
	}
	var item [2]int
	for k > 0 && p.Len() != 0 {
		fmt.Println(">>> p.PQ:", p.PQ)
		item = heap.Pop(p).([2]int)
		fmt.Println(">>> item:", item)
		res = append(res, []int{nums1[item[0]], nums2[item[1]]})
		fmt.Println(">>> res:", res)

		if item[1] < len(nums2)-1 {
			fmt.Println("add:", item[0], item[1]+1)
			heap.Push(p, [2]int{item[0], item[1] + 1})
			heap.Init(p)
		}
		k--
	}
	return res
}

type MinPQ struct {
	PQ   *[][2]int
	num1 []int
	num2 []int
}

func (p MinPQ) Len() int {
	return len((*p.PQ))
}

// 最小堆
func (p MinPQ) Less(i, j int) bool {
	return p.num1[(*p.PQ)[i][0]]+p.num2[(*p.PQ)[i][1]] < p.num1[(*p.PQ)[j][0]]+p.num2[(*p.PQ)[j][1]]
}

func (p MinPQ) Swap(i, j int) {
	(*p.PQ)[i], (*p.PQ)[j] = (*p.PQ)[j], (*p.PQ)[i]
}

func (p *MinPQ) Push(val interface{}) {
	*p.PQ = append(*p.PQ, val.([2]int))
}

func (p *MinPQ) Pop() interface{} {
	n := p.Len() - 1
	old := (*p.PQ)[n]
	(*p.PQ) = (*p.PQ)[:n]
	return old
}
