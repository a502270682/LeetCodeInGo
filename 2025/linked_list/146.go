package linked_list

type DLinkedNode struct {
	Key, Val   int
	Prev, Next *DLinkedNode
}

type LRUCache struct {
	size, cap int
	head      *DLinkedNode
	tail      *DLinkedNode
	cache     map[int]*DLinkedNode
}

func Constructor(capacity int) LRUCache {
	dummyHead := initDLinkedNode(0, 0)
	dummyTail := initDLinkedNode(0, 0)
	l := LRUCache{0, capacity, dummyHead, dummyTail, make(map[int]*DLinkedNode)}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key, value, nil, nil}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		node.Val = value
		this.moveToHead(node)
		return
	} else {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.cap {
			removed := this.removeTail()
			delete(this.cache, removed.Key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	next := this.head.Next
	this.head.Next = node
	node.Prev = this.head
	node.Next = next
	next.Prev = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
