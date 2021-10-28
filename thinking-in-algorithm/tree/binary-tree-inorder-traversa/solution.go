package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	ans = preOrder(root, ans)
	return ans
}

// 递归完成中序遍历
func preOrder(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}

	// 先遍历左结点的值
	ans = preOrder(root.Left, ans)
	ans = append(ans, root.Val)
	ans = preOrder(root.Right, ans)
	return ans
}

func main() {
	root := &TreeNode{Val: 1}

	ans := inorderTraversal(root)
	fmt.Println(ans)
}
