package bruteforce

// CountOccurrences returns how many times target appears in nums.
//
// nums is the slice being searched.
// target is the value to count.
func CountOccurrences(nums []int, target int) int {
	// count stores how many matching values have been found so far.
	count := 0

	// i walks through nums from the first index to the last index.
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}

	return count
}
