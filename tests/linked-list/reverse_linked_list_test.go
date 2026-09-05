package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "reverses several nodes",
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
		{
			name:  "keeps one node unchanged",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "handles an empty list",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linkedlist.ReverseList(listFromValues(tt.input))
			if values := valuesFromList(got); !reflect.DeepEqual(values, tt.want) {
				t.Fatalf("ReverseList() = %v, want %v", values, tt.want)
			}
		})
	}
}

func listFromValues(values []int) *linkedlist.ListNode {
	var head *linkedlist.ListNode
	var tail *linkedlist.ListNode

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

func valuesFromList(head *linkedlist.ListNode) []int {
	var values []int
	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}
	return values
}
