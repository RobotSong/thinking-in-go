package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 先遍历 root 的值
		for root != nil {
			stack = append(stack, root)
			ans = append(ans, root.Val)
			root = root.Left
		}

		// 当无法压栈的时候， 将 root.Right 进行出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right

	}

	return ans
}

func main() {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right
	ans := preorderTraversal(root)
	fmt.Println(ans)

}
