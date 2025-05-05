package memcontext

import "fmt"

var (
	ErrContextNotFound    = fmt.Errorf("context not found")
	ErrPoolAlreadyExists  = fmt.Errorf("pool already exists")
	ErrPoolNotFound       = fmt.Errorf("pool not found")
	ErrPoolNotInitialized = fmt.Errorf("pool not initialized")
	ErrContextManagerNil  = fmt.Errorf("context manager is nil")
)
