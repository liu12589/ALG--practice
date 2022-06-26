package main

import "fmt"

type LFUCache struct {
	size     int
	cap      int
	lfreq    int
	countMap map[int]*DLinked      // key:count; value:DLinked
	cache    map[int]*DLinkedNode1 // key:key; value:node
}

// DLinked 双向链表数据结构
type DLinked struct {
	head, tail *DLinkedNode1
}

// DLinkedNode1 双向链表节点数据结构
type DLinkedNode1 struct {
	key, value, freq int
	prev, next       *DLinkedNode1
}

// 初始化一个链表
func newDLink() *DLinked {
	link := &DLinked{
		head: newDLinkNode(0, 0),
		tail: newDLinkNode(0, 0),
	}
	link.head.next = link.tail
	link.tail.prev = link.head
	return link
}

// 初始化一个节点
func newDLinkNode(key, value int) *DLinkedNode1 {
	return &DLinkedNode1{
		key:   key,
		value: value,
		freq:  1,
	}
}

// Constructor1 初始化一个LFU对象
func Constructor1(capacity int) LFUCache {
	lfu := LFUCache{
		cap:      capacity,
		lfreq:    0,
		countMap: map[int]*DLinked{},
		cache:    map[int]*DLinkedNode1{},
	}
	return lfu
}

// Put 先判断 key 值是否已经存在,如果已经存在更新其值。注意 countMap 中链表中的节点和 cache 中的节点是同一个
// 然后判断是否满了，如果LRU已经满了，将队尾元素删除,并且cache中也删除 size--
// 然后创建一个node插入队首，并且放入cache中，size++
func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if _, ok := this.cache[key]; !ok {
		// 如果超出缓存，删除频次最低的最久未用的数据
		if this.size >= this.cap {
			// 找到最底层最后一个节点,然后删除该节点
			lfreq := this.lfreq
			link := this.countMap[lfreq]
			node := link.tail.prev
			this.removeNode(node)
			delete(this.cache, node.key)
			this.size--
		}
		// 如果不存在，增加节点，放入第一层，并更新 lfreq
		node := newDLinkNode(key, value)
		this.addToHead(node)
		this.lfreq = node.freq
		this.cache[key] = node
		this.size++
	} else {
		// 跟新该点值移动到更高层
		node := this.cache[key]
		node.value = value
		this.addToHighFreqLocation(node)
	}
}

// Get 获取一个数据，先判断存不存在如果不存在直接返回-1
func (this *LFUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.addToHighFreqLocation(node)
	return node.value
}

// 将该节点移动到更高频次的位置，并且放在队首
// 将元素放到上一层队首
func (this *LFUCache) addToHighFreqLocation(node *DLinkedNode1) {
	// 删除原来位置元素
	this.removeNode(node)
	node.freq++
	this.addToHead(node)
}

// 删除在原先的位置，删除完数据后检测一下，该层是否还有数据，如果没有了更新一下 lfreq
func (this *LFUCache) removeNode(node *DLinkedNode1) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// 检测删除后检测那一层是否还有数据，如果没有了，那么 lfreq+1
	link := this.countMap[node.freq]
	if link.head.next == link.tail && this.lfreq == node.freq {
		this.lfreq++
	}
}

// 增加该节点到队首，会先检测该层是否有队列
func (this *LFUCache) addToHead(node *DLinkedNode1) {
	// 检测更高层是否有队列
	freq := node.freq
	if _, ok := this.countMap[freq]; !ok {
		this.countMap[freq] = newDLink()
	}
	link := this.countMap[node.freq]
	// 这里将 node 节点连接到队首
	node.prev = link.head
	node.next = link.head.next
	link.head.next.prev = node
	link.head.next = node
}

func main() {
	lru := Constructor1(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	number2 := lru.Get(2)
	lru.Put(1, 1)
	lru.Put(4, 1)
	number3 := lru.Get(2)
	fmt.Println(number2, number3)
}
