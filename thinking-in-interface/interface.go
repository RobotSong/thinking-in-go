package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {

	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "???"
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	b := Binary(200)
	var s = Stringer(b)
	fmt.Println(s)
	b = Binary(5)
	fmt.Println(s)
	switch s.(type) {
	case Stringer:
		return
	default:
		panic("ds")
	}

}
