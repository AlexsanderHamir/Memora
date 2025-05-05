package memcontext

import (
	"sync"
)

type DefaultContextManager struct {
	mu       sync.RWMutex
	contexts map[string]*DefaultContext
}

type DefaultContext struct {
	mu      sync.RWMutex
	ctxName string
	pools   map[string]any
}
