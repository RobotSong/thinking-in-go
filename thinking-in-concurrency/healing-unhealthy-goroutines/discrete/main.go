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

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
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
							log.Printf("get wardHeartbeat")
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
	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case mayStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = mayStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}

				}
			}

		}()
		return valStream
	}

	doWorkFn := func(
		done <-chan interface{},
		intList ...int,
	) (startGoroutineFn, <-chan interface{}) { // 1.
		// 1. Here we’ll take in the values we want our ward to close over, and return any channels our ward will be using to communicate back on.
		intChanStream := make(chan (<-chan interface{})) // 2.
		// 2. This line creates our channel of channels as part of the bridge pattern
		intStream := bridge(done, intChanStream)
		doWork := func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} { // 3.
			// 3. Here we create the closure that will be started and monitored by our steward.
			intStream := make(chan interface{}) // 4.
			// 4. This is where we instantiate the channel we’ll communicate on within this instance of our ward’s goroutine.
			heartbeat := make(chan interface{})

			go func() {

				defer close(intStream)
				select {
				case intChanStream <- intStream: // 5.
				// 5. Here we let the bridge know about the new channel we’ll be communicating on.
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					//for _, intVal := range intList { // 重复发送值
					for {
						// 避免重复发送的值
						intVal := intList[0]
						intList = intList[1:]
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal) // 6.
							// 6. This line simulates an unhealthy ward by logging an error when we encounter a negative number and returning from the goroutine.
							return
						}

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}

				}

			}()

			return heartbeat
		}

		return doWork, intStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5, 8, 9, 10) //. 1
	// 1. This line creates our ward’s function, allowing it to close over our variadic slice of integers, and return a stream that it will communicate back on.
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork) // 2.
	// 2. Here we create our steward that will monitor the doWork closure. Because we expect failures fairly quickly, we’ll set the monitoring period at just one millisecond.
	doWorkWithSteward(done, 1*time.Hour) // 3.
	// 3. Here we tell the steward to start the ward and begin monitoring.
	// 如果需要避免重复发送的，又需要考虑 实际 take 到的值的数量， 需要大于数组
	for intVal := range take(done, intStream, 6) { // 4.
		// 4. Finally, we use one of the pipeline stages we developed and take the first six values from our intStream.
		log.Printf("Received: %v", intVal)
	}

}
