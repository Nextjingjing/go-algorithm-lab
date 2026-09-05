package linkedlist_test

import (
	"testing"

	linkedlist "github.com/Nextjingjing/go-algorithm-lab/algorithms/linked-list"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		cycleStart int
		want       bool
	}{
		{
			name:       "cycle returns to a middle node",
			values:     []int{3, 2, 0, -4},
			cycleStart: 1,
			want:       true,
		},
		{
			name:       "single node points to itself",
			values:     []int{1},
			cycleStart: 0,
			want:       true,
		},
		{
			name:       "ordinary list has no cycle",
			values:     []int{1, 2},
			cycleStart: -1,
			want:       false,
		},
		{
			name:       "empty list has no cycle",
			values:     nil,
			cycleStart: -1,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linkedlist.HasCycle(cycleListFromValues(tt.values, tt.cycleStart)); got != tt.want {
				t.Fatalf("HasCycle() = %t, want %t", got, tt.want)
			}
		})
	}
}

func cycleListFromValues(values []int, cycleStart int) *linkedlist.ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*linkedlist.ListNode, len(values))
	for i, value := range values {
		nodes[i] = &linkedlist.ListNode{Val: value}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if cycleStart >= 0 {
		nodes[len(nodes)-1].Next = nodes[cycleStart]
	}

	return nodes[0]
}
