package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 将链表改造成带假头的链表
	dummy := &ListNode{Next: head}
	// 已经跑过的链表长度
	preWalledSteps := 0
	// front 指针从 dummy 开始先走 k 步
	front := dummy
	// 注意 front 不能为空, 需要指向链表的最后一个结点 && front.Next
	for preWalledSteps < n && front != nil && front.Next != nil {
		front = front.Next
		preWalledSteps++
	}

	// back 指针一开始指向 dummy, 然后与 front 指针一起移动
	back := dummy
	for front != nil && front.Next != nil {
		back = back.Next
		front = front.Next
	}
	// 如果 preWalledSteps == k
	// 说明存在可以删除的结点
	if preWalledSteps == n {
		back.Next = back.Next.Next
	}

	// 返回新的链表
	return dummy.Next
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

	head := changeData([]int{1, 2, 3})
	ans := removeNthFromEnd(head, 4)
	printList(ans)

}
