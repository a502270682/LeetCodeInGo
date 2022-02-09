package tanxin

import "sort"

/*
如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。

给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：

s 是一个尽可能长的快乐字符串。
s 中 最多 有a 个字母 'a'、b个字母 'b'、c 个字母 'c' 。
s 中只含有 'a'、'b' 、'c' 三种字母。
如果不存在这样的字符串 s ，请返回一个空字符串 ""。



示例 1：

输入：a = 1, b = 1, c = 7
输出："ccaccbcc"
解释："ccbccacc" 也是一种正确答案。
示例 2：

输入：a = 2, b = 2, c = 1
输出："aabbc"
示例 3：

输入：a = 7, b = 1, c = 0
输出："aabaa"
解释：这是该测试用例的唯一正确答案。


提示：

0 <= a, b, c <= 100
a + b + c > 0

*/

func longestDiverseString(a int, b int, c int) string {
	type cnt struct {
		count int
		char  byte
	}
	countM := make([]cnt, 3)
	countM[0], countM[1], countM[2] = cnt{a, 'a'}, cnt{b, 'b'}, cnt{c, 'c'}

	var res []byte

	for {
		sort.Slice(countM, func(i, j int) bool {
			return countM[i].count > countM[j].count
		})
		hasNext := false
		for idx, this := range countM {
			if this.count <= 0 {
				break
			}
			m := len(res)
			if m >= 2 && res[m-2] == this.char && res[m-1] == this.char {
				// if last two char is same as this one, skip this one
				continue
			}
			// this step proof we have a char which could be appended
			hasNext = true
			res = append(res, this.char)
			countM[idx].count -= 1
			break
		}
		if !hasNext {
			break
		}
	}

	return string(res)
}
