package day_09

type PQ []int

func (p PQ) Len() int {
	return len(p)
}

func (p PQ) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQ) Push(val interface{}) {
	*p = append(*p, val.(int))
}

func (p *PQ) Pop() interface{} {
	n := p.Len() - 1
	x := (*p)[n]
	*p = (*p)[:n-1]
	return x
}
