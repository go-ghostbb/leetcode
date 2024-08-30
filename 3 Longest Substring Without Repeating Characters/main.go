package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	// Sliding Window
	var (
		left    int
		ans     int
		charMap = make(map[uint8]int) // map[char]index
	)

	for right := 0; right < len(s); right++ {
		char := s[right]
		if _, ok := charMap[char]; ok {
			// 如果在map裡面, 代表有出現過
			// 更新左指標
			left = max(left, charMap[char]+1)
		}

		charMap[char] = right

		ans = max(ans, right-left+1)
	}

	return ans
}
