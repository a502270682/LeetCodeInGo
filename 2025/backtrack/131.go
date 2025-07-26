package backtrack

import "slices"

func partition(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	var dfs func(str string)
	dfs = func(s string) {
		n := len(s)
		if n == 0 { // s 分割完毕
			ans = append(ans, slices.Clone(path))
			return
		}
		for i := 1; i <= n; i++ { // 枚举子串长度
			substr := s[:i]
			if isPalindrome(substr) {
				path = append(path, substr) // 分割！
				// 考虑剩余的 s[i:] 怎么分割
				dfs(s[i:])
				path = path[:len(path)-1] // 恢复现场
			}
		}

	}
	dfs(s)
	return ans
}

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
