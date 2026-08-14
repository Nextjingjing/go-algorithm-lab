package twopointers

// RemoveElement removes all occurrences of val from nums in place and
// returns the number of remaining elements.
//
// After the function returns, nums[:k] contains values different from val.
// The values after nums[:k] are unspecified.
func RemoveElement(nums []int, val int) int {
	w := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == val {
			continue
		}
		nums[w] = nums[r]
		w++
	}

	return w
}
