package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	s1 := head
	s2 := head

	for s2 != nil && s2.Next != nil {
		s1 = s1.Next
		s2 = s2.Next.Next
		// 如果存在环形链表, 则 s1 s2 不同的速度向前, s1 s2 最后会相等
		if s1 == s2 {
			return true
		}
	}

	return false
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	ans := hasCycle(head)
	fmt.Println(ans)
}
