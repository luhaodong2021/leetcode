package lrucache

import "container/list"

type LRUCache struct {
	Cap  int
	Map  map[int]*list.Element
	List *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*list.Element), list.New()}
}

func (lru *LRUCache) Get(key int) int {
	var e *list.Element
	var ok bool
	if e, ok = lru.Map[key]; ok {
		lru.List.MoveToFront(e)
		return e.Value.([2]int)[1]
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	var e *list.Element
	var ok bool

	if e, ok = lru.Map[key]; ok {
		lru.List.MoveToFront(e)
		e.Value = [2]int{key, value}
	} else {
		lru.List.PushFront([2]int{key, value})
		e = lru.List.Front()
	}
	lru.Map[key] = e
	if lru.List.Len() > lru.Cap {
		delete(lru.Map, lru.List.Back().Value.([2]int)[0])
		lru.List.Remove(lru.List.Back())
	}
}
