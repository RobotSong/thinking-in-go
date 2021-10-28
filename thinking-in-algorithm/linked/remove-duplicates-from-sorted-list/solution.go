package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 新链表 假头
	dummy := &ListNode{}
	// 末尾 节点
	tail := dummy
	for p := head; p != nil; p = p.Next {
		// 如果不是头节点, 并且 p 的 Val 等于 tail 的 Val,  则不入队
		if tail != dummy && p.Val == tail.Val {
			continue
		}
		tail.Next = p
		tail = p
	}
	// 使用新链表时, 末尾的 Next 必须设置为空
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

	data := []int{1, 1, 2, 3, 3}
	head := changeData(data)
	ans := deleteDuplicates(head)
	fmt.Print("[")
	for p := ans; p != nil; p = p.Next {
		fmt.Printf("%d,", p.Val)
	}
	fmt.Print("]\n")
}
