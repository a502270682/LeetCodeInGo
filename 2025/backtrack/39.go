package backtrack

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	comb := []int{}
	var dfs func(index, target int)
	dfs = func(index, target int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 不选取当前位置，直接跳过
		dfs(index+1, target)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			// 选取当前位置
			dfs(index, target-candidates[index])
			// 回溯当前位置
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, target)
	return ans
}
