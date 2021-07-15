#**Concurrency Patterns In Go**

# Confinement
编写安全的并发代码
When working with concurrent code, there are a few different options for safe operation. We’ve gone over two of them:
• Synchronization primitives for sharing memory (e.g., sync.Mutex)
• Synchronization via communicating (e.g., channels)  
还有两个是隐式并发安全的代码
However, there are a couple of other options that are implicitly safe within multiple
concurrent processes:  
• Immutable data  
• Data protected by confinement  

Confinement can also allow for a lighter cognitive load on the developer and smaller
critical sections.  

Confinement 是什么?  
Confinement is the simple yet powerful idea of ensuring information is only ever
available from one concurrent process. When this is achieved, a concurrent program
is implicitly safe and no synchronization is needed. There are two kinds of confinement possible: ad hoc and lexical.
1. Ad hoc confinement is when you achieve confinement through a convention—whether it be set by the languages community, the group you work within, or the codebase you work within.
就是通过惯例来约束开发人员来编写代码，比如下面的代码:  
   example: ./confinement/ad-hoc  
   We can see that the data slice of integers is available from both the loopData function and the loop over the handleData channel; however, by convention we’re only accessing it from the loopData function.
2. Lexical confinement involves using lexical scope to expose only the correct data and
   concurrency primitives for multiple concurrent processes to use. 
   通过词法来限制开发人员来编写代码，比如通过 channel 的只读、只写，如下代码：  
   example: //confinement/lexical
   
confinement的作用：
   So what’s the point? Why pursue confinement if we have synchronization available to
   us? The answer is improved performance and reduced cognitive load on developers.
   Synchronization comes with a cost, and if you can avoid it you won’t have any critical
   sections, and therefore you won’t have to pay the cost of synchronizing them.
   
# The for-select loop
Something you’ll see over and over again in Go programs is the for-select loop. It’s
nothing more than something like this:
```
for { // Either loop infinitely or range over something
  select {
  // Do some work with channels
  }
}
```
1. Sending iteration variables out on a channel
   Oftentimes you’ll want to convert something that can be iterated over into values on a channel. This is nothing fancy, 
   and usually looks something like this:
```text
for _, s := range []string{"a", "b", "c"} {
  select {
  case <-done:
    return
   case stringStream <- s: 
  }
}
```
2. Looping infinitely waiting to be stopped
   It’s very common to create goroutines that loop infinitely until they’re stopped.
   There are a couple variations of this one. 
   Which one you choose is purely a stylistic preference.  
   1. The first variation keeps the select statement as short as possible:
      If the done channel isn’t closed, we’ll exit the select statement and continue on to the rest of our for loop’s body.
```text
for {
  select {
  case: <- done:
    return
  default:
  }
  // Do non-preemptable work
}
``` 
   ii. The second variation embeds the work in a default clause of the select statement:
   ```text
    If the done channel isn’t closed, we’ll exit the select statement and continue on to the rest of our for loop’s body.
   ```
When we enter the select statement, if the done channel hasn’t been closed, we’ll
execute the default clause instead.

There’s nothing more to this pattern, but it shows up all over the place, and so it’sworth mentioning.

# Preventing Goroutine Leaks
The runtime handles multiplexing the goroutines onto any number of operating system threads 
so that we don’t often have to worry about that level of abstraction. 
But they do cost resources, and goroutines are not garbage collected by the runtime, so
regardless of how small their memory footprint is, we don’t want to leave them lying
about our process.  
Go采用多路复用技术，复用 goroutine 资源，但是不会在运行时回收 goroutine 资源。

The goroutine has a few paths to termination:
• When it has completed its work.
正常结束  
• When it cannot continue its work due to an unrecoverable error.
异常结束  
• When it’s told to stop working.
主动通知结束  

通常是父 goroutine 通知子 goroutine 终止任务。
The parent goroutine (often the main goroutine) with this full contextual knowledge should be
able to tell its child goroutines to terminate.

Let’s start with a simple example of a goroutine leak:
goroutine 泄漏例子:
```text
doWork := func(strings <-chan string) <-chan interface{} {
    completed := make(chan interface{})
    go func() {
      defer fmt.Println("doWork exit")
      defer close(completed)
      
      for s := range strings {
        // Do something interesting
        fmt.Println(s)
      }
    }()
    return completed
}

doWork(nil)
// Perhaps more work is done here.
fmt.Println("Done.")
```

