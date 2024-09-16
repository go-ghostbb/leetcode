package main

import "fmt"

func main() {
	fmt.Println(xorQueries([]int{16}, [][]int{{0, 0}, {0, 0}, {0, 0}}))
}

func xorQueries(arr []int, queries [][]int) []int {
	l := len(arr)
	ans := make([]int, len(queries))
	t := make([]int, l+1)

	for i := 1; i <= l; i++ {
		t[i] = t[i-1] ^ arr[i-1]
	}

	for i, q := range queries {
		ans[i] = t[q[1]+1] ^ t[q[0]]
	}

	return ans
}
