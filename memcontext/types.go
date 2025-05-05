package memcontext

import (
	"sync"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

type DefaultContextManager struct {
	mu       sync.RWMutex
	contexts map[string]any
}

type DefaultContext[T any] struct {
	mu      sync.RWMutex
	ctxName string
	pools   map[string]*pool.Pool[T]
}
