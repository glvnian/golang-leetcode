package day_09

import "container/heap"

func topKFrequentV1(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}

	h := &MapMinPQ{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret

}

type MapMinPQ [][2]int

func (p MapMinPQ) Len() int {
	return len(p)
}

func (p MapMinPQ) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p MapMinPQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *MapMinPQ) Push(val interface{}) {
	*p = append(*p, val.([2]int))
}

func (p *MapMinPQ) Pop() interface{} {
	n := p.Len() - 1
	old := (*p)[n]
	*p = (*p)[:n]
	return old
}
