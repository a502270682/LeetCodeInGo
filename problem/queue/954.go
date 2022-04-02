package queue

import (
	"math"
	"sort"
)

/*
给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <=i < len(arr) / 2，
都有 arr[2 * i + 1] = 2 * arr[2 * i]”时，返回 true；否则，返回 false。



示例 1：

输入：arr = [3,1,3,6]
输出：false
示例 2：

输入：arr = [2,1,2,6]
输出：false
示例 3：

输入：arr = [4,-2,2,-4]
输出：true
解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]

*/

func canReorderDoubled(arr []int) bool {
	k := make(map[int]int)
	for _, a := range arr {
		if count, ok := k[a]; ok {
			k[a] = count + 1
		} else {
			k[a] = 1
		}
	}
	targetNum := 0
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	for _, a := range arr {
		if count1, ok := k[a]; ok && count1 > 0 {
			if a == 0 && count1 <= 1 {
				continue
			}
			if count, ok := k[a*2]; ok && count > 0 {
				k[a*2] = k[a*2] - 1
				k[a] = k[a] - 1
				targetNum++
			}
		}
	}
	return targetNum >= len(arr)/2
}
