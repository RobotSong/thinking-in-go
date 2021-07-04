package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	// Notice how the loop doesn’t need an exit criteria,
	// and the range does not return the second boolean value.
	for i := range intStream {
		fmt.Printf("%v ", i)
	}
}
