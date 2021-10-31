package cache

type ID string

type Cacheable interface {
	Identify() ID
}

func (i ID) Identify() ID {
	return i
}

type Cache interface {
	Put(element Cacheable) Cacheable
	Get(id ID, compute func() Cacheable) Cacheable
}
