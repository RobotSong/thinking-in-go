package main

import (
	"context"
	"fmt"
	"runtime/debug"
)

type foo int
type bar int

type LowLevelErr struct {
	error
}

func main() {
	c := context.Background()
	fmt.Println(&c)

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(2)] = 2
	fmt.Printf("%v\n", m)
	fmt.Println(m[bar(1)])
	le := LowLevelErr{}
	fmt.Printf("error: %v", le.error)
	debug.Stack()
}
