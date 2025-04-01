package cache

import "time"

type CacheClient interface {
	Get(key string) (any, bool)
	Put(key string, value any) bool
	Delete(key string) bool
}

type CacheManager interface {
	NewCache(name string, ttl time.Duration) CacheClient
	Erase(name string) bool
	Delete(name string) bool
}
