package main

import "strconv"

/*
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 10^8]
*/

func maximumSwap(num int) int {
	numS := strconv.Itoa(num)
	left, right := -1, -1
	maxN := len(numS) - 1
	for i := len(numS) - 2; i >= 0; i-- {
		if numS[i] > numS[maxN] {
			maxN = i
		} else if numS[i] < numS[maxN] {
			left, right = i, maxN
		}
	}
	if left == -1 {
		return num
	}
	bNumS := []byte(numS)

	tmp := bNumS[left]
	bNumS[left] = bNumS[right]
	bNumS[right] = tmp
	ret, _ := strconv.Atoi(string(bNumS))
	return ret
}
