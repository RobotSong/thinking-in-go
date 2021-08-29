package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 先查找是否能找到相交的 结点 , 就是是否是环形链表
	if head == nil || head.Next == nil {
		return nil
	}

	s1 := head
	s2 := head
	for s2 != nil && s2.Next != nil {
		s1 = s1.Next
		s2 = s2.Next.Next
		if s1 == s2 {
			break
		}
	}

	if s1 != s2 {
		return nil
	}
	// s1 重新从头开始移动, s2 则从相遇的位置开始移动
	s1 = head
	// 然后两个指针一起走
	for s1 != s2 && s2 != nil {
		s1 = s1.Next
		s2 = s2.Next
	}
	// 返回环形的入口结点
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

func main() {
	data := []int{3, 2, 0, -4}
	head := changeData(data)
	// 做一个环形链表
	c := head.Next
	// 最后一个结点, 则指向第二个结点
	head.Next.Next.Next.Next = c

	ans := detectCycle(head)
	fmt.Println(ans.Val)
}
