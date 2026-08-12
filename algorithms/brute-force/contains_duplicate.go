package bruteforce

// ContainsDuplicate reports whether nums contains any repeated value.
//
// nums is the slice being checked.
func ContainsDuplicate(nums []int) bool {
	// i chooses the first value in a pair.
	for i := 0; i < len(nums); i++ {
		// j chooses the second value after i, so the same pair is not checked twice.
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
