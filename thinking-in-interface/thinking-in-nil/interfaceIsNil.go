package main

import "fmt"

// main 函数展示了如何检查接口变量是否为 nil。
// 在 Go 中，即使接口内部的值是 nil，接口变量本身也可能不是 nil。
func main() {
	var x interface{} = nil // x 是一个 nil 接口值
	var y *int = nil       // y 是一个指向 int 的 nil 指针
	interfaceIsNil(x)      // 调用函数检查 x 是否为 nil 接口
	interfaceIsNil(y)      // 调用函数检查 y 是否为 nil 接口
}

// interfaceIsNil 函数用于检查传入的接口变量是否为 nil。
// 如果接口变量为 nil，则打印 "empty interface"。
// 如果接口变量内部的值是 nil 但接口本身不是 nil，则打印 "non-empty interface"。
// 注意：在 Go 中，一个接口变量只有在类型和值都为 nil 时才等于 nil。
func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

