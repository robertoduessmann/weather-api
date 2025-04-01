package cache

import (
	"sync"
	"time"
)

type item[V any] struct {
	value  V
	expiry time.Time
}

func (i item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

type cacheClient struct {
	items map[string]item[any]
	ttl   time.Duration
	mu    sync.Mutex
}

func NewCacheClient(ttl time.Duration) CacheClient {
	client := &cacheClient{
		items: make(map[string]item[any]),
		ttl:   ttl,
	}

	go func() {
		for range time.Tick(10 * time.Second) {
			for key, item := range client.items {
				if item.isExpired() {
					client.Delete(key)
				}
			}
		}
	}()

	return client
}

func (c *cacheClient) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if item.isExpired() {
		return nil, false
	}

	return item.value, found
}

func (c *cacheClient) Put(key string, value any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[any]{
		value:  value,
		expiry: time.Now().Add(c.ttl),
	}
	return true
}

func (c *cacheClient) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return true
}
