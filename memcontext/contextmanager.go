package memcontext

import (
	"fmt"

	"github.com/AlexsanderHamir/PoolX/src/pool"
)

func NewContextManager() *DefaultContextManager {
	return &DefaultContextManager{
		contexts: make(map[string]any),
	}
}

func CreateContext[T any](cm *DefaultContextManager, name string) *DefaultContext[T] {
	ctx := &DefaultContext[T]{
		ctxName: name,
		pools:   make(map[string]*pool.Pool[T]),
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.contexts[name] = ctx

	return ctx
}

func GetContext[T any](dcm *DefaultContextManager, name string) (*DefaultContext[T], error) {
	dcm.mu.RLock()
	defer dcm.mu.RUnlock()

	val, ok := dcm.contexts[name]
	if !ok {
		return nil, fmt.Errorf("context %s not found", name)
	}

	ctx, ok := val.(*DefaultContext[T])
	if !ok {
		return nil, fmt.Errorf("context %s has unexpected type", name)
	}

	return ctx, nil
}

func GetOrCreateContext[T any](dcm *DefaultContextManager, name string) (*DefaultContext[T], error) {
	ctx, err := GetContext[T](dcm, name)
	if err != nil {
		return CreateContext[T](dcm, name), nil
	}
	return ctx, nil
}

func (dcm *DefaultContextManager) DeleteContext(name string) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	delete(dcm.contexts, name)
}
