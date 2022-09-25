package day_10

type TireNum struct {
	num      int
	children [26]*TireNum
}

type MapSum struct {
	root *TireNum
}

/** Initialize your data structure here. */
func SConstructor() MapSum {
	return MapSum{
		root: &TireNum{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.root
	for _, ch := range key {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &TireNum{}
		}
		node = node.children[ch-'a']
	}
	node.num = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, ch := range prefix {
		if node.children[ch-'a'] == nil {
			return 0
		}
		node = node.children[ch-'a']
	}
	return getSum(node)
}

func getSum(node *TireNum) int {
	if node == nil {
		return 0
	}
	var sum = node.num
	for _, child := range node.children {
		sum = sum + getSum(child)
	}
	return sum
}
