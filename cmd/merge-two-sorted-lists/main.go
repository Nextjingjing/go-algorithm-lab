package main

import (
	"fmt"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func main() {
	first := &linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 4,
		},
	}
	second := &linkedlist.ListNode{
		Val: 2,
		Next: &linkedlist.ListNode{
			Val: 3,
		},
	}

	merged := linkedlist.MergeTwoLists(first, second)
	fmt.Println("merged head:", merged.Val)
}
