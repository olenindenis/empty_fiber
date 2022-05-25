package mocks

import (
	"envs/pkg/cache"
)

type CacheItem struct {
}

func (ci CacheItem) GetKey() string {
	return ""
}

func (ci CacheItem) GetValue() []byte {
	return nil
}

func (ci CacheItem) GetExpiration() int32 {
	return 0
}

type Cache struct {
}

var _ cache.Service = (*Cache)(nil)

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Set(key string, value []byte, expiration int32) error {
	return nil
}

func (c Cache) Get(key string) (cache.CachedItem, error) {
	return CacheItem{}, nil
}

func (c Cache) Delete(key string) error {
	return nil
}

func (c Cache) Exists(key string) bool {
	return false
}

func (c Cache) Ping() error {
	return nil
}
