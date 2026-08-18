package slidingwindow

// LengthOfLongestSubstring returns the length of the longest substring that
// contains no repeated characters. The input string is not modified.
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	maxS := 0
	nowS := 0
	right := 0
	counts := map[byte]int{}

	for right < len(s) {
		for counts[s[right]]+1 > 1 { // Repeating Detected !!!
			counts[s[left]]--
			left++
			nowS--
		}
		counts[s[right]]++
		right++
		nowS++
		maxS = max(maxS, nowS)
	}
	return maxS
}
