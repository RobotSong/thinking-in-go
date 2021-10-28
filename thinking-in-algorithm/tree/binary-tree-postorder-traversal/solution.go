package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	preOrder(root.Left, ans)
	preOrder(root.Right, ans)
	*ans = append(*ans, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	preOrder(root, &ans)
	return ans
}

func main() {
	root := &TreeNode{Val: 1}

	ans := postorderTraversal(root)
	fmt.Println(ans)
}
