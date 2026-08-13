package hashmap

// IsAnagram reports whether t is an anagram of s.
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	x1 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		x1[s[i]]++
	}

	x2 := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		x2[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if x1[s[i]] != x2[s[i]] {
			return false
		}
	}
	return true
}
