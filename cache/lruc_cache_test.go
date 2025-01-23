package cache

import "testing"

func TestLRUCache(t *testing.T) {
	tests := []struct {
		capacity int
	}{
		{2},
		{3},
	}

	for _, test := range tests {
		ret := Constructor(test.capacity)
		if ret.capacity != test.capacity {
			t.Errorf("Got %d for input %d", ret.capacity, test.capacity)
		}
	}
}

func TestLRUCacheGet(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	ret := cache.Get(1)
	if ret != 1 {
		t.Errorf("Got %d for input %d", ret, 1)
	}
}

func TestLRUCachePut(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	ret := cache.Get(1)
	if ret != -1 {
		t.Errorf("Got %d for input %d", ret, 1)
	}
}

func TestLRUCachePut2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	ret := cache.Get(1)
	if ret != 1 {
		t.Errorf("Got %d for input %d", ret, 2)
	}
	cache.Put(3, 3)
	ret = cache.Get(2)
	if ret != -1 {
		t.Errorf("Got %d for input %d", ret, 2)
	}
	cache.Put(4, 4)
	ret = cache.Get(1)
	if ret != -1 {
		t.Errorf("Got %d for input %d", ret, 1)
	}
	ret = cache.Get(3)
	if ret != 3 {
		t.Errorf("Got %d for input %d", ret, 3)
	}
	ret = cache.Get(4)
	if ret != 4 {
		t.Errorf("Got %d for input %d", ret, 4)
	}
}

func TestLRUCachePut3(t *testing.T) {
	cache := Constructor(2)
	result := cache.Get(2)
	if result != -1 {
		t.Errorf("Got %d for input %d", result, 2)
	}

	cache.Put(2, 6)
	result = cache.Get(1)
	if result != -1 {
		t.Errorf("Got %d for input %d", result, 1)
	}

	cache.Put(1, 5)
	cache.Put(1, 2)
	result = cache.Get(1)
	if result != 2 {
		t.Errorf("Got %d for input %d", result, 1)
	}
	result = cache.Get(2)
	if result != 6 {
		t.Errorf("Got %d for input %d", result, 6)
	}

}
