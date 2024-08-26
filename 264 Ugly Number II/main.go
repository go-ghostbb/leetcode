package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(1))
}

func nthUglyNumber(n int) int {
	var (
		dp    = make([]int, n)
		two   int
		three int
		five  int
	)

	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = min(dp[two]*2, dp[three]*3, dp[five]*5)

		if dp[i] == dp[two]*2 {
			two++
		}

		if dp[i] == dp[three]*3 {
			three++
		}

		if dp[i] == dp[five]*5 {
			five++
		}
	}

	return dp[n-1]
}
