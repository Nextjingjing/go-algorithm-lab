package hashmap

// ContainsDuplicate reports whether nums contains any repeated value.
func ContainsDuplicate(nums []int) bool {
	counts := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if counts[nums[i]] == 0 {
			counts[nums[i]] = 1
		} else {
			counts[nums[i]] = counts[nums[i]] + 1
		}

		if counts[nums[i]] > 1 {
			return true
		}
	}
	return false
}
