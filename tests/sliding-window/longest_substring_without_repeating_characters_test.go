package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "repeated characters", s: "abcabcbb", want: 3},
		{name: "repeated single character", s: "bbbbb", want: 1},
		{name: "repeat after unique prefix", s: "pwwkew", want: 3},
		{name: "empty string", s: "", want: 0},
		{name: "all characters are unique", s: "abcdef", want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slidingwindow.LengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
