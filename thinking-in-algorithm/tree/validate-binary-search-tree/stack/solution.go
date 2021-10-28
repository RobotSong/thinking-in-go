package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	// 堆栈
	var stack []*TreeNode
	// 中序遍历中最小的值
	inorder := math.MinInt64

	for root != nil || len(stack) > 0 {
		// 先遍历所有左结点到堆栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 取出最后放到堆栈中的左结点的值
		root = stack[len(stack)-1]
		// 去掉最后一个
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			// 如果结点的值, 不符合升序的规则, 则不是二叉搜索树
			return false
		}
		inorder = root.Val
		root = root.Right
	}

	return true
}

func main() {
	root := &TreeNode{Val: 0}
	res := isValidBST(root)
	fmt.Println(res)
}
