package binary_tree

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func dfs(root *TreeNode, mx, mi int) int {
	if root == nil {
		return 0
	}
	diff := max(abs(root.Val-mx), abs(root.Val-mi))
	mx, mi = max(mx, root.Val), min(mi, root.Val)
	diff = max(dfs(root.Left, mx, mi), diff)
	diff = max(dfs(root.Right, mx, mi), diff)
	return diff
}
