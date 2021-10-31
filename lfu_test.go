package cache

import "testing"

type lfuTestStruct struct {
	ID
}

func TestLfu_Simple(t *testing.T) {
	lfu := NewInMemLFU(4)

	testStructs := []lfuTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs[0 : len(testStructs)-1] {
		lfu.Put(testStructs[i])
	}

	lfu.Get(testStructs[0].ID, nil)
	lfu.Put(testStructs[len(testStructs)-1])

	for i := range testStructs {
		res := lfu.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test5" && res == nil {
			t.Errorf("cache doesn't contain \"test5\"")
		}

		if testStructs[i].ID == "test1" && res == nil {
			t.Errorf("cache doesn't contain \"test1\"")
		}

		if res != nil && res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}

func TestLfu_WithComputeFunc(t *testing.T) {
	lfu := NewInMemLFU(4)

	testStructs := []lfuTestStruct{{"test1"}, {"test2"}, {"test3"}, {"test4"}, {"test5"}}

	for i := range testStructs[0 : len(testStructs)-1] {
		lfu.Put(testStructs[i])
	}

	for i := range testStructs[0 : len(testStructs)-2] {
		lfu.Get(testStructs[i].ID, nil)
	}

	lfu.Put(testStructs[len(testStructs)-1])

	lfu.Get(testStructs[len(testStructs)-2].ID, func() Cacheable { return testStructs[len(testStructs)-2] })

	for i := range testStructs {
		res := lfu.Get(testStructs[i].ID, nil)

		if testStructs[i].ID == "test5" && res != nil {
			t.Errorf("cache contains \"test5\"")
		}

		if testStructs[i].ID == "test4" && res == nil {
			t.Errorf("cache doesn't contain \"test4\"")
		}

		if res != nil && res.Identify() != testStructs[i].ID {
			t.Errorf("cache contains wrong element")
		}
	}
}
