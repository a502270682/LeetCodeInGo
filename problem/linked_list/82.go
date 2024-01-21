package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	cur := head
//	numM := make(map[int]int)
//	for cur.Next != nil {
//		if count, ok := numM[cur.Val]; ok {
//			numM[cur.Val] = count + 1
//		} else {
//			numM[cur.Val] = 1
//		}
//		cur = cur.Next
//	}
//	cur = head.Next
//	last := cur
//	for cur != nil {
//		if count, _ := numM[cur.Val]; count > 1 {
//			tmp := cur.Next
//			last.Next = tmp
//		} else {
//
//		}
//		cur = cur.Next
//
//	}
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
		} else {
			tmp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		}
	}
	return dummy.Next
}
