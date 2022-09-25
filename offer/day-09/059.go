package day_09

import "container/heap"

type KthLargest struct {
	p *PQ
	k int
}

func Constructor(k int, nums []int) KthLargest {
	k1 := KthLargest{
		p: &PQ{},
		k: k,
	}
	for _, num := range nums {
		k1.Add(num)
	}
	return k1

}

func (this *KthLargest) Add(val int) int {
	if this.p.Len() < this.k || val > (*this.p)[0] {
		heap.Push(this.p, val)
	}
	if this.p.Len() > this.k {
		heap.Pop(this.p)
	}

	return (*this.p)[0]
}
