package binary_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var helper func(root *TreeNode, maxInt, minInt int) bool
	helper = func(root *TreeNode, maxInt, minInt int) bool {
		if root == nil {
			return true
		}
		leftRes := helper(root.Left, root.Val, minInt)
		rightRes := helper(root.Right, maxInt, root.Val)
		return root.Val < maxInt && root.Val > minInt && leftRes && rightRes
	}
	return helper(root, math.MaxInt64, math.MinInt64)
}
