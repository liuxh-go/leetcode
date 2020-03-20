package simple

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode

	cur := head
	for cur != nil {
		next, cur.Next = cur.Next, pre
		pre, cur = cur, next
	}

	return pre
}
