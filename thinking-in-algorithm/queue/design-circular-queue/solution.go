package main

import "fmt"

type MyCircularQueue struct {
	used  int // 已经使用的元素个数
	front int // 第一个元素所在位置
	// rear 是 enQueue 可在存放的位置
	// 注意开闭原则
	// [front, rear)
	rear     int
	capacity int   // 循环队列最多可以存放的元素个数
	a        []int // 循环队列的存储空间
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k,
		a:        make([]int, k, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	// 如果已经满了， 则不加入队列
	if this.IsFull() {
		return false
	}
	// 入队
	this.a[this.rear] = value
	// 下一个元素入队 的位置
	this.rear = (this.rear + 1) % this.capacity
	// 已经使用的空间
	this.used++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	// 将第一个元素去除
	//var ret = this.a[this.front]
	this.front = (this.front + 1) % this.capacity
	// 使用元素减1
	this.used--
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
	tail := (this.rear - 1 + this.capacity) % this.capacity
	return this.a[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.used == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.used == this.capacity
}

func main() {
	// Your MyCircularQueue object will be instantiated and called as such:
	obj := Constructor(2)
	param1 := obj.EnQueue(12)
	fmt.Println(param1)
	param2 := obj.DeQueue()
	fmt.Println(param2)
	param3 := obj.Front()
	fmt.Println(param3)
	param4 := obj.Rear()
	fmt.Println(param4)
	param5 := obj.IsEmpty()
	fmt.Println(param5)
	param6 := obj.IsFull()
	fmt.Println(param6)

}
