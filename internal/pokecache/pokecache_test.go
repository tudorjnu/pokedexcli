package pokecache_test

import (
	"testing"
	"time"

	"github.com/tudorjnu/pokedexcli/internal/pokecache"
)

func TestCache_Add(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		interval time.Duration
		// Named input parameters for target function.
		key string
		val []byte
	}{
		{
			name:     "simple-case",
			interval: 5 * time.Minute,
			key:      "test-key",
			val:      []byte("test-value"),
		},
		{
			name:     "empty-value",
			interval: 1 * time.Minute,
			key:      "empty-key",
			val:      []byte(""),
		},
		{
			name:     "large-value",
			interval: 10 * time.Minute,
			key:      "large-key",
			val:      make([]byte, 1024*1024), // 1 MB of zero bytes
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := pokecache.NewCache(tt.interval)
			c.Add(tt.key, tt.val)
			storedVal, ok := c.Get(tt.key)
			if !ok {
				t.Errorf("Cache.Get() = _, false; want true")
			}
			if string(storedVal) != string(tt.val) {
				t.Errorf("Cache.Get() = %v; want %v", storedVal, tt.val)
			}
		})
	}
}

func TestCache_reapLoop(t *testing.T) {
	tests := []struct {
		name     string
		interval time.Duration
	}{
		// tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := pokecache.NewCache(tt.interval)
			c.Add("temp-key", []byte("temp-value"))
			time.Sleep(tt.interval + 100*time.Millisecond) // wait for the reap interval to pass
			_, ok := c.Get("temp-key")
			if ok {
				t.Errorf("Cache.Get() = _, true; want false after reap interval")
			}
		})
	}
}
