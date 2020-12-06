package main

import (
	"log"
	"sync"
)

func main() {
	cache := Constructor(2)
	log.Printf("cache: %v", cache)


	cache.Put(1, 1)

	cache.Put(2, 2)

	println(cache.Get(1))

	cache.Put(3, 3)

	println(cache.Get(2))

	cache.Put(4, 4)

	println(cache.Get(1))

	println(cache.Get(3))

	println(cache.Get(4))

	for node := cache.head; node != nil; node = node.next {
		println("key: ", node.key, "value:", node.value)
	}
}

type LRUCache struct {
	// map 提供 O(0) 的查询时间复杂度
	cacheData map[int]*CacheNode
	// cache cap
	capacity int
	// cache len
	length int
	// 链表的收尾节点
	head, tail *CacheNode
	// 读写锁，避免并发写入
	rwMutex sync.RWMutex
}

// data node
type CacheNode struct {
	key int
	value int
	next, prev *CacheNode
}

func Constructor(capacity int) LRUCache {
	var mutex sync.RWMutex
	cacheData := make(map[int]*CacheNode)
	cache := LRUCache{
		length: 0,
		capacity: capacity,
		rwMutex: mutex,
		cacheData: cacheData,
	}
	cache.InitCacheList()

	return cache
}

func (this *LRUCache) InitCacheList() {
	head := &CacheNode {}
	tail := &CacheNode {}

	head.next = tail
	tail.prev = head

	this.head = head
	this.tail = tail
}

func (this *LRUCache) addCacheNodeToHead(node *CacheNode) {
	if this.head == nil {
		panic("this head is null")
	}

	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) addCacheNodeToTail(node *CacheNode) {
	if this.tail ==  nil {
		panic("LRUCache tail is nil")
	}

	node.next = this.tail
	node.prev = this.tail.prev

	this.tail.prev.next = node
	this.tail.prev = node
}

func (this *LRUCache) removeCacheNode(node *CacheNode) {
	if node.prev == nil || node.next ==nil {
		panic("node is not in list")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) Get(key int) int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()

	node, ok := this.cacheData[key]
	if !ok {
		return -1
	}

	this.removeCacheNode(node)
	this.addCacheNodeToTail(node)
	return node.value

}

func (this *LRUCache) Put(key int, value int)  {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()

	node := &CacheNode{key: key, value: value}

	// 已经存在就更新
	if this.cacheData[key] != nil {
		this.removeCacheNode(this.cacheData[key])
		this.addCacheNodeToTail(node)
		this.cacheData[key] = node
		return
	}

	if this.length < this.capacity {
		this.length = this.length + 1

		// 插入节点到最后
		this.addCacheNodeToTail(node)
		this.cacheData[key] = node
	} else {
		delete(this.cacheData, this.head.next.key)
		this.removeCacheNode(this.head.next)
		this.cacheData[key] = node
		this.addCacheNodeToTail(node)
	}
}