package linkedlist

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses the links in the list beginning at head and returns its
// new head. A nil head returns nil.
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	stack := make([]*ListNode, 0)

	currentNode := head
	for currentNode != nil {
		stack = append(stack, currentNode)
		currentNode = currentNode.Next
		stack[len(stack)-1].Next = nil
	}

	newHead := stack[len(stack)-1]
	currentNode = newHead
	stack = stack[:len(stack)-1]
	for len(stack) > 0 {
		currentNode.Next = stack[len(stack)-1]
		currentNode = currentNode.Next
		stack = stack[:len(stack)-1]
	}
	return newHead
}
