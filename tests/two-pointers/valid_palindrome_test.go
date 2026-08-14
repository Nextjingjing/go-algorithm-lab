package twopointers_test

import (
	"testing"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "palindrome with spaces and punctuation",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "not a palindrome",
			s:    "race a car",
			want: false,
		},
		{
			name: "only spaces",
			s:    " ",
			want: true,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "mixed case",
			s:    "Aa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twopointers.IsPalindrome(tt.s)

			if got != tt.want {
				t.Fatalf("IsPalindrome(%q) = %t, want %t", tt.s, got, tt.want)
			}
		})
	}
}
