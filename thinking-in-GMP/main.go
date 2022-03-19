package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{} = []int{2, 6}
	fmt.Println("data 的数据类型是:", reflect.TypeOf(data))
	data = [2]int{2, 6}
	fmt.Println("data 的数据类型是:", reflect.TypeOf(data))
	data = [...]int{2, 6}
	fmt.Println("data 的数据类型是:", reflect.TypeOf(data))
}
