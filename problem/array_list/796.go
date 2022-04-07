package array_list

/*
给定两个字符串, s和goal。如果在若干次旋转操作之后，s能变成goal，那么返回true。

s的 旋转操作 就是将s 最左边的字符移动到最右边。

例如, 若s = 'abcde'，在旋转一次之后结果就是'bcdea'。


示例 1:

输入: s = "abcde", goal = "cdeab"
输出: true
示例 2:

输入: s = "abcde", goal = "abced"
输出: false

*/

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	length := len(s)
	first := goal[0]
	find := make([]int, 0)
	for i := 0; i < length; i++ {
		if s[i] == first {
			find = append(find, i)
		}
	}
	ret := false
	if len(find) > 0 {
		for _, thisFind := range find {
			j := 0
			for ; j < length; j++ {
				if goal[j] != s[(j+thisFind)%length] {
					break
				}
			}
			if j == length {
				return true
			}
		}
	} else {
		return false
	}
	return ret
}
