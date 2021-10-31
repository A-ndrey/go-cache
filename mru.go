package cache

import "container/list"

type mru struct {
	items    *list.List
	itemsMap map[ID]*list.Element
	size     int
}

func NewInMemMRU(size int) Cache {
	return &mru{
		items:    list.New(),
		itemsMap: make(map[ID]*list.Element),
		size:     size,
	}
}

func (m *mru) Put(element Cacheable) Cacheable {
	if le, ok := m.itemsMap[element.Identify()]; ok {
		m.items.MoveToFront(le)
		return le.Value.(Cacheable)
	}

	if m.items.Len() == m.size {
		last := m.items.Front()
		m.items.Remove(last)
		delete(m.itemsMap, last.Value.(Cacheable).Identify())
	}

	le := m.items.PushFront(element)
	m.itemsMap[element.Identify()] = le

	return element
}

func (m *mru) Get(id ID, compute func() Cacheable) Cacheable {
	if le, ok := m.itemsMap[id]; ok {
		m.items.MoveToFront(le)
		return le.Value.(Cacheable)
	}

	if compute == nil {
		return nil
	}
	element := compute()

	return m.Put(element)
}
