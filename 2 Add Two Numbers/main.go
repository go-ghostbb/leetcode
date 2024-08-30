package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	count := 0
	for {

		l1.Val += l2.Val + count // 相加

		// 進位
		count = l1.Val / 10
		l1.Val = l1.Val % 10

		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		// step
		l1 = l1.Next
		l2 = l2.Next
	}

	// 如果count!=0代表需要進位
	for count != 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1.Next.Val += count
		count = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10
		l1 = l1.Next
	}

	return head
}
