package backtrack

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(cur int)
	var path []int
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		path = append(path, nums[cur])
		dfs(cur + 1)
		path = path[:len(path)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res
}
