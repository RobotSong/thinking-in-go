package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 调用 Done 函数
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	// 2. 添加 Add
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)

	// 3. 执行 foreach ，等待所有的 goroutine 完成
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
