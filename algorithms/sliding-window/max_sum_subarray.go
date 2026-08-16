package slidingwindow

// MaxSumSubarray returns the largest sum of any contiguous subarray of size k.
// It returns 0 when k is not positive or is larger than len(nums).
// The input is not modified.
func MaxSumSubarray(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}
	left := 0
	right := k - 1
	nowSum := 0
	for i := 0; i <= right; i++ {
		nowSum += nums[i]
	}
	maxSum := nowSum
	for right < len(nums)-1 {
		right++
		nowSum += nums[right]
		nowSum -= nums[left]
		left++
		maxSum = max(nowSum, maxSum)
	}

	return maxSum
}
