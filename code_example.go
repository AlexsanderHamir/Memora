package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/AlexsanderHamir/Memora/memcontext"
	"github.com/AlexsanderHamir/PoolX/pool"
)

type Tuple struct {
	ID    int
	Value string
	Data  []byte
}

var (
	tupleAllocator = func() *Tuple {
		return &Tuple{ID: 1299229, Value: "test", Data: []byte("test")}
	}
	tupleCleaner = func(t *Tuple) {
		t.ID = 0
	}

	readerAllocator = func() *bytes.Reader {
		return bytes.NewReader([]byte("test"))
	}
	readerCleaner = func(r *bytes.Reader) {
		r.Reset([]byte("test"))
	}

	bufferAllocator = func() *bytes.Buffer {
		return bytes.NewBuffer(nil)
	}
	bufferCleaner = func(b *bytes.Buffer) {
		b.Reset()
	}
)

type RealWorld struct {
	cm memcontext.ContextManager
}

func main() {
	rw := &RealWorld{}
	rw.cm = memcontext.NewContextManager()

	ctx := rw.cm.CreateContext("row_group_example")

	// Create object pools
	tuplePool := mustCreatePool(ctx, tupleAllocator, tupleCleaner)
	readerPool := mustCreatePool(ctx, readerAllocator, readerCleaner)
	bufferPool := mustCreatePool(ctx, bufferAllocator, bufferCleaner)

	// Acquire pooled objects
	tuple := mustGetObj(tuplePool)
	reader := mustGetObj(readerPool)
	buffer := mustGetObj(bufferPool)

	// Serialize tuple into buffer
	buffer.WriteString(fmt.Sprintf("ID=%d;Value=%s;Data=%s", tuple.ID, tuple.Value, string(tuple.Data)))

	// Load buffer content into reader and read
	reader.Reset(buffer.Bytes())
	readBuf := make([]byte, buffer.Len())
	_, _ = reader.Read(readBuf)

	fmt.Println("Serialized Tuple:", string(readBuf))

	// Simulate processing delay
	time.Sleep(100 * time.Millisecond)

	// if want to use the obj again use the cleaner function
	tupleCleaner(tuple)
	readerCleaner(reader)
	bufferCleaner(buffer)

	// Return objects to their pools
	// The cleaner is called inside the put method, no need to call it before.
	tuplePool.Put(tuple)
	readerPool.Put(reader)
	bufferPool.Put(buffer)

	rw.cm.DeleteAllContexts()
}

// mustCreatePool creates a pool or panics on error.
func mustCreatePool[T any](ctx *memcontext.DefaultContext, alloc func() T, clean func(T)) pool.PoolObj[T] {
	config, err := pool.NewPoolConfigBuilder().Build()
	if err != nil {
		panic(err)
	}

	p, err := memcontext.CreatePool(ctx, config, alloc, clean)
	if err != nil {
		panic(err)
	}
	return p
}

// mustGet acquires an object from the pool or panics on error.
func mustGetObj[T any](p pool.PoolObj[T]) T {
	obj, err := p.Get()
	if err != nil {
		panic(err)
	}
	return obj
}
