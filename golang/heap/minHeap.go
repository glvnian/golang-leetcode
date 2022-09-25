package heap

import "container/heap"

// 优先队列，小顶堆
type PriorQ []int

type KthLargest struct {
	PQ   *PriorQ // 最小堆，堆顶元素即为第k大的数
	Size int     // 堆的大小
}

// Constructor 构造函数
func Constructor(k int, nums []int) KthLargest {
	p := &PriorQ{} // 最小堆
	kth := KthLargest{
		p,
		k,
	}
	for _, n := range nums { // 将初始值逐个加入堆
		kth.Add(n)
	}
	return kth
}

// Add 往最小堆里面插入数据，并输出topK
func (k *KthLargest) Add(val int) int {
	// 如果堆的长度小于k或者插入的值大于最小堆顶元素，则将该值插入堆中
	if k.PQ.Len() < k.Size || val > (*k.PQ)[0] {
		heap.Push(k.PQ, val)
	}

	// 保证最小堆的大小始终不超过k
	if k.PQ.Len() > k.Size {
		heap.Pop(k.PQ)
	}

	// 返回堆顶元素
	return (*k.PQ)[0]
}

// 以下方法为heap.Interface的模板方法，需要熟练掌握

func (p PriorQ) Len() int {
	return len(p)
}

func (p PriorQ) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PriorQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorQ) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *PriorQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}
