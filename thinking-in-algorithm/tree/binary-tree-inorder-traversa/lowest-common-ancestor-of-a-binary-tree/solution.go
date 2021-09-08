package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans = nil
	postOrder(root, p, q)
	return ans
}

var ans *TreeNode

func postOrder(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	// 查看子节点的统计个数
	lcnt := postOrder(root.Left, p, q)
	rcnt := postOrder(root.Right, p, q)
	// 利用子结点的统计个数
	// 如果左边一个，右边有一个，那么当前的 root 就是最低公共祖先
	if lcnt == 1 && rcnt == 1 {
		ans = root
	} else if lcnt == 1 || rcnt == 1 {
		// 如果 左边找到了一个 或者右边找到一个
		// 并且 root 等于其中一个结点 p | q
		// 那么当前 root 就是最低公共祖先
		if root == p || root == q {
			ans = root
		}
	}
	// 返回值为以root为根的子树, 统计里面的p,q结点的个数。
	var tcnt int
	if root == p || root == q {
		tcnt = 1
	}
	return lcnt + rcnt + tcnt
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	isAns := lowestCommonAncestor(root, root.Left, root.Right)
	fmt.Println(isAns)
}
