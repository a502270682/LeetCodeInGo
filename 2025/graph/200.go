package graph

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		// 出界，或者不是 '1'，就不再往下递归
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2' // 插旗！避免来回横跳无限递归
		dfs(i, j-1)      // 往左走
		dfs(i, j+1)      // 往右走
		dfs(i-1, j)      // 往上走
		dfs(i+1, j)      // 往下走
	}

	for i, row := range grid {
		for j, c := range row {
			if c == '1' { // 找到了一个新的岛
				dfs(i, j) // 把这个岛插满旗子，这样后面遍历到的 '1' 一定是新的岛
				ans++
			}
		}
	}
	return
}
