package linked_list

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		node1.Next = node2.Next
		node2.Next = node1
		cur.Next = node2
		cur = node1
	}
	return dummy.Next
}
