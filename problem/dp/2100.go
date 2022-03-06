package dp

/*
你和一群强盗准备打劫银行。给你一个下标从 0开始的整数数组security，其中security[i]是第 i天执勤警卫的数量。日子从 0开始编号。同时给你一个整数time。

如果第 i天满足以下所有条件，我们称它为一个适合打劫银行的日子：

第 i天前和后都分别至少有 time天。
第 i天前连续 time天警卫数目都是非递增的。
第 i天后连续 time天警卫数目都是非递减的。
更正式的，第 i 天是一个合适打劫银行的日子当且仅当：security[i - time] >= security[i - time + 1] >= ... >= security[i] <= ... <= security[i + time - 1] <= security[i + time].

请你返回一个数组，包含 所有 适合打劫银行的日子（下标从 0开始）。返回的日子可以 任意顺序排列。



示例 1：

输入：security = [5,3,3,3,5,6,2], time = 2
输出：[2,3]
解释：
第 2 天，我们有 security[0] >= security[1] >= security[2] <= security[3] <= security[4] 。
第 3 天，我们有 security[1] >= security[2] >= security[3] <= security[4] <= security[5] 。
没有其他日子符合这个条件，所以日子 2 和 3 是适合打劫银行的日子。

 */

// dp
// 取left和right两个数组，left记录连续比当前i位置大或等于的前置位个数； right记录连续比当前i小或等于的后置位个数
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	if n <= 1 {
		if n > time {
			return []int{0}
		}
		return []int{}
	}
	left, right := make([]int, n), make([]int, n)
	left[0] = 0
	right[n-1] = 0
	for i := 1 ;i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] +1
		}
	}
	for i := n-2 ;i>0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] +1
		}
	}
	res := make([]int, 0)
	for i:=time ; i<n-time; i++ {
		if left[i] >= time &&right[i]>= time {
			res = append(res, i)
		}
	}
	return res
}
