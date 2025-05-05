package memcontext

import (
	"fmt"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

func (dcm *DefaultContext[T]) CreatePool(name string, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	if _, ok := dcm.pools[name]; ok {
		return nil, ErrPoolAlreadyExists
	}

	pool, err := pool.NewPool(config, allocator, cleaner)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPoolNotInitialized, err)
	}

	dcm.pools[name] = pool

	return pool, nil
}

func (dcm *DefaultContext[T]) GetPool(name string) (*pool.Pool[T], error) {
	dcm.mu.RLock()
	defer dcm.mu.RUnlock()

	pool, ok := dcm.pools[name]
	if !ok {
		return nil, ErrPoolNotFound
	}

	return pool, nil
}

func (dcm *DefaultContext[T]) GetOrCreatePool(name string, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	pool, err := dcm.GetPool(name)
	if err != nil {
		return dcm.CreatePool(name, config, allocator, cleaner)
	}
	return pool, nil
}

func (dcm *DefaultContext[T]) DeletePool(name string) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	delete(dcm.pools, name)
}
