package simple

import "testing"

func TestReverseList(t *testing.T) {
	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	listNode := &ListNode{
		Val:  0,
		Next: node1,
	}

	result := reverseList(listNode)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
