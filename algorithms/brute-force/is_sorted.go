package bruteforce

// IsSorted reports whether nums is sorted in ascending order.
//
// nums is the slice being checked.
func IsSorted(nums []int) bool {
	// i compares each value with the next value after it.
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}
