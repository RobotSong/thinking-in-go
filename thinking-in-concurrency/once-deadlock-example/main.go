package main

import "sync"

// fatal error: all goroutines are asleep - deadlock!
func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
}
