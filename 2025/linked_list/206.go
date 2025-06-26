package linked_list

func reverseList(head *ListNode) *ListNode {
	var preLinked *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = preLinked
		preLinked = curr
		curr = next
	}
	return preLinked
}
