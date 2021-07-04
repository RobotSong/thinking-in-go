package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return mem
		},
	}
	// Seed the pool with 4KB
	calcPool.Put(calcPool.New)
	calcPool.Put(calcPool.New)
	calcPool.Put(calcPool.New)
	calcPool.Put(calcPool.New)

	var numWorker = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorker)
	for i := numWorker; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get()
			// 如果加上睡眠时间, 超过 62089 calculators were created.
			//time.Sleep(1 * time.Millisecond)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
