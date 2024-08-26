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

	fmt.Println(postorderTraversal(root))
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

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// init
	result := make([]int, 0)
	stack := make(Stack, 0)

	stack.Push(root)

	for stack.Len() > 0 {
		node := stack.Pop()

		result = append([]int{node.Val}, result...)

		if node.Left != nil {
			stack.Push(node.Left)
		}

		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	return result
}
