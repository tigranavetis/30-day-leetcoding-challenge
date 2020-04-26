package main

import (
	"container/list"
)

/**
*
* Design and implement a data structure for Least Recently Used (LRU) cache.
* It should support the following operations: get and put.
*
* get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
* put(key, value) - Set or insert the value if the key is not already present.
* When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
*
* The cache is initialized with a positive capacity.
*
* Follow up:
* Could you do both operations in O(1) time complexity?
*
* Example:
* LRUCache cache = new LRUCache(capacity);
*
* cache.put(1, 1);
* cache.put(2, 2);
* cache.get(1);       // returns 1
* cache.put(3, 3);    // evicts key 2
* cache.get(2);       // returns -1 (not found)
* cache.put(4, 4);    // evicts key 1
* cache.get(1);       // returns -1 (not found)
* cache.get(3);       // returns 3
* cache.get(4);       // returns 4
*
**/

type LRUCache struct {
	capacity int
	elements *list.List
	cache    map[int]*list.Element
}

type Node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element, capacity),
		elements: new(list.List),
	}
}

func (this *LRUCache) Get(key int) int {
	if el, ok := this.cache[key]; ok {
		value := el.Value.(*list.Element).Value.(Node).value
		this.elements.MoveToFront(el)
		return value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if el, ok := this.cache[key]; ok {
		this.elements.MoveToFront(el)
		el.Value.(*list.Element).Value = Node{key: key, value: value}
	} else {
		if this.capacity == this.elements.Len() {
			back := this.elements.Back()
			keyToDelete := back.Value.(*list.Element).Value.(Node).key
			delete(this.cache, keyToDelete)
			this.elements.Remove(back)
		}

		element := &list.Element{
			Value: Node{
				key:   key,
				value: value,
			},
		}

		p := this.elements.PushFront(element)
		this.cache[key] = p
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)    // returns 1
	lruCache.Put(3, 3) // evicts key 2
	lruCache.Get(2)    // returns -1 (not found)
	lruCache.Put(4, 4) // evicts key 1
	lruCache.Get(1)    // returns -1 (not found)
	lruCache.Get(3)    // returns 3
	lruCache.Get(4)    // returns 4
}
