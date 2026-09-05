package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{
			name:   "interleaves two lists",
			first:  []int{1, 2, 4},
			second: []int{1, 3, 4},
			want:   []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:   "accepts an empty first list",
			first:  nil,
			second: []int{0},
			want:   []int{0},
		},
		{
			name:   "accepts two empty lists",
			first:  nil,
			second: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linkedlist.MergeTwoLists(mergeListFromValues(tt.first), mergeListFromValues(tt.second))
			if values := mergeValuesFromList(got); !reflect.DeepEqual(values, tt.want) {
				t.Fatalf("MergeTwoLists() = %v, want %v", values, tt.want)
			}
		})
	}
}

func mergeListFromValues(values []int) *linkedlist.ListNode {
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

func mergeValuesFromList(head *linkedlist.ListNode) []int {
	var values []int
	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}
	return values
}
