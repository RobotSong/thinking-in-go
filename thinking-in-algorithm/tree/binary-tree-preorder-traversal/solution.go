package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = preOrder(root.Left, ans)
	ans = preOrder(root.Right, ans)
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int

	ans = preOrder(root, ans)
	return ans
}

func main() {
	root := &TreeNode{Val: 1}

	ans := preorderTraversal(root)
	fmt.Println(ans)
}
