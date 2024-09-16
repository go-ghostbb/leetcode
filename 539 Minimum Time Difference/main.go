package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}

func findMinDifference(timePoints []string) int {
	minuteList := make([]int, len(timePoints))

	for i, t := range timePoints {
		tmp, _ := time.Parse("15:04", t)
		minuteList[i] = tmp.Hour()*60 + tmp.Minute()
	}

	sort.Ints(minuteList)

	ans := minuteList[len(minuteList)-1] - minuteList[0]
	ans = min(ans, minuteList[0]-minuteList[len(minuteList)-1]+1440)

	for i := 0; i < len(minuteList)-1; i++ {
		ans = min(ans, minuteList[i+1]-minuteList[i], minuteList[i]-minuteList[i+1]+1440)
	}

	return ans
}
