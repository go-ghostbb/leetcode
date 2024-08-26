package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) Push(val *TreeNode) {
	*s = append(*s, val)
}

func (s *Stack) Pop() *TreeNode {
	l := len(*s) - 1
	res := (*s)[l]
	*s = (*s)[:l]
	return res
}

func (s *Stack) Len() int {
	return len(*s)
}

func inorderTraversal(root *TreeNode) []int {
	// init
	result := make([]int, 0)
	stack := make(Stack, 0)

	for root != nil || stack.Len() > 0 {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			// 如果左邊遍歷完了, 加進答案並換遍歷右邊
			node := stack.Pop()
			result = append(result, node.Val)
			root = node.Right
		}
	}

	return result
}
