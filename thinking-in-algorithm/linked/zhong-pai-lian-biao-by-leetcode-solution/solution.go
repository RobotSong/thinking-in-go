package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func findMiddle(head *ListNode) *ListNode {
	s1 := head
	s2 := head

	for s2 != nil && s2.Next != nil {
		s1 = s1.Next
		s2 = s2.Next.Next
	}

	return s1
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTmp
	}

	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		}
		if l2 != nil {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		}
	}
	return dummy.Next
}

func reorderList(head *ListNode) {

	// 查找中间结点
	mid := findMiddle(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	// 反转 结点
	l2 = reverse(l2)
	// 合并结点
	head = merge(l1, l2)
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
	data := changeData([]int{1, 2, 3, 4})
	reorderList(data)
	printList(data)

}
