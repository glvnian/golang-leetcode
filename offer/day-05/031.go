package day_05

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	nodeMap  map[int]*Node
}

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

func MConstructor(capacity int) LRUCache {
	var head = &Node{key: -1, val: -1}
	var tail = &Node{key: -1, val: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		nodeMap:  map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	var node *Node
	var ok bool
	if node, ok = this.nodeMap[key]; !ok {
		return -1
	}
	// 移动到最后
	this.moveTail(node, node.val)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	var node *Node
	var ok bool
	// map 中已经存在
	if node, ok = this.nodeMap[key]; ok {
		this.moveTail(node, value)
		return
	}

	// 不存在
	var size int
	size = len(this.nodeMap)
	if size >= this.capacity {
		var deleteNode = this.head.next
		this.deleteNode(deleteNode)
		delete(this.nodeMap, deleteNode.key)
	}
	node = &Node{
		key: key,
		val: value,
	}
	this.insertToTail(node)
	this.nodeMap[key] = node
}

func (this *LRUCache) moveTail(node *Node, newValue int) {
	this.deleteNode(node)

	node.val = newValue
	this.insertToTail(node)
}

// a > b > c
func (this *LRUCache) deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//  tail:  t-1 > -1  |  t-1 > node > -1
func (this *LRUCache) insertToTail(node *Node) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}
