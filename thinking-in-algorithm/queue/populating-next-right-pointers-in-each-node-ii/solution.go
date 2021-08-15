package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 使用一个队列解决问题
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 当前层 队列
	var cur []*Node
	cur = append(cur, root)
	// 4 , 5 , 7
	for len(cur) > 0 {

		qSize := len(cur)
		for i := 0; i < qSize; i++ {
			// 出队
			node := cur[0]
			cur = cur[1:]
			// 如果当前队列中不是最后一位， 则它指向右边
			if i < qSize-1 {
				node.Next = cur[0]
			}
			// 当前层级入队
			if node.Left != nil {
				cur = append(cur, node.Left)
			}

			if node.Right != nil {
				cur = append(cur, node.Right)
			}
		}
	}

	return root
}

// 使用两个数组解决 , LeetCode 上执行代码 OOM
func connectTwoList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var curList []*Node
	curList = append(curList, root)

	for len(curList) > 0 {
		var nextList []*Node
		for i, node := range curList {
			if i < len(curList)-1 {
				node.Next = curList[i]
			}

			// 当前层级入队
			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}

			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}
		curList = nextList
	}
	return root
}

func main() {
	root := Node{
		Val: 10,
	}

	ans := connect(&root)
	fmt.Println(ans)
}
