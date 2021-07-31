# Appendix 附录
Go 相关工具的使用
This appendix will discuss some of these tools and how they can aid you before, during, and after development.

## Anatomy of a Goroutine Error

For example, when this simple program is executed:
```go
package main

func main() {
	waitFoever := make(chan interface{})
	go func() {
		panic("test panic")
	}()
	<- waitFoever
}

```
The following stack trace is produced:
```text
panic: test panic

goroutine 5 [running]:
main.main.func1() 3
	E:/_self/go/thinking-in-go/thinking-in-concurrency/temp-main-2.go:6 +0x40  1
created by main.main
	E:/_self/go/thinking-in-go/thinking-in-concurrency/temp-main-2.go:5 +0x5f  2
```
1. Refers to where the panic occurred.  
2. Refers to where the goroutine was started.
3. Indicates the name of the function running as a goroutine. If it’s an anonymous
   function as in this example, an automatic and unique identifier is assigned.

If you’d like to see the stack traces of all the goroutines that were executing when the program panicked, you can enable the old behavior by setting the GOTRACEBACK environmental variable to all.

## Race Detection
In Go 1.1, a -race flag was added as a flag for most go commands:
* $ go test -race mypkg # test the package
* $ go run -race mysrc.go # compile and run the program
* $ go build -race mycmd # build the command
* $ go install -race mypkg # install the package

One caveat of using the race detector is that the
algorithm will only find races that are contained in code that is exercised. 
推荐在生产上使用
For this reason, the Go team recommends running a build of your application built with the
race flag under real-world load. 

For example:
```go
package main

import "fmt"

var data int
func main() {
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
```
使用 -race 命令
```text
go run -race temp-main-2.go
```
堆栈输出
```text
the value is 0.
==================
WARNING: DATA RACE
Write at 0x000000641ea8 by goroutine 7:
  main.main.func1()
      E:/_self/go/thinking-in-go/thinking-in-concurrency/temp-main-2.go:8 +0x5d  1.

Previous read at 0x000000641ea8 by main goroutine:
  main.main()
      E:/_self/go/thinking-in-go/thinking-in-concurrency/temp-main-2.go:11 +0x5d 2.

Goroutine 7 (running) created at:
  main.main()
      E:/_self/go/thinking-in-go/thinking-in-concurrency/temp-main-2.go:7 +0x4d
==================
Found 1 data race(s)

```

1. Signifies a goroutine that is attempting to write unsynchronized memory access.
2. Signifies a goroutine (in this case the main goroutine) trying to read this same
   memory
   
## pprof
In large codebases, it can sometimes be difficult to ascertain how your program is
performing at runtime. How many goroutines are running? Are your CPUs being
fully utilized? How’s memory usage doing? Profiling is a great way to answer these
questions, and Go has a package in the standard library to support a profiler named
“pprof.

The runtime/pprof package is pretty simple, and has predefined profiles to hook into
and display:
* goroutine - stack traces of all current goroutines
* heap - a sampling of all heap allocations
* threadcreate - stack traces that led to the creation of new OS threads
* block - stack traces that led to blocking on synchronization primitives
* mutex - stack traces of holders of contended mutexes

For example, here’s a goroutine that can help you detect goroutine leaks:
```go
package main

import (
	"log"
	"runtime/pprof"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.LUTC)

	// Every second, log how many goroutines are currently running.
	go func() {
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1 * time.Second) {
			log.Printf("goroutine count:%d\n", goroutines.Count())
		}
	}()

	// Create some goroutines which will never exit.
	var blockForever chan struct{}
	for i := 0; i < 10; i++ {
		go func() { <-blockForever }()
		time.Sleep(500 * time.Millisecond)
	}

}
```
输出
```text
10:28:31 goroutine count:5
10:28:32 goroutine count:6
10:28:33 goroutine count:8
10:28:34 goroutine count:10
10:28:35 goroutine count:12
```

These built-in profiles can really help you profile and diagnose issues with your pro‐
gram, but of course you can write custom profiles tailored to help you monitor your
programs:

```go
package main

import (
   "log"
   "runtime/pprof"
)

func newProfIfNotDef(name string) *pprof.Profile {
   prof := pprof.Lookup(name)
   if prof == nil {
      prof = pprof.NewProfile(name)
   }
   return prof
}

func main() {
   prof := newProfIfNotDef("my_package_namespace")
   log.Printf("%v", prof)
}
```