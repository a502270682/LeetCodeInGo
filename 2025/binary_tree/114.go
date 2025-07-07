package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var helper func(root *TreeNode)
	var list []*TreeNode
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	//list := preOrderTravel(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left = nil
		prev.Right = curr
	}
}

func preOrderTravel(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preOrderTravel(root.Left)...)
		list = append(list, preOrderTravel(root.Right)...)
	}
	return list
}
