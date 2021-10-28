package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(a, b *ListNode) *ListNode {
	ans := &ListNode{}
	tail := ans
	for a != nil || b != nil {
		if a != nil {
			tail.Next = a
			tail = a
			a = a.Next
		}
		if b != nil {
			tail.Next = b
			tail = b
			b = b.Next
		}
	}
	return ans.Next
}

func swapPairs(head *ListNode) *ListNode {
	// 分成 奇数、 偶数链表
	odd := &ListNode{}
	oddTail := odd
	even := &ListNode{}
	evenTail := even
	var index int
	for p := head; p != nil; p = p.Next {
		if index%2 == 0 {
			oddTail.Next = p
			oddTail = p
		} else {
			evenTail.Next = p
			evenTail = p
		}
		index++
	}
	// 注意两个链表的结尾 要设置为 nil
	oddTail.Next = nil
	evenTail.Next = nil
	// 合并两个 链表
	return mergeList(even.Next, odd.Next)
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

	data := []int{1, 2, 3, 4, 5, 7}
	head := changeData(data)
	ans := swapPairs(head)
	fmt.Print("[")
	for p := ans; p != nil; p = p.Next {
		fmt.Printf("%d,", p.Val)
	}
	fmt.Print("]\n")
}
