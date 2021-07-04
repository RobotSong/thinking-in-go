package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitTime = 0

func main() {

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(1)
	var now = time.Now()
	go func() {
		elapsed := time.Since(now)
		fmt.Println("Done, 耗时:", elapsed)
		wg.Done()
	}()

	go wait()

	wg.Wait()

	fmt.Println()
}

func wait() {
	// Wait: 114757 , 最终的执行次数 不一定
	waitTime++
	fmt.Println("Wait:", waitTime)
	go wait()
}
