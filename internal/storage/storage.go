package storage

import "time"

type Storage[K comparable, V any] interface {
	Set(key K, value V, ttl time.Duration) error
	Get(key K) (V, error)
	Close() error
}
