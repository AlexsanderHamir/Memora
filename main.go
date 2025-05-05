package main

import (
	"fmt"

	"github.com/AlexsanderHamir/Memora/memcontext"
	"github.com/AlexsanderHamir/PoolX/src/pool"
)

type Example struct {
	ID int
}

type Example2 struct {
	ID int
}

var allocator2 = func() *Example2 {
	return &Example2{
		ID: 2222,
	}
}

var cleaner2 = func(item *Example2) {
	item.ID = 0
}

var allocator = func() *Example {
	return &Example{
		ID: 1299229,
	}
}

var cleaner = func(item *Example) {
	item.ID = 0
}

func main() {
	ctxManager := memcontext.NewContextManager()
	ctx := memcontext.CreateContext(ctxManager, "test")

	poolConfig, err := pool.NewPoolConfigBuilder().Build()
	if err != nil {
		panic(err)
	}

	pool1, err := memcontext.CreatePool(ctx, "test", poolConfig, allocator, cleaner)
	if err != nil {
		panic(err)
	}

	obj, err := pool1.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("pool1", obj)

	poolInstance, err := memcontext.GetPool[*Example](ctx, "test")
	if err != nil {
		panic(err)
	}

	obj2, err := poolInstance.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("poolInstance", obj2)
}
