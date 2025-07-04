package datastructure

type ListNode struct {
	Val  int
	Next *ListNode
}

// single linked list
type LinkedList struct {
	Head *ListNode
}

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next // save next node for current
		curr.Next = prev  // reverse next node for current
		prev = curr       // move prev to current
		curr = next       // move current to next
	}
	return prev
}
