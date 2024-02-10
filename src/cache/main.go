package cache

type Cache interface {
	Connect() error
	Close() error
}

var cache *Cache
