package cache

import (
	"container/list"
)

type lru struct {
	items    *list.List
	itemsMap map[ID]*list.Element
	size     int
}

func NewInMemLRU(size int) Cache {
	return &lru{
		items:    list.New(),
		itemsMap: make(map[ID]*list.Element),
		size:     size,
	}
}

func (l *lru) Put(element Cacheable) Cacheable {
	if le, ok := l.itemsMap[element.Identify()]; ok {
		l.items.MoveToFront(le)
		return le.Value.(Cacheable)
	}

	if l.items.Len() == l.size {
		last := l.items.Back()
		l.items.Remove(last)
		delete(l.itemsMap, last.Value.(Cacheable).Identify())
	}

	le := l.items.PushFront(element)
	l.itemsMap[element.Identify()] = le

	return element
}

func (l *lru) Get(id ID, compute func() Cacheable) Cacheable {
	if le, ok := l.itemsMap[id]; ok {
		l.items.MoveToFront(le)
		return le.Value.(Cacheable)
	}

	if compute == nil {
		return nil
	}
	element := compute()

	return l.Put(element)
}
