package main

import (
	"fmt"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func main() {
	head := &linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 2,
			Next: &linkedlist.ListNode{
				Val: 3,
			},
		},
	}

	reversed := linkedlist.ReverseList(head)
	fmt.Println("new head:", reversed.Val)
}
