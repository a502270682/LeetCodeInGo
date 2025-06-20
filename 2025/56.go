package _025

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		length := len(ret)
		if intervals[i][0] > ret[length-1][1] {
			ret = append(ret, intervals[i])
		} else if ret[length-1][1] >= intervals[i][1] {
		} else {
			ret[length-1] = []int{ret[length-1][0], intervals[i][1]}
		}
	}
	return ret
}
