package cache

type LRUCache struct {
	stroage  map[int]*Node
	capacity int
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		stroage:  make(map[int]*Node),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.stroage[key]; ok {
		cache.Delete(node)
		cache.Insert(node)
		return node.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {

	if node, ok := cache.stroage[key]; ok {
		cache.Delete(node)
	}

	if len(cache.stroage) >= cache.capacity {
		cache.Delete(cache.tail.prev)
	}

	node := &Node{
		key:   key,
		value: value,
	}
	cache.Insert(node)
}

func (cache *LRUCache) Delete(node *Node) {
	delete(cache.stroage, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (cache *LRUCache) Insert(node *Node) {

	cache.stroage[node.key] = node

	nextNode := cache.head.next
	cache.head.next = node
	nextNode.prev = node
	node.prev = cache.head
	node.next = nextNode
}

