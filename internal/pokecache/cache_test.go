package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("a key", []byte("a value"))
	cache.Add("an key", []byte("an value"))
	if len(cache.cacheMap) != 2 {
		t.Errorf("cache.reapLoop: wrong length of the cacheMap")
	}
	time.Sleep(5 * time.Millisecond)
	cache.Add("not key", []byte("not value"))
	if len(cache.cacheMap) != 3 {
		t.Errorf("cache.reapLoop: wrong length of the cacheMap")
	}
	time.Sleep(6 * time.Millisecond)
	if len(cache.cacheMap) != 1 {
		t.Errorf("cache.reapLoop: wrong length of the cacheMap")
	}
	time.Sleep(6 * time.Millisecond)
	if len(cache.cacheMap) != 1 {
		t.Errorf("cache.reapLoop: wrong length of the cacheMap")
	}
	time.Sleep(5 * time.Millisecond)
	if len(cache.cacheMap) != 0 {
		t.Errorf("cache.reapLoop: wrong length of the cacheMap")
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("a key", []byte("a value"))

	result, exist := cache.Get("a key")
	if string(result) != "a value" || exist != true {
		t.Errorf("cache.Get: error getting existing value")
	}
	_, notexist := cache.Get("wrong key")
	if notexist != false {
		t.Errorf("cache.Get: error getting non-existing value")
	}
}
