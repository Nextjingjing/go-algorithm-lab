package linkedlist

// MiddleNode returns the middle node of the list beginning at head. For an
// even-length list, it returns the second of the two middle nodes. A nil head
// returns nil.
func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	currentNode := head
	nodes := make([]*ListNode, 0)
	for currentNode != nil {
		nodes = append(nodes, currentNode)
		currentNode = currentNode.Next
	}
	mid := len(nodes) / 2
	return nodes[mid]
}
