package main

type MyCircularQueue struct {
	// 队列的头部元素所在位置
	front int
	// 队列的尾巴
	// [MyCircularQueue.front, rear)
	rear int
	//
	a        []int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k + 1,
		a:        make([]int, k+1, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.a[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	// 出队之后， front 要向前移
	this.front = (this.front + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.a[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	// 由于我们使用的是前开后闭原则
	// [front, rear)
	// 所以在取最后一个元素时，应该是
	// (rear - 1 + capacity) % capacity;
	tail := (this.rear - 1 + this.capacity) % this.capacity
	return this.a[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	// 队列是否为空
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	// rear与front之间至少有一个空格
	// 当rear指向这个最后的一个空格时，
	// 队列就已经放满了!
	return (this.rear+1)%this.capacity == this.front
}
