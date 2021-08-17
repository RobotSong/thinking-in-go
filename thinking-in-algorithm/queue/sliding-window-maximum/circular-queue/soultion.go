package main

import "fmt"

type CircularQueue struct {
	front    int
	rear     int
	capacity int
	a        []int
}

func Construct(k int) CircularQueue {
	return CircularQueue{
		capacity: k + 1,
		a:        make([]int, k+1, k+1),
	}
}

func (q *CircularQueue) RearEnQueue(val int) bool {
	if q.IsFull() {
		return false
	}

	q.a[q.rear] = val
	q.rear = (q.rear + 1) % q.capacity
	return true
}

func (q *CircularQueue) FirstEnQueue(val int) bool {
	if q.IsFull() {
		return false
	}

	q.front = (q.front - 1 + q.capacity) % q.capacity
	q.a[q.front] = val
	return true
}

func (q *CircularQueue) RearDeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	// 从尾部出队
	q.rear = (q.rear - 1 + q.capacity) % q.capacity
	return true
}

func (q *CircularQueue) FirstDeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	// 从头部出队
	q.front = (q.front + 1) % q.capacity
	return true
}

func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.a[q.front]
}

func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	tail := (q.rear - 1 + q.capacity) % q.capacity
	return q.a[tail]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircularQueue) IsFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// 根据值是否相等，从头部出队
func (q *CircularQueue) popFront(val int) bool {
	if q.IsEmpty() {
		return false
	}

	if q.Front() != val {
		return false
	}
	// 循环队列的从头部出队
	q.FirstDeQueue()
	return true
}

// 根据入队的值，是否小于尾部的值，从尾部入队
func (q *CircularQueue) pushRear(val int) bool {

	for !q.IsEmpty() && q.Rear() < val {
		// 如果尾部的值小于入队的值，将尾部的值出队
		q.RearDeQueue()
	}

	q.RearEnQueue(val)
	return true
}

func maxSlidingWindow(nums []int, k int) []int {
	q := Construct(k)
	var ans []int
	for i := range nums {
		q.pushRear(nums[i])

		if i < k-1 {
			continue
		}

		ans = append(ans, q.Front())
		q.popFront(nums[i-k+1])
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
	nums = []int{1, 3, 1, 2, 0, 5}
	k = 3
	ans = maxSlidingWindow(nums, 3)
	// 预期结果 [3,3,2,5]
	fmt.Println(ans)
}
