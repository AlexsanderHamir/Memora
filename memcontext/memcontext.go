package memcontext

import (
	"fmt"
	"reflect"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

// Creates a new pool in the context
func CreatePool[T any](dc *DefaultContext, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := dc.pools[t]; ok {
		return nil, ErrPoolAlreadyExists
	}

	pool, err := pool.NewPool(config, allocator, cleaner)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPoolNotInitialized, err)
	}

	dc.pools[t] = pool

	return pool, nil
}

// Gets a pool from the context, if the pool does not exist, it returns an error
func GetPool[T any](dc *DefaultContext) (*pool.Pool[T], error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	Ipool, ok := dc.pools[t]
	if !ok {
		return nil, ErrPoolNotFound
	}

	pool, ok := Ipool.(*pool.Pool[T])
	if !ok {
		return nil, fmt.Errorf("pool for type %v has wrong type", t)
	}

	return pool, nil
}

// Gets a pool from the context, if the pool does not exist, it creates a new one
func GetOrCreatePool[T any](dc *DefaultContext, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	pool, err := GetPool[T](dc)
	if err != nil {
		return CreatePool(dc, config, allocator, cleaner)
	}
	return pool, nil
}

// Deletes a single pool from the context
func DeletePool[T any](dc *DefaultContext) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	delete(dc.pools, t)
}

// Closes all pools in the context
func (dc *DefaultContext) ClosePools() {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	for poolType, pool := range dc.pools {
		if p, ok := pool.(interface{ Close() }); ok {
			p.Close()
		}
		delete(dc.pools, poolType)
	}
}
