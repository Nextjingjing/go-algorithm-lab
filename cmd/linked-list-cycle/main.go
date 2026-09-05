package main

import (
	"fmt"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func main() {
	head := &linkedlist.ListNode{Val: 1}
	head.Next = &linkedlist.ListNode{Val: 2}
	head.Next.Next = head

	fmt.Println("has cycle:", linkedlist.HasCycle(head))
}
