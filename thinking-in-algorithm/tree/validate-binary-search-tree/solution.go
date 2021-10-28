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

func preOrder(root *TreeNode, l int64, r int64, ans bool) bool {
	// 如果 root 为 nil
	// 已经验证为非二叉搜索树
	if root == nil || !ans {
		return ans
	}
	// 如果结点的值不在范围内
	if !(int64(root.Val) > l && int64(root.Val) < r) {
		return false
	}
	// 遍历结点的 左子树, 它要小于父结点的值 但是要大于 传递下来的 左边的值
	// (l, root.Val)
	ans = preOrder(root.Left, l, int64(root.Val), ans)
	// 遍历结点的 右子树, 它要大于父结点的值, 但是要小于 传递下来的 右边的值
	ans = preOrder(root.Right, int64(root.Val), r, ans)
	return ans
}

func isValidBST(root *TreeNode) bool {
	// 递归验证
	ans := preOrder(root, math.MinInt64, math.MaxInt64, true)
	return ans
}

func main() {
	root := &TreeNode{Val: 0}
	res := isValidBST(root)
	fmt.Println(res)
}
