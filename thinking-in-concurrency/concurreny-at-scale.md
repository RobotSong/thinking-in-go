
# Error Propagation

参考 Error 封装传递项目地址: https://github.com/pkg/errors

Errors indicate that your system has entered a state in which it cannot fulfill an operation that a user either explicitly or implicitly requested. Because of this, it needs to
relay a few pieces of critical information 错误应该包含以下几个关键点 :
* What happened.  
  This is the part of the error that contains information about what happened, e.g.,
  “disk full,” “socket closed,” or “credentials expired.” This information is likely to
  be generated implicitly by whatever it was that generated the errors, although you
  can probably decorate this with some context that will help the user.
* When and where is occurred.  
  Errors should always contain a complete stack trace starting with how the call
  was initiated and ending with where the error was instantiated. The stack trace
  should not be contained in the error message (more on this in a bit), but should
  be easily accessible when handling the error up the stack.
  Further, the error should contain information regarding the context it’s running
  within. For example, in a distributed system, it should have some way of identify‐
  ing what machine the error occurred on. Later, when trying to understand what
  happened in your system, this information will be invaluable.
  In addition, the error should contain the time on the machine the error was
  instantiated on, in UTC
* A friendly user-facing message.  
  The message that gets displayed to the user should be customized to suit your
  system and its users. It should only contain abbreviated and relevant information
  from the previous two points. A friendly message is human-centric, gives some
  indication of whether the issue is transitory, and should be about one line of text.
* How the user can get more information.  
  At some point, someone will likely want to know, in detail, what happened when
  the error occurred. Errors that are presented to users should provide an ID that
  can be cross-referenced to a corresponding log that displays the full information
  of the error: time the error occurred (not the time the error was logged), the
  stack trace—everything you stuffed into the error when it was created. It can also
  be helpful to include a hash of the stack trace to aid in aggregating like issues in
  bug trackers
  
如何增加包装错误信息:  
Note that it is only necessary to wrap errors in this fashion at your own module boundaries—public functions/methods—or 
when your code can add valuable context.
```
func PostReport(id string) error {
  result, error := lowlevel.DoWork()
  if err != nil {
    if _, ok := err.(lowlevel.Error); ok {
      err = WrapErr(err, "cannot post report with id %q", id)
    }
    return err
  }
}
```

Let’s take a look at a complete example.  
path: ./error-propagation

# Timeouts and Cancellation

When working with concurrent code, timeouts and cancellations are going to turn up
frequently.As we’ll see in this section, among other things, timeouts are crucial to
creating a system with behavior you can understand. Cancellation is one natural
response to a timeout. We’ll also explore other reasons a concurrent process might be
canceled.

So what are the reasons we might want our concurrent processes to support timeouts? 
Here are a few:为什么需要超时:  
* System saturation
* Stale data
* Attempting to prevent deadlocks

并发取消的原因:
There are a number of reasons why a concurrent process might be canceled:
* Timeouts
* User intervention
* Parent cancellation
* Replicated requests

# Heartbeats

Heartbeats are a way for concurrent processes to signal life to outside parties.

There are two different types of heartbeats we’ll discuss in this section:
* Heartbeats that occur on a time interval.
* Heartbeats that occur at the beginning of a unit of work.

查看两个 时间间隔(time interval) 的心跳检查的代码例子:
正常心跳: ./heartbeats-normal
出现异常心跳: ./heartbeats-panic

一个开始工作时, 发送心跳的 代码例子:  
path: ./heartbeats-begin-work

