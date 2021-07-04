package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	salutation := "hello"
	a := 1
	go func() {
		defer wg.Done()
		salutation = "welcome"
		a = 2
	}()
	wg.Wait()
	// It turns out that goroutines execute within the same address space they
	// were created in, and so our program prints out the word “welcome.”
	fmt.Println(salutation) // welcome
	fmt.Println(a)          // 2

}
