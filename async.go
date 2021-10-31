package cache

type AsyncCache interface {
	Put(element Cacheable) chan Cacheable
	Get(id ID, compute func() Cacheable) chan Cacheable
}

type asyncCache struct {
	cache Cache
}

func NewAsync(cache Cache) AsyncCache {
	return &asyncCache{cache: cache}
}

func (a *asyncCache) Put(element Cacheable) chan Cacheable {
	ch := make(chan Cacheable, 1)

	go func() {
		ch <- a.cache.Put(element)
	}()

	return ch
}

func (a *asyncCache) Get(id ID, compute func() Cacheable) chan Cacheable {
	ch := make(chan Cacheable, 1)

	go func() {
		ch <- a.cache.Get(id, compute)
	}()

	return ch
}
