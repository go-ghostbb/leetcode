package main

import "fmt"

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
			},
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 6,
					},
					{
						Val: 7,
						Children: []*Node{
							{
								Val: 11,
								Children: []*Node{
									{
										Val: 14,
									},
								},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{
							{
								Val: 12,
							},
						},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
						Children: []*Node{
							{
								Val: 13,
							},
						},
					},
					{
						Val: 10,
					},
				},
			},
		},
	}

	fmt.Println(postorder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

type Stack []*Node

func (s *Stack) Push(val *Node) {
	*s = append(*s, val)
}

func (s *Stack) Pop() *Node {
	l := len(*s) - 1
	res := (*s)[l]
	*s = (*s)[:l]
	return res
}

func (s *Stack) Len() int {
	return len(*s)
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := make(Stack, 0)

	stack.Push(root)

	for stack.Len() > 0 {
		node := stack.Pop()

		result = append([]int{node.Val}, result...)

		for _, child := range node.Children {
			stack.Push(child)
		}
	}

	return result
}
