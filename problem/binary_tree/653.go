package binary_tree

/*
给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := make(map[int]int)
	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		m[root.Val] = 1
		if _, ok := m[k-root.Val]; ok && k != root.Val*2 {
			return true
		}
		var leftRes, rightRes bool
		if root.Left != nil {
			leftRes = helper(root.Left)
		}
		if root.Right != nil {
			rightRes = helper(root.Right)
		}
		return leftRes || rightRes
	}
	return helper(root)
}
