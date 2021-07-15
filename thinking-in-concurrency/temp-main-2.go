package main

import (
	"context"
	"fmt"
)

type foo int
type bar int

func main() {
	c := context.Background()
	fmt.Println(&c)

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(2)] = 2
	fmt.Printf("%v\n", m)
	fmt.Println(m[bar(1)])
}
