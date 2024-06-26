package bit_cal

/*
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x



示例 1：

输入：n = 16
输出：true

4的幂次方偶数位置为1
所以构造二进制位奇数位为1，与原n按位取与，结果为0且满足 n&(n-1) 则符合条件
mask=(10101010101010101010101010101010)
​


*/

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
