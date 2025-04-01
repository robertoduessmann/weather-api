package cache

import (
	"sync"
	"time"
)

type cacheManager struct {
	avaliablesCache map[string]CacheClient
	mu              sync.RWMutex
}

var (
	instance CacheManager
	once     sync.Once
)

func NewCacheManager() CacheManager {
	once.Do(func() {
		instance = &cacheManager{
			avaliablesCache: make(map[string]CacheClient),
		}
	})
	return instance
}

func (m *cacheManager) NewCache(name string, ttl time.Duration) CacheClient {
	m.mu.RLock()
	client, found := m.avaliablesCache[name]
	m.mu.RUnlock()
	if found {
		return client
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	cacheClient := NewCacheClient(ttl)
	m.avaliablesCache[name] = cacheClient
	return cacheClient
}

func (m *cacheManager) Erase(name string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.avaliablesCache[name]; exists {
		delete(m.avaliablesCache, name)
		return true
	}
	return false
}

func (m *cacheManager) Delete(name string) bool {
	return m.Erase(name)
}
