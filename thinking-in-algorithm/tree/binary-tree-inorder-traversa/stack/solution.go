package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 遍历到最后那个子节点 取值放到 结果中
		root = stack[len(stack)-1]
		ans = append(ans, root.Val)
		// 出栈
		stack = stack[:len(stack)-1]

		root = root.Right
	}

	return ans
}

func main() {
	root := &TreeNode{Val: 1}

	ans := inorderTraversal(root)
	fmt.Println(ans)
}
