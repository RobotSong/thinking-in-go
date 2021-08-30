package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// backtrace 遍历 树的结点
func backtrace(root *TreeNode, path []int, sum, target int, ans *[][]int) []int {
	if root == nil {
		return path
	}
	// 前序遍历, 加上累积的和
	sum += root.Val
	// 将当前结点添加到路径中，相当于 压栈
	path = append(path, root.Val)
	// 从根节点到叶子节点
	if root.Left == nil && root.Right == nil {
		if sum == target {
			// 当前路径符合要求
			ele := make([]int, len(path))
			copy(ele, path)
			*ans = append(*ans, ele)
		}
	} else {
		path = backtrace(root.Left, path, sum, target, ans)
		path = backtrace(root.Right, path, sum, target, ans)
	}

	// 遍历结束需要把当前值给出栈
	path = path[:len(path)-1]
	return path
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	// 遍历 二叉树 把每个结点的值 入栈， 遍历完该 结点 出栈
	var path []int
	var ans [][]int
	backtrace(root, path, 0, targetSum, &ans)

	return ans
}

func main() {
	left := &TreeNode{Val: 5}
	right := &TreeNode{Val: 5}
	root := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right
	ans := pathSum(root, 8)
	fmt.Println(ans)
}
