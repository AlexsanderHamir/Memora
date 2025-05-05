package memcontext

import (
	"fmt"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

func (dcm *DefaultContext) CreatePool(name string, config *pool.PoolConfig, allocator func() any, cleaner func(item any)) (any, error) {
	pool, err := pool.NewPool(config, allocator, cleaner)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPoolNotInitialized, err)
	}

	_, ok := dcm.pools.Load(name)
	if ok {
		return nil, ErrPoolAlreadyExists
	}

	dcm.pools.Store(name, pool)

	return pool, nil
}

func (dcm *DefaultContext) GetPool(name string) (any, error) {
	pool, ok := dcm.pools.Load(name)
	if !ok {
		return nil, ErrPoolNotFound
	}

	return pool, nil
}

func (dcm *DefaultContext) GetOrCreatePool(name string, config *pool.PoolConfig, allocator func() any, cleaner func(item any)) (any, error) {
	pool, err := dcm.GetPool(name)
	if err != nil {
		return dcm.CreatePool(name, config, allocator, cleaner)
	}
	return pool, nil
}

func (dcm *DefaultContext) DeletePool(name string) {
	dcm.pools.Delete(name)
}
