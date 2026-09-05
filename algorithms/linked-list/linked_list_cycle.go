package linkedlist

// HasCycle reports whether the singly linked list beginning at head contains a
// cycle. A nil head does not contain a cycle.
func HasCycle(head *ListNode) bool {
	addresses := make(map[*ListNode]bool)

	currentNode := head
	addresses[head] = true
	for currentNode != nil {
		currentNode = currentNode.Next
		if addresses[currentNode] {
			return true
		}
		addresses[currentNode] = true
	}

	return false
}
