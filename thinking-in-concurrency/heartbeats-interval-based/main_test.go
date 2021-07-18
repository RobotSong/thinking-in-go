package main_test

import (
	"fmt"
	"testing"
	"time"
)

func DoWork(
	done <-chan interface{},
	pulseInterval time.Duration,
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(pulseInterval)
	numLoops:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
						fmt.Println("send pulse")
					default:
					}
				case intStream <- n:
					fmt.Printf("send val: %v\n", n)
					continue numLoops
				}
			}
		}
	}()
	return heartbeat, intStream
}

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	const timeout = 2 * time.Second
	heartbeat, results := DoWork(done, timeout/2, intSlice...)

	<-heartbeat

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf(
					"index %v: expected: %v, but received %v, ",
					i,
					expected,
					r,
				)
			}
			i++
		case <-heartbeat:
		case <-time.After(timeout):
			t.Fatal("test timed out")

		}
	}
}
