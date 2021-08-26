package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 使用 假头 和 新链表
	dummy := &ListNode{}
	tail := dummy
	for p := head; p != nil; p = p.Next {
		if p.Val != val {
			tail.Next = p
			tail = p
		}
	}

	tail.Next = nil
	return dummy.Next
}

func main() {
	ans := removeElements(&ListNode{Val: 1}, 1)
	fmt.Println(ans)
}
