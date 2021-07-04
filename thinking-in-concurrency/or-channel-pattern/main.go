package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} { // 1. Here we have our function, or, which takes
		// in a variadic slice of channels and returns a single channel.
		switch len(channels) {
		case 0: // 2. Since this is a recursive function, we must set up termination criteria.
			// The first is that if the variadic slice is empty, we simply return a nil channel.
			// This is consistant with the idea of passing in no channels;
			// we wouldn’t expect a composite channel to do anything.
			return nil
		case 1: // 3. Our second termination criteria states that if our variadic slice only contains one element,
			// we just return that element.
			return channels[0]
		}
		orDone := make(chan interface{})
		go func() { // 4.Here is the main body of the function, and where the recursion happens.
			// We create a goroutine so that we can wait for messages on our channels without blocking.
			defer close(orDone)
			switch len(channels) {
			case 2: // 5. Because of how we’re recursing, every recursive call to or will at least have two channels.
				// As an optimization to keep the number of goroutines constrained, we place a special case here
				// for calls to or with only two channels.
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				// 三个意外的情况
				case <-or(append(channels[3:], orDone)...): // 6.Here we recursively create an or-channel from
					// all the channels in our slice after the third index, and then select from this.
					// This recurrence relation will destructure the rest of the slice into or-channels to form a tree
					// from which the first signal will return. We also pass in the orDone channel
					// so that when goroutines up the tree exit, goroutines down the tree also exit.
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Minute),
		sig(1*time.Second),
		sig(6*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}
