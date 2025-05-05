package memcontext

import (
	"sync"
)

type DefaultContextManager struct {
	contexts sync.Map
}

type DefaultContext struct {
	name  string
	pools sync.Map
}
