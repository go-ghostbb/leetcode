package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	newIntervals := make([][]int, 0)
	inserted := false

	for _, i := range intervals {
		if i[0] > newInterval[0] && !inserted {
			newIntervals = append(newIntervals, newInterval)
			inserted = true
		}
		newIntervals = append(newIntervals, i)
	}
	if inserted {
		return merge(newIntervals)
	}
	if !inserted && len(newInterval) != 0 {
		newIntervals = append(newIntervals, newInterval)
		return merge(newIntervals)
	}

	return newIntervals
}

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			// 代表需要合併
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}

	}
	return intervals
}
