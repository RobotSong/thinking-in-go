package main

import (
	"log"
	"os"
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
	// 1. Here we define the signature of a goroutine that can be monitored and restarted.
	// We see the familiar done channel, and pulseInterval and heartbeat from the
	// heartbeat pattern.
	type startGoroutineFn func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (heartbeat <-chan interface{}) // 1.

	newSteward := func(
		timeout time.Duration,
		startGoroutine startGoroutineFn,
	) startGoroutineFn { // 2.
		// 2. On this line we see that a steward takes in a timeout for the goroutine it will be
		// monitoring, and a function, startGoroutine, to start the goroutine it’s monitor‐
		// ing. Interestingly, the steward itself returns a startGoroutineFn indicating that
		// the steward itself is also monitorable.
		return func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}
				startWard := func() { // 3.
					// 3. Here we define a closure that encodes a consistent way to start the goroutine
					//	we’re monitoring.
					wardDone = make(chan interface{}) // 4.
					// 4.This is where we create a new channel that we’ll pass into the ward goroutine in
					// case we need to signal that it should halt.

					wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) // 5.
					// 5.Here we start the goroutine we’ll be monitoring. We want the ward goroutine to
					//	halt if either the steward is halted, or the steward wants to halt the ward gorou‐
					//	tine, so we wrap both done channels in a logical-or. The pulseInterval we pass
					//	in is half of the timeout period, although as we discussed in “Heartbeats” on page
					//	161, this can be tweaked.
				}
				startWard()
				pulse := time.Tick(pulseInterval)
			monitorLoop:
				for {
					timeoutSingal := time.After(timeout)

					for { // 6.
						// 6.This is our inner loop, which ensures that the steward can send out pulses of its
						//	own.
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat: // 7.
							// 7.Here we see that if we receive the ward’s pulse, we continue our monitoring loop.
							continue monitorLoop
						case <-timeoutSingal: // 8.
							// 8. This line indicates that if we don’t receive a pulse from the ward within our time‐
							//	out period, we request that the ward halt and we begin a new ward goroutine. We
							//	then continue monitoring
							log.Println("steward: ward unhealthy; restarting")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()
			return heartbeat
		}
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			<-done // 1.
			// 1. Here we see that this goroutine isn’t doing anything but waiting to be canceled.
			//	It’s also not sending out any pulses.
			log.Println("ward: I am halting.")
		}()
		return nil
	}

	doWorkWithSteward := newSteward(4*time.Second, doWork) // 2.
	// 2. This line creates a function that will create a steward for the goroutine doWork
	// starts. We set the timeout for doWork at four seconds.

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() { // 3.
		// 3. Here we halt the steward and its ward after nine seconds so that our example will end.
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) {
	} // 4.
	// 4. Finally, we start the steward and range over its pulses to prevent our example from halting
	log.Println("Done")
}
