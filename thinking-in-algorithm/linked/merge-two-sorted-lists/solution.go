package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 首先生成空链表 假头
	dummy := &ListNode{}
	tail := dummy
	// 遍历两个 链表
	for l1 != nil || l2 != nil {
		// 如果 l2 链表为空, 或者 l1 链表的值更小
		if l2 == nil || l1 != nil && l1.Val < l2.Val {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		} else {
			// 其他情况， 则把 l2 结点添加到新链表的尾部
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		}
	}

	tail.Next = nil
	return dummy.Next
}

func main() {
	var changeData = func(data []int) *ListNode {
		dummy := &ListNode{}
		tail := dummy
		for _, val := range data {
			p := &ListNode{Val: val}
			tail.Next = p
			tail = p
		}
		return dummy.Next
	}

	l1 := changeData([]int{1, 2, 4})
	l2 := changeData([]int{1, 3, 4})

	ans := mergeTwoLists(l1, l2)

	fmt.Print("[")
	for p := ans; p != nil; p = p.Next {
		fmt.Printf("%d,", p.Val)
	}
	fmt.Print("]\n")

}
