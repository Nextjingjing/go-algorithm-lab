package stack_test

import (
	"testing"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "matching brackets",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "nested brackets",
			s:    "{[]}",
			want: true,
		},
		{
			name: "mismatched closing bracket",
			s:    "(]",
			want: false,
		},
		{
			name: "wrong nesting order",
			s:    "([)]",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "unclosed bracket",
			s:    "(",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stack.IsValid(tt.s)

			if got != tt.want {
				t.Fatalf("IsValid(%q) = %t, want %t", tt.s, got, tt.want)
			}
		})
	}
}