The way to successfully mitigate this is to establish a signal 
between the parent goroutine and its children that allows the parent to signal cancellation to its children. By
convention, this signal is usually a read-only channel named done. 
The parent goroutine passes this channel to the child goroutine and then closes the channel when it
wants to cancel the child goroutine. 
解决这种问题的方案是，在父 goroutine 和子 goroutine 之间建立一个取消信号:
Here’s an example:
```
doWork := func(
  done <-chan interface{},
  strings <-chan string,
) <-chan interface{} {
  terminated := make(chan interface{})
  go func() {
    defer fmt.Println("doWork exit.")
    defer close(terminated)
    
    for {
      select {
      case s := <-strings:
        //Do someting interesting
        fmt.Println(s)
      case <-done:
        return  
      }
    }
  }()
  return terminated
}

done := make(chan interface{})
terminated := doWork(done, nil)

go func() {
  // Cancel the operation after 1 second.
  time.Sleep(1 * time.Second)
  fmt.Println("Canceling doWork goruotine...")
  close(done)
}()

<-terminated
fmt.Println("Done.")
```
如何预防写的时候
What if we’re dealing with the reverse situation: a goroutine blocked on attempting to write a value to a channel? 
```
newRandStream := func() <-chan int {
  randStream := make(chan int)
  go func() { 
    defer fmt.Println("newRandStream closure exited.")
    defer close(randStream)
    for {
      randStream <- rand.Int()
    }
    
  }()
  return randStream
}

randStream := newRandStream()
fmt.Println("3 random ints:")

for i := 1; i <= 3; i++ {
  fmt.Printf("%d: %d\n", i, <-randStream)
}
```

The solution, just like for the receiving case, 
is to provide the producer goroutine with a channel informing it to exit:  
example: ./perventing-groutine-leaks/prevent-write-leaks

# The or-channel
At times you may find yourself wanting to combine one or more done channels into a
single done channel that closes if any of its component channels close. It is perfectly
acceptable, albeit verbose, to write a select statement that performs this coupling;
however, sometimes you can’t know the number of done channels you’re working
with at runtime. In this case, or if you just prefer a one-liner, you can combine these
channels together using the or-channel pattern.
This pattern creates a composite done channel through recursion and goroutines.
Let’s have a look:
把一个或多个已完成的通道组合成一个单一的已完成的通道，如果它的任何一个组成通道关闭，这个通道就会关闭。  
example: ./or-channel-pattern/.

or-channel 模式会额外创建 goroutine . 
We achieve this terseness at the cost of additional goroutines —f(x)=⌊x/2⌋ where x is the number of goroutines.

This pattern is useful to employ at the intersection of modules in your system. 

还可以使用 The context Package 完成这个模式。

# Error Handling
When Go eschewed the popular exception model of errors, it made a statement that error handling was important, 
and that as we develop our programs, we should give our error paths the same attention we give our algorithms.
应该关注错误处理的方式

The most fundamental question when thinking about error handling is,
“Who should be responsible for handling the error?”
At some point, the program needs to stop ferrying the error up the stack and actually do something with it. 

I suggest you separate your concerns: in general, your concurrent processes should send their errors 
to another part of your program that has complete information about the state of your program, 
and can make a more informed decision about what to do.
通过 channel 将错误一起发送到程序的另一部分去：  
example: ./error-handling/send-complete-info

Again, the main takeaway here is that errors should be considered first-class citizens 
when constructing values to return from goroutines. 
错误应该是第一公民。

# Pipelines
A pipeline is just another tool you can use to form an abstraction in your system.

抽象的作用:
Partly to abstract away details that don’t matter to the greater flow, 
and partly so that we can work on one area of code without affecting other areas.  

Pipeline 可以做到:
You can modify stages independent of one another, 
you can mix and match how stages are combined independent of modifying the stages, 
you can process each stage concurrent to upstream or downstream stages, 
and you can fan-out, or rate-limit portions of your pipeline.

As mentioned previously, a stage is just something that takes data in, performs a
transformation on it, and sends the data back out. Here is a function that could be
considered a pipeline stage: 
流水线阶段的 take data in , sends the data back out 例子:
批处理例子:
```
mutiply := func(values []int, mutiplier int) []int {
   mutipliedValues := make([]int, len(values))
   for i, v := range values {
     mutipliedValues[i] = v * mutiplier
   }

   return mutipliedValues
}

add := func(values []int, additive int) []int {
   addValues := make([]int, len(values))
   for i, v := range values {
     addValues[i] = v + additive
   }
   
   return addValues
}

// combining them (add/mutiply)

ints := []int{1, 2, 3}
for _, v := range add(mutiply(ints, 2), 1) {
   fmt.Println(v)
}
```
流处理例子:
```
mutiply := func(value, multiplier int) int {
   return value * multiplier
}

add := func(value, additive int) int {
   return value + additive
}

ints := []int{1, 2, 3}
for _, v := range ints {
   fmt.Println(mutiply(add(mutiply(v, 2), 1), 2))
}

```
What are the properties of a pipeline stage?
1. A stage consumes and returns the same type.
2. A stage must be reified by the language so that it may be passed around. Functions in Go are reified and fit this purpose nicely.

批处理和流处理:  
1. 批处理是Notice how each stage is taking a slice of data and returning a slice of data? These
   stages are performing what we call batch processing. 
2. 流处理是This means that the stage receives and emits one element at a time.

批处理就是一次处理整个数组，流处理是一次处理一个元素。

# Best Practices for Constructing Pipelines
## Use channel primitive
Why use channel?  
Channels are uniquely suited to constructing pipelines in Go 
because they fulfill all of our basic requirements.
They can receive and emit values, 
they can safely be used concurrently, 
they can be ranged over, 
and they are reified by the language. 

## Some Handy Generators

