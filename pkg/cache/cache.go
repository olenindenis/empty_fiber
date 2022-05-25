package cache

type CachedItem interface {
	GetKey() string
	GetValue() []byte
	GetExpiration() int32
}

type Service interface {
	Set(key string, value []byte, expiration int32) error
	Get(key string) (CachedItem, error)
	Exists(key string) bool
	Delete(key string) error
	Ping() error
}
