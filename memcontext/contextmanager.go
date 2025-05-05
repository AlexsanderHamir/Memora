package memcontext

import (
	"fmt"
)

func NewContextManager() *DefaultContextManager {
	return &DefaultContextManager{
		contexts: make(map[string]*DefaultContext),
	}
}

// CreateContext creates a new context for the given name, if the context manager is nil,
// the context will be created, but not stored.
func CreateContext(cm *DefaultContextManager, name string) *DefaultContext {
	ctx := &DefaultContext{
		ctxName: name,
		pools:   make(map[string]any),
	}

	if cm != nil {
		cm.mu.Lock()
		defer cm.mu.Unlock()
		cm.contexts[name] = ctx
	}

	return ctx
}

// GetContext gets the context for the given name, if the context manager is nil,
// the context will not be returned, because it wasnt stored.
func GetContext(dcm *DefaultContextManager, name string) (*DefaultContext, error) {
	if dcm == nil {
		return nil, ErrContextManagerNil
	}

	dcm.mu.RLock()
	defer dcm.mu.RUnlock()

	val, ok := dcm.contexts[name]
	if !ok {
		return nil, fmt.Errorf("context %s not found", name)
	}

	return val, nil
}

// GetOrCreateContext gets the context for the given name, if the context manager is nil,
// the context will be created, but not stored.
func GetOrCreateContext(dcm *DefaultContextManager, name string) (*DefaultContext, error) {
	if dcm == nil {
		return CreateContext(dcm, name), nil
	}

	ctx, err := GetContext(dcm, name)
	if err != nil {
		return CreateContext(dcm, name), nil
	}

	return ctx, nil
}

func (dcm *DefaultContextManager) DeleteContext(name string) {
	dcm.mu.Lock()
	defer dcm.mu.Unlock()

	delete(dcm.contexts, name)
}
