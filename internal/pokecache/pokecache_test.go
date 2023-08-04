package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key",
			inputVal: []byte("val"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("Expected cache to have key")
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("Expected %s, got %s", string(cs.inputVal), string(actual))
		}
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val"))

	time.Sleep(interval / 4)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("Expected cache to not have key")
	}
}
