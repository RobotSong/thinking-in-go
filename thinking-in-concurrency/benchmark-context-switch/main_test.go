// Linux 下测试线程 上下文切换的命令
// taskset -c 0 perf bench sched pipe -T
//This produces:
//# Running 'sched/pipe' benchmark:
//# Executed 1000000 pipe operations between two threads
//Total time: 2.935 [sec]
//2.935784 usecs/op
//340624 ops/sec
//This benchmark actually measures the time it takes to send and receive a message on
//a thread, so we’ll take the result and divide it by two. That gives us 1.467 μs per con‐
//text switch. That doesn’t seem too bad, but let’s reserve judgment until we examine
//context switches between goroutines.
package main

import (
	"sync"
	"testing"
)

// BenchmarkContextSwitch   7591803               157 ns/op
// PASS
// ok      cyg.com/thinking-in-go/thinking-in-concurrency/benckmark-context-switch 1.640s
// 测试两个 goroutine 上下文切换的时间
func BenchmarkContextSwitch(b *testing.B) {
	// 1. 添加一个 发送/接收 的通道
	// 一个预备同时开始跑的 通道
	// 等待结束的 waitGroup
	var wg sync.WaitGroup
	begin := make(chan interface{})
	c := make(chan interface{})
	// 2. 发送 和 接收 benchmark 函数
	var token struct{}
	sender := func() {
		defer wg.Done()
		// 等待开始标识
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	// 3. 等待结束
	wg.Wait()

}
