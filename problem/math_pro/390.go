package math_pro

/*
列表 arr 由在范围 [1, n] 中的所有整数组成，并按严格递增排序。请你对 arr 应用下述算法：

从左到右，删除第一个数字，然后每隔一个数字删除一个，直到到达列表末尾。
重复上面的步骤，但这次是从右到左。也就是，删除最右侧的数字，然后剩下的数字每隔一个删除一个。
不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
给你整数 n ，返回 arr 最后剩下的数字。

*/

func lastRemaining(n int) int {
	// a1 等差数列的首， an 等差数列的尾
	a1, an := 1, n
	// step 为公差， 下一轮的等差数列的公差 = 上一轮的等差数列的公差 * 2
	// cnt 为剩余数字的总数量， 下一轮的cnt = 上一轮的cnt/2 (一定是向下取整)
	left, cnt, step := 0, n, 1
	for cnt > 1 {
		if left%2 == 0 { // 正向
			a1 += step
			if cnt%2 == 1 {
				an -= step
			}
		} else { // 反向
			if cnt%2 == 1 {
				a1 += step
			}
			an -= step
		}
		left++
		cnt >>= 1
		step <<= 1
	}
	return a1
}
