package main

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// findMiddleNode 查找中间结点
func findMiddleNode(head *ListNode) *ListNode {

	s1 := head
	s2 := head
	pre := s1
	for s2 != nil && s2.Next != nil {
		pre = s1
		s1 = s1.Next
		s2 = s2.Next.Next
	}

	if s2 != nil {
		return s1
	} else {
		return pre
	}
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddleNode(head)
	back := mid.Next
	mid.Next = nil
	// 后续遍历左右两个链表
	i := mergeSort(head)
	j := mergeSort(back)

	dummy := &ListNode{}

	tail := dummy

	for i != nil || j != nil {

		if j == nil || i != nil && i.Val <= j.Val {
			tail.Next = i
			tail = i
			i = i.Next
		} else {
			tail.Next = j
			tail = j
			j = j.Next
		}
	}
	// 习惯设置为 nil
	tail.Next = nil

	return dummy.Next
}

// sortList 使用合并排序
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
