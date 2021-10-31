package cache

import "container/heap"

type item struct {
	value      Cacheable
	usageCount int64
	index      int
}

type lfu struct {
	items    []*item
	itemsMap map[ID]*item
	size     int
}

func (l *lfu) Len() int {
	return len(l.items)
}

func (l *lfu) Less(i, j int) bool {
	return l.items[i].usageCount < l.items[j].usageCount
}

func (l *lfu) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
	l.items[i].index = i
	l.items[j].index = j
}

func (l *lfu) Push(x interface{}) {
	n := len(l.items)
	i := x.(*item)
	i.index = n
	l.items = append(l.items, i)
}

func (l *lfu) Pop() interface{} {
	old := l.items
	n := len(old)
	i := old[n-1]
	old[n-1] = nil
	i.index = -1
	l.items = old[0 : n-1]
	return i
}

func NewInMemLFU(size int) Cache {
	return &lfu{
		itemsMap: make(map[ID]*item),
		size:     size,
	}
}

func (l *lfu) Put(element Cacheable) Cacheable {
	if i, ok := l.itemsMap[element.Identify()]; ok {
		i.usageCount++
		heap.Fix(l, i.index)
		return element
	}

	if len(l.items) == l.size {
		i := heap.Pop(l).(*item)
		delete(l.itemsMap, i.value.Identify())
	}

	i := &item{value: element}
	heap.Push(l, i)
	l.itemsMap[element.Identify()] = i

	return element
}

func (l *lfu) Get(id ID, compute func() Cacheable) Cacheable {
	if i, ok := l.itemsMap[id]; ok {
		i.usageCount++
		heap.Fix(l, i.index)
		return i.value
	}

	if compute == nil {
		return nil
	}
	element := compute()

	return l.Put(element)
}
