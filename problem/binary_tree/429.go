package binary_tree

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。


*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node2 struct {
	Val      int
	Children []*Node2
}

func levelOrder(root *Node2) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	var helper func(root []*Node2)
	helper = func(root []*Node2) {
		if len(root) == 0 {
			return
		}
		next := make([]*Node2, 0)
		thisLevel := make([]int, 0)
		for _, this := range root {
			thisLevel = append(thisLevel, this.Val)
			if len(this.Children) > 0 {
				next = append(next, this.Children...)
			}
		}
		helper(next)
		ret = append(ret, thisLevel)
	}
	helper([]*Node2{root})
	return ret
}
