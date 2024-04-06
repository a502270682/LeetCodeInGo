package binary_tree

func allPossibleFBT(n int) []*TreeNode {
	var fullBinaryTrees []*TreeNode
	if n%2 == 0 {
		return fullBinaryTrees
	}
	if n == 1 {
		fullBinaryTrees = append(fullBinaryTrees, &TreeNode{Val: 0})
		return fullBinaryTrees
	}
	for i := 1; i <= n-1; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - 1 - i)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{
					Val:   0,
					Left:  l,
					Right: r,
				}
				fullBinaryTrees = append(fullBinaryTrees, root)
			}
		}
	}
	return fullBinaryTrees
}
