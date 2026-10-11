package divideconquer

// FindMin returns the minimum value in nums, not its index.
//
// nums must be non-empty and contain distinct integers sorted in ascending
// order then rotated. A full rotation (ascending order) is valid.
// FindMin does not modify nums. The exercise requires O(log n) time.
func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[left] > nums[len(nums)-1] && nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else if nums[right] <= nums[len(nums)-1] && nums[mid] <= nums[len(nums)-1] {
			right = mid
		}
	}
	return nums[left]
}
