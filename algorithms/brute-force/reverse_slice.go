package bruteforce

// ReverseSlice reverses nums in place.
//
// nums is the slice whose values should be reversed.
func ReverseSlice(nums []int) {
	// left starts at the first index.
	left := 0

	// right starts at the last index.
	right := len(nums) - 1

	for left < right {
		// temp stores nums[left] while the two values are swapped.
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp

		left++
		right--
	}
}
