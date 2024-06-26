package stack

/*
有效括号字符串为空 ""、"(" + A + ")" 或 A + B ，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。

例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
如果有效字符串 s 非空，且不存在将其拆分为 s = A + B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。

给出一个非空有效字符串 s，考虑将其进行原语化分解，使得：s = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。

对 s 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 s 。



示例 1：

输入：s = "(()())(())"
输出："()()()"
解释：
输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。

*/

func removeOuterParentheses(s string) string {
	var ans, st []rune
	for _, c := range s {
		if c == ')' {
			st = st[:len(st)-1]
		}
		// 如果stack一直有一个左括号，说明一个原语还没有结束，应该把当前的c计入ans中
		// 当stack为空后，又开始下一轮原语收集，注意最外层的是不会被收集到的，因为到最后一个右括号被找到时，len(stack) = 0了
		if len(st) > 0 {
			ans = append(ans, c)
		}
		if c == '(' {
			st = append(st, c)
		}
	}
	return string(ans)
}
