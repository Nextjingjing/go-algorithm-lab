package slidingwindow

// MaxConsecutiveOnes returns the length of the longest contiguous run of 1s.
// The input is expected to contain only 0 and 1 values and is not modified.
func MaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0
	maxCon1 := 0
	nowCon1 := 0
	for right < len(nums) {
		if nums[right] == 0 {
			left = right + 1
			right = left
			nowCon1 = 0
		} else {
			nowCon1++
			maxCon1 = max(maxCon1, nowCon1)
			right++
		}
	}
	return maxCon1
}
