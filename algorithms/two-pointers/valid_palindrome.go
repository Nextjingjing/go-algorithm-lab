package twopointers

import (
	"strings"
	"unicode"
)

// IsPalindrome reports whether s is a palindrome after ignoring
// non-alphanumeric characters and letter case.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	t := ""
	for i := 0; i < len(s); i++ {
		ch := rune(s[i])

		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			continue
		}
		t += string(ch)
	}

	left := 0
	right := len(t) - 1
	for left < right {
		if t[left] != t[right] {
			return false
		}
		left++
		right--
	}
	return true
}
