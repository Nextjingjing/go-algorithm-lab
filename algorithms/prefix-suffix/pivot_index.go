package prefixsuffix

// PivotIndex returns the leftmost index where the sum of the values to the
// left equals the sum of the values to the right. It returns -1 when no such
// index exists.
func PivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum := total - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
