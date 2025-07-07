package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	var helper func(preorder []int, inorder []int) *TreeNode
	helper = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		rootV := preorder[0]
		tmp := &TreeNode{
			Val:   rootV,
			Left:  nil,
			Right: nil,
		}
		index := 0
		for ; index < len(inorder); index++ {
			if inorder[index] == rootV {
				break
			}
		}
		preCut := len(inorder[:index]) + 1
		tmp.Left = helper(preorder[1:preCut], inorder[:index])
		tmp.Right = helper(preorder[preCut:], inorder[index+1:])
		return tmp
	}
	return helper(preorder, inorder)
}
