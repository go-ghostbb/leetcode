package main

import "fmt"

func main() {
	fmt.Println(minSteps(18))
}

func minSteps(n int) int {
	var ans int

	// 只要是因子n，動的步數都為n

	// 解題：做質因數分解，之後把們全部加起來就是答案

	// ex. n = 18
	// n = 2, ans = 2
	// n = 3, ans = 3
	// n = 4, ans = 2 + 2
	// n = 5, ans = 5
	// n = 6, ans = 2 + 3
	// n = 7, ans = 7
	// n = 8, ans = 2 + 2 + 2
	// ...
	// n = 18, ans = 2 + 3 + 3

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			ans += i
			n /= i
		}
	}

	return ans
}
