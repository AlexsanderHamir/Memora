package memcontext

import (
	"reflect"
	"sync"
)

type ContextManager interface {
	CreateContext(name string) *DefaultContext
	GetContext(name string) (*DefaultContext, error)
	GetOrCreateContext(name string) (*DefaultContext, error)
	DeleteContext(name string)
}

type DefaultContextManager struct {
	mu       sync.RWMutex
	contexts map[string]*DefaultContext
}

type DefaultContext struct {
	mu      sync.RWMutex
	ctxName string
	pools   map[reflect.Type]any
}
