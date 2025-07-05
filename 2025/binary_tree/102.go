package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelNodeCount := len(queue)
		levelVal := make([]int, 0)
		for i := 0; i < levelNodeCount; i++ {
			node := queue[0]
			queue = queue[1:]
			levelVal = append(levelVal, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, levelVal)
	}
	return ret
}
