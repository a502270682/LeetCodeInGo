package main

/*
自除数是指可以被它包含的每一位数整除的数。

例如，128 是一个 自除数 ，因为128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
自除数 不允许包含 0 。

给定两个整数left和right ，返回一个列表，列表的元素是范围[left, right]内所有的 自除数 。


示例 1：

输入：left = 1, right = 22
输出：[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
示例 2:

输入：left = 47, right = 85
输出：[48,55,66,77]

*/

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	needAdd := true
	for i := left; i <= right; i++ {
		tmp := i
		needAdd = true
		for tmp > 0 {
			if tmp%10 == 0 || i%(tmp%10) != 0 {
				needAdd = false
				break
			}
			tmp = tmp / 10
		}
		if needAdd {
			ret = append(ret, i)
		}
	}
	return ret
}
