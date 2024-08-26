package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{2, 6, 4, 1, 3, 1, 5}))
}

func findDuplicate(nums []int) int {
	/*
	 * 假設len(nums)為n, 數字範圍為[1, n-1]
	 *
	 * 轉換為linked list, 造成迴圈的點就是出現重複的數字
	 * 使用兩個變數, 使用迴圈, 一個每次動一步, 另一個每次動兩步
	 * 跑到後面必定重疊
	 * 接著將其中一個變數改為linked list的起點, 使用迴圈, 兩個變數每次動一步
	 * 重疊點必為linked list造成迴圈的點
	 *
	 * a -> 每次動一步
	 * b -> 每次動兩步
	 */
	a, b := nums[0], nums[0]
	a = nums[a]       // 動一步
	b = nums[nums[b]] // 動兩步

	for a != b {
		a = nums[a]       // 動一步
		b = nums[nums[b]] // 動兩步
	}

	a = nums[0]

	for a != b {
		a = nums[a]
		b = nums[b]
	}

	return a
}
