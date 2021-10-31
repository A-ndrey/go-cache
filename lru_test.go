package cache

import "testing"

type lruTestStruct struct {
	ID
}

func TestLru_Simple(t *testing.T) {
	lru := NewInMemLRU(4)

	testStructs := []lruTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs {
		lru.Put(testStructs[i])
	}

	for i := range testStructs {
		res := lru.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test1" {
			if res == nil {
				continue
			}
			t.Errorf("cache contains \"test1\"")
		}

		if res == nil {
			t.Errorf("cache doesn't contain element")
		}

		if res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}

func TestLru_WithComputeFunc(t *testing.T) {
	lru := NewInMemLRU(4)

	testStructs := []lruTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs {
		lru.Put(testStructs[i])
	}

	lru.Get(testStructs[0].ID, func() Cacheable { return testStructs[0] })

	for i := range testStructs {
		res := lru.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test2" {
			if res == nil {
				continue
			}
			t.Errorf("cache contains \"test2\"")
		}

		if res == nil {
			t.Errorf("cache doesn't contain element")
		}

		if res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}
