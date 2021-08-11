package main

import (
	"fmt"
	"reflect"
)

func findRightSmall(arr []int) []int {
	ans := make([]int, len(arr))
	stack := []int{}
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		// 每个元素都向左遍历栈中的元素完成消除动作
		sl := len(stack)
		for sl > 0 && arr[stack[sl-1]] > x {
			// 消除的时候，记录一下被谁消除了
			ans[stack[sl-1]] = i
			// 消除时候，值更大的需要从栈中消失
			stack = stack[:sl-1]
			sl = len(stack)
		}
		// 剩下的入栈
		stack = append(stack, i)
	}

	// 栈中剩下的元素，由于没有人能消除他们，因此，只能将结果设置为-1。
	for len(stack) > 0 {
		ans[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	return ans
}

func main() {
	data := []struct {
		arr  []int
		want []int
	}{
		{arr: []int{5, 4}, want: []int{1, -1}},
		{arr: []int{1, 2, 4, 9, 4, 0, 5}, want: []int{5, 5, 5, 4, 5, -1, -1}},
	}

	for _, d := range data {
		if got := findRightSmall(d.arr); !reflect.DeepEqual(got, d.want) {
			fmt.Printf("arr: %v -- got: %v, want: %v\n", d.arr, got, d.want)
		}
	}

}
