package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	// 存放遍历的结果
	var ans []int

	var stack []*TreeNode
	// pre 表示遍历时，前面一个已经遍历过的结点
	var pre *TreeNode
	// 如果栈中还有元素，或者 当前结点 t 非空
	for root != nil || len(stack) > 0 {
		// 顺着 左子树走， 并且将所有的元素 压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 当没有任何元素可以压栈的时候
		// 拿栈顶元素，注意这里并不将栈顶元素弹出
		// 因为在迭代时，根节点需要遍历两次，这里需要判断一下
		// 右子树是否遍历完毕
		root = stack[len(stack)-1]
		// 如果遍历当前结点，需要确保右子树已经遍历完毕
		// 1. 如果当前结点右子树为空，那么右子树没有遍历的必要
		// 需要将当前结点放到 ans 中
		// 2. 当 root.Right == pre 说明右子树 已经被打印过了
		// 那么此时需要将当前结点放到 nas 中
		if root.Right == nil || root.Right == pre {
			// 右子树已经遍历完毕，放到 ans 中
			ans = append(ans, root.Val)
			// 出栈
			stack = stack[:len(stack)-1]
			// y因为已经遍历了当前结点，所以需要更新 pre 结点
			pre = root
			// 已经打印完毕，需要设置为空，否则下一轮循环
			// 还会遍历 root 的 左子树。
			root = nil
		} else {
			// 第一次走到 root 结点， 不能放到 ans 中， 因为 root 的 右子树还没有遍历
			// 需要将 root 结点的右子树遍历
			root = root.Right
		}

	}

	return ans
}

func main() {
	root := &TreeNode{Val: 1}
	right := &TreeNode{Val: 2}
	root.Right = right
	rightLeft := &TreeNode{Val: 3}
	right.Left = rightLeft

	ans := postorderTraversal(root)
	fmt.Println(ans)
}
