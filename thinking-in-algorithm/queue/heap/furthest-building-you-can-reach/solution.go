package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[h.Len()-1]
	*h = a[:h.Len()-1]
	return v
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	if heights == nil || len(heights) == 0 {
		return 01
	}

	// 记录 需要跳跃的: 最大堆
	h := &hp{}
	// 当前已使用的 砖块 总数
	qSum := 0
	// 最后跳跃到的位置
	lastPos := 0
	// 当前脚下的高度
	preHeight := heights[0]
	for i := 1; i < len(heights); i++ {
		// 当前要往下跳的 高度
		curHeight := heights[i]
		// 如果是从高处往地处跳
		if preHeight > curHeight {
			lastPos = i
			// 更新位置的高度
			preHeight = curHeight
			continue
		}
		// 高度差
		delta := curHeight - preHeight
		// 记录使用的砖块
		heap.Push(h, delta)
		qSum += delta
		// 如果已经使用砖块总和 比 可使用的多了
		// 而且还有梯子，则替代最大的高度差
		for qSum > bricks && ladders > 0 {
			// 最大的高度差
			max := heap.Pop(h).(int)
			// 减去使用的砖块
			qSum -= max
			// 使用梯子
			ladders--
		}
		// 使用梯子后，还能跳到下一个位置
		if qSum <= bricks {
			lastPos = i
		} else {
			break
		}
		// 更新位置的高度
		preHeight = curHeight
	}

	return lastPos
}

func main() {
	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks := 10
	ladders := 2
	res := furthestBuilding(heights, bricks, ladders)
	fmt.Println(res)
}
