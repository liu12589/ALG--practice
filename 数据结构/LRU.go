package main

import "fmt"

// LRUCache LRU 数据结构
type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// DLinkedNode 双向链表数据结构
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 初始化一个双向链表
func newDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache: map[int]*DLinkedNode{},
		head:  newDLinkedNode(0, 0),
		tail:  newDLinkedNode(0, 0),
		cap:   capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get Get(key) 根据key获取值，如果key不存在返回-1
func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	// 使用了该节点的值，就把该节点移到队首
	lru.moveToHead(node)
	return node.value

}

// 将node节点移到队首
func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	// 首先将该节点移除
	lru.removeNode(node)
	// 将该节点插入队首
	lru.addToHead(node)
}

// 移除 node 节点
func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将 node 节点插入队首的位置
func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// Put put(key，value) 将key值放入LRU中
// 先判断 key 值是否已经存在
// 然后判断是否满了，如果LRU已经满了，将队尾元素删除,并且cache中也删除 size--
// 然后创建一个node插入队首，并且放入cache中，size++
func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := newDLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.cap {
			removeNode := lru.tail.prev
			lru.removeNode(removeNode)
			delete(lru.cache, removeNode.key)
			lru.size--
		}
	} else {
		// 如果已经存在改值了，则修改该值,并放入队首
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func main() {

	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	number1 := lru.Get(1)
	lru.Put(3, 3)
	number2 := lru.Get(2)
	lru.Put(4, 4)
	number3 := lru.Get(3)
	number4 := lru.Get(4)
	fmt.Println(number1, number2, number3, number4)
}
