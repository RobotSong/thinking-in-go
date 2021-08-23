package main

import (
	"container/heap"
	"fmt"
)

type apple struct {
	// 还剩 c 个苹果
	c int
	// 还剩 d 天苹果腐烂
	d int
}

type hp []apple

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.d < b.d || a.d == b.d && a.c < b.c
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(apple))
}

func (h *hp) Pop() interface{} {
	// 取出最后一位，就是为最小值
	// 这里要循环处理, 每个元素都要将所有值的 day - 1， 如果 day == 0，则需要把这个元素去掉
	// 首尾的 c - 1, 如果 c == 0 , 则需要把这个元素去掉
	a := *h
	v := a[h.Len()-1]
	*h = a[:h.Len()-1]
	return v
}

// First 最大/最小 的值, 取出来后，不一定需要移除
func (h *hp) First() *apple {
	a := *h
	return &a[0]
}

func eatenApples(apples []int, days []int) int {
	// 吃的苹果数量
	var count int
	h := &hp{}
	for i := 0; i < len(apples) || h.Len() > 0; i++ {
		// 移除过期的
		for h.Len() > 0 {
			a := h.First()
			if a.d <= i {
				// 过期了，移除
				heap.Pop(h)
			} else {
				// 不存在过期
				break
			}
		}
		if i < len(apples) && apples[i] > 0 {
			a := apple{
				c: apples[i],
				d: days[i] + i,
			}
			heap.Push(h, a)
		}
		// 如果 h 为空, 则没有苹果可以吃
		if h.Len() == 0 {
			continue
		}
		// 当前还有苹果可以吃
		v := h.First()
		v.c--
		// 苹果吃完了,去掉它
		if v.c == 0 {
			heap.Pop(h)
		}
		count++
	}

	return count
}

func main() {
	// 根据 apples 和 days 组成一个最小的 days 作为 最小堆
	apples := []int{3, 0, 0, 0, 0, 2}
	days := []int{3, 0, 0, 0, 0, 2}
	// 预期结果 7
	ans := eatenApples(apples, days)
	fmt.Println(ans)
}
