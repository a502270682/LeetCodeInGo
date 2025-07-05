package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	var helper func(root *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil || k == 0 {
			return
		}
		helper(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		helper(node.Right)
	}
	helper(root)
	return ans
}
