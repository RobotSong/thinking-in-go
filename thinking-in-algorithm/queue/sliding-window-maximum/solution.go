package main

import "fmt"

type ArrayDeque struct {
	a []int
}

func (d *ArrayDeque) getLast() int {
	if d.isEmpty() {
		return -1
	}

	return d.a[len(d.a)-1]
}

func (d *ArrayDeque) isEmpty() bool {
	return len(d.a) == 0
}

func (d *ArrayDeque) getFirst() int {
	if d.isEmpty() {
		return -1
	}

	return d.a[0]
}

func (d *ArrayDeque) addLast(val int) bool {
	d.a = append(d.a, val)
	return true
}

func (d *ArrayDeque) removeFirst() bool {
	if d.isEmpty() {
		return false
	}

	d.a = d.a[1:]
	return true
}

func (d *ArrayDeque) removeLast() bool {
	if d.isEmpty() {
		return false
	}

	d.a = d.a[:len(d.a)-1]
	return true
}

func maxSlidingWindow(nums []int, k int) []int {
	var Q ArrayDeque

	var push = func(val int) {
		// 入队的时候, last 方向入队， 但是入队的时候
		// 需要保证整个队列的数值是单调的
		// (在这个题里面我们需要是递减的)
		// 并且需要注意, 这里是 Q.getLast() < val
		// 如果写成 Q.getLast() <= val 就变成了严格单调递增
		for !Q.isEmpty() && Q.getLast() < val {
			Q.removeLast()
		}
		// 将元素入队
		Q.addLast(val)
	}

	var pop = func(val int) {
		// 出队的时候，要与头元素相等的时候，才会出队
		if !Q.isEmpty() && Q.getFirst() == val {
			Q.removeFirst()
		}
	}

	var ans []int
	for i := range nums {
		push(nums[i])
		// 如果添加队列的元素还少于 k 个
		// 那么这个时候，还不能去取最大值
		if i < k-1 {
			continue
		}
		ans = append(ans, Q.getFirst())

		pop(nums[i-k+1])
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
}
