package main

import (
	"fmt"
	"math/rand"
)

func main() {

	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartStream := make(chan interface{}, 1)
		workStream := make(chan int)
		go func() {
			defer close(heartStream)
			defer close(workStream)
			for i := 0; i < 10; i++ {
				select {
				case heartStream <- struct{}{}:
				default:
				}

				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}
		}()
		return heartStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeats, results := doWork(done)

	for {
		select {
		case _, ok := <-heartbeats:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}

	}
}
