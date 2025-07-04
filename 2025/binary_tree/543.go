package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var helper func(root *TreeNode) int
	maxNode := 1
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		maxNode = max(left+right+1, maxNode)
		return max(left, right) + 1
	}
	helper(root)
	return maxNode - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
