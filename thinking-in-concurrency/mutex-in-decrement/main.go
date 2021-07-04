package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 定义 增减的变量、mutex
	var count int
	var lock sync.Mutex
	// 2. 定义自增、自减函数
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// 3. 启动 goroutine 实现自增、自减
	var arithmetic sync.WaitGroup
	// Increment
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	// Decrement
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	// 4. 等待完成
	arithmetic.Wait()
	// 5. 打印完成
	fmt.Println("Arithmetic Complete.")
}
