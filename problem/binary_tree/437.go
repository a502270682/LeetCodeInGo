package binary_tree

/*
给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum2(root *TreeNode, targetSum int) int {
	preSum := map[int]int{0: 1} // 记录当前有的前缀和
	var dfs func(*TreeNode, int)
	var ans int
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}

		curr += node.Val
		if count, ok := preSum[curr-targetSum]; ok && count > 0 {
			// 按当前 前缀和 - targetSum 如果有等于这么多的某个前缀和，则结果叠加
			ans += count
		}
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return ans
}
