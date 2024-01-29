package client

import (
	"github.com/phuslu/lru"
	"time"
)

type Phuslu[K comparable, V any] struct {
	client *lru.Cache[K, V]
}

func (c *Phuslu[K, V]) Init(cap int) {
	c.client = lru.New[K, V](cap)
}

func (c *Phuslu[K, V]) Name() string {
	return "Phuslu"
}

func (c *Phuslu[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *Phuslu[K, V]) Set(key K, value V) {
	c.client.Set(key, value, time.Hour)
}

func (c *Phuslu[K, V]) Close() {
	c.client = nil
}
