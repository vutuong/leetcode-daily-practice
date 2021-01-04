package main

import "fmt"

/*
https://leetcode.com/problems/lru-cache/
146. LRU Cache
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.

Follow up:
Could you do get and put in O(1) time complexity?
*/

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Capacity int
	Head     *Node
	Tail     *Node
	Cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node)
	return LRUCache{Capacity: capacity, Cache: cache}
}

func (this *LRUCache) addToHead(node *Node) {
	node.Prev = nil
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
	return
}

func (this *LRUCache) removeNode(node *Node) {
	if node != this.Head {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}
	if node != this.Tail {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.removeNode(node)
		this.addToHead(node)
		return node.Val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Val = value
		this.removeNode(node)
		this.addToHead(node)
		return
	}
	node := &Node{Key: key, Val: value}
	this.Cache[key] = node
	this.addToHead(node)

	if len(this.Cache) > this.Capacity {
		delete(this.Cache, this.Tail.Key)
		this.removeNode(this.Tail)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lru_cache := Constructor(2)
	lru_cache.Put(1, 1)
	fmt.Println(lru_cache)
	lru_cache.Put(2, 2)
	fmt.Println(lru_cache)
	lru_cache.Get(1)
	fmt.Println(lru_cache)
	lru_cache.Put(3, 3)
	fmt.Println(lru_cache)
	lru_cache.Get(2)
	fmt.Println(lru_cache)
	lru_cache.Put(4, 4)
	fmt.Println(lru_cache)
	lru_cache.Get(1)
	fmt.Println(lru_cache)
	lru_cache.Get(3)
	fmt.Println(lru_cache)
	lru_cache.Get(4)
	fmt.Println(lru_cache)
}
