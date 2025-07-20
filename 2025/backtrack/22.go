package backtrack

func generateParenthesis(n int) []string {
	m := n * 2
	ans := []string{}
	path := ""
	dfs := func(open int) {}
	dfs = func(open int) {
		if m == len(path) {
			ans = append(ans, path)
			return
		}
		if open < n {
			path += "("
			dfs(open + 1)
			path = path[:len(path)-1]
		}
		// 右括号的数量还可以增加
		if len(path)-open < open {
			path += ")"
			dfs(open)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
