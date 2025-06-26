package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 翻转中间节点后的链表部分
	var preLinked *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = preLinked
		preLinked = slow
		slow = next
	}
	for preLinked != nil {
		if preLinked.Val != head.Val {
			return false
		}
		preLinked = preLinked.Next
		head = head.Next
	}
	return true
}
