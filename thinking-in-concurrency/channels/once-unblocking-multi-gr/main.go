// 一次解锁多个 goroutine 例子
package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has been begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...", cap(begin))
	close(begin)
	wg.Wait()
}
