package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swapVal(a, b *TreeNode) {
	t := a.Val
	a.Val = b.Val
	b.Val = t
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1.
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 如果已经没有了左右子树, 则直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil {
			// 还存在左子树
			// 那么需要从左子树中找较大值
			large := root.Left
			for large.Right != nil {
				large = large.Right
			}
			// 交换再删除
			swapVal(root, large)
			root.Left = deleteNode(root.Left, key)
		} else if root.Right != nil {
			// 还存在右子树
			// 那么还需要从右子树中找较小值
			small := root.Right
			for small.Left != nil {
				small = small.Left
			}

			// 交换后再删除
			swapVal(root, small)
			root.Right = deleteNode(root.Right, key)
		}
	}

	return root
}

func main() {
	root := &TreeNode{Val: 1}

	ans := deleteNode(root, 1)
	fmt.Println(ans)
}
