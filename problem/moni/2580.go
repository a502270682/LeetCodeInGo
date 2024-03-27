package moni

import (
	"sort"
)

/*
给你一个二维整数数组 ranges ，其中 ranges[i] = [starti, endi] 表示 starti 到 endi 之间（包括二者）的所有整数都包含在第 i 个区间中。

你需要将 ranges 分成 两个 组（可以为空），满足：

每个区间只属于一个组。
两个有 交集 的区间必须在 同一个 组内。
如果两个区间有至少 一个 公共整数，那么这两个区间是 有交集 的。

比方说，区间 [1, 3] 和 [2, 5] 有交集，因为 2 和 3 在两个区间中都被包含。
请你返回将 ranges 划分成两个组的 总方案数 。由于答案可能很大，将它对 (10^9 + 7) 取余 后返回
*/

func countWays(ranges [][]int) int {
	if len(ranges) <= 0 {
		return len(ranges) * 2
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	splitRange := make([][]int, 0)
	lastMerge := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if lastMerge[1] >= ranges[i][0] {
			lastMerge = []int{lastMerge[0], max(lastMerge[1], ranges[i][1])}
		} else {
			splitRange = append(splitRange, lastMerge)
			lastMerge = ranges[i]
		}
	}
	splitRange = append(splitRange, lastMerge)
	mod := 1000000000 + 7

	ret := 1
	for i := 1; i <= len(splitRange); i++ {
		ret = (ret * 2) % mod
	}
	return ret % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
