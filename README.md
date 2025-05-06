# Memora

[![GoDoc](https://pkg.go.dev/badge/github.com/AlexsanderHamir/Memora)](https://pkg.go.dev/github.com/AlexsanderHamir/Memora)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlexsanderHamir/Memora)](https://goreportcard.com/report/github.com/AlexsanderHamir/Memora)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

An object pool manager to facilitate the allocation and deallocation of memory.

## Overview

Memora provides a simple way to manage object pools with context-aware resource allocation and cleanup. It's designed to optimize memory usage and improve performance in applications that frequently create and destroy related objects.

For more details about the underlying pool implementation, check out [PoolX](https://github.com/AlexsanderHamir/PoolX).

## Features

- Context-aware object pool management
- Generic type support
- Custom allocators and cleaners
- Thread-safe operations

## Usage

For a complete working example, see `code_example.go`.

```go
import (
	"github.com/AlexsanderHamir/Memora/memcontext"
	"github.com/AlexsanderHamir/PoolX/src/pool"
)

// Create a context manager
cm := memcontext.NewContextManager()

// Create a context
ctx := cm.CreateContext("myContext")

// pool configuration
config, err := pool.NewPoolConfigBuilder().Build()
	if err != nil {
		panic(err)
}

// Create a pool with custom allocator and cleaner
pool, err := memcontext.CreatePool(ctx, config, allocator, cleaner)

// Get an object from the pool
obj, err := pool.Get()

// Use the object
// ...

// Return the object to the pool
pool.Put(obj)
```

## API

### ContextManager Interface

The ContextManager interface provides methods to manage object pool contexts:

```go
type ContextManager interface {
    // CreateContext creates a new context with the given name
    CreateContext(name string) *DefaultContext

    // GetContext retrieves an existing context by name
    // Returns an error if the context doesn't exist
    GetContext(name string) (*DefaultContext, error)

    // GetOrCreateContext retrieves an existing context or creates a new one if it doesn't exist
    GetOrCreateContext(name string) (*DefaultContext, error)

    // DeleteContext removes a context and its associated pools
    DeleteContext(name string)
}
```

### Context Methods

The DefaultContext provides methods to manage object pools within a context:

```go
// CreatePool creates a new pool in the context for a specific type
func CreatePool[T any](dc *DefaultContext, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error)

// GetPool retrieves an existing pool from the context
// Returns an error if the pool doesn't exist
func GetPool[T any](dc *DefaultContext) (*pool.Pool[T], error)

// GetOrCreatePool retrieves an existing pool or creates a new one if it doesn't exist
func GetOrCreatePool[T any](dc *DefaultContext, config *pool.PoolConfig, allocator func() T, cleaner func(T)) (*pool.Pool[T], error)

// DeletePool removes a specific pool from the context
func DeletePool[T any](dc *DefaultContext)

// ClosePools closes all pools in the context and removes them
func ClosePools(dc *DefaultContext)
```

## Installation

```bash
go get github.com/AlexsanderHamir/Memora
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
