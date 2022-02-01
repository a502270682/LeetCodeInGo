package move_window

import (
	"math/bits"
	"unicode"
)

/*
当一个字符串 包含的每一种字母的大写和小写形式 同时出现在中，就称这个字符串是 美好 字符串。比方说，"abABB"是美好字符串，因为'A' 和'a'同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。

给你一个字符串，请你返回最长的美好子字符串。如果有多个答案，请你返回最早出现的一个。如果不存在美好子字符串，请你返回一个空字符串。



示例 1：

输入：s = "YazaAay"
输出："aAa"
解释："aAa" 是一个美好字符串，因为这个子串中仅含一种字母，其小写形式 'a' 和大写形式 'A' 也同时出现了。
"aAa" 是最长的美好子字符串。
*/

func longestNiceSubstring(s string) (ans string) {
	mask := uint(0)
	for _, ch := range s {
		mask |= 1 << (unicode.ToLower(ch) - 'a')
	}
	maxTypeNum := bits.OnesCount(mask)

	for typeNum := 1; typeNum <= maxTypeNum; typeNum++ {
		var lowerCnt, upperCnt [26]int
		var total, cnt, l int
		for r, ch := range s {
			idx := unicode.ToLower(ch) - 'a'
			if unicode.IsLower(ch) {
				lowerCnt[idx]++
				if lowerCnt[idx] == 1 && upperCnt[idx] > 0 {
					cnt++
				}
			} else {
				upperCnt[idx]++
				if upperCnt[idx] == 1 && lowerCnt[idx] > 0 {
					cnt++
				}
			}
			if lowerCnt[idx]+upperCnt[idx] == 1 {
				total++
			}

			for total > typeNum {
				idx := unicode.ToLower(rune(s[l])) - 'a'
				if lowerCnt[idx]+upperCnt[idx] == 1 {
					total--
				}
				if unicode.IsLower(rune(s[l])) {
					lowerCnt[idx]--
					if lowerCnt[idx] == 0 && upperCnt[idx] > 0 {
						cnt--
					}
				} else {
					upperCnt[idx]--
					if upperCnt[idx] == 0 && lowerCnt[idx] > 0 {
						cnt--
					}
				}
				l++
			}

			if cnt == typeNum && r-l+1 > len(ans) {
				ans = s[l : r+1]
			}
		}
	}
	return
}
