package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	// 注意假头并不算是链表的一部分，所以这里是从 head 开始走
	s1 := head
	s2 := head

	// 两个指针同时走
	// 因为 s2 指针每次都要走两步， 所以需要对 s2.Next 这样判断
	for s2 != nil && s2.Next != nil {

		s1 = s1.Next
		s2 = s2.Next.Next
	}

	return s1
}

func changeData(data []int) *ListNode {
	dummy := &ListNode{}

	tail := dummy
	for _, val := range data {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	fmt.Print("[")
	for p := head; p != nil; p = p.Next {
		if p.Next == nil {
			fmt.Print(p.Val)
		} else {
			fmt.Printf("%d,", p.Val)
		}
	}
	fmt.Print("]")
}

func main() {
	data := changeData([]int{1, 2, 3, 4, 5})
	ans := middleNode(data)
	printList(ans)

}
