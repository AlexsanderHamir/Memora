package memcontext

import (
	"fmt"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

func CreatePool[T any](dc *DefaultContext, name string, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if _, ok := dc.pools[name]; ok {
		return nil, ErrPoolAlreadyExists
	}

	pool, err := pool.NewPool(config, allocator, cleaner)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPoolNotInitialized, err)
	}

	dc.pools[name] = pool

	return pool, nil
}

func GetPool[T any](dc *DefaultContext, name string) (*pool.Pool[T], error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	Ipool, ok := dc.pools[name]
	if !ok {
		return nil, ErrPoolNotFound
	}

	pool, ok := Ipool.(*pool.Pool[T])
	if !ok {
		return nil, fmt.Errorf("pool with name %q has wrong type", name)
	}

	return pool, nil
}

func GetOrCreatePool[T any](dc *DefaultContext, name string, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	pool, err := GetPool[T](dc, name)
	if err != nil {
		return CreatePool(dc, name, config, allocator, cleaner)
	}
	return pool, nil
}

func (dcm *DefaultContext) DeletePool(name string) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	delete(dcm.pools, name)
}
