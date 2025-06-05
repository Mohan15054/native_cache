package eviction

import (
	"container/list"
)

type EvictionPolicy interface {
	OnAccess(key string)
	OnInsert(key string)
	Evict() string
}

type lru struct {
	capacity int
	ll       *list.List
	cache    map[string]*list.Element
}

func NewLRU(capacity int) EvictionPolicy {
	return &lru{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (l *lru) OnAccess(key string) {
	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
	}
}

func (l *lru) OnInsert(key string) {
	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
		return
	}
	ele := l.ll.PushFront(key)
	l.cache[key] = ele
	if l.ll.Len() > l.capacity {
		l.Evict()
	}
}

func (l *lru) Evict() string {
	ele := l.ll.Back()
	if ele == nil {
		return ""
	}
	key := ele.Value.(string)
	l.ll.Remove(ele)
	delete(l.cache, key)
	return key
}
