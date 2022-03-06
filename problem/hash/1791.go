package hash

/*
有一个无向的 星型 图，由 n 个编号从 1 到 n 的节点组成。星型图有一个 中心 节点，并且恰有 n - 1 条边将中心节点与其他每个节点连接起来。

给你一个二维整数数组 edges ，其中edges[i] = [ui, vi] 表示在节点 ui 和 vi 之间存在一条边。请你找出并返回edges 所表示星型图的中心节点。

输入：edges = [[1,2],[5,1],[1,3],[1,4]]
输出：1
*/

func findCenter(edges [][]int) int {
	// 中心节点的度一定是n-1
	n := len(edges) + 1
	degrees := make([]int, n+1)
	for _, e := range edges {
		degrees[e[0]]++
		degrees[e[1]]++
	}
	for i, d := range degrees {
		if d == n-1 {
			return i
		}
	}
	return -1
}
