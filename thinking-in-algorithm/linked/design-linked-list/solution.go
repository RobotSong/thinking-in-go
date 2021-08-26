package main

import "fmt"

type ListNode struct {
	// val 用来存放链表中的数据
	val int
	// next 指向下一个节点
	next *ListNode
}

type MyLinkedList struct {
	// head 头节点
	head *ListNode
	// tail 尾节点
	tail *ListNode
	// length 链表的长度
	length int
}

// Constructor Initialize your data structure here.
func Constructor() MyLinkedList {
	h := &ListNode{}
	return MyLinkedList{
		head: h,
		tail: h,
	}
}

func (this *MyLinkedList) getPrevNode(index int) *ListNode {
	// 初始化
	front := this.head.next
	var back = this.head
	for i := 0; i < index && front != nil; i++ {
		back = front
		front = front.next
	}

	return back
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	// 因为 getPrevNode 总是返回有效的结点, 所以可以直接取值。
	return this.getPrevNode(index).next.val
}

// AddAtHead Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (this *MyLinkedList) AddAtHead(val int) {
	// 生成一个结点，存放的值 val
	p := &ListNode{val: val}
	// 将 p.next 指向第一个结点
	p.next = this.head.next
	// head.next 指向新结点, 使之变成第一个结点
	this.head.next = p
	if this.head == this.tail {
		this.tail = p
	}
	// 链表长度 + 1
	this.length++
}

// AddAtTail Append a node of value val to the last element of the linked list.
func (this *MyLinkedList) AddAtTail(val int) {
	// 尾部添加一个 新结点
	this.tail.next = &ListNode{val: val}
	// 移动 tail 指针
	this.tail = this.tail.next
	// 链表长度 + 1
	this.length++
}

// AddAtIndex Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index > this.length:
		// 如果 index 大于链表长度, 则不会插入结点
		return
	case index == this.length:
		// 如果 index 等于链表的长度, 则该结点附加到链表的末尾
		this.AddAtTail(val)
	case index <= 0:
		// 如果 index 小于等于 0, 则在头部插入结点
		this.AddAtHead(val)
	default:
		// 得到 index 之前的结点 pre
		pre := this.getPrevNode(index)
		// 在 pre 后面添加新结点
		p := &ListNode{val: val}
		p.next = pre.next
		pre.next = p
		// 添加链表长度
		this.length++
	}
}

// DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		// 如果 index 无效, 那么什么也不做。
		return
	}
	// 找到 index 前面的结点
	pre := this.getPrevNode(index)
	if this.tail == pre.next {
		// 如果删除的是最后一个结点, 那么需要更改 tail 指针
		this.tail = pre
	}
	// 进行删除操作, 并修改链表长度
	pre.next = pre.next.next
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	obj := Constructor()
	fmt.Println(obj.Get(0))
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1)) // 返回2
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1)) // 返回3
}
