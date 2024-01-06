package linked_list

/*
给你一个链表的头 head ，每个结点包含一个整数值。

在相邻结点之间，请你插入一个新的结点，结点值为这两个相邻结点值的 最大公约数 。

请你返回插入之后的链表。

两个数的 最大公约数 是可以被两个数字整除的最大正整数。
*/

func GCD(a, b int) int {
	if a == b {
		return a
	}
	getGCD := func(big, small int) int {
		for big%small != 0 {
			newBig := small
			small = big % small
			big = newBig
		}
		return small
	}
	if a > b {
		return getGCD(a, b)
	} else {
		return getGCD(b, a)
	}
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = insertGreatestCommonDivisors(head.Next)
	gcdNode := &ListNode{
		Val:  GCD(head.Val, head.Next.Val),
		Next: head.Next,
	}
	head.Next = gcdNode
	return head
}
