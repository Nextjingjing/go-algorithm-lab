package linkedlist

// MergeTwoLists merges two nondecreasing linked lists and returns the head of
// the merged nondecreasing list. Either input may be nil.
func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	current1 := list1
	current2 := list2
	var newHead *ListNode

	if current1.Val < current2.Val {
		newHead = &ListNode{Val: current1.Val}
		current1 = current1.Next
	} else {
		newHead = &ListNode{Val: current2.Val}
		current2 = current2.Next
	}

	currentNew := newHead

	for current1 != nil && current2 != nil {
		var newNode *ListNode

		if current1.Val <= current2.Val {
			newNode = &ListNode{Val: current1.Val}
			currentNew.Next = newNode
			currentNew = currentNew.Next
			current1 = current1.Next
		} else {
			newNode = &ListNode{Val: current2.Val}
			currentNew.Next = newNode
			currentNew = currentNew.Next
			current2 = current2.Next
		}
	}

	if current1 != nil {
		currentNew.Next = current1
	} else {
		currentNew.Next = current2
	}

	return newHead
}
