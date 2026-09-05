package linkedlist_test

import (
	"testing"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   *int
	}{
		{name: "odd length", values: []int{1, 2, 3, 4, 5}, want: intPointer(3)},
		{name: "even length returns second middle", values: []int{1, 2, 3, 4, 5, 6}, want: intPointer(4)},
		{name: "single node", values: []int{9}, want: intPointer(9)},
		{name: "empty list", values: nil, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linkedlist.MiddleNode(middleListFromValues(tt.values))
			if tt.want == nil {
				if got != nil {
					t.Fatalf("MiddleNode() = %d, want nil", got.Val)
				}
				return
			}
			if got == nil || got.Val != *tt.want {
				t.Fatalf("MiddleNode() = %v, want %d", got, *tt.want)
			}
		})
	}
}

func middleListFromValues(values []int) *linkedlist.ListNode {
	var head, tail *linkedlist.ListNode
	for _, value := range values {
		node := &linkedlist.ListNode{Val: value}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func intPointer(value int) *int { return &value }
