package synctool

import "sync"

type SyncMapManager[K comparable, V any] struct {
	cache sync.Map
}

func NewSyncMapManager[K comparable, V any]() *SyncMapManager[K, V] {
	return &SyncMapManager[K, V]{}
}

func (c *SyncMapManager[K, V]) has(value any) V {
	if value != nil {
		if v, ok := value.(V); ok {
			return v
		}
	}
	var zero V
	return zero
}

// Delete
func (c *SyncMapManager[K, V]) Load(key K) (V, bool) {
	value, ok := c.cache.Load(key)
	return c.has(value), ok
}

// Delete
func (c *SyncMapManager[K, V]) Delete(key K) {
	c.cache.Delete(key)
}

// Clear
func (c *SyncMapManager[K, V]) Clear() {
	c.cache.Clear()
}

// CompareAndDelete
func (c *SyncMapManager[K, V]) CompareAndDelete(key K, value V) bool {
	return c.cache.CompareAndDelete(key, value)
}

// CompareAndSwap
func (c *SyncMapManager[K, V]) CompareAndSwap(key K, oldValue, NewValue V) bool {
	return c.cache.CompareAndSwap(key, oldValue, NewValue)
}

// LoadAndDelete
func (c *SyncMapManager[K, V]) LoadAndDelete(key K) (V, bool) {
	value, load := c.cache.LoadAndDelete(key)
	return c.has(value), load
}

// LoadOrStore
func (c *SyncMapManager[K, V]) LoadOrStore(k K, v V) (V, bool) {
	value, load := c.cache.LoadOrStore(k, v)
	return c.has(value), load
}

// Store 插入
func (c *SyncMapManager[K, V]) Store(key K, value V) {
	c.cache.Store(key, value)
}

// Store
func (c *SyncMapManager[K, V]) Range(f func(key K, value V) bool) {
	c.cache.Range(func(k, v any) bool {
		kv, ok := k.(K)
		if !ok {
			return true
		}
		nv, ok := v.(V)
		if !ok {
			return true
		}

		return f(kv, nv)
	})
}

// Length
func (c *SyncMapManager[K, V]) Length() int {
	var l int
	c.Range(func(key K, value V) bool {
		l++
		return true
	})
	return l
}
