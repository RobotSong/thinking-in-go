package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var Q = []*TreeNode{root}
	var ans [][]int
	for len(Q) > 0 {
		var qSize = len(Q)
		// 当前层的结果保存 放于 tmp 表中
		var tmp []int
		for i := 0; i < qSize; i++ {
			// 将当前层前面的结点先出队
			cur := Q[0]
			Q = Q[1:]
			tmp = append(tmp, cur.Val)

			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}

			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
		}

		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	var root = TreeNode{
		Val: 10,
	}

	var ans = levelOrder(&root)
	fmt.Printf("%v\n", ans)
}
