package _025

func lexicalOrder(n int) []int {
	res := make([]int, 0)
	var dfs func(curValue, maxValue int)
	dfs = func(curValue, maxValue int) {
		if curValue > maxValue {
			return
		}
		res = append(res, curValue)
		for i := 0; i <= 9; i++ {
			dfs(curValue*10+i, maxValue)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}

	return res
}
