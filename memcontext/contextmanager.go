package memcontext

import (
	"fmt"
	"reflect"
)

// NewContextManager creates a new context manager
func NewContextManager() ContextManager {
	return &DefaultContextManager{
		contexts: make(map[string]*DefaultContext),
	}
}

// CreateContext creates a new context for the given name, if the context manager is nil,
// the context will be created, but not stored.
func (cm *DefaultContextManager) CreateContext(name string) *DefaultContext {
	ctx := &DefaultContext{
		ctxName: name,
		pools:   make(map[reflect.Type]any),
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.contexts[name] = ctx

	return ctx
}

// GetContext gets the context for the given name, if the context manager is nil,
// the context will not be returned, because it wasnt stored.
func (cm *DefaultContextManager) GetContext(name string) (*DefaultContext, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	val, ok := cm.contexts[name]
	if !ok {
		return nil, fmt.Errorf("context %s not found", name)
	}

	return val, nil
}

// GetOrCreateContext gets the context for the given name, if the context manager is nil,
// the context will be created, but not stored.
func (cm *DefaultContextManager) GetOrCreateContext(name string) (*DefaultContext, error) {
	ctx, err := cm.GetContext(name)
	if err != nil {
		return cm.CreateContext(name), nil
	}

	return ctx, nil
}

// DeleteContext deletes the context for the given name, if the context manager is nil,
// the context will not be deleted, because it wasnt stored.
func (cm *DefaultContextManager) DeleteContext(name string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	ctx, err := cm.GetContext(name)
	if err != nil {
		return
	}

	ctx.ClosePools()
	delete(cm.contexts, name)
}

// deletes all contexts
func (cm *DefaultContextManager) DeleteAllContexts() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, ctx := range cm.contexts {
		ctx.ClosePools()
		delete(cm.contexts, ctx.ctxName)
	}
}
