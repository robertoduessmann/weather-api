package cache

import (
	"sync"
	"time"
)

type cacheManager struct {
	avaliablesCache map[string]CacheClient
}

var instance CacheManager = nil

func NewCacheManager() CacheManager {
	sync.OnceFunc(func() {
		if instance == nil {
			instance = cacheManager{
				avaliablesCache: make(map[string]CacheClient),
			}
		}
	})()
	return instance
}

func (m cacheManager) NewCache(name string, ttl time.Duration) CacheClient {
	client, found := m.avaliablesCache[name]
	if found {
		return client
	}

	cacheClient := NewCacheClient(ttl)
	m.avaliablesCache[name] = cacheClient
	return cacheClient
}

func (m cacheManager) Erase(name string) bool {
	return false
}

func (m cacheManager) Delete(name string) bool {
	return false
}
