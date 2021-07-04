# understanding the interface
## What is an interface?
```text
In object-oriented programing, a protocol or interface is a common means for **unrelated objects** 
to **communicate** with each other.
    -- wikipedia
```
## What is a Go interface?
```text
abstract types
-------------
concrete types

**concrete types in Go**
    - they describe a memory layout
    - behavior attached to data through methods
        type Number int 
        func (n Number) Positive() bool {
            return n > 0
        }
        
**abstract types in Go**
    - they describe behavior
    - they define a set of methods, without specifying the receiver
        type Positiver interface {
            Positive() bool
        }
    
```

## interface{} says _nothing_

## Why do we use interfaces?
### - generic algorithms
### - hide implementation details
### - providing interception points

**The bigger the interface, the weaker the abstraction**
-- Rob Pike in his Go Proverbs

**Be conservative in what you do, be liberal in what you accept from others**
-- Robustness Principle

## Abstract Data Types
### Mathematical model for data types 
### Defined by its behavior in terms of:
### - possible values,
### - possible operations on data of this type,
### - and the behavior of these operations
## Example: stack ADT
## a Stack interface
```go
package stack

type Stack interface {
    Push(v interface{}) Stack
    Pop() Stack
    Empty() bool
}
```

##  In conclusion
### Interfaces provide:
#### - generic algorithms
#### - hidden implementation
#### - interceptions points
### Implicit satisfaction:
#### - break dependencies
### Type assertions:
#### - to extend behaviors
#### - to classify errors -> ./classify-errors.go
#### - to maintain compatibility -> ./maintain-compatibility.go
