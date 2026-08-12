package bruteforce

// SumArray returns the sum of all values in nums.
func SumArray(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	return count
}
