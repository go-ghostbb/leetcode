package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	fmt.Println(isSubPath(head, root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return checker(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func checker(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil || head.Val != root.Val {
		return false
	}

	return checker(head.Next, root.Left) || checker(head.Next, root.Right)
}
