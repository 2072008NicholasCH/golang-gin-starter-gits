package interfaces

import (
	"time"
)

// Cacheable is an interface for managing cache entries
type Cacheable interface {
	// Get returns the velue for the given key
	Get(key string) ([]byte, error)
	// Set sets the value for the given key
	Set(key string, value interface{}, ttl int) error
	// SetWithExpireAt set key value and update expire using unix timestamp
	SetWithExpireAt(key string, value interface{}, ttl time.Time) error
	// Exists Check if key is exist in redis
	Exists(key string) (bool, error)
	// Remove cache by cache key
	Remove(key string) error
	//BulkRemove cache by certain cache key pattern
	BulkRemove(keyPattern string) error
	// Scan all cache key with certain pattern
	Scan(keyPattern string) ([]string, error)
}
