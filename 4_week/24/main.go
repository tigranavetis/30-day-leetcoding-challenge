package main

import (
	"container/list"
)

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
