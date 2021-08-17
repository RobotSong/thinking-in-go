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

func maxResult(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k == 0 {
		return 0
	}
	// 单调递增循环队列
	q := Construct(k + 1)
	// 每个位置可以收集到的金币数目
	get := make([]int, len(nums))

	for i := range nums {
		//在取最大值之前，需要保证单调队列中都是有效值。
		//也就是都在区间里面的值
		//当要求get[i]的时候，
		//单调队列中应该是只能保存[i-k,i-1]这个范围
		if i-k > 0 && q.Front() == get[i-k-1] {
			q.FirstDeQueue()
		}
		//从单调队列中取得较大值
		var old int
		if q.IsEmpty() {
			old = 0
		} else {
			old = q.Front()
		}

		get[i] = old + nums[i]
		q.pushRear(get[i])
	}

	return get[len(nums)-1]
}

func main() {
	var nums []int
	var k int
	var ans int
	//nums = []int{1,-1,-2,4,-7,3}
	//k = 2
	//ans = maxResult(nums, k)
	//fmt.Println(ans)
	//
	nums = []int{0, -1, -2, -3, 1}
	k = 2
	// 预期结果 -1
	ans = maxResult(nums, k)
	fmt.Println(ans)
}
