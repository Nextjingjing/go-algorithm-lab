package hashmap_test

import (
	"testing"

	hashmap "github.com/Nextjingjing/go-algorithm-lab/algorithms/hash-map"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "same letters in different order",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "different letters",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "same letters with repeated values",
			s:    "aabb",
			t:    "abab",
			want: true,
		},
		{
			name: "different repeated counts",
			s:    "aabb",
			t:    "abbb",
			want: false,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			name: "both empty",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "case-sensitive characters",
			s:    "OOOOo",
			t:    "OOOOO",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashmap.IsAnagram(tt.s, tt.t)

			if got != tt.want {
				t.Fatalf("IsAnagram(%q, %q) = %t, want %t", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