As a reminder, a generator for a pipeline is 
any function that converts a set of discrete values 
into a stream of values on a channel. 
Let’s take a look at a generator called repeat:
```
repeatFn := func(
     done <-chan interface{},
     fn func() interface{},
 ) <-chan interface{} {
     valueStream := make(chan interface{})
     go func() {
         defer close(valueStream)
         for {
             select {
             case <-done:
                 return
             case valueStream <- fn():
             }

         }
     }()
     return valueStream
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
             case takeStream <- <- valueStream:
             }
         }
     }()
     return takeStream
 }

 done := make(chan interface{})
 defer close(done)

 randFn := func() interface{} { return rand.Int() }

 for num := range take(done, repeatFn(done, randFn), 10) {
     fmt.Println(num)
 }

```
Speaking of one stage being computationally expensive, how can we help mitigate this? 
Won’t it rate-limit the entire pipeline?  
For ways to help mitigate this, let’s discuss the fan-out, fan-in technique.

# Fan-out, Fan-in

Fan-out is a term to describe the process of starting multiple goroutines to handle input from the pipeline, 
and fan-in is a term to describe the process of combining multiple results into one channel.

Fan-out 是多个 goroutine 来处理来自管道的输入的过程。  
Fan-in  将多个结果合并到一个通道的过程。

So what makes a stage of a pipeline suited for utilizing this pattern? 
You might consider fanning out one of your stages if both of the following apply:
什么情况下适合这个模式？符合以下两点：
1. It doesn’t rely on values that the stage had calculated before.
2. It takes a long time to run.


# The or-done-channel
Unlike with pipelines, you can’t make any assertions about how a channel will behave
when code you’re working with is canceled via its done channel.
当需要 done channel 的时候，一般会写出像下面的代码：
```
loop:
for {
   select {
   case: <-done:
      break loop
   case maybeVal, ok := <- myChan:
      if ok == false {
         return // or maybe break from for
      }
      // Do something with val
   }
}

```
Continuing with
the theme of utilizing goroutines to write clearer concurrent code, 
and not prematurely optimizing, we can fix this with a single goroutine.
建议使用单独的 goroutine , 而不是过早的优化, 改为如下:
```
orDone := func(done , c <- chan interface{}) <- chan interface{} {
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


// Doing this allows us to get back to simple for loops, like so:
for val := range orDone(done, myChan) {
   // Do something with val
}
```

# The tee-channel
Imagine a channel of user commands: 
you might want to take in a stream of user commands on a channel, send
them to something that executes them, 
and also send them to something that logs the commands for later auditing.

Taking its name from the tee command in Unix-like systems, the tee-channel does
just this. You can pass it a channel to read from, and it will return two separate chan‐
nels that will get the same value:
```
tee := func(
   done <- chan interface{},
   in <- chan interface{},   
) (_, _ chan interface{}) {
   out1 := make(chan interface{})
   out2 := make(chan interface{})
   go func() {
      defer close(out1)
      defer close(out2)
      for val := range orDone(done, in) {
         var out1, out2 = out1, out2
         select {
         case <-done:
         case out1 <- val:
            out1 = nil
         case out2 <- val:
            out2 = nil   
         }
      }
   }()
   return out1, out2
}
```

# The bridge-channel
As a consumer, the code may not care about the fact that its values come from a
sequence of channels. In that case, dealing with a channel of channels can be cumbersome. 
If we instead define a function that can destructure the channel of channels
into a simple channel—a technique called bridging the channels—this will make it
much easier for the consumer to focus on the problem at hand. Here’s how we can
achieve that:
```go
package main

import "fmt"

func main() {

	orDone := func(done , c <- chan interface{}) <- chan interface{} {
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
	
	bridge := func(
		done <- chan interface{},
		chanStream <-chan <-chan interface{},
		) <- chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case mayStream, ok := <- chanStream:
					if ok == false {
						return
					}
					stream = mayStream
				case <- done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <- done:
					}
					
				}
			}

		}()
		return valStream
	}

	genVal := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 20; i++ {
				// chan 容量为1 很重要
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVal()) {
		fmt.Printf("%v ", v)
	}
}

```

# Queuing
Sometimes it’s useful to begin accepting work for your pipeline even though the pipeline is not yet ready for more. 
This process is called queuing.

Let’s begin by analyzing situations in which queuing can increase the overall performance of your system. 
The only applicable situations are: 队列可以提高性能的情况:
* If batching requests in a stage saves time.
* If delays in a stage produce a feedback loop into the system.

```
done := make(chan interface{})
defer close(done)
zeros := take(done, 3, repeat(done, 0))
short := sleep(done, 1*time.Second, zeros)
buffer := buffer(done, 2, short) // Buffers sends from short by 2
long := sleep(done, 4*time.Second, short)
pipeline := long
```

So from our examples we can begin to see a pattern emerge; queuing should be
implemented either:
* At the entrance to your pipeline.
* In stages where batching will lead to higher efficiency.


# The context package 
The context package serves two primary purpose:
* To provide an API for canceling branches of your call-graph.
* To provide a data-bag for transporting request-scope data through your call-graph.

Here is a program that concurrently prints a greeting and a farewell:  
path: 
