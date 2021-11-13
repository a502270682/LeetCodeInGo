package string

/*
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。



示例 1：

输入：word = "USA"
输出：true
示例 2：

输入：word = "FlaG"
输出：false

*/

func detectCapitalUse(word string) bool {
	needAllBig := false
	needAllSmall := false
	for idx, w := range word {
		if w <= 90 {
			if needAllSmall {
				return false
			}
			if idx == 0 {
				continue
			} else {
				needAllBig = true
				if needAllBig {
					continue
				} else {
					return false
				}
			}
		} else {
			needAllSmall = true
			if needAllBig {
				return false
			}
		}
	}
	return true
}
