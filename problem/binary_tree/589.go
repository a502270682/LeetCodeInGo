package binary_tree

type NNode struct {
	Val      int
	Children []*NNode
}

func preorder(root *NNode) []int {
	if root == nil {
		return nil
	}
	return preHelper(root, []int{})
}

func preHelper(root *NNode, arr []int) []int {
	arr = append(arr, root.Val)
	for _, child := range root.Children {
		arr = preHelper(child, arr)
	}
	return arr
}
