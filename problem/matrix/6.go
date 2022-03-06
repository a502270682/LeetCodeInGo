package matrix

import "strings"

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。

比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

 */

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	ret := make([]string, numRows)
	i, flag := 0, -1
	for _, this := range s {
		ret[i] += string(this)
		// 先从上到下，到转折点返回来从下到上，依次类推
		// 转向点在第0行和最后一行
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	return strings.Join(ret, "")
}