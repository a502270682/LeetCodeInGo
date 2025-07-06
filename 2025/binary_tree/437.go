package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	sum := thisRootSum(root, targetSum)
	sum += pathSum(root.Left, targetSum)
	sum += pathSum(root.Right, targetSum)
	return sum
}

func thisRootSum(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	sum := 0
	val := root.Val
	if val == target {
		sum++
	}
	sum += thisRootSum(root.Left, target-val)
	sum += thisRootSum(root.Right, target-val)
	return sum
}
