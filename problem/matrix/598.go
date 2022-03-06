package matrix

import "math"

/*
给定一个初始元素全部为0，大小为 m*n 的矩阵M以及在M上的一系列更新操作。

操作用二维数组表示，其中的每个操作用一个含有两个正整数a 和 b 的数组表示，含义是将所有符合0 <= i < a 以及 0 <= j < b 的元素M[i][j]的值都增加 1。

在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。

 */

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m*n
	}
	x, y := math.MaxInt64, math.MaxInt64
	for _, col := range ops {
		x = min(x, col[0])
		y = min(y, col[1])
		//for _, val := range col {
		//	y = min(y, val)
		//}
	}
	return x*y
}
