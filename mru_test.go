package cache

import "testing"

type mruTestStruct struct {
	ID
}

func TestMru_Simple(t *testing.T) {
	mru := NewInMemMRU(4)

	testStructs := []mruTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs {
		mru.Put(testStructs[i])
	}

	for i := range testStructs {
		res := mru.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test4" {
			if res == nil {
				continue
			}
			t.Errorf("cache contains \"test4\"")
		}

		if res == nil {
			t.Errorf("cache doesn't contain element")
		}

		if res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}

func TestMru_WithComputeFunc(t *testing.T) {
	mru := NewInMemMRU(4)

	testStructs := []mruTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs {
		mru.Put(testStructs[i])
	}

	mru.Get(testStructs[3].ID, func() Cacheable { return testStructs[3] })

	for i := range testStructs {
		res := mru.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test5" {
			if res == nil {
				continue
			}
			t.Errorf("cache contains \"test5\"")
		}

		if res == nil {
			t.Errorf("cache doesn't contain element")
		}

		if res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}
