package main

import "fmt"

func main() {
	// 3,0,2,6,8,1,7,9,4,2,5,5,0
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 2,
											Next: &ListNode{
												Val: 5,
												Next: &ListNode{
													Val: 5,
													Next: &ListNode{
														Val: 0,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println(spiralMatrix(3, 5, head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	// init
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}

	var (
		i     int // 0
		j     int // 0
		count int
	)

	for head != nil {
		ans[i][j] = head.Val
		head = head.Next

		switch count {
		case 0:
			// right
			// 如果到底了或下一個不等於-1, 往下走
			if j == n-1 || ans[i][j+1] != -1 {
				i++
				count = 1
				continue
			}
			j++
		case 1:
			if i == m-1 || ans[i+1][j] != -1 {
				j--
				count = 2
				continue
			}
			i++ // down
		case 2:
			if j == 0 || ans[i][j-1] != -1 {
				i--
				count = 3
				continue
			}
			j-- // left
		case 3:
			if j == n-1 || ans[i-1][j] != -1 {
				j++
				count = 0
				continue
			}
			i-- // up
		}
	}

	return ans
}
