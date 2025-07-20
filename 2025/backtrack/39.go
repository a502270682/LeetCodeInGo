package backtrack

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	dfs := func(idx, target int) {}
	dfs = func(idx, target int) {
		if idx == len(candidates) || target < 0 {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(idx+1, target)
		path = append(path, candidates[idx])
		dfs(idx, target-candidates[idx])
		path = path[:len(path)-1]
	}
	return ans
}
