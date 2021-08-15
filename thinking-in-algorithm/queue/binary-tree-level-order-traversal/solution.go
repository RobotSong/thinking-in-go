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

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var cur []*TreeNode
	cur = append(cur, root)
	for len(cur) > 0 {
		var next []*TreeNode
		var curResult []int
		for _, node := range cur {
			curResult = append(curResult, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		// 注意这里的迭代滚动前进
		cur = next

		ans = append(ans, curResult)
	}

	return ans
}

func main() {
	var root = TreeNode{
		Val: 10,
	}

	var ans = levelOrder(&root)
	fmt.Printf("%v\n", ans)
	ans = levelOrder1(&root)
	fmt.Printf("%v\n", ans)
}
