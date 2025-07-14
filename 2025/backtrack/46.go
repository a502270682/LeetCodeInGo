package backtrack

func permute(nums []int) [][]int {
	length := len(nums)
	used := make([]bool, length)
	var res [][]int

	dfs(nums, []int{}, used, length, &res)
	return res
}

func dfs(nums, path []int, used []bool, length int, res *[][]int) {
	if len(path) == length {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < length; i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, path, used, length, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
