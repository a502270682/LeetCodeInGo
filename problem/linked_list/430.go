package linked_list

/*
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func dfs(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		// 如果有子节点，那么首先处理子节点
		if cur.Child != nil {
			childLast := dfs(cur.Child)

			next = cur.Next
			// 将 node 与 child 相连
			cur.Next = cur.Child
			cur.Child.Prev = cur

			// 如果 next 不为空，就将 last 与 next 相连
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			// 将 child 置为空
			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}
