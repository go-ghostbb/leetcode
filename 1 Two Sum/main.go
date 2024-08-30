package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	temp := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		// 解題思路:找到target-nums[i]的位置即可
		// 利用map特性解題，結構為 map[int] = []int{num[i], i}
		if _, ok := temp[nums[i]]; ok {
			return []int{temp[nums[i]][1], i}
		} else {
			temp[target-nums[i]] = []int{nums[i], i}
		}
	}
	return nil
}
