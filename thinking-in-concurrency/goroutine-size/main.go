// 粗略的计算 goroutine 大小
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// 1. 设置一个 死锁 函数
	var c chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }
	// 2. 生成多个 goroutine
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	// 3. 未创建多个 goroutine 前的内存使用
	before := memConsumed()
	// 4. 启动
	for i := 0; i < numGoroutines; i++ {
		go noop()
	}
	// 5. 等待创建完成
	wg.Wait()
	after := memConsumed()
	fmt.Printf("now use %3fkb\n", float64(after-before)/1000)
	fmt.Printf("%3fkb\n", float64(after-before)/numGoroutines/1000)
}

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
