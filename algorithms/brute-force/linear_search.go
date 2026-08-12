package bruteforce

// LinearSearch returns the index of target in nums.
//
// nums is the slice being searched.
// target is the value to find.
// It returns -1 when target is not found.
func LinearSearch(nums []int, target int) int {
	// i walks through nums from the first index to the last index.
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
