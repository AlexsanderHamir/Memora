package memcontext

import "github.com/AlexsanderHamir/PoolX/src/pool"

type ContextManager interface {
	CreateContext(name string) Context
	GetContext(name string) (Context, error)
	GetOrCreateContext(name string) (Context, error)
	DeleteContext(name string)
}

type Context interface {
	CreatePool(name string, config *pool.PoolConfig, allocator func() any, cleaner func(item any)) (any, error)
	GetPool(name string) (any, error)
	GetOrCreatePool(name string, config *pool.PoolConfig, allocator func() any, cleaner func(item any)) (any, error)
	DeletePool(name string)
}
